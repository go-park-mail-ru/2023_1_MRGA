package delivery

import (
	"net/http"

	"github.com/go-park-mail-ru/2023_1_MRGA.git/utils/logger"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/utils/writer"
)

func (h *Handler) GetHashtags(w http.ResponseWriter, r *http.Request) {
	jobs, err := h.useCase.GetHashtags()
	if err != nil {
		logger.Log(http.StatusInternalServerError, err.Error(), r.Method, r.URL.Path)
		writer.ErrorRespond(w, r, err, http.StatusInternalServerError)
		return
	}

	result := make(map[string]interface{})
	result["hashtags"] = jobs
	logger.Log(http.StatusOK, "Success", r.Method, r.URL.Path)
	writer.Respond(w, r, result)
}

func (h *Handler) GetCities(w http.ResponseWriter, r *http.Request) {
	cities, err := h.useCase.GetCities()
	if err != nil {
		logger.Log(http.StatusInternalServerError, err.Error(), r.Method, r.URL.Path)
		writer.ErrorRespond(w, r, err, http.StatusInternalServerError)
		return
	}

	result := make(map[string]interface{})
	result["cities"] = cities
	logger.Log(http.StatusOK, "Success", r.Method, r.URL.Path)
	writer.Respond(w, r, result)
}

func (h *Handler) GetZodiac(w http.ResponseWriter, r *http.Request) {
	zodiac, err := h.useCase.GetZodiacs()
	if err != nil {
		logger.Log(http.StatusInternalServerError, err.Error(), r.Method, r.URL.Path)
		writer.ErrorRespond(w, r, err, http.StatusInternalServerError)
		return
	}

	result := make(map[string]interface{})
	result["zodiac"] = zodiac
	logger.Log(http.StatusOK, "Success", r.Method, r.URL.Path)
	writer.Respond(w, r, result)
}

func (h *Handler) GetJobs(w http.ResponseWriter, r *http.Request) {
	jobs, err := h.useCase.GetJobs()
	if err != nil {
		logger.Log(http.StatusInternalServerError, err.Error(), r.Method, r.URL.Path)
		writer.ErrorRespond(w, r, err, http.StatusInternalServerError)
		return
	}

	result := make(map[string]interface{})
	result["jobs"] = jobs
	logger.Log(http.StatusOK, "Success", r.Method, r.URL.Path)
	writer.Respond(w, r, result)
}

func (h *Handler) GetEducation(w http.ResponseWriter, r *http.Request) {
	education, err := h.useCase.GetEducation()
	if err != nil {
		logger.Log(http.StatusInternalServerError, err.Error(), r.Method, r.URL.Path)
		writer.ErrorRespond(w, r, err, http.StatusInternalServerError)
		return
	}

	result := make(map[string]interface{})
	result["education"] = education
	logger.Log(http.StatusOK, "Success", r.Method, r.URL.Path)
	writer.Respond(w, r, result)
}

func (h *Handler) GetReasons(w http.ResponseWriter, r *http.Request) {
	reasons, err := h.useCase.GetReasons()
	if err != nil {
		logger.Log(http.StatusInternalServerError, err.Error(), r.Method, r.URL.Path)
		writer.ErrorRespond(w, r, err, http.StatusInternalServerError)
		return
	}

	result := make(map[string]interface{})
	result["reasons"] = reasons
	logger.Log(http.StatusOK, "Success", r.Method, r.URL.Path)
	writer.Respond(w, r, result)
}
