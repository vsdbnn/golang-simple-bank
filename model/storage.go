package model

import (
	"sync"

	"github.com/shopspring/decimal"
)

type memAccountStorage struct {
	mutex   sync.RWMutex
	storage map[string]Account
}

func NewMemStorage() *memAccountStorage {
	return &memAccountStorage{
		mutex:   sync.RWMutex{},
		storage: map[string]Account{},
	}
}

func (s *memAccountStorage) GetAccount(id string) (Account, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	acc, ok := s.storage[id];
	if !ok {
		return Account{}, ErrNotFound
	}

	return acc, nil
}

func (s *memAccountStorage) AddAccount(a Account) error {
	if a.Balance.IsNegative() {
		return ErrIncorrectBalance
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()

	if _, ok := s.storage[a.ID]; ok {
		return ErrAlreadyExists
	}

	s.storage[a.ID] = a

	return nil
}

func (s *memAccountStorage) Transfer(
	sender, receiver string,
	amount decimal.Decimal,
) error {

	if sender == receiver {
		return ErrSelfPayment
	}

	if amount.IsNegative() || amount.IsZero() {
		return ErrNegativePayment
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()

	sAcc, ok := s.storage[sender]
	if !ok {
		return ErrNotFound
	}

	rAcc, ok := s.storage[receiver]
	if !ok {
		return ErrNotFound
	}

	if sAcc.Balance.LessThan(amount) {
		return ErrLowBalance
	}

	sAcc.Balance = sAcc.Balance.Sub(amount)
	rAcc.Balance = rAcc.Balance.Add(amount)

	s.storage[sender] = sAcc
	s.storage[receiver] = rAcc

	return nil
}
