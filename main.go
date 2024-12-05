package main

// #cgo LDFLAGS: -L${SRCDIR}/.libs -lrnnoise -lpthread
// #cgo CFLAGS: -I/home/hieu/Desktop/rnnoise_go_demo/include
// #include "include/rnnoise.h"
import "C"
import (
	"os"
	"unsafe"
)

const FRAME_SIZE = 480

func main() {
	st := C.rnnoise_create(nil)
	fout := C.fopen(C.CString(os.Args[1:][0]), C.CString("wb"))
	f1 := C.fopen(C.CString(os.Args[2:][0]), C.CString("rb"))
	x := make([]C.float, FRAME_SIZE)
	first := 1
	sizeOfShort := 2
	for {
		tmp := make([]C.short, FRAME_SIZE)
		C.fread(unsafe.Pointer(&tmp[0]), C.size_t(sizeOfShort), C.size_t(FRAME_SIZE), f1)
		if C.feof(f1) != C.int(0) {
			break
		}
		for i := 0; i < FRAME_SIZE; i++ {
			x[i] = C.float(tmp[i])
		}
		C.rnnoise_process_frame(st, (*C.float)(unsafe.Pointer(&x[0])), (*C.float)(unsafe.Pointer(&x[0])))
		for i := 0; i < FRAME_SIZE; i++ {
			tmp[i] = C.short(x[i])
		}
		if first == 0 {
			// fmt.Print("OK")
			C.fwrite(unsafe.Pointer(&tmp[0]), C.size_t(sizeOfShort), C.size_t(FRAME_SIZE), fout)
		}
		first = 0
	}

	C.rnnoise_destroy(st)
	C.fclose(f1)
	C.fclose(fout)
}
