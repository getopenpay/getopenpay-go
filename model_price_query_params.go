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

// checks if the PriceQueryParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PriceQueryParams{}

// PriceQueryParams struct for PriceQueryParams
type PriceQueryParams struct {
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
	ProductId NullableString `json:"product_id,omitempty"`
	IsActive NullableBool `json:"is_active,omitempty"`
	Currency NullableCurrencyEnum `json:"currency,omitempty"`
	PriceType NullablePriceTypeEnum `json:"price_type,omitempty"`
}

// NewPriceQueryParams instantiates a new PriceQueryParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPriceQueryParams() *PriceQueryParams {
	this := PriceQueryParams{}
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

// NewPriceQueryParamsWithDefaults instantiates a new PriceQueryParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPriceQueryParamsWithDefaults() *PriceQueryParams {
	this := PriceQueryParams{}
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
func (o *PriceQueryParams) GetPageNumber() int32 {
	if o == nil || IsNil(o.PageNumber) {
		var ret int32
		return ret
	}
	return *o.PageNumber
}

// GetPageNumberOk returns a tuple with the PageNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceQueryParams) GetPageNumberOk() (*int32, bool) {
	if o == nil || IsNil(o.PageNumber) {
		return nil, false
	}
	return o.PageNumber, true
}

// HasPageNumber returns a boolean if a field has been set.
func (o *PriceQueryParams) HasPageNumber() bool {
	if o != nil && !IsNil(o.PageNumber) {
		return true
	}

	return false
}

// SetPageNumber gets a reference to the given int32 and assigns it to the PageNumber field.
func (o *PriceQueryParams) SetPageNumber(v int32) {
	o.PageNumber = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *PriceQueryParams) GetPageSize() int32 {
	if o == nil || IsNil(o.PageSize) {
		var ret int32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceQueryParams) GetPageSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *PriceQueryParams) HasPageSize() bool {
	if o != nil && !IsNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.
func (o *PriceQueryParams) SetPageSize(v int32) {
	o.PageSize = &v
}

// GetSortKey returns the SortKey field value if set, zero value otherwise.
func (o *PriceQueryParams) GetSortKey() string {
	if o == nil || IsNil(o.SortKey) {
		var ret string
		return ret
	}
	return *o.SortKey
}

// GetSortKeyOk returns a tuple with the SortKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceQueryParams) GetSortKeyOk() (*string, bool) {
	if o == nil || IsNil(o.SortKey) {
		return nil, false
	}
	return o.SortKey, true
}

// HasSortKey returns a boolean if a field has been set.
func (o *PriceQueryParams) HasSortKey() bool {
	if o != nil && !IsNil(o.SortKey) {
		return true
	}

	return false
}

// SetSortKey gets a reference to the given string and assigns it to the SortKey field.
func (o *PriceQueryParams) SetSortKey(v string) {
	o.SortKey = &v
}

// GetSortDescending returns the SortDescending field value if set, zero value otherwise.
func (o *PriceQueryParams) GetSortDescending() bool {
	if o == nil || IsNil(o.SortDescending) {
		var ret bool
		return ret
	}
	return *o.SortDescending
}

// GetSortDescendingOk returns a tuple with the SortDescending field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceQueryParams) GetSortDescendingOk() (*bool, bool) {
	if o == nil || IsNil(o.SortDescending) {
		return nil, false
	}
	return o.SortDescending, true
}

// HasSortDescending returns a boolean if a field has been set.
func (o *PriceQueryParams) HasSortDescending() bool {
	if o != nil && !IsNil(o.SortDescending) {
		return true
	}

	return false
}

