package rolls

const forkliftableRoll = '@'

type Roll byte

func New(b byte) Roll {
	return Roll(b)
}

func (r Roll) IsForkliftable() bool {
	return r == forkliftableRoll
}
