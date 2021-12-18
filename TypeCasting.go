package main
import (
	"fmt"; 
	"strconv"
) 
/*
data types in go:
1. Bool
2. String
3. Int, Int8, Int16, Int32, Int64
3. Uint, Uint8, Uint16, Uint32, Uint64, Uintptr
byte // alias for uint8
rune // alias for int32
	// represents a Unicode code point
4. Float32, Float64
5. complex64, complex128

*/

func main() {
	var mystring = "1"
	var x = 10
	var f float32 = 2.1
	fmt.Println(mystring, x, f)
	number, _ := strconv.Atoi(mystring)
	fmt.Println(number)
	x = number
	fmt.Printf("%T", f)
	fmt.Printf( "%v", x)
}