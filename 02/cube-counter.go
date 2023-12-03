package _2

import (
	"aoc-2023/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type game struct {
	ID    int
	Games []subGame
}

type subGame struct {
	Cubes []cube
}

type cube struct {
	Color  string
	Amount int
}

func SumOfInvalidGameIDs() int {
	// Format game String -> { GameID, SubGames{ cube{ Color, Amount } } }
	values := utils.ReadStringArrayFromFile("02")
	games := formatInput(values)

	sumLowestAmounts := 0
	for _, g := range games {
		fmt.Println(isInvalidGame(g))
		red, green, blue := findLowestValue(g)
		power := red.Amount * green.Amount * blue.Amount
		sumLowestAmounts += power
	}

	return sumLowestAmounts
}

func isInvalidGame(g game) bool {
	redMax := 12
	greenMax := 13
	blueMax := 14

	for _, s := range g.Games {
		for _, c := range s.Cubes {
			if c.Color == "red" && c.Amount > redMax {
				return true
			}
			if c.Color == "green" && c.Amount > greenMax {
				return true
			}
			if c.Color == "blue" && c.Amount > blueMax {
				return true
			}
		}
	}
	return false
}

func findLowestValue(g game) (cube, cube, cube) {
	lowestRed := cube{
		Color:  "red",
		Amount: math.MinInt,
	}
	lowestGreen := cube{
		Color:  "green",
		Amount: math.MinInt,
	}
	lowestBlue := cube{
		Color:  "blue",
		Amount: math.MinInt,
	}

	for _, s := range g.Games {
		for _, c := range s.Cubes {
			if c.Color == "red" && c.Amount > lowestRed.Amount {
				lowestRed = c
			}
			if c.Color == "green" && c.Amount > lowestGreen.Amount {
				lowestGreen = c
			}
			if c.Color == "blue" && c.Amount > lowestBlue.Amount {
				lowestBlue = c
			}
		}
	}
	return lowestRed, lowestGreen, lowestBlue
}

func formatInput(values []string) []game {
	games := make([]game, 0)
	for _, value := range values {
		splitGameAndSubGames := strings.Split(value, ":")
		gameIDStr := strings.Split(splitGameAndSubGames[0], " ")[1]
		gameId, err := strconv.Atoi(gameIDStr)
		if err != nil {
			panic(err)
		}
		subGamesInp := strings.Split(splitGameAndSubGames[1], ";")
		subGames := make([]subGame, 0)
		for _, sg := range subGamesInp {
			cubeAndAmount := strings.Split(sg, ", ")
			cubes := make([]cube, 0)
			for _, ca := range cubeAndAmount {
				amountAndColor := strings.Split(strings.TrimSpace(ca), " ")
				amount, err := strconv.Atoi(amountAndColor[0])
				if err != nil {
					panic(err)
				}
				cubes = append(cubes, cube{
					Color:  amountAndColor[1],
					Amount: amount,
				})
			}
			subGaem := subGame{Cubes: cubes}
			subGames = append(subGames, subGaem)
		}
		games = append(games, game{
			ID:    gameId,
			Games: subGames,
		})
	}
	return games
}
