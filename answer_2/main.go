package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	var pemain, dadu int
	fmt.Print("Masukkan jumlah pemain: ")
	fmt.Scan(&pemain)
	fmt.Print("Masukkan jumlah dadu: ")
	fmt.Scan(&dadu)

	players := make([][]int, pemain)

	for i := range players {
		players[i] = make([]int, dadu)
	}

	fmt.Printf("Pemain %d , Dadu : %v\n", pemain, dadu)
	fmt.Println("==================")
	playerScore := make(map[int]int)

	for round := 1; ; round++ {
		fmt.Printf("Giliran %d lempar dadu\n", round)
		for i := range players {
			if len(players[i]) == 0 {
				continue
			}
			fmt.Printf("Pemain #%d (%d): %v\n", i+1, len(players[i]), players[i])
			for j := range players[i] {
				players[i][j] = rand.Intn(6) + 1
			}
		}
		fmt.Println(players)

		for i := range players {

			if len(players[i]) == 0 {
				continue
			}
			var newDice []int
			score := 0
			for j := range players[i] {
				if players[i][j] == 1 {
					if j+1 < len(players[i]) {
						players[(i+1)%pemain] = append(players[(i+1)%pemain], 1)
					}
				} else if players[i][j] != 6 {
					newDice = append(newDice, players[i][j])
				} else if players[i][j] == 6 {
					score += 1
				}

			}
			playerScore[i] = score
			players[i] = newDice
		}
		fmt.Println("Setelah evaluasi:")
		for i := range players {
			if len(players[i]) == 0 {
				fmt.Printf("Pemain #%d (%d): (Berhenti bermain karena tidak memiliki dadu)\n", i+1, len(players[i]))
			} else {
				fmt.Printf("Pemain #%d (%d): %v\n", i+1, len(players[i]), players[i])
			}
		}
		activePlayers := 0
		var maxScore int
		// playerScore := make(map[int]int)
		score := 0
		for i := range players {
			if len(players[i]) > 0 {
				activePlayers++
				for j := range players[i] {
					if players[i][j] == 6 {
					}
				}
				if score > maxScore {
					maxScore = score
				}
			}
		}

		if activePlayers <= 1 {
			fmt.Printf("Game berakhir karena hanya 1 pemain yang memiliki dadu\n")
			maxVal := playerScore[0]
			maxKey := 1

			for key, val := range playerScore {
				if val > maxVal {
					maxVal = val
					maxKey = key
				}
			}

			fmt.Printf("Game dimenangkan oleh pemain #%d karena memiliki poin lebih banyak dari pemain lainnya dengan skor %d.\n", maxKey, maxVal)
			break
		}

		fmt.Println("==================")
	}

}
