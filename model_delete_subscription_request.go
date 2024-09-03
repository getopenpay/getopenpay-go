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

// checks if the DeleteSubscriptionRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteSubscriptionRequest{}

// DeleteSubscriptionRequest struct for DeleteSubscriptionRequest
type DeleteSubscriptionRequest struct {
	CancellationDetails NullableSubscriptionCancellationDetails `json:"cancellation_details,omitempty"`
	// Mark unpaid invoices as uncollectible
	CancelUnpaidInvoices *bool `json:"cancel_unpaid_invoices,omitempty"`
	// Will generate a proration invoice_item that credits remaining unused time until the subscription period end, also creates invoice_item for un-invoiced metered usage.Setting this to false will not invoice for un-invoiced metered usage.
	Prorate *bool `json:"prorate,omitempty"`
	// Flag to decide whether full refund should be given or not.
	FullRefund *bool `json:"full_refund,omitempty"`
}

// NewDeleteSubscriptionRequest instantiates a new DeleteSubscriptionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteSubscriptionRequest() *DeleteSubscriptionRequest {
	this := DeleteSubscriptionRequest{}
	var cancelUnpaidInvoices bool = false
	this.CancelUnpaidInvoices = &cancelUnpaidInvoices
	var prorate bool = true
	this.Prorate = &prorate
	var fullRefund bool = false
	this.FullRefund = &fullRefund
	return &this
}

// NewDeleteSubscriptionRequestWithDefaults instantiates a new DeleteSubscriptionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteSubscriptionRequestWithDefaults() *DeleteSubscriptionRequest {
	this := DeleteSubscriptionRequest{}
	var cancelUnpaidInvoices bool = false
	this.CancelUnpaidInvoices = &cancelUnpaidInvoices
	var prorate bool = true
	this.Prorate = &prorate
	var fullRefund bool = false
	this.FullRefund = &fullRefund
	return &this
}

// GetCancellationDetails returns the CancellationDetails field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeleteSubscriptionRequest) GetCancellationDetails() SubscriptionCancellationDetails {
	if o == nil || IsNil(o.CancellationDetails.Get()) {
		var ret SubscriptionCancellationDetails
		return ret
	}
	return *o.CancellationDetails.Get()
}

// GetCancellationDetailsOk returns a tuple with the CancellationDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeleteSubscriptionRequest) GetCancellationDetailsOk() (*SubscriptionCancellationDetails, bool) {
	if o == nil {
		return nil, false
	}
	return o.CancellationDetails.Get(), o.CancellationDetails.IsSet()
}

// HasCancellationDetails returns a boolean if a field has been set.
func (o *DeleteSubscriptionRequest) HasCancellationDetails() bool {
	if o != nil && o.CancellationDetails.IsSet() {
		return true
	}

	return false
}

// SetCancellationDetails gets a reference to the given NullableSubscriptionCancellationDetails and assigns it to the CancellationDetails field.
func (o *DeleteSubscriptionRequest) SetCancellationDetails(v SubscriptionCancellationDetails) {
	o.CancellationDetails.Set(&v)
}
// SetCancellationDetailsNil sets the value for CancellationDetails to be an explicit nil
func (o *DeleteSubscriptionRequest) SetCancellationDetailsNil() {
	o.CancellationDetails.Set(nil)
}

// UnsetCancellationDetails ensures that no value is present for CancellationDetails, not even an explicit nil
func (o *DeleteSubscriptionRequest) UnsetCancellationDetails() {
	o.CancellationDetails.Unset()
}

// GetCancelUnpaidInvoices returns the CancelUnpaidInvoices field value if set, zero value otherwise.
func (o *DeleteSubscriptionRequest) GetCancelUnpaidInvoices() bool {
	if o == nil || IsNil(o.CancelUnpaidInvoices) {
		var ret bool
		return ret
	}
	return *o.CancelUnpaidInvoices
}

