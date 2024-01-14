package main

import (
	"aoc-shared/pkg/sharedcode"
	"aoc-shared/pkg/sharedstruct"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"sort"
	"strconv"
	"strings"
)

func getCurrentDirectory() string {
	_, filename, _, _ := runtime.Caller(0)
	dirname := filepath.Dir(filename)
	return dirname
}

// Default Input path is current directory + example-input
var inputPath = filepath.Join(getCurrentDirectory(), "example-input")
var isUsingExample = true

func main() {
	// If another cmd argument has been passed, use that as the input path:
	if len(os.Args) > 1 {
		inputPath = os.Args[1]
		isUsingExample = false
	}

	var contents, _ = sharedcode.ParseFile(inputPath)

	partOne(contents)
	partTwo(contents)
}

func partOne(contents string) {
	playerHealth := 100
	enemyStats := parseEnemyStats(contents)
	shopItems := parseItemShop()

	// Build up all the possibilities, then order by the lowest score. Work upwards until we defeat the enemy
	// The options are as follows:
	// a weapon MUST be used
	// armor is optional
	// rings can be doubled up (but not the same ring)

	// To handle optional armor, I'll just add another option "none" to that item list
	shopItems["Armor"] = append(shopItems["Armor"], shopItem{name: "none"})

	// Make this a map, with key of cost-damage-armor to deal with duplicates for free
	loadouts := make(map[string]loadout)

	// Now go through each option
	for _, weapon := range shopItems["Weapons"] {
		// 1. Choose Weapon
		chosenWeapon := weapon
		for _, armor := range shopItems["Armor"] {
			// 2. Choose Armor
			chosenArmour := armor
			for _, ring := range shopItems["Rings"] {
				// 3a. Choose Ring
				chosenRing := ring
				totalCost := chosenWeapon.cost + chosenArmour.cost + chosenRing.cost
				totalDamage := chosenWeapon.damage + chosenArmour.damage + chosenRing.damage
				totalArmor := chosenWeapon.armor + chosenArmour.armor + chosenRing.armor

				loadouts[makeKey(totalCost, totalDamage, totalArmor)] = loadout{
					cost:   totalCost,
					damage: totalDamage,
					armor:  totalArmor,
				}

				for _, secondRing := range shopItems["Rings"] {
					// 3b. Choose second ring
					if ring == secondRing {
						break // Can't wear the same ring
					}
					loadouts[makeKey(totalCost+secondRing.cost, totalDamage+secondRing.damage, totalArmor+secondRing.armor)] = loadout{
						cost:   totalCost + secondRing.cost,
						damage: totalDamage + secondRing.damage,
						armor:  totalArmor + secondRing.armor,
					}
				}
			}
		}
	}

	// We cna make it a slice now instead of a map, now duplicates have been handles by the map:
	loadoutsSlice := make([]loadout, 0)
	for _, v := range loadouts {
		loadoutsSlice = append(loadoutsSlice, v)
	}

	// Sort by lowest cost and then go upwards
	sort.Slice(loadoutsSlice, func(i, j int) bool {
		return loadoutsSlice[i].cost < loadoutsSlice[j].cost
	})

	lowestCost := math.MaxInt
	for _, loadout := range loadoutsSlice {
		// Does it win?
		playerOffence := max(1, loadout.damage-enemyStats.armor)
		enemyOffence := max(1, enemyStats.damage-loadout.armor)

		turnsForPlayerWin := math.Ceil(float64(enemyStats.hp) / float64(playerOffence))
		turnsForEnemyWin := math.Ceil(float64(playerHealth) / float64(enemyOffence))

		if turnsForPlayerWin <= turnsForEnemyWin {
			lowestCost = loadout.cost
			break
		}
	}

	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   21,
		Part:  1,
		Value: lowestCost,
	})
}

