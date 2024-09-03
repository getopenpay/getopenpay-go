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

// checks if the SubscriptionQueryParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionQueryParams{}

// SubscriptionQueryParams struct for SubscriptionQueryParams
type SubscriptionQueryParams struct {
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
	CustomerId NullableString `json:"customer_id,omitempty"`
	ProductId NullableString `json:"product_id,omitempty"`
	PriceId NullableString `json:"price_id,omitempty"`
	Status NullableSubscriptionStatusEnum `json:"status,omitempty"`
	CurrentPeriodStart NullableDateTimeFilter `json:"current_period_start,omitempty"`
	CurrentPeriodEnd NullableDateTimeFilter `json:"current_period_end,omitempty"`
	CouponId NullableString `json:"coupon_id,omitempty"`
	Search NullableSearchFilter `json:"search,omitempty"`
}

// NewSubscriptionQueryParams instantiates a new SubscriptionQueryParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionQueryParams() *SubscriptionQueryParams {
	this := SubscriptionQueryParams{}
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

// NewSubscriptionQueryParamsWithDefaults instantiates a new SubscriptionQueryParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionQueryParamsWithDefaults() *SubscriptionQueryParams {
	this := SubscriptionQueryParams{}
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
func (o *SubscriptionQueryParams) GetPageNumber() int32 {
	if o == nil || IsNil(o.PageNumber) {
		var ret int32
		return ret
	}
	return *o.PageNumber
}

// GetPageNumberOk returns a tuple with the PageNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionQueryParams) GetPageNumberOk() (*int32, bool) {
	if o == nil || IsNil(o.PageNumber) {
		return nil, false
	}
	return o.PageNumber, true
}

// HasPageNumber returns a boolean if a field has been set.
func (o *SubscriptionQueryParams) HasPageNumber() bool {
	if o != nil && !IsNil(o.PageNumber) {
		return true
	}

	return false
}

// SetPageNumber gets a reference to the given int32 and assigns it to the PageNumber field.
func (o *SubscriptionQueryParams) SetPageNumber(v int32) {
	o.PageNumber = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *SubscriptionQueryParams) GetPageSize() int32 {
	if o == nil || IsNil(o.PageSize) {
		var ret int32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionQueryParams) GetPageSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *SubscriptionQueryParams) HasPageSize() bool {
	if o != nil && !IsNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.
func (o *SubscriptionQueryParams) SetPageSize(v int32) {
	o.PageSize = &v
}

// GetSortKey returns the SortKey field value if set, zero value otherwise.
func (o *SubscriptionQueryParams) GetSortKey() string {
	if o == nil || IsNil(o.SortKey) {
		var ret string
		return ret
	}
	return *o.SortKey
}

// GetSortKeyOk returns a tuple with the SortKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionQueryParams) GetSortKeyOk() (*string, bool) {
	if o == nil || IsNil(o.SortKey) {
		return nil, false
	}
	return o.SortKey, true
}

// HasSortKey returns a boolean if a field has been set.
func (o *SubscriptionQueryParams) HasSortKey() bool {
	if o != nil && !IsNil(o.SortKey) {
		return true
	}

	return false
}

// SetSortKey gets a reference to the given string and assigns it to the SortKey field.
func (o *SubscriptionQueryParams) SetSortKey(v string) {
	o.SortKey = &v
}

// GetSortDescending returns the SortDescending field value if set, zero value otherwise.
func (o *SubscriptionQueryParams) GetSortDescending() bool {
	if o == nil || IsNil(o.SortDescending) {
		var ret bool
		return ret
	}
	return *o.SortDescending
}

// GetSortDescendingOk returns a tuple with the SortDescending field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionQueryParams) GetSortDescendingOk() (*bool, bool) {
	if o == nil || IsNil(o.SortDescending) {
		return nil, false
	}
	return o.SortDescending, true
}

// HasSortDescending returns a boolean if a field has been set.
func (o *SubscriptionQueryParams) HasSortDescending() bool {
	if o != nil && !IsNil(o.SortDescending) {
		return true
	}

	return false
}

