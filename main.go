package main

import (
	"log"
	"net/http"
	"os"
)

func streamHandler(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("video.mp4")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	done := make(chan struct{})
	go func() {
		defer close(done)

		w.Header().Set("Content-Type", "video/mp4")
		w.Header().Set("Content-Disposition", "attachment; filename=video.mp4")
		w.Header().Set("Transfer-Encoding", "chunked")

		log.Println("start streaming")
		buffer := make([]byte, 4096)
		for {
			chunk, err := file.Read(buffer)
			if err != nil {
				break
			}

			w.Write(buffer[:chunk])

			w.(http.Flusher).Flush()
		}
		log.Println("streaming completed")
	}()

	select {
	case <-done:
		return
	case <-r.Context().Done():
		return
	}
}

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET")
		next.ServeHTTP(w, r)
	})
}

func main() {
	log.Println("start server")
	http.Handle("/stream", enableCORS(http.HandlerFunc(streamHandler)))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
