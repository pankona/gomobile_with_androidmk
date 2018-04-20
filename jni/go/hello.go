package hello

/*
#cgo CFLAGS: -I../hello
#cgo arm64 LDFLAGS: ../../obj/local/arm64-v8a/libhello.a
#cgo arm   LDFLAGS: ../../obj/local/armeabi-v7a/libhello.a
#cgo 386   LDFLAGS: ../../obj/local/x86/libhello.a
#cgo amd64 LDFLAGS: ../../obj/local/x86_64/libhello.a
// gomobile bind doesn't support mips/mips64
// #cgo mips LDFLAGS: ../../obj/local/mips/libhello.a
// #cgo mips64 LDFLAGS: ../../obj/local/mips64/libhello.a
#include "hello.h"
*/
import "C"

func Hello() {
	C.print()
}
