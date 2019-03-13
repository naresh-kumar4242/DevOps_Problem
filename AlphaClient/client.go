package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"time"
)

type sshDetails struct {
	Hostname    string `json:"hostname"`
	Date        string `json:"date"`
	FailedLogin int    `json:"failed_login"`
}

func main() {
	fmt.Println("Client Started...")

	handleErrorFunc := func(err error) {
		if err != nil {
			fmt.Println(err)
		}

		time.Sleep(2 * time.Minute)
	}

	for {
		now := time.Now()
		ssh, err := getSSHLoginDatails(now)
		if err != nil {
			handleErrorFunc(err)
			continue
		}

		err = sendDataToServer(ssh)
		if err != nil {
			handleErrorFunc(err)
			continue
		}

		handleErrorFunc(nil)
	}
}

func sendDataToServer(ssh *sshDetails) error {
	jsonData, err := json.Marshal(ssh)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", "http://localhost:9001/ssh_details", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return errors.New("Unable to send data.")
	}

	return nil
}

func getSSHLoginDatails(t time.Time) (*sshDetails, error) {
	dayMonthStr := t.Format("Jan 2")
	regexp := fmt.Sprintf("%s.*Failed password", dayMonthStr)
	count, err := getSSHLoginsForRegexp(regexp)
	if err != nil {
		return nil, err
	}

	hostname, err := os.Hostname()
	if err != nil {
		return nil, err
	}

	ssh := &sshDetails{
		Hostname:    hostname,
		Date:        dayMonthStr,
		FailedLogin: count,
	}

	return ssh, nil
}

func getSSHLoginsForRegexp(regexp string) (int, error) {
	logFilePath := "/var/log/auth.log"

	cmd := exec.Command("grep", regexp, logFilePath)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
		return 0, nil
	}

	lines := countLines(string(out))
	return lines, nil
}

func countLines(s string) int {
	count := 0
	for _, v := range s {
		if v == '\n' {
			count++
		}
	}

	return count
}
