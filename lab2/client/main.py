from domain.models.computer import ComputerBuilder, ConfigurationManager
from domain.factory.computer_factory import GamingPCFactory, OfficePCFactory, ServerFactory
from domain.factory.component_factory import AMDFactory


def demonstrate_factory_method():
    print("=== Factory Method Pattern ===")
    
    gaming_factory = GamingPCFactory()
    gaming_pc = gaming_factory.create_computer()
    print(f"Gaming PC: {gaming_pc}")
    
    office_factory = OfficePCFactory()
    office_pc = office_factory.create_computer()
    print(f"Office PC: {office_pc}")
    
    server_factory = ServerFactory()
    server = server_factory.create_computer()
    print(f"Server: {server}")
    print()


def demonstrate_builder():
    print("=== Builder Pattern ===")
    
    custom_computer = (ComputerBuilder()
                       .set_cpu("AMD Ryzen 9")
                       .set_ram("64GB DDR5")
                       .set_storage("1TB NVMe + 2TB HDD")
                       .set_gpu("RTX 4080")
                       .set_case("Corsair 5000D")
                       .add_accessory("RGB Keyboard")
                       .add_accessory("Gaming Mouse")
                       .add_accessory("Monitor 27\"")
                       .build())
    
    print("Custom Computer:")
    print(f"  CPU: {custom_computer.cpu}")
    print(f"  RAM: {custom_computer.ram}")
    print(f"  Storage: {custom_computer.storage}")
    print(f"  GPU: {custom_computer.gpu}")
    print(f"  Case: {custom_computer.case}")
    print(f"  Accessories: {', '.join(custom_computer.accessories)}")
    print()


def demonstrate_singleton():
    print("=== Singleton Pattern ===")
    
    config1 = ConfigurationManager()
    config2 = ConfigurationManager()
    
    print(f"Same instance: {config1 is config2}")
    
    gaming_pc = GamingPCFactory().create_computer()
    config1.add_configuration("Gaming Setup", gaming_pc)
    
    print(f"Configurations: {config2.list_configurations()}")
    print()


def demonstrate_prototype():
    print("=== Prototype Pattern ===")
    
    base_pc = (ComputerBuilder()
               .set_cpu("Intel i7")
               .set_ram("16GB DDR4")
               .set_storage("1TB SSD")
               .build())
    
    cloned_pc1 = base_pc.clone()
    cloned_pc1.gpu = "RTX 4070"
    
    cloned_pc2 = base_pc.clone()
    cloned_pc2.gpu = "RTX 4060"
    
    print(f"Base PC: {base_pc}")
    print(f"Clone 1: {cloned_pc1}")
    print(f"Clone 2: {cloned_pc2}")
    print()


def demonstrate_abstract_factory():
    print("=== Abstract Factory Pattern ===")
    
    amd_factory = AMDFactory()
    
    amd_computer = (ComputerBuilder()
                    .set_cpu(amd_factory.create_cpu())
                    .set_ram(amd_factory.create_ram())
                    .set_storage(amd_factory.create_storage())
                    .build())
    
    print(f"AMD Computer: {amd_computer}")
    print()


if __name__ == "__main__":
    demonstrate_factory_method()
    demonstrate_builder()
    demonstrate_singleton()
    demonstrate_prototype()
    demonstrate_abstract_factory()
