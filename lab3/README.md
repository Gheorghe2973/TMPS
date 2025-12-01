# Computer Assembly - Structural Design Patterns Report

**Course:** Software Design Patterns
**Topic:** Structural Design Patterns

a---

## Objectives

1. Study and understand Structural Design Patterns
2. Extend previous laboratory work with new functionalities
3. Implement at least 3 structural design patterns in the project
4. Maintain proper project structure organized by responsibilities

---

## Domain Area

**Computer Assembly System** - An extended system for creating, configuring, and managing computer builds. The system now includes legacy component integration, dynamic feature enhancement, simplified ordering interface, and secure configuration management.

---

## Implementation

### 1. Adapter Pattern

**Purpose:** Converts the interface of a class into another interface clients expect. Adapter lets classes work together that couldn't otherwise because of incompatible interfaces.

**Location:** `domain/adapters/legacy_adapter.py`

**Code:**

```python
from abc import ABC, abstractmethod


class LegacyComponent:
    def __init__(self, component_type: str, specifications: str):
        self.component_type = component_type
        self.specifications = specifications
  
    def get_component_info(self) -> str:
        return f"{self.component_type}: {self.specifications}"


class ModernComponentInterface(ABC):
    @abstractmethod
    def get_name(self) -> str:
        pass
  
    @abstractmethod
    def get_specs(self) -> str:
        pass
  
    @abstractmethod
    def is_compatible(self) -> bool:
        pass


class LegacyComponentAdapter(ModernComponentInterface):
    def __init__(self, legacy_component: LegacyComponent):
        self._legacy_component = legacy_component
        self._compatibility_check()
  
    def get_name(self) -> str:
        return self._legacy_component.component_type
  
    def get_specs(self) -> str:
        return self._legacy_component.specifications
  
    def is_compatible(self) -> bool:
        return self._compatible
  
    def _compatibility_check(self):
        self._compatible = bool(self._legacy_component.specifications)
  
    def __str__(self):
        status = "✓ Compatible" if self._compatible else "✗ Incompatible"
        return f"{self.get_name()}: {self.get_specs()} [{status}]"
```

**Explanation:**
The Adapter Pattern bridges the gap between legacy and modern component systems. Legacy components use the `get_component_info()` method, while the modern system expects `get_name()`, `get_specs()`, and `is_compatible()` methods. The `LegacyComponentAdapter` wraps legacy components and translates their interface to match modern expectations. This allows seamless integration of older components without modifying either the legacy code or the modern system, following the Open/Closed Principle.

**Usage:**

```python
legacy_cpu = LegacyComponent("CPU", "Intel Core i7-9700K 8-core 3.6GHz")
adapted_cpu = LegacyComponentAdapter(legacy_cpu)

computer = (ComputerBuilder()
           .set_cpu(adapted_cpu.get_specs())
           .build())

print(adapted_cpu.is_compatible())  # Output: True
# Output: CPU: Intel Core i7-9700K 8-core 3.6GHz [✓ Compatible]
```

---

### 2. Decorator Pattern

**Purpose:** Attaches additional responsibilities to an object dynamically. Decorators provide a flexible alternative to subclassing for extending functionality.

**Location:** `domain/decorators/computer_decorator.py`

**Code:**

```python
from abc import ABC, abstractmethod
from domain.models.computer import Computer


class ComputerComponent(ABC):
    @abstractmethod
    def get_description(self) -> str:
        pass
  
    @abstractmethod
    def get_cost(self) -> float:
        pass


class BaseComputer(ComputerComponent):
    def __init__(self, computer: Computer):
        self._computer = computer
        self._base_cost = 1000.0
  
    def get_description(self) -> str:
        return str(self._computer)
  
    def get_cost(self) -> float:
        return self._base_cost


class ComputerDecorator(ComputerComponent):
    def __init__(self, computer_component: ComputerComponent):
        self._component = computer_component
  
    @abstractmethod
    def get_description(self) -> str:
        pass
  
    @abstractmethod
    def get_cost(self) -> float:
        pass


class WarrantyDecorator(ComputerDecorator):
    def __init__(self, computer_component: ComputerComponent, years: int = 3):
        super().__init__(computer_component)
        self._years = years
        self._warranty_cost = years * 50.0
  
    def get_description(self) -> str:
        return f"{self._component.get_description()} + {self._years}Y Warranty"
  
    def get_cost(self) -> float:
        return self._component.get_cost() + self._warranty_cost


class OverclockingDecorator(ComputerDecorator):
    def __init__(self, computer_component: ComputerComponent):
        super().__init__(computer_component)
        self._overclock_cost = 200.0
  
    def get_description(self) -> str:
        return f"{self._component.get_description()} + Professional Overclocking"
  
    def get_cost(self) -> float:
        return self._component.get_cost() + self._overclock_cost


class RGBLightingDecorator(ComputerDecorator):
    def __init__(self, computer_component: ComputerComponent):
        super().__init__(computer_component)
        self._rgb_cost = 150.0
  
    def get_description(self) -> str:
        return f"{self._component.get_description()} + RGB Lighting System"
  
    def get_cost(self) -> float:
        return self._component.get_cost() + self._rgb_cost


class CustomPaintDecorator(ComputerDecorator):
    def __init__(self, computer_component: ComputerComponent, color: str):
        super().__init__(computer_component)
        self._color = color
        self._paint_cost = 300.0
  
    def get_description(self) -> str:
        return f"{self._component.get_description()} + Custom {self._color} Paint"
  
    def get_cost(self) -> float:
        return self._component.get_cost() + self._paint_cost
```

