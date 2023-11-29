package config

import "log"

func Init() {
	err := LoadConfig()

	if err != nil {
		log.Fatalf("could not load env file : %v", err)
	}
}
