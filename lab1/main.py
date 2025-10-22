from abc import ABC, abstractmethod
from typing import List


class DatabaseConnection:
    _instance = None
    
    def __new__(cls):
        if cls._instance is None:
            cls._instance = super().__new__(cls)
            cls._instance.connection = "Connected to Database"
            print("Creating new database connection...")
        return cls._instance
    
    def query(self, sql: str):
        return f"Executing: {sql}"


class Computer(ABC):
    def __init__(self):
        self.cpu = ""
        self.ram = ""
        self.storage = ""
    
    @abstractmethod
    def get_specs(self) -> str:
        pass


class GamingComputer(Computer):
    def __init__(self):
        super().__init__()
        self.cpu = "Intel i9"
        self.ram = "32GB DDR5"
        self.storage = "2TB NVMe SSD"
        self.gpu = "RTX 4090"
    
    def get_specs(self) -> str:
        return f"Gaming PC: {self.cpu}, {self.ram}, {self.storage}, {self.gpu}"


class OfficeComputer(Computer):
    def __init__(self):
        super().__init__()
        self.cpu = "Intel i5"
        self.ram = "16GB DDR4"
        self.storage = "512GB SSD"
    
    def get_specs(self) -> str:
        return f"Office PC: {self.cpu}, {self.ram}, {self.storage}"


class ServerComputer(Computer):
    def __init__(self):
        super().__init__()
        self.cpu = "AMD EPYC"
        self.ram = "128GB ECC"
        self.storage = "4TB RAID"
    
    def get_specs(self) -> str:
        return f"Server: {self.cpu}, {self.ram}, {self.storage}"


class ComputerFactory:
    @staticmethod
    def create_computer(computer_type: str) -> Computer:
        if computer_type == "gaming":
            return GamingComputer()
        elif computer_type == "office":
            return OfficeComputer()
        elif computer_type == "server":
            return ServerComputer()
        else:
            raise ValueError(f"Unknown computer type: {computer_type}")


class CustomComputer:
    def __init__(self):
        self.cpu = ""
        self.ram = ""
        self.storage = ""
        self.gpu = ""
        self.case = ""
        self.accessories: List[str] = []
    
    def __str__(self):
        specs = f"Custom Computer:\n"
        specs += f"  CPU: {self.cpu}\n"
        specs += f"  RAM: {self.ram}\n"
        specs += f"  Storage: {self.storage}\n"
        specs += f"  GPU: {self.gpu}\n"
        specs += f"  Case: {self.case}\n"
        if self.accessories:
            specs += f"  Accessories: {', '.join(self.accessories)}"
        return specs


class ComputerBuilder:
    def __init__(self):
        self.computer = CustomComputer()
    
    def set_cpu(self, cpu: str):
        self.computer.cpu = cpu
        return self
    
    def set_ram(self, ram: str):
        self.computer.ram = ram
        return self
    
    def set_storage(self, storage: str):
        self.computer.storage = storage
        return self
    
    def set_gpu(self, gpu: str):
        self.computer.gpu = gpu
        return self
    
    def set_case(self, case: str):
        self.computer.case = case
        return self
    
    def add_accessory(self, accessory: str):
        self.computer.accessories.append(accessory)
        return self
    
    def build(self) -> CustomComputer:
        return self.computer


def main():
    print("=== CREATIONAL DESIGN PATTERNS DEMO ===\n")
    
    print("1. SINGLETON PATTERN - Database Connection")
    db1 = DatabaseConnection()
    db2 = DatabaseConnection()
    print(f"Same instance? {db1 is db2}")
    print(f"{db1.query('SELECT * FROM computers')}\n")
    
    print("2. FACTORY METHOD PATTERN - Computer Factory")
    gaming_pc = ComputerFactory.create_computer("gaming")
    office_pc = ComputerFactory.create_computer("office")
    server = ComputerFactory.create_computer("server")
    
    print(gaming_pc.get_specs())
    print(office_pc.get_specs())
    print(f"{server.get_specs()}\n")
    
    print("3. BUILDER PATTERN - Custom Computer Builder")
    custom_pc = (ComputerBuilder()
                 .set_cpu("AMD Ryzen 9")
                 .set_ram("64GB DDR5")
                 .set_storage("1TB NVMe + 2TB HDD")
                 .set_gpu("RTX 4080")
                 .set_case("Corsair 5000D")
                 .add_accessory("RGB Keyboard")
                 .add_accessory("Gaming Mouse")
                 .add_accessory("Monitor 27\"")
                 .build())
    
    print(custom_pc)


if __name__ == "__main__":
    main()