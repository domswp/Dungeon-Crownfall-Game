package game

import (
	"fmt"
	"math/rand"
	"time"
)

type Position struct {
	X int
	Y int
}

type Item struct {
	Name        string
	Description string
}

type Tile struct {
	Description string
	Visited     bool
	Enemy       *Enemy
	Loot        []Item
	IsPortal    bool
}

var WorldMap map[Position]*Tile
var CurrentLevel int = 1
var MapWidth int = 7
var MapHeight int = 7

func GenerateDungeon() {
	WorldMap = make(map[Position]*Tile)
	rand.Seed(time.Now().UnixNano())

	// Selalu mulai di {X: 0, Y: 0} untuk tiap level
	WorldMap[Position{0, 0}] = &Tile{
		Description: fmt.Sprintf("Kamu tiba di pintu masuk Lantai %d. Dindingnya berlumut dan lembab.", CurrentLevel),
		Visited:     true,
	}

	// Buat portal secara random, pastikan bukan di (0,0) kamar awal
	portalX := rand.Intn(MapWidth)
	portalY := rand.Intn(MapHeight)
	for portalX == 0 && portalY == 0 {
		portalX = rand.Intn(MapWidth)
		portalY = rand.Intn(MapHeight)
	}
	
	for y := 0; y < MapHeight; y++ {
		for x := 0; x < MapWidth; x++ {
			if x == 0 && y == 0 {
				continue
			}

			if x == portalX && y == portalY {
				WorldMap[Position{x, y}] = &Tile{
					Description: "Sebuah pusaran cahaya memancar terang. Ini adalah Portal menuju lantai selanjutnya!",
					IsPortal:    true,
				}
				continue
			}

			tile := &Tile{
				Description: "Sebuah koridor penjara yang dingin dan lapuk.",
			}

			prob := rand.Intn(100)
			if prob < 35 { // 35% men-generate Enemy
				enemyNames := []string{
					"Tikus Raksasa", "Tengkorak Hidup", "Orc Penjaga", 
					"Goblin Pencuri", "Slime Beracun", "Kelelawar Iblis", "Ksatria Korupsi",
				}
				
				name := enemyNames[rand.Intn(len(enemyNames))]
				
				// Skala kekuatan musuh bertambah seiring bertambahnya lantai
				hp := 10 + (CurrentLevel * 5) + rand.Intn(CurrentLevel)
				atk := 2 + (CurrentLevel * 2) + rand.Intn(3)
				def := 1 + CurrentLevel + rand.Intn(2)

				// Sistem Boss: Ada kemungkinan 20% bertemu boss setiap kelipatan 5 lantai
				isBoss := false
				if CurrentLevel%5 == 0 && rand.Intn(100) < 20 {
					isBoss = true
					name = "Raja " + name
					hp *= 3
					atk *= 2
					def *= 2
				}
				
				expReward := 10 + (CurrentLevel * 5)
				if isBoss {
					expReward *= 3
				}

				tile.Enemy = &Enemy{
					Name:      name,
					HP:        hp,
					Atk:       atk,
					Def:       def,
					ExpReward: expReward,
				}
			} else if prob < 50 { // 15% men-generate Loot
				tile.Loot = []Item{{Name: "Ramuan Penyembuh", Description: "Memulihkan sebagian HP-mu saat ditemukan."}}
			}

			WorldMap[Position{x, y}] = tile
		}
	}
}
