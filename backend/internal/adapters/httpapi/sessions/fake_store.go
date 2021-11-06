package sessions

import "github.com/quasoft/memstore"

func NewFakeStore() SessionStore {
	store := memstore.NewMemStore(
		[]byte("authkey123"),
		[]byte("enckey12341234567890123456789012"),
	)
	return NewGorillaSessionStore(store, "session_id")
}
