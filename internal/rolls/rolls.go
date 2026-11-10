package rolls

const forkliftableRoll = '@'

var forkliftableRollOffsets = [][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}

type Roll byte

func (r Roll) IsForkliftable() bool {
	return r == forkliftableRoll
}

type rollNeighbour struct {
	r  Roll
	ns []Roll
}

type RollNeighbours []rollNeighbour

func (rns *RollNeighbours) Append(r Roll) {
	*rns = append(*rns, rollNeighbour{r: r})
}

func (rns *RollNeighbours) IsForkliftableAt(l, i int) bool {
	c, x, y, mx, my := 0, i%l, i/l, l, len(*rns)/l
	if !(*rns)[i].r.IsForkliftable() {
		return false
	}
	for _, o := range forkliftableRollOffsets {
		if ox, oy := x+o[0], y+o[1]; oy >= 0 && oy < my && ox >= 0 && ox < mx && (*rns)[oy*l+ox].r.IsForkliftable() {
			c++
		}
	}
	return c < 4
}
