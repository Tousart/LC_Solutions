"""
:type n: int
:type edges: List[List[int]]
:rtype: int
"""

from collections import deque

class Solution:
    def magnificentSets(self, n, edges):
        # Функция, которая собирает компоненты связного графа и проверяет его на двудольность
        # Если длина (сумма элементов) color будет четной, то граф двудольный (в нем нет циклов нечетной длины)
        def is_bipartite(node, c, component):
            color[node] = c
            print(color)
            component.append(node)
            for neighbor in graph[node]:
                if color[neighbor] == c:
                    return False
                if color[neighbor] == -1 and not is_bipartite(neighbor, 1 - c, component):
                    return False
            return True

        # Для каждого связного графа (в составе несвязного, если граф несвязный соответственно)
        # ищем максимальный путь в ширину, пробегаясь по каждому узлу
        def max_groups_in_component(component):
            max_depth = 0
            for start in component:
                depth = [-1] * n
                q = deque()
                q.append(start)
                depth[start] = 0
                while q:
                    node = q.popleft()
                    for neighbor in graph[node]:
                        if depth[neighbor] == -1:
                            depth[neighbor] = depth[node] + 1
                            max_depth = max(max_depth, depth[neighbor])
                            q.append(neighbor)
            return max_depth + 1

        # Создаем граф компонентов
        graph = [[] for _ in range(n)]
        for u, v in edges:
            graph[u - 1].append(v - 1)
            graph[v - 1].append(u - 1)

        color = [-1] * n
        components = []

        # Собираем односвязные графы и проверяем на двудольность
        for i in range(n):
            if color[i] == -1:
                component = []
                if not is_bipartite(i, 0, component):
                    return -1
                components.append(component)

        total = 0
        for comp in components:
            total += max_groups_in_component(comp)

        return total


sl = Solution()
print(sl.magnificentSets(6, [[1, 2], [1, 4], [1, 5], [2, 3], [2, 6], [4, 6]]))