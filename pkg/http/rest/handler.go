package rest

import (
	"encoding/json"
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"ResourcePool/configs"
	"ResourcePool/pkg/dao"
	"ResourcePool/pkg/models"

	"github.com/gorilla/mux"
)

var cfg = config.Config{}
var da = dao.ImagesDAO{}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJson(w, code, map[string]string{"error": message})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func allImagesEndPoint(w http.ResponseWriter, r *http.Request) {
	images, err := da.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJson(w, http.StatusOK, images)
}

func findImagesEndPoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	image, err := da.FindById(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid image ID")
		return
	}
	respondWithJson(w, http.StatusOK, image)
}

func createImagesEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var image models.Image
	if err := json.NewDecoder(r.Body).Decode(&image); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	image.ID = bson.NewObjectId()

	if err := da.Insert(image); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJson(w, http.StatusCreated, image)
}

func updateImagesEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var image models.Image
	if err := json.NewDecoder(r.Body).Decode(&image); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := da.Update(image); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

func deleteImagesEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var image models.Image
	if err := json.NewDecoder(r.Body).Decode(&image); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	if err := da.Delete(image); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

func Init() {
	cfg.Read()

	da.Server = cfg.Server
	da.Database = cfg.Database
	da.Connect()
}

func Handler() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/images", allImagesEndPoint).Methods("GET")
	r.HandleFunc("/images", createImagesEndPoint).Methods("POST")
	r.HandleFunc("/images", updateImagesEndPoint).Methods("PUT")
	r.HandleFunc("/images", deleteImagesEndPoint).Methods("DELETE")
	r.HandleFunc("/images/{id}", findImagesEndPoint).Methods("GET")

	return r
}
