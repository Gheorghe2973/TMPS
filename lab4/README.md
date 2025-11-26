# Behavioral Design Patterns Laboratory Report

**Course:** Software Design Patterns
**Author:** George
**Topic:** Behavioral Design Patterns
**Laboratory Work:** Lab 4

---

## Objectives

1. Study and understand Behavioral Design Patterns
2. Extend previous laboratory work by implementing communication patterns between software entities
3. Implement at least 1 behavioral design pattern (implemented 4 patterns)
4. Maintain proper project structure organized by responsibilities
5. Demonstrate pattern integration and real-world applicability

---

## Domain Area

**Computer Assembly System** - An advanced system for managing computer orders with sophisticated communication patterns. The system now includes order status notifications, flexible pricing strategies, reversible operations, and validation pipelines. Behavioral patterns enable dynamic communication between system components, allowing observers to react to state changes, strategies to be swapped at runtime, commands to be undone, and validation requests to flow through chains of handlers.

---

## Theory

Behavioral design patterns are concerned with algorithms and the assignment of responsibilities between objects. They characterize complex control flow that's difficult to follow at runtime and shift focus away from control flow to let you concentrate on the way objects are interconnected.

### Key Concepts:

1. **Observer Pattern** - Defines a one-to-many dependency between objects so that when one object changes state, all its dependents are notified and updated automatically.
2. **Strategy Pattern** - Defines a family of algorithms, encapsulates each one, and makes them interchangeable. Strategy lets the algorithm vary independently from clients that use it.
3. **Command Pattern** - Encapsulates a request as an object, thereby letting you parameterize clients with different requests, queue or log requests, and support undoable operations.
4. **Chain of Responsibility Pattern** - Allows passing requests along a chain of handlers. Upon receiving a request, each handler decides either to process the request or to pass it to the next handler in the chain.

These patterns enable flexible, maintainable communication between objects while keeping them loosely coupled.

---

## Implementation

### 1. Observer Pattern

**Purpose:** Establishes a one-to-many dependency where multiple observers are automatically notified when the subject's state changes. This pattern is essential for maintaining consistency across distributed components without tight coupling.

**Location:** `patterns/observer/observer.go`

**Key Components:**

```go
// Observer interface defines the update method for observers
type Observer interface {
    Update(order *domain.Order)
    GetName() string
}

// Subject interface defines methods for managing observers
type Subject interface {
    Attach(observer Observer)
    Detach(observer Observer)
    Notify()
}

// OrderSubject is the concrete subject that holds order state
type OrderSubject struct {
    observers []Observer
    order     *domain.Order
}

func (s *OrderSubject) Attach(observer Observer) {
    s.observers = append(s.observers, observer)
}

func (s *OrderSubject) Notify() {
    for _, observer := range s.observers {
        observer.Update(s.order)
    }
}

func (s *OrderSubject) SetOrderStatus(status domain.OrderStatus) {
    s.order.Status = status
    s.Notify()
}
```

**Concrete Observers:**

```go
// CustomerObserver notifies customers about order changes
type CustomerObserver struct {
    name       string
    customerID string
}

func (o *CustomerObserver) Update(order *domain.Order) {
    if order.CustomerID == o.customerID {
        fmt.Printf("📧 Email sent: Your order %s is now %s\n", 
            order.ID, order.Status)
    }
}

// InventoryObserver updates inventory when order status changes
type InventoryObserver struct {
    name string
}

func (o *InventoryObserver) Update(order *domain.Order) {
    switch order.Status {
    case domain.OrderProcessing:
        fmt.Printf("📦 Reserved components for order %s\n", order.ID)
    case domain.OrderShipped:
        fmt.Printf("📤 Deducted components from inventory\n")
    case domain.OrderCancelled:
        fmt.Printf("↩️ Released reserved components\n")
    }
}

// AccountingObserver handles financial tracking
type AccountingObserver struct {
    name string
}

func (o *AccountingObserver) Update(order *domain.Order) {
    switch order.Status {
    case domain.OrderProcessing:
        fmt.Printf("💰 Payment of $%.2f authorized\n", order.TotalPrice)
    case domain.OrderDelivered:
        fmt.Printf("✅ Payment of $%.2f captured\n", order.TotalPrice)
    case domain.OrderCancelled:
        fmt.Printf("🔄 Refund of $%.2f initiated\n", order.TotalPrice)
    }
}
```

**Explanation:**

The Observer Pattern implements a publish-subscribe mechanism for order status changes. When an order transitions between states (PENDING → PROCESSING → SHIPPED → DELIVERED), multiple independent systems need to be notified: customers receive emails, inventory reserves/releases components, and accounting authorizes/captures payments.

**Key Design Decisions:**

- **Loose Coupling:** Observers don't know about each other; they only depend on the Subject interface
- **Dynamic Subscription:** Observers can be attached/detached at runtime using `Attach()` and `Detach()`
- **Automatic Notification:** The `Notify()` method is called automatically whenever order status changes
- **Multiple Observer Types:** Four concrete observers (Customer, Inventory, Accounting, Logging) each handle domain-specific concerns

