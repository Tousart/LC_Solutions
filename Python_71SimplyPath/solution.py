"""
:type path: str
:rtype: str
"""


class Solution(object):
    def simplifyPath(self, path):
        result = []
        path_list = path.split('/')

        for elem in path_list:
            if elem == '..' and (len(result) >= 1):
                result.pop()
            elif elem != '..' and elem != '' and elem != '.':
                result.append(elem)

        return '/' + '/'.join(result)


sl = Solution()
path = "/a/./b/../../c/"
path1 = "/../"
print(sl.simplifyPath(path))