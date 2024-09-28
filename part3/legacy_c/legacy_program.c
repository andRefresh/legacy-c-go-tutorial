#include "../mmf/mmf.h"
#include <stdio.h>

void main()
{
    struct Mmf *mmf;
    mmf = OpenMmf(MMF_NAME);

    mmf->Device.SomeValue = 0;
    printf("Registered to memory mapped file at address: %d\n", mmf);

    while (TRUE)
    {
        if (mmf->Device.IsActive) {
            mmf->Device.SomeValue++;
            printf("Device value: %d\n", mmf->Device.SomeValue);
        }

        Sleep(1000);
    };
}