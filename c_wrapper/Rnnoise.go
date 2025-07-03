package c_wrapper

// #cgo LDFLAGS: -L${SRCDIR}/../.libs -lrnnoise -lpthread
// #cgo CFLAGS: -I${SRCDIR}/../include
// #include "rnnoise.h"
import "C"
import (
	"unsafe"
)

const FRAME_SIZE = 480

func ExecuteRnnoise(inputName string, outputName string) {
	st := C.rnnoise_create(nil)
	f1 := C.fopen(C.CString(inputName), C.CString("rb"))
	fout := C.fopen(C.CString(outputName), C.CString("wb"))
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
