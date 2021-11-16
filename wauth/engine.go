package wauth

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/deeper-x/weblog/settings"
)

// IsAllowed check if signature is allowed
func IsAllowed(signature string) (bool, error) {
	var res bool

	entries, err := getJSONEntries()
	if err != nil {
		log.Println(err)
		return res, err
	}

	for _, entry := range entries {
		if entry.ID == signature {
			res = true
			break
		}
	}

	return res, nil
}

// getJSONEntries reads the json file
func getJSONEntries() ([]System, error) {
	var res = []System{}
	var list Whitelist

	jsonFile, err := os.Open(settings.GetAuthFile())
	if err != nil {
		log.Println(err)
		return res, err
	}

	defer jsonFile.Close()

	byteData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Println(err)
		return res, err
	}

	json.Unmarshal(byteData, &list)

	for _, system := range list.Systems {
		record := System{
			ID:          system.ID,
			Description: system.Description,
		}

		res = append(res, record)
	}

	return res, nil
}
