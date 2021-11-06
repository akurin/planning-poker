package sessions

import (
	"net/http"
)

type SessionStore interface {
	Get(req *http.Request) (Session, error)
	Save(w http.ResponseWriter, req *http.Request, session Session) error
}
