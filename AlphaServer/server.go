package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// map of hostname_date and details
var nodesInfo map[string]sshDetails

type sshDetails struct {
	Hostname    string `json:"hostname"`
	Date        string `json:"date"`
	FailedLogin int    `json:"failed_login"`
}

func main() {
	nodesInfo = make(map[string]sshDetails)

	http.HandleFunc("/ssh_details", handler)

	fmt.Println("Server started...")
	err := http.ListenAndServe(":9001", nil)
	if err != nil {
		fmt.Println(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		handlePost(w, r)
		return
	}

	if r.Method == "GET" {
		handleGet(w, r)
		return
	}

	w.WriteHeader(405)
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	var ssh sshDetails
	err = json.Unmarshal(body, &ssh)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	nodesInfo[ssh.Hostname+"_"+ssh.Date] = ssh
}

func handleGet(w http.ResponseWriter, r *http.Request) {
	data, err := json.Marshal(nodesInfo)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