// SetSortDescending gets a reference to the given bool and assigns it to the SortDescending field.
func (o *SubscriptionQueryParams) SetSortDescending(v bool) {
	o.SortDescending = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SubscriptionQueryParams) GetCreatedAt() DateTimeFilter {
	if o == nil || IsNil(o.CreatedAt.Get()) {
		var ret DateTimeFilter
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubscriptionQueryParams) GetCreatedAtOk() (*DateTimeFilter, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *SubscriptionQueryParams) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableDateTimeFilter and assigns it to the CreatedAt field.
func (o *SubscriptionQueryParams) SetCreatedAt(v DateTimeFilter) {
	o.CreatedAt.Set(&v)
}
// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *SubscriptionQueryParams) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *SubscriptionQueryParams) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SubscriptionQueryParams) GetUpdatedAt() DateTimeFilter {
	if o == nil || IsNil(o.UpdatedAt.Get()) {
		var ret DateTimeFilter
		return ret
	}
	return *o.UpdatedAt.Get()
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubscriptionQueryParams) GetUpdatedAtOk() (*DateTimeFilter, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedAt.Get(), o.UpdatedAt.IsSet()
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *SubscriptionQueryParams) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt.IsSet() {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given NullableDateTimeFilter and assigns it to the UpdatedAt field.
func (o *SubscriptionQueryParams) SetUpdatedAt(v DateTimeFilter) {
	o.UpdatedAt.Set(&v)
}
// SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil
func (o *SubscriptionQueryParams) SetUpdatedAtNil() {
	o.UpdatedAt.Set(nil)
}

// UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
func (o *SubscriptionQueryParams) UnsetUpdatedAt() {
	o.UpdatedAt.Unset()
}

// GetExpand returns the Expand field value if set, zero value otherwise.
func (o *SubscriptionQueryParams) GetExpand() []string {
	if o == nil || IsNil(o.Expand) {
		var ret []string
		return ret
	}
	return o.Expand
}

// GetExpandOk returns a tuple with the Expand field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionQueryParams) GetExpandOk() ([]string, bool) {
	if o == nil || IsNil(o.Expand) {
		return nil, false
	}
	return o.Expand, true
}

// HasExpand returns a boolean if a field has been set.
func (o *SubscriptionQueryParams) HasExpand() bool {
	if o != nil && !IsNil(o.Expand) {
		return true
	}

	return false
}

// SetExpand gets a reference to the given []string and assigns it to the Expand field.
func (o *SubscriptionQueryParams) SetExpand(v []string) {
	o.Expand = v
}

// GetCustomerId returns the CustomerId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SubscriptionQueryParams) GetCustomerId() string {
	if o == nil || IsNil(o.CustomerId.Get()) {
		var ret string
		return ret
	}
	return *o.CustomerId.Get()
}

// GetCustomerIdOk returns a tuple with the CustomerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubscriptionQueryParams) GetCustomerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomerId.Get(), o.CustomerId.IsSet()
}

// HasCustomerId returns a boolean if a field has been set.
func (o *SubscriptionQueryParams) HasCustomerId() bool {
	if o != nil && o.CustomerId.IsSet() {
		return true
	}

	return false
}

// SetCustomerId gets a reference to the given NullableString and assigns it to the CustomerId field.
func (o *SubscriptionQueryParams) SetCustomerId(v string) {
	o.CustomerId.Set(&v)
}
// SetCustomerIdNil sets the value for CustomerId to be an explicit nil
func (o *SubscriptionQueryParams) SetCustomerIdNil() {
	o.CustomerId.Set(nil)
}

// UnsetCustomerId ensures that no value is present for CustomerId, not even an explicit nil
func (o *SubscriptionQueryParams) UnsetCustomerId() {
	o.CustomerId.Unset()
}

// GetProductId returns the ProductId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SubscriptionQueryParams) GetProductId() string {
	if o == nil || IsNil(o.ProductId.Get()) {
		var ret string
		return ret
	}
	return *o.ProductId.Get()
}

// GetProductIdOk returns a tuple with the ProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubscriptionQueryParams) GetProductIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProductId.Get(), o.ProductId.IsSet()
}

// HasProductId returns a boolean if a field has been set.
func (o *SubscriptionQueryParams) HasProductId() bool {
	if o != nil && o.ProductId.IsSet() {
		return true
	}

	return false
}

// SetProductId gets a reference to the given NullableString and assigns it to the ProductId field.
func (o *SubscriptionQueryParams) SetProductId(v string) {
	o.ProductId.Set(&v)
}
// SetProductIdNil sets the value for ProductId to be an explicit nil
func (o *SubscriptionQueryParams) SetProductIdNil() {
	o.ProductId.Set(nil)
}

// UnsetProductId ensures that no value is present for ProductId, not even an explicit nil
func (o *SubscriptionQueryParams) UnsetProductId() {
	o.ProductId.Unset()
}

// GetPriceId returns the PriceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SubscriptionQueryParams) GetPriceId() string {
	if o == nil || IsNil(o.PriceId.Get()) {
		var ret string
		return ret
	}
	return *o.PriceId.Get()
}

