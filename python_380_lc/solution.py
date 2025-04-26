"""
Your RandomizedSet object will be instantiated and called as such:
obj = RandomizedSet()
param_1 = obj.insert(val)
:type val: int
:rtype: bool

param_2 = obj.remove(val)
:type val: int
:rtype: bool

param_3 = obj.getRandom()
:rtype: int
"""

from random import choice

class RandomizedSet(object):

    def __init__(self):
        self.elems = {}

    def insert(self, val):
        if val in self.elems:
            return False
        self.elems[val] = 0
        return True
        

    def remove(self, val):
        if self.elems.pop(val, None) is None:
            return False
        return True
        

    def getRandom(self):
        return choice(list(self.elems.keys()))
        


obj = RandomizedSet()
param_1 = obj.insert(1)
print(param_1)
param_11 = obj.insert(1)
print(param_11)
param_111 = obj.insert(2)
print(param_111)
param_2 = obj.remove(1)
print(param_2)
param_22 = obj.remove(1)
print(param_22)
param_3 = obj.getRandom()
print(param_3)