**Benefits:**

- Eliminates tight coupling between order management and notification systems
- Allows adding new observers without modifying existing code (Open/Closed Principle)
- Each observer has a single responsibility (Single Responsibility Principle)
- Enables real-time reactive updates across the system

**Usage Example:**

```go
// Create order and subject
order := &domain.Order{ID: "ORD001", CustomerID: "CUST001", Status: domain.OrderPending}
orderSubject := observer.NewOrderSubject(order)

// Attach observers
orderSubject.Attach(observer.NewCustomerObserver("CUST001"))
orderSubject.Attach(observer.NewInventoryObserver())
orderSubject.Attach(observer.NewAccountingObserver())

// Change status - all observers are notified automatically
orderSubject.SetOrderStatus(domain.OrderProcessing)
orderSubject.SetOrderStatus(domain.OrderShipped)
orderSubject.SetOrderStatus(domain.OrderDelivered)
```

---

### 2. Strategy Pattern

**Purpose:** Encapsulates different pricing algorithms and makes them interchangeable. This pattern allows selecting pricing strategies at runtime without modifying client code, enabling flexible business logic.

**Location:** `patterns/strategy/strategy.go`

**Key Components:**

```go
// PricingStrategy defines the interface for different pricing strategies
type PricingStrategy interface {
    CalculatePrice(computer *domain.Computer) float64
    GetName() string
}

// PriceCalculator is the context that uses pricing strategies
type PriceCalculator struct {
    strategy PricingStrategy
}

func NewPriceCalculator(strategy PricingStrategy) *PriceCalculator {
    return &PriceCalculator{strategy: strategy}
}

// SetStrategy allows changing the pricing strategy at runtime
func (pc *PriceCalculator) SetStrategy(strategy PricingStrategy) {
    pc.strategy = strategy
}

// CalculateFinalPrice calculates the price using the current strategy
func (pc *PriceCalculator) CalculateFinalPrice(computer *domain.Computer) float64 {
    return pc.strategy.CalculatePrice(computer)
}
```

**Concrete Strategies:**

```go
// RetailPricingStrategy is the standard pricing for retail customers
type RetailPricingStrategy struct{}

func (s *RetailPricingStrategy) CalculatePrice(computer *domain.Computer) float64 {
    computer.CalculateTotal()
    return computer.TotalPrice // No discount
}

// WholesalePricingStrategy offers discounts for bulk/business customers
type WholesalePricingStrategy struct {
    discountPercent float64
}

func (s *WholesalePricingStrategy) CalculatePrice(computer *domain.Computer) float64 {
    computer.CalculateTotal()
    discount := computer.TotalPrice * (s.discountPercent / 100)
    return computer.TotalPrice - discount
}

// SeasonalPricingStrategy offers time-limited discounts
type SeasonalPricingStrategy struct {
    seasonName      string
    discountPercent float64
}

func (s *SeasonalPricingStrategy) CalculatePrice(computer *domain.Computer) float64 {
    computer.CalculateTotal()
    discount := computer.TotalPrice * (s.discountPercent / 100)
    return computer.TotalPrice - discount
}

// StudentPricingStrategy offers educational discounts
type StudentPricingStrategy struct {
    baseDiscount float64
}

func (s *StudentPricingStrategy) CalculatePrice(computer *domain.Computer) float64 {
    computer.CalculateTotal()
  
    // Additional 5% discount on components over $200
    additionalDiscount := 0.0
    if computer.GPU.Price > 200 {
        additionalDiscount += computer.GPU.Price * 0.05
    }
    if computer.CPU.Price > 200 {
        additionalDiscount += computer.CPU.Price * 0.05
    }
  
    baseDiscount := computer.TotalPrice * (s.baseDiscount / 100)
    return computer.TotalPrice - baseDiscount - additionalDiscount
}
```

**Explanation:**

The Strategy Pattern encapsulates different pricing algorithms into separate strategy classes. Instead of using conditional statements (`if customer == "retail" then... else if customer == "wholesale"...`), each pricing method is its own class implementing the `PricingStrategy` interface.

**Key Design Decisions:**

- **Algorithm Encapsulation:** Each pricing strategy is isolated in its own class
- **Runtime Flexibility:** The strategy can be changed at any time using `SetStrategy()`
- **Open for Extension:** New pricing strategies can be added without modifying existing code
- **Single Responsibility:** Each strategy class has one reason to change - its specific pricing logic

**Benefits:**

- Eliminates complex conditional logic in pricing calculations
- Makes it easy to add new pricing strategies (just implement the interface)
- Allows testing each pricing algorithm independently
- Enables dynamic pricing based on customer type, season, or promotions
- Follows the Open/Closed Principle perfectly

**Usage Example:**

