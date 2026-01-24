package closure

import "fmt"

// Example 1: A simple closure that maintains a counter
func NewCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

// Example 2: MakeHandler creates a callback function that tracks the number of times it is called.
type Button struct {
	Label string
}

func (b *Button) OnClick(handle func()) {
	// Handle button click
	handle()
}

// The handler using closure
func MakeHandler(name string) func() {
	calls := 0
	return func() {
		calls++
		fmt.Printf("%s was called %d times\n", name, calls)
	}
}

// The handler using struct and method
type Handler struct {
	name  string
	calls int
}

func NewHandler(name string) *Handler {
	return &Handler{name: name, calls: 0}
}

func (h *Handler) Handle() {
	h.calls++
	fmt.Printf("%s was called %d times\n", h.name, h.calls)
}

// Example 3: Price strategies using closures and factory functions.
// Each factory function returns a pricing strategy that captures
// variables from the outer function (off, coupon, strategies).

// makeAdder is a factory function that returns a closure.
func makeAdder(base int) func(int) int {
	return func(x int) int {
		return base + x
	}
}

// add10 := makeAdder(10)
// add100 := makeAdder(100)

// fmt.Println(add10(5))   // 15
// fmt.Println(add100(5))  // 105

type PriceStrategy func(price float64) float64

func DiscountStrategy(off float64) PriceStrategy {
	return func(price float64) float64 {
		return price * (1 - off)
	}
}

func WithCoupon(coupon float64) PriceStrategy {
	return func(price float64) float64 {
		return price - coupon
	}
}

func Combine(strategies ...PriceStrategy) PriceStrategy {
	return func(price float64) float64 {
		for _, s := range strategies {
			price = s(price)
		}
		return price
	}
}
