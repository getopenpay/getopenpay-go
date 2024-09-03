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

// checks if the ProductQueryParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductQueryParams{}

// ProductQueryParams struct for ProductQueryParams
type ProductQueryParams struct {
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
	IsActive NullableBool `json:"is_active,omitempty"`
	CreatedBy NullableString `json:"created_by,omitempty"`
	NameContains NullableString `json:"name_contains,omitempty"`
	Search NullableSearchFilter `json:"search,omitempty"`
}

// NewProductQueryParams instantiates a new ProductQueryParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductQueryParams() *ProductQueryParams {
	this := ProductQueryParams{}
	var pageNumber int32 = 1
	this.PageNumber = &pageNumber
	var pageSize int32 = 100
	this.PageSize = &pageSize
	var sortKey string = "created_at"
	this.SortKey = &sortKey
	var sortDescending bool = false
	this.SortDescending = &sortDescending
	return &this
}

// NewProductQueryParamsWithDefaults instantiates a new ProductQueryParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductQueryParamsWithDefaults() *ProductQueryParams {
	this := ProductQueryParams{}
	var pageNumber int32 = 1
	this.PageNumber = &pageNumber
	var pageSize int32 = 100
	this.PageSize = &pageSize
	var sortKey string = "created_at"
	this.SortKey = &sortKey
	var sortDescending bool = false
	this.SortDescending = &sortDescending
	return &this
}

// GetPageNumber returns the PageNumber field value if set, zero value otherwise.
func (o *ProductQueryParams) GetPageNumber() int32 {
	if o == nil || IsNil(o.PageNumber) {
		var ret int32
		return ret
	}
	return *o.PageNumber
}

// GetPageNumberOk returns a tuple with the PageNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductQueryParams) GetPageNumberOk() (*int32, bool) {
	if o == nil || IsNil(o.PageNumber) {
		return nil, false
	}
	return o.PageNumber, true
}

// HasPageNumber returns a boolean if a field has been set.
func (o *ProductQueryParams) HasPageNumber() bool {
	if o != nil && !IsNil(o.PageNumber) {
		return true
	}

	return false
}

// SetPageNumber gets a reference to the given int32 and assigns it to the PageNumber field.
func (o *ProductQueryParams) SetPageNumber(v int32) {
	o.PageNumber = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *ProductQueryParams) GetPageSize() int32 {
	if o == nil || IsNil(o.PageSize) {
		var ret int32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductQueryParams) GetPageSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *ProductQueryParams) HasPageSize() bool {
	if o != nil && !IsNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.
func (o *ProductQueryParams) SetPageSize(v int32) {
	o.PageSize = &v
}

// GetSortKey returns the SortKey field value if set, zero value otherwise.
func (o *ProductQueryParams) GetSortKey() string {
	if o == nil || IsNil(o.SortKey) {
		var ret string
		return ret
	}
	return *o.SortKey
}

// GetSortKeyOk returns a tuple with the SortKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductQueryParams) GetSortKeyOk() (*string, bool) {
	if o == nil || IsNil(o.SortKey) {
		return nil, false
	}
	return o.SortKey, true
}

// HasSortKey returns a boolean if a field has been set.
func (o *ProductQueryParams) HasSortKey() bool {
	if o != nil && !IsNil(o.SortKey) {
		return true
	}

	return false
}

// SetSortKey gets a reference to the given string and assigns it to the SortKey field.
func (o *ProductQueryParams) SetSortKey(v string) {
	o.SortKey = &v
}

// GetSortDescending returns the SortDescending field value if set, zero value otherwise.
func (o *ProductQueryParams) GetSortDescending() bool {
	if o == nil || IsNil(o.SortDescending) {
		var ret bool
		return ret
	}
	return *o.SortDescending
}

// GetSortDescendingOk returns a tuple with the SortDescending field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductQueryParams) GetSortDescendingOk() (*bool, bool) {
	if o == nil || IsNil(o.SortDescending) {
		return nil, false
	}
	return o.SortDescending, true
}

