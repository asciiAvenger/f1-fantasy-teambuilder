package models

// structure of the data.yaml file
type Data struct {
	Players []Unit `json:"players"`
	Drivers []Unit
	Teams   []Unit
}
