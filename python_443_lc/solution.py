"""
:type chars: List[str]
:rtype: int
"""


class Solution(object):
    def compress(self, chars):
        w, r = 0, 0

        while r < (len(chars)):
            count = 0
            while r < len(chars) and (chars[r] == chars[w]):
                count += 1
                r += 1

            w += 1
            if count > 1:
                for i in str(count):
                    chars[w] = i
                    w += 1
                del chars[w:r]
                r = w
        
        return w, chars


sl = Solution()
print(sl.compress(["a","a","b","b","c","c","c","c"]))