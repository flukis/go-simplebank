// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	context "context"

	db "github.com/flukis/simplebank/db/sqlc"
	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// Store is an autogenerated mock type for the Store type
type Store struct {
	mock.Mock
}

// AddBalanceAccount provides a mock function with given fields: ctx, arg
func (_m *Store) AddBalanceAccount(ctx context.Context, arg db.AddBalanceAccountParams) (db.Account, error) {
	ret := _m.Called(ctx, arg)

	var r0 db.Account
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, db.AddBalanceAccountParams) (db.Account, error)); ok {
		return rf(ctx, arg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, db.AddBalanceAccountParams) db.Account); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Get(0).(db.Account)
	}

	if rf, ok := ret.Get(1).(func(context.Context, db.AddBalanceAccountParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateAccount provides a mock function with given fields: ctx, arg
func (_m *Store) CreateAccount(ctx context.Context, arg db.CreateAccountParams) (db.Account, error) {
	ret := _m.Called(ctx, arg)

	var r0 db.Account
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, db.CreateAccountParams) (db.Account, error)); ok {
		return rf(ctx, arg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, db.CreateAccountParams) db.Account); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Get(0).(db.Account)
	}

	if rf, ok := ret.Get(1).(func(context.Context, db.CreateAccountParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateEntry provides a mock function with given fields: ctx, arg
func (_m *Store) CreateEntry(ctx context.Context, arg db.CreateEntryParams) (db.Entry, error) {
	ret := _m.Called(ctx, arg)

	var r0 db.Entry
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, db.CreateEntryParams) (db.Entry, error)); ok {
		return rf(ctx, arg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, db.CreateEntryParams) db.Entry); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Get(0).(db.Entry)
	}

	if rf, ok := ret.Get(1).(func(context.Context, db.CreateEntryParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateTransfer provides a mock function with given fields: ctx, arg
func (_m *Store) CreateTransfer(ctx context.Context, arg db.CreateTransferParams) (db.Transfer, error) {
	ret := _m.Called(ctx, arg)

	var r0 db.Transfer
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, db.CreateTransferParams) (db.Transfer, error)); ok {
		return rf(ctx, arg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, db.CreateTransferParams) db.Transfer); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Get(0).(db.Transfer)
	}

	if rf, ok := ret.Get(1).(func(context.Context, db.CreateTransferParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateUser provides a mock function with given fields: ctx, arg
func (_m *Store) CreateUser(ctx context.Context, arg db.CreateUserParams) (db.User, error) {
	ret := _m.Called(ctx, arg)

	var r0 db.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, db.CreateUserParams) (db.User, error)); ok {
		return rf(ctx, arg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, db.CreateUserParams) db.User); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Get(0).(db.User)
	}

	if rf, ok := ret.Get(1).(func(context.Context, db.CreateUserParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteAccount provides a mock function with given fields: ctx, id
func (_m *Store) DeleteAccount(ctx context.Context, id int64) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FetchAccounts provides a mock function with given fields: ctx, arg
func (_m *Store) FetchAccounts(ctx context.Context, arg db.FetchAccountsParams) ([]db.Account, error) {
	ret := _m.Called(ctx, arg)

	var r0 []db.Account
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, db.FetchAccountsParams) ([]db.Account, error)); ok {
		return rf(ctx, arg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, db.FetchAccountsParams) []db.Account); ok {
		r0 = rf(ctx, arg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]db.Account)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, db.FetchAccountsParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FetchEntries provides a mock function with given fields: ctx, arg
func (_m *Store) FetchEntries(ctx context.Context, arg db.FetchEntriesParams) ([]db.Entry, error) {
	ret := _m.Called(ctx, arg)

	var r0 []db.Entry
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, db.FetchEntriesParams) ([]db.Entry, error)); ok {
		return rf(ctx, arg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, db.FetchEntriesParams) []db.Entry); ok {
		r0 = rf(ctx, arg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]db.Entry)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, db.FetchEntriesParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FetchTransfer provides a mock function with given fields: ctx, arg
