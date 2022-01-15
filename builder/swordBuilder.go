package builder

type SwordBuilder struct {
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

func NewSwordBuilder() *SwordBuilder {
	return &SwordBuilder{}
}

func (sb *SwordBuilder) SetName(name string) {
	sb.Name = name
}

func (sb *SwordBuilder) SetHandleMaterial(handleMaterial string) {
	sb.HandleMaterial = handleMaterial
}

func (sb *SwordBuilder) SetHandleSize() {
	sb.HandleSize = SMALL
}

func (sb *SwordBuilder) SetBladeMaterial(bladeMaterial string) {
	sb.BladeMaterial = bladeMaterial
}

func (sb *SwordBuilder) SetBladeSize() {
	sb.BladeSize = LARGE
}

func (sb *SwordBuilder) SetHeadMaterial(headMaterial string) {
	sb.HeadMaterial = ""
}

func (sb *SwordBuilder) setHeadSize() {
	sb.HeadSize = ""
}

func (sb *SwordBuilder) SetHeadType() {
	sb.HeadType = ""
}

func (sb *SwordBuilder) SetDurability(durability int) {
	sb.Durability = durability
}

func (sb *SwordBuilder) SetDamage(damage int) {
	sb.Damage = damage
}

func (sb *SwordBuilder) GetWeapon() Weapon {
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
