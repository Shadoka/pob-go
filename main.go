package main

import (
	"log"

	"github.com/Shadoka/pob-go/data_import"
)

func main() {
	character_data, err := data_import.ImportCharacterPassives("TheShadoka", "CrashingShadoka", "pc")
	if err != nil {
		log.Fatal("Importing character data was unsuccessful: ", err)
	}

	log.Printf("Imported character has %v jewels socketed", len(character_data.Items))
}
