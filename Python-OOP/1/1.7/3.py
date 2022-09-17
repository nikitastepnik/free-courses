import re
from string import ascii_lowercase, digits


class CardCheck:
    CHARS_FOR_NAME = ascii_lowercase.upper() + digits

    @classmethod
    def check_card_number(cls, number):

        if re.match("(\d\d\d\d-){3}\d\d\d\d", number) and len(number)==19:
            return True
        else:
            return False

    @classmethod
    def check_name(cls, name):
        name = name.split()

        if len(name) != 2:
            return False

        for item in name:
            for elem in item:
                if elem not in cls.CHARS_FOR_NAME:
                    return False
        return True