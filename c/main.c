#include <stdio.h>
#include <stdint.h>
struct A
{
    int32_t a;
    int64_t b;
};

struct B
{
    int64_t c;
    int32_t d;
    int e;
};

int main()
{
    struct A x = {1, 2};
    printf("x.a = %d, x.b = %ld\n", x.a, x.b);

    struct B *q;
    q = (struct B *)&x;
    q->c = 10;
    q->d = 20;
    printf("x.a = %d, x.b = %ld\n", x.a, x.b);
    return 0;
}