from abc import ABC, abstractmethod


class ComponentFactory(ABC):
    @abstractmethod
    def create_cpu(self):
        pass

    @abstractmethod
    def create_ram(self):
        pass

    @abstractmethod
    def create_storage(self):
        pass


class IntelAMDFactory(ComponentFactory):
    def create_cpu(self):
        return "Intel i7"

    def create_ram(self):
        return "16GB DDR4"

    def create_storage(self):
        return "1TB SSD"


class AMDFactory(ComponentFactory):
    def create_cpu(self):
        return "AMD Ryzen 9"

    def create_ram(self):
        return "32GB DDR5"

    def create_storage(self):
        return "2TB NVMe SSD"


class BudgetFactory(ComponentFactory):
    def create_cpu(self):
        return "Intel i3"

    def create_ram(self):
        return "8GB DDR4"

    def create_storage(self):
        return "256GB SSD"