// SetSortDescending gets a reference to the given bool and assigns it to the SortDescending field.
func (o *PriceQueryParams) SetSortDescending(v bool) {
	o.SortDescending = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PriceQueryParams) GetCreatedAt() DateTimeFilter {
	if o == nil || IsNil(o.CreatedAt.Get()) {
		var ret DateTimeFilter
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PriceQueryParams) GetCreatedAtOk() (*DateTimeFilter, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *PriceQueryParams) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableDateTimeFilter and assigns it to the CreatedAt field.
func (o *PriceQueryParams) SetCreatedAt(v DateTimeFilter) {
	o.CreatedAt.Set(&v)
}
// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *PriceQueryParams) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *PriceQueryParams) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PriceQueryParams) GetUpdatedAt() DateTimeFilter {
	if o == nil || IsNil(o.UpdatedAt.Get()) {
		var ret DateTimeFilter
		return ret
	}
	return *o.UpdatedAt.Get()
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PriceQueryParams) GetUpdatedAtOk() (*DateTimeFilter, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedAt.Get(), o.UpdatedAt.IsSet()
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *PriceQueryParams) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt.IsSet() {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given NullableDateTimeFilter and assigns it to the UpdatedAt field.
func (o *PriceQueryParams) SetUpdatedAt(v DateTimeFilter) {
	o.UpdatedAt.Set(&v)
}
// SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil
func (o *PriceQueryParams) SetUpdatedAtNil() {
	o.UpdatedAt.Set(nil)
}

// UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
func (o *PriceQueryParams) UnsetUpdatedAt() {
	o.UpdatedAt.Unset()
}

// GetExpand returns the Expand field value if set, zero value otherwise.
func (o *PriceQueryParams) GetExpand() []string {
	if o == nil || IsNil(o.Expand) {
		var ret []string
		return ret
	}
	return o.Expand
}

// GetExpandOk returns a tuple with the Expand field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceQueryParams) GetExpandOk() ([]string, bool) {
	if o == nil || IsNil(o.Expand) {
		return nil, false
	}
	return o.Expand, true
}

// HasExpand returns a boolean if a field has been set.
func (o *PriceQueryParams) HasExpand() bool {
	if o != nil && !IsNil(o.Expand) {
		return true
	}

	return false
}

// SetExpand gets a reference to the given []string and assigns it to the Expand field.
func (o *PriceQueryParams) SetExpand(v []string) {
	o.Expand = v
}

// GetProductId returns the ProductId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PriceQueryParams) GetProductId() string {
	if o == nil || IsNil(o.ProductId.Get()) {
		var ret string
		return ret
	}
	return *o.ProductId.Get()
}

// GetProductIdOk returns a tuple with the ProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PriceQueryParams) GetProductIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProductId.Get(), o.ProductId.IsSet()
}

// HasProductId returns a boolean if a field has been set.
func (o *PriceQueryParams) HasProductId() bool {
	if o != nil && o.ProductId.IsSet() {
		return true
	}

	return false
}

// SetProductId gets a reference to the given NullableString and assigns it to the ProductId field.
func (o *PriceQueryParams) SetProductId(v string) {
	o.ProductId.Set(&v)
}
// SetProductIdNil sets the value for ProductId to be an explicit nil
func (o *PriceQueryParams) SetProductIdNil() {
	o.ProductId.Set(nil)
}

// UnsetProductId ensures that no value is present for ProductId, not even an explicit nil
func (o *PriceQueryParams) UnsetProductId() {
	o.ProductId.Unset()
}

// GetIsActive returns the IsActive field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PriceQueryParams) GetIsActive() bool {
	if o == nil || IsNil(o.IsActive.Get()) {
		var ret bool
		return ret
	}
	return *o.IsActive.Get()
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PriceQueryParams) GetIsActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsActive.Get(), o.IsActive.IsSet()
}

