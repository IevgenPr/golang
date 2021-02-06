package main

import (
	"fmt"
	"math/big"
)

// const (
// 	prefix byte = 1 << 7
// 	mask   byte = prefix - 1
// )

//BigIntToByteArray blah
// func BigIntToByteArray(num *big.Int) (buf []byte) {
// 	var curBits, remainder big.Int
// 	var tmp byte
// 	remainder.Set(num) //make a working copy of num to remainder
// 	zero := big.NewInt(0)
// 	var mask = big.NewInt(0b01111111)
// 	for {
// 		curBits.And(&remainder, mask)
// 		remainder.Rsh(&remainder, 7)
// 		tmp = byte(curBits.Int64())
// 		if remainder.Cmp(zero) <= 0 {
// 			buf = append(buf, tmp)
// 			break
// 		}
// 		buf = append(buf, tmp|1<<7)
// 	}
// 	return buf
// }
func BigIntToBytes(num *big.Int) (bytes []byte) {
	var (
		b        byte
		bts, buf big.Int
		mask     = big.NewInt(0b01111111)
		zero     = big.NewInt(0)
	)
	buf.Set(num) // make num copy to prevent modification
	for {
		bts.And(&buf, mask)
		buf.Rsh(&buf, 7)
		b = byte(bts.Int64())
		if buf.Cmp(zero) == 0 {
			bytes = append(bytes, b)
			break
		}
		bytes = append(bytes, b|1<<7)
	}
	return
}

// ByteArrayToBigInt bytes must be non empty slice
func BytesToBigInt(bytes []byte) *big.Int {
	var (
		bg   = big.NewInt(0)
		bits int64
	)
	for i := len(bytes) - 1; i >= 0; i-- {
		bits = int64(bytes[i] & 0b01111111) // get 7 bits
		bg.Lsh(bg, 7)
		bg.Or(bg, big.NewInt(bits))
	}
	return bg
}

func DecodeBigInts(bytes []byte) (r []*big.Int) {
	buf := make([]byte, 0)
	for _, b := range bytes {
		buf = append(buf, b)
		if b&(1<<7) == 0 {
			r = append(r, BytesToBigInt(buf))
			buf = nil
		}
	}
	return
}

func EncodeBigInts(bigInts []*big.Int) (bytes []byte) {
	for _, bg := range bigInts {
		bytes = append(bytes, BigIntToBytes(bg)...)
	}
	return
}

func main() {

	num := big.NewInt(103400)

	got := BigIntToBytes(num)
	fmt.Println("XXX", got)
	for j, i := range got {
		fmt.Printf("%d >>> %.8b\n", j, i)
	}
	back := BytesToBigInt(got)
	fmt.Println("main bitset :", num.Text(2))
	fmt.Println("bitset back :", back.Text(2))

	i1 := big.NewInt(222222)
	i2 := big.NewInt(555555)
	arr := []*big.Int{i1, i2}
	fmt.Println(arr)
	for _, i := range arr {
		fmt.Println(BigIntToBytes(i))
	}
	x := EncodeBigInts(arr)
	fmt.Println(x)
	y := DecodeBigInts(x)
	fmt.Println(y)

}
