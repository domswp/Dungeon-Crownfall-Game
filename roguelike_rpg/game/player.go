package game

import "fmt"

type Player struct {
	Name      string
	Class     string
	Level     int
	Exp       int
	MaxExp    int
	MaxHP     int
	HP        int
	Atk       int
	Def       int
	Position  Position
	Inventory []Item
}

func NewPlayer(name string, class string) *Player {
	player := &Player{
		Name:      name,
		Class:     class,
		Position:  Position{X: 0, Y: 0},
		Inventory: []Item{},
	}

	player.Level = 1
	player.Exp = 0
	player.MaxExp = 20

	switch class {
	case "Ksatria":
		player.MaxHP = 50
		player.Atk = 8
		player.Def = 5
	case "Pencuri":
		player.MaxHP = 35
		player.Atk = 10
		player.Def = 2
	case "Penyihir":
		player.MaxHP = 25
		player.Atk = 15
		player.Def = 1
	default:
		player.MaxHP = 20
		player.Atk = 5
		player.Def = 1
	}

	player.HP = player.MaxHP
	return player
}

func (p *Player) CheckLevelUp() bool {
	leveledUp := false
	for p.Exp >= p.MaxExp {
		p.Exp -= p.MaxExp
		p.Level++
		p.MaxExp += 15 // Kebutuhan exp bertambah
		
		// Status bertambah
		p.MaxHP += 5
		if p.Class == "Ksatria" {
			p.Def += 2
			p.Atk += 1
		} else if p.Class == "Penyihir" {
			p.Atk += 3
		} else {
			p.Atk += 2
			p.Def += 1
		}
		
		p.HP = p.MaxHP // Pulih max HP
		
		fmt.Printf("\n=== SELAMAT! %s Naik ke Level %d! ===\n", p.Name, p.Level)
		fmt.Printf("Max HP meningkat menjadi %d!\n", p.MaxHP)
		fmt.Printf("Attack: %d, Defense: %d\n", p.Atk, p.Def)
		fmt.Println("=========================================")
		
		leveledUp = true
	}
	return leveledUp
}

func (p *Player) Move(dx int, dy int) bool {
	newX := p.Position.X + dx
	newY := p.Position.Y + dy

	// Cegah player keluar dari batas Peta
	if newX < 0 || newX >= MapWidth || newY < 0 || newY >= MapHeight {
		return false
	}

	p.Position.X = newX
	p.Position.Y = newY
	return true
}
