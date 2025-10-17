# Creational Design Patterns

## Author: [Numele Tău]

----

## Objectives:

* Get familiar with the Creational Design Patterns;
* Choose a specific domain;
* Implement at least 3 Creational Design Patterns for the specific domain;

## Used Design Patterns:

* **Singleton Pattern** - Database Connection
* **Factory Method Pattern** - Computer Factory
* **Builder Pattern** - Custom Computer Builder

## Implementation

### Domain Description
Am ales domeniul unui **Computer Shop** care gestionează diferite tipuri de calculatoare și conexiuni la baze de date. Sistemul permite crearea de calculatoare predefinite (gaming, office, server) și construirea de calculatoare custom personalizate.

### 1. Singleton Pattern - Database Connection

Singleton Pattern asigură că există o singură instanță a conexiunii la baza de date în întreaga aplicație, economisind resurse și evitând probleme de sincronizare.

```python
class DatabaseConnection:
    _instance = None
    
    def __new__(cls):
        if cls._instance is None:
            cls._instance = super().__new__(cls)
            cls._instance.connection = "Connected to Database"
            print("Creating new database connection...")
        return cls._instance
```

Când creăm două instanțe ale clasei DatabaseConnection, ambele referă același obiect în memorie, demonstrând astfel pattern-ul Singleton.

### 2. Factory Method Pattern - Computer Factory

Factory Method Pattern permite crearea diferitelor tipuri de calculatoare fără a specifica clasa exactă a obiectului care va fi creat. Acest pattern oferă o interfață pentru crearea obiectelor, dar permite subclaselor să decidă ce clasă să instanțieze.

```python
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
```

Factory-ul centralizează logica de creare a obiectelor și face codul mai ușor de extins cu noi tipuri de calculatoare.

### 3. Builder Pattern - Custom Computer Builder

Builder Pattern permite construirea pas cu pas a unui obiect complex. Acest pattern separă construcția unui obiect de reprezentarea sa, permițând același proces de construcție să creeze diferite reprezentări.

```python
class ComputerBuilder:
    def __init__(self):
        self.computer = CustomComputer()
    
    def set_cpu(self, cpu: str):
        self.computer.cpu = cpu
        return self
    
    def set_ram(self, ram: str):
        self.computer.ram = ram
        return self
    
    def build(self) -> CustomComputer:
        return self.computer
```

Builder-ul oferă o interfață fluent (method chaining) pentru construirea obiectelor complexe într-un mod intuitiv și ușor de citit.

## Conclusions / Results

### Output Example:
```
=== CREATIONAL DESIGN PATTERNS DEMO ===

1. SINGLETON PATTERN - Database Connection
Creating new database connection...
Same instance? True
Executing: SELECT * FROM computers

2. FACTORY METHOD PATTERN - Computer Factory
Gaming PC: Intel i9, 32GB DDR5, 2TB NVMe SSD, RTX 4090
Office PC: Intel i5, 16GB DDR4, 512GB SSD
Server: AMD EPYC, 128GB ECC, 4TB RAID

3. BUILDER PATTERN - Custom Computer Builder
Custom Computer:
  CPU: AMD Ryzen 9
  RAM: 64GB DDR5
  Storage: 1TB NVMe + 2TB HDD
  GPU: RTX 4080
  Case: Corsair 5000D
  Accessories: RGB Keyboard, Gaming Mouse, Monitor 27"
```

### Conclusions:

În acest laborator am implementat cu succes trei Creational Design Patterns fundamentale în contextul unui Computer Shop:

**Singleton Pattern** s-a dovedit esențial pentru gestionarea resurselor partajate, în cazul nostru conexiunea la baza de date. Acest pattern garantează că o singură instanță există în sistem, prevenind crearea multiplă de conexiuni și asigurând consistența datelor.

**Factory Method Pattern** a simplificat semnificativ procesul de creare a obiectelor. În loc să instanțiem direct clasele specifice (GamingComputer, OfficeComputer, ServerComputer), factory-ul oferă o interfață uniformă care ascunde complexitatea și face codul mai ușor de menținut și extins.

**Builder Pattern** a demonstrat utilitatea sa în construirea obiectelor complexe cu multe proprietăți opționale. Sintaxa fluent (method chaining) face codul mult mai lizibil comparativ cu constructori cu mulți parametri, iar procesul de construcție devine explicit și ușor de urmărit.

Aceste pattern-uri creaționale îmbunătățesc calitatea codului prin: încapsularea logicii de creare, reducerea coupling-ului între clase, creșterea flexibilității și mentenabilității, și oferirea unei structuri clare și consistente pentru instanțierea obiectelor. Implementarea lor în Python demonstrează cum principiile OOP pot fi aplicate elegant pentru a rezolva probleme comune de design software.

---

### How to Run:
```bash
python main.py
```

### Requirements:
* Python 3.7+
* No external dependencies required