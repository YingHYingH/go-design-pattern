package _7_decorator

type IDraw interface {
	Draw() string
}

type Square struct {
}

func (s *Square) Draw() string {
	return "Square"
}

type ColorSquare struct {
	square IDraw
	color  string
}

func NewColorSquare(square IDraw, color string) *ColorSquare {
	return &ColorSquare{
		square: square,
		color:  color,
	}
}

func (c *ColorSquare) Draw() string {
	return c.square.Draw() + " with color " + c.color
}
