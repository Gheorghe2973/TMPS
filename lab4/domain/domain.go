package domain

import "fmt"

// Component represents a computer component
type Component struct {
	ID    string
	Name  string
	Type  string
	Price float64
	Stock int
}

// Computer represents an assembled computer
type Computer struct {
	ID          string
	CPU         Component
	GPU         Component
	RAM         Component
	Storage     Component
	Motherboard Component
	TotalPrice  float64
}

func (c *Computer) CalculateTotal() {
	c.TotalPrice = c.CPU.Price + c.GPU.Price + c.RAM.Price +
		c.Storage.Price + c.Motherboard.Price
}

func (c *Computer) String() string {
	return fmt.Sprintf("Computer[ID=%s, CPU=%s, GPU=%s, RAM=%s, Storage=%s, MB=%s, Price=%.2f]",
		c.ID, c.CPU.Name, c.GPU.Name, c.RAM.Name, c.Storage.Name, c.Motherboard.Name, c.TotalPrice)
}

// OrderStatus represents the current status of an order
type OrderStatus string

const (
	OrderPending    OrderStatus = "PENDING"
	OrderProcessing OrderStatus = "PROCESSING"
	OrderShipped    OrderStatus = "SHIPPED"
	OrderDelivered  OrderStatus = "DELIVERED"
	OrderCancelled  OrderStatus = "CANCELLED"
)

// Order represents a customer order
type Order struct {
	ID         string
	CustomerID string
	Computer   Computer
	Status     OrderStatus
	TotalPrice float64
	CreatedAt  string
}

func (o *Order) String() string {
	return fmt.Sprintf("Order[ID=%s, Customer=%s, Status=%s, Price=%.2f]",
		o.ID, o.CustomerID, o.Status, o.TotalPrice)
}
