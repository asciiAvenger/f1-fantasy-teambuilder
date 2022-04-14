package models

// driver or team
type Unit struct {
	FirstName   string  `json:"first_name"`
	LastName    string  `json:"last_name"`
	SeasonScore float64 `json:"season_score"`
	Price       float64 `json:"price"`
	Position    string  `json:"position"`
}

func (u Unit) GetScore() float64 {
	// maybe rework the scoring later
	return u.SeasonScore / u.Price
}
