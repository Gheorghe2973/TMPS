import sys
import os

sys.path.insert(0, os.path.join(os.path.dirname(__file__), '..'))

from domain.models import Computer, ComputerBuilder
from domain.adapters import LegacyComponent, LegacyComponentAdapter
from domain.decorators import (
    BaseComputer,
    WarrantyDecorator,
    OverclockingDecorator,
    RGBLightingDecorator,
    CustomPaintDecorator
)
from domain.facades import ComputerAssemblyFacade
from domain.proxies import RealComputerConfiguration, ComputerConfigurationProxy


def demonstrate_adapter_pattern():
    print("=" * 70)
    print("ADAPTER PATTERN - Legacy Component Integration")
    print("=" * 70)
    
    legacy_cpu = LegacyComponent("CPU", "Intel Core i7-9700K 8-core 3.6GHz")
    legacy_ram = LegacyComponent("RAM", "Corsair Vengeance 32GB DDR4-3200")
    legacy_gpu = LegacyComponent("GPU", "NVIDIA GTX 1080 Ti 11GB")
    
    print("\n📦 Integrating legacy components into modern system...\n")
    
    adapted_cpu = LegacyComponentAdapter(legacy_cpu)
    adapted_ram = LegacyComponentAdapter(legacy_ram)
    adapted_gpu = LegacyComponentAdapter(legacy_gpu)
    
    print("Adapted Components:")
    print(f"  • {adapted_cpu}")
    print(f"  • {adapted_ram}")
    print(f"  • {adapted_gpu}")
    
    if all([adapted_cpu.is_compatible(), adapted_ram.is_compatible(), adapted_gpu.is_compatible()]):
        print("\n✓ All legacy components are compatible with modern system!")
        print("  Building computer with legacy components...")
        
        computer = (ComputerBuilder()
                   .set_cpu(adapted_cpu.get_specs())
                   .set_ram(adapted_ram.get_specs())
                   .set_gpu(adapted_gpu.get_specs())
                   .build())
        
        print(f"\n  Result: {computer}")
    
    print()


def demonstrate_decorator_pattern():
    print("=" * 70)
    print("DECORATOR PATTERN - Dynamic Feature Enhancement")
    print("=" * 70)
    
    gaming_pc = (ComputerBuilder()
                .set_cpu("Intel i9")
                .set_ram("32GB DDR5")
                .set_storage("2TB NVMe SSD")
                .set_gpu("RTX 4090")
                .set_case("NZXT H710")
                .build())
    
    print("\n🖥️  Base Configuration:")
    base = BaseComputer(gaming_pc)
    print(f"  Description: {base.get_description()}")
    print(f"  Cost: ${base.get_cost():.2f}")
    
    print("\n➕ Adding 3-year warranty...")
    with_warranty = WarrantyDecorator(base, years=3)
    print(f"  Description: {with_warranty.get_description()}")
    print(f"  Cost: ${with_warranty.get_cost():.2f}")
    
    print("\n➕ Adding professional overclocking...")
    with_overclock = OverclockingDecorator(with_warranty)
    print(f"  Description: {with_overclock.get_description()}")
    print(f"  Cost: ${with_overclock.get_cost():.2f}")
    
    print("\n➕ Adding RGB lighting system...")
    with_rgb = RGBLightingDecorator(with_overclock)
    print(f"  Description: {with_rgb.get_description()}")
    print(f"  Cost: ${with_rgb.get_cost():.2f}")
    
    print("\n➕ Adding custom paint job...")
    final_config = CustomPaintDecorator(with_rgb, "Matte Black")
    print(f"  Description: {final_config.get_description()}")
    print(f"  Cost: ${final_config.get_cost():.2f}")
    
    print(f"\n💰 Price increase: ${final_config.get_cost() - base.get_cost():.2f}")
    print()


