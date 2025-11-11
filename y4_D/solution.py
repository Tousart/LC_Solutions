import sys


def main():
    n = int(input())
    tbls = list(map(int, input().split()))
    i, j = 0, n-1
    count_V, count_M = tbls[i], tbls[j]
    
    result = abs(count_V - count_M)
    
    l = i+1
    r = j+1
    while i < j-1:
        if count_V <= count_M:
            i += 1
            count_V += tbls[i]
        else:
            j -= 1
            count_M += tbls[j]
        diff = abs(count_V - count_M)
        if diff < result:
            result = diff
            l = i+1
            r = j+1
    
    print(result, l, r)


if __name__ == "__main__":
    main()
