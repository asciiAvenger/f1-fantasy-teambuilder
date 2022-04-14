package reader

import (
	"encoding/json"
	"f1-fantasy-teambuilder/models"
	"io/ioutil"
	"log"
	"net/http"
)

// collect the data from the f1-fantasy public (unauthenticated) api
func CollectData() models.Data {
	// make the GET request
	// this includes drivers and teams
	resp, err := http.Get("https://fantasy-api.formula1.com/f1/2022/players")

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	var data models.Data

	// unmarshal into the data struct to filter out unnecessary fields
	// and allows for easier computation
	err = json.Unmarshal(body, &data)

	if err != nil {
		log.Fatal(err)
	}

	// sort players (teams and drivers) into separate slices inside the data
	for _, p := range data.Players {
		if p.Position == "Driver" {
			data.Drivers = append(data.Drivers, p)
		} else {
			data.Teams = append(data.Teams, p)
		}
	}

	return data
}
