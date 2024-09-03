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

// checks if the CreateCreditNoteRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateCreditNoteRequest{}

// CreateCreditNoteRequest struct for CreateCreditNoteRequest
type CreateCreditNoteRequest struct {
	// ID of the invoice
	InvoiceId string `json:"invoice_id"`
	Reason NullableCreditNoteReason `json:"reason,omitempty"`
	// The int amount representing the total amount of the credit note.
	TotalAmountAtom int32 `json:"total_amount_atom"`
	// The integer amount representing the amount to credit the customer’s balance, which will be automatically applied to their next invoice.
	CreditAmountAtom *int32 `json:"credit_amount_atom,omitempty"`
	// The integer amount representing the amount to refund. If set, a refund will be created for the charge associated with the invoice.
	RefundAmountAtom *int32 `json:"refund_amount_atom,omitempty"`
	Lines []CreateCreditNoteLine `json:"lines"`
	Currency *CurrencyEnum `json:"currency,omitempty"`
}

type _CreateCreditNoteRequest CreateCreditNoteRequest

// NewCreateCreditNoteRequest instantiates a new CreateCreditNoteRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateCreditNoteRequest(invoiceId string, totalAmountAtom int32, lines []CreateCreditNoteLine) *CreateCreditNoteRequest {
	this := CreateCreditNoteRequest{}
	this.InvoiceId = invoiceId
	this.TotalAmountAtom = totalAmountAtom
	var creditAmountAtom int32 = 0
	this.CreditAmountAtom = &creditAmountAtom
	var refundAmountAtom int32 = 0
	this.RefundAmountAtom = &refundAmountAtom
	this.Lines = lines
	var currency CurrencyEnum = CURRENCYENUM_USD
	this.Currency = &currency
	return &this
}

// NewCreateCreditNoteRequestWithDefaults instantiates a new CreateCreditNoteRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateCreditNoteRequestWithDefaults() *CreateCreditNoteRequest {
	this := CreateCreditNoteRequest{}
	var creditAmountAtom int32 = 0
	this.CreditAmountAtom = &creditAmountAtom
	var refundAmountAtom int32 = 0
	this.RefundAmountAtom = &refundAmountAtom
	var currency CurrencyEnum = CURRENCYENUM_USD
	this.Currency = &currency
	return &this
}

// GetInvoiceId returns the InvoiceId field value
func (o *CreateCreditNoteRequest) GetInvoiceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InvoiceId
}

// GetInvoiceIdOk returns a tuple with the InvoiceId field value
// and a boolean to check if the value has been set.
func (o *CreateCreditNoteRequest) GetInvoiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InvoiceId, true
}

// SetInvoiceId sets field value
func (o *CreateCreditNoteRequest) SetInvoiceId(v string) {
	o.InvoiceId = v
}

// GetReason returns the Reason field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateCreditNoteRequest) GetReason() CreditNoteReason {
	if o == nil || IsNil(o.Reason.Get()) {
		var ret CreditNoteReason
		return ret
	}
	return *o.Reason.Get()
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateCreditNoteRequest) GetReasonOk() (*CreditNoteReason, bool) {
	if o == nil {
		return nil, false
	}
	return o.Reason.Get(), o.Reason.IsSet()
}

// HasReason returns a boolean if a field has been set.
func (o *CreateCreditNoteRequest) HasReason() bool {
	if o != nil && o.Reason.IsSet() {
		return true
	}

	return false
}

// SetReason gets a reference to the given NullableCreditNoteReason and assigns it to the Reason field.
func (o *CreateCreditNoteRequest) SetReason(v CreditNoteReason) {
	o.Reason.Set(&v)
}
// SetReasonNil sets the value for Reason to be an explicit nil
func (o *CreateCreditNoteRequest) SetReasonNil() {
	o.Reason.Set(nil)
}

// UnsetReason ensures that no value is present for Reason, not even an explicit nil
func (o *CreateCreditNoteRequest) UnsetReason() {
	o.Reason.Unset()
}

// GetTotalAmountAtom returns the TotalAmountAtom field value
func (o *CreateCreditNoteRequest) GetTotalAmountAtom() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalAmountAtom
}

// GetTotalAmountAtomOk returns a tuple with the TotalAmountAtom field value
// and a boolean to check if the value has been set.
func (o *CreateCreditNoteRequest) GetTotalAmountAtomOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalAmountAtom, true
}

// SetTotalAmountAtom sets field value
func (o *CreateCreditNoteRequest) SetTotalAmountAtom(v int32) {
	o.TotalAmountAtom = v
}

// GetCreditAmountAtom returns the CreditAmountAtom field value if set, zero value otherwise.
func (o *CreateCreditNoteRequest) GetCreditAmountAtom() int32 {
	if o == nil || IsNil(o.CreditAmountAtom) {
		var ret int32
		return ret
	}
	return *o.CreditAmountAtom
}