// GetPriceIdOk returns a tuple with the PriceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubscriptionQueryParams) GetPriceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PriceId.Get(), o.PriceId.IsSet()
}

// HasPriceId returns a boolean if a field has been set.
func (o *SubscriptionQueryParams) HasPriceId() bool {
	if o != nil && o.PriceId.IsSet() {
		return true
	}

	return false
}

// SetPriceId gets a reference to the given NullableString and assigns it to the PriceId field.
func (o *SubscriptionQueryParams) SetPriceId(v string) {
	o.PriceId.Set(&v)
}
// SetPriceIdNil sets the value for PriceId to be an explicit nil
func (o *SubscriptionQueryParams) SetPriceIdNil() {
	o.PriceId.Set(nil)
}

// UnsetPriceId ensures that no value is present for PriceId, not even an explicit nil
func (o *SubscriptionQueryParams) UnsetPriceId() {
	o.PriceId.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SubscriptionQueryParams) GetStatus() SubscriptionStatusEnum {
	if o == nil || IsNil(o.Status.Get()) {
		var ret SubscriptionStatusEnum
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubscriptionQueryParams) GetStatusOk() (*SubscriptionStatusEnum, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *SubscriptionQueryParams) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableSubscriptionStatusEnum and assigns it to the Status field.
func (o *SubscriptionQueryParams) SetStatus(v SubscriptionStatusEnum) {
	o.Status.Set(&v)
}
// SetStatusNil sets the value for Status to be an explicit nil
func (o *SubscriptionQueryParams) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *SubscriptionQueryParams) UnsetStatus() {
	o.Status.Unset()
}

// GetCurrentPeriodStart returns the CurrentPeriodStart field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SubscriptionQueryParams) GetCurrentPeriodStart() DateTimeFilter {
	if o == nil || IsNil(o.CurrentPeriodStart.Get()) {
		var ret DateTimeFilter
		return ret
	}
	return *o.CurrentPeriodStart.Get()
}

// GetCurrentPeriodStartOk returns a tuple with the CurrentPeriodStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubscriptionQueryParams) GetCurrentPeriodStartOk() (*DateTimeFilter, bool) {
	if o == nil {
		return nil, false
	}
	return o.CurrentPeriodStart.Get(), o.CurrentPeriodStart.IsSet()
}

// HasCurrentPeriodStart returns a boolean if a field has been set.
func (o *SubscriptionQueryParams) HasCurrentPeriodStart() bool {
	if o != nil && o.CurrentPeriodStart.IsSet() {
		return true
	}

	return false
}

// SetCurrentPeriodStart gets a reference to the given NullableDateTimeFilter and assigns it to the CurrentPeriodStart field.
func (o *SubscriptionQueryParams) SetCurrentPeriodStart(v DateTimeFilter) {
	o.CurrentPeriodStart.Set(&v)
}
// SetCurrentPeriodStartNil sets the value for CurrentPeriodStart to be an explicit nil
func (o *SubscriptionQueryParams) SetCurrentPeriodStartNil() {
	o.CurrentPeriodStart.Set(nil)
}

// UnsetCurrentPeriodStart ensures that no value is present for CurrentPeriodStart, not even an explicit nil
func (o *SubscriptionQueryParams) UnsetCurrentPeriodStart() {
	o.CurrentPeriodStart.Unset()
}

// GetCurrentPeriodEnd returns the CurrentPeriodEnd field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SubscriptionQueryParams) GetCurrentPeriodEnd() DateTimeFilter {
	if o == nil || IsNil(o.CurrentPeriodEnd.Get()) {
		var ret DateTimeFilter
		return ret
	}
	return *o.CurrentPeriodEnd.Get()
}

// GetCurrentPeriodEndOk returns a tuple with the CurrentPeriodEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubscriptionQueryParams) GetCurrentPeriodEndOk() (*DateTimeFilter, bool) {
	if o == nil {
		return nil, false
	}
	return o.CurrentPeriodEnd.Get(), o.CurrentPeriodEnd.IsSet()
}

// HasCurrentPeriodEnd returns a boolean if a field has been set.
func (o *SubscriptionQueryParams) HasCurrentPeriodEnd() bool {
	if o != nil && o.CurrentPeriodEnd.IsSet() {
		return true
	}

	return false
}

