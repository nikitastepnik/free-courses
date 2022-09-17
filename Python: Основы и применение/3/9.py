import re
import sys

c = r"\b[Aa]+\b"
for line in sys.stdin:
    line = line.rstrip()
    if re.search(c, line) is not None:
        line = re.sub(c, "argh", line, count=1)
        print(line)
