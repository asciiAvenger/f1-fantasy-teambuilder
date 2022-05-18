package builder

import (
	"f1-fantasy-teambuilder/models"
	"fmt"
	"sort"
)

func BuildTeam(data models.Data, budget float64, topTeams int) []models.FantasyTeam {
	fmt.Println("Creating permutations...")
	permutations := CreatePermutations(data, budget)
	fmt.Println("Sorting...")
	sortPermutations(permutations)
	return permutations
}

func sortPermutations(permutations []models.FantasyTeam) {
	sort.Slice(permutations, func(i, j int) bool {
		return permutations[i].GetTeamSeasonScore() > permutations[j].GetTeamSeasonScore()
	})
}
