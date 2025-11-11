import sys
from collections import deque


def main():
    _, x = map(int, input().split())
    experience = list(map(int, input().split()))
    m = int(input())
    deq = deque(experience)
    prefix = [0]
    
    count = 0
    for event in experience:
        if event >= x:
            count += 1
        prefix.append(count)
    
    o = 0
    for _ in range(m):
        event = input().split()
        t = int(event[0])
        
        if t == 1:
            v = int(event[1])
            deq.append(v)
            count = prefix[-1] + (1 if v >= x else 0)
            prefix.append(count)
        elif t == 2:
            deq.popleft()
            o += 1
        elif t == 3:
            k = int(event[1])
            if k == 0:
                print(0)
            else:
                print(prefix[o+k] - prefix[o])


if __name__ == "__main__":
    main()