```go
computer := createSampleComputer("PC001")

// Start with retail pricing
retailStrategy := strategy.NewRetailPricingStrategy()
calculator := strategy.NewPriceCalculator(retailStrategy)
retailPrice := calculator.CalculateFinalPrice(&computer)
// Output: $1669.95

// Customer upgrades to wholesale account
wholesaleStrategy := strategy.NewWholesalePricingStrategy(15.0)
calculator.SetStrategy(wholesaleStrategy)
wholesalePrice := calculator.CalculateFinalPrice(&computer)
// Output: $1419.46 (15% discount)

// Black Friday sale
seasonalStrategy := strategy.NewSeasonalPricingStrategy("Black Friday", 20.0)
calculator.SetStrategy(seasonalStrategy)
salePrice := calculator.CalculateFinalPrice(&computer)
// Output: $1335.96 (20% discount)
```

---

### 3. Command Pattern

**Purpose:** Encapsulates requests as objects, enabling parameterization of clients with different requests, queuing operations, logging requests, and supporting undo/redo functionality.

**Location:** `patterns/command/command.go`

**Key Components:**

```go
// Command interface defines the execute and undo operations
type Command interface {
    Execute() error
    Undo() error
    GetDescription() string
}

// OrderReceiver manages the actual order operations
type OrderReceiver struct {
    orders map[string]*domain.Order
}

func (r *OrderReceiver) CreateOrder(order *domain.Order) error {
    if _, exists := r.orders[order.ID]; exists {
        return fmt.Errorf("order %s already exists", order.ID)
    }
    r.orders[order.ID] = order
    return nil
}

func (r *OrderReceiver) UpdateOrderStatus(orderID string, status domain.OrderStatus) (domain.OrderStatus, error) {
    order, exists := r.orders[orderID]
    if !exists {
        return "", fmt.Errorf("order %s not found", orderID)
    }
    oldStatus := order.Status
    order.Status = status
    return oldStatus, nil
}
```

**Concrete Commands:**

```go
// CreateOrderCommand creates a new order
type CreateOrderCommand struct {
    receiver *OrderReceiver
    order    *domain.Order
}

func (c *CreateOrderCommand) Execute() error {
    return c.receiver.CreateOrder(c.order)
}

func (c *CreateOrderCommand) Undo() error {
    _, err := c.receiver.DeleteOrder(c.order.ID)
    return err
}

// ShipOrderCommand marks an order as shipped
type ShipOrderCommand struct {
    receiver   *OrderReceiver
    orderID    string
    prevStatus domain.OrderStatus
}

func (c *ShipOrderCommand) Execute() error {
    prevStatus, err := c.receiver.UpdateOrderStatus(c.orderID, domain.OrderShipped)
    if err != nil {
        return err
    }
    c.prevStatus = prevStatus
    return nil
}

func (c *ShipOrderCommand) Undo() error {
    _, err := c.receiver.UpdateOrderStatus(c.orderID, c.prevStatus)
    return err
}

// CancelOrderCommand cancels an existing order
type CancelOrderCommand struct {
    receiver   *OrderReceiver
    orderID    string
    prevStatus domain.OrderStatus
}

func (c *CancelOrderCommand) Execute() error {
    prevStatus, err := c.receiver.UpdateOrderStatus(c.orderID, domain.OrderCancelled)
    if err != nil {
        return err
    }
    c.prevStatus = prevStatus
    return nil
}

func (c *CancelOrderCommand) Undo() error {
    _, err := c.receiver.UpdateOrderStatus(c.orderID, c.prevStatus)
    return err
}
```

**Invoker:**

```go
// OrderInvoker executes commands and maintains history
type OrderInvoker struct {
    history []Command
}

func (i *OrderInvoker) ExecuteCommand(cmd Command) error {
    err := cmd.Execute()
    if err != nil {
        return err
    }
    i.history = append(i.history, cmd)
    return nil
}

func (i *OrderInvoker) UndoLastCommand() error {
    if len(i.history) == 0 {
        return fmt.Errorf("no commands to undo")
    }
  
    lastCmd := i.history[len(i.history)-1]
    err := lastCmd.Undo()
    if err != nil {
        return err
    }
  
    i.history = i.history[:len(i.history)-1]
    return nil
}

func (i *OrderInvoker) ShowHistory() {
    for idx, cmd := range i.history {
        fmt.Printf("%d. %s\n", idx+1, cmd.GetDescription())
    }
}
```

**Macro Command:**

```go
// MacroCommand executes multiple commands as a single unit
type MacroCommand struct {
    commands []Command
    name     string
}

func (m *MacroCommand) Execute() error {
    for _, cmd := range m.commands {
        if err := cmd.Execute(); err != nil {
            return err
        }
    }
    return nil
}

func (m *MacroCommand) Undo() error {
    // Undo in reverse order
    for i := len(m.commands) - 1; i >= 0; i-- {
        if err := m.commands[i].Undo(); err != nil {
            return err
        }
    }
    return nil
}
```

**Explanation:**

The Command Pattern transforms operations into standalone objects. Each command encapsulates all information needed to perform an action: the receiver object, the method to call, and the method arguments. This enables powerful features like undo/redo, operation logging, and batch processing.

