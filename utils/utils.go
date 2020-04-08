package utils

import (
	"log"
	"os/user"
)

func HomeDir() string {
	user, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return user.HomeDir
}
