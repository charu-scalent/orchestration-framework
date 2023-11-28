package main

import "log"

func main() {
	err := initConfig()
	if err != nil {
		log.Fatal(err)
	}

	registry, err := initServer(&config)
	if err != nil {
		log.Fatal(err)
	}

	err = registry.StartServer()
	if err != nil {
		log.Fatal(err)
	}
}
