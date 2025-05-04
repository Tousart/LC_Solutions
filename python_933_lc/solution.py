# Your RecentCounter object will be instantiated and called as such:

class RecentCounter:
    def __init__(self):
        self.requests = []

    def ping(self, t):
        self.requests.append(t)
        i = 0
        while self.requests[i] < (t - 3000):
            i += 1
        del self.requests[:i]
        return len(self.requests)


obj = RecentCounter()
param_1 = obj.ping(1178)
print(param_1)
param_2 = obj.ping(1178)
print(param_2)
param_3 = obj.ping(1563)
print(param_3)
param_4 = obj.ping(4054)
print(param_4)
param_5 = obj.ping(4152)
print(param_5)