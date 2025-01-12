package handler

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
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

func (h Handler) getPitHandler(w http.ResponseWriter, r *http.Request) error {
	name := chi.URLParam(r, "name")
	var (
		err   error
		pits  []models.Pitition
		bytes []byte
	)
	pits, err = h.Pit.Get(name)

	bytes, err = json.Marshal(&pits)

	_, err = fmt.Fprint(w, bytes)

	return err
}

func (h Handler) addLikeHandler(w http.ResponseWriter, r *http.Request) error {
	return h.Pit.Like(chi.URLParam(r, "name"))
}
