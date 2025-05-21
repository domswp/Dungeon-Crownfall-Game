package game

import (
	"fmt"
	"math/rand"
)

type Position struct {
	X, Y int
}

type Tile struct {
	Description string
	Enemy       *Enemy
	Loot        []Item
	Visited     bool
}

var WorldMap = map[Position]*Tile{}

func GenerateDungeon() {
	for x := -3; x <= 3; x++ {
		for y := -3; y <= 3; y++ {
			pos := Position{X: x, Y: y}
			if rand.Float64() < 0.5 {
				WorldMap[pos] = &Tile{
					Description: generateRoomDescription(),
					Enemy:       maybeSpawnEnemy(),
					Loot:        maybeDropLoot(),
					Visited:     false,
				}
			}
		}
	}
	fmt.Println("Penjara bawah tanah telah terbentuk.")
}

func generateRoomDescription() string {
	descs := []string{
		"Ruangan gelap dan lembap dengan dinding batu berlumut.",
		"Lorong sempit dengan jejak darah di lantainya.",
		"Reruntuhan kapel tua, sunyi dan penuh aura aneh.",
		"Ruangan terbuka yang dipenuhi tengkorak dan tulang-tulang.",
		"Sudut penjara dengan simbol kuno tergurat di dinding.",
	}
	return descs[rand.Intn(len(descs))]
}

func maybeSpawnEnemy() *Enemy {
	r := rand.Float64()
	switch {
	case r < 0.2:
		return NewGoblin()
	case r < 0.4:
		return NewSkeleton()
	case r < 0.5:
		return NewDarkMage()
	default:
		return nil
	}
}

func maybeDropLoot() []Item {
	if rand.Float64() < 0.4 {
		return []Item{
			{Name: "Ramuan Penyembuh", Description: "Mengembalikan 20 HP saat digunakan."},
		}
	}
	return nil
}