// SetCurrentPeriodEnd gets a reference to the given NullableDateTimeFilter and assigns it to the CurrentPeriodEnd field.
func (o *SubscriptionQueryParams) SetCurrentPeriodEnd(v DateTimeFilter) {
	o.CurrentPeriodEnd.Set(&v)
}
// SetCurrentPeriodEndNil sets the value for CurrentPeriodEnd to be an explicit nil
func (o *SubscriptionQueryParams) SetCurrentPeriodEndNil() {
	o.CurrentPeriodEnd.Set(nil)
}

// UnsetCurrentPeriodEnd ensures that no value is present for CurrentPeriodEnd, not even an explicit nil
func (o *SubscriptionQueryParams) UnsetCurrentPeriodEnd() {
	o.CurrentPeriodEnd.Unset()
}

// GetCouponId returns the CouponId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SubscriptionQueryParams) GetCouponId() string {
	if o == nil || IsNil(o.CouponId.Get()) {
		var ret string
		return ret
	}
	return *o.CouponId.Get()
}

// GetCouponIdOk returns a tuple with the CouponId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubscriptionQueryParams) GetCouponIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CouponId.Get(), o.CouponId.IsSet()
}

// HasCouponId returns a boolean if a field has been set.
func (o *SubscriptionQueryParams) HasCouponId() bool {
	if o != nil && o.CouponId.IsSet() {
		return true
	}

	return false
}

// SetCouponId gets a reference to the given NullableString and assigns it to the CouponId field.
func (o *SubscriptionQueryParams) SetCouponId(v string) {
	o.CouponId.Set(&v)
}
// SetCouponIdNil sets the value for CouponId to be an explicit nil
func (o *SubscriptionQueryParams) SetCouponIdNil() {
	o.CouponId.Set(nil)
}

// UnsetCouponId ensures that no value is present for CouponId, not even an explicit nil
func (o *SubscriptionQueryParams) UnsetCouponId() {
	o.CouponId.Unset()
}

// GetSearch returns the Search field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SubscriptionQueryParams) GetSearch() SearchFilter {
	if o == nil || IsNil(o.Search.Get()) {
		var ret SearchFilter
		return ret
	}
	return *o.Search.Get()
}

// GetSearchOk returns a tuple with the Search field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubscriptionQueryParams) GetSearchOk() (*SearchFilter, bool) {
	if o == nil {
		return nil, false
	}
	return o.Search.Get(), o.Search.IsSet()
}

// HasSearch returns a boolean if a field has been set.
func (o *SubscriptionQueryParams) HasSearch() bool {
	if o != nil && o.Search.IsSet() {
		return true
	}

	return false
}

// SetSearch gets a reference to the given NullableSearchFilter and assigns it to the Search field.
func (o *SubscriptionQueryParams) SetSearch(v SearchFilter) {
	o.Search.Set(&v)
}
// SetSearchNil sets the value for Search to be an explicit nil
func (o *SubscriptionQueryParams) SetSearchNil() {
	o.Search.Set(nil)
}

// UnsetSearch ensures that no value is present for Search, not even an explicit nil
func (o *SubscriptionQueryParams) UnsetSearch() {
	o.Search.Unset()
}

func (o SubscriptionQueryParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionQueryParams) ToMap() (map[string]interface{}, error) {
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
	if o.CustomerId.IsSet() {
		toSerialize["customer_id"] = o.CustomerId.Get()
	}
	if o.ProductId.IsSet() {
		toSerialize["product_id"] = o.ProductId.Get()
	}
	if o.PriceId.IsSet() {
		toSerialize["price_id"] = o.PriceId.Get()
	}
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
	}
	if o.CurrentPeriodStart.IsSet() {
		toSerialize["current_period_start"] = o.CurrentPeriodStart.Get()
	}
	if o.CurrentPeriodEnd.IsSet() {
		toSerialize["current_period_end"] = o.CurrentPeriodEnd.Get()
	}
	if o.CouponId.IsSet() {
		toSerialize["coupon_id"] = o.CouponId.Get()
	}
	if o.Search.IsSet() {
		toSerialize["search"] = o.Search.Get()
	}
	return toSerialize, nil
}

type NullableSubscriptionQueryParams struct {
	value *SubscriptionQueryParams
	isSet bool
}

func (v NullableSubscriptionQueryParams) Get() *SubscriptionQueryParams {
	return v.value
}

func (v *NullableSubscriptionQueryParams) Set(val *SubscriptionQueryParams) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionQueryParams) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionQueryParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionQueryParams(val *SubscriptionQueryParams) *NullableSubscriptionQueryParams {
	return &NullableSubscriptionQueryParams{value: val, isSet: true}
}

func (v NullableSubscriptionQueryParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionQueryParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


