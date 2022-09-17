import re
import sys

c = r"z.{3}z"
for line in sys.stdin:
    line = line.rstrip()
    if re.search(c, line) is not None:
        print(line)
