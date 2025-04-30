package storage

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2/log"
	"os"
	"sync"
	"time"

	"github.com/edwingeng/deque/v2"
)

type Message struct {
	Content string    `json:"content"`
	Time    time.Time `json:"time"`
}

var messagesToWrite *deque.Deque[Message] = deque.NewDeque[Message]()
var fileMutex sync.Mutex = sync.Mutex{}
var DataFilename string

func write() {
	fileMutex.Lock()
	defer fileMutex.Unlock()

	if messagesToWrite.IsEmpty() {
		return
	}

	content, err := os.ReadFile(DataFilename)
	if err != nil {
		log.Fatal("Could not read the file: ", err)
	}
	var messages []Message
	err = json.Unmarshal(content, &messages)
	if err != nil {
		log.Fatal("Could not parse the file: ", err)
	}

	for !messagesToWrite.IsEmpty() {
		messages = append(messages, messagesToWrite.PopFront())
	}

	newcontent, err := json.Marshal(messages)
	if err != nil {
		log.Fatal("Could not parse the file: ", err)
	}

	err = os.WriteFile(DataFilename, newcontent, 0444)
	if err != nil {
		log.Fatal("Could not write the file: ", err)
	}
}

func WriteMessage(msg Message) {
	messagesToWrite.PushBack(msg)
	write()
}

func WriteMessages(msgs []Message) {
	for _, msg := range msgs {
		messagesToWrite.PushBack(msg)
	}
	go write()
}

func read() []Message {
	fileMutex.Lock()
	defer fileMutex.Unlock()

	content, err := os.ReadFile(DataFilename)
	if err != nil {
		log.Fatal("Could not read the file: ", err)
	}

	var messages []Message
	err = json.Unmarshal(content, &messages)
	if err != nil {
		log.Fatal("Could not parse the file: ", err)
	}

	return messages
}

func ReadMessages() []Message {
	return read()
}
