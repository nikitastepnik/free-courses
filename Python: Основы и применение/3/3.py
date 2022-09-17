import re
import sys
c = r".*cat.*cat.*"
for line in sys.stdin:
    line = line.rstrip()
    if re.match(c, line) is not None:
        print(line)

