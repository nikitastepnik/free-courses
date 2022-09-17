import time


class Loggable:
    def log(self, msg):
        print(str(time.ctime()) + ": " + str(msg))


class LoggableList(list, Loggable):
    def append(self, n):
        super(LoggableList, self).append(n)
        msg = self[-1]
        self.log(msg)
