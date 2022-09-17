s = input()
a = input()
b = input()

i = 0
while a in s:
    s = s.replace(a, b)
    if i < 1001:
        i += 1
    else:
        break
if i < 1001:
    print(i)
else:
    print("Impossible")
