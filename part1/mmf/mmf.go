package mmf

//#include "mmf.h"
import "C"

func OpenMmf(mmfPath string) (mmf *C.struct_Mmf) {
	mmf = C.OpenMmf(C.CString(mmfPath))
	return
}
