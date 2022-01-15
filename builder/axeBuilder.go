package builder

type AxeBuilder struct {
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

func NewAxeBuilder() *AxeBuilder {
	return &AxeBuilder{}
}

func (sb *AxeBuilder) SetName(name string) {
	sb.Name = name
}

func (sb *AxeBuilder) SetHandleMaterial(handleMaterial string) {
	sb.HandleMaterial = handleMaterial
}

func (sb *AxeBuilder) SetHandleSize() {
	sb.HandleSize = MEDIUM
}

func (sb *AxeBuilder) SetBladeMaterial(bladeMaterial string) {
	sb.BladeMaterial = bladeMaterial
}

func (sb *AxeBuilder) SetBladeSize() {
	sb.BladeSize = MEDIUM
}

func (sb *AxeBuilder) SetHeadMaterial(headMaterial string) {
	sb.HeadMaterial = headMaterial
}

func (sb *AxeBuilder) setHeadSize() {
	sb.HeadSize = MEDIUM
}

func (sb *AxeBuilder) SetHeadType() {
	sb.HeadType = AXE
}

func (sb *AxeBuilder) SetDurability(durability int) {
	sb.Durability = durability
}

func (sb *AxeBuilder) SetDamage(damage int) {
	sb.Damage = damage
}

func (sb *AxeBuilder) GetWeapon() Weapon {
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