**Explanation:**
The Decorator Pattern allows adding features to computers dynamically without modifying the original `Computer` class. Each decorator (Warranty, Overclocking, RGB, CustomPaint) wraps a `ComputerComponent` and adds its own functionality while delegating to the wrapped component. Decorators can be stacked in any combination, providing extreme flexibility. This approach avoids creating separate subclasses for every possible combination of features (which would result in a class explosion) and follows the Single Responsibility Principle by keeping each feature in its own class.

**Usage:**

```python
gaming_pc = (ComputerBuilder()
            .set_cpu("Intel i9")
            .set_ram("32GB DDR5")
            .set_gpu("RTX 4090")
            .build())

base = BaseComputer(gaming_pc)
with_warranty = WarrantyDecorator(base, years=3)
with_rgb = RGBLightingDecorator(with_warranty)
final = CustomPaintDecorator(with_rgb, "Matte Black")

print(final.get_cost())  # Output: $1800.00
# Base: $1000 + Warranty: $150 + RGB: $150 + Paint: $300
```

---

### 3. Facade Pattern

**Purpose:** Provides a unified interface to a set of interfaces in a subsystem. Facade defines a higher-level interface that makes the subsystem easier to use.

**Location:** `domain/facades/assembly_facade.py` and `domain/facades/subsystems.py`

**Code:**

```python
# subsystems.py
from domain.models.computer import Computer


class ComponentValidator:
    def validate_cpu_ram_compatibility(self, cpu: str, ram: str) -> bool:
        ddr5_cpus = ["Intel i9", "AMD Ryzen 9", "Intel i7 13th"]
        if any(cpu_name in cpu for cpu_name in ddr5_cpus):
            return "DDR5" in ram
        return "DDR4" in ram
  
    def validate_gpu_power(self, gpu: str, ram: str) -> bool:
        high_end_gpus = ["RTX 4090", "RTX 4080", "RTX 4070"]
        if any(gpu_name in gpu for gpu_name in high_end_gpus):
            ram_size = int(''.join(filter(str.isdigit, ram)))
            return ram_size >= 32
        return True


class PriceCalculator:
    def __init__(self):
        self._prices = {
            "Intel i9": 500, "Intel i7": 350, "Intel i5": 200,
            "AMD Ryzen 9": 450, "AMD EPYC": 800,
            "RTX 4090": 1600, "RTX 4080": 1200, "RTX 4070": 600,
            "32GB DDR5": 200, "16GB DDR4": 80, "64GB DDR5": 400,
            "128GB ECC": 800, "2TB NVMe SSD": 200, "1TB SSD": 100,
            "512GB SSD": 60, "4TB RAID": 600
        }
  
    def calculate_total_price(self, computer: Computer) -> float:
        total = 0.0
        if computer.cpu:
            total += self.calculate_component_price(computer.cpu)
        if computer.ram:
            total += self.calculate_component_price(computer.ram)
        if computer.storage:
            total += self.calculate_component_price(computer.storage)
        if computer.gpu:
            total += self.calculate_component_price(computer.gpu)
        if computer.case:
            total += 100.0
        total += len(computer.accessories) * 50.0
        return total


class AssemblyScheduler:
    def __init__(self):
        self._queue = []
  
    def schedule_assembly(self, order_id: str, computer: Computer) -> str:
        position = len(self._queue) + 1
        self._queue.append(order_id)
    
        complexity = sum([
            1 if computer.cpu else 0,
            1 if computer.ram else 0,
            1 if computer.storage else 0,
            1 if computer.gpu else 0,
            1 if computer.case else 0,
            len(computer.accessories)
        ])
    
        hours = complexity * 0.5 + 1
        return f"Order #{order_id} scheduled. Position: {position}. Estimated: {hours}h"


# assembly_facade.py
from typing import Dict
from domain.models.computer import Computer
from domain.facades.subsystems import ComponentValidator, PriceCalculator, AssemblyScheduler


class ComputerAssemblyFacade:
    def __init__(self):
        self._validator = ComponentValidator()
        self._price_calculator = PriceCalculator()
        self._scheduler = AssemblyScheduler()
        self._order_counter = 1000
  
    def order_computer(self, computer: Computer) -> Dict[str, any]:
        order_id = str(self._order_counter)
        self._order_counter += 1
    
        is_valid = True
        validation_messages = []
    
        if computer.cpu and computer.ram:
            if not self._validator.validate_cpu_ram_compatibility(computer.cpu, computer.ram):
                is_valid = False
                validation_messages.append("⚠ CPU and RAM may not be compatible")
    
        if computer.gpu and computer.ram:
            if not self._validator.validate_gpu_power(computer.gpu, computer.ram):
                is_valid = False
                validation_messages.append("⚠ Insufficient RAM for this GPU")
    
        total_price = self._price_calculator.calculate_total_price(computer)
        schedule_info = self._scheduler.schedule_assembly(order_id, computer)
    
        return {
            "order_id": order_id,
            "computer": str(computer),
            "is_valid": is_valid,
            "validation_messages": validation_messages,
            "total_price": f"${total_price:.2f}",
            "schedule": schedule_info,
            "status": "✓ Order placed successfully" if is_valid else "⚠ Order placed with warnings"
        }
  
    def get_quick_quote(self, computer: Computer) -> str:
        price = self._price_calculator.calculate_total_price(computer)
        return f"${price:.2f}"
```

