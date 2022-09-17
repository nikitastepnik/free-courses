s = input()
t = input()
c = len(t)
r = 0
for i in range(0, len(s) - c + 1):
    p = s[i:c]
    c += 1
    if (p == t):
        r += 1
print(r)
