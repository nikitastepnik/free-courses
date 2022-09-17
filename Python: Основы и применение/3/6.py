import re
import sys

c = r"\\"
for line in sys.stdin:
    line = line.rstrip()
    if re.search(c, line) is not None:
        print(line)