// HasSortDescending returns a boolean if a field has been set.
func (o *ProductQueryParams) HasSortDescending() bool {
	if o != nil && !IsNil(o.SortDescending) {
		return true
	}

	return false
}

// SetSortDescending gets a reference to the given bool and assigns it to the SortDescending field.
func (o *ProductQueryParams) SetSortDescending(v bool) {
	o.SortDescending = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProductQueryParams) GetCreatedAt() DateTimeFilter {
	if o == nil || IsNil(o.CreatedAt.Get()) {
		var ret DateTimeFilter
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProductQueryParams) GetCreatedAtOk() (*DateTimeFilter, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ProductQueryParams) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableDateTimeFilter and assigns it to the CreatedAt field.
func (o *ProductQueryParams) SetCreatedAt(v DateTimeFilter) {
	o.CreatedAt.Set(&v)
}
// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *ProductQueryParams) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *ProductQueryParams) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProductQueryParams) GetUpdatedAt() DateTimeFilter {
	if o == nil || IsNil(o.UpdatedAt.Get()) {
		var ret DateTimeFilter
		return ret
	}
	return *o.UpdatedAt.Get()
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProductQueryParams) GetUpdatedAtOk() (*DateTimeFilter, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedAt.Get(), o.UpdatedAt.IsSet()
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ProductQueryParams) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt.IsSet() {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given NullableDateTimeFilter and assigns it to the UpdatedAt field.
func (o *ProductQueryParams) SetUpdatedAt(v DateTimeFilter) {
	o.UpdatedAt.Set(&v)
}
// SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil
func (o *ProductQueryParams) SetUpdatedAtNil() {
	o.UpdatedAt.Set(nil)
}

// UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
func (o *ProductQueryParams) UnsetUpdatedAt() {
	o.UpdatedAt.Unset()
}

// GetExpand returns the Expand field value if set, zero value otherwise.
func (o *ProductQueryParams) GetExpand() []string {
	if o == nil || IsNil(o.Expand) {
		var ret []string
		return ret
	}
	return o.Expand
}

// GetExpandOk returns a tuple with the Expand field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductQueryParams) GetExpandOk() ([]string, bool) {
	if o == nil || IsNil(o.Expand) {
		return nil, false
	}
	return o.Expand, true
}

// HasExpand returns a boolean if a field has been set.
func (o *ProductQueryParams) HasExpand() bool {
	if o != nil && !IsNil(o.Expand) {
		return true
	}

	return false
}

// SetExpand gets a reference to the given []string and assigns it to the Expand field.
func (o *ProductQueryParams) SetExpand(v []string) {
	o.Expand = v
}

// GetIsActive returns the IsActive field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProductQueryParams) GetIsActive() bool {
	if o == nil || IsNil(o.IsActive.Get()) {
		var ret bool
		return ret
	}
	return *o.IsActive.Get()
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProductQueryParams) GetIsActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsActive.Get(), o.IsActive.IsSet()
}

