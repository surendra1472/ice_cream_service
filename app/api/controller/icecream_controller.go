package controller

import (
	"github.com/go-pg/pg"
	is "ic-indexer-service/app/model/request"
	"ic-service/app/api/dataaccessor"
	"ic-service/app/api/preprocessor"
	"ic-service/app/config"
	"ic-service/app/config/locales/local_config"
	"ic-service/app/model/request"
	"ic-service/app/processor"
	"ic-service/app/response_handler"
	"log"
	"net/http"
)

func AddIcecream(w http.ResponseWriter, r *http.Request) {

	icecreamRequest := *request.NewIcecreamRequest()

	err := preprocessor.DecodeAndValidateRequestParams(r, &icecreamRequest)

	if err != nil {
		msg := local_config.GetTranslationMessage(r.Context(), "invalid_request_params")
		response_handler.WriteErrorResponseAsJson(w, r, http.StatusBadRequest, msg)
		return
	}

	icecreamService := processor.GetNewIcecreamService(getIcecreamConstructorParams())
	icecreamDetails, insertError := icecreamService.Save(r.Context(), icecreamRequest)

	if insertError != nil || icecreamDetails == nil { //when both are nil, then we have recovered from a panic
		msg := local_config.GetTranslationMessage(r.Context(), "invalid_request_params")
		response_handler.WriteErrorResponseAsJson(w, r, http.StatusBadRequest, msg)
		return
	}

	log.Print(r.Context(), *icecreamDetails)
	response_handler.WriteResponseMapAsJson(w, r, http.StatusOK, map[string]interface{}{"data": icecreamDetails})

}

func UpdateIcecream(w http.ResponseWriter, r *http.Request) {

	icecreamUpdateRequest := request.IcecreamUpdateRequest{}
	icecreamClientRequest := request.IcecreamRequest{}
	var requestMap map[string]interface{}

	_, err := preprocessor.BindAndUnMarshallRequest(r, &requestMap, &icecreamUpdateRequest, &icecreamClientRequest)

	if err != nil {
		msg := local_config.GetTranslationMessage(r.Context(), "invalid_request_params")
		response_handler.WriteErrorResponseAsJson(w, r, http.StatusBadRequest, msg)
		return
	}

	icecreamService := processor.GetNewIcecreamService(getIcecreamConstructorParams())
	_, updateErr := icecreamService.PartialUpdate(r.Context(), icecreamUpdateRequest)

	if updateErr != nil { //when both are nil, then we have recovered from a panic
		msg := local_config.GetTranslationMessage(r.Context(), "invalid_request_params")
		response_handler.WriteErrorResponseAsJson(w, r, http.StatusBadRequest, msg)
		return
	}

	response_handler.WriteResponseMapAsJson(w, r, http.StatusOK, map[string]interface{}{"data": "updated successfully"})

}

func DeleteIcecream(w http.ResponseWriter, r *http.Request) {

	icecreamRequest := is.IcecreamDelete{}

	err := preprocessor.DecodeAndValidateRequestParams(r, &icecreamRequest)

	if err != nil {
		msg := local_config.GetTranslationMessage(r.Context(), "invalid_request_params")
		response_handler.WriteErrorResponseAsJson(w, r, http.StatusBadRequest, msg)
		return
	}

	icecreamService := processor.GetNewIcecreamService(getIcecreamConstructorParams())
	_, deleteErr := icecreamService.Delete(r.Context(), icecreamRequest)

	if deleteErr != nil { //when both are nil, then we have recovered from a panic
		msg := local_config.GetTranslationMessage(r.Context(), "invalid_request_params")
		response_handler.WriteErrorResponseAsJson(w, r, http.StatusBadRequest, msg)
		return
	}

	response_handler.WriteResponseMapAsJson(w, r, http.StatusOK, map[string]interface{}{"data": "deleted successfully"})

}

func getIcecreamConstructorParams() (*pg.DB, dataaccessor.IcecreamDataAccessor) {
	return config.GetDBConnection(), dataaccessor.NewIcecreamDataAccessor()
}
