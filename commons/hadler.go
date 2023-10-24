package commons

import "net/http"

func SendReponse(write http.ResponseWriter, status int, data []byte) {
	write.Header().Set("Conten-Type", "application/json")
	write.WriteHeader(status)
	write.Write(data)
}

func SendError(write http.ResponseWriter, status int) {
	data := []byte(`{}`)
	write.Header().Set("Conten-Type", "application/json")
	write.WriteHeader(status)
	write.Write(data)
}
