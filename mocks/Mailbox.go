// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	client "github.com/textileio/textile/v2/api/users/client"
	cmd "github.com/textileio/textile/v2/cmd"

	context "context"

	local "github.com/textileio/textile/v2/mail/local"

	mock "github.com/stretchr/testify/mock"

	thread "github.com/textileio/go-threads/core/thread"
)

// Mailbox is an autogenerated mock type for the Mailbox type
type Mailbox struct {
	mock.Mock
}

// Identity provides a mock function with given fields:
func (_m *Mailbox) Identity() thread.Identity {
	ret := _m.Called()

	var r0 thread.Identity
	if rf, ok := ret.Get(0).(func() thread.Identity); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(thread.Identity)
		}
	}

	return r0
}

// ListInboxMessages provides a mock function with given fields: ctx, opts
func (_m *Mailbox) ListInboxMessages(ctx context.Context, opts ...client.ListOption) ([]client.Message, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 []client.Message
	if rf, ok := ret.Get(0).(func(context.Context, ...client.ListOption) []client.Message); ok {
		r0 = rf(ctx, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]client.Message)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, ...client.ListOption) error); ok {
		r1 = rf(ctx, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SendMessage provides a mock function with given fields: ctx, to, body
func (_m *Mailbox) SendMessage(ctx context.Context, to thread.PubKey, body []byte) (client.Message, error) {
	ret := _m.Called(ctx, to, body)

	var r0 client.Message
	if rf, ok := ret.Get(0).(func(context.Context, thread.PubKey, []byte) client.Message); ok {
		r0 = rf(ctx, to, body)
	} else {
		r0 = ret.Get(0).(client.Message)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, thread.PubKey, []byte) error); ok {
		r1 = rf(ctx, to, body)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WatchInbox provides a mock function with given fields: ctx, mevents, offline
func (_m *Mailbox) WatchInbox(ctx context.Context, mevents chan<- local.MailboxEvent, offline bool) (<-chan cmd.WatchState, error) {
	ret := _m.Called(ctx, mevents, offline)

	var r0 <-chan cmd.WatchState
	if rf, ok := ret.Get(0).(func(context.Context, chan<- local.MailboxEvent, bool) <-chan cmd.WatchState); ok {
		r0 = rf(ctx, mevents, offline)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan cmd.WatchState)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, chan<- local.MailboxEvent, bool) error); ok {
		r1 = rf(ctx, mevents, offline)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
