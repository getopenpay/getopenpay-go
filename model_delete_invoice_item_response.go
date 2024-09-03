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

// checks if the DeleteInvoiceItemResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteInvoiceItemResponse{}

// DeleteInvoiceItemResponse struct for DeleteInvoiceItemResponse
type DeleteInvoiceItemResponse struct {
	// Message describing result of API call.
	Message *string `json:"message,omitempty"`
	// Unique identifier of the invoice_item.
	InvoiceItemId string `json:"invoice_item_id"`
}

type _DeleteInvoiceItemResponse DeleteInvoiceItemResponse

// NewDeleteInvoiceItemResponse instantiates a new DeleteInvoiceItemResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteInvoiceItemResponse(invoiceItemId string) *DeleteInvoiceItemResponse {
	this := DeleteInvoiceItemResponse{}
	var message string = "Invoice Item Deleted successfully."
	this.Message = &message
	this.InvoiceItemId = invoiceItemId
	return &this
}

// NewDeleteInvoiceItemResponseWithDefaults instantiates a new DeleteInvoiceItemResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteInvoiceItemResponseWithDefaults() *DeleteInvoiceItemResponse {
	this := DeleteInvoiceItemResponse{}
	var message string = "Invoice Item Deleted successfully."
	this.Message = &message
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *DeleteInvoiceItemResponse) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteInvoiceItemResponse) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *DeleteInvoiceItemResponse) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *DeleteInvoiceItemResponse) SetMessage(v string) {
	o.Message = &v
}

// GetInvoiceItemId returns the InvoiceItemId field value
func (o *DeleteInvoiceItemResponse) GetInvoiceItemId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InvoiceItemId
}

// GetInvoiceItemIdOk returns a tuple with the InvoiceItemId field value
// and a boolean to check if the value has been set.
func (o *DeleteInvoiceItemResponse) GetInvoiceItemIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InvoiceItemId, true
}

// SetInvoiceItemId sets field value
func (o *DeleteInvoiceItemResponse) SetInvoiceItemId(v string) {
	o.InvoiceItemId = v
}

func (o DeleteInvoiceItemResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteInvoiceItemResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	toSerialize["invoice_item_id"] = o.InvoiceItemId
	return toSerialize, nil
}

func (o *DeleteInvoiceItemResponse) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"invoice_item_id",
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

	varDeleteInvoiceItemResponse := _DeleteInvoiceItemResponse{}

	err = json.Unmarshal(bytes, &varDeleteInvoiceItemResponse)

	if err != nil {
		return err
	}

	*o = DeleteInvoiceItemResponse(varDeleteInvoiceItemResponse)

	return err
}

type NullableDeleteInvoiceItemResponse struct {
	value *DeleteInvoiceItemResponse
	isSet bool
}

func (v NullableDeleteInvoiceItemResponse) Get() *DeleteInvoiceItemResponse {
	return v.value
}

func (v *NullableDeleteInvoiceItemResponse) Set(val *DeleteInvoiceItemResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteInvoiceItemResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteInvoiceItemResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteInvoiceItemResponse(val *DeleteInvoiceItemResponse) *NullableDeleteInvoiceItemResponse {
	return &NullableDeleteInvoiceItemResponse{value: val, isSet: true}
}

func (v NullableDeleteInvoiceItemResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteInvoiceItemResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


