package ascii

import "fmt"

func Ascii(text, banner string) (string, string, string, string, string, string, string, string) {
	var a, b, c, d, e, f, g, h []int = Calculator(text)
	ab, bc, cd, de, ef, fg, gh, hi := Concatenator(a, b, c, d, e, f, g, h)
	return ab, bc, cd, de, ef, fg, gh, hi
}

func Combinator(str1, str2, str3, str4, str5, str6, str7, str8 string) (result string, err error) {
	result = fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s", str1, str2, str3, str4, str5, str6, str7, str8)
	return result, nil
}