func (_m *Store) FetchTransfer(ctx context.Context, arg db.FetchTransferParams) ([]db.Transfer, error) {
	ret := _m.Called(ctx, arg)

	var r0 []db.Transfer
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, db.FetchTransferParams) ([]db.Transfer, error)); ok {
		return rf(ctx, arg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, db.FetchTransferParams) []db.Transfer); ok {
		r0 = rf(ctx, arg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]db.Transfer)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, db.FetchTransferParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAccount provides a mock function with given fields: ctx, id
func (_m *Store) GetAccount(ctx context.Context, id int64) (db.Account, error) {
	ret := _m.Called(ctx, id)

	var r0 db.Account
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (db.Account, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) db.Account); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(db.Account)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAccountForUpdate provides a mock function with given fields: ctx, id
func (_m *Store) GetAccountForUpdate(ctx context.Context, id int64) (db.Account, error) {
	ret := _m.Called(ctx, id)

	var r0 db.Account
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (db.Account, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) db.Account); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(db.Account)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetEntry provides a mock function with given fields: ctx, id
func (_m *Store) GetEntry(ctx context.Context, id int64) (db.Entry, error) {
	ret := _m.Called(ctx, id)

	var r0 db.Entry
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (db.Entry, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) db.Entry); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(db.Entry)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTransfer provides a mock function with given fields: ctx, id
func (_m *Store) GetTransfer(ctx context.Context, id int64) (db.Transfer, error) {
	ret := _m.Called(ctx, id)

	var r0 db.Transfer
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (db.Transfer, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) db.Transfer); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(db.Transfer)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUser provides a mock function with given fields: ctx, id
func (_m *Store) GetUser(ctx context.Context, id uuid.UUID) (db.User, error) {
	ret := _m.Called(ctx, id)

	var r0 db.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) (db.User, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) db.User); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(db.User)
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserByEmail provides a mock function with given fields: ctx, email
func (_m *Store) GetUserByEmail(ctx context.Context, email string) (db.User, error) {
	ret := _m.Called(ctx, email)

	var r0 db.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (db.User, error)); ok {
		return rf(ctx, email)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) db.User); ok {
		r0 = rf(ctx, email)
	} else {
		r0 = ret.Get(0).(db.User)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserByUsername provides a mock function with given fields: ctx, username
func (_m *Store) GetUserByUsername(ctx context.Context, username string) (db.User, error) {
	ret := _m.Called(ctx, username)

	var r0 db.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (db.User, error)); ok {
		return rf(ctx, username)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) db.User); ok {
		r0 = rf(ctx, username)
	} else {
		r0 = ret.Get(0).(db.User)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TransferTx provides a mock function with given fields: ctx, arg
func (_m *Store) TransferTx(ctx context.Context, arg db.TransferTxParams) (db.TransferTxResult, error) {
	ret := _m.Called(ctx, arg)

	var r0 db.TransferTxResult
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, db.TransferTxParams) (db.TransferTxResult, error)); ok {
		return rf(ctx, arg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, db.TransferTxParams) db.TransferTxResult); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Get(0).(db.TransferTxResult)
	}

	if rf, ok := ret.Get(1).(func(context.Context, db.TransferTxParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateBalanceAccount provides a mock function with given fields: ctx, arg
func (_m *Store) UpdateBalanceAccount(ctx context.Context, arg db.UpdateBalanceAccountParams) (db.Account, error) {
	ret := _m.Called(ctx, arg)

	var r0 db.Account
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, db.UpdateBalanceAccountParams) (db.Account, error)); ok {
		return rf(ctx, arg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, db.UpdateBalanceAccountParams) db.Account); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Get(0).(db.Account)
	}

	if rf, ok := ret.Get(1).(func(context.Context, db.UpdateBalanceAccountParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewStore interface {
	mock.TestingT
	Cleanup(func())
}

// NewStore creates a new instance of Store. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewStore(t mockConstructorTestingTNewStore) *Store {
	mock := &Store{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
