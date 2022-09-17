import re

import requests

a = input()
b = input()
k = 0
# str1 = r'<a href="(\w+)">(\d)</a>'
# str1 = r'<a href="https://stepic.org/media/attachments/lesson/(\d+)/sample(\d+).html">(\d)</a>'
str3 = 'https:.*html'
str2 = '<a href="' + b + '">'
r = requests.get(a)
# print(r.text)
if re.search(str3, r.text) is not None:
    p = re.findall(str3, r.text)
    for i in range(0, len(p)):
        #    print(p[i])
        c = requests.get(p[i])
        if (c.status_code == 200) and (re.search(str2, c.text) is not None):
            k = 1
if k == 1:
    print("Yes")
else:
    print("No")
