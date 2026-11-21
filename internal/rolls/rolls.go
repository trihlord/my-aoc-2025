package rolls

const forkliftableRoll = '@'

var precedingRollNeighbourOffsets = [][2]int{{-1, 0}, {-1, -1}, {1, -1}, {0, -1}}

type Roll byte

func (r Roll) IsForkliftable() bool {
	return r == forkliftableRoll
}

type rollNeighbour struct {
	roll  Roll
	rolls []Roll
}

type RollWarehouse struct {
	width          int
	rollNeighbours []rollNeighbour
}

func NewRollWarehouse(width int) *RollWarehouse {
	return &RollWarehouse{width: width}
}

func (rw RollWarehouse) Len() int {
	return len(rw.rollNeighbours)
}

func (rw *RollWarehouse) precedingRollNeighboursAt(i int) []*rollNeighbour {
	var rns []*rollNeighbour
	x, y := i%rw.width, i/rw.width
	for _, o := range precedingRollNeighbourOffsets {
		if ox, oy := x+o[0], y+o[1]; ox >= 0 && ox < rw.width && oy >= 0 {
			if j := ox + oy*rw.width; j >= 0 && j < i {
				rns = append(rns, &rw.rollNeighbours[j])
			}
		}
	}
	return rns
}

func (rw *RollWarehouse) Append(roll Roll) {
	rn := rollNeighbour{roll: roll}
	rw.rollNeighbours = append(rw.rollNeighbours, rn)
	l := len(rw.rollNeighbours) - 1
	for _, prn := range rw.precedingRollNeighboursAt(l) {
		rw.rollNeighbours[l].rolls = append(rw.rollNeighbours[l].rolls, prn.roll)
		prn.rolls = append(prn.rolls, roll)
	}
}

func (rw RollWarehouse) IsForkliftableAt(i int) bool {
	rn := rw.rollNeighbours[i]
	if !rn.roll.IsForkliftable() {
		return false
	}
	c := 0
	for _, r := range rn.rolls {
		if r.IsForkliftable() {
			c++
		}
	}
	return c < 4
}
