package builder

import "errors"

type Weapon struct {
	Name           string
	HandleMaterial string
	HandleSize     string
	BladeMaterial  string
	BladeSize      string
	HeadMaterial   string
	HeadSize       string
	HeadType       string
	Durability     int
	Damage         int
}

const (
	SWORD  = "Sword"
	SPEAR  = "Spear"
	AXE    = "Axe"
	HAMMER = "Hammer"
	SMALL  = "Small"
	MEDIUM = "Medium"
	LARGE  = "Large"
	IRON   = "Iron"
	STEEL  = "Steel"
	WOOD   = "Wood"
)

func GetBuilder(builderType string) (IBuilder, error) {
	switch builderType {
	case SWORD:
		return NewSwordBuilder(), nil
	case SPEAR:
		return NewSpearBuilder(), nil
	case AXE:
		return NewAxeBuilder(), nil
	case HAMMER:
		return NewHammerBuilder(), nil
	default:
		return nil, errors.New("builder type not found")
	}
}
