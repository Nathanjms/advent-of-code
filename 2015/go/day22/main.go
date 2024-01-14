package main

import (
	"aoc-shared/pkg/sharedcode"
	"aoc-shared/pkg/sharedstruct"
	"math"
	"math/rand"
	"os"
	"path/filepath"
	"runtime"
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

func main() {
	// If another cmd argument has been passed, use that as the input path:
	if len(os.Args) > 1 {
		inputPath = os.Args[1]
	}

	var _, contents = sharedcode.ParseFile(inputPath)

	partOne(contents)
	partTwo(contents)
}

func partOne(contents []string) {
	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   22,
		Part:  1,
		Value: solve(contents, false),
	})
}

func partTwo(contents []string) {
	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   22,
		Part:  2,
		Value: solve(contents, true),
	})
}

var spellsAndCost = map[string]int{
	"Magic Missile": 53,
	"Drain":         73,
	"Shield":        113,
	"Poison":        173,
	"Recharge":      229,
}

func solve(contents []string, isPartTwo bool) int {
	playerHp := 50
	playerMp := 500
	enemyStats := parseEnemyStats(contents)

	// Do we just create a simulation and go through it? Could do the randomness trick again maybe
	frequencyByManaKey := make(map[int]int)
	resultsCount := 0
	lowestMana := math.MaxInt
	for {
		result := simulateGame(playerHp, playerMp, enemyStats, spellsAndCost, isPartTwo)

		if result.won {
			resultsCount++
			if result.manaUsed < lowestMana {
				lowestMana = result.manaUsed
			}
			if _, ok := frequencyByManaKey[result.manaUsed]; !ok {
				frequencyByManaKey[result.manaUsed] = 1
			} else {
				frequencyByManaKey[result.manaUsed]++
			}
		}

		// 1000 successful results should hopefully be enough
		if resultsCount > 1000 {
			return lowestMana
		}
	}
}

type stats struct {
	hp     int
	damage int
	armor  int
}

type usedUpMana struct {
	won      bool
	manaUsed int
}

type ongoingSpell struct {
	name           string
	remainingTurns int
}

func parseEnemyStats(contents []string) stats {
	var hp, damage int
	// HP is row 0
	tmp := strings.Split(contents[0], "Hit Points: ")
	hp, _ = strconv.Atoi(tmp[1])
	tmp = strings.Split(contents[1], "Damage: ")
	damage, _ = strconv.Atoi(tmp[1])

	return stats{hp: hp, damage: damage, armor: 0}
}

func getRandomSpell(spells map[string]int, remainingMana int, ongoingSpells []ongoingSpell) string {
	// Filter to the ones which you have mana for
	filteredSpells := make([]string, 0)
	for spell, manaCost := range spells {
		isActive := false
		for _, ongoingSpl := range ongoingSpells {
			if ongoingSpl.name == spell && ongoingSpl.remainingTurns > 0 {
				isActive = true
				break
			}
		}
		if remainingMana >= manaCost && !isActive {
			filteredSpells = append(filteredSpells, spell)
		}
	}
	if len(filteredSpells) == 0 {
		return "" // Not enough mana for any spell!
	}

	// Get random spell
	return filteredSpells[rand.Intn(len(filteredSpells))]
}

func simulateGame(playerHp int, playerMp int, enemyStats stats, spellsAndCost map[string]int, isPartTwo bool) usedUpMana {
	enemyHp := enemyStats.hp
	manaUsed := 0
	ongoingSpells := make([]ongoingSpell, 0)
	playerTurn := true

	for {
		// Pt2: lose 1hp per turn
		if isPartTwo && playerTurn {
			playerHp -= 1
			// Check after the 1hp drop effects
			if playerHp <= 0 {
				return usedUpMana{won: false}
			}
		}

		playerArmor := 0
		// First handle any ongoing spells
		for i, spell := range ongoingSpells {
			if spell.remainingTurns <= 0 {
				continue // skip if this spell is spent, could probably remove it from array but lets do this
			}
			switch spell.name {
			case "Magic Missile", "Drain":
				// Instant, so these should never be here. Put in for completeness
			case "Shield":
				playerArmor = 7
			case "Poison":
				enemyHp -= 3
			case "Recharge":
				playerMp += 101
			}
			ongoingSpells[i].remainingTurns--
		}
		// Check after the status effects
		if playerHp <= 0 || enemyHp <= 0 {
			break
		}

		if playerTurn {
			spell := getRandomSpell(spellsAndCost, playerMp, ongoingSpells)

			switch spell {
			case "Magic Missile":
				enemyHp -= 4
			case "Drain":
				enemyHp -= 2
				playerHp += 2
			case "Shield":
				ongoingSpells = append(ongoingSpells, ongoingSpell{spell, 6})
			case "Poison":
				ongoingSpells = append(ongoingSpells, ongoingSpell{spell, 6})
			case "Recharge":
				ongoingSpells = append(ongoingSpells, ongoingSpell{spell, 5})
			case "":
				return usedUpMana{won: false}
			}

			playerMp -= spellsAndCost[spell]
			manaUsed += spellsAndCost[spell]

			// Check after the status effects
			if playerHp <= 0 || enemyHp <= 0 {
				break
			}
		} else {
			// Now handle boss attack, accounting for possible magical armor
			playerHp = playerHp - max(1, enemyStats.damage-playerArmor)
			if playerHp <= 0 || enemyHp <= 0 {
				break
			}
		}
		playerTurn = !playerTurn
	}

	return usedUpMana{won: enemyHp <= 0, manaUsed: manaUsed}
}
