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

// checks if the CreateCheckoutSessionRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateCheckoutSessionRequest{}

// CreateCheckoutSessionRequest struct for CreateCheckoutSessionRequest
type CreateCheckoutSessionRequest struct {
	ClientReferenceId NullableString `json:"client_reference_id,omitempty"`
	CouponId NullableString `json:"coupon_id,omitempty"`
	Currency NullableCurrencyEnum `json:"currency,omitempty"`
	CustomerId NullableString `json:"customer_id,omitempty"`
	CustomerEmail NullableString `json:"customer_email,omitempty"`
	LineItems []CreateCheckoutLineItem `json:"line_items,omitempty"`
	Mode CheckoutMode `json:"mode"`
	ReturnUrl NullableString `json:"return_url,omitempty"`
	SuccessUrl NullableString `json:"success_url,omitempty"`
	TrialEnd NullableTime `json:"trial_end,omitempty"`
	TrialPeriodDays NullableInt32 `json:"trial_period_days,omitempty"`
	TrialFromPrice NullableBool `json:"trial_from_price,omitempty"`
}

type _CreateCheckoutSessionRequest CreateCheckoutSessionRequest

// NewCreateCheckoutSessionRequest instantiates a new CreateCheckoutSessionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateCheckoutSessionRequest(mode CheckoutMode) *CreateCheckoutSessionRequest {
	this := CreateCheckoutSessionRequest{}
	this.Mode = mode
	return &this
}

// NewCreateCheckoutSessionRequestWithDefaults instantiates a new CreateCheckoutSessionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateCheckoutSessionRequestWithDefaults() *CreateCheckoutSessionRequest {
	this := CreateCheckoutSessionRequest{}
	return &this
}

// GetClientReferenceId returns the ClientReferenceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateCheckoutSessionRequest) GetClientReferenceId() string {
	if o == nil || IsNil(o.ClientReferenceId.Get()) {
		var ret string
		return ret
	}
	return *o.ClientReferenceId.Get()
}

// GetClientReferenceIdOk returns a tuple with the ClientReferenceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateCheckoutSessionRequest) GetClientReferenceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClientReferenceId.Get(), o.ClientReferenceId.IsSet()
}

// HasClientReferenceId returns a boolean if a field has been set.
func (o *CreateCheckoutSessionRequest) HasClientReferenceId() bool {
	if o != nil && o.ClientReferenceId.IsSet() {
		return true
	}

	return false
}

// SetClientReferenceId gets a reference to the given NullableString and assigns it to the ClientReferenceId field.
func (o *CreateCheckoutSessionRequest) SetClientReferenceId(v string) {
	o.ClientReferenceId.Set(&v)
}
// SetClientReferenceIdNil sets the value for ClientReferenceId to be an explicit nil
func (o *CreateCheckoutSessionRequest) SetClientReferenceIdNil() {
	o.ClientReferenceId.Set(nil)
}

// UnsetClientReferenceId ensures that no value is present for ClientReferenceId, not even an explicit nil
func (o *CreateCheckoutSessionRequest) UnsetClientReferenceId() {
	o.ClientReferenceId.Unset()
}

// GetCouponId returns the CouponId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateCheckoutSessionRequest) GetCouponId() string {
	if o == nil || IsNil(o.CouponId.Get()) {
		var ret string
		return ret
	}
	return *o.CouponId.Get()
}

// GetCouponIdOk returns a tuple with the CouponId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateCheckoutSessionRequest) GetCouponIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CouponId.Get(), o.CouponId.IsSet()
}

// HasCouponId returns a boolean if a field has been set.
func (o *CreateCheckoutSessionRequest) HasCouponId() bool {
	if o != nil && o.CouponId.IsSet() {
		return true
	}

	return false
}

// SetCouponId gets a reference to the given NullableString and assigns it to the CouponId field.
func (o *CreateCheckoutSessionRequest) SetCouponId(v string) {
	o.CouponId.Set(&v)
}
// SetCouponIdNil sets the value for CouponId to be an explicit nil
func (o *CreateCheckoutSessionRequest) SetCouponIdNil() {
	o.CouponId.Set(nil)
}

// UnsetCouponId ensures that no value is present for CouponId, not even an explicit nil
func (o *CreateCheckoutSessionRequest) UnsetCouponId() {
	o.CouponId.Unset()
}

// GetCurrency returns the Currency field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateCheckoutSessionRequest) GetCurrency() CurrencyEnum {
	if o == nil || IsNil(o.Currency.Get()) {
		var ret CurrencyEnum
		return ret
	}
	return *o.Currency.Get()
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateCheckoutSessionRequest) GetCurrencyOk() (*CurrencyEnum, bool) {
	if o == nil {
		return nil, false
	}
	return o.Currency.Get(), o.Currency.IsSet()
}

