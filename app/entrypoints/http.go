package entrypoints

import (
	"challange-20211024/app/entities"
	"challange-20211024/app/usecase"
	"encoding/json"
	"net/http"
)

func GetHandlers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		entity, err := entities.NewGet(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		respon, err := usecase.NewDataGetter().GetData(entity)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		resultRespon, _ := json.Marshal(respon)
		w.Write(resultRespon)
	default:
		http.Error(w, "", http.StatusBadRequest)
	}
}
