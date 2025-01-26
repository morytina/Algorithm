def make_prime():
    is_prime = [True] * 1001  # 2부터 1000까지의 소수 여부를 저장할 리스트
    for i in range(2, 1001):
        if is_prime[i]:
            cur = i * 2
            while cur < 1001:
                is_prime[cur] = False
                cur += i
    return is_prime


def main():
    is_prime = make_prime()
    
    # 입력 처리
    import sys
    input = sys.stdin.read
    data = input().splitlines()
    
    N = int(data[0])  # 첫 번째 줄: N 값
    numbers = map(int, data[1].split())  # 두 번째 줄: 숫자 리스트
    
    # 소수 개수 카운트
    ans = sum(1 for num in numbers if is_prime[num])
    print(ans)


if __name__ == "__main__":
    main()
