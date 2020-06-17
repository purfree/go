// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !386

// TODO finish intrinsifying 386, deadcode the assembly, remove build tags, merge w/ intrinsics_common
// TODO replace all uses of CtzXX with TrailingZerosXX; they are the same.

package sys

// Using techniques from http://supertech.csail.mit.edu/papers/debruijn.pdf
// 1000011000101000111001001011001101001111010101110110111111
// 根据公式 uint64(1<<i) * deBruijn64ctz >> 58 预生成i=1～63内的值，即deBruijnIdx64ctz
const deBruijn64ctz = 0x0218a392cd3d5dbf

var deBruijnIdx64ctz = [64]byte{
	0, 1, 2, 7, 3, 13, 8, 19,
	4, 25, 14, 28, 9, 34, 20, 40,
	5, 17, 26, 38, 15, 46, 29, 48,
	10, 31, 35, 54, 21, 50, 41, 57,
	63, 6, 12, 18, 24, 27, 33, 39,
	16, 37, 45, 47, 30, 53, 49, 56,
	62, 11, 23, 32, 36, 44, 52, 55,
	61, 22, 43, 51, 60, 42, 59, 58,
}

const deBruijn32ctz = 0x04653adf

var deBruijnIdx32ctz = [32]byte{
	0, 1, 2, 6, 3, 11, 7, 16,
	4, 14, 12, 21, 8, 23, 17, 26,
	31, 5, 10, 15, 13, 20, 22, 25,
	30, 9, 19, 24, 29, 18, 28, 27,
}

// Ctz64 counts trailing (low-order) zeroes,
// and if all are zero, then 64.
func Ctz64(x uint64) int {
	// 这里将x从右向左，第一个1的位置向前的值全部清零
	// 例如x=58 二进制为111010，则结果为2(10)
	// x=48 二进制为110000，则结果为16(10000)
	x &= -x // isolate low-order bit
	// 德布鲁因序列(De Bruijn sequence)，记为B(k, n)，是 k 元素构成的循环序列。所有长度为 n 的 k 元素构成序列都在它的子序列（以环状形式）中，出现并且仅出现一次。
	// deBruijn64ctz是一个deBruijn序列b(2,6)
	// x是2^n次方，所以x * deBruijn64ctz等于将deBruijn64ctz右移n位
	// 而deBruijn64ctz是一个6位的循环序列，所以右移58位
	y := x * deBruijn64ctz >> 58 // extract part of deBruijn sequence
	// deBruijnIdx64ctz是计算好的deBruijnIdx64ctz指定序列位置值的末尾0数量
	// 因为deBruijnIdx64ctz是64位循环的，只需要计算deBruijn64ctz偏移0～63的值，即deBruijnIdx64ctz，这样硬编码后，就可以快速读取
	i := int(deBruijnIdx64ctz[y]) // convert to bit index
	// 当x小于(1<<57)时，因为正数右移不会为负，0右移仍为0，所有0&64=0
	// 当x大于(1<<57)时，因为最大64位，右移57位后，最高只有7位，但x-1，导致只有6位，而64=1<<6为7位，结果仍为0
	// 当x为0时，x - 1为负数，int(-1&64)=64
	// 所有这里是将0变为64
	z := int((x - 1) >> 57 & 64) // adjustment if zero
	return i + z
}

// Ctz32 counts trailing (low-order) zeroes,
// and if all are zero, then 32.
func Ctz32(x uint32) int {
	x &= -x                       // isolate low-order bit
	y := x * deBruijn32ctz >> 27  // extract part of deBruijn sequence
	i := int(deBruijnIdx32ctz[y]) // convert to bit index
	z := int((x - 1) >> 26 & 32)  // adjustment if zero
	return i + z
}

// Ctz8 returns the number of trailing zero bits in x; the result is 8 for x == 0.
func Ctz8(x uint8) int {
	return int(ntz8tab[x])
}

// Bswap64 returns its input with byte order reversed
// 0x0102030405060708 -> 0x0807060504030201
func Bswap64(x uint64) uint64 {
	c8 := uint64(0x00ff00ff00ff00ff)
	a := x >> 8 & c8
	b := (x & c8) << 8
	x = a | b
	c16 := uint64(0x0000ffff0000ffff)
	a = x >> 16 & c16
	b = (x & c16) << 16
	x = a | b
	c32 := uint64(0x00000000ffffffff)
	a = x >> 32 & c32
	b = (x & c32) << 32
	x = a | b
	return x
}

// Bswap32 returns its input with byte order reversed
// 0x01020304 -> 0x04030201
func Bswap32(x uint32) uint32 {
	c8 := uint32(0x00ff00ff)
	a := x >> 8 & c8
	b := (x & c8) << 8
	x = a | b
	c16 := uint32(0x0000ffff)
	a = x >> 16 & c16
	b = (x & c16) << 16
	x = a | b
	return x
}
