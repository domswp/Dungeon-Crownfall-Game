package game

import (
	"fmt"
)

func PrintDivider() {
	fmt.Println("----------------------------------------")
}

func DisplayPlayerStatus(p *Player) {
	fmt.Println()
	PrintDivider()
	fmt.Printf("⚔️  %s sang %s\n", p.Name, p.Class)
	fmt.Printf("❤️  HP: %d/%d\n", p.HP, p.MaxHP)
	fmt.Printf("⭐  Level: %d  |  XP: %d/%d\n", p.Level, p.XP, p.NextXP)
	fmt.Printf("🎒  Jumlah Item: %d\n", len(p.Inventory))
	PrintDivider()
}

