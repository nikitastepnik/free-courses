import re
import sys

c = r"\b([0-9a-zA-Z]+)(\1)\b"
for line in sys.stdin:
    line = line.rstrip()
    if re.search(c, line) is not None:
        print(line)
