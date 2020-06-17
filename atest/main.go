package main

func main() {

	return

	//fmt.Println(1*(1<<47) + (^0x0a00000000000000+1)&(1<<64-1)*0)

	//fmt.Printf("%b \n", 0x0a00000000000000)
	//fmt.Printf("%b \n", 0x0afffffffffffff)
	//fmt.Printf("%b \n", 1<<47)
	//fmt.Printf("%b \n", -1)

	//fmt.Println(2-1 < 1)
	//m := make(map[bool]bool, 9)
	//m := new(A)
	//m2 := new(uint16)
	//fmt.Println(m, reflect.TypeOf(m).Size())
	//v, _ := strconv.ParseInt("11000000", 2, 64)
	//fmt.Println(strconv.ParseInt("11000000", 2, 64))
	//fmt.Println(Ctz64(uint64(v)))
	//fmt.Println(deBruijnIdx64ctz[33])
	//var m []*uint16
	//for i := uint64(0); i < 1024000; i++ {
	//	m = append(m, new(uint16))
	//	x := uint64(1 << i)
	//	y := x * deBruijn64ctz >> 58
	//	fmt.Printf("%d %d %d\n", i, y, deBruijnIdx64ctz[y])
	//	//fmt.Printf("%b %d %d\n", i, y, Ctz64(uint64(i)))
	//fmt.Printf("%b \n", 9223372036854775807)
	//fmt.Printf("%b \n", 2305843009213693951)
	//}
	//fmt.Println("end")
}

//const deBruijn64ctz = 0x0218a392cd3d5dbf
//
//var deBruijnIdx64ctz = [64]byte{
//	0, 1, 2, 7, 3, 13, 8, 19,
//	4, 25, 14, 28, 9, 34, 20, 40,
//	5, 17, 26, 38, 15, 46, 29, 48,
//	10, 31, 35, 54, 21, 50, 41, 57,
//	63, 6, 12, 18, 24, 27, 33, 39,
//	16, 37, 45, 47, 30, 53, 49, 56,
//	62, 11, 23, 32, 36, 44, 52, 55,
//	61, 22, 43, 51, 60, 42, 59, 58,
//}
//
//func Ctz64(x uint64) int {
//	x &= -x                       // isolate low-order bit
//	y := x * deBruijn64ctz >> 58  // extract part of deBruijn sequence
//	i := int(deBruijnIdx64ctz[y]) // convert to bit index
//	z := int((x - 1) >> 57 & 64)  // adjustment if zero
//	return i + z
//}
