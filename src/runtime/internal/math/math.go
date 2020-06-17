// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package math

import "runtime/internal/sys"

const MaxUintptr = ^uintptr(0)

// MulUintptr returns a * b and whether the multiplication overflowed.
// On supported platforms this is an intrinsic lowered by the compiler.
// a: 元素的数量，b: 元素的大小，a*b即需要的内存大小
func MulUintptr(a, b uintptr) (uintptr, bool) {
	// a|b的意义是什么???为什么a|b < 1<<(4*sys.PtrSize)可以检测内存是否溢出???
	// 1<<(4*sys.PtrSize)是考虑32位系统，以最小32位溢出做判断???
	// 猜测b是一个很小的数字，b是bucket的大小，当a最大为(1<<(4*sys.PtrSize)-1)时，a*b约等于a，所有a|b可以近似看成a*b，
	// 因为这里当目的只是检测是否内存溢出，而不是计算a*b的值，乘法运算比较复杂???
	// 1<<(4*sys.PtrSize) 2^32
	if a|b < 1<<(4*sys.PtrSize) || a == 0 {
		return a * b, false
	}
	overflow := b > MaxUintptr/a
	return a * b, overflow
}
