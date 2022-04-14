package models

type FantasyTeam struct {
	Drivers []Unit
	Team    Unit
}

func (f FantasyTeam) GetTeamSeasonScore() float64 {
	sum := f.Team.SeasonScore
	for _, d := range f.Drivers {
		sum += d.SeasonScore
	}
	return sum
}

func (f FantasyTeam) GetTeamPrice() float64 {
	sum := f.Team.Price
	for _, d := range f.Drivers {
		sum += d.Price
	}
	return sum
}
