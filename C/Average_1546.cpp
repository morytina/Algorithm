#include <stdio.h>

int main() {
    int N;
    float max, sum;
    float score[1000];

    scanf("%d", &N);
    max = sum = 0;
    for(int i=0; i<N; i++) {
        scanf("%f", &score[i]);
        if(score[i] > max) max = score[i];
    }

    for(int i=0; i<N; i++) {
        score[i] = score[i] / max * 100;
        sum += score[i];
    }

    printf("%f", sum / N);
    return 0;
}