**Explanation:**
The Facade Pattern simplifies the complex computer ordering process by hiding three subsystems (ComponentValidator, PriceCalculator, AssemblyScheduler) behind a single, easy-to-use interface. Without the facade, clients would need to manually interact with all three subsystems, understand their APIs, and coordinate the workflow. The `ComputerAssemblyFacade` handles all this complexity internally, providing a simple `order_computer()` method that performs validation, pricing, and scheduling automatically. This reduces coupling between the client and subsystems and makes the system much easier to use.

**Usage:**

```python
facade = ComputerAssemblyFacade()

gaming_pc = (ComputerBuilder()
            .set_cpu("Intel i9")
            .set_ram("32GB DDR5")
            .set_gpu("RTX 4090")
            .build())

result = facade.order_computer(gaming_pc)
print(f"Order ID: {result['order_id']}")
print(f"Price: {result['total_price']}")
print(f"Status: {result['status']}")
# Output:
# Order ID: 1000
# Price: $2600.00
# Status: ✓ Order placed successfully
```

---

### 4. Proxy Pattern

**Purpose:** Provides a surrogate or placeholder for another object to control access to it.

**Location:** `domain/proxies/configuration_proxy.py`

**Code:**

```python
from abc import ABC, abstractmethod
from typing import Dict, List
from domain.models.computer import Computer


class ComputerConfigurationInterface(ABC):
    @abstractmethod
    def get_configuration(self, name: str) -> Computer:
        pass
  
    @abstractmethod
    def save_configuration(self, name: str, computer: Computer) -> bool:
        pass
  
    @abstractmethod
    def delete_configuration(self, name: str) -> bool:
        pass


class RealComputerConfiguration(ComputerConfigurationInterface):
    def __init__(self):
        self._configurations: Dict[str, Computer] = {}
  
    def get_configuration(self, name: str) -> Computer:
        if name in self._configurations:
            return self._configurations[name]
        raise ValueError(f"Configuration '{name}' not found")
  
    def save_configuration(self, name: str, computer: Computer) -> bool:
        self._configurations[name] = computer
        return True
  
    def delete_configuration(self, name: str) -> bool:
        if name in self._configurations:
            del self._configurations[name]
            return True
        return False


class ComputerConfigurationProxy(ComputerConfigurationInterface):
    def __init__(self, real_configuration: RealComputerConfiguration):
        self._real_configuration = real_configuration
        self._cache: Dict[str, Computer] = {}
        self._access_log: List[str] = []
        self._authenticated_users: set = set()
        self._current_user: str = None
  
    def authenticate(self, username: str, password: str) -> bool:
        if password == "admin":
            self._authenticated_users.add(username)
            self._current_user = username
            self._log(f"User '{username}' authenticated")
            return True
        return False
  
    def logout(self):
        if self._current_user:
            self._log(f"User '{self._current_user}' logged out")
            self._current_user = None
  
    def get_configuration(self, name: str) -> Computer:
        if not self._check_access():
            raise PermissionError("Authentication required")
    
        if name in self._cache:
            self._log(f"Retrieved '{name}' from cache")
            return self._cache[name]
    
        config = self._real_configuration.get_configuration(name)
        self._cache[name] = config
        self._log(f"Retrieved '{name}' from storage")
    
        return config
  
    def save_configuration(self, name: str, computer: Computer) -> bool:
        if not self._check_access():
            raise PermissionError("Authentication required")
    
        result = self._real_configuration.save_configuration(name, computer)
    
        if result:
            self._cache[name] = computer
            self._log(f"Saved configuration '{name}'")
    
        return result
  
    def delete_configuration(self, name: str) -> bool:
        if not self._check_access():
            raise PermissionError("Authentication required")
    
        result = self._real_configuration.delete_configuration(name)
    
        if result:
            if name in self._cache:
                del self._cache[name]
            self._log(f"Deleted configuration '{name}'")
    
        return result
  
    def get_access_log(self) -> List[str]:
        if not self._check_access():
            raise PermissionError("Authentication required")
        return self._access_log.copy()
  
    def _check_access(self) -> bool:
        return self._current_user in self._authenticated_users
  
    def _log(self, message: str):
        import datetime
        timestamp = datetime.datetime.now().strftime("%Y-%m-%d %H:%M:%S")
        user = self._current_user or "anonymous"
        log_entry = f"[{timestamp}] {user}: {message}"
        self._access_log.append(log_entry)
```

