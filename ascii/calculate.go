package ascii

func Calculator(line string) (a, b, c, d, e, f, g, h []int) {
	for _, char := range line {
		a = append(a, FindHeadLine(char))
		b = append(b, Line2(char))
		c = append(c, Line3(char))
		d = append(d, Line4(char))
		e = append(e, Line5(char))
		f = append(f, Line6(char))
		g = append(g, Line7(char))
		h = append(h, Line8(char))
	}
	return
}
