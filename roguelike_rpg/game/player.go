package game

type Player struct {
	Name      string
	Class     string
	HP        int
	MaxHP     int
	Level     int
	XP        int
	NextXP    int
	Position  Position
	Inventory []Item
}

func NewPlayer(name, class string) *Player {
	return &Player{
		Name:     name,
		Class:    class,
		HP:       100,
		MaxHP:    100,
		Level:    1,
		XP:       0,
		NextXP:   50,
		Position: Position{X: 0, Y: 0},
		Inventory: []Item{
			{Name: "Rusty Sword", Description: "Pedang tua yang masih bisa digunakan untuk bertahan hidup."},
		},
	}
}

func (p *Player) Move(dx, dy int) {
	p.Position.X += dx
	p.Position.Y += dy
}

func (p *Player) GainXP(amount int) {
	p.XP += amount
	println()
	println("Kamu mendapatkan", amount, "XP! [", p.XP, "/", p.NextXP, "]")
	if p.XP >= p.NextXP {
		p.LevelUp()
	}
}

func (p *Player) LevelUp() {
	p.Level++
	p.XP -= p.NextXP
	p.NextXP += 25
	p.MaxHP += 20
	p.HP = p.MaxHP
	println("\n=== NAIK LEVEL! Sekarang kamu level", p.Level, "! ===")
	println("Tubuhmu terasa lebih kuat dan penuh semangat baru.")
}

