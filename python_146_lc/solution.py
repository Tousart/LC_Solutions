"""
Your LRUCache object will be instantiated and called as such:
obj = LRUCache(capacity)

param_1 = obj.get(key)
:type key: int
:rtype: int

obj.put(key,value)
:type key: int
:type value: int
:rtype: None
"""

class LRUCache(object):
    class Node:
        def __init__(self, key, val):
            self.key = key
            self.val = val
            self.next = None
            self.prev = None

    def __init__(self, capacity):
        self.cache = dict()
        self.head = self.Node(0, 0)
        self.tail = self.Node(0, 0)
        self.head.next = self.tail
        self.tail.prev = self.head
        self.cap = capacity

    def add_elem(self, elem):
        self.cache[elem.key] = elem
        elem.prev = self.head
        elem.next = self.head.next
        elem.next.prev = elem
        self.head.next = elem

    def del_elem(self, elem):
        nextt = elem.next
        prevv = elem.prev
        nextt.prev = prevv
        prevv.next = nextt

    def print(self):
        current = self.head
        while current:
            print([current.key, current.val], end=" ")
            current = current.next
        print()

    def get(self, key):
        if key in self.cache:
            elem = self.cache[key]
            self.del_elem(elem)
            self.add_elem(elem)
            self.cache[key] = self.head.next
            return elem.val
        return -1

    def put(self, key, value):
        if key in self.cache:
            elem = self.cache[key]
            del self.cache[key]
            self.del_elem(elem)

        if len(self.cache) == self.cap:
            del self.cache[self.tail.prev.key]
            self.del_elem(self.tail.prev)
        
        self.add_elem(self.Node(key, value))
        self.cache[key] = self.head.next




obj = LRUCache(2)
obj.put(2, 6)
obj.print()
obj.put(1, 5)
obj.print()
obj.put(1, 2)
obj.print()
print(obj.get(1))
obj.print()
print(obj.get(2))
obj.print()
