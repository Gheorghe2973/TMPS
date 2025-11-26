package main

import (
	"fmt"
	"lab4/domain"
	"lab4/patterns/chain"
	"lab4/patterns/command"
	"lab4/patterns/observer"
	"lab4/patterns/strategy"
	"strings"
	"time"
)

func printSeparator(title string) {
	fmt.Println("\n" + strings.Repeat("=", 70))
	fmt.Printf("  %s\n", title)
	fmt.Println(strings.Repeat("=", 70))
}

func createSampleComputer(id string) domain.Computer {
	return domain.Computer{
		ID: id,
		CPU: domain.Component{
			ID:    "CPU001",
			Name:  "Intel Core i7-13700K",
			Type:  "LGA1700",
			Price: 409.99,
			Stock: 15,
		},
		GPU: domain.Component{
			ID:    "GPU001",
			Name:  "NVIDIA RTX 4070",
			Type:  "PCIe 4.0",
			Price: 599.99,
			Stock: 8,
		},
		RAM: domain.Component{
			ID:    "RAM001",
			Name:  "Corsair Vengeance 32GB",
			Type:  "DDR4",
			Price: 129.99,
			Stock: 25,
		},
		Storage: domain.Component{
			ID:    "SSD001",
			Name:  "Samsung 980 Pro 1TB",
			Type:  "NVMe",
			Price: 149.99,
			Stock: 30,
		},
		Motherboard: domain.Component{
			ID:    "MB001",
			Name:  "ASUS ROG Strix Z790",
			Type:  "LGA1700",
			Price: 379.99,
			Stock: 10,
		},
	}
}

func demonstrateObserverPattern() {
	printSeparator("OBSERVER PATTERN DEMONSTRATION")

	fmt.Println("\nScenario: Order status changes trigger notifications to multiple observers")
	fmt.Println("Pattern: One-to-many dependency where observers are notified automatically")

	// Create a computer and order
	computer := createSampleComputer("PC001")
	computer.CalculateTotal()

	order := &domain.Order{
		ID:         "ORD001",
		CustomerID: "CUST001",
		Computer:   computer,
		Status:     domain.OrderPending,
		TotalPrice: computer.TotalPrice,
		CreatedAt:  time.Now().Format(time.RFC3339),
	}

	// Create the subject
	orderSubject := observer.NewOrderSubject(order)

	// Create and attach observers
	fmt.Println("\n--- Setting up Observers ---")
	customerObs := observer.NewCustomerObserver("CUST001")
	inventoryObs := observer.NewInventoryObserver()
	accountingObs := observer.NewAccountingObserver()
	loggingObs := observer.NewLoggingObserver()

	orderSubject.Attach(customerObs)
	orderSubject.Attach(inventoryObs)
	orderSubject.Attach(accountingObs)
	orderSubject.Attach(loggingObs)

	// Simulate order lifecycle
	fmt.Println("\n--- Order Lifecycle Simulation ---")
	time.Sleep(500 * time.Millisecond)
	orderSubject.SetOrderStatus(domain.OrderProcessing)

	time.Sleep(500 * time.Millisecond)
	orderSubject.SetOrderStatus(domain.OrderShipped)

	time.Sleep(500 * time.Millisecond)
	orderSubject.SetOrderStatus(domain.OrderDelivered)

	// Demonstrate detaching an observer
	fmt.Println("\n--- Detaching an Observer ---")
	orderSubject.Detach(loggingObs)

	// Try another status change
	time.Sleep(500 * time.Millisecond)
	order2 := &domain.Order{
		ID:         "ORD002",
		CustomerID: "CUST001",
		Computer:   computer,
		Status:     domain.OrderPending,
		TotalPrice: computer.TotalPrice,
		CreatedAt:  time.Now().Format(time.RFC3339),
	}
	orderSubject2 := observer.NewOrderSubject(order2)
	orderSubject2.Attach(customerObs)
	orderSubject2.Attach(inventoryObs)
	orderSubject2.SetOrderStatus(domain.OrderCancelled)
}

