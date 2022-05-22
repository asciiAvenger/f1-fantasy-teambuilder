package printer

import (
	"f1-fantasy-teambuilder/models"
	"fmt"
	"math"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
)

func PrintResults(results []models.FantasyTeam, topTeams int, budget float64) {
	fmt.Printf("Displaying top %d teams out of %d possible teams...\n", topTeams, len(results))
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "name", "points", "price"})
	for i := 0; i < topTeams; i++ {
		t.AppendRow(table.Row{i + 1, results[i].Team.FirstName, results[i].Team.SeasonScore, results[i].Team.Price})
		var drivers []table.Row
		for _, driver := range results[i].Drivers {
			drivers = append(drivers, table.Row{"", fmt.Sprintf("%s %s", driver.FirstName, driver.LastName), driver.SeasonScore, driver.Price})
		}
		t.AppendRows(drivers)
		t.AppendRow(table.Row{"", "", roundNum(results[i].GetTeamSeasonScore()), roundNum(results[i].GetTeamPrice())})
		t.AppendSeparator()
	}
	t.Render()
}

func roundNum(num float64) float64 {
	return math.Round(num*100) / 100
}
