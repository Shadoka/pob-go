package data_import

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/Shadoka/pob-go/model"
)

func ImportCharacterPassives(accountName string, characterName string, realm string) (*model.CharacterData, error) {
	log.Printf("Importing from account:\t\t%v", accountName)
	log.Printf("Importing character with name:\t%v", characterName)
	log.Printf("Importing from realm:\t\t%v", realm)

	content, err := ioutil.ReadFile("C:\\go_workspaces/pob-go/resources/character_data.json")
	if err != nil {
		return nil, err
	}

	var character_data model.CharacterData
	err = json.Unmarshal(content, &character_data)
	if err != nil {
		return nil, err
	}

	return &character_data, nil
}
