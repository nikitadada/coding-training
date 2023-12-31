package main

import (
	"reflect"
	"unsafe"
)

//func main() {
//	var str = "go"
//	modifyStr(str)
//}

func modifyStr(str string) {
	// Именно эта операция позволяет изменить строку.
	// Мы, конечно, меняем не исходную строку, а новую, так как после конкатенации будет создана новая строка в памяти.
	// Но в любом случае изменение строки получается выполнить.
	// Без конкатенации даже с помощью reflect и unsafe не получится выполнить модификацию, потому что
	// компилятор помечает исходную строку как константу и помещает ее в постоянное адресное пространство,
	// где изменения запрещены.
	str += "go"

	strHeader := (*reflect.StringHeader)(unsafe.Pointer(&str))
	sliceHeader := &reflect.SliceHeader{
		Data: strHeader.Data,
		Len:  strHeader.Len,
		Cap:  strHeader.Len,
	}

	slice := *(*[]byte)(unsafe.Pointer(sliceHeader))
	slice[0] = 'G'
	slice[2] = 'G'
	println(str)
}
