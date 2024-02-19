package handler

import (
	"encoding/json"
	"net/http"

	"github.com/MatheusNP/fc-ports-adapters/adapters/dto"
	"github.com/MatheusNP/fc-ports-adapters/application"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func MakeProductandlers(
	r *mux.Router,
	n *negroni.Negroni,
	service application.ProductServiceInterface,
) {

	r.Handle("/product/{id}", n.With(
		negroni.Wrap(getProduct(service)),
	)).Methods("GET", "OPTIONS")

	r.Handle("/product", n.With(
		negroni.Wrap(createProduct(service)),
	)).Methods("POST", "OPTIONS")

	r.Handle("/product/{id}/enable", n.With(
		negroni.Wrap(enableProduct(service, true)),
	)).Methods("POST", "OPTIONS")

	r.Handle("/product/{id}/disable", n.With(
		negroni.Wrap(enableProduct(service, false)),
	)).Methods("POST", "OPTIONS")

}

func getProduct(service application.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(r)
		id := vars["id"]

		product, err := service.Get(id)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		if err := json.NewEncoder(w).Encode(product); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError(err.Error()))
			return
		}
	})
}

func createProduct(service application.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var productDto dto.Product

		if err := json.NewDecoder(r.Body).Decode(&productDto); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError(err.Error()))
			return
		}

		product, err := service.Create(productDto.Name, productDto.Price)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError(err.Error()))
			return
		}

		if err := json.NewEncoder(w).Encode(product); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError(err.Error()))
			return
		}
	})
}

func enableProduct(service application.ProductServiceInterface, enable bool) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		id := vars["id"]

		product, err := service.Get(id)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		var result application.ProductInterface
		if enable {
			result, err = service.Enable(product)
		} else {
			result, err = service.Disable(product)
		}
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError(err.Error()))
			return
		}

		if err := json.NewEncoder(w).Encode(result); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError(err.Error()))
			return
		}
	})
}
