package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/IraIvanishak/shopping-list-app/storage"
	"github.com/go-chi/chi"
)

func AddGoodToListHandler(w http.ResponseWriter, r *http.Request) {
	var newGood storage.Good
	err := json.NewDecoder(r.Body).Decode(&newGood)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = storage.CreateGood(&newGood)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func GetAllGoodsHandler(w http.ResponseWriter, r *http.Request) {
	allGoods, err := storage.GetAllGoods()
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	allGoodsJSON, err := json.Marshal(allGoods)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(allGoodsJSON)
}

func UpdateGoodHandler(w http.ResponseWriter, r *http.Request) {
	IDString := chi.URLParam(r, "id")
	ID, err := strconv.Atoi(IDString)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var updatedGood storage.Good
	err = json.NewDecoder(r.Body).Decode(&updatedGood)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if ID != updatedGood.ID {
		log.Printf("IDs does not match")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = storage.UpdateGood(&updatedGood)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func DeleteGoodHandler(w http.ResponseWriter, r *http.Request) {
	IDString := chi.URLParam(r, "id")
	ID, err := strconv.Atoi(IDString)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = storage.DeleteGood(ID)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