// GetCreditAmountAtomOk returns a tuple with the CreditAmountAtom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCreditNoteRequest) GetCreditAmountAtomOk() (*int32, bool) {
	if o == nil || IsNil(o.CreditAmountAtom) {
		return nil, false
	}
	return o.CreditAmountAtom, true
}

// HasCreditAmountAtom returns a boolean if a field has been set.
func (o *CreateCreditNoteRequest) HasCreditAmountAtom() bool {
	if o != nil && !IsNil(o.CreditAmountAtom) {
		return true
	}

	return false
}

// SetCreditAmountAtom gets a reference to the given int32 and assigns it to the CreditAmountAtom field.
func (o *CreateCreditNoteRequest) SetCreditAmountAtom(v int32) {
	o.CreditAmountAtom = &v
}

// GetRefundAmountAtom returns the RefundAmountAtom field value if set, zero value otherwise.
func (o *CreateCreditNoteRequest) GetRefundAmountAtom() int32 {
	if o == nil || IsNil(o.RefundAmountAtom) {
		var ret int32
		return ret
	}
	return *o.RefundAmountAtom
}

// GetRefundAmountAtomOk returns a tuple with the RefundAmountAtom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCreditNoteRequest) GetRefundAmountAtomOk() (*int32, bool) {
	if o == nil || IsNil(o.RefundAmountAtom) {
		return nil, false
	}
	return o.RefundAmountAtom, true
}

// HasRefundAmountAtom returns a boolean if a field has been set.
func (o *CreateCreditNoteRequest) HasRefundAmountAtom() bool {
	if o != nil && !IsNil(o.RefundAmountAtom) {
		return true
	}

	return false
}

// SetRefundAmountAtom gets a reference to the given int32 and assigns it to the RefundAmountAtom field.
func (o *CreateCreditNoteRequest) SetRefundAmountAtom(v int32) {
	o.RefundAmountAtom = &v
}

// GetLines returns the Lines field value
func (o *CreateCreditNoteRequest) GetLines() []CreateCreditNoteLine {
	if o == nil {
		var ret []CreateCreditNoteLine
		return ret
	}

	return o.Lines
}

// GetLinesOk returns a tuple with the Lines field value
// and a boolean to check if the value has been set.
func (o *CreateCreditNoteRequest) GetLinesOk() ([]CreateCreditNoteLine, bool) {
	if o == nil {
		return nil, false
	}
	return o.Lines, true
}

// SetLines sets field value
func (o *CreateCreditNoteRequest) SetLines(v []CreateCreditNoteLine) {
	o.Lines = v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *CreateCreditNoteRequest) GetCurrency() CurrencyEnum {
	if o == nil || IsNil(o.Currency) {
		var ret CurrencyEnum
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCreditNoteRequest) GetCurrencyOk() (*CurrencyEnum, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *CreateCreditNoteRequest) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given CurrencyEnum and assigns it to the Currency field.
func (o *CreateCreditNoteRequest) SetCurrency(v CurrencyEnum) {
	o.Currency = &v
}

func (o CreateCreditNoteRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateCreditNoteRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["invoice_id"] = o.InvoiceId
	if o.Reason.IsSet() {
		toSerialize["reason"] = o.Reason.Get()
	}
	toSerialize["total_amount_atom"] = o.TotalAmountAtom
	if !IsNil(o.CreditAmountAtom) {
		toSerialize["credit_amount_atom"] = o.CreditAmountAtom
	}
	if !IsNil(o.RefundAmountAtom) {
		toSerialize["refund_amount_atom"] = o.RefundAmountAtom
	}
	toSerialize["lines"] = o.Lines
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	return toSerialize, nil
}

func (o *CreateCreditNoteRequest) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"invoice_id",
		"total_amount_atom",
		"lines",
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

	varCreateCreditNoteRequest := _CreateCreditNoteRequest{}

	err = json.Unmarshal(bytes, &varCreateCreditNoteRequest)

	if err != nil {
		return err
	}

	*o = CreateCreditNoteRequest(varCreateCreditNoteRequest)

	return err
}

type NullableCreateCreditNoteRequest struct {
	value *CreateCreditNoteRequest
	isSet bool
}

func (v NullableCreateCreditNoteRequest) Get() *CreateCreditNoteRequest {
	return v.value
}

func (v *NullableCreateCreditNoteRequest) Set(val *CreateCreditNoteRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateCreditNoteRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateCreditNoteRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateCreditNoteRequest(val *CreateCreditNoteRequest) *NullableCreateCreditNoteRequest {
	return &NullableCreateCreditNoteRequest{value: val, isSet: true}
}

func (v NullableCreateCreditNoteRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateCreditNoteRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


