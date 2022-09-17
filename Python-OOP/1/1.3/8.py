class Person:
    name = 'Сергей Балакирев'
    job = 'Программист'
    city = 'Москва'


p1 = Person()

if 'job' in p1.__dict__:
    print(True)
else:
    print(False)

