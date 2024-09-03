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

// checks if the CreatePortalSessionRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreatePortalSessionRequest{}

// CreatePortalSessionRequest struct for CreatePortalSessionRequest
type CreatePortalSessionRequest struct {
	// The ID of an existing customer.
	CustomerId string `json:"customer_id"`
	// The default URL to redirect customers to when they click on the portal's link to return to your website.
	ReturnUrl string `json:"return_url"`
}

type _CreatePortalSessionRequest CreatePortalSessionRequest

// NewCreatePortalSessionRequest instantiates a new CreatePortalSessionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreatePortalSessionRequest(customerId string, returnUrl string) *CreatePortalSessionRequest {
	this := CreatePortalSessionRequest{}
	this.CustomerId = customerId
	this.ReturnUrl = returnUrl
	return &this
}

// NewCreatePortalSessionRequestWithDefaults instantiates a new CreatePortalSessionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreatePortalSessionRequestWithDefaults() *CreatePortalSessionRequest {
	this := CreatePortalSessionRequest{}
	return &this
}

// GetCustomerId returns the CustomerId field value
func (o *CreatePortalSessionRequest) GetCustomerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value
// and a boolean to check if the value has been set.
func (o *CreatePortalSessionRequest) GetCustomerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomerId, true
}

// SetCustomerId sets field value
func (o *CreatePortalSessionRequest) SetCustomerId(v string) {
	o.CustomerId = v
}

// GetReturnUrl returns the ReturnUrl field value
func (o *CreatePortalSessionRequest) GetReturnUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReturnUrl
}

// GetReturnUrlOk returns a tuple with the ReturnUrl field value
// and a boolean to check if the value has been set.
func (o *CreatePortalSessionRequest) GetReturnUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReturnUrl, true
}

// SetReturnUrl sets field value
func (o *CreatePortalSessionRequest) SetReturnUrl(v string) {
	o.ReturnUrl = v
}

func (o CreatePortalSessionRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreatePortalSessionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["customer_id"] = o.CustomerId
	toSerialize["return_url"] = o.ReturnUrl
	return toSerialize, nil
}

func (o *CreatePortalSessionRequest) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"customer_id",
		"return_url",
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

	varCreatePortalSessionRequest := _CreatePortalSessionRequest{}

	err = json.Unmarshal(bytes, &varCreatePortalSessionRequest)

	if err != nil {
		return err
	}

	*o = CreatePortalSessionRequest(varCreatePortalSessionRequest)

	return err
}

type NullableCreatePortalSessionRequest struct {
	value *CreatePortalSessionRequest
	isSet bool
}

func (v NullableCreatePortalSessionRequest) Get() *CreatePortalSessionRequest {
	return v.value
}

func (v *NullableCreatePortalSessionRequest) Set(val *CreatePortalSessionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreatePortalSessionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreatePortalSessionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreatePortalSessionRequest(val *CreatePortalSessionRequest) *NullableCreatePortalSessionRequest {
	return &NullableCreatePortalSessionRequest{value: val, isSet: true}
}

func (v NullableCreatePortalSessionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreatePortalSessionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

