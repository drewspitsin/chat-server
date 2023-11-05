package mocks

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

//go:generate minimock -i github.com/drewspitsin/chat-server/internal/service.ChatService -o ./mocks/chat_service_minimock.go -n ChatServiceMock

import (
	"context"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/drewspitsin/chat-server/internal/model"
	"github.com/gojuno/minimock/v3"
)

// ChatServiceMock implements service.ChatService
type ChatServiceMock struct {
	t minimock.Tester

	funcCreate          func(ctx context.Context, info *model.Chat) (i1 int64, err error)
	inspectFuncCreate   func(ctx context.Context, info *model.Chat)
	afterCreateCounter  uint64
	beforeCreateCounter uint64
	CreateMock          mChatServiceMockCreate

	funcDelete          func(ctx context.Context, id int64) (err error)
	inspectFuncDelete   func(ctx context.Context, id int64)
	afterDeleteCounter  uint64
	beforeDeleteCounter uint64
	DeleteMock          mChatServiceMockDelete

	funcSend          func(ctx context.Context, info *model.Chat) (err error)
	inspectFuncSend   func(ctx context.Context, info *model.Chat)
	afterSendCounter  uint64
	beforeSendCounter uint64
	SendMock          mChatServiceMockSend
}

// NewChatServiceMock returns a mock for service.ChatService
func NewChatServiceMock(t minimock.Tester) *ChatServiceMock {
	m := &ChatServiceMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.CreateMock = mChatServiceMockCreate{mock: m}
	m.CreateMock.callArgs = []*ChatServiceMockCreateParams{}

	m.DeleteMock = mChatServiceMockDelete{mock: m}
	m.DeleteMock.callArgs = []*ChatServiceMockDeleteParams{}

	m.SendMock = mChatServiceMockSend{mock: m}
	m.SendMock.callArgs = []*ChatServiceMockSendParams{}

	return m
}

type mChatServiceMockCreate struct {
	mock               *ChatServiceMock
	defaultExpectation *ChatServiceMockCreateExpectation
	expectations       []*ChatServiceMockCreateExpectation

	callArgs []*ChatServiceMockCreateParams
	mutex    sync.RWMutex
}

// ChatServiceMockCreateExpectation specifies expectation struct of the ChatService.Create
type ChatServiceMockCreateExpectation struct {
	mock    *ChatServiceMock
	params  *ChatServiceMockCreateParams
	results *ChatServiceMockCreateResults
	Counter uint64
}

// ChatServiceMockCreateParams contains parameters of the ChatService.Create
type ChatServiceMockCreateParams struct {
	ctx  context.Context
	info *model.Chat
}

// ChatServiceMockCreateResults contains results of the ChatService.Create
type ChatServiceMockCreateResults struct {
	i1  int64
	err error
}

// Expect sets up expected params for ChatService.Create
func (mmCreate *mChatServiceMockCreate) Expect(ctx context.Context, info *model.Chat) *mChatServiceMockCreate {
	if mmCreate.mock.funcCreate != nil {
		mmCreate.mock.t.Fatalf("ChatServiceMock.Create mock is already set by Set")
	}

	if mmCreate.defaultExpectation == nil {
		mmCreate.defaultExpectation = &ChatServiceMockCreateExpectation{}
	}

	mmCreate.defaultExpectation.params = &ChatServiceMockCreateParams{ctx, info}
	for _, e := range mmCreate.expectations {
		if minimock.Equal(e.params, mmCreate.defaultExpectation.params) {
			mmCreate.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmCreate.defaultExpectation.params)
		}
	}

	return mmCreate
}

// Inspect accepts an inspector function that has same arguments as the ChatService.Create
func (mmCreate *mChatServiceMockCreate) Inspect(f func(ctx context.Context, info *model.Chat)) *mChatServiceMockCreate {
	if mmCreate.mock.inspectFuncCreate != nil {
		mmCreate.mock.t.Fatalf("Inspect function is already set for ChatServiceMock.Create")
	}

	mmCreate.mock.inspectFuncCreate = f

	return mmCreate
}