// GetCancelUnpaidInvoicesOk returns a tuple with the CancelUnpaidInvoices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteSubscriptionRequest) GetCancelUnpaidInvoicesOk() (*bool, bool) {
	if o == nil || IsNil(o.CancelUnpaidInvoices) {
		return nil, false
	}
	return o.CancelUnpaidInvoices, true
}

// HasCancelUnpaidInvoices returns a boolean if a field has been set.
func (o *DeleteSubscriptionRequest) HasCancelUnpaidInvoices() bool {
	if o != nil && !IsNil(o.CancelUnpaidInvoices) {
		return true
	}

	return false
}

// SetCancelUnpaidInvoices gets a reference to the given bool and assigns it to the CancelUnpaidInvoices field.
func (o *DeleteSubscriptionRequest) SetCancelUnpaidInvoices(v bool) {
	o.CancelUnpaidInvoices = &v
}

// GetProrate returns the Prorate field value if set, zero value otherwise.
func (o *DeleteSubscriptionRequest) GetProrate() bool {
	if o == nil || IsNil(o.Prorate) {
		var ret bool
		return ret
	}
	return *o.Prorate
}

// GetProrateOk returns a tuple with the Prorate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteSubscriptionRequest) GetProrateOk() (*bool, bool) {
	if o == nil || IsNil(o.Prorate) {
		return nil, false
	}
	return o.Prorate, true
}

// HasProrate returns a boolean if a field has been set.
func (o *DeleteSubscriptionRequest) HasProrate() bool {
	if o != nil && !IsNil(o.Prorate) {
		return true
	}

	return false
}

// SetProrate gets a reference to the given bool and assigns it to the Prorate field.
func (o *DeleteSubscriptionRequest) SetProrate(v bool) {
	o.Prorate = &v
}

// GetFullRefund returns the FullRefund field value if set, zero value otherwise.
func (o *DeleteSubscriptionRequest) GetFullRefund() bool {
	if o == nil || IsNil(o.FullRefund) {
		var ret bool
		return ret
	}
	return *o.FullRefund
}

// GetFullRefundOk returns a tuple with the FullRefund field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteSubscriptionRequest) GetFullRefundOk() (*bool, bool) {
	if o == nil || IsNil(o.FullRefund) {
		return nil, false
	}
	return o.FullRefund, true
}

// HasFullRefund returns a boolean if a field has been set.
func (o *DeleteSubscriptionRequest) HasFullRefund() bool {
	if o != nil && !IsNil(o.FullRefund) {
		return true
	}

	return false
}

// SetFullRefund gets a reference to the given bool and assigns it to the FullRefund field.
func (o *DeleteSubscriptionRequest) SetFullRefund(v bool) {
	o.FullRefund = &v
}

func (o DeleteSubscriptionRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteSubscriptionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.CancellationDetails.IsSet() {
		toSerialize["cancellation_details"] = o.CancellationDetails.Get()
	}
	if !IsNil(o.CancelUnpaidInvoices) {
		toSerialize["cancel_unpaid_invoices"] = o.CancelUnpaidInvoices
	}
	if !IsNil(o.Prorate) {
		toSerialize["prorate"] = o.Prorate
	}
	if !IsNil(o.FullRefund) {
		toSerialize["full_refund"] = o.FullRefund
	}
	return toSerialize, nil
}

type NullableDeleteSubscriptionRequest struct {
	value *DeleteSubscriptionRequest
	isSet bool
}

func (v NullableDeleteSubscriptionRequest) Get() *DeleteSubscriptionRequest {
	return v.value
}

func (v *NullableDeleteSubscriptionRequest) Set(val *DeleteSubscriptionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteSubscriptionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteSubscriptionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteSubscriptionRequest(val *DeleteSubscriptionRequest) *NullableDeleteSubscriptionRequest {
	return &NullableDeleteSubscriptionRequest{value: val, isSet: true}
}

func (v NullableDeleteSubscriptionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteSubscriptionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

