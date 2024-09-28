#include "mmf.h"

struct Mmf *Mmf;

struct Mmf *OpenMmf(char *mmfName)
{

    HANDLE fileHandle = CreateFile(mmfName,
                                   GENERIC_WRITE | GENERIC_READ,
                                   FILE_SHARE_READ | FILE_SHARE_WRITE,
                                   NULL,
                                   CREATE_NEW,
                                   FILE_ATTRIBUTE_TEMPORARY,
                                   NULL);

    if (GetLastError() == ERROR_FILE_EXISTS)
    {
        fileHandle = CreateFile(mmfName,
                                GENERIC_WRITE | GENERIC_READ,
                                FILE_SHARE_READ | FILE_SHARE_WRITE,
                                NULL,
                                OPEN_EXISTING,
                                FILE_ATTRIBUTE_TEMPORARY,
                                NULL);
    }

    HANDLE mmfFileMapping = CreateFileMapping(fileHandle, NULL, PAGE_READWRITE, 0, sizeof(struct Mmf), NULL);
    Mmf = (struct Mmf *)MapViewOfFile(mmfFileMapping, FILE_MAP_ALL_ACCESS, 0, 0, 0);

    return Mmf;
}

int CloseMmf(HANDLE mmf)
{
    CloseHandle(mmf);
}