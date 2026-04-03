package game

import (
	"fmt"
	"time"
)

func StartBattle(player *Player, enemy *Enemy) {
	fmt.Printf("\n--- PERTEMPURAN DIMULAI: %s vs %s ---\n", player.Name, enemy.Name)

	for player.HP > 0 && enemy.HP > 0 {
		// --- Giliran Pemain ---
		validAction := false
		for !validAction {
			fmt.Println("\nPilihan Giliranmu:")
			fmt.Println("1. ⚔️ Serang dengan Senjata")
			fmt.Println("2. 🍾 Minum Potion")
			fmt.Printf("3. 💥 Gunakan Skill (Sisa MP: %d/%d)\n", player.MP, player.MaxMP)
			fmt.Print("Pilih Aksi (1/2/3): ")
			
			var aksi string
			fmt.Scanln(&aksi)

			switch aksi {
			case "1":
				playerDamage := player.Atk - enemy.Def
				if playerDamage < 0 {
					playerDamage = 0
				}
				enemy.HP -= playerDamage
				fmt.Printf("\n🗡️ Kamu menyerang %s dan memberikan %d *damage*!\n", enemy.Name, playerDamage)
				fmt.Printf("(HP %s tersisa: %d)\n", enemy.Name, enemy.HP)
				validAction = true

			case "2":
				used := player.UsePotion()
				if used {
					validAction = true
				} else {
					fmt.Println("Silakan pilih tindakan lain.")
				}
			case "3":
				used := player.UseSkill(enemy)
				if used {
					validAction = true
				} else {
					fmt.Println("Silakan pilih tindakan lain.")
				}
			default:
				fmt.Println("❌ Aksi tidak valid. Tekan tombol 1, 2, atau 3 lalu enter.")
			}
		}

		if enemy.HP <= 0 {
			fmt.Printf("\n🎉 Kamu berhasil mengalahkan %s!\n", enemy.Name)
			break
		}
		time.Sleep(800 * time.Millisecond)

		// --- Giliran Musuh ---
		enemyDamage := enemy.Atk - player.Def
		if enemyDamage < 0 {
			enemyDamage = 0
		}
		player.HP -= enemyDamage
		fmt.Printf("\n🩸 %s membalas menyerangmu sebesar %d *damage*!\n", enemy.Name, enemyDamage)
		fmt.Printf("(❤️ HP Kamu tersisa: %d)\n", player.HP)

		if player.HP <= 0 {
			fmt.Println("\n⚰️ Kamu telah kalah dalam pertempuran...")
			break
		}
		fmt.Println("-------------------------------------------")
		time.Sleep(800 * time.Millisecond)
	}
}
