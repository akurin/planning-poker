package fakehandler

import "net/http"

type fakeHandler struct {
	wasCalled bool
}

func New() *fakeHandler {
	return &fakeHandler{}
}

func (f *fakeHandler) ServeHTTP(_ http.ResponseWriter, _ *http.Request) {
	f.wasCalled = true
}

func (f *fakeHandler) WasCalled() bool {
	return f.wasCalled
}
