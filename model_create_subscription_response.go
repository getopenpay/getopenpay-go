/*
OpenPay API

super charge your subscription management.

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package getopenpay

import (
	"encoding/json"
	"fmt"
)

// checks if the CreateSubscriptionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateSubscriptionResponse{}

// CreateSubscriptionResponse struct for CreateSubscriptionResponse
type CreateSubscriptionResponse struct {
	// List of subscriptions created.
	Created []SubscriptionExternal `json:"created"`
	// List of invoices created
	Invoices []InvoiceExternal `json:"invoices"`
	// List of successful processor IDs used for creating the subscriptions
	ProcessorsUsed []string `json:"processors_used"`
	AdditionalProperties map[string]interface{}
}

type _CreateSubscriptionResponse CreateSubscriptionResponse

// NewCreateSubscriptionResponse instantiates a new CreateSubscriptionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateSubscriptionResponse(created []SubscriptionExternal, invoices []InvoiceExternal, processorsUsed []string) *CreateSubscriptionResponse {
	this := CreateSubscriptionResponse{}
	this.Created = created
	this.Invoices = invoices
	this.ProcessorsUsed = processorsUsed
	return &this
}

// NewCreateSubscriptionResponseWithDefaults instantiates a new CreateSubscriptionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateSubscriptionResponseWithDefaults() *CreateSubscriptionResponse {
	this := CreateSubscriptionResponse{}
	return &this
}

// GetCreated returns the Created field value
func (o *CreateSubscriptionResponse) GetCreated() []SubscriptionExternal {
	if o == nil {
		var ret []SubscriptionExternal
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *CreateSubscriptionResponse) GetCreatedOk() ([]SubscriptionExternal, bool) {
	if o == nil {
		return nil, false
	}
	return o.Created, true
}

// SetCreated sets field value
func (o *CreateSubscriptionResponse) SetCreated(v []SubscriptionExternal) {
	o.Created = v
}

// GetInvoices returns the Invoices field value
func (o *CreateSubscriptionResponse) GetInvoices() []InvoiceExternal {
	if o == nil {
		var ret []InvoiceExternal
		return ret
	}

	return o.Invoices
}

// GetInvoicesOk returns a tuple with the Invoices field value
// and a boolean to check if the value has been set.
func (o *CreateSubscriptionResponse) GetInvoicesOk() ([]InvoiceExternal, bool) {
	if o == nil {
		return nil, false
	}
	return o.Invoices, true
}

// SetInvoices sets field value
func (o *CreateSubscriptionResponse) SetInvoices(v []InvoiceExternal) {
	o.Invoices = v
}

// GetProcessorsUsed returns the ProcessorsUsed field value
func (o *CreateSubscriptionResponse) GetProcessorsUsed() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ProcessorsUsed
}

// GetProcessorsUsedOk returns a tuple with the ProcessorsUsed field value
// and a boolean to check if the value has been set.
func (o *CreateSubscriptionResponse) GetProcessorsUsedOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProcessorsUsed, true
}

// SetProcessorsUsed sets field value
func (o *CreateSubscriptionResponse) SetProcessorsUsed(v []string) {
	o.ProcessorsUsed = v
}

func (o CreateSubscriptionResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateSubscriptionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["created"] = o.Created
	toSerialize["invoices"] = o.Invoices
	toSerialize["processors_used"] = o.ProcessorsUsed

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateSubscriptionResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"created",
		"invoices",
		"processors_used",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCreateSubscriptionResponse := _CreateSubscriptionResponse{}

	err = json.Unmarshal(data, &varCreateSubscriptionResponse)

	if err != nil {
		return err
	}

	*o = CreateSubscriptionResponse(varCreateSubscriptionResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "created")
		delete(additionalProperties, "invoices")
		delete(additionalProperties, "processors_used")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateSubscriptionResponse struct {
	value *CreateSubscriptionResponse
	isSet bool
}

func (v NullableCreateSubscriptionResponse) Get() *CreateSubscriptionResponse {
	return v.value
}

func (v *NullableCreateSubscriptionResponse) Set(val *CreateSubscriptionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSubscriptionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSubscriptionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSubscriptionResponse(val *CreateSubscriptionResponse) *NullableCreateSubscriptionResponse {
	return &NullableCreateSubscriptionResponse{value: val, isSet: true}
}

func (v NullableCreateSubscriptionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateSubscriptionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


