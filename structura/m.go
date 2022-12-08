

type Color struct {
	R, G, B uint8
}

type Car struct {
	color  Color
	engine float64
}

func (c Color) getColor() int {
	return int(c.R) + int(c.G) + int(c.B)
}

func (c Car) run(speed float64) (float64, float64) {
	return c.engine, (c.engine * 2) + speed
}

func main() {
	var BMW = Car{
		color:  Color{0, 128, 255},
		engine: 1200,
	}

	fmt.Println(BMW.color.getColor())

	fmt.Println(BMW.run(20))
}