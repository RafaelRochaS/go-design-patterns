package factory

import "fmt"

type Sword struct {
	Name         string
	DamageNumber int
	Level        int
	Durability   int
}

func (s Sword) Damage(target string) string {
	return fmt.Sprintf("%s took %d of damage!", target, s.DamageNumber)
}

func (s Sword) Repair(repairLevel int) int {
	return s.Durability % repairLevel
}

func (s Sword) Upgrade() (int, int) {
	return s.Level + 1, s.DamageNumber + 5
}

func (s Sword) GetName() string {
	return s.Name
}

func (s Sword) GetDamageNumber() int {
	return s.DamageNumber
}

// Factory method for basic, default sword
func NewSword(name string) ISword {
	return Sword{
		Name:         name,
		DamageNumber: 10,
		Level:        1,
		Durability:   60,
	}
}