**Key Design Decisions:**

- **Operation Encapsulation:** Each operation (create, ship, cancel) is a separate command object
- **Undo Support:** Commands store previous state to enable reversal
- **Command History:** The invoker maintains a history of executed commands
- **Macro Commands:** Multiple commands can be grouped and executed as one unit
- **Separation of Concerns:** Command objects separate the requester from the performer

**Benefits:**

- Enables undo/redo functionality naturally
- Allows queuing, scheduling, and logging operations
- Supports macro commands (composite pattern integration)
- Decouples the object that invokes the operation from the one that performs it
- Makes it easy to add new commands without changing existing code

**Usage Example:**

```go
receiver := command.NewOrderReceiver()
invoker := command.NewOrderInvoker()

// Create and execute commands
order1 := &domain.Order{ID: "ORD001", Status: domain.OrderPending}
createCmd := command.NewCreateOrderCommand(receiver, order1)
invoker.ExecuteCommand(createCmd)

shipCmd := command.NewShipOrderCommand(receiver, "ORD001")
invoker.ExecuteCommand(shipCmd)

// View command history
invoker.ShowHistory()
// Output:
// 1. Create Order: ORD001
// 2. Ship Order: ORD001

// Undo last command
invoker.UndoLastCommand()
// Order status reverted back to PENDING

// Macro command - batch processing
macro := command.NewMacroCommand("Process Orders")
macro.AddCommand(createCmd1)
macro.AddCommand(createCmd2)
macro.AddCommand(shipCmd1)
invoker.ExecuteCommand(macro)
```

---

### 4. Chain of Responsibility Pattern

**Purpose:** Allows passing requests along a chain of handlers. Each handler decides whether to process the request or pass it to the next handler. This pattern decouples senders and receivers of requests.

**Location:** `patterns/chain/chain.go`

**Key Components:**

```go
// ValidationHandler interface for the chain of responsibility
type ValidationHandler interface {
    SetNext(handler ValidationHandler) ValidationHandler
    Handle(order *domain.Order) error
}

// BaseHandler provides default chaining behavior
type BaseHandler struct {
    next ValidationHandler
}

func (h *BaseHandler) SetNext(handler ValidationHandler) ValidationHandler {
    h.next = handler
    return handler
}

func (h *BaseHandler) Handle(order *domain.Order) error {
    if h.next != nil {
        return h.next.Handle(order)
    }
    return nil
}
```

**Concrete Handlers:**

```go
// CompatibilityChecker validates component compatibility
type CompatibilityChecker struct {
    BaseHandler
}

func (h *CompatibilityChecker) Handle(order *domain.Order) error {
    // Check CPU-Motherboard socket compatibility
    if order.Computer.CPU.Type != order.Computer.Motherboard.Type {
        return fmt.Errorf("CPU socket %s incompatible with motherboard socket %s",
            order.Computer.CPU.Type, order.Computer.Motherboard.Type)
    }
  
    // Check RAM compatibility
    if order.Computer.RAM.Type != "DDR4" && order.Computer.RAM.Type != "DDR5" {
        return fmt.Errorf("unsupported RAM type: %s", order.Computer.RAM.Type)
    }
  
    // Check storage interface
    if order.Computer.Storage.Type != "NVMe" && order.Computer.Storage.Type != "SATA" {
        return fmt.Errorf("unsupported storage interface: %s", order.Computer.Storage.Type)
    }
  
    // Pass to next handler
    if h.next != nil {
        return h.next.Handle(order)
    }
    return nil
}

// BudgetValidator checks if order meets minimum/maximum budget constraints
type BudgetValidator struct {
    BaseHandler
    minBudget float64
    maxBudget float64
}

func (h *BudgetValidator) Handle(order *domain.Order) error {
    totalPrice := order.TotalPrice
  
    if totalPrice < h.minBudget {
        return fmt.Errorf("order total $%.2f is below minimum budget $%.2f",
            totalPrice, h.minBudget)
    }
  
    if totalPrice > h.maxBudget {
        return fmt.Errorf("order total $%.2f exceeds maximum budget $%.2f",
            totalPrice, h.maxBudget)
    }
  
    // Pass to next handler
    if h.next != nil {
        return h.next.Handle(order)
    }
    return nil
}

// StockValidator checks if all components are in stock
type StockValidator struct {
    BaseHandler
}

func (h *StockValidator) Handle(order *domain.Order) error {
    components := []domain.Component{
        order.Computer.CPU,
        order.Computer.GPU,
        order.Computer.RAM,
        order.Computer.Storage,
        order.Computer.Motherboard,
    }
  
    allAvailable := true
    for _, component := range components {
        if component.Stock <= 0 {
            fmt.Printf("❌ %s is out of stock\n", component.Name)
            allAvailable = false
        }
    }
  
    if !allAvailable {
        return fmt.Errorf("some components are out of stock")
    }
  
    // Pass to next handler
    if h.next != nil {
        return h.next.Handle(order)
    }
    return nil
}

// CustomerCreditChecker validates customer credit/payment method
type CustomerCreditChecker struct {
    BaseHandler
    approvedCustomers map[string]bool
}

func (h *CustomerCreditChecker) Handle(order *domain.Order) error {
    if !h.approvedCustomers[order.CustomerID] {
        return fmt.Errorf("customer %s is not approved for credit", order.CustomerID)
    }
  
    // Pass to next handler
    if h.next != nil {
        return h.next.Handle(order)
    }
    return nil
}

// FinalApprovalHandler is the last handler that approves the order
type FinalApprovalHandler struct {
    BaseHandler
}

func (h *FinalApprovalHandler) Handle(order *domain.Order) error {
    fmt.Printf("🎉 Order %s is APPROVED for processing\n", order.ID)
    return nil
}
```