func demonstrateStrategyPattern() {
	printSeparator("STRATEGY PATTERN DEMONSTRATION")

	fmt.Println("\nScenario: Calculate prices using different pricing strategies")
	fmt.Println("Pattern: Encapsulate algorithms and make them interchangeable")

	computer := createSampleComputer("PC002")
	computer.CalculateTotal()

	fmt.Printf("\nBase Computer Configuration:\n")
	fmt.Printf("  CPU: %s ($%.2f)\n", computer.CPU.Name, computer.CPU.Price)
	fmt.Printf("  GPU: %s ($%.2f)\n", computer.GPU.Name, computer.GPU.Price)
	fmt.Printf("  RAM: %s ($%.2f)\n", computer.RAM.Name, computer.RAM.Price)
	fmt.Printf("  Storage: %s ($%.2f)\n", computer.Storage.Name, computer.Storage.Price)
	fmt.Printf("  Motherboard: %s ($%.2f)\n", computer.Motherboard.Name, computer.Motherboard.Price)

	// Create different pricing strategies
	retailStrategy := strategy.NewRetailPricingStrategy()
	wholesaleStrategy := strategy.NewWholesalePricingStrategy(15.0)
	seasonalStrategy := strategy.NewSeasonalPricingStrategy("Black Friday", 20.0)
	studentStrategy := strategy.NewStudentPricingStrategy(10.0)

	// Compare all strategies
	strategies := []strategy.PricingStrategy{
		retailStrategy,
		wholesaleStrategy,
		seasonalStrategy,
		studentStrategy,
	}

	strategy.CompareStrategies(&computer, strategies)

	// Demonstrate runtime strategy switching
	fmt.Println("\n--- Runtime Strategy Switching ---")
	calculator := strategy.NewPriceCalculator(retailStrategy)

	fmt.Println("\nCustomer Type: Retail")
	retailPrice := calculator.CalculateFinalPrice(&computer)

	fmt.Println("\nCustomer upgrades to Business (Wholesale):")
	calculator.SetStrategy(wholesaleStrategy)
	wholesalePrice := calculator.CalculateFinalPrice(&computer)

	fmt.Printf("Savings from upgrade: $%.2f\n", retailPrice-wholesalePrice)
}

func demonstrateCommandPattern() {
	printSeparator("COMMAND PATTERN DEMONSTRATION")

	fmt.Println("\nScenario: Execute order operations with undo capability")
	fmt.Println("Pattern: Encapsulate requests as objects to support undo/redo")

	// Create receiver and invoker
	receiver := command.NewOrderReceiver()
	invoker := command.NewOrderInvoker()

	// Create orders
	computer1 := createSampleComputer("PC003")
	computer1.CalculateTotal()

	order1 := &domain.Order{
		ID:         "ORD003",
		CustomerID: "CUST002",
		Computer:   computer1,
		Status:     domain.OrderPending,
		TotalPrice: computer1.TotalPrice,
		CreatedAt:  time.Now().Format(time.RFC3339),
	}

	order2 := &domain.Order{
		ID:         "ORD004",
		CustomerID: "CUST002",
		Computer:   computer1,
		Status:     domain.OrderPending,
		TotalPrice: computer1.TotalPrice,
		CreatedAt:  time.Now().Format(time.RFC3339),
	}

	// Execute commands
	fmt.Println("\n--- Executing Commands ---")

	createCmd1 := command.NewCreateOrderCommand(receiver, order1)
	invoker.ExecuteCommand(createCmd1)

	time.Sleep(300 * time.Millisecond)

	createCmd2 := command.NewCreateOrderCommand(receiver, order2)
	invoker.ExecuteCommand(createCmd2)

	time.Sleep(300 * time.Millisecond)

	shipCmd := command.NewShipOrderCommand(receiver, "ORD003")
	invoker.ExecuteCommand(shipCmd)

	// Show history
	invoker.ShowHistory()

	// Undo last command
	fmt.Println("\n--- Undoing Last Command ---")
	invoker.UndoLastCommand()

	time.Sleep(300 * time.Millisecond)

	// Show updated history
	invoker.ShowHistory()

	// Demonstrate macro command
	fmt.Println("\n--- Macro Command (Batch Processing) ---")
	macro := command.NewMacroCommand("Process Multiple Orders")

	order3 := &domain.Order{
		ID:         "ORD005",
		CustomerID: "CUST003",
		Computer:   computer1,
		Status:     domain.OrderPending,
		TotalPrice: computer1.TotalPrice,
		CreatedAt:  time.Now().Format(time.RFC3339),
	}

	macro.AddCommand(command.NewCreateOrderCommand(receiver, order3))
	macro.AddCommand(command.NewShipOrderCommand(receiver, "ORD005"))

	invoker.ExecuteCommand(macro)

	invoker.ShowHistory()
}

