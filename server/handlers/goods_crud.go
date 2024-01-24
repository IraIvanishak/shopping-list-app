package handlers

import "net/http"

func AddGoodToListHandler(w http.ResponseWriter, r *http.Request) {

}

func GetAllGoodsHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello?"))
}

func UpdateGoodHandler(w http.ResponseWriter, r *http.Request) {

}

func DeleteGoodHandler(w http.ResponseWriter, r *http.Request) {

}
