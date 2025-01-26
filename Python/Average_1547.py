def main():
    # 입력 처리
    import sys
    N = int(sys.stdin.readline())
    score = list(map(float, input().split()))
    max = sum = 0
    for i in range(0, N):
        if score[i] > max:
            max = score[i]

    for i in range(0, N):
        score[i] = score[i] / max * 100
        sum += score[i]

    print(sum / N)

if __name__ == "__main__":
    main()