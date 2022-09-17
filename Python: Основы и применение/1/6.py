class Buffer:
    def __init__(self):
        self.list = []
        self.len = 0
        self.i = 0

    def add(self, *a):
        sum = 0
        self.list.extend(a)
        if len(self.list) > 4:
            while len(self.list) > 4:
                buffer_head = self.list[0:5]
                buffer_tail = self.list[5:]
                for i in buffer_head:
                    sum = sum + i
                print(sum)
                self.list = buffer_tail
                sum = 0
        else:
            pass

    def get_current_part(self):
        return (self.list)
