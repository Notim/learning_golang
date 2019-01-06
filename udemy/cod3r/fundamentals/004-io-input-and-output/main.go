package main

import "fmt"

func main() {
	var in1 int64
	fmt.Println("it is a print line ftm.Println()")

	fmt.Println("write a number, example: 8984")
	_, err := fmt.Scanf("%d", &in1)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("----------your values examples-----------")
	fmt.Printf("Value ==> %d\n", in1)
	fmt.Printf("Type ==> %T\n", in1)
	fmt.Printf("binary value ==> %b\n", in1)
	fmt.Printf("hexadecimal value ==> %x\n", in1)
	fmt.Printf("unicode character representantion ==> %c\n", in1)
	fmt.Printf("unicode character representantion with quotes ==> %q\n", in1)
	fmt.Printf("unicode value representantion ==> %U\n", in1)

	fmt.Println("----------another examples-----------")
	fmt.Printf("for boolean types use %%t: %t\n", true)
	fmt.Printf("for float types use %%f: %f\n", 5.1)
	fmt.Printf("for float types with scientific notation use %%e: %e\n", 5.1)
	fmt.Printf("for float types with large exponents use %%g: %g\n", 5.1)
	fmt.Printf("for float types with especific format use %%.10f: %.2f\n", 5.1)
	fmt.Printf("for string types use %%s: %s\n", "Hello Worls")
	fmt.Printf("for string types with hexadecimals use %%X: %X\n", "Hello Worls")

	/*
	   --- float types ---
	   %e - scientific notation, e.g. -1.234456e+78
	   %E - scientific notation, e.g. -1.234456E+78
	   %f - decimal point but no exponent, e.g. 123.456
	   %F - synonym for %f
	   %g - %e for large exponents, %f otherwise. Precision is discussed below.
	   %G - %E for large exponents, %F otherwise

	   ---bool--
	   %t - show boolean types

	   %T - show the var type

	   %d - integer values or base 10
	   %b - binary representation (base 2)
	   %o - octal representation (base 8)
	   %x - hexadecimal representation (base 16)
	   %X - hexadecimal representation with uppercase (base 16)

	   %c - show unicode character representation
	   %q - show unicode character with quotes representation
	   %U - show unicode value representation like U+8984

	   --- string ---
	   %s	the uninterpreted bytes of the string or slice
	   %x	base 16, lower-case, two characters per byte
	   %X	base 16, upper-case, two characters per byte

	   --- slice ---
	   %p - address of 0th element in base 16 notation, with leading 0x

	*/

	fmt.Println("<Notim_ 2019>")
}
