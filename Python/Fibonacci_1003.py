def main():
    F0 = [0] * 41; # F0[4] = Fibo(4) 가 호출하는 Fibo(0) 횟수
    F1 = [0] * 41; # F1[4] = Fibo(4) 가 호출하는 Fibo(1) 횟수

    F0[0] = F0[2] = F1[1] = F1[2] = 1
    for i in range(3,41):
        F0[i] = F0[i-1] + F0[i-2]
        F1[i] = F1[i-1] + F1[i-2]

    # 입력 처리
    import sys
    T = int(sys.stdin.readline())
    for _ in range(T):
        N = int(sys.stdin.readline())
        print(F0[N],F1[N])


if __name__ == "__main__":
    main()