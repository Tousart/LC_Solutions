import sys


def main():
    n, m, x = map(int, input().split())
    times = []
    
    for _ in range(n):
        a, b, v = map(int, input().split())
        
        if a < b:
            time1 = (x-b)/v
            time2 = (x-a)/v
        else:
            time1 = (a-x)/v
            time2 = (b-x)/v
        
        s = max(0.0, min(time1, time2))
        f = max(time1, time2)
        if f >= 0:
            times.append((s, f))
    
    times.sort()
    merged = []
    for s, f in times:
        if not merged or merged[-1][1] < s:
            merged.append([s, f])
        else:
            merged[-1][1] = max(merged[-1][1], f)
    
    times = list(map(int, input().split()))
    ind_times = sorted([(t, i) for i, t in enumerate(times)])
    
    result = [0.0] * m
    i = 0
    for t, ind in ind_times:
        while i < len(merged) and merged[i][1] < t:
            i += 1
        
        if i < len(merged) and merged[i][0] <= t <= merged[i][1]:
            result[ind] = merged[i][1]
        else:
            result[ind] = t
    
    for i in result:
        print(f"{i:.9f}")


if __name__ == "__main__":
    main()
