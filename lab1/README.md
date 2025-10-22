# Creational Design Patterns

## Author: Gurschi Gheorghe

## Objectives:

* Get familiar with the Creational Design Patterns;
* Choose a specific domain;
* Implement at least 3 Creational Design Patterns for the specific domain;

## Used Design Patterns:

* **Singleton Pattern** - Database Connection
* **Factory Method Pattern** - Computer Factory
* **Builder Pattern** - Custom Computer Builder

## SOLID Principles Implementation

Pe lângă Design Patterns, acest proiect respectă și implementează 3 principii SOLID fundamentale:

### 1. Single Responsibility Principle (SRP)

Fiecare clasă din proiect are o singură responsabilitate bine definită:

* `DatabaseConnection` - gestionează exclusiv conexiunea la baza de date
* `Computer` - definește interfața comună pentru toate tipurile de computere
* `GamingComputer`, `OfficeComputer`, `ServerComputer` - fiecare gestionează propriile specificații
* `ComputerFactory` - responsabil doar pentru crearea instanțelor de calculatoare
* `CustomComputer` - stochează datele unui computer personalizat
* `ComputerBuilder` - construiește pas cu pas computere personalizate

**Exemplu din cod:**

```python
class DatabaseConnection:  # O singură responsabilitate: gestionare conexiune
    def query(self, sql: str):
        return f"Executing: {sql}"

class ComputerFactory:  # O singură responsabilitate: creare computere
    @staticmethod
    def create_computer(computer_type: str) -> Computer:
        # Logică de creare...
```

### 2. Open/Closed Principle (OCP)

Sistemul este deschis pentru extindere, dar închis pentru modificare. Putem adăuga noi tipuri de calculatoare fără a modifica clasele existente:

**Exemplu de extindere:**

```python
# Clasa de bază rămâne neschimbată (closed for modification)
class Computer(ABC):
    @abstractmethod
    def get_specs(self) -> str:
        pass

# Putem extinde cu noi tipuri (open for extension)
class GamingComputer(Computer):  # Tip existent
    pass

class WorkstationComputer(Computer):  # Tip nou - fără modificări în cod existent!
    def __init__(self):
        super().__init__()
        self.cpu = "Intel Xeon"
        self.ram = "64GB ECC"
        self.storage = "2TB NVMe"
  
    def get_specs(self) -> str:
        return f"Workstation: {self.cpu}, {self.ram}, {self.storage}"
```

Acest principiu este evidențiat în special prin Factory Method Pattern, care permite adăugarea de noi produse fără modificarea structurii existente.

### 3. Liskov Substitution Principle (LSP)

Obiectele claselor derivate pot înlocui obiectele clasei de bază fără a afecta corectitudinea programului. Toate subclasele `Computer` pot fi folosite interschimbabil:

**Exemplu din cod:**

```python
# Factory returnează tipul de bază Computer
def create_computer(computer_type: str) -> Computer:
    if computer_type == "gaming":
        return GamingComputer()  # Subclasă care înlocuiește baza
    elif computer_type == "office":
        return OfficeComputer()  # Subclasă care înlocuiește baza
    elif computer_type == "server":
        return ServerComputer()  # Subclasă care înlocuiește baza

# Toate pot fi tratate uniform ca Computer
gaming_pc: Computer = ComputerFactory.create_computer("gaming")
office_pc: Computer = ComputerFactory.create_computer("office")
server: Computer = ComputerFactory.create_computer("server")

# Toate implementează get_specs() și funcționează identic
print(gaming_pc.get_specs())  # ✅ Funcționează
print(office_pc.get_specs())  # ✅ Funcționează
print(server.get_specs())     # ✅ Funcționează
```

Fiecare subclasă respectă contractul definit de clasa abstractă `Computer`, asigurând substituibilitatea perfectă.

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

Când creăm două instanțe ale clasei DatabaseConnection, ambele referă același obiect în memorie, demonstrând astfel pattern-ul Singleton. **Acest pattern respectă SRP** - clasa are responsabilitatea unică de a gestiona conexiunea la baza de date.

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

Factory-ul centralizează logica de creare a obiectelor și face codul mai ușor de extins cu noi tipuri de calculatoare. **Acest pattern demonstrează perfect OCP** - putem adăuga noi tipuri de computere fără a modifica interfața `Computer` sau alte clase existente.

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

Builder-ul oferă o interfață fluent (method chaining) pentru construirea obiectelor complexe într-un mod intuitiv și ușor de citit. **Acest pattern respectă SRP** - responsabilitatea unică este construirea pas cu pas a obiectelor complexe.

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

În acest laborator am implementat cu succes trei Creational Design Patterns fundamentale în contextul unui Computer Shop, respectând totodată trei principii SOLID esențiale.

**Singleton Pattern** s-a dovedit esențial pentru gestionarea resurselor partajate, în cazul nostru conexiunea la baza de date. Acest pattern garantează că o singură instanță există în sistem, prevenind crearea multiplă de conexiuni și asigurând consistența datelor. Implementarea respectă **Single Responsibility Principle**, având o responsabilitate unică și bine definită.

**Factory Method Pattern** a simplificat semnificativ procesul de creare a obiectelor. În loc să instanțiem direct clasele specifice (GamingComputer, OfficeComputer, ServerComputer), factory-ul oferă o interfață uniformă care ascunde complexitatea și face codul mai ușor de menținut și extins. Acest pattern demonstrează excelent **Open/Closed Principle** - sistemul este deschis pentru extindere (putem adăuga noi tipuri de computere) dar închis pentru modificare (nu trebuie să modificăm clasele existente).

**Builder Pattern** a demonstrat utilitatea sa în construirea obiectelor complexe cu multe proprietăți opționale. Sintaxa fluent (method chaining) face codul mult mai lizibil comparativ cu constructori cu mulți parametri, iar procesul de construcție devine explicit și ușor de urmărit. Pattern-ul respectă **Single Responsibility Principle** prin separarea clară a responsabilităților.

**Liskov Substitution Principle** este implementat consistent în ierarhia de clase `Computer`, unde toate subclasele (GamingComputer, OfficeComputer, ServerComputer) pot înlocui clasa de bază fără a afecta corectitudinea programului. Acest lucru este evidențiat în Factory Method, unde toate obiectele create pot fi tratate uniform prin interfața comună.

Aceste pattern-uri creaționale, împreună cu principiile SOLID, îmbunătățesc calitatea codului prin:

* **Încapsularea** logicii de creare
* **Reducerea coupling-ului** între clase
* **Creșterea flexibilității** și mentenabilității
* **Oferirea unei structuri** clare și consistente pentru instanțierea obiectelor
* **Respectarea principiilor** de design orientat pe obiect

Implementarea lor în Python demonstrează cum principiile OOP și SOLID pot fi aplicate elegant pentru a rezolva probleme comune de design software, creând cod robust, extensibil și ușor de întreținut.

---

### How to Run:

```bash
python main.py
```

### Requirements:

* Python 3.7+
* No external dependencies required
