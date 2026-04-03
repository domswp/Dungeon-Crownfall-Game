package game

import (
	"fmt"
	"time"
)

func StartBattle(player *Player, enemy *Enemy) {
	fmt.Printf("\n--- PERTEMPURAN DIMULAI: %s vs %s ---\n", player.Name, enemy.Name)

	for player.HP > 0 && enemy.HP > 0 {
		// --- Giliran Pemain ---
		playerDamage := player.Atk - enemy.Def
		if playerDamage < 0 {
			playerDamage = 0
		}
		enemy.HP -= playerDamage
		fmt.Printf("Kamu menyerang %s dan memberikan %d *damage*!\n", enemy.Name, playerDamage)
		fmt.Printf("(HP %s tersisa: %d)\n", enemy.Name, enemy.HP)

		if enemy.HP <= 0 {
			fmt.Printf("\nKamu berhasil mengalahkan %s!\n", enemy.Name)
			break
		}
		time.Sleep(800 * time.Millisecond)

		// --- Giliran Musuh ---
		enemyDamage := enemy.Atk - player.Def
		if enemyDamage < 0 {
			enemyDamage = 0
		}
		player.HP -= enemyDamage
		fmt.Printf("%s membalas menyerangmu sebesar %d *damage*!\n", enemy.Name, enemyDamage)
		fmt.Printf("(HP Kamu tersisa: %d)\n", player.HP)

		if player.HP <= 0 {
			fmt.Println("\nKamu telah kalah dalam pertempuran...")
			break
		}
		fmt.Println("-------------------------------------------")
		time.Sleep(800 * time.Millisecond)
	}
}
