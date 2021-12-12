package factory

type MagicSword struct {
	Sword
	Element         string
	ElementalDamage int
}

func NewMagicSword(name string) ISword {
	return MagicSword{
		Sword: Sword{
			Name:         name,
			DamageNumber: 5,
			Level:        1,
			Durability:   50,
		},
		ElementalDamage: 10,
	}
}
