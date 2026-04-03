package game

import "fmt"

type Player struct {
	Name      string
	Class     string
	Level     int
	Exp       int
	MaxExp    int
	MaxHP     int
	MaxMP     int
	HP        int
	MP        int
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
		player.MaxMP = 20
		player.Atk = 10
		player.Def = 6
	case "Pencuri":
		player.MaxHP = 35
		player.MaxMP = 30
		player.Atk = 12
		player.Def = 3
	case "Penyihir":
		player.MaxHP = 25
		player.MaxMP = 50
		player.Atk = 4
		player.Def = 1
	default:
		player.MaxHP = 20
		player.MaxMP = 10
		player.Atk = 5
		player.Def = 1
	}

	player.HP = player.MaxHP
	player.MP = player.MaxMP
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
		switch p.Class {
		case "Ksatria":
			p.Def += 2
			p.Atk += 2
			p.MaxMP += 5
		case "Penyihir":
			p.Atk += 3
			p.MaxMP += 15
		default:
			p.Atk += 3
			p.Def += 1
			p.MaxMP += 10
		}
		
		p.HP = p.MaxHP // Pulih max HP
		p.MP = p.MaxMP // Pulih max MP
		
		fmt.Printf("\n=== SELAMAT! %s Naik ke Level %d! ===\n", p.Name, p.Level)
		fmt.Printf("Max HP meningkat menjadi %d! (Max MP: %d)\n", p.MaxHP, p.MaxMP)
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

func (p *Player) UsePotion() bool {
	for i, item := range p.Inventory {
		if item.Name == "Ramuan Penyembuh" {
			p.Inventory = append(p.Inventory[:i], p.Inventory[i+1:]...)
			
			healAmount := 20
			p.HP += healAmount
			if p.HP > p.MaxHP {
				p.HP = p.MaxHP
			}
			fmt.Printf("\n🍾 Kamu meminum %s dan memulihkan %d HP!\n", item.Name, healAmount)
			fmt.Printf("❤️  HP Kamu sekarang: %d/%d\n", p.HP, p.MaxHP)
			return true
		}
	}
	fmt.Println("\n❌ Kamu tidak memiliki Ramuan Penyembuh di dalam tas!")
	return false
}

func (p *Player) UseSkill(enemy *Enemy) bool {
	switch p.Class {
	case "Ksatria":
		if p.MP < 5 {
			fmt.Println("❌ MP tidak cukup! (Butuh 5 MP)")
			return false
		}
		p.MP -= 5
		damage := p.Atk * 2
		enemy.HP -= damage
		fmt.Printf("\n💥 [HEAVY STRIKE] Kamu menghantam %s dgn perisaimu memberikan %d *damage* murni!\n", enemy.Name, damage)
		fmt.Printf("(HP %s tersisa: %d)\n", enemy.Name, enemy.HP)
		return true

	case "Pencuri":
		if p.MP < 8 {
			fmt.Println("❌ MP tidak cukup! (Butuh 8 MP)")
			return false
		}
		p.MP -= 8
		damage := p.Atk
		enemy.HP -= damage
		lifesteal := damage
		p.HP += lifesteal
		if p.HP > p.MaxHP {
			p.HP = p.MaxHP
		}
		fmt.Printf("\n🧛 [VAMPIRIC STAB] Kamu menusuk %s, menimbulkan %d *damage* dan menyedot %d HP!\n", enemy.Name, damage, lifesteal)
		fmt.Printf("(HP %s tersisa: %d)\n", enemy.Name, enemy.HP)
		fmt.Printf("(❤️ HP Kamu tersisa: %d)\n", p.HP)
		return true

	case "Penyihir":
		if p.MP < 15 {
			fmt.Println("❌ MP tidak cukup! (Butuh 15 MP)")
			return false
		}
		p.MP -= 15
		damage := p.Atk * 6
		enemy.HP -= damage
		fmt.Printf("\n☄️ [METEOR STRIKE] Kamu merapal sihir bola api ke %s dan memberikan %d *MAHAMENDAMAGE*!\n", enemy.Name, damage)
		fmt.Printf("(HP %s tersisa: %d)\n", enemy.Name, enemy.HP)
		return true

	default:
		// Petani atau unhandled
		fmt.Println("❌ Kelasmu belum memiliki Skill khusus.")
		return false
	}
}