func demonstrateChainOfResponsibilityPattern() {
	printSeparator("CHAIN OF RESPONSIBILITY PATTERN DEMONSTRATION")

	fmt.Println("\nScenario: Validate orders through a chain of validators")
	fmt.Println("Pattern: Pass requests along a chain until one handles it")

	// Create validator
	validator := chain.NewOrderValidator()

	// Test Case 1: Valid order
	fmt.Println("\n--- Test Case 1: Valid Order ---")
	computer1 := createSampleComputer("PC004")
	computer1.CalculateTotal()

	order1 := &domain.Order{
		ID:         "ORD006",
		CustomerID: "CUST001",
		Computer:   computer1,
		Status:     domain.OrderPending,
		TotalPrice: computer1.TotalPrice,
		CreatedAt:  time.Now().Format(time.RFC3339),
	}

	err := validator.Validate(order1)
	if err == nil {
		fmt.Println("✅ Order can proceed to fulfillment")
	}

	time.Sleep(500 * time.Millisecond)

	// Test Case 2: Incompatible components
	fmt.Println("\n--- Test Case 2: Incompatible Components ---")
	computer2 := createSampleComputer("PC005")
	computer2.Motherboard.Type = "LGA1200" // Different socket
	computer2.CalculateTotal()

	order2 := &domain.Order{
		ID:         "ORD007",
		CustomerID: "CUST001",
		Computer:   computer2,
		Status:     domain.OrderPending,
		TotalPrice: computer2.TotalPrice,
		CreatedAt:  time.Now().Format(time.RFC3339),
	}

	err = validator.Validate(order2)
	if err != nil {
		fmt.Printf("⚠️  Order rejected: %v\n", err)
	}

	time.Sleep(500 * time.Millisecond)

	// Test Case 3: Out of stock
	fmt.Println("\n--- Test Case 3: Component Out of Stock ---")
	computer3 := createSampleComputer("PC006")
	computer3.GPU.Stock = 0 // Out of stock
	computer3.CalculateTotal()

	order3 := &domain.Order{
		ID:         "ORD008",
		CustomerID: "CUST001",
		Computer:   computer3,
		Status:     domain.OrderPending,
		TotalPrice: computer3.TotalPrice,
		CreatedAt:  time.Now().Format(time.RFC3339),
	}

	err = validator.Validate(order3)
	if err != nil {
		fmt.Printf("⚠️  Order rejected: %v\n", err)
	}

	time.Sleep(500 * time.Millisecond)

	// Test Case 4: Budget exceeded
	fmt.Println("\n--- Test Case 4: Budget Exceeded ---")
	computer4 := createSampleComputer("PC007")
	computer4.GPU.Price = 8000.0 // Very expensive GPU
	computer4.CalculateTotal()

	order4 := &domain.Order{
		ID:         "ORD009",
		CustomerID: "CUST001",
		Computer:   computer4,
		Status:     domain.OrderPending,
		TotalPrice: computer4.TotalPrice,
		CreatedAt:  time.Now().Format(time.RFC3339),
	}

	err = validator.Validate(order4)
	if err != nil {
		fmt.Printf("⚠️  Order rejected: %v\n", err)
	}
}

