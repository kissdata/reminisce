package main

import (
	"log"
)

func main() {

	if err := connPG14(); err != nil {
		log.Println("fail to connect to pg14, stop running! err=", err)
		return
	}

	test221009()
}