// HasCurrency returns a boolean if a field has been set.
func (o *CreateCheckoutSessionRequest) HasCurrency() bool {
	if o != nil && o.Currency.IsSet() {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given NullableCurrencyEnum and assigns it to the Currency field.
func (o *CreateCheckoutSessionRequest) SetCurrency(v CurrencyEnum) {
	o.Currency.Set(&v)
}
// SetCurrencyNil sets the value for Currency to be an explicit nil
func (o *CreateCheckoutSessionRequest) SetCurrencyNil() {
	o.Currency.Set(nil)
}

// UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
func (o *CreateCheckoutSessionRequest) UnsetCurrency() {
	o.Currency.Unset()
}

// GetCustomerId returns the CustomerId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateCheckoutSessionRequest) GetCustomerId() string {
	if o == nil || IsNil(o.CustomerId.Get()) {
		var ret string
		return ret
	}
	return *o.CustomerId.Get()
}

// GetCustomerIdOk returns a tuple with the CustomerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateCheckoutSessionRequest) GetCustomerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomerId.Get(), o.CustomerId.IsSet()
}

// HasCustomerId returns a boolean if a field has been set.
func (o *CreateCheckoutSessionRequest) HasCustomerId() bool {
	if o != nil && o.CustomerId.IsSet() {
		return true
	}

	return false
}

// SetCustomerId gets a reference to the given NullableString and assigns it to the CustomerId field.
func (o *CreateCheckoutSessionRequest) SetCustomerId(v string) {
	o.CustomerId.Set(&v)
}
// SetCustomerIdNil sets the value for CustomerId to be an explicit nil
func (o *CreateCheckoutSessionRequest) SetCustomerIdNil() {
	o.CustomerId.Set(nil)
}

// UnsetCustomerId ensures that no value is present for CustomerId, not even an explicit nil
func (o *CreateCheckoutSessionRequest) UnsetCustomerId() {
	o.CustomerId.Unset()
}

// GetCustomerEmail returns the CustomerEmail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateCheckoutSessionRequest) GetCustomerEmail() string {
	if o == nil || IsNil(o.CustomerEmail.Get()) {
		var ret string
		return ret
	}
	return *o.CustomerEmail.Get()
}

// GetCustomerEmailOk returns a tuple with the CustomerEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateCheckoutSessionRequest) GetCustomerEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomerEmail.Get(), o.CustomerEmail.IsSet()
}

// HasCustomerEmail returns a boolean if a field has been set.
func (o *CreateCheckoutSessionRequest) HasCustomerEmail() bool {
	if o != nil && o.CustomerEmail.IsSet() {
		return true
	}

	return false
}

// SetCustomerEmail gets a reference to the given NullableString and assigns it to the CustomerEmail field.
func (o *CreateCheckoutSessionRequest) SetCustomerEmail(v string) {
	o.CustomerEmail.Set(&v)
}
// SetCustomerEmailNil sets the value for CustomerEmail to be an explicit nil
func (o *CreateCheckoutSessionRequest) SetCustomerEmailNil() {
	o.CustomerEmail.Set(nil)
}

// UnsetCustomerEmail ensures that no value is present for CustomerEmail, not even an explicit nil
func (o *CreateCheckoutSessionRequest) UnsetCustomerEmail() {
	o.CustomerEmail.Unset()
}

// GetLineItems returns the LineItems field value if set, zero value otherwise.
func (o *CreateCheckoutSessionRequest) GetLineItems() []CreateCheckoutLineItem {
	if o == nil || IsNil(o.LineItems) {
		var ret []CreateCheckoutLineItem
		return ret
	}
	return o.LineItems
}

// GetLineItemsOk returns a tuple with the LineItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCheckoutSessionRequest) GetLineItemsOk() ([]CreateCheckoutLineItem, bool) {
	if o == nil || IsNil(o.LineItems) {
		return nil, false
	}
	return o.LineItems, true
}

// HasLineItems returns a boolean if a field has been set.
func (o *CreateCheckoutSessionRequest) HasLineItems() bool {
	if o != nil && !IsNil(o.LineItems) {
		return true
	}

	return false
}

// SetLineItems gets a reference to the given []CreateCheckoutLineItem and assigns it to the LineItems field.
func (o *CreateCheckoutSessionRequest) SetLineItems(v []CreateCheckoutLineItem) {
	o.LineItems = v
}

// GetMode returns the Mode field value
func (o *CreateCheckoutSessionRequest) GetMode() CheckoutMode {
	if o == nil {
		var ret CheckoutMode
		return ret
	}

	return o.Mode
}

// GetModeOk returns a tuple with the Mode field value
// and a boolean to check if the value has been set.
func (o *CreateCheckoutSessionRequest) GetModeOk() (*CheckoutMode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mode, true
}

// SetMode sets field value
func (o *CreateCheckoutSessionRequest) SetMode(v CheckoutMode) {
	o.Mode = v
}

