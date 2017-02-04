package demo2176

type sup struct {
	indices  []byte
	Position struct {
		X, Y int
	}
}

func _(s, space sup, i, j int) {
	s.indices[i+1] = uint16(j)
	if space.Position.X > (800 - 16) {
		space.Position.X = 400 - 16
	} else if space.Position.Y < 0 {

	}
}

func _(vl ...string) []string {
	if len(vl)%2 != 0 {
		panic("ErrImbalancedPairs")
	}
	vls := make([]string, len(vl)/2)
	j := 0
	for i := 0; i < len(vl); i = i + 2 {
		vls[j] = vl[i] + vl[i+1]
		j++
	}
	return vls
}
