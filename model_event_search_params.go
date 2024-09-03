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

// checks if the EventSearchParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EventSearchParams{}

// EventSearchParams struct for EventSearchParams
type EventSearchParams struct {
	// Page number
	PageNumber *int32 `json:"page_number,omitempty"`
	// Page size
	PageSize *int32 `json:"page_size,omitempty"`
	// Key name based on which data is sorted.
	SortKey *string `json:"sort_key,omitempty"`
	// Sort direction.
	SortDescending *bool `json:"sort_descending,omitempty"`
	CreatedAt NullableDateTimeFilter `json:"created_at,omitempty"`
	UpdatedAt NullableDateTimeFilter `json:"updated_at,omitempty"`
	Expand []string `json:"expand,omitempty"`
	// Unique ID of the object.
	ObjectId string `json:"object_id"`
	Types []EventType `json:"types,omitempty"`
	// Maximum number of documents to retrieve.
	Limit *int32 `json:"limit,omitempty"`
}

type _EventSearchParams EventSearchParams

// NewEventSearchParams instantiates a new EventSearchParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventSearchParams(objectId string) *EventSearchParams {
	this := EventSearchParams{}
	var pageNumber int32 = 1
	this.PageNumber = &pageNumber
	var pageSize int32 = 100
	this.PageSize = &pageSize
	var sortKey string = "created_at"
	this.SortKey = &sortKey
	var sortDescending bool = false
	this.SortDescending = &sortDescending
	this.ObjectId = objectId
	var limit int32 = 50
	this.Limit = &limit
	return &this
}

// NewEventSearchParamsWithDefaults instantiates a new EventSearchParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventSearchParamsWithDefaults() *EventSearchParams {
	this := EventSearchParams{}
	var pageNumber int32 = 1
	this.PageNumber = &pageNumber
	var pageSize int32 = 100
	this.PageSize = &pageSize
	var sortKey string = "created_at"
	this.SortKey = &sortKey
	var sortDescending bool = false
	this.SortDescending = &sortDescending
	var limit int32 = 50
	this.Limit = &limit
	return &this
}

// GetPageNumber returns the PageNumber field value if set, zero value otherwise.
func (o *EventSearchParams) GetPageNumber() int32 {
	if o == nil || IsNil(o.PageNumber) {
		var ret int32
		return ret
	}
	return *o.PageNumber
}

// GetPageNumberOk returns a tuple with the PageNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventSearchParams) GetPageNumberOk() (*int32, bool) {
	if o == nil || IsNil(o.PageNumber) {
		return nil, false
	}
	return o.PageNumber, true
}

// HasPageNumber returns a boolean if a field has been set.
func (o *EventSearchParams) HasPageNumber() bool {
	if o != nil && !IsNil(o.PageNumber) {
		return true
	}

	return false
}

// SetPageNumber gets a reference to the given int32 and assigns it to the PageNumber field.
func (o *EventSearchParams) SetPageNumber(v int32) {
	o.PageNumber = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *EventSearchParams) GetPageSize() int32 {
	if o == nil || IsNil(o.PageSize) {
		var ret int32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventSearchParams) GetPageSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *EventSearchParams) HasPageSize() bool {
	if o != nil && !IsNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.
func (o *EventSearchParams) SetPageSize(v int32) {
	o.PageSize = &v
}

// GetSortKey returns the SortKey field value if set, zero value otherwise.
func (o *EventSearchParams) GetSortKey() string {
	if o == nil || IsNil(o.SortKey) {
		var ret string
		return ret
	}
	return *o.SortKey
}

// GetSortKeyOk returns a tuple with the SortKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventSearchParams) GetSortKeyOk() (*string, bool) {
	if o == nil || IsNil(o.SortKey) {
		return nil, false
	}
	return o.SortKey, true
}

// HasSortKey returns a boolean if a field has been set.
func (o *EventSearchParams) HasSortKey() bool {
	if o != nil && !IsNil(o.SortKey) {
		return true
	}

	return false
}

// SetSortKey gets a reference to the given string and assigns it to the SortKey field.
func (o *EventSearchParams) SetSortKey(v string) {
	o.SortKey = &v
}

// GetSortDescending returns the SortDescending field value if set, zero value otherwise.
func (o *EventSearchParams) GetSortDescending() bool {
	if o == nil || IsNil(o.SortDescending) {
		var ret bool
		return ret
	}
	return *o.SortDescending
}

// GetSortDescendingOk returns a tuple with the SortDescending field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventSearchParams) GetSortDescendingOk() (*bool, bool) {
	if o == nil || IsNil(o.SortDescending) {
		return nil, false
	}
	return o.SortDescending, true
}

// HasSortDescending returns a boolean if a field has been set.
func (o *EventSearchParams) HasSortDescending() bool {
	if o != nil && !IsNil(o.SortDescending) {
		return true
	}

	return false
}

