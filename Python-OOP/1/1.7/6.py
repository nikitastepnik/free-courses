class Message:

    def __init__(self, text):
        self.text = text
        self.fl_like = False


class Viber:
    lst_msgs = []

    @staticmethod
    def add_message(msg):
        """добавление нового сообщения в список сообщений;
        """
        Viber.lst_msgs.append(msg)

    @staticmethod
    def remove_message(msg):
        """удаление сообщения из списка;
        """
        Viber.lst_msgs.remove(msg)

    @staticmethod
    def set_like(msg: Message):
        """поставить/убрать лайк для сообщения msg
        (если лайка нет то он ставится, если уже есть, то убирается);
        """
        if not msg.fl_like:
            msg.fl_like = True
        else:
            msg.fl_like = False

    @staticmethod
    def show_last_message(count):
        """отображение последних сообщений;
        """
        return Viber.lst_msgs[:count]

    @staticmethod
    def total_messages():
        """Возвращает общее число сообщений.
        """
        return len(Viber.lst_msgs)





