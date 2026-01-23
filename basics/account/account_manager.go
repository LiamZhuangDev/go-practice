package account

import (
	"errors"
)

type Account struct {
	ID      int
	Balance float64
	Name    string
}

type Manager struct {
	accounts []Account
}

func NewManager() *Manager {
	return &Manager{accounts: []Account{}}
}

func (m *Manager) OpenAccount(name string) (a Account) {
	a = Account{
		ID:      len(m.accounts) + 1,
		Balance: 0.0,
		Name:    name,
	}
	m.accounts = append(m.accounts, a)
	return a
}

func (m *Manager) GetBalance(id int) (float64, error) {
	for _, a := range m.accounts {
		if a.ID == id {
			return a.Balance, nil
		}
	}
	return 0.0, errors.New("account not found")
}

func (m *Manager) Deposit(id int, amount float64) error {
	for i := range m.accounts {
		if m.accounts[i].ID == id {
			m.accounts[i].Balance += amount
			return nil
		}
	}
	return errors.New("account not found")
}

func (m *Manager) WithDraw(id int, amount float64) error {
	for i := range m.accounts {
		if m.accounts[i].ID == id {
			if m.accounts[i].Balance < amount {
				return errors.New("insufficient funds")
			}
			m.accounts[i].Balance -= amount
			return nil
		}
	}
	return errors.New("account not found")
}
