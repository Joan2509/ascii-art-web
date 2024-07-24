package ascii

//liner function calculates mathematically the line number of each rune contained in the input string.
func FindHeadLine(r rune) int {
	nums := int(((r - 32) * 9) + 2)
	return nums
}
func Line2(y rune) int {
	num := int(((y - 32) * 9) + 3)
	return num
}
func Line3(n rune) int {
	nm := int(((n - 32) * 9) + 4)
	return nm
}
func Line4(m rune) int {
	nms := int(((m - 32) * 9) + 5)
	return nms
}
func Line5(t rune) int {
	s := int(((t - 32) * 9) + 6)
	return s
}
func Line6(u rune) int {
	digits := int(((u - 32) * 9) + 7)
	return digits
}
func Line7(r rune) int {
	digit := int(((r - 32) * 9) + 8)
	return digit
}
func Line8(z rune) int {
	numbs := int(((z - 32) * 9) + 9)
	return numbs
}
