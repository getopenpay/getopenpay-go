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

// checks if the RecurringDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecurringDetails{}

// RecurringDetails struct for RecurringDetails
type RecurringDetails struct {
	UsageType UsageTypeEnum `json:"usage_type"`
	AggregateUsage NullableUsageAggMethodEnum `json:"aggregate_usage,omitempty"`
	TrialPeriodDays *int32 `json:"trial_period_days,omitempty"`
}

type _RecurringDetails RecurringDetails

// NewRecurringDetails instantiates a new RecurringDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecurringDetails(usageType UsageTypeEnum) *RecurringDetails {
	this := RecurringDetails{}
	this.UsageType = usageType
	var trialPeriodDays int32 = 0
	this.TrialPeriodDays = &trialPeriodDays
	return &this
}

// NewRecurringDetailsWithDefaults instantiates a new RecurringDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecurringDetailsWithDefaults() *RecurringDetails {
	this := RecurringDetails{}
	var trialPeriodDays int32 = 0
	this.TrialPeriodDays = &trialPeriodDays
	return &this
}

// GetUsageType returns the UsageType field value
func (o *RecurringDetails) GetUsageType() UsageTypeEnum {
	if o == nil {
		var ret UsageTypeEnum
		return ret
	}

	return o.UsageType
}

// GetUsageTypeOk returns a tuple with the UsageType field value
// and a boolean to check if the value has been set.
func (o *RecurringDetails) GetUsageTypeOk() (*UsageTypeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UsageType, true
}

// SetUsageType sets field value
func (o *RecurringDetails) SetUsageType(v UsageTypeEnum) {
	o.UsageType = v
}

// GetAggregateUsage returns the AggregateUsage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecurringDetails) GetAggregateUsage() UsageAggMethodEnum {
	if o == nil || IsNil(o.AggregateUsage.Get()) {
		var ret UsageAggMethodEnum
		return ret
	}
	return *o.AggregateUsage.Get()
}

// GetAggregateUsageOk returns a tuple with the AggregateUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecurringDetails) GetAggregateUsageOk() (*UsageAggMethodEnum, bool) {
	if o == nil {
		return nil, false
	}
	return o.AggregateUsage.Get(), o.AggregateUsage.IsSet()
}

// HasAggregateUsage returns a boolean if a field has been set.
func (o *RecurringDetails) HasAggregateUsage() bool {
	if o != nil && o.AggregateUsage.IsSet() {
		return true
	}

	return false
}

// SetAggregateUsage gets a reference to the given NullableUsageAggMethodEnum and assigns it to the AggregateUsage field.
func (o *RecurringDetails) SetAggregateUsage(v UsageAggMethodEnum) {
	o.AggregateUsage.Set(&v)
}
// SetAggregateUsageNil sets the value for AggregateUsage to be an explicit nil
func (o *RecurringDetails) SetAggregateUsageNil() {
	o.AggregateUsage.Set(nil)
}

// UnsetAggregateUsage ensures that no value is present for AggregateUsage, not even an explicit nil
func (o *RecurringDetails) UnsetAggregateUsage() {
	o.AggregateUsage.Unset()
}

// GetTrialPeriodDays returns the TrialPeriodDays field value if set, zero value otherwise.
func (o *RecurringDetails) GetTrialPeriodDays() int32 {
	if o == nil || IsNil(o.TrialPeriodDays) {
		var ret int32
		return ret
	}
	return *o.TrialPeriodDays
}

// GetTrialPeriodDaysOk returns a tuple with the TrialPeriodDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecurringDetails) GetTrialPeriodDaysOk() (*int32, bool) {
	if o == nil || IsNil(o.TrialPeriodDays) {
		return nil, false
	}
	return o.TrialPeriodDays, true
}

// HasTrialPeriodDays returns a boolean if a field has been set.
func (o *RecurringDetails) HasTrialPeriodDays() bool {
	if o != nil && !IsNil(o.TrialPeriodDays) {
		return true
	}

	return false
}

// SetTrialPeriodDays gets a reference to the given int32 and assigns it to the TrialPeriodDays field.
func (o *RecurringDetails) SetTrialPeriodDays(v int32) {
	o.TrialPeriodDays = &v
}

func (o RecurringDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecurringDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["usage_type"] = o.UsageType
	if o.AggregateUsage.IsSet() {
		toSerialize["aggregate_usage"] = o.AggregateUsage.Get()
	}
	if !IsNil(o.TrialPeriodDays) {
		toSerialize["trial_period_days"] = o.TrialPeriodDays
	}
	return toSerialize, nil
}

func (o *RecurringDetails) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"usage_type",
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

	varRecurringDetails := _RecurringDetails{}

	err = json.Unmarshal(bytes, &varRecurringDetails)

	if err != nil {
		return err
	}

	*o = RecurringDetails(varRecurringDetails)

	return err
}

type NullableRecurringDetails struct {
	value *RecurringDetails
	isSet bool
}

func (v NullableRecurringDetails) Get() *RecurringDetails {
	return v.value
}

func (v *NullableRecurringDetails) Set(val *RecurringDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableRecurringDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableRecurringDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecurringDetails(val *RecurringDetails) *NullableRecurringDetails {
	return &NullableRecurringDetails{value: val, isSet: true}
}

func (v NullableRecurringDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecurringDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


