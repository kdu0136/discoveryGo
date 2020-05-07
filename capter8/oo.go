package capter8

type Shape interface {
	Area() float32
}

type Square struct {
	Size float32
}

func (s Square) Area() float32 {
	return s.Size * s.Size
}

type Rectangle struct {
	Width, Height float32
}

func (r Rectangle) Area() float32 {
	return r.Height * r.Width
}

type Triangle struct {
	Width, Height float32
}

func (t Triangle) Area() float32 {
	return t.Width * t.Height * 0.5
}

func TotalArea(shapes []Shape) float32 {
	var total float32
	for _, shape := range shapes {
		total += shape.Area()
	}
	return total
}
