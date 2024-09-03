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

// checks if the ListResponseCheckoutSessionExternal type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListResponseCheckoutSessionExternal{}

// ListResponseCheckoutSessionExternal struct for ListResponseCheckoutSessionExternal
type ListResponseCheckoutSessionExternal struct {
	Data []CheckoutSessionExternal `json:"data"`
	TotalObjects int32 `json:"total_objects"`
	PageNumber int32 `json:"page_number"`
	PageSize int32 `json:"page_size"`
}

type _ListResponseCheckoutSessionExternal ListResponseCheckoutSessionExternal

// NewListResponseCheckoutSessionExternal instantiates a new ListResponseCheckoutSessionExternal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListResponseCheckoutSessionExternal(data []CheckoutSessionExternal, totalObjects int32, pageNumber int32, pageSize int32) *ListResponseCheckoutSessionExternal {
	this := ListResponseCheckoutSessionExternal{}
	this.Data = data
	this.TotalObjects = totalObjects
	this.PageNumber = pageNumber
	this.PageSize = pageSize
	return &this
}

// NewListResponseCheckoutSessionExternalWithDefaults instantiates a new ListResponseCheckoutSessionExternal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListResponseCheckoutSessionExternalWithDefaults() *ListResponseCheckoutSessionExternal {
	this := ListResponseCheckoutSessionExternal{}
	return &this
}

// GetData returns the Data field value
func (o *ListResponseCheckoutSessionExternal) GetData() []CheckoutSessionExternal {
	if o == nil {
		var ret []CheckoutSessionExternal
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ListResponseCheckoutSessionExternal) GetDataOk() ([]CheckoutSessionExternal, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *ListResponseCheckoutSessionExternal) SetData(v []CheckoutSessionExternal) {
	o.Data = v
}

// GetTotalObjects returns the TotalObjects field value
func (o *ListResponseCheckoutSessionExternal) GetTotalObjects() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalObjects
}

// GetTotalObjectsOk returns a tuple with the TotalObjects field value
// and a boolean to check if the value has been set.
func (o *ListResponseCheckoutSessionExternal) GetTotalObjectsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalObjects, true
}

// SetTotalObjects sets field value
func (o *ListResponseCheckoutSessionExternal) SetTotalObjects(v int32) {
	o.TotalObjects = v
}

// GetPageNumber returns the PageNumber field value
func (o *ListResponseCheckoutSessionExternal) GetPageNumber() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PageNumber
}

// GetPageNumberOk returns a tuple with the PageNumber field value
// and a boolean to check if the value has been set.
func (o *ListResponseCheckoutSessionExternal) GetPageNumberOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PageNumber, true
}

// SetPageNumber sets field value
func (o *ListResponseCheckoutSessionExternal) SetPageNumber(v int32) {
	o.PageNumber = v
}

// GetPageSize returns the PageSize field value
func (o *ListResponseCheckoutSessionExternal) GetPageSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value
// and a boolean to check if the value has been set.
func (o *ListResponseCheckoutSessionExternal) GetPageSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PageSize, true
}

// SetPageSize sets field value
func (o *ListResponseCheckoutSessionExternal) SetPageSize(v int32) {
	o.PageSize = v
}

func (o ListResponseCheckoutSessionExternal) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListResponseCheckoutSessionExternal) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["total_objects"] = o.TotalObjects
	toSerialize["page_number"] = o.PageNumber
	toSerialize["page_size"] = o.PageSize
	return toSerialize, nil
}

func (o *ListResponseCheckoutSessionExternal) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
		"total_objects",
		"page_number",
		"page_size",
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

	varListResponseCheckoutSessionExternal := _ListResponseCheckoutSessionExternal{}

	err = json.Unmarshal(bytes, &varListResponseCheckoutSessionExternal)

	if err != nil {
		return err
	}

	*o = ListResponseCheckoutSessionExternal(varListResponseCheckoutSessionExternal)

	return err
}

type NullableListResponseCheckoutSessionExternal struct {
	value *ListResponseCheckoutSessionExternal
	isSet bool
}

func (v NullableListResponseCheckoutSessionExternal) Get() *ListResponseCheckoutSessionExternal {
	return v.value
}

func (v *NullableListResponseCheckoutSessionExternal) Set(val *ListResponseCheckoutSessionExternal) {
	v.value = val
	v.isSet = true
}

func (v NullableListResponseCheckoutSessionExternal) IsSet() bool {
	return v.isSet
}

func (v *NullableListResponseCheckoutSessionExternal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListResponseCheckoutSessionExternal(val *ListResponseCheckoutSessionExternal) *NullableListResponseCheckoutSessionExternal {
	return &NullableListResponseCheckoutSessionExternal{value: val, isSet: true}
}

func (v NullableListResponseCheckoutSessionExternal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListResponseCheckoutSessionExternal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


