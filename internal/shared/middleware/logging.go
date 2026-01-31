package middleware

import (
    "log"
    "net/http"
    "time"
)

func Logging(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        
        // Log request masuk
        log.Printf("%s %s %s", r.Method, r.URL.Path, r.RemoteAddr)
        
        // Call handler asli
        next.ServeHTTP(w, r)
        
        // Log durasi
        log.Printf("%s %s %s - %v", 
            r.Method, r.URL.Path, r.RemoteAddr, 
            time.Since(start))
    })
}