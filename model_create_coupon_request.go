/*
OpenPay API

super charge your subscription management.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package getopenpay

import (
	"encoding/json"
	"time"
	"fmt"
)

// checks if the CreateCouponRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateCouponRequest{}

// CreateCouponRequest struct for CreateCouponRequest
type CreateCouponRequest struct {
	// Name of the coupon displayed to customers on, for instance invoices, or receipts.
	Name string `json:"name"`
	AmountAtomOff NullableInt32 `json:"amount_atom_off,omitempty"`
	PercentOff NullableInt32 `json:"percent_off,omitempty"`
	Currency NullableCurrencyEnum `json:"currency,omitempty"`
	Duration *CouponDuration `json:"duration,omitempty"`
	DurationInMonths NullableInt32 `json:"duration_in_months,omitempty"`
	ProductIds []string `json:"product_ids,omitempty"`
	ProductFamilyIds []string `json:"product_family_ids,omitempty"`
	MaxRedemptions NullableInt32 `json:"max_redemptions,omitempty"`
	RedeemBy NullableTime `json:"redeem_by,omitempty"`
	// Whether the coupon is available to be redeemed.
	IsActive *bool `json:"is_active,omitempty"`
}

type _CreateCouponRequest CreateCouponRequest

// NewCreateCouponRequest instantiates a new CreateCouponRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateCouponRequest(name string) *CreateCouponRequest {
	this := CreateCouponRequest{}
	this.Name = name
	var duration CouponDuration = COUPONDURATION_ONCE
	this.Duration = &duration
	var isActive bool = true
	this.IsActive = &isActive
	return &this
}

// NewCreateCouponRequestWithDefaults instantiates a new CreateCouponRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateCouponRequestWithDefaults() *CreateCouponRequest {
	this := CreateCouponRequest{}
	var duration CouponDuration = COUPONDURATION_ONCE
	this.Duration = &duration
	var isActive bool = true
	this.IsActive = &isActive
	return &this
}

// GetName returns the Name field value
func (o *CreateCouponRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateCouponRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateCouponRequest) SetName(v string) {
	o.Name = v
}

// GetAmountAtomOff returns the AmountAtomOff field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateCouponRequest) GetAmountAtomOff() int32 {
	if o == nil || IsNil(o.AmountAtomOff.Get()) {
		var ret int32
		return ret
	}
	return *o.AmountAtomOff.Get()
}

// GetAmountAtomOffOk returns a tuple with the AmountAtomOff field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateCouponRequest) GetAmountAtomOffOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AmountAtomOff.Get(), o.AmountAtomOff.IsSet()
}

// HasAmountAtomOff returns a boolean if a field has been set.
func (o *CreateCouponRequest) HasAmountAtomOff() bool {
	if o != nil && o.AmountAtomOff.IsSet() {
		return true
	}

	return false
}

// SetAmountAtomOff gets a reference to the given NullableInt32 and assigns it to the AmountAtomOff field.
func (o *CreateCouponRequest) SetAmountAtomOff(v int32) {
	o.AmountAtomOff.Set(&v)
}
// SetAmountAtomOffNil sets the value for AmountAtomOff to be an explicit nil
func (o *CreateCouponRequest) SetAmountAtomOffNil() {
	o.AmountAtomOff.Set(nil)
}

// UnsetAmountAtomOff ensures that no value is present for AmountAtomOff, not even an explicit nil
func (o *CreateCouponRequest) UnsetAmountAtomOff() {
	o.AmountAtomOff.Unset()
}

// GetPercentOff returns the PercentOff field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateCouponRequest) GetPercentOff() int32 {
	if o == nil || IsNil(o.PercentOff.Get()) {
		var ret int32
		return ret
	}
	return *o.PercentOff.Get()
}

// GetPercentOffOk returns a tuple with the PercentOff field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateCouponRequest) GetPercentOffOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PercentOff.Get(), o.PercentOff.IsSet()
}

// HasPercentOff returns a boolean if a field has been set.
func (o *CreateCouponRequest) HasPercentOff() bool {
	if o != nil && o.PercentOff.IsSet() {
		return true
	}

	return false
}

// SetPercentOff gets a reference to the given NullableInt32 and assigns it to the PercentOff field.
func (o *CreateCouponRequest) SetPercentOff(v int32) {
	o.PercentOff.Set(&v)
}
// SetPercentOffNil sets the value for PercentOff to be an explicit nil
func (o *CreateCouponRequest) SetPercentOffNil() {
	o.PercentOff.Set(nil)
}

// UnsetPercentOff ensures that no value is present for PercentOff, not even an explicit nil
func (o *CreateCouponRequest) UnsetPercentOff() {
	o.PercentOff.Unset()
}

// GetCurrency returns the Currency field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateCouponRequest) GetCurrency() CurrencyEnum {
	if o == nil || IsNil(o.Currency.Get()) {
		var ret CurrencyEnum
		return ret
	}
	return *o.Currency.Get()
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateCouponRequest) GetCurrencyOk() (*CurrencyEnum, bool) {
	if o == nil {
		return nil, false
	}
	return o.Currency.Get(), o.Currency.IsSet()
}

// HasCurrency returns a boolean if a field has been set.
func (o *CreateCouponRequest) HasCurrency() bool {
	if o != nil && o.Currency.IsSet() {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given NullableCurrencyEnum and assigns it to the Currency field.
func (o *CreateCouponRequest) SetCurrency(v CurrencyEnum) {
	o.Currency.Set(&v)
}
// SetCurrencyNil sets the value for Currency to be an explicit nil
func (o *CreateCouponRequest) SetCurrencyNil() {
	o.Currency.Set(nil)
}

// UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
func (o *CreateCouponRequest) UnsetCurrency() {
	o.Currency.Unset()
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *CreateCouponRequest) GetDuration() CouponDuration {
	if o == nil || IsNil(o.Duration) {
		var ret CouponDuration
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCouponRequest) GetDurationOk() (*CouponDuration, bool) {
	if o == nil || IsNil(o.Duration) {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *CreateCouponRequest) HasDuration() bool {
	if o != nil && !IsNil(o.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given CouponDuration and assigns it to the Duration field.
func (o *CreateCouponRequest) SetDuration(v CouponDuration) {
	o.Duration = &v
}

// GetDurationInMonths returns the DurationInMonths field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateCouponRequest) GetDurationInMonths() int32 {
	if o == nil || IsNil(o.DurationInMonths.Get()) {
		var ret int32
		return ret
	}
	return *o.DurationInMonths.Get()
}

// GetDurationInMonthsOk returns a tuple with the DurationInMonths field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateCouponRequest) GetDurationInMonthsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.DurationInMonths.Get(), o.DurationInMonths.IsSet()
}

// HasDurationInMonths returns a boolean if a field has been set.
func (o *CreateCouponRequest) HasDurationInMonths() bool {
	if o != nil && o.DurationInMonths.IsSet() {
		return true
	}

	return false
}

// SetDurationInMonths gets a reference to the given NullableInt32 and assigns it to the DurationInMonths field.
func (o *CreateCouponRequest) SetDurationInMonths(v int32) {
	o.DurationInMonths.Set(&v)
}
// SetDurationInMonthsNil sets the value for DurationInMonths to be an explicit nil
func (o *CreateCouponRequest) SetDurationInMonthsNil() {
	o.DurationInMonths.Set(nil)
}

// UnsetDurationInMonths ensures that no value is present for DurationInMonths, not even an explicit nil
func (o *CreateCouponRequest) UnsetDurationInMonths() {
	o.DurationInMonths.Unset()
}

// GetProductIds returns the ProductIds field value if set, zero value otherwise.
func (o *CreateCouponRequest) GetProductIds() []string {
	if o == nil || IsNil(o.ProductIds) {
		var ret []string
		return ret
	}
	return o.ProductIds
}

// GetProductIdsOk returns a tuple with the ProductIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCouponRequest) GetProductIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.ProductIds) {
		return nil, false
	}
	return o.ProductIds, true
}

// HasProductIds returns a boolean if a field has been set.
func (o *CreateCouponRequest) HasProductIds() bool {
	if o != nil && !IsNil(o.ProductIds) {
		return true
	}

	return false
}

// SetProductIds gets a reference to the given []string and assigns it to the ProductIds field.
func (o *CreateCouponRequest) SetProductIds(v []string) {
	o.ProductIds = v
}

// GetProductFamilyIds returns the ProductFamilyIds field value if set, zero value otherwise.
func (o *CreateCouponRequest) GetProductFamilyIds() []string {
	if o == nil || IsNil(o.ProductFamilyIds) {
		var ret []string
		return ret
	}
	return o.ProductFamilyIds
}

// GetProductFamilyIdsOk returns a tuple with the ProductFamilyIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCouponRequest) GetProductFamilyIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.ProductFamilyIds) {
		return nil, false
	}
	return o.ProductFamilyIds, true
}

// HasProductFamilyIds returns a boolean if a field has been set.
func (o *CreateCouponRequest) HasProductFamilyIds() bool {
	if o != nil && !IsNil(o.ProductFamilyIds) {
		return true
	}

	return false
}

// SetProductFamilyIds gets a reference to the given []string and assigns it to the ProductFamilyIds field.
func (o *CreateCouponRequest) SetProductFamilyIds(v []string) {
	o.ProductFamilyIds = v
}

// GetMaxRedemptions returns the MaxRedemptions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateCouponRequest) GetMaxRedemptions() int32 {
	if o == nil || IsNil(o.MaxRedemptions.Get()) {
		var ret int32
		return ret
	}
	return *o.MaxRedemptions.Get()
}

// GetMaxRedemptionsOk returns a tuple with the MaxRedemptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateCouponRequest) GetMaxRedemptionsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxRedemptions.Get(), o.MaxRedemptions.IsSet()
}

// HasMaxRedemptions returns a boolean if a field has been set.
func (o *CreateCouponRequest) HasMaxRedemptions() bool {
	if o != nil && o.MaxRedemptions.IsSet() {
		return true
	}

	return false
}

// SetMaxRedemptions gets a reference to the given NullableInt32 and assigns it to the MaxRedemptions field.
func (o *CreateCouponRequest) SetMaxRedemptions(v int32) {
	o.MaxRedemptions.Set(&v)
}
// SetMaxRedemptionsNil sets the value for MaxRedemptions to be an explicit nil
func (o *CreateCouponRequest) SetMaxRedemptionsNil() {
	o.MaxRedemptions.Set(nil)
}

// UnsetMaxRedemptions ensures that no value is present for MaxRedemptions, not even an explicit nil
func (o *CreateCouponRequest) UnsetMaxRedemptions() {
	o.MaxRedemptions.Unset()
}

// GetRedeemBy returns the RedeemBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateCouponRequest) GetRedeemBy() time.Time {
	if o == nil || IsNil(o.RedeemBy.Get()) {
		var ret time.Time
		return ret
	}
	return *o.RedeemBy.Get()
}

// GetRedeemByOk returns a tuple with the RedeemBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateCouponRequest) GetRedeemByOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.RedeemBy.Get(), o.RedeemBy.IsSet()
}

// HasRedeemBy returns a boolean if a field has been set.
func (o *CreateCouponRequest) HasRedeemBy() bool {
	if o != nil && o.RedeemBy.IsSet() {
		return true
	}

	return false
}

// SetRedeemBy gets a reference to the given NullableTime and assigns it to the RedeemBy field.
func (o *CreateCouponRequest) SetRedeemBy(v time.Time) {
	o.RedeemBy.Set(&v)
}
// SetRedeemByNil sets the value for RedeemBy to be an explicit nil
func (o *CreateCouponRequest) SetRedeemByNil() {
	o.RedeemBy.Set(nil)
}

// UnsetRedeemBy ensures that no value is present for RedeemBy, not even an explicit nil
func (o *CreateCouponRequest) UnsetRedeemBy() {
	o.RedeemBy.Unset()
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *CreateCouponRequest) GetIsActive() bool {
	if o == nil || IsNil(o.IsActive) {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCouponRequest) GetIsActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.IsActive) {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *CreateCouponRequest) HasIsActive() bool {
	if o != nil && !IsNil(o.IsActive) {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *CreateCouponRequest) SetIsActive(v bool) {
	o.IsActive = &v
}

func (o CreateCouponRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateCouponRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if o.AmountAtomOff.IsSet() {
		toSerialize["amount_atom_off"] = o.AmountAtomOff.Get()
	}
	if o.PercentOff.IsSet() {
		toSerialize["percent_off"] = o.PercentOff.Get()
	}
	if o.Currency.IsSet() {
		toSerialize["currency"] = o.Currency.Get()
	}
	if !IsNil(o.Duration) {
		toSerialize["duration"] = o.Duration
	}
	if o.DurationInMonths.IsSet() {
		toSerialize["duration_in_months"] = o.DurationInMonths.Get()
	}
	if !IsNil(o.ProductIds) {
		toSerialize["product_ids"] = o.ProductIds
	}
	if !IsNil(o.ProductFamilyIds) {
		toSerialize["product_family_ids"] = o.ProductFamilyIds
	}
	if o.MaxRedemptions.IsSet() {
		toSerialize["max_redemptions"] = o.MaxRedemptions.Get()
	}
	if o.RedeemBy.IsSet() {
		toSerialize["redeem_by"] = o.RedeemBy.Get()
	}
	if !IsNil(o.IsActive) {
		toSerialize["is_active"] = o.IsActive
	}
	return toSerialize, nil
}

func (o *CreateCouponRequest) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varCreateCouponRequest := _CreateCouponRequest{}

	err = json.Unmarshal(bytes, &varCreateCouponRequest)

	if err != nil {
		return err
	}

	*o = CreateCouponRequest(varCreateCouponRequest)

	return err
}

type NullableCreateCouponRequest struct {
	value *CreateCouponRequest
	isSet bool
}

func (v NullableCreateCouponRequest) Get() *CreateCouponRequest {
	return v.value
}

func (v *NullableCreateCouponRequest) Set(val *CreateCouponRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateCouponRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateCouponRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateCouponRequest(val *CreateCouponRequest) *NullableCreateCouponRequest {
	return &NullableCreateCouponRequest{value: val, isSet: true}
}

func (v NullableCreateCouponRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateCouponRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

