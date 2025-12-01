# Computer Assembly - Creational Design Patterns Report

**Course:** Software Design Patterns
**Topic:** Creational Design Patterns

---

## Objectives

1. Study and understand Creational Design Patterns
2. Choose a domain and define main classes/models/entities
3. Implement multiple creational design patterns in a sample project

---

## Domain Area

**Computer Assembly System** - A system for creating and configuring computer builds with various specifications and purposes (Gaming, Office, Server, Custom builds).

---

## Implementation

### 1. Factory Method Pattern

**Purpose:** Defines an interface for creating objects, but lets subclasses decide which class to instantiate.

**Location:** `domain/factory/computer_factory.py`

**Code:**

```python
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
```

**Explanation:**
The `ComputerFactory` abstract class defines the interface with the `create_computer()` method. Three concrete factories (`GamingPCFactory`, `OfficePCFactory`, `ServerFactory`) implement this method to create specific computer types. Each factory encapsulates the creation logic for a particular computer configuration. This allows the client code to create computers without knowing the specific configuration details, and new computer types can be added by creating new factory classes without modifying existing code.

**Usage:**

```python
gaming_factory = GamingPCFactory()
gaming_pc = gaming_factory.create_computer()
print(f"Gaming PC: {gaming_pc}")
# Output: Gaming PC: Intel i9, 32GB DDR5, 2TB NVMe SSD, RTX 4090
```

---

### 2. Builder Pattern

**Purpose:** Separates the construction of a complex object from its representation, allowing the same construction process to create different representations.

**Location:** `domain/models/computer.py`

**Code:**

```python
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
```

**Explanation:**
The `ComputerBuilder` class provides a fluent interface for constructing `Computer` objects step-by-step. Each setter method returns `self`, enabling method chaining. The `Computer` class represents the complex object being built with multiple optional components (CPU, RAM, GPU, storage, case, accessories). This pattern is particularly useful when an object has many optional parameters, avoiding constructor telescoping and making the code more readable.

**Usage:**

```python
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
# Output: Custom Computer with all specified components
```

---

### 3. Singleton Pattern

**Purpose:** Ensures a class has only one instance and provides a global point of access to it.

**Location:** `domain/models/computer.py`

**Code:**

```python
class ConfigurationManager:
    _instance = None

    def __new__(cls):
        if cls._instance is None:
            cls._instance = super().__new__(cls)
            cls._instance.configurations = {}
        return cls._instance

    def add_configuration(self, name, computer):
        self.configurations[name] = computer

    def get_configuration(self, name):
        return self.configurations.get(name)

    def list_configurations(self):
        return list(self.configurations.keys())
```

**Explanation:**
The `ConfigurationManager` uses Python's `__new__` method to control instance creation. The `_instance` class variable stores the single instance. When `__new__` is called, it checks if an instance already exists; if not, it creates one and initializes the configurations dictionary. All subsequent calls return the same instance. This ensures that all parts of the application share the same configuration data, preventing inconsistencies and reducing memory usage.

**Usage:**

```python
config1 = ConfigurationManager()
config2 = ConfigurationManager()
print(f"Same instance: {config1 is config2}")  # Output: True

gaming_pc = GamingPCFactory().create_computer()
config1.add_configuration("Gaming Setup", gaming_pc)
print(config2.list_configurations())  # Output: ['Gaming Setup']
```

---

### 4. Prototype Pattern

**Purpose:** Creates new objects by copying an existing object (prototype) rather than creating from scratch.

**Location:** `domain/models/computer.py`

**Code:**

```python
class Computer:
    # ... (previous code)
  
    def clone(self):
        new_computer = Computer()
        new_computer.cpu = self.cpu
        new_computer.ram = self.ram
        new_computer.storage = self.storage
        new_computer.gpu = self.gpu
        new_computer.case = self.case
        new_computer.accessories = self.accessories.copy()
        return new_computer
```

**Explanation:**
The `clone()` method creates a new `Computer` instance and copies all attributes from the current instance to the new one. The accessories list is copied using `.copy()` to create a shallow copy, preventing shared reference issues. This pattern is useful when creating objects is expensive or complex, or when you want to create variations of a base configuration. Instead of rebuilding an entire computer from scratch, you clone a template and modify only what's needed.

**Usage:**

