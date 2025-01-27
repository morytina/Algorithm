def main():
    # 입력 처리
    import sys
    N = int(sys.stdin.readline())
    data = [list(map(str, input().split())) for _ in range(N)]
    stack = [];

    for i in range(0, N):
        if(data[i][0] == "push"):
            stack.append(data[i][1])
        if(data[i][0] == "pop"):
            if(len(stack) == 0): 
                print(-1)
            else:
                print(stack[-1])
                stack.pop()
        if(data[i][0] == "top"):
            if(len(stack) == 0): 
                print(-1)
            else:
                print(stack[-1])
        if(data[i][0] == "size"):
            print(len(stack))
        if(data[i][0] == "empty"):
            if(len(stack) == 0): 
                print(1)
            else:
                print(0)

if __name__ == "__main__":
    main()