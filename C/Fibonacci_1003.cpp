#include <stdio.h>

int main() {
    int F0[41]; // F0[4] = Fibo(4) 가 호출하는 Fibo(0) 횟수
    int F1[41]; // F1[4] = Fibo(4) 가 호출하는 Fibo(1) 횟수

    F0[1] = F1[0] = 0;
    F0[0] = F0[2] = F1[1] = F1[2] = 1;
    for(int i=3; i<=40; i++) {
        F0[i] = F0[i-1] + F0[i-2];            
        F1[i] = F1[i-1] + F1[i-2];
    }

    int T, N;
    scanf("%d", &T);

    for(int i=0; i<T; i++) {
        scanf("%d", &N);
        printf("%d %d\n", F0[N], F1[N]);
    }
}