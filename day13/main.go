package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const MAX_INT = 9223372036854775807

type Game struct {
	A     [2]int
	B     [2]int
	Price [2]int
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	output := strings.Split(string(data), "\n")

	fmt.Printf("Part 1 Tokens: %d\n", getTokens(output))
	fmt.Printf("Part 2 Tokens: %d\n", getTokens2(output, 10000000000000))
}

func parseLine(s string, idx int) (x int, y int) {
	partsA := strings.Split(s[idx:], ", ")
	x, _ = strconv.Atoi(partsA[0][2:])
	y, _ = strconv.Atoi(partsA[1][2:])
	return x, y
}

func parseInput(s []string) []Game {
	games := make([]Game, (len(s)+1)/4)
	for idx := 0; idx < len(s); idx += 4 {
		aX, aY := parseLine(s[idx], 10)
		bX, bY := parseLine(s[idx+1], 10)
		pX, pY := parseLine(s[idx+2], 7)
		games[idx/4] = Game{
			A:     [2]int{aX, aY},
			B:     [2]int{bX, bY},
			Price: [2]int{pX, pY},
		}
	}
	return games
}

func getPrice(game Game) int {
	minimum := MAX_INT
	for a := 0; a < 100; a++ {
		for b := 0; b < 100; b++ {
			x := game.A[0]*a + game.B[0]*b
			y := game.A[1]*a + game.B[1]*b
			if x == game.Price[0] && y == game.Price[1] && (a*3+b) < minimum {
				minimum = a*3 + b
			}
		}
	}
	if minimum != MAX_INT {
		return minimum
	}
	return 0
}

func getTokens(s []string) int {
	games := parseInput(s)
	result := 0
	for _, game := range games {
		result += getPrice(game)
	}
	return result
}

func getPriceEquation(game Game, delta int) int {
	pX := game.Price[0] + delta
	pY := game.Price[1] + delta
	aX := game.A[0]
	aY := game.A[1]
	bX := game.B[0]
	bY := game.B[1]

	a := float64(pX*bY-pY*bX) / float64(aX*bY-aY*bX)
	b := float64(pY*aX-pX*aY) / float64(aX*bY-aY*bX)

	// if there is no decimals is valid
	if a == math.Trunc(a) && b == math.Trunc(b) {
		return int(a*3 + b)
	}
	return 0
}

func getTokens2(s []string, delta int) int {
	games := parseInput(s)
	result := 0
	for _, game := range games {
		result += getPriceEquation(game, delta)
	}
	return result
}