// HasIsActive returns a boolean if a field has been set.
func (o *ProductQueryParams) HasIsActive() bool {
	if o != nil && o.IsActive.IsSet() {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given NullableBool and assigns it to the IsActive field.
func (o *ProductQueryParams) SetIsActive(v bool) {
	o.IsActive.Set(&v)
}
// SetIsActiveNil sets the value for IsActive to be an explicit nil
func (o *ProductQueryParams) SetIsActiveNil() {
	o.IsActive.Set(nil)
}

// UnsetIsActive ensures that no value is present for IsActive, not even an explicit nil
func (o *ProductQueryParams) UnsetIsActive() {
	o.IsActive.Unset()
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProductQueryParams) GetCreatedBy() string {
	if o == nil || IsNil(o.CreatedBy.Get()) {
		var ret string
		return ret
	}
	return *o.CreatedBy.Get()
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProductQueryParams) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedBy.Get(), o.CreatedBy.IsSet()
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *ProductQueryParams) HasCreatedBy() bool {
	if o != nil && o.CreatedBy.IsSet() {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given NullableString and assigns it to the CreatedBy field.
func (o *ProductQueryParams) SetCreatedBy(v string) {
	o.CreatedBy.Set(&v)
}
// SetCreatedByNil sets the value for CreatedBy to be an explicit nil
func (o *ProductQueryParams) SetCreatedByNil() {
	o.CreatedBy.Set(nil)
}

// UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
func (o *ProductQueryParams) UnsetCreatedBy() {
	o.CreatedBy.Unset()
}

// GetNameContains returns the NameContains field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProductQueryParams) GetNameContains() string {
	if o == nil || IsNil(o.NameContains.Get()) {
		var ret string
		return ret
	}
	return *o.NameContains.Get()
}

// GetNameContainsOk returns a tuple with the NameContains field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProductQueryParams) GetNameContainsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NameContains.Get(), o.NameContains.IsSet()
}

// HasNameContains returns a boolean if a field has been set.
func (o *ProductQueryParams) HasNameContains() bool {
	if o != nil && o.NameContains.IsSet() {
		return true
	}

	return false
}

// SetNameContains gets a reference to the given NullableString and assigns it to the NameContains field.
func (o *ProductQueryParams) SetNameContains(v string) {
	o.NameContains.Set(&v)
}
// SetNameContainsNil sets the value for NameContains to be an explicit nil
func (o *ProductQueryParams) SetNameContainsNil() {
	o.NameContains.Set(nil)
}

// UnsetNameContains ensures that no value is present for NameContains, not even an explicit nil
func (o *ProductQueryParams) UnsetNameContains() {
	o.NameContains.Unset()
}

// GetSearch returns the Search field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProductQueryParams) GetSearch() SearchFilter {
	if o == nil || IsNil(o.Search.Get()) {
		var ret SearchFilter
		return ret
	}
	return *o.Search.Get()
}

// GetSearchOk returns a tuple with the Search field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProductQueryParams) GetSearchOk() (*SearchFilter, bool) {
	if o == nil {
		return nil, false
	}
	return o.Search.Get(), o.Search.IsSet()
}

// HasSearch returns a boolean if a field has been set.
func (o *ProductQueryParams) HasSearch() bool {
	if o != nil && o.Search.IsSet() {
		return true
	}

	return false
}

// SetSearch gets a reference to the given NullableSearchFilter and assigns it to the Search field.
func (o *ProductQueryParams) SetSearch(v SearchFilter) {
	o.Search.Set(&v)
}
// SetSearchNil sets the value for Search to be an explicit nil
func (o *ProductQueryParams) SetSearchNil() {
	o.Search.Set(nil)
}

// UnsetSearch ensures that no value is present for Search, not even an explicit nil
func (o *ProductQueryParams) UnsetSearch() {
	o.Search.Unset()
}

func (o ProductQueryParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductQueryParams) ToMap() (map[string]interface{}, error) {
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
	if o.IsActive.IsSet() {
		toSerialize["is_active"] = o.IsActive.Get()
	}
	if o.CreatedBy.IsSet() {
		toSerialize["created_by"] = o.CreatedBy.Get()
	}
	if o.NameContains.IsSet() {
		toSerialize["name_contains"] = o.NameContains.Get()
	}
	if o.Search.IsSet() {
		toSerialize["search"] = o.Search.Get()
	}
	return toSerialize, nil
}

type NullableProductQueryParams struct {
	value *ProductQueryParams
	isSet bool
}

func (v NullableProductQueryParams) Get() *ProductQueryParams {
	return v.value
}

func (v *NullableProductQueryParams) Set(val *ProductQueryParams) {
	v.value = val
	v.isSet = true
}

func (v NullableProductQueryParams) IsSet() bool {
	return v.isSet
}

func (v *NullableProductQueryParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductQueryParams(val *ProductQueryParams) *NullableProductQueryParams {
	return &NullableProductQueryParams{value: val, isSet: true}
}

func (v NullableProductQueryParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductQueryParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