**Chain Construction:**

```go
// OrderValidator orchestrates the validation chain
type OrderValidator struct {
    chain ValidationHandler
}

func NewOrderValidator() *OrderValidator {
    // Build the chain
    compatibility := NewCompatibilityChecker()
    budget := NewBudgetValidator(500.0, 10000.0)
    stock := NewStockValidator()
    credit := NewCustomerCreditChecker()
    approval := NewFinalApprovalHandler()
  
    // Chain them together
    compatibility.SetNext(budget).
        SetNext(stock).
        SetNext(credit).
        SetNext(approval)
  
    return &OrderValidator{
        chain: compatibility,
    }
}

// Validate runs the entire validation chain
func (v *OrderValidator) Validate(order *domain.Order) error {
    return v.chain.Handle(order)
}
```

**Explanation:**

The Chain of Responsibility Pattern creates a pipeline of validation handlers. An order must pass through multiple checks: component compatibility, budget constraints, stock availability, and customer credit. Each handler either processes the request and passes it along, or stops the chain by returning an error.

**Key Design Decisions:**

- **Sequential Processing:** Validators are executed in a specific order
- **Early Termination:** Chain stops at the first failure
- **Extensibility:** New validators can be added to the chain without modifying existing ones
- **Loose Coupling:** Each handler only knows about the next handler in the chain
- **Single Responsibility:** Each handler performs one specific validation

**Benefits:**

- Reduces coupling between sender and receivers
- Adds/removes responsibilities dynamically
- Follows the Single Responsibility Principle (each validator has one job)
- Follows the Open/Closed Principle (add new validators without changing existing code)
- Provides clear, readable validation flow
- Makes it easy to reorder or skip validations

**Usage Example:**

```go
validator := chain.NewOrderValidator()

// Test Case 1: Valid order
order1 := &domain.Order{
    ID: "ORD001",
    Computer: computer1,
    TotalPrice: 1669.95,
    CustomerID: "CUST001",
}

err := validator.Validate(order1)
// Output: ✅ VALIDATION SUCCESSFUL

// Test Case 2: Incompatible components
order2 := &domain.Order{
    ID: "ORD002",
    Computer: computer2, // CPU socket doesn't match motherboard
    TotalPrice: 1669.95,
    CustomerID: "CUST001",
}

err = validator.Validate(order2)
// Output: ❌ VALIDATION FAILED: CPU socket LGA1700 incompatible with motherboard socket LGA1200

// Test Case 3: Out of stock
order3 := &domain.Order{
    ID: "ORD003",
    Computer: computer3, // GPU out of stock
    TotalPrice: 1669.95,
    CustomerID: "CUST001",
}

err = validator.Validate(order3)
// Output: ❌ VALIDATION FAILED: some components are out of stock
```

---

## Project Structure

```
lab4/
├── client/
│   ├── main.go                     # Main client application
│   └── demonstrations.go           # Pattern demonstrations
├── domain/
│   └── models.go                   # Domain models (Computer, Order, Component)
├── patterns/
│   ├── observer/
│   │   └── observer.go             # Observer Pattern implementation
│   ├── strategy/
│   │   └── strategy.go             # Strategy Pattern implementation
│   ├── command/
│   │   └── command.go              # Command Pattern implementation
│   └── chain/
│       └── chain.go                # Chain of Responsibility implementation
├── go.mod                          # Go module definition
└── README.md                       # Project documentation
```

**Organization Principles:**

- **Domain Layer:** Core business models (Computer, Order, OrderStatus)
- **Patterns Layer:** Each behavioral pattern in its own package
- **Client Layer:** Demonstration code showing pattern usage
- **Separation of Concerns:** Each package has a single, well-defined responsibility

---

## Output Results

### Observer Pattern Demonstration

