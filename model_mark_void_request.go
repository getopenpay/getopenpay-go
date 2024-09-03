/*
OpenPay API

super charge your subscription management.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package getopenpay

import (
	"encoding/json"
)

// checks if the MarkVoidRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MarkVoidRequest{}

// MarkVoidRequest struct for MarkVoidRequest
type MarkVoidRequest struct {
	Comment NullableString `json:"comment,omitempty"`
}

// NewMarkVoidRequest instantiates a new MarkVoidRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarkVoidRequest() *MarkVoidRequest {
	this := MarkVoidRequest{}
	return &this
}

// NewMarkVoidRequestWithDefaults instantiates a new MarkVoidRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarkVoidRequestWithDefaults() *MarkVoidRequest {
	this := MarkVoidRequest{}
	return &this
}

// GetComment returns the Comment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MarkVoidRequest) GetComment() string {
	if o == nil || IsNil(o.Comment.Get()) {
		var ret string
		return ret
	}
	return *o.Comment.Get()
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MarkVoidRequest) GetCommentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Comment.Get(), o.Comment.IsSet()
}

// HasComment returns a boolean if a field has been set.
func (o *MarkVoidRequest) HasComment() bool {
	if o != nil && o.Comment.IsSet() {
		return true
	}

	return false
}

// SetComment gets a reference to the given NullableString and assigns it to the Comment field.
func (o *MarkVoidRequest) SetComment(v string) {
	o.Comment.Set(&v)
}
// SetCommentNil sets the value for Comment to be an explicit nil
func (o *MarkVoidRequest) SetCommentNil() {
	o.Comment.Set(nil)
}

// UnsetComment ensures that no value is present for Comment, not even an explicit nil
func (o *MarkVoidRequest) UnsetComment() {
	o.Comment.Unset()
}

func (o MarkVoidRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MarkVoidRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Comment.IsSet() {
		toSerialize["comment"] = o.Comment.Get()
	}
	return toSerialize, nil
}

type NullableMarkVoidRequest struct {
	value *MarkVoidRequest
	isSet bool
}

func (v NullableMarkVoidRequest) Get() *MarkVoidRequest {
	return v.value
}

func (v *NullableMarkVoidRequest) Set(val *MarkVoidRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableMarkVoidRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableMarkVoidRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarkVoidRequest(val *MarkVoidRequest) *NullableMarkVoidRequest {
	return &NullableMarkVoidRequest{value: val, isSet: true}
}

func (v NullableMarkVoidRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarkVoidRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


