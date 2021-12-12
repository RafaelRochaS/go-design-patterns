package factory

type ISword interface {
	Damage(string) string
	Repair(int) int
	Upgrade() (int, int)
	GetName() string
	GetDamageNumber() int
}