func partTwo(contents string) {
	playerHealth := 100
	enemyStats := parseEnemyStats(contents)
	shopItems := parseItemShop()

	// To handle optional armor, I'll just add another option "none" to that item list
	shopItems["Armor"] = append(shopItems["Armor"], shopItem{name: "none"})

	// Make this a map, with key of cost-damage-armor to deal with duplicates for free
	loadouts := make(map[string]loadout)

	// Now go through each option
	for _, weapon := range shopItems["Weapons"] {
		// 1. Choose Weapon
		chosenWeapon := weapon
		for _, armor := range shopItems["Armor"] {
			// 2. Choose Armor
			chosenArmour := armor
			for _, ring := range shopItems["Rings"] {
				// 3a. Choose Ring
				chosenRing := ring
				totalCost := chosenWeapon.cost + chosenArmour.cost + chosenRing.cost
				totalDamage := chosenWeapon.damage + chosenArmour.damage + chosenRing.damage
				totalArmor := chosenWeapon.armor + chosenArmour.armor + chosenRing.armor

				loadouts[makeKey(totalCost, totalDamage, totalArmor)] = loadout{
					cost:   totalCost,
					damage: totalDamage,
					armor:  totalArmor,
				}

				for _, secondRing := range shopItems["Rings"] {
					// 3b. Choose second ring
					if ring == secondRing {
						break // Can't wear the same ring
					}
					loadouts[makeKey(totalCost+secondRing.cost, totalDamage+secondRing.damage, totalArmor+secondRing.armor)] = loadout{
						cost:   totalCost + secondRing.cost,
						damage: totalDamage + secondRing.damage,
						armor:  totalArmor + secondRing.armor,
					}
				}
			}
		}
	}

	// We cna make it a slice now instead of a map, now duplicates have been handles by the map:
	loadoutsSlice := make([]loadout, 0)
	for _, v := range loadouts {
		loadoutsSlice = append(loadoutsSlice, v)
	}

	// Sort by highest cost and then go downwards
	sort.Slice(loadoutsSlice, func(i, j int) bool {
		return loadoutsSlice[i].cost > loadoutsSlice[j].cost
	})

	highestCost := 0
	for _, loadout := range loadoutsSlice {
		// Does it win?
		playerOffence := max(1, loadout.damage-enemyStats.armor)
		enemyOffence := max(1, enemyStats.damage-loadout.armor)

		turnsForPlayerWin := math.Ceil(float64(enemyStats.hp) / float64(playerOffence))
		turnsForEnemyWin := math.Ceil(float64(playerHealth) / float64(enemyOffence))

		if turnsForPlayerWin > turnsForEnemyWin {
			highestCost = loadout.cost
			break
		}
	}
	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   21,
		Part:  2,
		Value: highestCost,
	})
}

type stats struct {
	hp     int
	damage int
	armor  int
}

type loadout struct {
	cost   int
	damage int
	armor  int
}

type shopItem struct {
	name   string
	cost   int
	damage int
	armor  int
}

func parseEnemyStats(contents string) stats {
	var hp, damage, armor int
	contentsByLine := strings.Split(contents, "\n")
	// HP is row 0
	tmp := strings.Split(contentsByLine[0], "Hit Points: ")
	hp, _ = strconv.Atoi(tmp[1])
	tmp = strings.Split(contentsByLine[1], "Damage: ")
	damage, _ = strconv.Atoi(tmp[1])
	tmp = strings.Split(contentsByLine[2], "Armor: ")
	armor, _ = strconv.Atoi(tmp[1])

	return stats{hp: hp, damage: damage, armor: armor}
}

func parseItemShop() map[string][]shopItem {
	itemShopContents, _ := sharedcode.ParseFile(getCurrentDirectory() + "/item-shop")

	// First we split by 2 new lines, to get three arrays of weapons, armor and rings:

	splits := strings.Split(itemShopContents, "\n\n")

	itemsByType := make(map[string][]shopItem)
	currentKey := ""
	for _, split := range splits {
		// For each type, we now want to explode by \n
		items := strings.Split(split, "\n")
		for j, item := range items {
			if j == 0 {
				// We can find the key!
				lineOne := strings.Split(item, ":")
				currentKey = lineOne[0]
				itemsByType[currentKey] = make([]shopItem, 0)
			} else {
				re := regexp.MustCompile(`^(\w+(?: \+\d)?)\s*(\w+)\s*(\w+)\s*(\w+)`)
				var matches = re.FindStringSubmatch(item)

				name := matches[1]
				cost, _ := strconv.Atoi(matches[2])
				damage, _ := strconv.Atoi(matches[3])
				armor, _ := strconv.Atoi(matches[4])

				itemsByType[currentKey] = append(itemsByType[currentKey], shopItem{name, cost, damage, armor})
			}
		}
	}
	return itemsByType
}

func makeKey(cost int, damage int, armor int) string {
	return fmt.Sprintf("%d-%d-%d", cost, damage, armor)
}
