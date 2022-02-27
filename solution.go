package square

import "math"

type sides int

const (
	SidesSquare   sides = 4
	SidesTriangle sides = 3
	SidesCircle   sides = 0
)

func CalcSquare(sideLen float64, sidesNum sides) (result float64) {
	switch sidesNum {
	case SidesSquare:
		result = math.Pow(sideLen, 2.0)
	case SidesTriangle:
		result = (math.Pow(sideLen, 2.0) * math.Sqrt(3.0)) / 4
	case SidesCircle:
		result = math.Pi * math.Pow(sideLen, 2.0)
	default:
		result = 0
	}
	return result
}
