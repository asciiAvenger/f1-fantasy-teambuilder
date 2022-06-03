package ranker

import (
	"f1-fantasy-teambuilder/models"
	"fmt"
	"sort"
)

func RankDrivers(data models.Data, reverse bool) []models.Unit {
	fmt.Println("Ranking drivers...")
	drivers := data.Drivers
	if reverse {
		sortUnitsReverse(drivers)
	} else {
		sortUnits(drivers)
	}
	return drivers
}

func RankTeams(data models.Data, reverse bool) []models.Unit {
	fmt.Println("Ranking drivers...")
	teams := data.Teams
	if reverse {
		sortUnitsReverse(teams)
	} else {
		sortUnits(teams)
	}
	return teams
}

func sortUnits(units []models.Unit) {
	sort.Slice(units, func(i, j int) bool {
		return units[i].GetScore() > units[j].GetScore()
	})
}

func sortUnitsReverse(units []models.Unit) {
	sort.Slice(units, func(i, j int) bool {
		return units[i].GetScore() < units[j].GetScore()
	})
}
