package graph

type Slope struct {
	Rise, Run int
}

func (s Slope) val() float64 {
	return float64(s.Rise) / float64(s.Run)
}

func CalculateSlope(origin, target Point) Slope {
	return Slope{
		target.Y - origin.Y,
		target.X - origin.X,
	}
}

type direction int

const (
	dirRight direction = iota
	dirRightDown
	dirDown
	dirLeftDown
	dirLeft
	dirLeftUp
	dirUp
	dirRightUp
	dirNone
)

func (d direction) ToString() string {
	switch d {
	case dirRight:
		return "right"
	case dirRightDown:
		return "right & down"
	case dirDown:
		return "down"
	case dirLeftDown:
		return "left & down"
	case dirLeft:
		return "left"
	case dirLeftUp:
		return "left & up"
	case dirUp:
		return "up"
	case dirRightUp:
		return "right & up"
	default:
		return "standstill"
	}
}

func slopeDirection(m Slope) direction {
	if m.Rise == 0 {
		if m.Run > 0 {
			return dirRight
		} else if m.Run < 0 {
			return dirLeft
		} else {
			return dirNone
		}
	} else if m.Rise > 0 {
		if m.Run > 0 {
			return dirRightDown
		} else if m.Run < 0 {
			return dirLeftDown
		} else {
			return dirDown
		}
	} else {
		if m.Run > 0 {
			return dirRightUp
		} else if m.Run < 0 {
			return dirLeftUp
		} else {
			return dirUp
		}
	}
}
