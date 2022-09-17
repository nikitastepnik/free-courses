class multifilter():

    def judge_any(self, pos, neg):
        if pos >= 1:
            return True
        else:
            return False

    def __init__(self, a, *funcs, judge=judge_any):
        self.iterable = a
        self.funcs = funcs
        self.judge = judge

    def judge_half(self, pos, neg):
        if pos >= neg:
            return True
        else:
            return False

    def judge_all(self, pos, neg):
        if neg == 0:
            return True
        else:
            return False

    def __iter__(self):
        pos = 0
        neg = 0
        for i in range(0, len(self.iterable)):
            for k in range(0, len(self.funcs)):
                if self.funcs[k](self.iterable[i]):
                    pos += 1
                else:
                    neg += 1
                k += 1
            k = 0
            if self.judge(self, pos, neg):
                pos = 0
                neg = 0
                yield self.iterable[i]
            else:
                pos = 0
                neg = 0
