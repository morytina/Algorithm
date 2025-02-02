#include<stdio.h>

int main() {
    int N;

    int D[1000001];
    D[0] = D[1] = 0;
    D[2] = D[3] = 1;

    for(int i=4; i<=1000000; i++) {
        D[i] = D[i-1] + 1;
        if(i%2 == 0) D[i] = D[i/2]+1 < D[i] ? D[i/2]+1 : D[i];
        if(i%3 == 0) D[i] = D[i/3]+1 < D[i] ? D[i/3]+1 : D[i];
    }

    scanf("%d", &N);
    printf("%d\n", D[N]);

    return 0;
}