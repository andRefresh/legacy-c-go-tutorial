#include "shm.h"

struct Shm *Shm;

struct Shm *OpenShm(char *shmName)
{

        HANDLE fileHandle = CreateFile(shmName,
                                       GENERIC_WRITE | GENERIC_READ,
                                       FILE_SHARE_READ | FILE_SHARE_WRITE,
                                       NULL,
                                       CREATE_NEW,
                                       FILE_ATTRIBUTE_TEMPORARY,
                                       NULL);

        if (GetLastError() == ERROR_FILE_EXISTS) {
            fileHandle = CreateFile(shmName,
                                       GENERIC_WRITE | GENERIC_READ,
                                       FILE_SHARE_READ | FILE_SHARE_WRITE,
                                       NULL,
                                       OPEN_EXISTING,
                                       FILE_ATTRIBUTE_TEMPORARY,
                                       NULL);
        }

        HANDLE shmFileMapping = CreateFileMapping(fileHandle, NULL, PAGE_READWRITE, 0, sizeof(struct Shm), NULL);
        Shm = (struct Shm *)MapViewOfFile(shmFileMapping, FILE_MAP_ALL_ACCESS, 0, 0, 0);

    return Shm;
}

int CloseShm(HANDLE shm)
{
    CloseHandle(shm);
}