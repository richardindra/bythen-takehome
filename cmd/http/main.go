package main

import (
	"bythen-takehome/internal/boot"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	time.Sleep(30 * time.Second)
	if err := boot.HTTP(); err != nil {
		log.Println("[HTTP] failed to boot http server due to " + err.Error())
	}
}
