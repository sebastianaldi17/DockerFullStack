package hello

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/sebastianaldi17/dockerfullstack/server/code/entity"
)

func (h *Handler) Hello(w http.ResponseWriter, r *http.Request) {
	err := h.helloBusiness.Hello()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("A problem occured, please try again later."))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello there!"))
}

func (h *Handler) HelloCron() {
	log.Println("Hello Cron @ ", time.Now())
}

func (h *Handler) GetLogs(w http.ResponseWriter, r *http.Request) {
	var req entity.HelloLogsRequest
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("An error occured while trying to read body: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("A problem occured, please try again later."))
		return
	}

	json.Unmarshal(reqBody, &req)
	if err != nil {
		log.Println("An error occured while trying to decode body: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("A problem occured, please try again later."))
		return
	}

	log.Printf("Request body: %+v\n", req)
	logs, err := h.helloBusiness.GetLogs(req)
	if err != nil {
		log.Println("An error occured while trying to get logs: ", err)
	}

	w.WriteHeader(logs.Metadata.StatusCode)
	json.NewEncoder(w).Encode(logs)
}
