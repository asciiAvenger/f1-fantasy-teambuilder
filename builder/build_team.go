package builder

import (
	"f1-fantasy-teambuilder/models"
	"fmt"
	"sort"
)

func BuildTeam(data models.Data, budget float64, topTeams int) {
	fmt.Println("Creating permutations...")
	permutations := CreatePermutations(data, budget)
	fmt.Println("Sorting...")
	sortPermutations(permutations)
	fmt.Printf("Top %d Fantasy Teams out of a total of %d possible teams:\n", topTeams, len(permutations))
	for i := 0; i < topTeams; i++ {
		fmt.Println(permutations[i])
		fmt.Printf("Team cost: %.2f - Total fantasy points: %.2f\n", permutations[i].GetTeamCost(), permutations[i].GetTeamPoints())
	}
}

func sortPermutations(permutations []models.FantasyTeam) {
	sort.Slice(permutations, func(i, j int) bool {
		return permutations[i].GetTeamPoints() > permutations[j].GetTeamPoints()
	})
}
