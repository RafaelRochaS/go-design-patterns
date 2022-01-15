package builder

type HammerBuilder struct {
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

func NewHammerBuilder() *HammerBuilder {
	return &HammerBuilder{}
}

func (sb *HammerBuilder) SetName(name string) {
	sb.Name = name
}

func (sb *HammerBuilder) SetHandleMaterial(handleMaterial string) {
	sb.HandleMaterial = handleMaterial
}

func (sb *HammerBuilder) SetHandleSize() {
	sb.HandleSize = MEDIUM
}

func (sb *HammerBuilder) SetBladeMaterial(bladeMaterial string) {
	sb.BladeMaterial = ""
}

func (sb *HammerBuilder) SetBladeSize() {
	sb.BladeSize = ""
}

func (sb *HammerBuilder) SetHeadMaterial(headMaterial string) {
	sb.HeadMaterial = headMaterial
}

func (sb *HammerBuilder) setHeadSize() {
	sb.HeadSize = LARGE
}

func (sb *HammerBuilder) SetHeadType() {
	sb.HeadType = HAMMER
}

func (sb *HammerBuilder) SetDurability(durability int) {
	sb.Durability = durability
}

func (sb *HammerBuilder) SetDamage(damage int) {
	sb.Damage = damage
}

func (sb *HammerBuilder) GetWeapon() Weapon {
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