// HasIsActive returns a boolean if a field has been set.
func (o *PriceQueryParams) HasIsActive() bool {
	if o != nil && o.IsActive.IsSet() {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given NullableBool and assigns it to the IsActive field.
func (o *PriceQueryParams) SetIsActive(v bool) {
	o.IsActive.Set(&v)
}
// SetIsActiveNil sets the value for IsActive to be an explicit nil
func (o *PriceQueryParams) SetIsActiveNil() {
	o.IsActive.Set(nil)
}

// UnsetIsActive ensures that no value is present for IsActive, not even an explicit nil
func (o *PriceQueryParams) UnsetIsActive() {
	o.IsActive.Unset()
}

// GetCurrency returns the Currency field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PriceQueryParams) GetCurrency() CurrencyEnum {
	if o == nil || IsNil(o.Currency.Get()) {
		var ret CurrencyEnum
		return ret
	}
	return *o.Currency.Get()
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PriceQueryParams) GetCurrencyOk() (*CurrencyEnum, bool) {
	if o == nil {
		return nil, false
	}
	return o.Currency.Get(), o.Currency.IsSet()
}

// HasCurrency returns a boolean if a field has been set.
func (o *PriceQueryParams) HasCurrency() bool {
	if o != nil && o.Currency.IsSet() {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given NullableCurrencyEnum and assigns it to the Currency field.
func (o *PriceQueryParams) SetCurrency(v CurrencyEnum) {
	o.Currency.Set(&v)
}
// SetCurrencyNil sets the value for Currency to be an explicit nil
func (o *PriceQueryParams) SetCurrencyNil() {
	o.Currency.Set(nil)
}

// UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
func (o *PriceQueryParams) UnsetCurrency() {
	o.Currency.Unset()
}

// GetPriceType returns the PriceType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PriceQueryParams) GetPriceType() PriceTypeEnum {
	if o == nil || IsNil(o.PriceType.Get()) {
		var ret PriceTypeEnum
		return ret
	}
	return *o.PriceType.Get()
}

// GetPriceTypeOk returns a tuple with the PriceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PriceQueryParams) GetPriceTypeOk() (*PriceTypeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return o.PriceType.Get(), o.PriceType.IsSet()
}

// HasPriceType returns a boolean if a field has been set.
func (o *PriceQueryParams) HasPriceType() bool {
	if o != nil && o.PriceType.IsSet() {
		return true
	}

	return false
}

// SetPriceType gets a reference to the given NullablePriceTypeEnum and assigns it to the PriceType field.
func (o *PriceQueryParams) SetPriceType(v PriceTypeEnum) {
	o.PriceType.Set(&v)
}
// SetPriceTypeNil sets the value for PriceType to be an explicit nil
func (o *PriceQueryParams) SetPriceTypeNil() {
	o.PriceType.Set(nil)
}

// UnsetPriceType ensures that no value is present for PriceType, not even an explicit nil
func (o *PriceQueryParams) UnsetPriceType() {
	o.PriceType.Unset()
}

func (o PriceQueryParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PriceQueryParams) ToMap() (map[string]interface{}, error) {
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
	if o.ProductId.IsSet() {
		toSerialize["product_id"] = o.ProductId.Get()
	}
	if o.IsActive.IsSet() {
		toSerialize["is_active"] = o.IsActive.Get()
	}
	if o.Currency.IsSet() {
		toSerialize["currency"] = o.Currency.Get()
	}
	if o.PriceType.IsSet() {
		toSerialize["price_type"] = o.PriceType.Get()
	}
	return toSerialize, nil
}

type NullablePriceQueryParams struct {
	value *PriceQueryParams
	isSet bool
}

func (v NullablePriceQueryParams) Get() *PriceQueryParams {
	return v.value
}

func (v *NullablePriceQueryParams) Set(val *PriceQueryParams) {
	v.value = val
	v.isSet = true
}

func (v NullablePriceQueryParams) IsSet() bool {
	return v.isSet
}

func (v *NullablePriceQueryParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePriceQueryParams(val *PriceQueryParams) *NullablePriceQueryParams {
	return &NullablePriceQueryParams{value: val, isSet: true}
}

func (v NullablePriceQueryParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePriceQueryParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

