#include <windows.h>

#define SHM_NAME "my_shared_memory.shm"

struct Device {
    int IsActive;
    int SomeValue;
};

struct Shm {
    struct Device Device;
};

struct Shm *OpenShm(char *shmName);
int CloseShm(HANDLE shm);

extern struct Shm *Shm;