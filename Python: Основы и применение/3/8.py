import re
import sys

c = r"human"
for line in sys.stdin:
    line = line.rstrip()
    if re.search(c, line) is not None:
        line = re.sub(c, "computer", line)
    print(line)
