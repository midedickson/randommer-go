package randommer

import "log"

var req *requester

func Init(apiKey string) {
	log.Println("Initializing randommer.io...")
	req = newRequester("https://randommer.io/api", apiKey)
}
