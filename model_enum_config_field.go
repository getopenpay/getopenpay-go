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

// checks if the EnumConfigField type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EnumConfigField{}

// EnumConfigField struct for EnumConfigField
type EnumConfigField struct {
	// The name of the field.
	Name string `json:"name"`
	// The key of the field in the config.
	Key string `json:"key"`
	Description NullableString `json:"description,omitempty"`
	Type *string `json:"type,omitempty"`
	Value NullableString `json:"value"`
	DefaultValue NullableString `json:"default_value,omitempty"`
	Options []string `json:"options"`
}

type _EnumConfigField EnumConfigField

// NewEnumConfigField instantiates a new EnumConfigField object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnumConfigField(name string, key string, value NullableString, options []string) *EnumConfigField {
	this := EnumConfigField{}
	this.Name = name
	this.Key = key
	var type_ string = "enum"
	this.Type = &type_
	this.Value = value
	this.Options = options
	return &this
}

// NewEnumConfigFieldWithDefaults instantiates a new EnumConfigField object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnumConfigFieldWithDefaults() *EnumConfigField {
	this := EnumConfigField{}
	var type_ string = "enum"
	this.Type = &type_
	return &this
}

// GetName returns the Name field value
func (o *EnumConfigField) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EnumConfigField) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *EnumConfigField) SetName(v string) {
	o.Name = v
}

// GetKey returns the Key field value
func (o *EnumConfigField) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *EnumConfigField) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *EnumConfigField) SetKey(v string) {
	o.Key = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnumConfigField) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnumConfigField) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *EnumConfigField) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *EnumConfigField) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *EnumConfigField) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *EnumConfigField) UnsetDescription() {
	o.Description.Unset()
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *EnumConfigField) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnumConfigField) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *EnumConfigField) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *EnumConfigField) SetType(v string) {
	o.Type = &v
}

// GetValue returns the Value field value
// If the value is explicit nil, the zero value for string will be returned
func (o *EnumConfigField) GetValue() string {
	if o == nil || o.Value.Get() == nil {
		var ret string
		return ret
	}

	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnumConfigField) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// SetValue sets field value
func (o *EnumConfigField) SetValue(v string) {
	o.Value.Set(&v)
}

// GetDefaultValue returns the DefaultValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnumConfigField) GetDefaultValue() string {
	if o == nil || IsNil(o.DefaultValue.Get()) {
		var ret string
		return ret
	}
	return *o.DefaultValue.Get()
}

// GetDefaultValueOk returns a tuple with the DefaultValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnumConfigField) GetDefaultValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DefaultValue.Get(), o.DefaultValue.IsSet()
}

// HasDefaultValue returns a boolean if a field has been set.
func (o *EnumConfigField) HasDefaultValue() bool {
	if o != nil && o.DefaultValue.IsSet() {
		return true
	}

	return false
}

// SetDefaultValue gets a reference to the given NullableString and assigns it to the DefaultValue field.
func (o *EnumConfigField) SetDefaultValue(v string) {
	o.DefaultValue.Set(&v)
}
// SetDefaultValueNil sets the value for DefaultValue to be an explicit nil
func (o *EnumConfigField) SetDefaultValueNil() {
	o.DefaultValue.Set(nil)
}

// UnsetDefaultValue ensures that no value is present for DefaultValue, not even an explicit nil
func (o *EnumConfigField) UnsetDefaultValue() {
	o.DefaultValue.Unset()
}

// GetOptions returns the Options field value
func (o *EnumConfigField) GetOptions() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *EnumConfigField) GetOptionsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Options, true
}

// SetOptions sets field value
func (o *EnumConfigField) SetOptions(v []string) {
	o.Options = v
}

func (o EnumConfigField) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnumConfigField) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["key"] = o.Key
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	toSerialize["value"] = o.Value.Get()
	if o.DefaultValue.IsSet() {
		toSerialize["default_value"] = o.DefaultValue.Get()
	}
	toSerialize["options"] = o.Options
	return toSerialize, nil
}

func (o *EnumConfigField) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"key",
		"value",
		"options",
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

	varEnumConfigField := _EnumConfigField{}

	err = json.Unmarshal(bytes, &varEnumConfigField)

	if err != nil {
		return err
	}

	*o = EnumConfigField(varEnumConfigField)

	return err
}

type NullableEnumConfigField struct {
	value *EnumConfigField
	isSet bool
}

func (v NullableEnumConfigField) Get() *EnumConfigField {
	return v.value
}

func (v *NullableEnumConfigField) Set(val *EnumConfigField) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumConfigField) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumConfigField) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumConfigField(val *EnumConfigField) *NullableEnumConfigField {
	return &NullableEnumConfigField{value: val, isSet: true}
}

func (v NullableEnumConfigField) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumConfigField) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

