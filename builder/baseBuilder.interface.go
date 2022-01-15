package builder

type IBuilder interface { // Make any type of melee weapon
	SetName(string)
	SetHandleMaterial(string)
	SetHandleSize()
	SetBladeMaterial(string)
	SetBladeSize()
	SetHeadMaterial(string)
	setHeadSize()
	SetHeadType()
	SetDurability(int)
	SetDamage(int)
	GetWeapon() Weapon
}
