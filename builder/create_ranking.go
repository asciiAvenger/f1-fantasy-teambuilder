package builder

import (
	"f1-fantasy-teambuilder/models"
	"fmt"
)

func CreateRanking(data models.Data) {
	fmt.Println("##########")
	fmt.Println("# Teams: #")
	fmt.Println("##########")
	rankUnits(data.Teams)
	fmt.Println("############")
	fmt.Println("# Drivers: #")
	fmt.Println("############")
	rankUnits(data.Drivers)
}

func rankUnits(units []models.Unit) {
	sortUnitsDesc(units)
	fmt.Println("NAME\t\t\tPRICE\t\tSCORE")
	for _, unit := range units {
		extraTabs := ""
		if len(unit.Name) < 8 {
			extraTabs = "\t"
		}
		fmt.Printf("%s\t\t%s%.1f\t\t%.2f\n", unit.Name, extraTabs, unit.Cost, unit.GetScore())
	}
}
