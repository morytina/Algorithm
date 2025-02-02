def main():
    import sys
    N = int(sys.stdin.readline())
    D = [0] * 1000001
    D[2] = D[3] = 1

    for i in range(4,1000001):
        D[i] = D[i-1] + 1
        if(i%2 == 0):
            D[i] = min(D[i], D[int(i/2)]+1)
        if(i%3 == 0):
            D[i] = min(D[i], D[int(i/3)]+1)

    print(D[N])

if __name__ == "__main__":
    main()