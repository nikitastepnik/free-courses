import random


class Cell:
    def __init__(self, mine=False, around_mines=0):
        self.around_mines = around_mines
        self.mine = mine
        self.fl_open = False


class GamePole:
    vars = ((-1, -1), (-1, 0), (-1, 1), (0, 1), (1, 1), (1, 0), (1, -1), (0, -1))

    def __init__(self, N, M):
        self.N = N
        self.M = M
        self.init()

    def check_mine(self, idx_str, idx_col):
        for elem in self.vars:
            if idx_str + elem[0] >= 0 and idx_col + elem[1] >= 0:
                try:
                    if self.pole[idx_str + elem[0]][idx_col + elem[1]].mine is True:
                        self.pole[idx_str][idx_col].around_mines += 1
                except:
                    pass

    def init(self):
        self.pole = [[Cell() for i in range(self.M)] for i in range(self.N)]
        for i in range(self.M):
            n = random.randint(0, self.N - 1)
            m = random.randint(0, self.M - 1)
            self.pole[n][m].mine = True
            self.pole[n][m].fl_open = True
        for i in range(self.N):
            for j in range(self.M):
                self.check_mine(i, j)

    def show(self):
        lst = [['0' for i in range(self.M)] for i in range(self.N)]
        for i in range(self.N):
            for j in range(self.M):
                if self.pole[i][j].mine is True:
                    lst[i][j] = "*"
                elif self.pole[i][j].around_mines:
                    lst[i][j] = str(self.pole[i][j].around_mines)
                else:
                    lst[i][j] = "#"

        return lst


pole_game = GamePole(10, 12)
pole_game.init()