import sys


def main():

    def find_tax(p):
        l, r = 0, n-1
        while l < r:
            mid = (l+r+1) // 2
            if table[mid][0] < p:
                l = mid
            else:
                r = mid-1
        return table[l][1]


    n = int(input())
    table = []

    for _ in range(n):
        p, t = map(int, input().split())
        table.append([p, t])

    m = int(input())
    cars = [int(input()) for _ in range(m)]

    for power in cars:
        print(find_tax(power) * power)


if __name__ == "__main__":
    main()
