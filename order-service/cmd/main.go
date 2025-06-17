package main

import (
	"github.com/google/uuid"
	"log"
	"time"
)

func main() {
	log.Println(uuid.New().String())
	time.Sleep(time.Hour)
}
