package sessions

import (
	"github.com/gorilla/sessions"
	"net/http"
)

type gorillaSessionStore struct {
	store       sessions.Store
	sessionName string
}

func NewGorillaSessionStore(store sessions.Store, sessionName string) SessionStore {
	return &gorillaSessionStore{
		store:       store,
		sessionName: sessionName,
	}
}

func (g *gorillaSessionStore) Get(req *http.Request) (Session, error) {
	session, err := g.store.Get(req, g.sessionName)
	if err != nil {
		return nil, err
	}
	return NewGorillaSession(session), nil
}

func (g *gorillaSessionStore) Save(w http.ResponseWriter, req *http.Request, session Session) error {
	value, ok := session.(*gorillaSession)
	if !ok {
		panic("session type is not gorillaSession")
	}
	return g.store.Save(req, w, value.session)
}
