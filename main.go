package main

// ./app | logmgr
// open http://localhost:19271/

// go run main.go open.go
// go build && echo "hello testing 123" | ./logmgr

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type logEntry struct {
	Line      string    `json:"line"`
	Timestamp time.Time `json:"timestamp"`
}

const port = 18192

func main() {
	broker := NewBroker()

	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			test := scanner.Text()
			le := logEntry{Line: test, Timestamp: time.Now()}

			output, err := json.Marshal(le)
			if err != nil {
				continue // TODO: What's better?
			}

			broker.Notifier <- []byte(output)
		}
	}()

	//Open(fmt.Sprintf("http://localhost:%d", port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), broker))
}
