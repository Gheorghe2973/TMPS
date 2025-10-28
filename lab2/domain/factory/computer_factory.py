from abc import ABC, abstractmethod
from domain.models.computer import ComputerBuilder


class ComputerFactory(ABC):
    @abstractmethod
    def create_computer(self):
        pass


class GamingPCFactory(ComputerFactory):
    def create_computer(self):
        return (ComputerBuilder()
                .set_cpu("Intel i9")
                .set_ram("32GB DDR5")
                .set_storage("2TB NVMe SSD")
                .set_gpu("RTX 4090")
                .build())


class OfficePCFactory(ComputerFactory):
    def create_computer(self):
        return (ComputerBuilder()
                .set_cpu("Intel i5")
                .set_ram("16GB DDR4")
                .set_storage("512GB SSD")
                .build())


class ServerFactory(ComputerFactory):
    def create_computer(self):
        return (ComputerBuilder()
                .set_cpu("AMD EPYC")
                .set_ram("128GB ECC")
                .set_storage("4TB RAID")
                .build())