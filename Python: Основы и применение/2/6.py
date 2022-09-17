from math import sqrt


def primes():
    i = 1
    while True:
        k = 1
        i += 1
        while k <= sqrt(i):
            if i % k == 0 and k != 1:
                break
            k += 1
        if k > sqrt(i):
            yield i
