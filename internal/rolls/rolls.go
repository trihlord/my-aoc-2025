package rolls

const forkliftableRoll = '@'

type Roll byte

func (r Roll) IsForkliftable() bool {
	return r == forkliftableRoll
}

// TODO: rolls struct of array of tupple current roll, next rolls
