package shm

//#include "shm.h"
import "C"

func OpenShm(shmPath string) (shm *C.struct_Shm) {
	shm = C.OpenShm(C.CString(shmPath))
	return
}
