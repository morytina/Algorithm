#include <stdio.h>

int main() {
    int N, ans, num;
    bool isPrime[1001];
    for(int i = 2; i < 1001; i++) isPrime[i] = true;
    for(int i = 2; i < 1001; i++) {
        if(isPrime[i]) {
            int cur = i << 1;
            while(cur < 1001) {
                isPrime[cur] = false;
                cur += i;
            }
        }
    }

    scanf("%d", &N);
    ans = 0;
    for(int i=0; i < N; i++) {
        scanf("%d", &num);
        if(isPrime[num]) ans++;
    }

    printf("%d\n", ans);
    return 0;
}