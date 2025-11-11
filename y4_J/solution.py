from fractions import Fraction
import sys

def main():
    N, L, W = map(int, input().split())
    models = [tuple(map(int, input().split())) for _ in range(N)]

    times_map = {}


    def add_time(t, idx):
        if t is not None and t > 0:
            times_map.setdefault(t, set()).add(idx)


    # бортики и финиш
    for i, (x, y, vx, vy) in enumerate(models):
        if vy > 0:
            t = Fraction(W - y, vy)
            add_time(t, i)
        elif vy < 0:
            t = Fraction(-y, vy)
            add_time(t, i)

        if vx > 0:
            t = Fraction(L - x, vx)
            add_time(t, i)

    # столкновения
    for i in range(N):
        x1, y1, vx1, vy1 = models[i]
        for j in range(i + 1, N):
            x2, y2, vx2, vy2 = models[j]
            dx, dy = x1 - x2, y1 - y2
            dvx, dvy = vx1 - vx2, vy1 - vy2

            if dvx == 0 and dvy == 0:
                continue

            t = None
            if dvx == 0:
                if dx == 0 and dvy != 0:
                    t_candidate = Fraction(-dy, dvy)
                    if t_candidate > 0:
                        t = t_candidate
            elif dvy == 0:
                if dy == 0:
                    t_candidate = Fraction(-dx, dvx)
                    if t_candidate > 0:
                        t = t_candidate
            else:
                tx = Fraction(-dx, dvx)
                ty = Fraction(-dy, dvy)
                if tx == ty and tx > 0:
                    t = tx

            if t is not None:
                add_time(t, i)
                add_time(t, j)

    times = sorted(times_map.keys())
    alive = [True] * N
    finished_time = [None] * N

    for t in times:
        idxs = [i for i in times_map[t] if alive[i]]
        if not idxs:
            continue

        positions = {}
        for i in idxs:
            x, y, vx, vy = models[i]
            x_at = Fraction(x) + Fraction(vx) * t
            y_at = Fraction(y) + Fraction(vy) * t
            positions.setdefault((x_at, y_at), []).append(i)

        to_remove = set()
        to_finish = set()

        for (x_at, y_at), group in positions.items():
            if len(group) >= 2:
                to_remove.update(group)
                continue

            m = group[0]
            if y_at == 0 or y_at == W:
                to_remove.add(m)
            elif x_at == L:
                to_finish.add(m)

        for m in to_remove:
            alive[m] = False
            finished_time[m] = None

        for m in to_finish:
            if alive[m]:
                finished_time[m] = t
                alive[m] = False

    finishers = [(i, finished_time[i]) for i in range(N) if finished_time[i] is not None]
    if not finishers:
        print(0)
        return

    min_t = min(t for _, t in finishers)
    res = [i + 1 for i, t in finishers if t == min_t]
    print(len(res))
    print(*sorted(res))


if __name__ == "__main__":
    main()

# from fractions import Fraction
# import sys

# def new_situation(situations, t, loc, typ, idx):
#     if t is None:
#         return
#     d = situations.setdefault(t, {})
#     d.setdefault(loc, []).append((typ, idx))


# def main():
#     N, L, W = map(int, input().split())
#     models = [tuple(map(int, input().split())) for _ in range(N)]
#     situations = {}

#     # бортики и финиш
#     for i, (x, y, vx, vy) in enumerate(models):
#         t_border = None
#         if vy > 0:
#             t = Fraction(W - y, vy)
#             if t > 0:
#                 t_border = t
#         elif vy < 0:
#             t = Fraction(-y, vy)
#             if t > 0:
#                 t_border = t
#         if t_border is not None:
#             x_at = Fraction(x) + Fraction(vx) * t_border
#             y_at = Fraction(y) + Fraction(vy) * t_border
#             new_situation(situations, t_border, (x_at, y_at), 'border', i)

#         if vx > 0:
#             t_finish = Fraction(L - x, vx)
#             if t_finish > 0:
#                 y_at = Fraction(y) + Fraction(vy) * t_finish
#                 new_situation(situations, t_finish, (Fraction(L), y_at), 'finish', i)

#     for i in range(N):
#         x1, y1, vx1, vy1 = models[i]
#         for j in range(i + 1, N):
#             x2, y2, vx2, vy2 = models[j]
#             dx = x1 - x2
#             dy = y1 - y2
#             dvx = vx1 - vx2
#             dvy = vy1 - vy2

#             if dvx == 0 and dvy == 0:
#                 continue

#             t = None
#             if dvx == 0:
#                 if dx != 0:
#                     continue
#                 if dvy == 0:
#                     continue
#                 t_candidate = Fraction(-dy, dvy)
#                 if t_candidate > 0:
#                     t = t_candidate
#             elif dvy == 0:
#                 if dy != 0:
#                     continue
#                 t_candidate = Fraction(-dx, dvx)
#                 if t_candidate > 0:
#                     t = t_candidate
#             else:
#                 tx = Fraction(-dx, dvx)
#                 ty = Fraction(-dy, dvy)
#                 if tx == ty and tx > 0:
#                     t = tx

#             if t is not None:
#                 x_at = Fraction(x1) + Fraction(vx1) * t
#                 y_at = Fraction(y1) + Fraction(vy1) * t
#                 new_situation(situations, t, (x_at, y_at), 'collision', i)
#                 new_situation(situations, t, (x_at, y_at), 'collision', j)

#     times = sorted(situations.keys())
#     alive = [True] * N
#     finished_time = [None] * N

#     for t in times:
#         loc_dict = situations[t]
#         to_remove = set()
#         to_finish = set()

#         for loc, evs in loc_dict.items():
#             models_here = []
#             types_here = {}
#             for typ, idx in evs:
#                 if alive[idx]:
#                     models_here.append(idx)
#                     types_here.setdefault(idx, []).append(typ)

#             if not models_here:
#                 continue

#             # столкновение
#             if len(models_here) >= 2:
#                 for m in models_here:
#                     to_remove.add(m)
#                 continue

#             m = models_here[0]
#             x_at, y_at = loc

#             if y_at == 0 or y_at == W:
#                 to_remove.add(m)
#                 continue

#             if 'border' in types_here.get(m, []):
#                 to_remove.add(m)
#                 continue

#             if 'finish' in types_here.get(m, []):
#                 to_finish.add(m)

#         for m in to_remove:
#             alive[m] = False
#             finished_time[m] = None

#         for m in to_finish:
#             if alive[m]:
#                 finished_time[m] = t
#                 alive[m] = False

#     finishers = [(i, finished_time[i]) for i in range(N) if finished_time[i] is not None]
#     if not finishers:
#         print(0)
#         return

#     min_t = min(t for _, t in finishers)
#     winners = [i + 1 for i, t in finishers if t == min_t]
#     winners.sort()

#     print(len(winners))
#     print(' '.join(map(str, winners)))


# if __name__ == "__main__":
#     main()
