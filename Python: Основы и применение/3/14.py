import re

import requests

a = str(input())
r = requests.get(a)
str1 = r'(?:<a.*href=")(?:[\w-]*://)?(\w[\w.-]+)'
p = re.findall(str1, r.text)
set1 = set(p)
list1 = list(set1)
list1.sort()
for i in range(0, len(list1)):
    print(list1[i])
