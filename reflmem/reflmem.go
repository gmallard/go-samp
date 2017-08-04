/*
Reflection and memory layout.
Ref:
https://syslog.ravelin.com/go-and-memory-layout-6ef30c730d51
*/
package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type MyData struct {
	aByte       byte
	anotherByte byte
	aShort      int16
	anInt32     int32
	aSlice      []byte
}

type SliceHeader struct {
	Data uintptr
	Len  int
	Cap  int
}

func main() {
	// First ask Go to give us some information about the MyData type
	typ := reflect.TypeOf(MyData{})
	fmt.Printf("Struct is %d bytes long\n", typ.Size())

	// We can run through the fields in the structure in order
	n := typ.NumField()
	for i := 0; i < n; i++ {
		field := typ.Field(i)
		fmt.Printf("%s at offset %v, size=%d, align=%d\n",
			field.Name, field.Offset, field.Type.Size(),
			field.Type.Align())
	}
	// A real instance of the struct
	data := MyData{
		aByte:   0x1,
		aShort:  0x0203,
		anInt32: 0x04050607,
		aSlice: []byte{
			0x08, 0x09, 0x0a,
		},
	}
	// Show it.
	// 32 is hardcoded based on previous output.
	dataBytes := (*[32]byte)(unsafe.Pointer(&data))
	fmt.Printf("Bytes are %#v\n", dataBytes)
	//
	afield := (unsafe.Pointer(&data.aByte))
	fmt.Printf("aByte data is %#v\n",
		(*[1]byte)(unsafe.Pointer(afield)))
	//
	afield = (unsafe.Pointer(&data.aShort))
	work2 := (*[2]byte)(unsafe.Pointer(afield))
	fmt.Printf("aShort data is %#v, BigEnd: %#v %#v\n",
		work2, work2[1], work2[0])
	//
	afield = (unsafe.Pointer(&data.anInt32))
	work4 := (*[4]byte)(unsafe.Pointer(afield))
	fmt.Printf("anInt32 data is %#v, BigEnd: %#v %#v %#v %#v\n",
		work4, work4[3], work4[2], work4[1], work4[0])
	//
	dataslice := *(*reflect.SliceHeader)(unsafe.Pointer(&data.aSlice))
	work3 := (*[3]byte)(unsafe.Pointer(dataslice.Data))
	fmt.Printf("Slice data is %#v\n",
		work3)
	//
}
