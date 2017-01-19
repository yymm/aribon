from typing import List

Array = List[int]


def q(n: int, a: Array) -> int:
    a.sort(reverse=True)
    for i in range(0, len(a)-2):
        if a[i] < a[i+1] + a[i+2]:
            return a[i] + a[i+1] + a[i+2]
    return 0


n = 5
a = [2, 3, 4, 5, 10]
print(q(n, a))
