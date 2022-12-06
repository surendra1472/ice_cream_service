package response_handler

import (
	"context"
	"encoding/json"
	"ic-service/app/config/locales/local_config"
	"ic-service/icecream_middleware"
	"log"
	"net/http"
)

func WriteResponseMapAsJson(w http.ResponseWriter, r *http.Request, statusCode int, data map[string]interface{}) {
	log.Print(r.Context(), "Writing response for the request")
	update(data, statusCode, icecream_middleware.GetRequestIDFromRequest(r))
	js, err := json.Marshal(data)
	if err != nil {
		statusCode = http.StatusInternalServerError
		message := local_config.GetTranslationMessage(r.Context(), "something_went_wrong")
		data = getErrorData(message)
		update(data, statusCode, icecream_middleware.GetRequestIDFromRequest(r))
		js, _ = json.Marshal(data)
	}
	addCustomHeaders(w)
	w.WriteHeader(statusCode)
	w.Write(js)
	setStatusInContext(r, statusCode)
}

func WriteErrorResponseAsJson(w http.ResponseWriter, r *http.Request, statusCode int, errorMessage string) {
	data := getErrorData(errorMessage)
	WriteResponseMapAsJson(w, r, statusCode, data)
}

func setStatusInContext(r *http.Request, statusCode int) {
	log.Print(r.Context(), "Setting status code for the request")
	ctx := r.Context()
	ctx = context.WithValue(ctx, "status", statusCode)
	r.WithContext(ctx)
}

func getErrorData(errorMessage string) map[string]interface{} {
	data := make(map[string]interface{})
	data["error"] = errorMessage
	return data
}

func addCustomHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "max-age=0, private, must-revalidate")
}

func update(data map[string]interface{}, statusCode int, requestId string) {
	data["code"] = statusCode
	data["request_id"] = requestId
}
