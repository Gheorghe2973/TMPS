package strategy

import (
	"fmt"
	"lab4/domain"
	"strings"
)

// PricingStrategy defines the interface for different pricing strategies
type PricingStrategy interface {
	CalculatePrice(computer *domain.Computer) float64
	GetName() string
}

// RetailPricingStrategy is the standard pricing for retail customers
type RetailPricingStrategy struct{}

func NewRetailPricingStrategy() *RetailPricingStrategy {
	return &RetailPricingStrategy{}
}

func (s *RetailPricingStrategy) CalculatePrice(computer *domain.Computer) float64 {
	computer.CalculateTotal()
	// Standard retail price - no discount
	return computer.TotalPrice
}

func (s *RetailPricingStrategy) GetName() string {
	return "Retail Pricing (No discount)"
}

// WholesalePricingStrategy offers discounts for bulk/business customers
type WholesalePricingStrategy struct {
	discountPercent float64
}

func NewWholesalePricingStrategy(discountPercent float64) *WholesalePricingStrategy {
	return &WholesalePricingStrategy{
		discountPercent: discountPercent,
	}
}

func (s *WholesalePricingStrategy) CalculatePrice(computer *domain.Computer) float64 {
	computer.CalculateTotal()
	discount := computer.TotalPrice * (s.discountPercent / 100)
	return computer.TotalPrice - discount
}

func (s *WholesalePricingStrategy) GetName() string {
	return fmt.Sprintf("Wholesale Pricing (%.0f%% discount)", s.discountPercent)
}

// SeasonalPricingStrategy offers time-limited discounts
type SeasonalPricingStrategy struct {
	seasonName      string
	discountPercent float64
}

func NewSeasonalPricingStrategy(seasonName string, discountPercent float64) *SeasonalPricingStrategy {
	return &SeasonalPricingStrategy{
		seasonName:      seasonName,
		discountPercent: discountPercent,
	}
}

func (s *SeasonalPricingStrategy) CalculatePrice(computer *domain.Computer) float64 {
	computer.CalculateTotal()
	discount := computer.TotalPrice * (s.discountPercent / 100)
	return computer.TotalPrice - discount
}

func (s *SeasonalPricingStrategy) GetName() string {
	return fmt.Sprintf("%s Sale (%.0f%% discount)", s.seasonName, s.discountPercent)
}

// StudentPricingStrategy offers educational discounts
type StudentPricingStrategy struct {
	baseDiscount float64
}

func NewStudentPricingStrategy(baseDiscount float64) *StudentPricingStrategy {
	return &StudentPricingStrategy{
		baseDiscount: baseDiscount,
	}
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

func (s *StudentPricingStrategy) GetName() string {
	return fmt.Sprintf("Student Pricing (%.0f%% + extra discounts)", s.baseDiscount)
}

// PriceCalculator is the context that uses pricing strategies
type PriceCalculator struct {
	strategy PricingStrategy
}

func NewPriceCalculator(strategy PricingStrategy) *PriceCalculator {
	return &PriceCalculator{
		strategy: strategy,
	}
}

// SetStrategy allows changing the pricing strategy at runtime
func (pc *PriceCalculator) SetStrategy(strategy PricingStrategy) {
	pc.strategy = strategy
	fmt.Printf("[PriceCalculator] Strategy changed to: %s\n", strategy.GetName())
}

// CalculateFinalPrice calculates the price using the current strategy
func (pc *PriceCalculator) CalculateFinalPrice(computer *domain.Computer) float64 {
	fmt.Printf("\n[PriceCalculator] Calculating price using: %s\n", pc.strategy.GetName())

	originalPrice := computer.CPU.Price + computer.GPU.Price + computer.RAM.Price +
		computer.Storage.Price + computer.Motherboard.Price

	finalPrice := pc.strategy.CalculatePrice(computer)
	savings := originalPrice - finalPrice

	fmt.Printf("  Original Price: $%.2f\n", originalPrice)
	if savings > 0 {
		fmt.Printf("  Discount: -$%.2f (%.1f%%)\n", savings, (savings/originalPrice)*100)
	}
	fmt.Printf("  Final Price: $%.2f\n\n", finalPrice)

	return finalPrice
}

// CompareStrategies shows pricing across different strategies
func CompareStrategies(computer *domain.Computer, strategies []PricingStrategy) {
	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("PRICING COMPARISON")
	fmt.Println(strings.Repeat("=", 60))

	for _, strategy := range strategies {
		calculator := NewPriceCalculator(strategy)
		calculator.CalculateFinalPrice(computer)
	}
}
