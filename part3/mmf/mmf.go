package mmf

//#include "mmf.h"
import "C"
import "unsafe"

// Here i create aliases, in this way I can export the data definition to an external package,
// and define methods on it, making int more readable and hinding some C logic
type Mmf = C.struct_Mmf

func OpenMmf(mmfPath string) (mmf *Mmf) {
	c_mmf_ptr := C.OpenMmf(C.CString(mmfPath))
	mmf_ptr := unsafe.Pointer(c_mmf_ptr)

	// with a casting of a pointer I can export my own alias type
	mmf = (*Mmf)(mmf_ptr)
	return
}

func (mmf *Mmf) GetDevice() (d *Device) {
	d = &mmf.Device
	return
}
