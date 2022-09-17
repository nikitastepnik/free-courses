from typing import List


class CPU:
    def __init__(self, name, fr):
        self.name = name
        self.fr = fr


class Memory:
    def __init__(self, name, volume):
        self.name = name
        self.volume = volume


class MotherBoard:
    def __init__(self, name, cpu, *slots):
        self.name = name
        self.cpu: CPU = cpu
        self.total_mem_slots = 4
        self.mem_slots: List[Memory] = [elem for elem in slots]

    def get_config(self):
        last_str = 'Память: '
        for i in range(len(self.mem_slots)):
            last_str += '{name} - {volume}'.format(name=self.mem_slots[i].name,
                                                               volume=self.mem_slots[i].volume)
            if i != len(self.mem_slots) - 1:
                last_str += ';'
        return ['Материнская плата: {name_mother_board}'.format(name_mother_board=self.name),
                'Центральный процессор: {name_cpu}, {fr}'.format(name_cpu=self.cpu.name, fr=self.cpu.fr),
                'Слотов памяти: {total_mem_slots}'.format(total_mem_slots=self.total_mem_slots),
                last_str
                ]


cpu = CPU("наименование", "30")
mem1 = Memory("наименование1", "50")
mem2 = Memory("наименование2", "100")

mb = MotherBoard("наименование3", cpu, mem1, mem2)

