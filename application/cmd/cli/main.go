package main

import (
	"fmt"
	poker "learn_go_with_test/application"
	"log"
	"os"
)

const dbFileName = "game.db.json"

func main() {
	store, close, err := poker.FileSystemPlayerStoreFromFile(dbFileName)

	if err != nil {
		log.Fatal(err)
	}
	defer close()

	fmt.Println("Lets play poker")
	fmt.Println("Type '{Name} wins' to record a win")
	poker.NewCLI(store, os.Stdin).PlayPoker()
}
