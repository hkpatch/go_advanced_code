#include "number.h"

int number_add_mod(int a, int b, int mod) {
    return (a+b)%mod;
}


// gcc -shared -o libnumber.so number.c