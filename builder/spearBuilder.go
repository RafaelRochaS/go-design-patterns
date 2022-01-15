package builder

type SpearBuilder struct {
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

func NewSpearBuilder() *SpearBuilder {
	return &SpearBuilder{}
}

func (sb *SpearBuilder) SetName(name string) {
	sb.Name = name
}

func (sb *SpearBuilder) SetHandleMaterial(handleMaterial string) {
	sb.HandleMaterial = handleMaterial
}

func (sb *SpearBuilder) SetHandleSize() {
	sb.HandleSize = LARGE
}

func (sb *SpearBuilder) SetBladeMaterial(bladeMaterial string) {
	sb.BladeMaterial = bladeMaterial
}

func (sb *SpearBuilder) SetBladeSize() {
	sb.BladeSize = SMALL
}

func (sb *SpearBuilder) SetHeadMaterial(headMaterial string) {
	sb.HeadMaterial = ""
}

func (sb *SpearBuilder) setHeadSize() {
	sb.HeadSize = ""
}

func (sb *SpearBuilder) SetHeadType() {
	sb.HeadType = ""
}

func (sb *SpearBuilder) SetDurability(durability int) {
	sb.Durability = durability
}

func (sb *SpearBuilder) SetDamage(damage int) {
	sb.Damage = damage
}

func (sb *SpearBuilder) GetWeapon() Weapon {
	return Weapon{
		Name:           sb.Name,
		HandleMaterial: sb.HandleMaterial,
		HandleSize:     sb.HandleSize,
		BladeMaterial:  sb.BladeMaterial,
		BladeSize:      sb.BladeSize,
		HeadMaterial:   sb.HeadMaterial,
		HeadSize:       sb.HeadSize,
		HeadType:       sb.HeadType,
		Durability:     sb.Durability,
		Damage:         sb.Damage,
	}
}