```python
base_pc = (ComputerBuilder()
           .set_cpu("Intel i7")
           .set_ram("16GB DDR4")
           .set_storage("1TB SSD")
           .build())

cloned_pc1 = base_pc.clone()
cloned_pc1.gpu = "RTX 4070"

cloned_pc2 = base_pc.clone()
cloned_pc2.gpu = "RTX 4060"

# Output:
# Base PC: Intel i7, 16GB DDR4, 1TB SSD
# Clone 1: Intel i7, 16GB DDR4, 1TB SSD, RTX 4070
# Clone 2: Intel i7, 16GB DDR4, 1TB SSD, RTX 4060
```

---

### 5. Abstract Factory Pattern

**Purpose:** Provides an interface for creating families of related objects without specifying their concrete classes.

**Location:** `domain/factory/component_factory.py`

**Code:**

```python
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


class AMDFactory(ComponentFactory):
    def create_cpu(self):
        return "AMD Ryzen 9"

    def create_ram(self):
        return "32GB DDR5"

    def create_storage(self):
        return "2TB NVMe SSD"


class IntelAMDFactory(ComponentFactory):
    def create_cpu(self):
        return "Intel i7"

    def create_ram(self):
        return "16GB DDR4"

    def create_storage(self):
        return "1TB SSD"


class BudgetFactory(ComponentFactory):
    def create_cpu(self):
        return "Intel i3"

    def create_ram(self):
        return "8GB DDR4"

    def create_storage(self):
        return "256GB SSD"
```

**Explanation:**
The `ComponentFactory` abstract class defines methods for creating a family of related components (CPU, RAM, Storage). Each concrete factory (`AMDFactory`, `IntelAMDFactory`, `BudgetFactory`) implements these methods to create components that work well together. This ensures compatibility between components - an AMD factory creates AMD-compatible parts, a Budget factory creates budget-tier parts. The client can switch between entire component families by changing the factory instance, without modifying the construction code.

**Usage:**

```python
amd_factory = AMDFactory()

amd_computer = (ComputerBuilder()
                .set_cpu(amd_factory.create_cpu())
                .set_ram(amd_factory.create_ram())
                .set_storage(amd_factory.create_storage())
                .build())

print(f"AMD Computer: {amd_computer}")
# Output: AMD Computer: AMD Ryzen 9, 32GB DDR5, 2TB NVMe SSD
```

---

## Project Structure

```
lab2/
├── domain/
│   ├── __init__.py
│   ├── models/
│   │   ├── __init__.py
│   │   └── computer.py          # Computer, ComputerBuilder, ConfigurationManager
│   └── factory/
│       ├── __init__.py
│       ├── computer_factory.py   # Factory Method implementations
│       └── component_factory.py  # Abstract Factory implementations
├── client/
│   ├── __init__.py
│   └── main.py                   # Demonstrates all patterns
├── .gitignore
└── README.md
```

---

## Output Results

```
=== Factory Method Pattern ===
Gaming PC: Intel i9, 32GB DDR5, 2TB NVMe SSD, RTX 4090
Office PC: Intel i5, 16GB DDR4, 512GB SSD
Server: AMD EPYC, 128GB ECC, 4TB RAID

=== Builder Pattern ===
Custom Computer:
  CPU: AMD Ryzen 9
  RAM: 64GB DDR5
  Storage: 1TB NVMe + 2TB HDD
  GPU: RTX 4080
  Case: Corsair 5000D
  Accessories: RGB Keyboard, Gaming Mouse, Monitor 27"

=== Singleton Pattern ===
Same instance: True
Configurations: ['Gaming Setup']

=== Prototype Pattern ===
Base PC: Intel i7, 16GB DDR4, 1TB SSD
Clone 1: Intel i7, 16GB DDR4, 1TB SSD, RTX 4070
Clone 2: Intel i7, 16GB DDR4, 1TB SSD, RTX 4060

=== Abstract Factory Pattern ===
AMD Computer: AMD Ryzen 9, 32GB DDR5, 2TB NVMe SSD
```

---

## Conclusions

This project successfully demonstrates 5 creational design patterns applied to a computer assembly domain:

1. **Factory Method** - Encapsulates object creation logic for different computer types (Gaming, Office, Server), allowing easy extension without modifying existing code.
2. **Builder** - Provides a fluent interface for constructing complex Computer objects with many optional components, improving code readability and maintainability.
3. **Singleton** - Ensures single instance of ConfigurationManager for consistent state management across the application.
4. **Prototype** - Enables efficient cloning of base computer configurations, reducing initialization overhead when creating similar objects.
5. **Abstract Factory** - Creates families of compatible components (AMD, Intel, Budget), ensuring component compatibility and allowing easy switching between product families.

Each pattern addresses specific object creation challenges while maintaining clean architecture and following SOLID principles, particularly the Open/Closed Principle (open for extension, closed for modification).
