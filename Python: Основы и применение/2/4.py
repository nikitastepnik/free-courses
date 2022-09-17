from datetime import datetime, timedelta

y, m, d = input().split()

a = datetime(int(y), int(m), int(d))
delta = int(input())
b = timedelta(days=delta)
a = a + b
print(a.year, a.month, a.day)
