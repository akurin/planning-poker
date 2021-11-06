package sessionauth

import (
	"backend/internal/adapters/httpapi/sessions"
	"net/http"
)

type sessionMiddleware struct {
	next         http.Handler
	sessionStore sessions.SessionStore
}

func NewSessionMiddleware(next http.Handler, sessionStore sessions.SessionStore) *sessionMiddleware {
	return &sessionMiddleware{
		next:         next,
		sessionStore: sessionStore,
	}
}

func (s *sessionMiddleware) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	session, _ := s.sessionStore.Get(req)

	if !session.IsAuthenticated() {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	s.next.ServeHTTP(w, req)
}
