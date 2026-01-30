package interface_example

import "fmt"

type Payment interface {
	Pay(amount float64) error
	Name() string
}

type WeChatPay struct {
	Account string
}

func (wp *WeChatPay) Pay(amount float64) error {
	fmt.Printf("[WeChatPay] %s paid %.2f yuan\n", wp.Account, amount)
	return nil
}

func (wp *WeChatPay) Name() string {
	return "WeChatPay"
}

type AliPay struct {
	Account string
}

func (ap *AliPay) Pay(amount float64) error {
	fmt.Printf("[AliPay] %s paid %.2f yuan\n", ap.Account, amount)
	return nil
}

func (ap *AliPay) Name() string {
	return "AliPay"
}

type BankCard struct {
	CardNumber string
}

func (b *BankCard) Pay(amount float64) error {
	fmt.Printf("[BankCard] %s paid %.2f yuan\n", b.CardNumber, amount)
	return nil
}

func (c *BankCard) Name() string {
	return "BankCard"
}

func ProcessPayment(p Payment, amount float64) error {
	fmt.Printf("Processing payment via %s with amount %.2f\n", p.Name(), amount)
	return p.Pay(amount)
}

func PaymentTest() {
	paymentMethods := []Payment{
		&WeChatPay{Account: "wx-123456"},
		&AliPay{Account: "alice@alipay"},
		&BankCard{CardNumber: "62223098456009"},
	}

	for _, p := range paymentMethods {
		if err := ProcessPayment(p, 100); err != nil {
			fmt.Println("Payment failed: ", err)
		}
	}
}
