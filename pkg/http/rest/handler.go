package rest

import (
	"encoding/json"
	"log"
	"net/http"

	//	"net/http"

	//	"gopkg.in/mgo.v2/bson"

	//	"ResourcePool/configs"
	//	"ResourcePool/pkg/dao"
	//	"ResourcePool/pkg/models"

	"github.com/gorilla/mux"
	minio "github.com/minio/minio-go"
	//minio "github.com/minio/minio-go"
)

var minioClient *minio.Client
var location = "us-east-1"

func Init() {
	endpoint := "play.minio.io:9000"
	accessKeyID := "Q3AM3UQ867SPQQA43P2F"
	secretAccessKey := "zuf+tfteSlswRu7BJ86wekitnifILbZam1KYY3TG"
	useSSL := true

	var err error
	minioClient, err = minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
	if err != nil {
		log.Fatalln(err)
		return
	}

	log.Printf("%#v\n", minioClient)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJson(w, code, map[string]string{"error": message})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func Handler() http.Handler {
	r := mux.NewRouter()

	// bucket api
	r.HandleFunc("/buckets", allBuckets).Methods("GET")             // get bucket list
	r.HandleFunc("/bucket", createBucket).Methods("POST")           // create bucket
	r.HandleFunc("/buckets/{name}", updateBucket).Methods("PUT")    // update bucket
	r.HandleFunc("/buckets/{name}", deleteBucket).Methods("DELETE") // delete bucket
	r.HandleFunc("/buckets/{name}", listBucket).Methods("GET")      // list bucket items

	// object api
	r.HandleFunc("/bucket/{bucketname}", allObjects).Methods("GET")
	r.HandleFunc("/bucket/{bucketname}/object", createObject).Methods("POST")
	r.HandleFunc("/bucket/{bucketname}/objects/{objectname}", updateObject).Methods("PUT")
	r.HandleFunc("/bucket/{bucketname}/objects/{objectname}", deleteObject).Methods("DELETE")
	r.HandleFunc("/bucket/{bucketname}/objects/{objectname}", findObject).Methods("GET")

	return r
}

func allBuckets(w http.ResponseWriter, r *http.Request) {
	buckets, err := minioClient.ListBuckets()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJson(w, http.StatusOK, buckets)
}

func createBucket(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bucketName := params["bucket"]
	if err := minioClient.MakeBucket(bucketName, location); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	respondWithJson(w, http.StatusCreated, nil)
}

func updateBucket(w http.ResponseWriter, r *http.Request) {
	//	defer r.Body.Close()
	//	var image models.Image
	//	if err := json.NewDecoder(r.Body).Decode(&image); err != nil {
	//		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
	//		return
	//	}
	//	if err := da.Update(image); err != nil {
	//		respondWithError(w, http.StatusInternalServerError, err.Error())
	//		return
	//	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

func deleteBucket(w http.ResponseWriter, r *http.Request) {
	//	defer r.Body.Close()
	//	var image models.Image
	//	if err := json.NewDecoder(r.Body).Decode(&image); err != nil {
	//		respondWithError(w, http.StatusBadRequest, err.Error())
	//		return
	//	}
	//	if err := da.Delete(image); err != nil {
	//		respondWithError(w, http.StatusInternalServerError, err.Error())
	//		return
	//	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

func listBucket(w http.ResponseWriter, r *http.Request) {
	//	params := mux.Vars(r)
	//	image, err := da.FindById(params["id"])
	//	if err != nil {
	//		respondWithError(w, http.StatusBadRequest, "Invalid image ID")
	//		return
	//	}
	respondWithJson(w, http.StatusOK, nil)
}

func createObject(w http.ResponseWriter, r *http.Request) {
	//	defer r.Body.Close()
	//	var image models.Image
	//	if err := json.NewDecoder(r.Body).Decode(&image); err != nil {
	//		respondWithError(w, http.StatusBadRequest, err.Error())
	//		return
	//	}

	//	image.ID = bson.NewObjectId()

	//	if err := da.Insert(image); err != nil {
	//		respondWithError(w, http.StatusInternalServerError, err.Error())
	//		return
	//	}

	respondWithJson(w, http.StatusCreated, nil)
}

func updateObject(w http.ResponseWriter, r *http.Request) {
	//	defer r.Body.Close()
	//	var image models.Image
	//	if err := json.NewDecoder(r.Body).Decode(&image); err != nil {
	//		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
	//		return
	//	}
	//	if err := da.Update(image); err != nil {
	//		respondWithError(w, http.StatusInternalServerError, err.Error())
	//		return
	//	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

func deleteObject(w http.ResponseWriter, r *http.Request) {
	//	defer r.Body.Close()
	//	var image models.Image
	//	if err := json.NewDecoder(r.Body).Decode(&image); err != nil {
	//		respondWithError(w, http.StatusBadRequest, err.Error())
	//		return
	//	}
	//	if err := da.Delete(image); err != nil {
	//		respondWithError(w, http.StatusInternalServerError, err.Error())
	//		return
	//	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

func listObject(w http.ResponseWriter, r *http.Request) {
	//	params := mux.Vars(r)
	//	image, err := da.FindById(params["id"])
	//	if err != nil {
	//		respondWithError(w, http.StatusBadRequest, "Invalid image ID")
	///		return
	//	}
	respondWithJson(w, http.StatusOK, nil)
}

func findObject(w http.ResponseWriter, r *http.Request) {
	//	params := mux.Vars(r)
	//	image, err := da.FindById(params["id"])
	//	if err != nil {
	//		respondWithError(w, http.StatusBadRequest, "Invalid image ID")
	///		return
	//	}
	respondWithJson(w, http.StatusOK, nil)
}

func allObjects(w http.ResponseWriter, r *http.Request) {
	//	params := mux.Vars(r)
	//	image, err := da.FindById(params["id"])
	//	if err != nil {
	//		respondWithError(w, http.StatusBadRequest, "Invalid image ID")
	///		return
	//	}
	respondWithJson(w, http.StatusOK, nil)
}
