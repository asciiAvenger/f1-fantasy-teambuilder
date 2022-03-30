package reader

import (
	"f1-fantasy-teambuilder/models"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

func ReadData(src string) models.Data {
	yfile, err := ioutil.ReadFile(src)

	if err != nil {
		log.Fatal(err)
	}

	var data models.Data

	err = yaml.Unmarshal(yfile, &data)

	if err != nil {
		log.Fatal(err)
	}

	return data
}
