package delivery

import (
	"net/http"

	"github.com/go-park-mail-ru/2023_1_MRGA.git/utils/logger"
	tracejaeger "github.com/go-park-mail-ru/2023_1_MRGA.git/utils/trace_jaeger"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/utils/writer"
)

func (h *Handler) GetHashtags(w http.ResponseWriter, r *http.Request) {
	_, parentSpan := tracejaeger.NewSpan(r.Context(), "mainServer", "GetHashtagsHandler", nil)
	defer parentSpan.End()

	hashtags, err := h.useCase.GetHashtags()
	if err != nil {
		logger.Log(http.StatusInternalServerError, err.Error(), r.Method, r.URL.Path, true)
		writer.ErrorRespond(w, r, err, http.StatusInternalServerError)
		return
	}

	result := make(map[string]interface{})
	result["hashtags"] = hashtags
	logger.Log(http.StatusOK, "Success", r.Method, r.URL.Path, false)
	writer.Respond(w, r, result)
}

func (h *Handler) GetCities(w http.ResponseWriter, r *http.Request) {
	_, parentSpan := tracejaeger.NewSpan(r.Context(), "mainServer", "GetCitiesHandler", nil)
	defer parentSpan.End()

	cities, err := h.useCase.GetCities()
	if err != nil {
		logger.Log(http.StatusInternalServerError, err.Error(), r.Method, r.URL.Path, true)
		writer.ErrorRespond(w, r, err, http.StatusInternalServerError)
		return
	}

	result := make(map[string]interface{})
	result["cities"] = cities
	logger.Log(http.StatusOK, "Success", r.Method, r.URL.Path, false)
	writer.Respond(w, r, result)
}

func (h *Handler) GetZodiac(w http.ResponseWriter, r *http.Request) {
	_, parentSpan := tracejaeger.NewSpan(r.Context(), "mainServer", "GetZodiacHandler", nil)
	defer parentSpan.End()

	zodiac, err := h.useCase.GetZodiacs()
	if err != nil {
		logger.Log(http.StatusInternalServerError, err.Error(), r.Method, r.URL.Path, true)
		writer.ErrorRespond(w, r, err, http.StatusInternalServerError)
		return
	}

	result := make(map[string]interface{})
	result["zodiac"] = zodiac
	logger.Log(http.StatusOK, "Success", r.Method, r.URL.Path, false)
	writer.Respond(w, r, result)
}

func (h *Handler) GetJobs(w http.ResponseWriter, r *http.Request) {
	_, parentSpan := tracejaeger.NewSpan(r.Context(), "mainServer", "GetJobsHandler", nil)
	defer parentSpan.End()

	jobs, err := h.useCase.GetJobs()
	if err != nil {
		logger.Log(http.StatusInternalServerError, err.Error(), r.Method, r.URL.Path, true)
		writer.ErrorRespond(w, r, err, http.StatusInternalServerError)
		return
	}

	result := make(map[string]interface{})
	result["jobs"] = jobs
	logger.Log(http.StatusOK, "Success", r.Method, r.URL.Path, false)
	writer.Respond(w, r, result)
}

func (h *Handler) GetEducation(w http.ResponseWriter, r *http.Request) {
	_, parentSpan := tracejaeger.NewSpan(r.Context(), "mainServer", "GetEducationHandler", nil)
	defer parentSpan.End()

	education, err := h.useCase.GetEducation()
	if err != nil {
		logger.Log(http.StatusInternalServerError, err.Error(), r.Method, r.URL.Path, true)
		writer.ErrorRespond(w, r, err, http.StatusInternalServerError)
		return
	}

	result := make(map[string]interface{})
	result["education"] = education
	logger.Log(http.StatusOK, "Success", r.Method, r.URL.Path, false)
	writer.Respond(w, r, result)
}

func (h *Handler) GetReasons(w http.ResponseWriter, r *http.Request) {
	_, parentSpan := tracejaeger.NewSpan(r.Context(), "mainServer", "GetReasonsHandler", nil)
	defer parentSpan.End()

	reasons, err := h.useCase.GetReasons()
	if err != nil {
		logger.Log(http.StatusInternalServerError, err.Error(), r.Method, r.URL.Path, true)
		writer.ErrorRespond(w, r, err, http.StatusInternalServerError)
		return
	}

	result := make(map[string]interface{})
	result["reasons"] = reasons
	logger.Log(http.StatusOK, "Success", r.Method, r.URL.Path, false)
	writer.Respond(w, r, result)
}

func (h *Handler) GetStatuses(w http.ResponseWriter, r *http.Request) {
	_, parentSpan := tracejaeger.NewSpan(r.Context(), "mainServer", "GetStatusesHandler", nil)
	defer parentSpan.End()

	statuses, err := h.useCase.GetStatuses()
	if err != nil {
		logger.Log(http.StatusInternalServerError, err.Error(), r.Method, r.URL.Path, true)
		writer.ErrorRespond(w, r, err, http.StatusInternalServerError)
		return
	}

	result := make(map[string]interface{})
	result["statuses"] = statuses
	logger.Log(http.StatusOK, "Success", r.Method, r.URL.Path, false)
	writer.Respond(w, r, result)
}