// Return sets up results that will be returned by ChatService.Create
func (mmCreate *mChatServiceMockCreate) Return(i1 int64, err error) *ChatServiceMock {
	if mmCreate.mock.funcCreate != nil {
		mmCreate.mock.t.Fatalf("ChatServiceMock.Create mock is already set by Set")
	}

	if mmCreate.defaultExpectation == nil {
		mmCreate.defaultExpectation = &ChatServiceMockCreateExpectation{mock: mmCreate.mock}
	}
	mmCreate.defaultExpectation.results = &ChatServiceMockCreateResults{i1, err}
	return mmCreate.mock
}

// Set uses given function f to mock the ChatService.Create method
func (mmCreate *mChatServiceMockCreate) Set(f func(ctx context.Context, info *model.Chat) (i1 int64, err error)) *ChatServiceMock {
	if mmCreate.defaultExpectation != nil {
		mmCreate.mock.t.Fatalf("Default expectation is already set for the ChatService.Create method")
	}

	if len(mmCreate.expectations) > 0 {
		mmCreate.mock.t.Fatalf("Some expectations are already set for the ChatService.Create method")
	}

	mmCreate.mock.funcCreate = f
	return mmCreate.mock
}

// When sets expectation for the ChatService.Create which will trigger the result defined by the following
// Then helper
func (mmCreate *mChatServiceMockCreate) When(ctx context.Context, info *model.Chat) *ChatServiceMockCreateExpectation {
	if mmCreate.mock.funcCreate != nil {
		mmCreate.mock.t.Fatalf("ChatServiceMock.Create mock is already set by Set")
	}

	expectation := &ChatServiceMockCreateExpectation{
		mock:   mmCreate.mock,
		params: &ChatServiceMockCreateParams{ctx, info},
	}
	mmCreate.expectations = append(mmCreate.expectations, expectation)
	return expectation
}

// Then sets up ChatService.Create return parameters for the expectation previously defined by the When method
func (e *ChatServiceMockCreateExpectation) Then(i1 int64, err error) *ChatServiceMock {
	e.results = &ChatServiceMockCreateResults{i1, err}
	return e.mock
}

// Create implements service.ChatService
func (mmCreate *ChatServiceMock) Create(ctx context.Context, info *model.Chat) (i1 int64, err error) {
	mm_atomic.AddUint64(&mmCreate.beforeCreateCounter, 1)
	defer mm_atomic.AddUint64(&mmCreate.afterCreateCounter, 1)

	if mmCreate.inspectFuncCreate != nil {
		mmCreate.inspectFuncCreate(ctx, info)
	}

	mm_params := &ChatServiceMockCreateParams{ctx, info}

	// Record call args
	mmCreate.CreateMock.mutex.Lock()
	mmCreate.CreateMock.callArgs = append(mmCreate.CreateMock.callArgs, mm_params)
	mmCreate.CreateMock.mutex.Unlock()

	for _, e := range mmCreate.CreateMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.i1, e.results.err
		}
	}

	if mmCreate.CreateMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmCreate.CreateMock.defaultExpectation.Counter, 1)
		mm_want := mmCreate.CreateMock.defaultExpectation.params
		mm_got := ChatServiceMockCreateParams{ctx, info}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmCreate.t.Errorf("ChatServiceMock.Create got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmCreate.CreateMock.defaultExpectation.results
		if mm_results == nil {
			mmCreate.t.Fatal("No results are set for the ChatServiceMock.Create")
		}
		return (*mm_results).i1, (*mm_results).err
	}
	if mmCreate.funcCreate != nil {
		return mmCreate.funcCreate(ctx, info)
	}
	mmCreate.t.Fatalf("Unexpected call to ChatServiceMock.Create. %v %v", ctx, info)
	return
}

// CreateAfterCounter returns a count of finished ChatServiceMock.Create invocations
func (mmCreate *ChatServiceMock) CreateAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCreate.afterCreateCounter)
}

// CreateBeforeCounter returns a count of ChatServiceMock.Create invocations
func (mmCreate *ChatServiceMock) CreateBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCreate.beforeCreateCounter)
}

