package models

type FantasyTeam struct {
	Drivers []Unit
	Team    Unit
}

func (f FantasyTeam) GetTeamPoints() float64 {
	sum := f.Team.GetPoints()
	for _, d := range f.Drivers {
		sum += d.GetPoints()
	}
	return sum
}

func (f FantasyTeam) GetTeamCost() float64 {
	sum := f.Team.Cost
	for _, d := range f.Drivers {
		sum += d.Cost
	}
	return sum
}
