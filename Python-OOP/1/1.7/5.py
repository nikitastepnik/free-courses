from typing import List


class Application:
    def __init__(self, name):
        self.name = name
        self.blocked = False


class AppStore:
    """Объявите класс AppStore - интернет-магазин приложений для устройств
    под iOS. В этом классе должны быть реализованы следующие методы:
    """
    lst_appl: List[Application] = []

    def add_application(self, app):
        """добавление нового приложения app в магазин
        """
        self.lst_appl.append(app)

    def remove_application(self, app):
        """удаление приложения app из магазина
        """
        self.lst_appl.remove(app)

    def block_application(self, app: Application):
        """блокировка приложения app (устанавливает локальное
        свойство blocked объекта app в значение True)
        """
        app.blocked = True

    def total_apps(self):
        """возвращает общее число приложений в магазине
        """
        return len(self.lst_appl)