```
======================================================================
OBSERVER PATTERN DEMONSTRATION
======================================================================

Scenario: Order status changes trigger notifications to multiple observers
Pattern: One-to-many dependency where observers are notified automatically

--- Setting up Observers ---
[OrderSubject] Attached observer: CustomerObserver-CUST001
[OrderSubject] Attached observer: InventoryObserver
[OrderSubject] Attached observer: AccountingObserver
[OrderSubject] Attached observer: LoggingObserver

--- Order Lifecycle Simulation ---
[OrderSubject] Order ORD001 status changed: PENDING -> PROCESSING

[OrderSubject] Notifying all observers...
  [CustomerObserver-CUST001] 📧 Email sent: Your order ORD001 is now PROCESSING
  [InventoryObserver] 📦 Reserved components for order ORD001
  [AccountingObserver] 💰 Payment of $1669.95 authorized for order ORD001
  [LoggingObserver] 📝 LOG: Order ORD001 | Status=PROCESSING | Customer=CUST001 | Price=$1669.95

[OrderSubject] Order ORD001 status changed: PROCESSING -> SHIPPED

[OrderSubject] Notifying all observers...
  [CustomerObserver-CUST001] 📧 Email sent: Your order ORD001 is now SHIPPED
  [InventoryObserver] 📤 Deducted components from inventory for order ORD001
  [LoggingObserver] 📝 LOG: Order ORD001 | Status=SHIPPED | Customer=CUST001 | Price=$1669.95

[OrderSubject] Order ORD001 status changed: SHIPPED -> DELIVERED

[OrderSubject] Notifying all observers...
  [CustomerObserver-CUST001] 📧 Email sent: Your order ORD001 is now DELIVERED
  [AccountingObserver] ✅ Payment of $1669.95 captured for order ORD001
  [LoggingObserver] 📝 LOG: Order ORD001 | Status=DELIVERED | Customer=CUST001 | Price=$1669.95

--- Detaching an Observer ---
[OrderSubject] Detached observer: LoggingObserver
[OrderSubject] Order ORD002 status changed: PENDING -> CANCELLED

[OrderSubject] Notifying all observers...
  [CustomerObserver-CUST001] 📧 Email sent: Your order ORD002 is now CANCELLED
  [InventoryObserver] ↩️  Released reserved components for order ORD002
```

### Strategy Pattern Demonstration

```
======================================================================
STRATEGY PATTERN DEMONSTRATION
======================================================================

Scenario: Calculate prices using different pricing strategies
Pattern: Encapsulate algorithms and make them interchangeable

Base Computer Configuration:
  CPU: Intel Core i7-13700K ($409.99)
  GPU: NVIDIA RTX 4070 ($599.99)
  RAM: Corsair Vengeance 32GB ($129.99)
  Storage: Samsung 980 Pro 1TB ($149.99)
  Motherboard: ASUS ROG Strix Z790 ($379.99)

============================================================
PRICING COMPARISON
============================================================

[PriceCalculator] Calculating price using: Retail Pricing (No discount)
  Original Price: $1669.95
  Final Price: $1669.95

[PriceCalculator] Calculating price using: Wholesale Pricing (15% discount)
  Original Price: $1669.95
  Discount: -$250.49 (15.0%)
  Final Price: $1419.46

[PriceCalculator] Calculating price using: Black Friday Sale (20% discount)
  Original Price: $1669.95
  Discount: -$333.99 (20.0%)
  Final Price: $1335.96

[PriceCalculator] Calculating price using: Student Pricing (10% + extra discounts)
  Original Price: $1669.95
  Discount: -$217.49 (13.0%)
  Final Price: $1452.46

--- Runtime Strategy Switching ---

Customer Type: Retail
  Final Price: $1669.95

Customer upgrades to Business (Wholesale):
[PriceCalculator] Strategy changed to: Wholesale Pricing (15% discount)
  Final Price: $1419.46

Savings from upgrade: $250.49
```

### Command Pattern Demonstration

```
======================================================================
COMMAND PATTERN DEMONSTRATION
======================================================================

Scenario: Execute order operations with undo capability
Pattern: Encapsulate requests as objects to support undo/redo

--- Executing Commands ---

[Invoker] Executing: Create Order: ORD003
[CreateOrderCommand] Executing...
  [Receiver] Created order: ORD003
[Invoker] ✅ Command executed successfully

[Invoker] Executing: Create Order: ORD004
[CreateOrderCommand] Executing...
  [Receiver] Created order: ORD004
[Invoker] ✅ Command executed successfully

[Invoker] Executing: Ship Order: ORD003
[ShipOrderCommand] Executing...
  [Receiver] Updated order ORD003 status: PENDING -> SHIPPED
[Invoker] ✅ Command executed successfully

============================================================
COMMAND HISTORY
============================================================
  1. Create Order: ORD003
  2. Create Order: ORD004
  3. Ship Order: ORD003
============================================================

--- Undoing Last Command ---

[Invoker] Undoing: Ship Order: ORD003
[ShipOrderCommand] Undoing...
  [Receiver] Updated order ORD003 status: SHIPPED -> PENDING
[Invoker] ✅ Command undone successfully

============================================================
COMMAND HISTORY
============================================================
  1. Create Order: ORD003
  2. Create Order: ORD004
============================================================

--- Macro Command (Batch Processing) ---

[Invoker] Executing: Macro: Process Multiple Orders (2 commands)
[MacroCommand] Executing macro: Process Multiple Orders
[CreateOrderCommand] Executing...
  [Receiver] Created order: ORD005
[ShipOrderCommand] Executing...
  [Receiver] Updated order ORD005 status: PENDING -> SHIPPED
[Invoker] ✅ Command executed successfully
```