// Calls returns a list of arguments used in each call to ChatServiceMock.Create.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmCreate *mChatServiceMockCreate) Calls() []*ChatServiceMockCreateParams {
	mmCreate.mutex.RLock()

	argCopy := make([]*ChatServiceMockCreateParams, len(mmCreate.callArgs))
	copy(argCopy, mmCreate.callArgs)

	mmCreate.mutex.RUnlock()

	return argCopy
}

// MinimockCreateDone returns true if the count of the Create invocations corresponds
// the number of defined expectations
func (m *ChatServiceMock) MinimockCreateDone() bool {
	for _, e := range m.CreateMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.CreateMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterCreateCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcCreate != nil && mm_atomic.LoadUint64(&m.afterCreateCounter) < 1 {
		return false
	}
	return true
}

// MinimockCreateInspect logs each unmet expectation
func (m *ChatServiceMock) MinimockCreateInspect() {
	for _, e := range m.CreateMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to ChatServiceMock.Create with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.CreateMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterCreateCounter) < 1 {
		if m.CreateMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to ChatServiceMock.Create")
		} else {
			m.t.Errorf("Expected call to ChatServiceMock.Create with params: %#v", *m.CreateMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcCreate != nil && mm_atomic.LoadUint64(&m.afterCreateCounter) < 1 {
		m.t.Error("Expected call to ChatServiceMock.Create")
	}
}

type mChatServiceMockDelete struct {
	mock               *ChatServiceMock
	defaultExpectation *ChatServiceMockDeleteExpectation
	expectations       []*ChatServiceMockDeleteExpectation

	callArgs []*ChatServiceMockDeleteParams
	mutex    sync.RWMutex
}

// ChatServiceMockDeleteExpectation specifies expectation struct of the ChatService.Delete
type ChatServiceMockDeleteExpectation struct {
	mock    *ChatServiceMock
	params  *ChatServiceMockDeleteParams
	results *ChatServiceMockDeleteResults
	Counter uint64
}

// ChatServiceMockDeleteParams contains parameters of the ChatService.Delete
type ChatServiceMockDeleteParams struct {
	ctx context.Context
	id  int64
}

// ChatServiceMockDeleteResults contains results of the ChatService.Delete
type ChatServiceMockDeleteResults struct {
	err error
}

// Expect sets up expected params for ChatService.Delete
func (mmDelete *mChatServiceMockDelete) Expect(ctx context.Context, id int64) *mChatServiceMockDelete {
	if mmDelete.mock.funcDelete != nil {
		mmDelete.mock.t.Fatalf("ChatServiceMock.Delete mock is already set by Set")
	}

	if mmDelete.defaultExpectation == nil {
		mmDelete.defaultExpectation = &ChatServiceMockDeleteExpectation{}
	}

	mmDelete.defaultExpectation.params = &ChatServiceMockDeleteParams{ctx, id}
	for _, e := range mmDelete.expectations {
		if minimock.Equal(e.params, mmDelete.defaultExpectation.params) {
			mmDelete.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmDelete.defaultExpectation.params)
		}
	}

	return mmDelete
}

// Inspect accepts an inspector function that has same arguments as the ChatService.Delete
func (mmDelete *mChatServiceMockDelete) Inspect(f func(ctx context.Context, id int64)) *mChatServiceMockDelete {
	if mmDelete.mock.inspectFuncDelete != nil {
		mmDelete.mock.t.Fatalf("Inspect function is already set for ChatServiceMock.Delete")
	}

	mmDelete.mock.inspectFuncDelete = f

	return mmDelete
}

// Return sets up results that will be returned by ChatService.Delete
func (mmDelete *mChatServiceMockDelete) Return(err error) *ChatServiceMock {
	if mmDelete.mock.funcDelete != nil {
		mmDelete.mock.t.Fatalf("ChatServiceMock.Delete mock is already set by Set")
	}

	if mmDelete.defaultExpectation == nil {
		mmDelete.defaultExpectation = &ChatServiceMockDeleteExpectation{mock: mmDelete.mock}
	}
	mmDelete.defaultExpectation.results = &ChatServiceMockDeleteResults{err}
	return mmDelete.mock
}

