package rest

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

type StatusRecorder struct {
	http.ResponseWriter
	statusCode int
}

func (rw *StatusRecorder) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// wr := responseWriter(w)
		rec := &StatusRecorder{ResponseWriter: w, statusCode: 200}

		next.ServeHTTP(rec, r) // Call the next handler
		//statusCode := wr.statusCode
		log.WithFields(log.Fields{
			"method":     r.Method,
			"uri":        r.RequestURI,
			"remoteAddr": r.RemoteAddr,
			"code":       rec.statusCode,
		}).Info()
	})

}