def demonstrate_facade_pattern():
    print("=" * 70)
    print("FACADE PATTERN - Simplified System Interface")
    print("=" * 70)
    
    facade = ComputerAssemblyFacade()
    
    print("\n🏪 Computer Assembly Shop - Easy Ordering System\n")
    
    print("📋 Order 1: Gaming PC")
    gaming_pc = (ComputerBuilder()
                .set_cpu("Intel i9")
                .set_ram("32GB DDR5")
                .set_storage("2TB NVMe SSD")
                .set_gpu("RTX 4090")
                .set_case("NZXT H710")
                .build())
    
    order1 = facade.order_computer(gaming_pc)
    print(f"  Order ID: {order1['order_id']}")
    print(f"  Computer: {order1['computer']}")
    print(f"  Price: {order1['total_price']}")
    print(f"  {order1['schedule']}")
    print(f"  {order1['status']}")
    
    print("\n📋 Order 2: Office PC (with compatibility warning)")
    office_pc = (ComputerBuilder()
                .set_cpu("Intel i9")
                .set_ram("16GB DDR4")
                .set_storage("512GB SSD")
                .build())
    
    order2 = facade.order_computer(office_pc)
    print(f"  Order ID: {order2['order_id']}")
    print(f"  Computer: {order2['computer']}")
    print(f"  Price: {order2['total_price']}")
    for msg in order2['validation_messages']:
        print(f"  {msg}")
    print(f"  {order2['schedule']}")
    print(f"  {order2['status']}")
    
    print("\n💵 Quick Quote Request:")
    server = (ComputerBuilder()
             .set_cpu("AMD EPYC")
             .set_ram("128GB ECC")
             .set_storage("4TB RAID")
             .build())
    
    quote = facade.get_quick_quote(server)
    print(f"  Server Configuration Quote: {quote}")
    print()


def demonstrate_proxy_pattern():
    print("=" * 70)
    print("PROXY PATTERN - Secure Configuration Access")
    print("=" * 70)
    
    real_config = RealComputerConfiguration()
    proxy = ComputerConfigurationProxy(real_config)
    
    print("\n🔒 Configuration Management System (Secured)\n")
    
    print("❌ Attempting to access without authentication:")
    try:
        proxy.get_configuration("gaming_setup")
    except PermissionError as e:
        print(f"   {e}")
    
    print("\n🔐 Authenticating user...")
    if proxy.authenticate("john_doe", "admin"):
        print("   ✓ Authentication successful!")
    
    print("\n💾 Saving configurations:")
    
    gaming_pc = (ComputerBuilder()
                .set_cpu("Intel i9")
                .set_ram("32GB DDR5")
                .set_gpu("RTX 4090")
                .build())
    
    office_pc = (ComputerBuilder()
                .set_cpu("Intel i5")
                .set_ram("16GB DDR4")
                .build())
    
    proxy.save_configuration("gaming_setup", gaming_pc)
    print("   ✓ Saved 'gaming_setup'")
    
    proxy.save_configuration("office_setup", office_pc)
    print("   ✓ Saved 'office_setup'")
    
    print("\n📖 Retrieving 'gaming_setup' (first time):")
    config1 = proxy.get_configuration("gaming_setup")
    print(f"   {config1}")
    
    print("\n📖 Retrieving 'gaming_setup' (second time - cached):")
    config2 = proxy.get_configuration("gaming_setup")
    print(f"   {config2}")
    
    print("\n🗑️  Deleting 'office_setup':")
    if proxy.delete_configuration("office_setup"):
        print("   ✓ Configuration deleted")
    
    print("\n📊 Access Log:")
    log = proxy.get_access_log()
    for entry in log[-5:]:
        print(f"   {entry}")
    
    print("\n👋 Logging out...")
    proxy.logout()
    print("   ✓ User logged out")
    
    print("\n❌ Attempting to access after logout:")
    try:
        proxy.get_configuration("gaming_setup")
    except PermissionError as e:
        print(f"   {e}")
    
    print()


def main():
    print("\n")
    print("╔" + "=" * 68 + "╗")
    print("║" + " " * 15 + "STRUCTURAL DESIGN PATTERNS LAB" + " " * 23 + "║")
    print("║" + " " * 20 + "Author: Drumea Vasile" + " " * 27 + "║")
    print("║" + " " * 17 + "Computer Assembly System" + " " * 27 + "║")
    print("╚" + "=" * 68 + "╝")
    print("\n")
    
    demonstrate_adapter_pattern()
    input("Press Enter to continue...\n")
    
    demonstrate_decorator_pattern()
    input("Press Enter to continue...\n")
    
    demonstrate_facade_pattern()
    input("Press Enter to continue...\n")
    
    demonstrate_proxy_pattern()
    
    print("=" * 70)
    print("✓ All structural patterns demonstrated successfully!")
    print("=" * 70)
    print()


if __name__ == "__main__":
    main()