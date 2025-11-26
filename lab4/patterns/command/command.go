package command

import (
	"fmt"
	"lab4/domain"
	"strings"
	"time"
)

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

func NewOrderReceiver() *OrderReceiver {
	return &OrderReceiver{
		orders: make(map[string]*domain.Order),
	}
}

func (r *OrderReceiver) CreateOrder(order *domain.Order) error {
	if _, exists := r.orders[order.ID]; exists {
		return fmt.Errorf("order %s already exists", order.ID)
	}
	r.orders[order.ID] = order
	fmt.Printf("  [Receiver] Created order: %s\n", order.ID)
	return nil
}

func (r *OrderReceiver) DeleteOrder(orderID string) (*domain.Order, error) {
	order, exists := r.orders[orderID]
	if !exists {
		return nil, fmt.Errorf("order %s not found", orderID)
	}
	delete(r.orders, orderID)
	fmt.Printf("  [Receiver] Deleted order: %s\n", orderID)
	return order, nil
}

func (r *OrderReceiver) UpdateOrderStatus(orderID string, status domain.OrderStatus) (domain.OrderStatus, error) {
	order, exists := r.orders[orderID]
	if !exists {
		return "", fmt.Errorf("order %s not found", orderID)
	}
	oldStatus := order.Status
	order.Status = status
	fmt.Printf("  [Receiver] Updated order %s status: %s -> %s\n", orderID, oldStatus, status)
	return oldStatus, nil
}

func (r *OrderReceiver) GetOrder(orderID string) (*domain.Order, error) {
	order, exists := r.orders[orderID]
	if !exists {
		return nil, fmt.Errorf("order %s not found", orderID)
	}
	return order, nil
}

// CreateOrderCommand creates a new order
type CreateOrderCommand struct {
	receiver *OrderReceiver
	order    *domain.Order
}

func NewCreateOrderCommand(receiver *OrderReceiver, order *domain.Order) *CreateOrderCommand {
	return &CreateOrderCommand{
		receiver: receiver,
		order:    order,
	}
}

func (c *CreateOrderCommand) Execute() error {
	fmt.Println("[CreateOrderCommand] Executing...")
	return c.receiver.CreateOrder(c.order)
}

func (c *CreateOrderCommand) Undo() error {
	fmt.Println("[CreateOrderCommand] Undoing...")
	_, err := c.receiver.DeleteOrder(c.order.ID)
	return err
}

func (c *CreateOrderCommand) GetDescription() string {
	return fmt.Sprintf("Create Order: %s", c.order.ID)
}

// CancelOrderCommand cancels an existing order
type CancelOrderCommand struct {
	receiver   *OrderReceiver
	orderID    string
	prevStatus domain.OrderStatus
}

func NewCancelOrderCommand(receiver *OrderReceiver, orderID string) *CancelOrderCommand {
	return &CancelOrderCommand{
		receiver: receiver,
		orderID:  orderID,
	}
}

func (c *CancelOrderCommand) Execute() error {
	fmt.Println("[CancelOrderCommand] Executing...")
	prevStatus, err := c.receiver.UpdateOrderStatus(c.orderID, domain.OrderCancelled)
	if err != nil {
		return err
	}
	c.prevStatus = prevStatus
	return nil
}

func (c *CancelOrderCommand) Undo() error {
	fmt.Println("[CancelOrderCommand] Undoing...")
	_, err := c.receiver.UpdateOrderStatus(c.orderID, c.prevStatus)
	return err
}

func (c *CancelOrderCommand) GetDescription() string {
	return fmt.Sprintf("Cancel Order: %s", c.orderID)
}

// ShipOrderCommand marks an order as shipped
type ShipOrderCommand struct {
	receiver   *OrderReceiver
	orderID    string
	prevStatus domain.OrderStatus
}

func NewShipOrderCommand(receiver *OrderReceiver, orderID string) *ShipOrderCommand {
	return &ShipOrderCommand{
		receiver: receiver,
		orderID:  orderID,
	}
}

func (c *ShipOrderCommand) Execute() error {
	fmt.Println("[ShipOrderCommand] Executing...")
	prevStatus, err := c.receiver.UpdateOrderStatus(c.orderID, domain.OrderShipped)
	if err != nil {
		return err
	}
	c.prevStatus = prevStatus
	return nil
}

func (c *ShipOrderCommand) Undo() error {
	fmt.Println("[ShipOrderCommand] Undoing...")
	_, err := c.receiver.UpdateOrderStatus(c.orderID, c.prevStatus)
	return err
}

func (c *ShipOrderCommand) GetDescription() string {
	return fmt.Sprintf("Ship Order: %s", c.orderID)
}

// OrderInvoker executes commands and maintains history
type OrderInvoker struct {
	history []Command
}

func NewOrderInvoker() *OrderInvoker {
	return &OrderInvoker{
		history: make([]Command, 0),
	}
}

// ExecuteCommand runs a command and adds it to history
func (i *OrderInvoker) ExecuteCommand(cmd Command) error {
	fmt.Printf("\n[Invoker] Executing: %s\n", cmd.GetDescription())
	err := cmd.Execute()
	if err != nil {
		fmt.Printf("[Invoker] ❌ Execution failed: %v\n", err)
		return err
	}
	i.history = append(i.history, cmd)
	fmt.Printf("[Invoker] ✅ Command executed successfully\n")
	return nil
}

// UndoLastCommand reverts the most recent command
func (i *OrderInvoker) UndoLastCommand() error {
	if len(i.history) == 0 {
		return fmt.Errorf("no commands to undo")
	}

	lastCmd := i.history[len(i.history)-1]
	fmt.Printf("\n[Invoker] Undoing: %s\n", lastCmd.GetDescription())

	err := lastCmd.Undo()
	if err != nil {
		fmt.Printf("[Invoker] ❌ Undo failed: %v\n", err)
		return err
	}

	i.history = i.history[:len(i.history)-1]
	fmt.Printf("[Invoker] ✅ Command undone successfully\n")
	return nil
}

// ShowHistory displays all executed commands
func (i *OrderInvoker) ShowHistory() {
	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("COMMAND HISTORY")
	fmt.Println(strings.Repeat("=", 60))
	if len(i.history) == 0 {
		fmt.Println("  No commands executed yet")
	} else {
		for idx, cmd := range i.history {
			fmt.Printf("  %d. %s\n", idx+1, cmd.GetDescription())
		}
	}
	fmt.Println(strings.Repeat("=", 60) + "\n")
}

// MacroCommand executes multiple commands as a single unit
type MacroCommand struct {
	commands []Command
	name     string
}

func NewMacroCommand(name string) *MacroCommand {
	return &MacroCommand{
		commands: make([]Command, 0),
		name:     name,
	}
}

func (m *MacroCommand) AddCommand(cmd Command) {
	m.commands = append(m.commands, cmd)
}

func (m *MacroCommand) Execute() error {
	fmt.Printf("[MacroCommand] Executing macro: %s\n", m.name)
	for _, cmd := range m.commands {
		if err := cmd.Execute(); err != nil {
			return err
		}
		time.Sleep(100 * time.Millisecond) // Simulate processing time
	}
	return nil
}

func (m *MacroCommand) Undo() error {
	fmt.Printf("[MacroCommand] Undoing macro: %s\n", m.name)
	// Undo in reverse order
	for i := len(m.commands) - 1; i >= 0; i-- {
		if err := m.commands[i].Undo(); err != nil {
			return err
		}
	}
	return nil
}

func (m *MacroCommand) GetDescription() string {
	return fmt.Sprintf("Macro: %s (%d commands)", m.name, len(m.commands))
}
