#include <stdio.h>
#include <string.h>
#include <stack>

int main() {
    std::stack<int> s;
    int N;
    // freopen("input.txt", "r", stdin);
    scanf("%d", &N);

    char cmd[100]; int value;
    for(int i=0; i<N; i++) {
        scanf("%s", cmd);
        if(strcmp(cmd, "push") == 0) {
            scanf("%d", &value);
            s.push(value);
        } else if(strcmp(cmd, "top") == 0) {
            if(s.size() == 0)   printf("-1\n");
            else                printf("%d\n", s.top());
        } else if(strcmp(cmd, "pop") == 0) {
            if(s.size() == 0)   printf("-1\n");
            else {
                printf("%d\n", s.top());
                s.pop();
            }
        } else if(strcmp(cmd, "empty") == 0) {
            if(s.size() == 0)   printf("1\n");
            else                printf("0\n");
        } else if(strcmp(cmd, "size") == 0) {
            printf("%d\n", (int)s.size());
        }
    }
    
    return 0;
}