**Explanation:**
The Proxy Pattern controls access to the real configuration storage by adding three crucial features: authentication (access control), caching (performance optimization), and logging (audit trail). The `ComputerConfigurationProxy` acts as an intermediary between clients and the `RealComputerConfiguration`. It checks user authentication before allowing any operation, caches frequently accessed configurations to improve performance, and logs all operations for security auditing. This pattern allows adding these cross-cutting concerns without modifying the real configuration class, following the Open/Closed Principle.

**Usage:**

```python
real_config = RealComputerConfiguration()
proxy = ComputerConfigurationProxy(real_config)

# Must authenticate first
proxy.authenticate("john_doe", "admin")

# Save configuration
gaming_pc = ComputerBuilder().set_cpu("Intel i9").build()
proxy.save_configuration("gaming_setup", gaming_pc)

# First retrieval - from storage
config1 = proxy.get_configuration("gaming_setup")

# Second retrieval - from cache (faster)
config2 = proxy.get_configuration("gaming_setup")

# View access log
log = proxy.get_access_log()
# Output shows all operations with timestamps and user info
```

---

## Project Structure

```
lab3/
├── run.py                          # Entry point
├── client/
│   ├── __init__.py
│   └── main.py                     # Client demonstrations
└── domain/
    ├── __init__.py
    ├── models/
    │   ├── __init__.py
    │   └── computer.py             # Computer, ComputerBuilder
    ├── adapters/
    │   ├── __init__.py
    │   └── legacy_adapter.py       # Adapter Pattern
    ├── decorators/
    │   ├── __init__.py
    │   └── computer_decorator.py   # Decorator Pattern
    ├── facades/
    │   ├── __init__.py
    │   ├── subsystems.py           # Subsystem classes
    │   └── assembly_facade.py      # Facade Pattern
    └── proxies/
        ├── __init__.py
        └── configuration_proxy.py  # Proxy Pattern
```

---

## Output Results

