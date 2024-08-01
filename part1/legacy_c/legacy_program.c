#include "../shm/shm.h"
#include <stdio.h>

void main()
{
    struct Shm *shm;
    shm = OpenShm(SHM_NAME);

    shm->Device.IsActive = 0;
    shm->Device.SomeValue = 0;
    printf("Registered to shm at address: %d\n", shm);

    shm->Device.IsActive = 1;
    while (TRUE)
    {
        shm->Device.SomeValue++;
        printf("Device value: %d\n", shm->Device.SomeValue);

        Sleep(1000);
    };
}