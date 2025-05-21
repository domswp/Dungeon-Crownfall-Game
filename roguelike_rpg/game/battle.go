package game

import (
	"fmt"
	"math/rand"
)

func StartBattle(player *Player, enemy *Enemy) {
	fmt.Printf("\n⚔️  PERTARUNGAN DIMULAI: %s vs %s ⚔️\n", player.Name, enemy.Name)

	for player.HP > 0 && enemy.HP > 0 {
		DisplayPlayerStatus(player)
		fmt.Printf("\nMusuh: %s | HP: %d/%d\n", enemy.Name, enemy.HP, enemy.MaxHP)
		fmt.Println("\nApa yang ingin kamu lakukan?")
		fmt.Println("1. Serang")
		fmt.Println("2. Gunakan Item")
		fmt.Println("3. Lari")
		fmt.Print("Masukkan pilihan (1-3): ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			// Serangan pemain
			damage := rand.Intn(10) + 5
			fmt.Printf("🔪 Kamu menyerang %s dan menyebabkan %d kerusakan!\n", enemy.Name, damage)
			enemy.HP -= damage
		case 2:
			// Gunakan item
			if len(player.Inventory) == 0 {
				fmt.Println("Tasmu kosong. Tidak ada item yang bisa digunakan.")
			} else {
				fmt.Println("🎒 Daftar item:")
				for i, item := range player.Inventory {
					fmt.Printf("%d. %s - %s\n", i+1, item.Name, item.Description)
				}
				fmt.Print("Pilih item yang ingin digunakan: ")
				var idx int
				fmt.Scanln(&idx)
				if idx > 0 && idx <= len(player.Inventory) {
					item := player.Inventory[idx-1]
					fmt.Printf("Kamu menggunakan %s.\n", item.Name)

					if item.Name == "Ramuan Penyembuh" || item.Name == "Healing Herb" {
						player.HP += 20
						if player.HP > player.MaxHP {
							player.HP = player.MaxHP
						}
						fmt.Println("🌿 Kamu merasa segar kembali. +20 HP.")
						player.Inventory = append(player.Inventory[:idx-1], player.Inventory[idx:]...)
					} else {
						fmt.Println("Tidak terjadi apa-apa...")
					}
				} else {
					fmt.Println("Pilihan item tidak valid.")
				}
			}
		case 3:
			// Coba melarikan diri
			if rand.Float64() < 0.5 {
				fmt.Println("🏃‍♂️ Kamu berhasil melarikan diri dari pertempuran!")
				return
			} else {
				fmt.Println("‼️ Kamu mencoba kabur, tapi musuh menghalangimu!")
			}
		default:
			fmt.Println("Kamu ragu-ragu... waktu terbuang percuma.")
		}

		// Efek spesial musuh (jika ada)
		if enemy.HP > 0 && enemy.Special != nil {
			specialEffect := enemy.Special(player)
			if specialEffect != "" {
				fmt.Println("🔥 " + specialEffect)
			}
		}

		// Serangan musuh
		if enemy.HP > 0 {
			damage := enemy.Attack(player)
			fmt.Printf("💥 %s menyerangmu dan menyebabkan %d kerusakan!\n", enemy.Name, damage)
		}
	}

	// Hasil akhir pertarungan
	if player.HP <= 0 {
		fmt.Println("\n🩸 Kamu telah gugur dalam pertempuran... Rohmu menghilang ke kegelapan.")
	} else if enemy.HP <= 0 {
		fmt.Printf("\n✅ Kamu mengalahkan %s!\n", enemy.Name)

		// XP
		player.GainXP(enemy.XPReward)

		// Loot
		if len(enemy.DropLoot) > 0 {
			fmt.Println("🎁 Kamu menemukan:")
			for _, item := range enemy.DropLoot {
				fmt.Printf("- %s: %s\n", item.Name, item.Description)
				player.Inventory = append(player.Inventory, item)
			}
		}
	}
}

