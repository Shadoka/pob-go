package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/Shadoka/pob-go/model"
)

func main() {
	content, err := ioutil.ReadFile("mock_data/character_data.json")
	if err != nil {
		log.Fatal("error opening file: ", err)
	}

	var character_data model.CharacterData
	err = json.Unmarshal(content, &character_data)
	if err != nil {
		log.Fatal("error unmarshalling server response: ", err)
	}

	log.Printf("Imported character has %v jewels socketed", len(character_data.Items))
}
