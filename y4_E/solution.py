import sys


def main():
    n, m, k = map(int, input().split())
    holes = list(map(int, input().split()))
    
    diff = [0] * (n+2)
    for _ in range(m):
        l, r = map(int, input().split())
        diff[l] += 1
        diff[r+1] -= 1
    
    current = 0
    count = [0] * n
    for i in range(n):
        current += diff[i+1]
        count[i] = current

    comfort = 0
    for i in range(n):
        comfort += holes[i] * count[i]

    ind = list(range(n))
    ind.sort(key=lambda i: count[i], reverse=True)
    for i in ind:
        if k == 0 or count[i] == 0:
            break
        
        rep = min(holes[i], k)
        comfort -= rep * count[i]
        k -= rep
    
    print(comfort)


if __name__ == "__main__":
    main()
