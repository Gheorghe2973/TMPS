package observer

import (
	"fmt"
	"lab4/domain"
)

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

func NewOrderSubject(order *domain.Order) *OrderSubject {
	return &OrderSubject{
		observers: make([]Observer, 0),
		order:     order,
	}
}

// Attach adds an observer to the list
func (s *OrderSubject) Attach(observer Observer) {
	s.observers = append(s.observers, observer)
	fmt.Printf("[OrderSubject] Attached observer: %s\n", observer.GetName())
}

// Detach removes an observer from the list
func (s *OrderSubject) Detach(observer Observer) {
	for i, obs := range s.observers {
		if obs.GetName() == observer.GetName() {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
			fmt.Printf("[OrderSubject] Detached observer: %s\n", observer.GetName())
			return
		}
	}
}

// Notify alerts all observers about the change
func (s *OrderSubject) Notify() {
	fmt.Println("\n[OrderSubject] Notifying all observers...")
	for _, observer := range s.observers {
		observer.Update(s.order)
	}
	fmt.Println()
}

// SetOrderStatus changes the order status and notifies observers
func (s *OrderSubject) SetOrderStatus(status domain.OrderStatus) {
	fmt.Printf("[OrderSubject] Order %s status changed: %s -> %s\n",
		s.order.ID, s.order.Status, status)
	s.order.Status = status
	s.Notify()
}

func (s *OrderSubject) GetOrder() *domain.Order {
	return s.order
}

// CustomerObserver notifies customers about order changes
type CustomerObserver struct {
	name       string
	customerID string
}

func NewCustomerObserver(customerID string) *CustomerObserver {
	return &CustomerObserver{
		name:       fmt.Sprintf("CustomerObserver-%s", customerID),
		customerID: customerID,
	}
}

func (o *CustomerObserver) Update(order *domain.Order) {
	if order.CustomerID == o.customerID {
		fmt.Printf("  [%s] 📧 Email sent: Your order %s is now %s\n",
			o.name, order.ID, order.Status)
	}
}

func (o *CustomerObserver) GetName() string {
	return o.name
}

// InventoryObserver updates inventory when order status changes
type InventoryObserver struct {
	name string
}

func NewInventoryObserver() *InventoryObserver {
	return &InventoryObserver{
		name: "InventoryObserver",
	}
}

func (o *InventoryObserver) Update(order *domain.Order) {
	switch order.Status {
	case domain.OrderProcessing:
		fmt.Printf("  [%s] 📦 Reserved components for order %s\n", o.name, order.ID)
	case domain.OrderShipped:
		fmt.Printf("  [%s] 📤 Deducted components from inventory for order %s\n", o.name, order.ID)
	case domain.OrderCancelled:
		fmt.Printf("  [%s] ↩️  Released reserved components for order %s\n", o.name, order.ID)
	}
}

func (o *InventoryObserver) GetName() string {
	return o.name
}

// AccountingObserver handles financial tracking
type AccountingObserver struct {
	name string
}

func NewAccountingObserver() *AccountingObserver {
	return &AccountingObserver{
		name: "AccountingObserver",
	}
}

func (o *AccountingObserver) Update(order *domain.Order) {
	switch order.Status {
	case domain.OrderProcessing:
		fmt.Printf("  [%s] 💰 Payment of $%.2f authorized for order %s\n",
			o.name, order.TotalPrice, order.ID)
	case domain.OrderDelivered:
		fmt.Printf("  [%s] ✅ Payment of $%.2f captured for order %s\n",
			o.name, order.TotalPrice, order.ID)
	case domain.OrderCancelled:
		fmt.Printf("  [%s] 🔄 Refund of $%.2f initiated for order %s\n",
			o.name, order.TotalPrice, order.ID)
	}
}

func (o *AccountingObserver) GetName() string {
	return o.name
}

// LoggingObserver logs all order changes
type LoggingObserver struct {
	name string
}

func NewLoggingObserver() *LoggingObserver {
	return &LoggingObserver{
		name: "LoggingObserver",
	}
}

func (o *LoggingObserver) Update(order *domain.Order) {
	fmt.Printf("  [%s] 📝 LOG: Order %s | Status=%s | Customer=%s | Price=$%.2f\n",
		o.name, order.ID, order.Status, order.CustomerID, order.TotalPrice)
}

func (o *LoggingObserver) GetName() string {
	return o.name
}
