import re
import sys

c = r"(\w)(\w)(\w*)"
for line in sys.stdin:
    s = []
    s1 = ' '
    res = ''
    line = line.rstrip()
    if re.search(c, line) is not None:
        line = line.split()
        for i in line:
            s.append(re.sub(c, r"\2\1\3", i))
        res = s1.join(s)
        print(res)
