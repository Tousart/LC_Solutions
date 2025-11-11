import sys


def main():
    n = int(input())
    series = list(map(int, input().split()))
    coeff = list(map(int, input().split()))

    
    ind = sorted(range(n), key=lambda i: series[i])
    series_s = [series[i] for i in ind]
    coeff_s = [coeff[i] for i in ind]

    p = (sum(coeff)+1)//2
    coeff_power = 0
    
    for i in range(n):
        coeff_power += coeff_s[i]
        if coeff_power >= p:
            episodes = series_s[i]
            break
    
    cost = 0
    for i in range(n):
        cost += abs(episodes - series[i]) * coeff[i]
    
    print(episodes, cost)


if __name__ =="__main__":
    main()