### Chain of Responsibility Pattern Demonstration

```
======================================================================
CHAIN OF RESPONSIBILITY PATTERN DEMONSTRATION
======================================================================

Scenario: Validate orders through a chain of validators
Pattern: Pass requests along a chain until one handles it

--- Test Case 1: Valid Order ---

============================================================
VALIDATING ORDER: ORD006
============================================================

[CompatibilityChecker] Validating component compatibility...
  ✅ CPU and Motherboard are compatible
  ✅ RAM type is supported
  ✅ Storage interface is supported
[CompatibilityChecker] ✅ All components are compatible

[BudgetValidator] Validating order budget...
  ✅ Order meets minimum budget ($1669.95 >= $500.00)
  ✅ Order within maximum budget ($1669.95 <= $10000.00)
[BudgetValidator] ✅ Budget validation passed

[StockValidator] Validating component stock availability...
  ✅ Intel Core i7-13700K (LGA1700) - 15 units available
  ✅ NVIDIA RTX 4070 (PCIe 4.0) - 8 units available
  ✅ Corsair Vengeance 32GB (DDR4) - 25 units available
  ✅ Samsung 980 Pro 1TB (NVMe) - 30 units available
  ✅ ASUS ROG Strix Z790 (LGA1700) - 10 units available
[StockValidator] ✅ All components are in stock

[CustomerCreditChecker] Validating customer credit...
  ✅ Customer CUST001 has approved credit
[CustomerCreditChecker] ✅ Credit validation passed

[FinalApprovalHandler] All validations passed!
  🎉 Order ORD006 is APPROVED for processing

============================================================
✅ VALIDATION SUCCESSFUL
============================================================

--- Test Case 2: Incompatible Components ---

============================================================
VALIDATING ORDER: ORD007
============================================================

[CompatibilityChecker] Validating component compatibility...

============================================================
❌ VALIDATION FAILED: CPU socket LGA1700 incompatible with motherboard socket LGA1200
============================================================

--- Test Case 3: Component Out of Stock ---

============================================================
VALIDATING ORDER: ORD008
============================================================

[CompatibilityChecker] Validating component compatibility...
  ✅ CPU and Motherboard are compatible
  ✅ RAM type is supported
  ✅ Storage interface is supported
[CompatibilityChecker] ✅ All components are compatible

[BudgetValidator] Validating order budget...
  ✅ Order meets minimum budget ($1669.95 >= $500.00)
  ✅ Order within maximum budget ($1669.95 <= $10000.00)
[BudgetValidator] ✅ Budget validation passed

[StockValidator] Validating component stock availability...
  ✅ Intel Core i7-13700K (LGA1700) - 15 units available
  ❌ NVIDIA RTX 4070 (PCIe 4.0) is out of stock
  ✅ Corsair Vengeance 32GB (DDR4) - 25 units available
  ✅ Samsung 980 Pro 1TB (NVMe) - 30 units available
  ✅ ASUS ROG Strix Z790 (LGA1700) - 10 units available

============================================================
❌ VALIDATION FAILED: some components are out of stock
============================================================
```

### Integrated Scenario - All Patterns Working Together

```
======================================================================
INTEGRATED SCENARIO - ALL PATTERNS WORKING TOGETHER
======================================================================

Scenario: Complete order processing pipeline
Combining: Chain → Strategy → Command → Observer

--- STEP 1: Apply Pricing Strategy ---

[PriceCalculator] Calculating price using: Student Pricing (12% + extra discounts)
  Original Price: $1669.95
  Discount: -$250.89 (15.0%)
  Final Price: $1419.06

--- STEP 2: Validate Order (Chain of Responsibility) ---

============================================================
VALIDATING ORDER: ORD010
============================================================

[CompatibilityChecker] Validating component compatibility...
  ✅ All components are compatible

[BudgetValidator] Validating order budget...
  ✅ Order meets budget constraints

[StockValidator] Validating component stock availability...
  ✅ All components are in stock

[CustomerCreditChecker] Validating customer credit...
  ✅ Customer CUST003 has approved credit

[FinalApprovalHandler] All validations passed!
  🎉 Order ORD010 is APPROVED for processing

============================================================
✅ VALIDATION SUCCESSFUL
============================================================

--- STEP 3: Setup Observers ---
[OrderSubject] Attached observer: CustomerObserver-CUST003
[OrderSubject] Attached observer: InventoryObserver
[OrderSubject] Attached observer: AccountingObserver

--- STEP 4: Process Order (Command Pattern) ---

[Invoker] Executing: Create Order: ORD010
[CreateOrderCommand] Executing...
  [Receiver] Created order: ORD010
[Invoker] ✅ Command executed successfully

[OrderSubject] Order ORD010 status changed: PENDING -> PROCESSING

[OrderSubject] Notifying all observers...
  [CustomerObserver-CUST003] 📧 Email sent: Your order ORD010 is now PROCESSING
  [InventoryObserver] 📦 Reserved components for order ORD010
  [AccountingObserver] 💰 Payment of $1419.06 authorized for order ORD010

[Invoker] Executing: Ship Order: ORD010
[ShipOrderCommand] Executing...
  [Receiver] Updated order ORD010 status: PROCESSING -> SHIPPED
[Invoker] ✅ Command executed successfully

[OrderSubject] Order ORD010 status changed: SHIPPED -> SHIPPED

[OrderSubject] Notifying all observers...
  [CustomerObserver-CUST003] 📧 Email sent: Your order ORD010 is now SHIPPED
  [InventoryObserver] 📤 Deducted components from inventory for order ORD010

[OrderSubject] Order ORD010 status changed: SHIPPED -> DELIVERED

[OrderSubject] Notifying all observers...
  [CustomerObserver-CUST003] 📧 Email sent: Your order ORD010 is now DELIVERED
  [AccountingObserver] ✅ Payment of $1419.06 captured for order ORD010

--- Final Command History ---

============================================================
COMMAND HISTORY
============================================================
  1. Create Order: ORD010
  2. Ship Order: ORD010
============================================================

✅ Order successfully processed through all stages!
```

