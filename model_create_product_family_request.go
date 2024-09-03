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

// checks if the CreateProductFamilyRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateProductFamilyRequest{}

// CreateProductFamilyRequest struct for CreateProductFamilyRequest
type CreateProductFamilyRequest struct {
	// The name of the product family.
	Name string `json:"name"`
	Description NullableString `json:"description,omitempty"`
	Products []string `json:"products"`
	// A JSON object representing the hierarchy within the family.
	Hierarchy *string `json:"hierarchy,omitempty"`
}

type _CreateProductFamilyRequest CreateProductFamilyRequest

// NewCreateProductFamilyRequest instantiates a new CreateProductFamilyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateProductFamilyRequest(name string, products []string) *CreateProductFamilyRequest {
	this := CreateProductFamilyRequest{}
	this.Name = name
	this.Products = products
	var hierarchy string = "{}"
	this.Hierarchy = &hierarchy
	return &this
}

// NewCreateProductFamilyRequestWithDefaults instantiates a new CreateProductFamilyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateProductFamilyRequestWithDefaults() *CreateProductFamilyRequest {
	this := CreateProductFamilyRequest{}
	var hierarchy string = "{}"
	this.Hierarchy = &hierarchy
	return &this
}

// GetName returns the Name field value
func (o *CreateProductFamilyRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateProductFamilyRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateProductFamilyRequest) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateProductFamilyRequest) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateProductFamilyRequest) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateProductFamilyRequest) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *CreateProductFamilyRequest) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *CreateProductFamilyRequest) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *CreateProductFamilyRequest) UnsetDescription() {
	o.Description.Unset()
}

// GetProducts returns the Products field value
func (o *CreateProductFamilyRequest) GetProducts() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Products
}

// GetProductsOk returns a tuple with the Products field value
// and a boolean to check if the value has been set.
func (o *CreateProductFamilyRequest) GetProductsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Products, true
}

// SetProducts sets field value
func (o *CreateProductFamilyRequest) SetProducts(v []string) {
	o.Products = v
}

// GetHierarchy returns the Hierarchy field value if set, zero value otherwise.
func (o *CreateProductFamilyRequest) GetHierarchy() string {
	if o == nil || IsNil(o.Hierarchy) {
		var ret string
		return ret
	}
	return *o.Hierarchy
}

// GetHierarchyOk returns a tuple with the Hierarchy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateProductFamilyRequest) GetHierarchyOk() (*string, bool) {
	if o == nil || IsNil(o.Hierarchy) {
		return nil, false
	}
	return o.Hierarchy, true
}

// HasHierarchy returns a boolean if a field has been set.
func (o *CreateProductFamilyRequest) HasHierarchy() bool {
	if o != nil && !IsNil(o.Hierarchy) {
		return true
	}

	return false
}

// SetHierarchy gets a reference to the given string and assigns it to the Hierarchy field.
func (o *CreateProductFamilyRequest) SetHierarchy(v string) {
	o.Hierarchy = &v
}

func (o CreateProductFamilyRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateProductFamilyRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	toSerialize["products"] = o.Products
	if !IsNil(o.Hierarchy) {
		toSerialize["hierarchy"] = o.Hierarchy
	}
	return toSerialize, nil
}

func (o *CreateProductFamilyRequest) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"products",
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

	varCreateProductFamilyRequest := _CreateProductFamilyRequest{}

	err = json.Unmarshal(bytes, &varCreateProductFamilyRequest)

	if err != nil {
		return err
	}

	*o = CreateProductFamilyRequest(varCreateProductFamilyRequest)

	return err
}

type NullableCreateProductFamilyRequest struct {
	value *CreateProductFamilyRequest
	isSet bool
}

func (v NullableCreateProductFamilyRequest) Get() *CreateProductFamilyRequest {
	return v.value
}

func (v *NullableCreateProductFamilyRequest) Set(val *CreateProductFamilyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateProductFamilyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateProductFamilyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateProductFamilyRequest(val *CreateProductFamilyRequest) *NullableCreateProductFamilyRequest {
	return &NullableCreateProductFamilyRequest{value: val, isSet: true}
}

func (v NullableCreateProductFamilyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateProductFamilyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


