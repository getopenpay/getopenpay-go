/*
OpenPay API

super charge your subscription management.

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package getopenpay

import (
	"encoding/json"
)

// checks if the CreateDisputeRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateDisputeRequest{}

// CreateDisputeRequest struct for CreateDisputeRequest
type CreateDisputeRequest struct {
	Meta map[string]interface{} `json:"meta,omitempty"`
	PaymentIntentId NullableString `json:"payment_intent_id,omitempty"`
	Reason NullableString `json:"reason,omitempty"`
	TheirDisputeId NullableString `json:"their_dispute_id,omitempty"`
	TheirPaymentIntentId NullableString `json:"their_payment_intent_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateDisputeRequest CreateDisputeRequest

// NewCreateDisputeRequest instantiates a new CreateDisputeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateDisputeRequest() *CreateDisputeRequest {
	this := CreateDisputeRequest{}
	return &this
}

// NewCreateDisputeRequestWithDefaults instantiates a new CreateDisputeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateDisputeRequestWithDefaults() *CreateDisputeRequest {
	this := CreateDisputeRequest{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateDisputeRequest) GetMeta() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateDisputeRequest) GetMetaOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Meta) {
		return map[string]interface{}{}, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *CreateDisputeRequest) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given map[string]interface{} and assigns it to the Meta field.
func (o *CreateDisputeRequest) SetMeta(v map[string]interface{}) {
	o.Meta = v
}

// GetPaymentIntentId returns the PaymentIntentId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateDisputeRequest) GetPaymentIntentId() string {
	if o == nil || IsNil(o.PaymentIntentId.Get()) {
		var ret string
		return ret
	}
	return *o.PaymentIntentId.Get()
}

// GetPaymentIntentIdOk returns a tuple with the PaymentIntentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateDisputeRequest) GetPaymentIntentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PaymentIntentId.Get(), o.PaymentIntentId.IsSet()
}

// HasPaymentIntentId returns a boolean if a field has been set.
func (o *CreateDisputeRequest) HasPaymentIntentId() bool {
	if o != nil && o.PaymentIntentId.IsSet() {
		return true
	}

	return false
}

// SetPaymentIntentId gets a reference to the given NullableString and assigns it to the PaymentIntentId field.
func (o *CreateDisputeRequest) SetPaymentIntentId(v string) {
	o.PaymentIntentId.Set(&v)
}
// SetPaymentIntentIdNil sets the value for PaymentIntentId to be an explicit nil
func (o *CreateDisputeRequest) SetPaymentIntentIdNil() {
	o.PaymentIntentId.Set(nil)
}

// UnsetPaymentIntentId ensures that no value is present for PaymentIntentId, not even an explicit nil
func (o *CreateDisputeRequest) UnsetPaymentIntentId() {
	o.PaymentIntentId.Unset()
}

// GetReason returns the Reason field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateDisputeRequest) GetReason() string {
	if o == nil || IsNil(o.Reason.Get()) {
		var ret string
		return ret
	}
	return *o.Reason.Get()
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateDisputeRequest) GetReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Reason.Get(), o.Reason.IsSet()
}

// HasReason returns a boolean if a field has been set.
func (o *CreateDisputeRequest) HasReason() bool {
	if o != nil && o.Reason.IsSet() {
		return true
	}

	return false
}

// SetReason gets a reference to the given NullableString and assigns it to the Reason field.
func (o *CreateDisputeRequest) SetReason(v string) {
	o.Reason.Set(&v)
}
// SetReasonNil sets the value for Reason to be an explicit nil
func (o *CreateDisputeRequest) SetReasonNil() {
	o.Reason.Set(nil)
}

// UnsetReason ensures that no value is present for Reason, not even an explicit nil
func (o *CreateDisputeRequest) UnsetReason() {
	o.Reason.Unset()
}

// GetTheirDisputeId returns the TheirDisputeId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateDisputeRequest) GetTheirDisputeId() string {
	if o == nil || IsNil(o.TheirDisputeId.Get()) {
		var ret string
		return ret
	}
	return *o.TheirDisputeId.Get()
}

// GetTheirDisputeIdOk returns a tuple with the TheirDisputeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateDisputeRequest) GetTheirDisputeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TheirDisputeId.Get(), o.TheirDisputeId.IsSet()
}

// HasTheirDisputeId returns a boolean if a field has been set.
func (o *CreateDisputeRequest) HasTheirDisputeId() bool {
	if o != nil && o.TheirDisputeId.IsSet() {
		return true
	}

	return false
}

// SetTheirDisputeId gets a reference to the given NullableString and assigns it to the TheirDisputeId field.
func (o *CreateDisputeRequest) SetTheirDisputeId(v string) {
	o.TheirDisputeId.Set(&v)
}
// SetTheirDisputeIdNil sets the value for TheirDisputeId to be an explicit nil
func (o *CreateDisputeRequest) SetTheirDisputeIdNil() {
	o.TheirDisputeId.Set(nil)
}

// UnsetTheirDisputeId ensures that no value is present for TheirDisputeId, not even an explicit nil
func (o *CreateDisputeRequest) UnsetTheirDisputeId() {
	o.TheirDisputeId.Unset()
}

// GetTheirPaymentIntentId returns the TheirPaymentIntentId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateDisputeRequest) GetTheirPaymentIntentId() string {
	if o == nil || IsNil(o.TheirPaymentIntentId.Get()) {
		var ret string
		return ret
	}
	return *o.TheirPaymentIntentId.Get()
}

// GetTheirPaymentIntentIdOk returns a tuple with the TheirPaymentIntentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateDisputeRequest) GetTheirPaymentIntentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TheirPaymentIntentId.Get(), o.TheirPaymentIntentId.IsSet()
}

// HasTheirPaymentIntentId returns a boolean if a field has been set.
func (o *CreateDisputeRequest) HasTheirPaymentIntentId() bool {
	if o != nil && o.TheirPaymentIntentId.IsSet() {
		return true
	}

	return false
}

// SetTheirPaymentIntentId gets a reference to the given NullableString and assigns it to the TheirPaymentIntentId field.
func (o *CreateDisputeRequest) SetTheirPaymentIntentId(v string) {
	o.TheirPaymentIntentId.Set(&v)
}
// SetTheirPaymentIntentIdNil sets the value for TheirPaymentIntentId to be an explicit nil
func (o *CreateDisputeRequest) SetTheirPaymentIntentIdNil() {
	o.TheirPaymentIntentId.Set(nil)
}

// UnsetTheirPaymentIntentId ensures that no value is present for TheirPaymentIntentId, not even an explicit nil
func (o *CreateDisputeRequest) UnsetTheirPaymentIntentId() {
	o.TheirPaymentIntentId.Unset()
}

func (o CreateDisputeRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateDisputeRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}
	if o.PaymentIntentId.IsSet() {
		toSerialize["payment_intent_id"] = o.PaymentIntentId.Get()
	}
	if o.Reason.IsSet() {
		toSerialize["reason"] = o.Reason.Get()
	}
	if o.TheirDisputeId.IsSet() {
		toSerialize["their_dispute_id"] = o.TheirDisputeId.Get()
	}
	if o.TheirPaymentIntentId.IsSet() {
		toSerialize["their_payment_intent_id"] = o.TheirPaymentIntentId.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateDisputeRequest) UnmarshalJSON(data []byte) (err error) {
	varCreateDisputeRequest := _CreateDisputeRequest{}

	err = json.Unmarshal(data, &varCreateDisputeRequest)

	if err != nil {
		return err
	}

	*o = CreateDisputeRequest(varCreateDisputeRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "meta")
		delete(additionalProperties, "payment_intent_id")
		delete(additionalProperties, "reason")
		delete(additionalProperties, "their_dispute_id")
		delete(additionalProperties, "their_payment_intent_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateDisputeRequest struct {
	value *CreateDisputeRequest
	isSet bool
}

func (v NullableCreateDisputeRequest) Get() *CreateDisputeRequest {
	return v.value
}

func (v *NullableCreateDisputeRequest) Set(val *CreateDisputeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateDisputeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateDisputeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateDisputeRequest(val *CreateDisputeRequest) *NullableCreateDisputeRequest {
	return &NullableCreateDisputeRequest{value: val, isSet: true}
}

func (v NullableCreateDisputeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateDisputeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


