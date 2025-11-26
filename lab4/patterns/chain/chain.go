package chain

import (
	"fmt"
	"lab4/domain"
	"strings"
)

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

// CompatibilityChecker validates component compatibility
type CompatibilityChecker struct {
	BaseHandler
}

func NewCompatibilityChecker() *CompatibilityChecker {
	return &CompatibilityChecker{}
}

func (h *CompatibilityChecker) Handle(order *domain.Order) error {
	fmt.Println("\n[CompatibilityChecker] Validating component compatibility...")

	// Check CPU-Motherboard socket compatibility (simplified)
	if order.Computer.CPU.Type != order.Computer.Motherboard.Type {
		return fmt.Errorf("CPU socket %s incompatible with motherboard socket %s",
			order.Computer.CPU.Type, order.Computer.Motherboard.Type)
	}
	fmt.Println("  ✅ CPU and Motherboard are compatible")

	// Check RAM compatibility
	if order.Computer.RAM.Type != "DDR4" && order.Computer.RAM.Type != "DDR5" {
		return fmt.Errorf("unsupported RAM type: %s", order.Computer.RAM.Type)
	}
	fmt.Println("  ✅ RAM type is supported")

	// Check storage interface
	if order.Computer.Storage.Type != "NVMe" && order.Computer.Storage.Type != "SATA" {
		return fmt.Errorf("unsupported storage interface: %s", order.Computer.Storage.Type)
	}
	fmt.Println("  ✅ Storage interface is supported")

	fmt.Println("[CompatibilityChecker] ✅ All components are compatible")

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

func NewBudgetValidator(minBudget, maxBudget float64) *BudgetValidator {
	return &BudgetValidator{
		minBudget: minBudget,
		maxBudget: maxBudget,
	}
}

func (h *BudgetValidator) Handle(order *domain.Order) error {
	fmt.Println("\n[BudgetValidator] Validating order budget...")

	totalPrice := order.TotalPrice

	if totalPrice < h.minBudget {
		return fmt.Errorf("order total $%.2f is below minimum budget $%.2f",
			totalPrice, h.minBudget)
	}
	fmt.Printf("  ✅ Order meets minimum budget ($%.2f >= $%.2f)\n", totalPrice, h.minBudget)

	if totalPrice > h.maxBudget {
		return fmt.Errorf("order total $%.2f exceeds maximum budget $%.2f",
			totalPrice, h.maxBudget)
	}
	fmt.Printf("  ✅ Order within maximum budget ($%.2f <= $%.2f)\n", totalPrice, h.maxBudget)

	fmt.Println("[BudgetValidator] ✅ Budget validation passed")

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

func NewStockValidator() *StockValidator {
	return &StockValidator{}
}

func (h *StockValidator) Handle(order *domain.Order) error {
	fmt.Println("\n[StockValidator] Validating component stock availability...")

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
			fmt.Printf("  ❌ %s (%s) is out of stock\n", component.Name, component.Type)
			allAvailable = false
		} else {
			fmt.Printf("  ✅ %s (%s) - %d units available\n",
				component.Name, component.Type, component.Stock)
		}
	}

	if !allAvailable {
		return fmt.Errorf("some components are out of stock")
	}

	fmt.Println("[StockValidator] ✅ All components are in stock")

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

func NewCustomerCreditChecker() *CustomerCreditChecker {
	return &CustomerCreditChecker{
		approvedCustomers: map[string]bool{
			"CUST001": true,
			"CUST002": true,
			"CUST003": true,
		},
	}
}

func (h *CustomerCreditChecker) Handle(order *domain.Order) error {
	fmt.Println("\n[CustomerCreditChecker] Validating customer credit...")

	if !h.approvedCustomers[order.CustomerID] {
		return fmt.Errorf("customer %s is not approved for credit", order.CustomerID)
	}

	fmt.Printf("  ✅ Customer %s has approved credit\n", order.CustomerID)
	fmt.Println("[CustomerCreditChecker] ✅ Credit validation passed")

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

func NewFinalApprovalHandler() *FinalApprovalHandler {
	return &FinalApprovalHandler{}
}

func (h *FinalApprovalHandler) Handle(order *domain.Order) error {
	fmt.Println("\n[FinalApprovalHandler] All validations passed!")
	fmt.Printf("  🎉 Order %s is APPROVED for processing\n", order.ID)
	return nil
}

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
	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Printf("VALIDATING ORDER: %s\n", order.ID)
	fmt.Println(strings.Repeat("=", 60))

	err := v.chain.Handle(order)

	fmt.Println("\n" + strings.Repeat("=", 60))
	if err != nil {
		fmt.Printf("❌ VALIDATION FAILED: %v\n", err)
	} else {
		fmt.Println("✅ VALIDATION SUCCESSFUL")
	}
	fmt.Println(strings.Repeat("=", 60) + "\n")

	return err
}
