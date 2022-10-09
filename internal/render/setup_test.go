package render

import (
	"encoding/gob"
	"github.com/alexedwards/scs/v2"
	"github.com/sarkartanmay393/RoomReservation-WebApp/internal/config"
	"github.com/sarkartanmay393/RoomReservation-WebApp/internal/models"
	"net/http"
	"os"
	"testing"
	"time"
)

var appConfTest config.AppConfig
var sessionTest *scs.SessionManager

func TestMain(m *testing.M) {
	gob.Register(&models.Reservation{})
	gob.Register(&models.ChosenDates{})
	// Session Management Implementation
	sessionTest = scs.New()
	sessionTest.Lifetime = 24 * time.Hour
	sessionTest.Cookie.Persist = true
	sessionTest.Cookie.SameSite = http.SameSiteLaxMode
	sessionTest.Cookie.Secure = false

	appConfTest.SessionManager = sessionTest // Transfers this session object to app config.
	appConf = &appConfTest

	os.Exit(m.Run())
}

type httpResponseWriter struct {
}

func (rw *httpResponseWriter) Header() http.Header {
	return http.Header{}
}
func (rw *httpResponseWriter) Write([]byte) (int, error) {
	return 0, nil
}
func (rw *httpResponseWriter) WriteHeader(statusCode int) {
}
