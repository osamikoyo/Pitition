package handler

import (
	"encoding/json"
	"github.com/osamikoyo/pitition/internal/data/models"
	"io/ioutil"
	"net/http"
)

func (h Handler) addPitHandler(w http.ResponseWriter, r *http.Request) error {
	var pit models.Pitition

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(body, &pit); err != nil {
		return err
	}

	return h.Pit.Add(pit)
}
