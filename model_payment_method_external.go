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

// checks if the PaymentMethodExternal type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentMethodExternal{}

// PaymentMethodExternal struct for PaymentMethodExternal
type PaymentMethodExternal struct {
	// Unique Identifier of the payment_method.
	Id string `json:"id"`
	Object *ObjectName `json:"object,omitempty"`
	// DateTime at which the object was created, in 'ISO 8601' format.
	CreatedAt time.Time `json:"created_at"`
	// DateTime at which the object was updated, in 'ISO 8601' format.
	UpdatedAt time.Time `json:"updated_at"`
	// If true, indicates that this object has been deleted
	IsDeleted *bool `json:"is_deleted,omitempty"`
	CustomerId NullableString `json:"customer_id,omitempty"`
	BillingAddress NullableCompleteAddress `json:"billing_address,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	Provider PaymentProviderType `json:"provider"`
	CardType NullableString `json:"card_type,omitempty"`
	LastFour NullableString `json:"last_four,omitempty"`
	// Display name for the payment method to show on the UI.
	DisplayName string `json:"display_name"`
}

type _PaymentMethodExternal PaymentMethodExternal

// NewPaymentMethodExternal instantiates a new PaymentMethodExternal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentMethodExternal(id string, createdAt time.Time, updatedAt time.Time, provider PaymentProviderType, displayName string) *PaymentMethodExternal {
	this := PaymentMethodExternal{}
	this.Id = id
	var object ObjectName = OBJECTNAME_PAYMENT_METHOD
	this.Object = &object
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	var isDeleted bool = false
	this.IsDeleted = &isDeleted
	this.Provider = provider
	this.DisplayName = displayName
	return &this
}

// NewPaymentMethodExternalWithDefaults instantiates a new PaymentMethodExternal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentMethodExternalWithDefaults() *PaymentMethodExternal {
	this := PaymentMethodExternal{}
	var object ObjectName = OBJECTNAME_PAYMENT_METHOD
	this.Object = &object
	var isDeleted bool = false
	this.IsDeleted = &isDeleted
	return &this
}

// GetId returns the Id field value
func (o *PaymentMethodExternal) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PaymentMethodExternal) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PaymentMethodExternal) SetId(v string) {
	o.Id = v
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *PaymentMethodExternal) GetObject() ObjectName {
	if o == nil || IsNil(o.Object) {
		var ret ObjectName
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodExternal) GetObjectOk() (*ObjectName, bool) {
	if o == nil || IsNil(o.Object) {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *PaymentMethodExternal) HasObject() bool {
	if o != nil && !IsNil(o.Object) {
		return true
	}

	return false
}

// SetObject gets a reference to the given ObjectName and assigns it to the Object field.
func (o *PaymentMethodExternal) SetObject(v ObjectName) {
	o.Object = &v
}

// GetCreatedAt returns the CreatedAt field value
func (o *PaymentMethodExternal) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *PaymentMethodExternal) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *PaymentMethodExternal) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *PaymentMethodExternal) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *PaymentMethodExternal) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *PaymentMethodExternal) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetIsDeleted returns the IsDeleted field value if set, zero value otherwise.
func (o *PaymentMethodExternal) GetIsDeleted() bool {
	if o == nil || IsNil(o.IsDeleted) {
		var ret bool
		return ret
	}
	return *o.IsDeleted
}

// GetIsDeletedOk returns a tuple with the IsDeleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodExternal) GetIsDeletedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDeleted) {
		return nil, false
	}
	return o.IsDeleted, true
}

// HasIsDeleted returns a boolean if a field has been set.
func (o *PaymentMethodExternal) HasIsDeleted() bool {
	if o != nil && !IsNil(o.IsDeleted) {
		return true
	}

	return false
}

// SetIsDeleted gets a reference to the given bool and assigns it to the IsDeleted field.
func (o *PaymentMethodExternal) SetIsDeleted(v bool) {
	o.IsDeleted = &v
}

// GetCustomerId returns the CustomerId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentMethodExternal) GetCustomerId() string {
	if o == nil || IsNil(o.CustomerId.Get()) {
		var ret string
		return ret
	}
	return *o.CustomerId.Get()
}

// GetCustomerIdOk returns a tuple with the CustomerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentMethodExternal) GetCustomerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomerId.Get(), o.CustomerId.IsSet()
}

// HasCustomerId returns a boolean if a field has been set.
func (o *PaymentMethodExternal) HasCustomerId() bool {
	if o != nil && o.CustomerId.IsSet() {
		return true
	}

	return false
}

// SetCustomerId gets a reference to the given NullableString and assigns it to the CustomerId field.
func (o *PaymentMethodExternal) SetCustomerId(v string) {
	o.CustomerId.Set(&v)
}
// SetCustomerIdNil sets the value for CustomerId to be an explicit nil
func (o *PaymentMethodExternal) SetCustomerIdNil() {
	o.CustomerId.Set(nil)
}

// UnsetCustomerId ensures that no value is present for CustomerId, not even an explicit nil
func (o *PaymentMethodExternal) UnsetCustomerId() {
	o.CustomerId.Unset()
}

// GetBillingAddress returns the BillingAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentMethodExternal) GetBillingAddress() CompleteAddress {
	if o == nil || IsNil(o.BillingAddress.Get()) {
		var ret CompleteAddress
		return ret
	}
	return *o.BillingAddress.Get()
}

// GetBillingAddressOk returns a tuple with the BillingAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentMethodExternal) GetBillingAddressOk() (*CompleteAddress, bool) {
	if o == nil {
		return nil, false
	}
	return o.BillingAddress.Get(), o.BillingAddress.IsSet()
}

// HasBillingAddress returns a boolean if a field has been set.
func (o *PaymentMethodExternal) HasBillingAddress() bool {
	if o != nil && o.BillingAddress.IsSet() {
		return true
	}

	return false
}

// SetBillingAddress gets a reference to the given NullableCompleteAddress and assigns it to the BillingAddress field.
func (o *PaymentMethodExternal) SetBillingAddress(v CompleteAddress) {
	o.BillingAddress.Set(&v)
}
// SetBillingAddressNil sets the value for BillingAddress to be an explicit nil
func (o *PaymentMethodExternal) SetBillingAddressNil() {
	o.BillingAddress.Set(nil)
}

// UnsetBillingAddress ensures that no value is present for BillingAddress, not even an explicit nil
func (o *PaymentMethodExternal) UnsetBillingAddress() {
	o.BillingAddress.Unset()
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentMethodExternal) GetMetadata() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentMethodExternal) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *PaymentMethodExternal) HasMetadata() bool {
	if o != nil && IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *PaymentMethodExternal) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetProvider returns the Provider field value
func (o *PaymentMethodExternal) GetProvider() PaymentProviderType {
	if o == nil {
		var ret PaymentProviderType
		return ret
	}

	return o.Provider
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
func (o *PaymentMethodExternal) GetProviderOk() (*PaymentProviderType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Provider, true
}

// SetProvider sets field value
func (o *PaymentMethodExternal) SetProvider(v PaymentProviderType) {
	o.Provider = v
}

// GetCardType returns the CardType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentMethodExternal) GetCardType() string {
	if o == nil || IsNil(o.CardType.Get()) {
		var ret string
		return ret
	}
	return *o.CardType.Get()
}

// GetCardTypeOk returns a tuple with the CardType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentMethodExternal) GetCardTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CardType.Get(), o.CardType.IsSet()
}

// HasCardType returns a boolean if a field has been set.
func (o *PaymentMethodExternal) HasCardType() bool {
	if o != nil && o.CardType.IsSet() {
		return true
	}

	return false
}

// SetCardType gets a reference to the given NullableString and assigns it to the CardType field.
func (o *PaymentMethodExternal) SetCardType(v string) {
	o.CardType.Set(&v)
}
// SetCardTypeNil sets the value for CardType to be an explicit nil
func (o *PaymentMethodExternal) SetCardTypeNil() {
	o.CardType.Set(nil)
}

// UnsetCardType ensures that no value is present for CardType, not even an explicit nil
func (o *PaymentMethodExternal) UnsetCardType() {
	o.CardType.Unset()
}

// GetLastFour returns the LastFour field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentMethodExternal) GetLastFour() string {
	if o == nil || IsNil(o.LastFour.Get()) {
		var ret string
		return ret
	}
	return *o.LastFour.Get()
}

// GetLastFourOk returns a tuple with the LastFour field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentMethodExternal) GetLastFourOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastFour.Get(), o.LastFour.IsSet()
}

// HasLastFour returns a boolean if a field has been set.
func (o *PaymentMethodExternal) HasLastFour() bool {
	if o != nil && o.LastFour.IsSet() {
		return true
	}

	return false
}

// SetLastFour gets a reference to the given NullableString and assigns it to the LastFour field.
func (o *PaymentMethodExternal) SetLastFour(v string) {
	o.LastFour.Set(&v)
}
// SetLastFourNil sets the value for LastFour to be an explicit nil
func (o *PaymentMethodExternal) SetLastFourNil() {
	o.LastFour.Set(nil)
}

// UnsetLastFour ensures that no value is present for LastFour, not even an explicit nil
func (o *PaymentMethodExternal) UnsetLastFour() {
	o.LastFour.Unset()
}

// GetDisplayName returns the DisplayName field value
func (o *PaymentMethodExternal) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *PaymentMethodExternal) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *PaymentMethodExternal) SetDisplayName(v string) {
	o.DisplayName = v
}

func (o PaymentMethodExternal) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentMethodExternal) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.Object) {
		toSerialize["object"] = o.Object
	}
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_at"] = o.UpdatedAt
	if !IsNil(o.IsDeleted) {
		toSerialize["is_deleted"] = o.IsDeleted
	}
	if o.CustomerId.IsSet() {
		toSerialize["customer_id"] = o.CustomerId.Get()
	}
	if o.BillingAddress.IsSet() {
		toSerialize["billing_address"] = o.BillingAddress.Get()
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	toSerialize["provider"] = o.Provider
	if o.CardType.IsSet() {
		toSerialize["card_type"] = o.CardType.Get()
	}
	if o.LastFour.IsSet() {
		toSerialize["last_four"] = o.LastFour.Get()
	}
	toSerialize["display_name"] = o.DisplayName
	return toSerialize, nil
}

func (o *PaymentMethodExternal) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"created_at",
		"updated_at",
		"provider",
		"display_name",
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

	varPaymentMethodExternal := _PaymentMethodExternal{}

	err = json.Unmarshal(bytes, &varPaymentMethodExternal)

	if err != nil {
		return err
	}

	*o = PaymentMethodExternal(varPaymentMethodExternal)

	return err
}

type NullablePaymentMethodExternal struct {
	value *PaymentMethodExternal
	isSet bool
}

func (v NullablePaymentMethodExternal) Get() *PaymentMethodExternal {
	return v.value
}

func (v *NullablePaymentMethodExternal) Set(val *PaymentMethodExternal) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentMethodExternal) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentMethodExternal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentMethodExternal(val *PaymentMethodExternal) *NullablePaymentMethodExternal {
	return &NullablePaymentMethodExternal{value: val, isSet: true}
}

func (v NullablePaymentMethodExternal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentMethodExternal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


