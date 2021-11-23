package shapes

type Rectangle struct {
	Width  float64
	Height float64
}

func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}
