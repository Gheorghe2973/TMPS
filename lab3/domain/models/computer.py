from typing import List


class Computer:
    def __init__(self):
        self.cpu = None
        self.ram = None
        self.storage = None
        self.gpu = None
        self.case = None
        self.accessories = []

    def __str__(self):
        parts = []
        if self.cpu:
            parts.append(self.cpu)
        if self.ram:
            parts.append(self.ram)
        if self.storage:
            parts.append(self.storage)
        if self.gpu:
            parts.append(self.gpu)
        if self.case:
            parts.append(f"Case: {self.case}")
        if self.accessories:
            parts.append(f"Accessories: {', '.join(self.accessories)}")
        return ', '.join(parts)


class ComputerBuilder:
    def __init__(self):
        self.computer = Computer()

    def set_cpu(self, cpu):
        self.computer.cpu = cpu
        return self

    def set_ram(self, ram):
        self.computer.ram = ram
        return self

    def set_storage(self, storage):
        self.computer.storage = storage
        return self

    def set_gpu(self, gpu):
        self.computer.gpu = gpu
        return self

    def set_case(self, case):
        self.computer.case = case
        return self

    def add_accessory(self, accessory):
        self.computer.accessories.append(accessory)
        return self

    def build(self):
        return self.computer