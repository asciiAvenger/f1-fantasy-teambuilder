package builder

import (
	"f1-fantasy-teambuilder/models"
	"sort"
	"strconv"
	"strings"
)

func CreatePermutations(data models.Data, budget float64) []models.FantasyTeam {
	permutations := []models.FantasyTeam{}
	numDrivers := len(data.Drivers)
	usedTeams := make(map[int]bool)
	for i1 := 0; i1 < numDrivers; i1++ {
		for i2 := 0; i2 < numDrivers; i2++ {
			for i3 := 0; i3 < numDrivers; i3++ {
				for i4 := 0; i4 < numDrivers; i4++ {
					for i5 := 0; i5 < numDrivers; i5++ {
						nums := []int{i1, i2, i3, i4, i5}
						sort.Ints(nums)
						var numsConv []string
						for _, n := range nums {
							numsConv = append(numsConv, strconv.Itoa(n))
						}
						keyString := strings.Join(numsConv, "")
						key, _ := strconv.Atoi(keyString)
						if checkTwoSameNums(i1, i2, i3, i4, i5) || usedTeams[key] {
							continue
						}
						usedTeams[key] = true
						team := models.FantasyTeam{Drivers: []models.Unit{
							data.Drivers[i1],
							data.Drivers[i2],
							data.Drivers[i3],
							data.Drivers[i4],
							data.Drivers[i5],
						}}
						for _, t := range data.Teams {
							team.Team = t
							if team.GetTeamCost() < budget {
								permutations = append(permutations, team)
							}
						}
					}
				}
			}
		}
	}
	return permutations
}

func checkTwoSameNums(n1 int, n2 int, n3 int, n4 int, n5 int) bool {
	return n1 == n2 || n1 == n3 || n1 == n4 || n1 == n5 || n2 == n3 || n2 == n4 || n2 == n5 || n3 == n4 || n3 == n5 || n4 == n5
}
