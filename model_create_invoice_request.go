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

// checks if the CreateInvoiceRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateInvoiceRequest{}

// CreateInvoiceRequest struct for CreateInvoiceRequest
type CreateInvoiceRequest struct {
	SubscriptionId NullableString `json:"subscription_id,omitempty"`
	PaymentMethodId NullableString `json:"payment_method_id,omitempty"`
	CollectionMethod NullableCollectionMethodEnum `json:"collection_method,omitempty"`
	// Description for newly created invoice
	Description *string `json:"description,omitempty"`
	CouponId NullableString `json:"coupon_id,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	SelectedProductPriceQuantity []SelectedPriceQuantity `json:"selected_product_price_quantity,omitempty"`
	InvoiceType *InvoiceType `json:"invoice_type,omitempty"`
	// The external id of the customer.
	CustomerId string `json:"customer_id"`
	NetD NullableInt32 `json:"net_d,omitempty"`
	EmailInvoiceOnFinalization NullableBool `json:"email_invoice_on_finalization,omitempty"`
	FinalizeInvoiceImmediately NullableBool `json:"finalize_invoice_immediately,omitempty"`
}

type _CreateInvoiceRequest CreateInvoiceRequest

// NewCreateInvoiceRequest instantiates a new CreateInvoiceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateInvoiceRequest(customerId string) *CreateInvoiceRequest {
	this := CreateInvoiceRequest{}
	var description string = "Manual creation of invoice"
	this.Description = &description
	var invoiceType InvoiceType = INVOICETYPE_STANDARD
	this.InvoiceType = &invoiceType
	this.CustomerId = customerId
	return &this
}

// NewCreateInvoiceRequestWithDefaults instantiates a new CreateInvoiceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateInvoiceRequestWithDefaults() *CreateInvoiceRequest {
	this := CreateInvoiceRequest{}
	var description string = "Manual creation of invoice"
	this.Description = &description
	var invoiceType InvoiceType = INVOICETYPE_STANDARD
	this.InvoiceType = &invoiceType
	return &this
}

// GetSubscriptionId returns the SubscriptionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateInvoiceRequest) GetSubscriptionId() string {
	if o == nil || IsNil(o.SubscriptionId.Get()) {
		var ret string
		return ret
	}
	return *o.SubscriptionId.Get()
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateInvoiceRequest) GetSubscriptionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SubscriptionId.Get(), o.SubscriptionId.IsSet()
}

// HasSubscriptionId returns a boolean if a field has been set.
func (o *CreateInvoiceRequest) HasSubscriptionId() bool {
	if o != nil && o.SubscriptionId.IsSet() {
		return true
	}

	return false
}

// SetSubscriptionId gets a reference to the given NullableString and assigns it to the SubscriptionId field.
func (o *CreateInvoiceRequest) SetSubscriptionId(v string) {
	o.SubscriptionId.Set(&v)
}
// SetSubscriptionIdNil sets the value for SubscriptionId to be an explicit nil
func (o *CreateInvoiceRequest) SetSubscriptionIdNil() {
	o.SubscriptionId.Set(nil)
}

// UnsetSubscriptionId ensures that no value is present for SubscriptionId, not even an explicit nil
func (o *CreateInvoiceRequest) UnsetSubscriptionId() {
	o.SubscriptionId.Unset()
}

// GetPaymentMethodId returns the PaymentMethodId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateInvoiceRequest) GetPaymentMethodId() string {
	if o == nil || IsNil(o.PaymentMethodId.Get()) {
		var ret string
		return ret
	}
	return *o.PaymentMethodId.Get()
}

// GetPaymentMethodIdOk returns a tuple with the PaymentMethodId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateInvoiceRequest) GetPaymentMethodIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PaymentMethodId.Get(), o.PaymentMethodId.IsSet()
}

// HasPaymentMethodId returns a boolean if a field has been set.
func (o *CreateInvoiceRequest) HasPaymentMethodId() bool {
	if o != nil && o.PaymentMethodId.IsSet() {
		return true
	}

	return false
}

// SetPaymentMethodId gets a reference to the given NullableString and assigns it to the PaymentMethodId field.
func (o *CreateInvoiceRequest) SetPaymentMethodId(v string) {
	o.PaymentMethodId.Set(&v)
}
// SetPaymentMethodIdNil sets the value for PaymentMethodId to be an explicit nil
func (o *CreateInvoiceRequest) SetPaymentMethodIdNil() {
	o.PaymentMethodId.Set(nil)
}

// UnsetPaymentMethodId ensures that no value is present for PaymentMethodId, not even an explicit nil
func (o *CreateInvoiceRequest) UnsetPaymentMethodId() {
	o.PaymentMethodId.Unset()
}

// GetCollectionMethod returns the CollectionMethod field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateInvoiceRequest) GetCollectionMethod() CollectionMethodEnum {
	if o == nil || IsNil(o.CollectionMethod.Get()) {
		var ret CollectionMethodEnum
		return ret
	}
	return *o.CollectionMethod.Get()
}

// GetCollectionMethodOk returns a tuple with the CollectionMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateInvoiceRequest) GetCollectionMethodOk() (*CollectionMethodEnum, bool) {
	if o == nil {
		return nil, false
	}
	return o.CollectionMethod.Get(), o.CollectionMethod.IsSet()
}

// HasCollectionMethod returns a boolean if a field has been set.
func (o *CreateInvoiceRequest) HasCollectionMethod() bool {
	if o != nil && o.CollectionMethod.IsSet() {
		return true
	}

	return false
}

// SetCollectionMethod gets a reference to the given NullableCollectionMethodEnum and assigns it to the CollectionMethod field.
func (o *CreateInvoiceRequest) SetCollectionMethod(v CollectionMethodEnum) {
	o.CollectionMethod.Set(&v)
}
// SetCollectionMethodNil sets the value for CollectionMethod to be an explicit nil
func (o *CreateInvoiceRequest) SetCollectionMethodNil() {
	o.CollectionMethod.Set(nil)
}

// UnsetCollectionMethod ensures that no value is present for CollectionMethod, not even an explicit nil
func (o *CreateInvoiceRequest) UnsetCollectionMethod() {
	o.CollectionMethod.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateInvoiceRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateInvoiceRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateInvoiceRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateInvoiceRequest) SetDescription(v string) {
	o.Description = &v
}

// GetCouponId returns the CouponId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateInvoiceRequest) GetCouponId() string {
	if o == nil || IsNil(o.CouponId.Get()) {
		var ret string
		return ret
	}
	return *o.CouponId.Get()
}

// GetCouponIdOk returns a tuple with the CouponId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateInvoiceRequest) GetCouponIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CouponId.Get(), o.CouponId.IsSet()
}

// HasCouponId returns a boolean if a field has been set.
func (o *CreateInvoiceRequest) HasCouponId() bool {
	if o != nil && o.CouponId.IsSet() {
		return true
	}

	return false
}

// SetCouponId gets a reference to the given NullableString and assigns it to the CouponId field.
func (o *CreateInvoiceRequest) SetCouponId(v string) {
	o.CouponId.Set(&v)
}
// SetCouponIdNil sets the value for CouponId to be an explicit nil
func (o *CreateInvoiceRequest) SetCouponIdNil() {
	o.CouponId.Set(nil)
}

// UnsetCouponId ensures that no value is present for CouponId, not even an explicit nil
func (o *CreateInvoiceRequest) UnsetCouponId() {
	o.CouponId.Unset()
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateInvoiceRequest) GetCustomFields() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateInvoiceRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *CreateInvoiceRequest) HasCustomFields() bool {
	if o != nil && IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *CreateInvoiceRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetSelectedProductPriceQuantity returns the SelectedProductPriceQuantity field value if set, zero value otherwise.
func (o *CreateInvoiceRequest) GetSelectedProductPriceQuantity() []SelectedPriceQuantity {
	if o == nil || IsNil(o.SelectedProductPriceQuantity) {
		var ret []SelectedPriceQuantity
		return ret
	}
	return o.SelectedProductPriceQuantity
}

// GetSelectedProductPriceQuantityOk returns a tuple with the SelectedProductPriceQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateInvoiceRequest) GetSelectedProductPriceQuantityOk() ([]SelectedPriceQuantity, bool) {
	if o == nil || IsNil(o.SelectedProductPriceQuantity) {
		return nil, false
	}
	return o.SelectedProductPriceQuantity, true
}

// HasSelectedProductPriceQuantity returns a boolean if a field has been set.
func (o *CreateInvoiceRequest) HasSelectedProductPriceQuantity() bool {
	if o != nil && !IsNil(o.SelectedProductPriceQuantity) {
		return true
	}

	return false
}

// SetSelectedProductPriceQuantity gets a reference to the given []SelectedPriceQuantity and assigns it to the SelectedProductPriceQuantity field.
func (o *CreateInvoiceRequest) SetSelectedProductPriceQuantity(v []SelectedPriceQuantity) {
	o.SelectedProductPriceQuantity = v
}

// GetInvoiceType returns the InvoiceType field value if set, zero value otherwise.
func (o *CreateInvoiceRequest) GetInvoiceType() InvoiceType {
	if o == nil || IsNil(o.InvoiceType) {
		var ret InvoiceType
		return ret
	}
	return *o.InvoiceType
}

// GetInvoiceTypeOk returns a tuple with the InvoiceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateInvoiceRequest) GetInvoiceTypeOk() (*InvoiceType, bool) {
	if o == nil || IsNil(o.InvoiceType) {
		return nil, false
	}
	return o.InvoiceType, true
}

// HasInvoiceType returns a boolean if a field has been set.
func (o *CreateInvoiceRequest) HasInvoiceType() bool {
	if o != nil && !IsNil(o.InvoiceType) {
		return true
	}

	return false
}

// SetInvoiceType gets a reference to the given InvoiceType and assigns it to the InvoiceType field.
func (o *CreateInvoiceRequest) SetInvoiceType(v InvoiceType) {
	o.InvoiceType = &v
}

// GetCustomerId returns the CustomerId field value
func (o *CreateInvoiceRequest) GetCustomerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value
// and a boolean to check if the value has been set.
func (o *CreateInvoiceRequest) GetCustomerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomerId, true
}

// SetCustomerId sets field value
func (o *CreateInvoiceRequest) SetCustomerId(v string) {
	o.CustomerId = v
}

// GetNetD returns the NetD field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateInvoiceRequest) GetNetD() int32 {
	if o == nil || IsNil(o.NetD.Get()) {
		var ret int32
		return ret
	}
	return *o.NetD.Get()
}

// GetNetDOk returns a tuple with the NetD field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateInvoiceRequest) GetNetDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.NetD.Get(), o.NetD.IsSet()
}

// HasNetD returns a boolean if a field has been set.
func (o *CreateInvoiceRequest) HasNetD() bool {
	if o != nil && o.NetD.IsSet() {
		return true
	}

	return false
}

// SetNetD gets a reference to the given NullableInt32 and assigns it to the NetD field.
func (o *CreateInvoiceRequest) SetNetD(v int32) {
	o.NetD.Set(&v)
}
// SetNetDNil sets the value for NetD to be an explicit nil
func (o *CreateInvoiceRequest) SetNetDNil() {
	o.NetD.Set(nil)
}

// UnsetNetD ensures that no value is present for NetD, not even an explicit nil
func (o *CreateInvoiceRequest) UnsetNetD() {
	o.NetD.Unset()
}

// GetEmailInvoiceOnFinalization returns the EmailInvoiceOnFinalization field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateInvoiceRequest) GetEmailInvoiceOnFinalization() bool {
	if o == nil || IsNil(o.EmailInvoiceOnFinalization.Get()) {
		var ret bool
		return ret
	}
	return *o.EmailInvoiceOnFinalization.Get()
}

// GetEmailInvoiceOnFinalizationOk returns a tuple with the EmailInvoiceOnFinalization field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateInvoiceRequest) GetEmailInvoiceOnFinalizationOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.EmailInvoiceOnFinalization.Get(), o.EmailInvoiceOnFinalization.IsSet()
}

// HasEmailInvoiceOnFinalization returns a boolean if a field has been set.
func (o *CreateInvoiceRequest) HasEmailInvoiceOnFinalization() bool {
	if o != nil && o.EmailInvoiceOnFinalization.IsSet() {
		return true
	}

	return false
}

// SetEmailInvoiceOnFinalization gets a reference to the given NullableBool and assigns it to the EmailInvoiceOnFinalization field.
func (o *CreateInvoiceRequest) SetEmailInvoiceOnFinalization(v bool) {
	o.EmailInvoiceOnFinalization.Set(&v)
}
// SetEmailInvoiceOnFinalizationNil sets the value for EmailInvoiceOnFinalization to be an explicit nil
func (o *CreateInvoiceRequest) SetEmailInvoiceOnFinalizationNil() {
	o.EmailInvoiceOnFinalization.Set(nil)
}

// UnsetEmailInvoiceOnFinalization ensures that no value is present for EmailInvoiceOnFinalization, not even an explicit nil
func (o *CreateInvoiceRequest) UnsetEmailInvoiceOnFinalization() {
	o.EmailInvoiceOnFinalization.Unset()
}

// GetFinalizeInvoiceImmediately returns the FinalizeInvoiceImmediately field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateInvoiceRequest) GetFinalizeInvoiceImmediately() bool {
	if o == nil || IsNil(o.FinalizeInvoiceImmediately.Get()) {
		var ret bool
		return ret
	}
	return *o.FinalizeInvoiceImmediately.Get()
}

// GetFinalizeInvoiceImmediatelyOk returns a tuple with the FinalizeInvoiceImmediately field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateInvoiceRequest) GetFinalizeInvoiceImmediatelyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.FinalizeInvoiceImmediately.Get(), o.FinalizeInvoiceImmediately.IsSet()
}

// HasFinalizeInvoiceImmediately returns a boolean if a field has been set.
func (o *CreateInvoiceRequest) HasFinalizeInvoiceImmediately() bool {
	if o != nil && o.FinalizeInvoiceImmediately.IsSet() {
		return true
	}

	return false
}

// SetFinalizeInvoiceImmediately gets a reference to the given NullableBool and assigns it to the FinalizeInvoiceImmediately field.
func (o *CreateInvoiceRequest) SetFinalizeInvoiceImmediately(v bool) {
	o.FinalizeInvoiceImmediately.Set(&v)
}
// SetFinalizeInvoiceImmediatelyNil sets the value for FinalizeInvoiceImmediately to be an explicit nil
func (o *CreateInvoiceRequest) SetFinalizeInvoiceImmediatelyNil() {
	o.FinalizeInvoiceImmediately.Set(nil)
}

// UnsetFinalizeInvoiceImmediately ensures that no value is present for FinalizeInvoiceImmediately, not even an explicit nil
func (o *CreateInvoiceRequest) UnsetFinalizeInvoiceImmediately() {
	o.FinalizeInvoiceImmediately.Unset()
}

func (o CreateInvoiceRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateInvoiceRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.SubscriptionId.IsSet() {
		toSerialize["subscription_id"] = o.SubscriptionId.Get()
	}
	if o.PaymentMethodId.IsSet() {
		toSerialize["payment_method_id"] = o.PaymentMethodId.Get()
	}
	if o.CollectionMethod.IsSet() {
		toSerialize["collection_method"] = o.CollectionMethod.Get()
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if o.CouponId.IsSet() {
		toSerialize["coupon_id"] = o.CouponId.Get()
	}
	if o.CustomFields != nil {
		toSerialize["custom_fields"] = o.CustomFields
	}
	if !IsNil(o.SelectedProductPriceQuantity) {
		toSerialize["selected_product_price_quantity"] = o.SelectedProductPriceQuantity
	}
	if !IsNil(o.InvoiceType) {
		toSerialize["invoice_type"] = o.InvoiceType
	}
	toSerialize["customer_id"] = o.CustomerId
	if o.NetD.IsSet() {
		toSerialize["net_d"] = o.NetD.Get()
	}
	if o.EmailInvoiceOnFinalization.IsSet() {
		toSerialize["email_invoice_on_finalization"] = o.EmailInvoiceOnFinalization.Get()
	}
	if o.FinalizeInvoiceImmediately.IsSet() {
		toSerialize["finalize_invoice_immediately"] = o.FinalizeInvoiceImmediately.Get()
	}
	return toSerialize, nil
}

func (o *CreateInvoiceRequest) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"customer_id",
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

	varCreateInvoiceRequest := _CreateInvoiceRequest{}

	err = json.Unmarshal(bytes, &varCreateInvoiceRequest)

	if err != nil {
		return err
	}

	*o = CreateInvoiceRequest(varCreateInvoiceRequest)

	return err
}

type NullableCreateInvoiceRequest struct {
	value *CreateInvoiceRequest
	isSet bool
}

func (v NullableCreateInvoiceRequest) Get() *CreateInvoiceRequest {
	return v.value
}

func (v *NullableCreateInvoiceRequest) Set(val *CreateInvoiceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateInvoiceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateInvoiceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateInvoiceRequest(val *CreateInvoiceRequest) *NullableCreateInvoiceRequest {
	return &NullableCreateInvoiceRequest{value: val, isSet: true}
}

func (v NullableCreateInvoiceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateInvoiceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