// Set uses given function f to mock the ChatService.Delete method
func (mmDelete *mChatServiceMockDelete) Set(f func(ctx context.Context, id int64) (err error)) *ChatServiceMock {
	if mmDelete.defaultExpectation != nil {
		mmDelete.mock.t.Fatalf("Default expectation is already set for the ChatService.Delete method")
	}

	if len(mmDelete.expectations) > 0 {
		mmDelete.mock.t.Fatalf("Some expectations are already set for the ChatService.Delete method")
	}

	mmDelete.mock.funcDelete = f
	return mmDelete.mock
}

// When sets expectation for the ChatService.Delete which will trigger the result defined by the following
// Then helper
func (mmDelete *mChatServiceMockDelete) When(ctx context.Context, id int64) *ChatServiceMockDeleteExpectation {
	if mmDelete.mock.funcDelete != nil {
		mmDelete.mock.t.Fatalf("ChatServiceMock.Delete mock is already set by Set")
	}

	expectation := &ChatServiceMockDeleteExpectation{
		mock:   mmDelete.mock,
		params: &ChatServiceMockDeleteParams{ctx, id},
	}
	mmDelete.expectations = append(mmDelete.expectations, expectation)
	return expectation
}

// Then sets up ChatService.Delete return parameters for the expectation previously defined by the When method
func (e *ChatServiceMockDeleteExpectation) Then(err error) *ChatServiceMock {
	e.results = &ChatServiceMockDeleteResults{err}
	return e.mock
}

// Delete implements service.ChatService
func (mmDelete *ChatServiceMock) Delete(ctx context.Context, id int64) (err error) {
	mm_atomic.AddUint64(&mmDelete.beforeDeleteCounter, 1)
	defer mm_atomic.AddUint64(&mmDelete.afterDeleteCounter, 1)

	if mmDelete.inspectFuncDelete != nil {
		mmDelete.inspectFuncDelete(ctx, id)
	}

	mm_params := &ChatServiceMockDeleteParams{ctx, id}

	// Record call args
	mmDelete.DeleteMock.mutex.Lock()
	mmDelete.DeleteMock.callArgs = append(mmDelete.DeleteMock.callArgs, mm_params)
	mmDelete.DeleteMock.mutex.Unlock()

	for _, e := range mmDelete.DeleteMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmDelete.DeleteMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmDelete.DeleteMock.defaultExpectation.Counter, 1)
		mm_want := mmDelete.DeleteMock.defaultExpectation.params
		mm_got := ChatServiceMockDeleteParams{ctx, id}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmDelete.t.Errorf("ChatServiceMock.Delete got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmDelete.DeleteMock.defaultExpectation.results
		if mm_results == nil {
			mmDelete.t.Fatal("No results are set for the ChatServiceMock.Delete")
		}
		return (*mm_results).err
	}
	if mmDelete.funcDelete != nil {
		return mmDelete.funcDelete(ctx, id)
	}
	mmDelete.t.Fatalf("Unexpected call to ChatServiceMock.Delete. %v %v", ctx, id)
	return
}

// DeleteAfterCounter returns a count of finished ChatServiceMock.Delete invocations
func (mmDelete *ChatServiceMock) DeleteAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmDelete.afterDeleteCounter)
}

// DeleteBeforeCounter returns a count of ChatServiceMock.Delete invocations
func (mmDelete *ChatServiceMock) DeleteBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmDelete.beforeDeleteCounter)
}

// Calls returns a list of arguments used in each call to ChatServiceMock.Delete.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmDelete *mChatServiceMockDelete) Calls() []*ChatServiceMockDeleteParams {
	mmDelete.mutex.RLock()

	argCopy := make([]*ChatServiceMockDeleteParams, len(mmDelete.callArgs))
	copy(argCopy, mmDelete.callArgs)

	mmDelete.mutex.RUnlock()

	return argCopy
}

// MinimockDeleteDone returns true if the count of the Delete invocations corresponds
// the number of defined expectations
func (m *ChatServiceMock) MinimockDeleteDone() bool {
	for _, e := range m.DeleteMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.DeleteMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterDeleteCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcDelete != nil && mm_atomic.LoadUint64(&m.afterDeleteCounter) < 1 {
		return false
	}
	return true
}

