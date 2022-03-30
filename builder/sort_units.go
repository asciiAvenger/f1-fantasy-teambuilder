package builder

import (
	"f1-fantasy-teambuilder/models"
	"sort"
)

func sortUnitsDesc(units []models.Unit) {
	sort.Slice(units, func(i, j int) bool {
		return units[i].GetScore() > units[j].GetScore()
	})
}
