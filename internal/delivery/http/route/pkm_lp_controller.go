package route

import (
	"encoding/json"
	"net/http"
	"strconv"
	"tugasakhir/internal/model"
	"tugasakhir/internal/usecase"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type PKMLPController struct {
	UseCase *usecase.PKMLPUseCase
	Log     *logrus.Logger
}

func NewPKMLPController(useCase *usecase.PKMLPUseCase, log *logrus.Logger) *PKMLPController {
	return &PKMLPController{
		UseCase: useCase,
		Log:     log,
	}
}

func (c *PKMLPController) Create(w http.ResponseWriter, r *http.Request) {

	request := new(model.CreatePKMLPRequest)

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		c.Log.Warnf("Failed to parse request body: %+v", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	response, err := c.UseCase.Create(r.Context(), request)
	if err != nil {
		c.Log.WithError(err).Error("error creating contact")
		http.Error(w, "Internal Server", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(model.WebResponse[*model.PKMLPResponse]{Data: response}); err != nil {
		c.Log.Warnf("Failed to write response: %+v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (c *PKMLPController) List(w http.ResponseWriter, r *http.Request) {

	responses, err := c.UseCase.FindAll(r.Context())
	if err != nil {
		c.Log.WithError(err).Error("error get all pkm LP")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// paging := &model.PageMetadata{
	// 	Page:      1,
	// 	Size:      1,
	// 	TotalItem: 0,
	// 	TotalPage: int64(math.Ceil(float64(total) / float64(request.Size))),
	// }

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(model.WebResponse[[]model.PKMLPResponse]{Data: responses}); err != nil {
		c.Log.Warnf("Failed to write response: %+v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (c *PKMLPController) Update(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]
	idUint, err := strconv.Atoi(id)
	if err != nil {
		c.Log.Warnf("Failed to parse id: %+v", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	request := new(model.UpdatePKMLPRequest)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		c.Log.Warnf("Failed to parse request body: %+v", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	request.ID = uint(idUint)

	response, err := c.UseCase.Update(r.Context(), request)
	if err != nil {
		c.Log.WithError(err).Error("error updating contact")
		http.Error(w, "Internal server", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(model.WebResponse[*model.PKMLPResponse]{Data: response}); err != nil {
		c.Log.Warnf("Failed to write response: %+v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (c *PKMLPController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	idUint, err := strconv.Atoi(id)
	if err != nil {
		c.Log.Warnf("Failed to parse id: %+v", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	request := &model.DeletePKMLPRequest{
		ID: uint(idUint),
	}

	if err := c.UseCase.Delete(r.Context(), request); err != nil {
		c.Log.WithError(err).Error("error deleting pkm LP")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(model.WebResponse[bool]{Data: true}); err != nil {
		c.Log.Warnf("Failed to write response: %+v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