// GetReturnUrl returns the ReturnUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateCheckoutSessionRequest) GetReturnUrl() string {
	if o == nil || IsNil(o.ReturnUrl.Get()) {
		var ret string
		return ret
	}
	return *o.ReturnUrl.Get()
}

// GetReturnUrlOk returns a tuple with the ReturnUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateCheckoutSessionRequest) GetReturnUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReturnUrl.Get(), o.ReturnUrl.IsSet()
}

// HasReturnUrl returns a boolean if a field has been set.
func (o *CreateCheckoutSessionRequest) HasReturnUrl() bool {
	if o != nil && o.ReturnUrl.IsSet() {
		return true
	}

	return false
}

// SetReturnUrl gets a reference to the given NullableString and assigns it to the ReturnUrl field.
func (o *CreateCheckoutSessionRequest) SetReturnUrl(v string) {
	o.ReturnUrl.Set(&v)
}
// SetReturnUrlNil sets the value for ReturnUrl to be an explicit nil
func (o *CreateCheckoutSessionRequest) SetReturnUrlNil() {
	o.ReturnUrl.Set(nil)
}

// UnsetReturnUrl ensures that no value is present for ReturnUrl, not even an explicit nil
func (o *CreateCheckoutSessionRequest) UnsetReturnUrl() {
	o.ReturnUrl.Unset()
}

// GetSuccessUrl returns the SuccessUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateCheckoutSessionRequest) GetSuccessUrl() string {
	if o == nil || IsNil(o.SuccessUrl.Get()) {
		var ret string
		return ret
	}
	return *o.SuccessUrl.Get()
}

// GetSuccessUrlOk returns a tuple with the SuccessUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateCheckoutSessionRequest) GetSuccessUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SuccessUrl.Get(), o.SuccessUrl.IsSet()
}

// HasSuccessUrl returns a boolean if a field has been set.
func (o *CreateCheckoutSessionRequest) HasSuccessUrl() bool {
	if o != nil && o.SuccessUrl.IsSet() {
		return true
	}

	return false
}

// SetSuccessUrl gets a reference to the given NullableString and assigns it to the SuccessUrl field.
func (o *CreateCheckoutSessionRequest) SetSuccessUrl(v string) {
	o.SuccessUrl.Set(&v)
}
// SetSuccessUrlNil sets the value for SuccessUrl to be an explicit nil
func (o *CreateCheckoutSessionRequest) SetSuccessUrlNil() {
	o.SuccessUrl.Set(nil)
}

// UnsetSuccessUrl ensures that no value is present for SuccessUrl, not even an explicit nil
func (o *CreateCheckoutSessionRequest) UnsetSuccessUrl() {
	o.SuccessUrl.Unset()
}

// GetTrialEnd returns the TrialEnd field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateCheckoutSessionRequest) GetTrialEnd() time.Time {
	if o == nil || IsNil(o.TrialEnd.Get()) {
		var ret time.Time
		return ret
	}
	return *o.TrialEnd.Get()
}

// GetTrialEndOk returns a tuple with the TrialEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateCheckoutSessionRequest) GetTrialEndOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.TrialEnd.Get(), o.TrialEnd.IsSet()
}

// HasTrialEnd returns a boolean if a field has been set.
func (o *CreateCheckoutSessionRequest) HasTrialEnd() bool {
	if o != nil && o.TrialEnd.IsSet() {
		return true
	}

	return false
}

// SetTrialEnd gets a reference to the given NullableTime and assigns it to the TrialEnd field.
func (o *CreateCheckoutSessionRequest) SetTrialEnd(v time.Time) {
	o.TrialEnd.Set(&v)
}
// SetTrialEndNil sets the value for TrialEnd to be an explicit nil
func (o *CreateCheckoutSessionRequest) SetTrialEndNil() {
	o.TrialEnd.Set(nil)
}

// UnsetTrialEnd ensures that no value is present for TrialEnd, not even an explicit nil
func (o *CreateCheckoutSessionRequest) UnsetTrialEnd() {
	o.TrialEnd.Unset()
}

// GetTrialPeriodDays returns the TrialPeriodDays field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateCheckoutSessionRequest) GetTrialPeriodDays() int32 {
	if o == nil || IsNil(o.TrialPeriodDays.Get()) {
		var ret int32
		return ret
	}
	return *o.TrialPeriodDays.Get()
}

// GetTrialPeriodDaysOk returns a tuple with the TrialPeriodDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateCheckoutSessionRequest) GetTrialPeriodDaysOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TrialPeriodDays.Get(), o.TrialPeriodDays.IsSet()
}

// HasTrialPeriodDays returns a boolean if a field has been set.
func (o *CreateCheckoutSessionRequest) HasTrialPeriodDays() bool {
	if o != nil && o.TrialPeriodDays.IsSet() {
		return true
	}

	return false
}