```
╔====================================================================╗
║               STRUCTURAL DESIGN PATTERNS LAB                       ║
║                    Author: Drumea Vasile                           ║
║                 Computer Assembly System                           ║
╚====================================================================╝


======================================================================
ADAPTER PATTERN - Legacy Component Integration
======================================================================

📦 Integrating legacy components into modern system...

Adapted Components:
  • CPU: Intel Core i7-9700K 8-core 3.6GHz [✓ Compatible]
  • RAM: Corsair Vengeance 32GB DDR4-3200 [✓ Compatible]
  • GPU: NVIDIA GTX 1080 Ti 11GB [✓ Compatible]

✓ All legacy components are compatible with modern system!
  Building computer with legacy components...

  Result: Intel Core i7-9700K 8-core 3.6GHz, Corsair Vengeance 32GB DDR4-3200, NVIDIA GTX 1080 Ti 11GB

======================================================================
DECORATOR PATTERN - Dynamic Feature Enhancement
======================================================================

🖥️  Base Configuration:
  Description: Intel i9, 32GB DDR5, 2TB NVMe SSD, RTX 4090, Case: NZXT H710
  Cost: $1000.00

➕ Adding 3-year warranty...
  Cost: $1150.00

➕ Adding professional overclocking...
  Cost: $1350.00

➕ Adding RGB lighting system...
  Cost: $1500.00

➕ Adding custom paint job...
  Cost: $1800.00

💰 Price increase: $800.00

======================================================================
FACADE PATTERN - Simplified System Interface
======================================================================

🏪 Computer Assembly Shop - Easy Ordering System

📋 Order 1: Gaming PC
  Order ID: 1000
  Computer: Intel i9, 32GB DDR5, 2TB NVMe SSD, RTX 4090, Case: NZXT H710
  Price: $2600.00
  Order #1000 scheduled. Position: 1. Estimated: 3.5h
  ✓ Order placed successfully

📋 Order 2: Office PC (with compatibility warning)
  Order ID: 1001
  Computer: Intel i9, 16GB DDR4, 512GB SSD
  Price: $640.00
  ⚠ CPU and RAM may not be compatible
  Order #1001 scheduled. Position: 2. Estimated: 2.5h
  ⚠ Order placed with warnings

💵 Quick Quote Request:
  Server Configuration Quote: $2200.00

======================================================================
PROXY PATTERN - Secure Configuration Access
======================================================================

🔒 Configuration Management System (Secured)

❌ Attempting to access without authentication:
   Authentication required

🔐 Authenticating user...
   ✓ Authentication successful!

💾 Saving configurations:
   ✓ Saved 'gaming_setup'
   ✓ Saved 'office_setup'

📖 Retrieving 'gaming_setup' (first time):
   Intel i9, 32GB DDR5, RTX 4090

📖 Retrieving 'gaming_setup' (second time - cached):
   Intel i9, 32GB DDR5, RTX 4090

🗑️  Deleting 'office_setup':
   ✓ Configuration deleted

📊 Access Log:
   [2025-11-02 20:49:05] john_doe: Saved configuration 'gaming_setup'
   [2025-11-02 20:49:05] john_doe: Saved configuration 'office_setup'
   [2025-11-02 20:49:05] john_doe: Retrieved 'gaming_setup' from cache
   [2025-11-02 20:49:05] john_doe: Retrieved 'gaming_setup' from cache
   [2025-11-02 20:49:05] john_doe: Deleted configuration 'office_setup'

👋 Logging out...
   ✓ User logged out

❌ Attempting to access after logout:
   Authentication required

======================================================================
✓ All structural patterns demonstrated successfully!
======================================================================
```

---

## Conclusions

This project successfully demonstrates 4 structural design patterns applied to the Computer Assembly System:

1. **Adapter Pattern** - Enables seamless integration of legacy components with incompatible interfaces into the modern system. This pattern allows reusing existing code without modification, following the Open/Closed Principle. It demonstrates how to bridge different systems that need to work together but have incompatible interfaces.
2. **Decorator Pattern** - Provides dynamic feature enhancement without subclassing. Instead of creating separate classes for every feature combination (which would result in dozens of classes), decorators can be stacked in any order to add warranties, overclocking, RGB lighting, and custom paint. This pattern exemplifies composition over inheritance and keeps each feature's responsibility separate.
3. **Facade Pattern** - Dramatically simplifies the complex ordering workflow by hiding three subsystems (validation, pricing, scheduling) behind a unified interface. Clients can order computers with a single method call instead of manually coordinating multiple subsystems. This reduces coupling and makes the system much more maintainable and user-friendly.
4. **Proxy Pattern** - Adds critical cross-cutting concerns (authentication, caching, logging) without modifying the core configuration storage. The proxy controls access, improves performance through caching, and maintains a complete audit trail. This demonstrates how to add functionality transparently while maintaining the same interface.

Each pattern addresses specific structural challenges while maintaining clean architecture and following SOLID principles:

- **Single Responsibility Principle**: Each class has one clear purpose
- **Open/Closed Principle**: Open for extension, closed for modification
- **Liskov Substitution Principle**: Subtypes can replace base types seamlessly
- **Interface Segregation Principle**: Clients depend only on methods they use
- **Dependency Inversion Principle**: Depend on abstractions, not concretions

The structural patterns complement the creational patterns from previous labs, creating a comprehensive, well-architected system that is flexible, maintainable, and extensible.
