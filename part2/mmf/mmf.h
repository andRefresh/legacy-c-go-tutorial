#include <windows.h>

#define MMF_NAME "memory_mapped_file.mmf"

struct Device {
    int IsActive;
    int SomeValue;
};

struct Mmf {
    struct Device Device;
};

struct Mmf *OpenMmf(char *mmfName);
int CloseMmf(HANDLE mmf);


extern struct Mmf *Mmf;