package game

import "math/rand"

type Enemy struct {
	Name        string
	HP          int
	MaxHP       int
	Damage      int
	Type        string
	Description string
	DropLoot    []Item
	XPReward    int
	Special     func(*Player) string
}

func (e *Enemy) IsAlive() bool {
	return e.HP > 0
}

func (e *Enemy) Attack(p *Player) int {
	damage := rand.Intn(e.Damage) + 1
	p.HP -= damage
	return damage
}

// === Musuh-Musuh Crownfall ===

func NewGoblin() *Enemy {
	return &Enemy{
		Name:        "Goblin",
		HP:          25,
		MaxHP:       25,
		Damage:      6,
		Type:        "Binatang",
		Description: "Makhluk kecil yang gesit dan licik, menyukai tempat gelap dan bau darah.",
		XPReward:    20,
		DropLoot: []Item{
			{Name: "Taring Goblin", Description: "Taring tajam dari goblin hutan."},
		},
		Special: func(p *Player) string {
			if rand.Float64() < 0.2 {
				p.HP -= 5
				return "Goblin menyelinap dan menyerang dari belakang! Kau kehilangan 5 HP!"
			}
			return ""
		},
	}
}

func NewSkeleton() *Enemy {
	return &Enemy{
		Name:        "Skeleton",
		HP:          40,
		MaxHP:       40,
		Damage:      8,
		Type:        "Mayat Hidup",
		Description: "Tulang-tulang hidup yang dikutuk, tanpa jiwa tapi haus darah.",
		XPReward:    35,
		DropLoot: []Item{
			{Name: "Serpihan Tulang", Description: "Tulang terkutuk dari Skeleton."},
		},
		Special: func(p *Player) string {
			if rand.Float64() < 0.3 {
				p.HP -= 3
				return "Skeleton mengayunkan pedang beku! Tubuhmu menggigil... -3 HP."
			}
			return ""
		},
	}
}

func NewDarkMage() *Enemy {
	return &Enemy{
		Name:        "Dark Mage",
		HP:          30,
		MaxHP:       30,
		Damage:      5,
		Type:        "Ahli Sihir",
		Description: "Penyihir hitam yang memanggil kekuatan dari balik bayangan abadi.",
		XPReward:    50,
		DropLoot: []Item{
			{Name: "Kitab Hitam", Description: "Kitab sihir terlarang yang menyimpan kekuatan jahat."},
		},
		Special: func(p *Player) string {
			if rand.Float64() < 0.4 {
				p.HP -= 10
				return "Dark Mage melempar peluru bayangan! Jiwamu terkoyak. -10 HP!"
			}
			return ""
		},
	}
}

