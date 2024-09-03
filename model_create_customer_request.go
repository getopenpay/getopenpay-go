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

// checks if the CreateCustomerRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateCustomerRequest{}

// CreateCustomerRequest struct for CreateCustomerRequest
type CreateCustomerRequest struct {
	FirstName NullableString `json:"first_name,omitempty"`
	LastName NullableString `json:"last_name,omitempty"`
	Line1 NullableString `json:"line1,omitempty"`
	Line2 NullableString `json:"line2,omitempty"`
	Line3 NullableString `json:"line3,omitempty"`
	City NullableString `json:"city,omitempty"`
	State NullableString `json:"state,omitempty"`
	Country NullableString `json:"country,omitempty"`
	ZipCode NullableString `json:"zip_code,omitempty"`
	CouponId NullableString `json:"coupon_id,omitempty"`
	PromotionCodeId NullableString `json:"promotion_code_id,omitempty"`
	Notes NullableString `json:"notes,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	// The customer’s email address.
	Email string `json:"email"`
}

type _CreateCustomerRequest CreateCustomerRequest

// NewCreateCustomerRequest instantiates a new CreateCustomerRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateCustomerRequest(email string) *CreateCustomerRequest {
	this := CreateCustomerRequest{}
	this.Email = email
	return &this
}

// NewCreateCustomerRequestWithDefaults instantiates a new CreateCustomerRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateCustomerRequestWithDefaults() *CreateCustomerRequest {
	this := CreateCustomerRequest{}
	return &this
}

// GetFirstName returns the FirstName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateCustomerRequest) GetFirstName() string {
	if o == nil || IsNil(o.FirstName.Get()) {
		var ret string
		return ret
	}
	return *o.FirstName.Get()
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateCustomerRequest) GetFirstNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FirstName.Get(), o.FirstName.IsSet()
}

// HasFirstName returns a boolean if a field has been set.
func (o *CreateCustomerRequest) HasFirstName() bool {
	if o != nil && o.FirstName.IsSet() {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given NullableString and assigns it to the FirstName field.
func (o *CreateCustomerRequest) SetFirstName(v string) {
	o.FirstName.Set(&v)
}
// SetFirstNameNil sets the value for FirstName to be an explicit nil
func (o *CreateCustomerRequest) SetFirstNameNil() {
	o.FirstName.Set(nil)
}

// UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
func (o *CreateCustomerRequest) UnsetFirstName() {
	o.FirstName.Unset()
}

// GetLastName returns the LastName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateCustomerRequest) GetLastName() string {
	if o == nil || IsNil(o.LastName.Get()) {
		var ret string
		return ret
	}
	return *o.LastName.Get()
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateCustomerRequest) GetLastNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastName.Get(), o.LastName.IsSet()
}

// HasLastName returns a boolean if a field has been set.
func (o *CreateCustomerRequest) HasLastName() bool {
	if o != nil && o.LastName.IsSet() {
		return true
	}

	return false
}

// SetLastName gets a reference to the given NullableString and assigns it to the LastName field.
func (o *CreateCustomerRequest) SetLastName(v string) {
	o.LastName.Set(&v)
}
// SetLastNameNil sets the value for LastName to be an explicit nil
func (o *CreateCustomerRequest) SetLastNameNil() {
	o.LastName.Set(nil)
}

// UnsetLastName ensures that no value is present for LastName, not even an explicit nil
func (o *CreateCustomerRequest) UnsetLastName() {
	o.LastName.Unset()
}

// GetLine1 returns the Line1 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateCustomerRequest) GetLine1() string {
	if o == nil || IsNil(o.Line1.Get()) {
		var ret string
		return ret
	}
	return *o.Line1.Get()
}

// GetLine1Ok returns a tuple with the Line1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateCustomerRequest) GetLine1Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Line1.Get(), o.Line1.IsSet()
}

// HasLine1 returns a boolean if a field has been set.
func (o *CreateCustomerRequest) HasLine1() bool {
	if o != nil && o.Line1.IsSet() {
		return true
	}

	return false
}

// SetLine1 gets a reference to the given NullableString and assigns it to the Line1 field.
func (o *CreateCustomerRequest) SetLine1(v string) {
	o.Line1.Set(&v)
}
// SetLine1Nil sets the value for Line1 to be an explicit nil
func (o *CreateCustomerRequest) SetLine1Nil() {
	o.Line1.Set(nil)
}

// UnsetLine1 ensures that no value is present for Line1, not even an explicit nil
func (o *CreateCustomerRequest) UnsetLine1() {
	o.Line1.Unset()
}

// GetLine2 returns the Line2 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateCustomerRequest) GetLine2() string {
	if o == nil || IsNil(o.Line2.Get()) {
		var ret string
		return ret
	}
	return *o.Line2.Get()
}

// GetLine2Ok returns a tuple with the Line2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateCustomerRequest) GetLine2Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Line2.Get(), o.Line2.IsSet()
}

// HasLine2 returns a boolean if a field has been set.
func (o *CreateCustomerRequest) HasLine2() bool {
	if o != nil && o.Line2.IsSet() {
		return true
	}

	return false
}

// SetLine2 gets a reference to the given NullableString and assigns it to the Line2 field.
func (o *CreateCustomerRequest) SetLine2(v string) {
	o.Line2.Set(&v)
}
// SetLine2Nil sets the value for Line2 to be an explicit nil
func (o *CreateCustomerRequest) SetLine2Nil() {
	o.Line2.Set(nil)
}

// UnsetLine2 ensures that no value is present for Line2, not even an explicit nil
func (o *CreateCustomerRequest) UnsetLine2() {
	o.Line2.Unset()
}

// GetLine3 returns the Line3 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateCustomerRequest) GetLine3() string {
	if o == nil || IsNil(o.Line3.Get()) {
		var ret string
		return ret
	}
	return *o.Line3.Get()
}

// GetLine3Ok returns a tuple with the Line3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateCustomerRequest) GetLine3Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Line3.Get(), o.Line3.IsSet()
}

// HasLine3 returns a boolean if a field has been set.
func (o *CreateCustomerRequest) HasLine3() bool {
	if o != nil && o.Line3.IsSet() {
		return true
	}

	return false
}

// SetLine3 gets a reference to the given NullableString and assigns it to the Line3 field.
func (o *CreateCustomerRequest) SetLine3(v string) {
	o.Line3.Set(&v)
}
// SetLine3Nil sets the value for Line3 to be an explicit nil
func (o *CreateCustomerRequest) SetLine3Nil() {
	o.Line3.Set(nil)
}

// UnsetLine3 ensures that no value is present for Line3, not even an explicit nil
func (o *CreateCustomerRequest) UnsetLine3() {
	o.Line3.Unset()
}

// GetCity returns the City field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateCustomerRequest) GetCity() string {
	if o == nil || IsNil(o.City.Get()) {
		var ret string
		return ret
	}
	return *o.City.Get()
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateCustomerRequest) GetCityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.City.Get(), o.City.IsSet()
}

// HasCity returns a boolean if a field has been set.
func (o *CreateCustomerRequest) HasCity() bool {
	if o != nil && o.City.IsSet() {
		return true
	}

	return false
}

// SetCity gets a reference to the given NullableString and assigns it to the City field.
func (o *CreateCustomerRequest) SetCity(v string) {
	o.City.Set(&v)
}
// SetCityNil sets the value for City to be an explicit nil
func (o *CreateCustomerRequest) SetCityNil() {
	o.City.Set(nil)
}

// UnsetCity ensures that no value is present for City, not even an explicit nil
func (o *CreateCustomerRequest) UnsetCity() {
	o.City.Unset()
}

// GetState returns the State field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateCustomerRequest) GetState() string {
	if o == nil || IsNil(o.State.Get()) {
		var ret string
		return ret
	}
	return *o.State.Get()
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateCustomerRequest) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.State.Get(), o.State.IsSet()
}

// HasState returns a boolean if a field has been set.
func (o *CreateCustomerRequest) HasState() bool {
	if o != nil && o.State.IsSet() {
		return true
	}

	return false
}

// SetState gets a reference to the given NullableString and assigns it to the State field.
func (o *CreateCustomerRequest) SetState(v string) {
	o.State.Set(&v)
}
// SetStateNil sets the value for State to be an explicit nil
func (o *CreateCustomerRequest) SetStateNil() {
	o.State.Set(nil)
}

// UnsetState ensures that no value is present for State, not even an explicit nil
func (o *CreateCustomerRequest) UnsetState() {
	o.State.Unset()
}

// GetCountry returns the Country field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateCustomerRequest) GetCountry() string {
	if o == nil || IsNil(o.Country.Get()) {
		var ret string
		return ret
	}
	return *o.Country.Get()
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateCustomerRequest) GetCountryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Country.Get(), o.Country.IsSet()
}

// HasCountry returns a boolean if a field has been set.
func (o *CreateCustomerRequest) HasCountry() bool {
	if o != nil && o.Country.IsSet() {
		return true
	}

	return false
}

// SetCountry gets a reference to the given NullableString and assigns it to the Country field.
func (o *CreateCustomerRequest) SetCountry(v string) {
	o.Country.Set(&v)
}
// SetCountryNil sets the value for Country to be an explicit nil
func (o *CreateCustomerRequest) SetCountryNil() {
	o.Country.Set(nil)
}

// UnsetCountry ensures that no value is present for Country, not even an explicit nil
func (o *CreateCustomerRequest) UnsetCountry() {
	o.Country.Unset()
}

// GetZipCode returns the ZipCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateCustomerRequest) GetZipCode() string {
	if o == nil || IsNil(o.ZipCode.Get()) {
		var ret string
		return ret
	}
	return *o.ZipCode.Get()
}

// GetZipCodeOk returns a tuple with the ZipCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateCustomerRequest) GetZipCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ZipCode.Get(), o.ZipCode.IsSet()
}

// HasZipCode returns a boolean if a field has been set.
func (o *CreateCustomerRequest) HasZipCode() bool {
	if o != nil && o.ZipCode.IsSet() {
		return true
	}

	return false
}

// SetZipCode gets a reference to the given NullableString and assigns it to the ZipCode field.
func (o *CreateCustomerRequest) SetZipCode(v string) {
	o.ZipCode.Set(&v)
}
// SetZipCodeNil sets the value for ZipCode to be an explicit nil
func (o *CreateCustomerRequest) SetZipCodeNil() {
	o.ZipCode.Set(nil)
}

// UnsetZipCode ensures that no value is present for ZipCode, not even an explicit nil
func (o *CreateCustomerRequest) UnsetZipCode() {
	o.ZipCode.Unset()
}

// GetCouponId returns the CouponId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateCustomerRequest) GetCouponId() string {
	if o == nil || IsNil(o.CouponId.Get()) {
		var ret string
		return ret
	}
	return *o.CouponId.Get()
}

// GetCouponIdOk returns a tuple with the CouponId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateCustomerRequest) GetCouponIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CouponId.Get(), o.CouponId.IsSet()
}

// HasCouponId returns a boolean if a field has been set.
func (o *CreateCustomerRequest) HasCouponId() bool {
	if o != nil && o.CouponId.IsSet() {
		return true
	}

	return false
}

// SetCouponId gets a reference to the given NullableString and assigns it to the CouponId field.
func (o *CreateCustomerRequest) SetCouponId(v string) {
	o.CouponId.Set(&v)
}
// SetCouponIdNil sets the value for CouponId to be an explicit nil
func (o *CreateCustomerRequest) SetCouponIdNil() {
	o.CouponId.Set(nil)
}

// UnsetCouponId ensures that no value is present for CouponId, not even an explicit nil
func (o *CreateCustomerRequest) UnsetCouponId() {
	o.CouponId.Unset()
}

// GetPromotionCodeId returns the PromotionCodeId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateCustomerRequest) GetPromotionCodeId() string {
	if o == nil || IsNil(o.PromotionCodeId.Get()) {
		var ret string
		return ret
	}
	return *o.PromotionCodeId.Get()
}

// GetPromotionCodeIdOk returns a tuple with the PromotionCodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateCustomerRequest) GetPromotionCodeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PromotionCodeId.Get(), o.PromotionCodeId.IsSet()
}

// HasPromotionCodeId returns a boolean if a field has been set.
func (o *CreateCustomerRequest) HasPromotionCodeId() bool {
	if o != nil && o.PromotionCodeId.IsSet() {
		return true
	}

	return false
}

// SetPromotionCodeId gets a reference to the given NullableString and assigns it to the PromotionCodeId field.
func (o *CreateCustomerRequest) SetPromotionCodeId(v string) {
	o.PromotionCodeId.Set(&v)
}
// SetPromotionCodeIdNil sets the value for PromotionCodeId to be an explicit nil
func (o *CreateCustomerRequest) SetPromotionCodeIdNil() {
	o.PromotionCodeId.Set(nil)
}

// UnsetPromotionCodeId ensures that no value is present for PromotionCodeId, not even an explicit nil
func (o *CreateCustomerRequest) UnsetPromotionCodeId() {
	o.PromotionCodeId.Unset()
}

// GetNotes returns the Notes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateCustomerRequest) GetNotes() string {
	if o == nil || IsNil(o.Notes.Get()) {
		var ret string
		return ret
	}
	return *o.Notes.Get()
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateCustomerRequest) GetNotesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Notes.Get(), o.Notes.IsSet()
}

// HasNotes returns a boolean if a field has been set.
func (o *CreateCustomerRequest) HasNotes() bool {
	if o != nil && o.Notes.IsSet() {
		return true
	}

	return false
}

// SetNotes gets a reference to the given NullableString and assigns it to the Notes field.
func (o *CreateCustomerRequest) SetNotes(v string) {
	o.Notes.Set(&v)
}
// SetNotesNil sets the value for Notes to be an explicit nil
func (o *CreateCustomerRequest) SetNotesNil() {
	o.Notes.Set(nil)
}

// UnsetNotes ensures that no value is present for Notes, not even an explicit nil
func (o *CreateCustomerRequest) UnsetNotes() {
	o.Notes.Unset()
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateCustomerRequest) GetCustomFields() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateCustomerRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *CreateCustomerRequest) HasCustomFields() bool {
	if o != nil && IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *CreateCustomerRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetEmail returns the Email field value
func (o *CreateCustomerRequest) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *CreateCustomerRequest) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *CreateCustomerRequest) SetEmail(v string) {
	o.Email = v
}

func (o CreateCustomerRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateCustomerRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.FirstName.IsSet() {
		toSerialize["first_name"] = o.FirstName.Get()
	}
	if o.LastName.IsSet() {
		toSerialize["last_name"] = o.LastName.Get()
	}
	if o.Line1.IsSet() {
		toSerialize["line1"] = o.Line1.Get()
	}
	if o.Line2.IsSet() {
		toSerialize["line2"] = o.Line2.Get()
	}
	if o.Line3.IsSet() {
		toSerialize["line3"] = o.Line3.Get()
	}
	if o.City.IsSet() {
		toSerialize["city"] = o.City.Get()
	}
	if o.State.IsSet() {
		toSerialize["state"] = o.State.Get()
	}
	if o.Country.IsSet() {
		toSerialize["country"] = o.Country.Get()
	}
	if o.ZipCode.IsSet() {
		toSerialize["zip_code"] = o.ZipCode.Get()
	}
	if o.CouponId.IsSet() {
		toSerialize["coupon_id"] = o.CouponId.Get()
	}
	if o.PromotionCodeId.IsSet() {
		toSerialize["promotion_code_id"] = o.PromotionCodeId.Get()
	}
	if o.Notes.IsSet() {
		toSerialize["notes"] = o.Notes.Get()
	}
	if o.CustomFields != nil {
		toSerialize["custom_fields"] = o.CustomFields
	}
	toSerialize["email"] = o.Email
	return toSerialize, nil
}

func (o *CreateCustomerRequest) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"email",
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

	varCreateCustomerRequest := _CreateCustomerRequest{}

	err = json.Unmarshal(bytes, &varCreateCustomerRequest)

	if err != nil {
		return err
	}

	*o = CreateCustomerRequest(varCreateCustomerRequest)

	return err
}

type NullableCreateCustomerRequest struct {
	value *CreateCustomerRequest
	isSet bool
}

func (v NullableCreateCustomerRequest) Get() *CreateCustomerRequest {
	return v.value
}

func (v *NullableCreateCustomerRequest) Set(val *CreateCustomerRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateCustomerRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateCustomerRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateCustomerRequest(val *CreateCustomerRequest) *NullableCreateCustomerRequest {
	return &NullableCreateCustomerRequest{value: val, isSet: true}
}

func (v NullableCreateCustomerRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateCustomerRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


