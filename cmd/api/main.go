package main

import (
	boostrap "jagch/auth-go/cmd/api/bootstrap"
	"log"
)

func main() {
	if err := boostrap.Run(); err != nil {
		log.Fatal(err)
	}
}
