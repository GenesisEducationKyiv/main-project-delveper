// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mocks

import (
	"context"
	"github.com/GenesisEducationKyiv/main-project-delveper/internal/rate"
	"sync"
)

// Ensure, that ExchangeRateServiceMock does implement rate.ExchangeRateService.
// If this is not the case, regenerate this file with moq.
var _ rate.ExchangeRateService = &ExchangeRateServiceMock{}

// ExchangeRateServiceMock is a mock implementation of rate.ExchangeRateService.
//
//	func TestSomethingThatUsesExchangeRateService(t *testing.T) {
//
//		// make and configure a mocked rate.ExchangeRateService
//		mockedExchangeRateService := &ExchangeRateServiceMock{
//			GetFunc: func(ctx context.Context, currency rate.CurrencyPair) (*rate.ExchangeRate, error) {
//				panic("mock out the Get method")
//			},
//		}
//
//		// use mockedExchangeRateService in code that requires rate.ExchangeRateService
//		// and then make assertions.
//
//	}
type ExchangeRateServiceMock struct {
	// GetFunc mocks the Get method.
	GetFunc func(ctx context.Context, currency rate.CurrencyPair) (*rate.ExchangeRate, error)

	// calls tracks calls to the methods.
	calls struct {
		// Get holds details about calls to the Get method.
		Get []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Currency is the currency argument value.
			Currency rate.CurrencyPair
		}
	}
	lockGet sync.RWMutex
}

// Get calls GetFunc.
func (mock *ExchangeRateServiceMock) Get(ctx context.Context, currency rate.CurrencyPair) (*rate.ExchangeRate, error) {
	if mock.GetFunc == nil {
		panic("ExchangeRateServiceMock.GetFunc: method is nil but ExchangeRateService.Get was just called")
	}
	callInfo := struct {
		Ctx      context.Context
		Currency rate.CurrencyPair
	}{
		Ctx:      ctx,
		Currency: currency,
	}
	mock.lockGet.Lock()
	mock.calls.Get = append(mock.calls.Get, callInfo)
	mock.lockGet.Unlock()
	return mock.GetFunc(ctx, currency)
}

// GetCalls gets all the calls that were made to Get.
// Check the length with:
//
//	len(mockedExchangeRateService.GetCalls())
func (mock *ExchangeRateServiceMock) GetCalls() []struct {
	Ctx      context.Context
	Currency rate.CurrencyPair
} {
	var calls []struct {
		Ctx      context.Context
		Currency rate.CurrencyPair
	}
	mock.lockGet.RLock()
	calls = mock.calls.Get
	mock.lockGet.RUnlock()
	return calls
}