package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/heldeen/aoc2024/cmd"
)

func main() {

	date := time.Now().Format("20060102150405")
	logger, err := os.OpenFile(fmt.Sprintf("AoC.%s.log", date), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0660)
	if err != nil {
		log.Fatalln(err)
	}
	log.SetOutput(logger)

	if err := cmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}