func demonstrateIntegratedScenario() {
	printSeparator("INTEGRATED SCENARIO - ALL PATTERNS WORKING TOGETHER")

	fmt.Println("\nScenario: Complete order processing pipeline")
	fmt.Println("Combining: Chain → Strategy → Command → Observer")

	// Step 1: Create computer and calculate base price
	computer := createSampleComputer("PC008")
	computer.CalculateTotal()

	fmt.Println("\n--- STEP 1: Apply Pricing Strategy ---")
	studentStrategy := strategy.NewStudentPricingStrategy(12.0)
	calculator := strategy.NewPriceCalculator(studentStrategy)
	finalPrice := calculator.CalculateFinalPrice(&computer)

	// Step 2: Create order
	order := &domain.Order{
		ID:         "ORD010",
		CustomerID: "CUST003",
		Computer:   computer,
		Status:     domain.OrderPending,
		TotalPrice: finalPrice,
		CreatedAt:  time.Now().Format(time.RFC3339),
	}

	// Step 3: Validate through chain
	fmt.Println("\n--- STEP 2: Validate Order (Chain of Responsibility) ---")
	validator := chain.NewOrderValidator()
	err := validator.Validate(order)

	if err != nil {
		fmt.Printf("\n❌ Order validation failed: %v\n", err)
		return
	}

	// Step 4: Setup observers
	fmt.Println("\n--- STEP 3: Setup Observers ---")
	orderSubject := observer.NewOrderSubject(order)
	orderSubject.Attach(observer.NewCustomerObserver("CUST003"))
	orderSubject.Attach(observer.NewInventoryObserver())
	orderSubject.Attach(observer.NewAccountingObserver())

	// Step 5: Execute commands with notifications
	fmt.Println("\n--- STEP 4: Process Order (Command Pattern) ---")
	receiver := command.NewOrderReceiver()
	invoker := command.NewOrderInvoker()

	createCmd := command.NewCreateOrderCommand(receiver, order)
	invoker.ExecuteCommand(createCmd)

	time.Sleep(500 * time.Millisecond)
	orderSubject.SetOrderStatus(domain.OrderProcessing)

	time.Sleep(500 * time.Millisecond)
	shipCmd := command.NewShipOrderCommand(receiver, order.ID)
	invoker.ExecuteCommand(shipCmd)
	orderSubject.SetOrderStatus(domain.OrderShipped)

	time.Sleep(500 * time.Millisecond)
	orderSubject.SetOrderStatus(domain.OrderDelivered)

	fmt.Println("\n--- Final Command History ---")
	invoker.ShowHistory()

	fmt.Println("\n✅ Order successfully processed through all stages!")
}

func main() {
	fmt.Println("\n╔══════════════════════════════════════════════════════════════════╗")
	fmt.Println("║  BEHAVIORAL DESIGN PATTERNS - Computer Assembly System          ║")
	fmt.Println("╚══════════════════════════════════════════════════════════════════╝")

	// Demonstrate each pattern individually
	demonstrateObserverPattern()
	time.Sleep(1 * time.Second)

	demonstrateStrategyPattern()
	time.Sleep(1 * time.Second)

	demonstrateCommandPattern()
	time.Sleep(1 * time.Second)

	demonstrateChainOfResponsibilityPattern()
	time.Sleep(1 * time.Second)

	// Demonstrate all patterns working together
	demonstrateIntegratedScenario()

	printSeparator("DEMONSTRATION COMPLETE")
	fmt.Println("\nAll behavioral patterns demonstrated successfully!")
	fmt.Println("Each pattern solves a specific communication/interaction problem:")
	fmt.Println("  • Observer: Automatic notifications on state changes")
	fmt.Println("  • Strategy: Flexible algorithm selection at runtime")
	fmt.Println("  • Command: Encapsulated operations with undo capability")
	fmt.Println("  • Chain of Responsibility: Sequential validation pipeline")
	fmt.Println()
}
