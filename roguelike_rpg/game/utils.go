package game

import "fmt"

func DrawMap(player *Player) {
	fmt.Println("\n==================================")
	fmt.Printf(" 🗺️  Lantai %d             \n", CurrentLevel)
	fmt.Println("----------------------------------")
	fmt.Printf(" 👤 %s si %s (Lvl %d)\n", player.Name, player.Class, player.Level)
	fmt.Printf(" ❤️  HP: %d/%d   ⭐️ Exp: %d/%d\n", player.HP, player.MaxHP, player.Exp, player.MaxExp)
	fmt.Printf(" ⚔️  Atk: %d     🛡️  Def: %d     🌠 MP: %d/%d\n", player.Atk, player.Def, player.MP, player.MaxMP)
	fmt.Println("==================================")
	// Y (Baris)
	for y := player.Position.Y - 2; y <= player.Position.Y + 2; y++ {
		// X (Kolom)
		for x := player.Position.X - 2; x <= player.Position.X + 2; x++ {
			pos := Position{X: x, Y: y}
			
			// Jika titik rendering ada di luar batas grid (0-6), cetak icon tembok
			if x < 0 || x >= MapWidth || y < 0 || y >= MapHeight {
				fmt.Print("[#]") // Tembok Pembatas
				continue
			}

			if player.Position == pos {
				fmt.Print("[P]") // Posisi Pemain
			} else if tile, exists := WorldMap[pos]; exists {
				if !tile.Visited {
					fmt.Print("[?]") // Ruang Misterius tersembunyi
				} else if tile.IsPortal {
					fmt.Print("[O]") // Simbol Portal (Hanya tampil jika sudah dilewati)
				} else if tile.Enemy != nil {
					fmt.Print("[E]") // Musuh bersiaga
				} else if len(tile.Loot) > 0 {
					fmt.Print("[*]") // Harta
				} else {
					fmt.Print("[.]") // Kosong / Telah Dijelajahi
				}
			} else {
				fmt.Print("[ ]") // Ruang null (jaga-jaga)
			}
		}
		fmt.Println() 
	}
	fmt.Println("==================================")
}
