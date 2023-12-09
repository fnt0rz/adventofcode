package utils

import (
	"log"
	"os"
	"time"
)

func Timer(name string) func() {
	startTime := time.Now()
	return func() {
		log.Printf("%s took %v\n", name, time.Since(startTime))
	}
}

func GetData(day string, test bool) string {
	path := "input/day" + day + "/"
	if test {
		path += "test.txt"
	} else {
		path += "input.txt"
	}

	res, err := os.ReadFile(path)

	if err != nil {
		log.Println("Error while loading data: ")
		log.Println(err)
	}

	return string(res)
}