// SetTrialPeriodDays gets a reference to the given NullableInt32 and assigns it to the TrialPeriodDays field.
func (o *CreateCheckoutSessionRequest) SetTrialPeriodDays(v int32) {
	o.TrialPeriodDays.Set(&v)
}
// SetTrialPeriodDaysNil sets the value for TrialPeriodDays to be an explicit nil
func (o *CreateCheckoutSessionRequest) SetTrialPeriodDaysNil() {
	o.TrialPeriodDays.Set(nil)
}

// UnsetTrialPeriodDays ensures that no value is present for TrialPeriodDays, not even an explicit nil
func (o *CreateCheckoutSessionRequest) UnsetTrialPeriodDays() {
	o.TrialPeriodDays.Unset()
}

// GetTrialFromPrice returns the TrialFromPrice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateCheckoutSessionRequest) GetTrialFromPrice() bool {
	if o == nil || IsNil(o.TrialFromPrice.Get()) {
		var ret bool
		return ret
	}
	return *o.TrialFromPrice.Get()
}

// GetTrialFromPriceOk returns a tuple with the TrialFromPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateCheckoutSessionRequest) GetTrialFromPriceOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.TrialFromPrice.Get(), o.TrialFromPrice.IsSet()
}

// HasTrialFromPrice returns a boolean if a field has been set.
func (o *CreateCheckoutSessionRequest) HasTrialFromPrice() bool {
	if o != nil && o.TrialFromPrice.IsSet() {
		return true
	}

	return false
}

// SetTrialFromPrice gets a reference to the given NullableBool and assigns it to the TrialFromPrice field.
func (o *CreateCheckoutSessionRequest) SetTrialFromPrice(v bool) {
	o.TrialFromPrice.Set(&v)
}
// SetTrialFromPriceNil sets the value for TrialFromPrice to be an explicit nil
func (o *CreateCheckoutSessionRequest) SetTrialFromPriceNil() {
	o.TrialFromPrice.Set(nil)
}

// UnsetTrialFromPrice ensures that no value is present for TrialFromPrice, not even an explicit nil
func (o *CreateCheckoutSessionRequest) UnsetTrialFromPrice() {
	o.TrialFromPrice.Unset()
}

func (o CreateCheckoutSessionRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateCheckoutSessionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientReferenceId.IsSet() {
		toSerialize["client_reference_id"] = o.ClientReferenceId.Get()
	}
	if o.CouponId.IsSet() {
		toSerialize["coupon_id"] = o.CouponId.Get()
	}
	if o.Currency.IsSet() {
		toSerialize["currency"] = o.Currency.Get()
	}
	if o.CustomerId.IsSet() {
		toSerialize["customer_id"] = o.CustomerId.Get()
	}
	if o.CustomerEmail.IsSet() {
		toSerialize["customer_email"] = o.CustomerEmail.Get()
	}
	if !IsNil(o.LineItems) {
		toSerialize["line_items"] = o.LineItems
	}
	toSerialize["mode"] = o.Mode
	if o.ReturnUrl.IsSet() {
		toSerialize["return_url"] = o.ReturnUrl.Get()
	}
	if o.SuccessUrl.IsSet() {
		toSerialize["success_url"] = o.SuccessUrl.Get()
	}
	if o.TrialEnd.IsSet() {
		toSerialize["trial_end"] = o.TrialEnd.Get()
	}
	if o.TrialPeriodDays.IsSet() {
		toSerialize["trial_period_days"] = o.TrialPeriodDays.Get()
	}
	if o.TrialFromPrice.IsSet() {
		toSerialize["trial_from_price"] = o.TrialFromPrice.Get()
	}
	return toSerialize, nil
}

func (o *CreateCheckoutSessionRequest) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"mode",
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

	varCreateCheckoutSessionRequest := _CreateCheckoutSessionRequest{}

	err = json.Unmarshal(bytes, &varCreateCheckoutSessionRequest)

	if err != nil {
		return err
	}

	*o = CreateCheckoutSessionRequest(varCreateCheckoutSessionRequest)

	return err
}

type NullableCreateCheckoutSessionRequest struct {
	value *CreateCheckoutSessionRequest
	isSet bool
}

func (v NullableCreateCheckoutSessionRequest) Get() *CreateCheckoutSessionRequest {
	return v.value
}

func (v *NullableCreateCheckoutSessionRequest) Set(val *CreateCheckoutSessionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateCheckoutSessionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateCheckoutSessionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateCheckoutSessionRequest(val *CreateCheckoutSessionRequest) *NullableCreateCheckoutSessionRequest {
	return &NullableCreateCheckoutSessionRequest{value: val, isSet: true}
}

func (v NullableCreateCheckoutSessionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateCheckoutSessionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


