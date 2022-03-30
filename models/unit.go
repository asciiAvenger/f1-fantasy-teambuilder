package models

// driver or team
type Unit struct {
	Name    string    `yaml:"name"`
	Results []float64 `yaml:"results"`
	Cost    float64   `yaml:"cost"`
}

func (u Unit) GetScore() float64 {
	sum := 0.0
	for _, r := range u.Results {
		sum += r
	}
	avg := sum / float64(len(u.Results))
	return avg / u.Cost
}

func (u Unit) GetPoints() float64 {
	sum := 0.0
	for _, r := range u.Results {
		sum += r
	}
	return sum
}
