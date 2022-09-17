from string import ascii_lowercase, digits

class Base:
    CHARS = "абвгдеёжзийклмнопрстуфхцчшщьыъэюя " + ascii_lowercase
    CHARS_CORRECT = CHARS + CHARS.upper() + digits

    def __init__(self, name, size=10):
        if self.check_name(name):
            self.name = name
        self.size = size

    def get_html(self):
        if isinstance(self, TextInput):
            return f"<p class='login'>{self.name}: <input type='text' size={self.size} />"
        else:
            return f"<p class='password'>{self.name}: <input type='text' size={self.size} />"

    @classmethod
    def check_name(cls, name):
        if not 3 <= len(name) <= 50:
            raise ValueError("некорректное поле name")
        for elem in name:
            if elem not in cls.CHARS_CORRECT:
                raise ValueError("некорректное поле name")

        return True


class TextInput(Base):
    pass


class PasswordInput(Base):
    pass

class FormLogin:
    def __init__(self, lgn, psw):
        self.login = lgn
        self.password = psw

    def render_template(self):
        return "\n".join(['<form action="#">', self.login.get_html(), self.password.get_html(), '</form>'])


# эти строчки не менять
login = FormLogin(TextInput("Логин"), PasswordInput("Пароль"))
html = login.render_template()