// MinimockDeleteInspect logs each unmet expectation
func (m *ChatServiceMock) MinimockDeleteInspect() {
	for _, e := range m.DeleteMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to ChatServiceMock.Delete with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.DeleteMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterDeleteCounter) < 1 {
		if m.DeleteMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to ChatServiceMock.Delete")
		} else {
			m.t.Errorf("Expected call to ChatServiceMock.Delete with params: %#v", *m.DeleteMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcDelete != nil && mm_atomic.LoadUint64(&m.afterDeleteCounter) < 1 {
		m.t.Error("Expected call to ChatServiceMock.Delete")
	}
}

type mChatServiceMockSend struct {
	mock               *ChatServiceMock
	defaultExpectation *ChatServiceMockSendExpectation
	expectations       []*ChatServiceMockSendExpectation

	callArgs []*ChatServiceMockSendParams
	mutex    sync.RWMutex
}

// ChatServiceMockSendExpectation specifies expectation struct of the ChatService.Send
type ChatServiceMockSendExpectation struct {
	mock    *ChatServiceMock
	params  *ChatServiceMockSendParams
	results *ChatServiceMockSendResults
	Counter uint64
}

// ChatServiceMockSendParams contains parameters of the ChatService.Send
type ChatServiceMockSendParams struct {
	ctx  context.Context
	info *model.Chat
}

// ChatServiceMockSendResults contains results of the ChatService.Send
type ChatServiceMockSendResults struct {
	err error
}

// Expect sets up expected params for ChatService.Send
func (mmSend *mChatServiceMockSend) Expect(ctx context.Context, info *model.Chat) *mChatServiceMockSend {
	if mmSend.mock.funcSend != nil {
		mmSend.mock.t.Fatalf("ChatServiceMock.Send mock is already set by Set")
	}

	if mmSend.defaultExpectation == nil {
		mmSend.defaultExpectation = &ChatServiceMockSendExpectation{}
	}

	mmSend.defaultExpectation.params = &ChatServiceMockSendParams{ctx, info}
	for _, e := range mmSend.expectations {
		if minimock.Equal(e.params, mmSend.defaultExpectation.params) {
			mmSend.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmSend.defaultExpectation.params)
		}
	}

	return mmSend
}

// Inspect accepts an inspector function that has same arguments as the ChatService.Send
func (mmSend *mChatServiceMockSend) Inspect(f func(ctx context.Context, info *model.Chat)) *mChatServiceMockSend {
	if mmSend.mock.inspectFuncSend != nil {
		mmSend.mock.t.Fatalf("Inspect function is already set for ChatServiceMock.Send")
	}

	mmSend.mock.inspectFuncSend = f

	return mmSend
}

// Return sets up results that will be returned by ChatService.Send
func (mmSend *mChatServiceMockSend) Return(err error) *ChatServiceMock {
	if mmSend.mock.funcSend != nil {
		mmSend.mock.t.Fatalf("ChatServiceMock.Send mock is already set by Set")
	}

	if mmSend.defaultExpectation == nil {
		mmSend.defaultExpectation = &ChatServiceMockSendExpectation{mock: mmSend.mock}
	}
	mmSend.defaultExpectation.results = &ChatServiceMockSendResults{err}
	return mmSend.mock
}

// Set uses given function f to mock the ChatService.Send method
func (mmSend *mChatServiceMockSend) Set(f func(ctx context.Context, info *model.Chat) (err error)) *ChatServiceMock {
	if mmSend.defaultExpectation != nil {
		mmSend.mock.t.Fatalf("Default expectation is already set for the ChatService.Send method")
	}

	if len(mmSend.expectations) > 0 {
		mmSend.mock.t.Fatalf("Some expectations are already set for the ChatService.Send method")
	}

	mmSend.mock.funcSend = f
	return mmSend.mock
}

