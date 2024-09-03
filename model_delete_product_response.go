/*
OpenPay API

super charge your subscription management.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package getopenpay

import (
	"encoding/json"
	"fmt"
)

// checks if the DeleteProductResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteProductResponse{}

// DeleteProductResponse struct for DeleteProductResponse
type DeleteProductResponse struct {
	// Message describing result of API call.
	Message *string `json:"message,omitempty"`
	// Unique identifier of the product.
	ProductId string `json:"product_id"`
}

type _DeleteProductResponse DeleteProductResponse

// NewDeleteProductResponse instantiates a new DeleteProductResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteProductResponse(productId string) *DeleteProductResponse {
	this := DeleteProductResponse{}
	var message string = "Product deleted Successfully."
	this.Message = &message
	this.ProductId = productId
	return &this
}

// NewDeleteProductResponseWithDefaults instantiates a new DeleteProductResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteProductResponseWithDefaults() *DeleteProductResponse {
	this := DeleteProductResponse{}
	var message string = "Product deleted Successfully."
	this.Message = &message
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *DeleteProductResponse) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteProductResponse) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *DeleteProductResponse) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *DeleteProductResponse) SetMessage(v string) {
	o.Message = &v
}

// GetProductId returns the ProductId field value
func (o *DeleteProductResponse) GetProductId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProductId
}

// GetProductIdOk returns a tuple with the ProductId field value
// and a boolean to check if the value has been set.
func (o *DeleteProductResponse) GetProductIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductId, true
}

// SetProductId sets field value
func (o *DeleteProductResponse) SetProductId(v string) {
	o.ProductId = v
}

func (o DeleteProductResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteProductResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	toSerialize["product_id"] = o.ProductId
	return toSerialize, nil
}

func (o *DeleteProductResponse) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"product_id",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varDeleteProductResponse := _DeleteProductResponse{}

	err = json.Unmarshal(bytes, &varDeleteProductResponse)

	if err != nil {
		return err
	}

	*o = DeleteProductResponse(varDeleteProductResponse)

	return err
}

type NullableDeleteProductResponse struct {
	value *DeleteProductResponse
	isSet bool
}

func (v NullableDeleteProductResponse) Get() *DeleteProductResponse {
	return v.value
}

func (v *NullableDeleteProductResponse) Set(val *DeleteProductResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteProductResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteProductResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteProductResponse(val *DeleteProductResponse) *NullableDeleteProductResponse {
	return &NullableDeleteProductResponse{value: val, isSet: true}
}

func (v NullableDeleteProductResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteProductResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

