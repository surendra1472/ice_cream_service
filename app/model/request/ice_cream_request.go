package request

import (
	"ic-service/app/utils"
	"time"
)

// Add icecream Request
// swagger:model icecreamRequest
type IcecreamRequest struct {
	// a icecream request
	// in: body
	Id                    int64    `json:"id"`
	Name                  string   `json:"name" validate:"required" conform:"trim,omitempty"`
	ImageClosed           string   `json:"image_closed" conform:"trim,omitempty"`
	ImageOpen             string   `json:"image_open" conform:"trim,omitempty"`
	Description           string   `json:"description" conform:"trim,omitempty"`
	Story                 string   `json:"story" conform:"trim,omitempty"`
	AllergyInfo           string   `json:"allergy_info" conform:"trim,omitempty"`
	SourcingValues        []string `json:"sourcing_values" conform:"trim,omitempty"`
	Ingredients           []string `json:"ingredients" conform:"trim,omitempty"`
	DietaryCertifications string   `json:"dietary_certifications" conform:"trim,omitempty"`
	ProductId             string   `json:"product_id" validate:"required" conform:"trim,omitempty"`
}

func NewIcecreamRequest() *IcecreamRequest {
	return &IcecreamRequest{}
}

type IcecreamUpdateRequest struct {
	Id                    CusInt64       `json:"id" validate:"gt=0"`
	Name                  CusString      `json:"name,omitempty"`
	ImageClosed           CusString      `json:"image_closed,omitempty"`
	ImageOpen             CusString      `json:"image_open,omitempty"`
	Description           CusString      `json:"description,omitempty"`
	Story                 CusString      `json:"story,omitempty"`
	AllergyInfo           CusString      `json:"allergy_info,omitempty"`
	SourcingValues        CusArrayString `json:"sourcing_values,omitempty"`
	Ingredients           CusArrayString `json:"ingredients,omitempty"`
	DietaryCertifications CusString      `json:"dietary_certifications,omitempty"`
	ProductId             CusString      `json:"product_id,omitempty"`
}

// swagger:model genericError
type GenericError struct {
	Code      int    `json:"code"`
	Error     string `json:"error"`
	RequestId string `json:"request_id"`
}

// swagger:model genericModel
type GenericResponse struct {
	Data      map[string]interface{} `json:"data"`
	Code      int                    `json:"code"`
	RequestId string                 `json:"request_id"`
	Error     string                 `json:"error"`
}

type CusTime struct {
	Value *time.Time
	Set   bool
}

type CusArrayString struct {
	Value *[]string
	Set   bool
}

type CusInt64 struct {
	Value *int64
	Set   bool
}

type CusString struct {
	Value *string
	Set   bool
}

func NewCusInt64(value *int64) *CusInt64 {
	return &CusInt64{Value: value, Set: true}
}

func (x *CusString) UnmarshalJSON(data []byte) error {
	return utils.UnmarshalJSON(data, new(string), &x.Value, &x.Set)
}

func (x *CusInt64) UnmarshalJSON(data []byte) error {
	return utils.UnmarshalJSON(data, new(int64), &x.Value, &x.Set)
}
func (x *CusTime) UnmarshalJSON(data []byte) error {
	return utils.UnmarshalJSON(data, new(time.Time), &x.Value, &x.Set)
}

func (x *CusArrayString) UnmarshalJSON(data []byte) error {
	return utils.UnmarshalJSON(data, new([]string), &x.Value, &x.Set)
}