// When sets expectation for the ChatService.Send which will trigger the result defined by the following
// Then helper
func (mmSend *mChatServiceMockSend) When(ctx context.Context, info *model.Chat) *ChatServiceMockSendExpectation {
	if mmSend.mock.funcSend != nil {
		mmSend.mock.t.Fatalf("ChatServiceMock.Send mock is already set by Set")
	}

	expectation := &ChatServiceMockSendExpectation{
		mock:   mmSend.mock,
		params: &ChatServiceMockSendParams{ctx, info},
	}
	mmSend.expectations = append(mmSend.expectations, expectation)
	return expectation
}

// Then sets up ChatService.Send return parameters for the expectation previously defined by the When method
func (e *ChatServiceMockSendExpectation) Then(err error) *ChatServiceMock {
	e.results = &ChatServiceMockSendResults{err}
	return e.mock
}

// Send implements service.ChatService
func (mmSend *ChatServiceMock) Send(ctx context.Context, info *model.Chat) (err error) {
	mm_atomic.AddUint64(&mmSend.beforeSendCounter, 1)
	defer mm_atomic.AddUint64(&mmSend.afterSendCounter, 1)

	if mmSend.inspectFuncSend != nil {
		mmSend.inspectFuncSend(ctx, info)
	}

	mm_params := &ChatServiceMockSendParams{ctx, info}

	// Record call args
	mmSend.SendMock.mutex.Lock()
	mmSend.SendMock.callArgs = append(mmSend.SendMock.callArgs, mm_params)
	mmSend.SendMock.mutex.Unlock()

	for _, e := range mmSend.SendMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmSend.SendMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmSend.SendMock.defaultExpectation.Counter, 1)
		mm_want := mmSend.SendMock.defaultExpectation.params
		mm_got := ChatServiceMockSendParams{ctx, info}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmSend.t.Errorf("ChatServiceMock.Send got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmSend.SendMock.defaultExpectation.results
		if mm_results == nil {
			mmSend.t.Fatal("No results are set for the ChatServiceMock.Send")
		}
		return (*mm_results).err
	}
	if mmSend.funcSend != nil {
		return mmSend.funcSend(ctx, info)
	}
	mmSend.t.Fatalf("Unexpected call to ChatServiceMock.Send. %v %v", ctx, info)
	return
}

// SendAfterCounter returns a count of finished ChatServiceMock.Send invocations
func (mmSend *ChatServiceMock) SendAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmSend.afterSendCounter)
}

// SendBeforeCounter returns a count of ChatServiceMock.Send invocations
func (mmSend *ChatServiceMock) SendBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmSend.beforeSendCounter)
}

// Calls returns a list of arguments used in each call to ChatServiceMock.Send.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmSend *mChatServiceMockSend) Calls() []*ChatServiceMockSendParams {
	mmSend.mutex.RLock()

	argCopy := make([]*ChatServiceMockSendParams, len(mmSend.callArgs))
	copy(argCopy, mmSend.callArgs)

	mmSend.mutex.RUnlock()

	return argCopy
}

// MinimockSendDone returns true if the count of the Send invocations corresponds
// the number of defined expectations
func (m *ChatServiceMock) MinimockSendDone() bool {
	for _, e := range m.SendMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.SendMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterSendCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcSend != nil && mm_atomic.LoadUint64(&m.afterSendCounter) < 1 {
		return false
	}
	return true
}

// MinimockSendInspect logs each unmet expectation
func (m *ChatServiceMock) MinimockSendInspect() {
	for _, e := range m.SendMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to ChatServiceMock.Send with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.SendMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterSendCounter) < 1 {
		if m.SendMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to ChatServiceMock.Send")
		} else {
			m.t.Errorf("Expected call to ChatServiceMock.Send with params: %#v", *m.SendMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcSend != nil && mm_atomic.LoadUint64(&m.afterSendCounter) < 1 {
		m.t.Error("Expected call to ChatServiceMock.Send")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *ChatServiceMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockCreateInspect()

		m.MinimockDeleteInspect()

		m.MinimockSendInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *ChatServiceMock) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *ChatServiceMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockCreateDone() &&
		m.MinimockDeleteDone() &&
		m.MinimockSendDone()
}
