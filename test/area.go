package area

type Rectangle struct {
	Length  float64
	Breadth float64
}

const Pie float64 = 3.14

func AreaOfCircle(radius float64) float64 {
	return 3.14 * radius * radius
}

func AreaOfRectangle(rect Rectangle) float64 {
	return rect.Length * rect.Breadth
}