---

## Conclusions

This laboratory work successfully demonstrates the implementation and integration of four behavioral design patterns in Go, applied to a Computer Assembly System. Each pattern addresses specific communication and behavioral challenges:

### Pattern Analysis:

1. **Observer Pattern** - Successfully implements a publish-subscribe mechanism for order status changes. The pattern eliminates tight coupling between order management and notification systems, allowing multiple independent observers (Customer, Inventory, Accounting, Logging) to react to state changes automatically. This demonstrates the power of event-driven architecture and reactive programming.
2. **Strategy Pattern** - Provides flexible pricing algorithms that can be selected and changed at runtime. Instead of hardcoding pricing logic with complex conditionals, each strategy is encapsulated in its own class. This makes the system highly extensible - new pricing strategies can be added without modifying existing code, perfectly following the Open/Closed Principle.
3. **Command Pattern** - Transforms operations into first-class objects, enabling powerful features like undo/redo, operation logging, and batch processing (macro commands). This pattern provides a clean separation between invokers and receivers, making the system more maintainable and testable. The ability to reverse operations is particularly valuable in order management systems.
4. **Chain of Responsibility Pattern** - Creates a flexible validation pipeline where each handler performs a specific check and passes the request along. The chain can be easily extended or reordered without affecting existing validators. This pattern demonstrates excellent use of the Single Responsibility Principle, with each validator focused on one specific concern.

### Integration Benefits:

The integrated scenario demonstrates how all four patterns work together seamlessly:

- **Chain of Responsibility** validates the order through multiple checks
- **Strategy** calculates the final price with the appropriate discount
- **Command** executes reversible operations with full history tracking
- **Observer** notifies all relevant systems about order status changes

This integration shows that behavioral patterns complement each other, creating a robust, maintainable, and extensible architecture.

### SOLID Principles Adherence:

- **Single Responsibility Principle**: Each class has one clear purpose (e.g., CustomerObserver only handles customer notifications)
- **Open/Closed Principle**: New observers, strategies, commands, and validators can be added without modifying existing code
- **Liskov Substitution Principle**: All implementations can be substituted for their interfaces without breaking the system
- **Interface Segregation Principle**: Interfaces are focused and minimal (Observer, Command, PricingStrategy, ValidationHandler)
- **Dependency Inversion Principle**: Components depend on abstractions (interfaces) rather than concrete implementations

### Technical Achievements:

1. **Loose Coupling**: Components communicate through well-defined interfaces, making the system flexible and testable
2. **High Cohesion**: Related functionality is grouped together logically
3. **Extensibility**: New features can be added with minimal changes to existing code
4. **Maintainability**: Clear separation of concerns makes the codebase easy to understand and modify
5. **Testability**: Each pattern component can be tested independently

### Real-World Applicability:

These patterns solve common real-world problems:

- **Observer**: Event notification systems, real-time updates, monitoring
- **Strategy**: Payment processing, sorting algorithms, compression algorithms
- **Command**: Transaction systems, job queues, macro recording
- **Chain of Responsibility**: Request filtering, logging, validation pipelines

The implementation in Go demonstrates that behavioral patterns are language-agnostic and provide value regardless of the programming paradigm.

### Lessons Learned:

1. Behavioral patterns excel at managing complex interactions and responsibilities
2. Pattern combination creates more powerful solutions than individual patterns
3. Proper abstraction (interfaces) is crucial for pattern implementation
4. Go's interface system provides excellent support for design patterns
5. Behavioral patterns significantly improve code organization and maintainability

This laboratory work provides a solid foundation for understanding how behavioral design patterns enable flexible, maintainable communication between objects in complex software systems.
