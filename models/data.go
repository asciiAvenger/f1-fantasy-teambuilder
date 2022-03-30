package models

// structure of the data.yaml file
type Data struct {
	Drivers []Unit `yaml:"drivers"`
	Teams   []Unit `yaml:"teams"`
}
