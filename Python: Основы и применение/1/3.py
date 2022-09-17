n, k = map(int, input().split())


def C(n, k):
    if (n != 0) and (k == 0):
        return 1
    if n < k:
        return None
    if n == k:
        return 1
    else:
        return C(n - 1, k) + C(n - 1, k - 1)


res = C(n, k)
print(res)