// SetSortDescending gets a reference to the given bool and assigns it to the SortDescending field.
func (o *EventSearchParams) SetSortDescending(v bool) {
	o.SortDescending = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventSearchParams) GetCreatedAt() DateTimeFilter {
	if o == nil || IsNil(o.CreatedAt.Get()) {
		var ret DateTimeFilter
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EventSearchParams) GetCreatedAtOk() (*DateTimeFilter, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *EventSearchParams) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableDateTimeFilter and assigns it to the CreatedAt field.
func (o *EventSearchParams) SetCreatedAt(v DateTimeFilter) {
	o.CreatedAt.Set(&v)
}
// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *EventSearchParams) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *EventSearchParams) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventSearchParams) GetUpdatedAt() DateTimeFilter {
	if o == nil || IsNil(o.UpdatedAt.Get()) {
		var ret DateTimeFilter
		return ret
	}
	return *o.UpdatedAt.Get()
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EventSearchParams) GetUpdatedAtOk() (*DateTimeFilter, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedAt.Get(), o.UpdatedAt.IsSet()
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *EventSearchParams) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt.IsSet() {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given NullableDateTimeFilter and assigns it to the UpdatedAt field.
func (o *EventSearchParams) SetUpdatedAt(v DateTimeFilter) {
	o.UpdatedAt.Set(&v)
}
// SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil
func (o *EventSearchParams) SetUpdatedAtNil() {
	o.UpdatedAt.Set(nil)
}

// UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
func (o *EventSearchParams) UnsetUpdatedAt() {
	o.UpdatedAt.Unset()
}

// GetExpand returns the Expand field value if set, zero value otherwise.
func (o *EventSearchParams) GetExpand() []string {
	if o == nil || IsNil(o.Expand) {
		var ret []string
		return ret
	}
	return o.Expand
}

// GetExpandOk returns a tuple with the Expand field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventSearchParams) GetExpandOk() ([]string, bool) {
	if o == nil || IsNil(o.Expand) {
		return nil, false
	}
	return o.Expand, true
}

// HasExpand returns a boolean if a field has been set.
func (o *EventSearchParams) HasExpand() bool {
	if o != nil && !IsNil(o.Expand) {
		return true
	}

	return false
}

// SetExpand gets a reference to the given []string and assigns it to the Expand field.
func (o *EventSearchParams) SetExpand(v []string) {
	o.Expand = v
}

// GetObjectId returns the ObjectId field value
func (o *EventSearchParams) GetObjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectId
}

// GetObjectIdOk returns a tuple with the ObjectId field value
// and a boolean to check if the value has been set.
func (o *EventSearchParams) GetObjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectId, true
}

// SetObjectId sets field value
func (o *EventSearchParams) SetObjectId(v string) {
	o.ObjectId = v
}

// GetTypes returns the Types field value if set, zero value otherwise.
func (o *EventSearchParams) GetTypes() []EventType {
	if o == nil || IsNil(o.Types) {
		var ret []EventType
		return ret
	}
	return o.Types
}

// GetTypesOk returns a tuple with the Types field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventSearchParams) GetTypesOk() ([]EventType, bool) {
	if o == nil || IsNil(o.Types) {
		return nil, false
	}
	return o.Types, true
}

// HasTypes returns a boolean if a field has been set.
func (o *EventSearchParams) HasTypes() bool {
	if o != nil && !IsNil(o.Types) {
		return true
	}

	return false
}

// SetTypes gets a reference to the given []EventType and assigns it to the Types field.
func (o *EventSearchParams) SetTypes(v []EventType) {
	o.Types = v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *EventSearchParams) GetLimit() int32 {
	if o == nil || IsNil(o.Limit) {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventSearchParams) GetLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *EventSearchParams) HasLimit() bool {
	if o != nil && !IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *EventSearchParams) SetLimit(v int32) {
	o.Limit = &v
}

func (o EventSearchParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EventSearchParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PageNumber) {
		toSerialize["page_number"] = o.PageNumber
	}
	if !IsNil(o.PageSize) {
		toSerialize["page_size"] = o.PageSize
	}
	if !IsNil(o.SortKey) {
		toSerialize["sort_key"] = o.SortKey
	}
	if !IsNil(o.SortDescending) {
		toSerialize["sort_descending"] = o.SortDescending
	}
	if o.CreatedAt.IsSet() {
		toSerialize["created_at"] = o.CreatedAt.Get()
	}
	if o.UpdatedAt.IsSet() {
		toSerialize["updated_at"] = o.UpdatedAt.Get()
	}
	if !IsNil(o.Expand) {
		toSerialize["expand"] = o.Expand
	}
	toSerialize["object_id"] = o.ObjectId
	if !IsNil(o.Types) {
		toSerialize["types"] = o.Types
	}
	if !IsNil(o.Limit) {
		toSerialize["limit"] = o.Limit
	}
	return toSerialize, nil
}

func (o *EventSearchParams) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"object_id",
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

	varEventSearchParams := _EventSearchParams{}

	err = json.Unmarshal(bytes, &varEventSearchParams)

	if err != nil {
		return err
	}

	*o = EventSearchParams(varEventSearchParams)

	return err
}

type NullableEventSearchParams struct {
	value *EventSearchParams
	isSet bool
}

func (v NullableEventSearchParams) Get() *EventSearchParams {
	return v.value
}

func (v *NullableEventSearchParams) Set(val *EventSearchParams) {
	v.value = val
	v.isSet = true
}

func (v NullableEventSearchParams) IsSet() bool {
	return v.isSet
}

func (v *NullableEventSearchParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventSearchParams(val *EventSearchParams) *NullableEventSearchParams {
	return &NullableEventSearchParams{value: val, isSet: true}
}

func (v NullableEventSearchParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventSearchParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


