// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package keystoretest

import (
	"sync"
)

// Service is a mock implementation of keystore.KeyStorer.
//
//	func TestSomethingThatUsesKeyStorer(t *testing.T) {
//
//		// make and configure a mocked keystore.KeyStorer
//		mockedKeyStorer := &Service{
//			AddKeyFunc: func(key string, value []byte) error {
//				panic("mock out the AddKey method")
//			},
//			GetKeyFunc: func(key string) ([]byte, error) {
//				panic("mock out the GetKey method")
//			},
//		}
//
//		// use mockedKeyStorer in code that requires keystore.KeyStorer
//		// and then make assertions.
//
//	}
type Service struct {
	// AddKeyFunc mocks the AddKey method.
	AddKeyFunc func(key string, value []byte) error

	// GetKeyFunc mocks the GetKey method.
	GetKeyFunc func(key string) ([]byte, error)

	// calls tracks calls to the methods.
	calls struct {
		// AddKey holds details about calls to the AddKey method.
		AddKey []struct {
			// Key is the key argument value.
			Key string
			// Value is the value argument value.
			Value []byte
		}
		// GetKey holds details about calls to the GetKey method.
		GetKey []struct {
			// Key is the key argument value.
			Key string
		}
	}
	lockAddKey sync.RWMutex
	lockGetKey sync.RWMutex
}

// AddKey calls AddKeyFunc.
func (mock *Service) AddKey(key string, value []byte) error {
	callInfo := struct {
		Key   string
		Value []byte
	}{
		Key:   key,
		Value: value,
	}
	mock.lockAddKey.Lock()
	mock.calls.AddKey = append(mock.calls.AddKey, callInfo)
	mock.lockAddKey.Unlock()
	if mock.AddKeyFunc == nil {
		var (
			errOut error
		)
		return errOut
	}
	return mock.AddKeyFunc(key, value)
}

// AddKeyCalls gets all the calls that were made to AddKey.
// Check the length with:
//
//	len(mockedKeyStorer.AddKeyCalls())
func (mock *Service) AddKeyCalls() []struct {
	Key   string
	Value []byte
} {
	var calls []struct {
		Key   string
		Value []byte
	}
	mock.lockAddKey.RLock()
	calls = mock.calls.AddKey
	mock.lockAddKey.RUnlock()
	return calls
}

// ResetAddKeyCalls reset all the calls that were made to AddKey.
func (mock *Service) ResetAddKeyCalls() {
	mock.lockAddKey.Lock()
	mock.calls.AddKey = nil
	mock.lockAddKey.Unlock()
}

// GetKey calls GetKeyFunc.
func (mock *Service) GetKey(key string) ([]byte, error) {
	callInfo := struct {
		Key string
	}{
		Key: key,
	}
	mock.lockGetKey.Lock()
	mock.calls.GetKey = append(mock.calls.GetKey, callInfo)
	mock.lockGetKey.Unlock()
	if mock.GetKeyFunc == nil {
		var (
			bytesOut []byte
			errOut   error
		)
		return bytesOut, errOut
	}
	return mock.GetKeyFunc(key)
}

// GetKeyCalls gets all the calls that were made to GetKey.
// Check the length with:
//
//	len(mockedKeyStorer.GetKeyCalls())
func (mock *Service) GetKeyCalls() []struct {
	Key string
} {
	var calls []struct {
		Key string
	}
	mock.lockGetKey.RLock()
	calls = mock.calls.GetKey
	mock.lockGetKey.RUnlock()
	return calls
}

// ResetGetKeyCalls reset all the calls that were made to GetKey.
func (mock *Service) ResetGetKeyCalls() {
	mock.lockGetKey.Lock()
	mock.calls.GetKey = nil
	mock.lockGetKey.Unlock()
}

// ResetCalls reset all the calls that were made to all mocked methods.
func (mock *Service) ResetCalls() {
	mock.lockAddKey.Lock()
	mock.calls.AddKey = nil
	mock.lockAddKey.Unlock()

	mock.lockGetKey.Lock()
	mock.calls.GetKey = nil
	mock.lockGetKey.Unlock()
}
