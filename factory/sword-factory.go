package factory

import "errors"

func CreateSword(swordType string, name string) (ISword, error) {
	switch swordType {
	case SWORD:
		return NewSword(name), nil
	case MAGIC:
		return NewMagicSword(name), nil
	default:
		return nil, errors.New("sword type not found")
	}
}

const (
	SWORD = "sword"
	MAGIC = "magic"
)
