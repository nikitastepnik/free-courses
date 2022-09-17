def f1(dict1, k, list2):
    i = 0
    p = 0
    while (dict1.get(k)) is not None and i != len(dict1.get(k)):
        p = dict1.get(k)[i]
        if p in list2:
            pass
        else:
            list2.append(p)
        i += 1
    if p != 0:
        list2 = f1(dict1, p, list2)
    return list2


f = 1
dict1 = {}
list1 = []
list2 = []
k = int(input())
for i in range(0, k):
    a, *b = input().replace(":", " ").split()
    dict1[a] = b
k = int(input())
for i in range(0, k):
    a = input()
    list1.append(a)
for i in range(1, k):
    a = dict1.get(list1[i])
    c = len(a)
    for p in range(0, i):
        if (list1[p] in a) and (f == 1):
            print(list1[i])
            f = 0
    for p in range(0, c):
        k = a[p]
        r = dict1.get(k)
        for h in range(0, i):
            if (list1[h] in r) and (f == 1):
                print(list1[i])
                f = 0
    if f == 1:
        list2 = f1(dict1, list1[i], list2)
        for h in range(0, i):
            if list1[h] in list2:
                print(list1[i])
    f = 1
    list2 = []
