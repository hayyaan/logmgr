package main

// ./app | logmgr
// open http://localhost:19271/

// go run main.go open.go
// go build && echo "hello testing 123" | ./logmgr

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

// brainstorm struct
// type logInstance struct {
// 	name       string
// 	port       int
// 	instanceID string //some sort of unique ID for inter-OS stuff
// 	createdTime
// 	ownerProcessID
// 	ownerShell
// }

type logEntry struct {
	line      string
	timestamp time.Time
}

const port = 18192

func logEntryToOutput(le logEntry) string {
	return fmt.Sprintf("timestamp=%s, content=%s\n", le.timestamp, le.line)
}

func main() {
	fmt.Println("hello world")

	var logs []logEntry
	var logsMutex sync.Mutex

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "")
		//dont wait for channel
		logsMutex.Lock()
		for _, le := range logs {
			fmt.Fprint(w, logEntryToOutput(le))
		}
		logsMutex.Unlock()
	})

	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			test := scanner.Text()
			fmt.Println(test)
			logsMutex.Lock()
			logs = append(logs, logEntry{line: test, timestamp: time.Now()})
			logsMutex.Unlock()
		}
	}()

	Open("http://localhost:18192")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
