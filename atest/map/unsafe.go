package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

// A header for a Go map.
type hmap struct {
	// Note: the format of the hmap is also encoded in cmd/compile/internal/gc/reflect.go.
	// Make sure this stays in sync with the compiler's definition.
	count     int // # live cells == size of map.  Must be first (used by len() builtin)
	flags     uint8
	B         uint8  // log_2 of # of buckets (can hold up to loadFactor * 2^B items)
	noverflow uint16 // approximate number of overflow buckets; see incrnoverflow for details
	hash0     uint32 // hash seed

	buckets    unsafe.Pointer // array of 2^B Buckets. may be nil if count==0.
	oldbuckets unsafe.Pointer // previous bucket array of half the size, non-nil only when growing
	nevacuate  uintptr        // progress counter for evacuation (buckets less than this have been evacuated)
	extra      uintptr        // optional fields

}

func unsafeMap() {
	m := make(map[string]string, 0)
	m["a"] = "a"
	m["b"] = "a"
	m["c"] = "a"
	m["d"] = "a"
	m["e"] = "a"
	m["f"] = "a"
	m["g"] = "a"
	m["h"] = "a"
	m["i"] = "a"

	rm := *(**hmap)(unsafe.Pointer(&m))
	fmt.Printf("%+v\n", rm)

	hm := *(**reflect.Type)(unsafe.Pointer(&m))
	count := **(**int)(unsafe.Pointer(&hm))
	fmt.Println("count: ", count)

	flags := *(*uint8)(unsafe.Pointer((uintptr(unsafe.Pointer(hm)) + uintptr(8))))
	fmt.Println("flags: ", flags)

	B := *(*uint8)(unsafe.Pointer((uintptr(unsafe.Pointer(hm)) + uintptr(9))))
	fmt.Println("B: ", B)

	noverflow := *(*uint16)(unsafe.Pointer((uintptr(unsafe.Pointer(hm)) + uintptr(10))))
	fmt.Println("noverflow: ", noverflow)

	hash0 := *(*uint32)(unsafe.Pointer((uintptr(unsafe.Pointer(hm)) + uintptr(12))))
	fmt.Println("hash0: ", hash0)

	fmt.Println("buckets: ", *(*unsafe.Pointer)(unsafe.Pointer((uintptr(unsafe.Pointer(hm)) + uintptr(16)))))

	fmt.Println("oldbuckets: ", *(*unsafe.Pointer)(unsafe.Pointer((uintptr(unsafe.Pointer(hm)) + uintptr(24)))))

	fmt.Println("nevacuate: ", *(*uintptr)(unsafe.Pointer((uintptr(unsafe.Pointer(hm)) + uintptr(32)))))

	fmt.Println("extra: ", *(*uintptr)(unsafe.Pointer((uintptr(unsafe.Pointer(hm)) + uintptr(40)))))
}
