package main

import (
  "os/exec"
  "fmt"
  "time"
)

func main() {
  now := time.Now()
  regexp := getRegexpForDay(now)
  logFilePath := "/var/log/auth.log"

  cmd := exec.Command("grep", regexp, logFilePath)
  out, err := cmd.CombinedOutput()
  if err != nil {
    fmt.Println(err)
    return
  }

  lines := countLines(string(out))
  fmt.Println(lines)
}

func getRegexpForDay(t time.Time) string {
  dayMonthStr := t.Format("Jan 2")
  return fmt.Sprintf("%s.*Failed password", dayMonthStr)
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
