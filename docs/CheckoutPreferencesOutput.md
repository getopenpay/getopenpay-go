# CheckoutPreferencesOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FallbackCascadePreferences** | Pointer to [**[]FallbackConfigurationOutput**](FallbackConfigurationOutput.md) | This object contains a list of price groups that will be used to fall back to if the selected product price quantity fails to be created. Processor preferences here will override the processor preferences in the processor_preferences field. | [optional] [default to []]
**FallbackCascadeSelectedProductPriceQuantity** | Pointer to [**[][]SelectedPriceQuantity**]([]SelectedPriceQuantity.md) | This object contains a list of price groups that will be used to fall back to if the selected product price quantity fails to be created. | [optional] [default to []]
**OverrideFeatureRolloutSettings** | Pointer to **map[string]interface{}** |  | [optional] 
**ProcessorPreferences** | Pointer to [**NullableCheckoutProcessorsPreferences**](CheckoutProcessorsPreferences.md) |  | [optional] 
**Unused** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCheckoutPreferencesOutput

`func NewCheckoutPreferencesOutput() *CheckoutPreferencesOutput`

NewCheckoutPreferencesOutput instantiates a new CheckoutPreferencesOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckoutPreferencesOutputWithDefaults

`func NewCheckoutPreferencesOutputWithDefaults() *CheckoutPreferencesOutput`

NewCheckoutPreferencesOutputWithDefaults instantiates a new CheckoutPreferencesOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFallbackCascadePreferences

`func (o *CheckoutPreferencesOutput) GetFallbackCascadePreferences() []FallbackConfigurationOutput`

GetFallbackCascadePreferences returns the FallbackCascadePreferences field if non-nil, zero value otherwise.

### GetFallbackCascadePreferencesOk

`func (o *CheckoutPreferencesOutput) GetFallbackCascadePreferencesOk() (*[]FallbackConfigurationOutput, bool)`

GetFallbackCascadePreferencesOk returns a tuple with the FallbackCascadePreferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackCascadePreferences

`func (o *CheckoutPreferencesOutput) SetFallbackCascadePreferences(v []FallbackConfigurationOutput)`

SetFallbackCascadePreferences sets FallbackCascadePreferences field to given value.

### HasFallbackCascadePreferences

`func (o *CheckoutPreferencesOutput) HasFallbackCascadePreferences() bool`

HasFallbackCascadePreferences returns a boolean if a field has been set.

### GetFallbackCascadeSelectedProductPriceQuantity

`func (o *CheckoutPreferencesOutput) GetFallbackCascadeSelectedProductPriceQuantity() [][]SelectedPriceQuantity`

GetFallbackCascadeSelectedProductPriceQuantity returns the FallbackCascadeSelectedProductPriceQuantity field if non-nil, zero value otherwise.

### GetFallbackCascadeSelectedProductPriceQuantityOk

`func (o *CheckoutPreferencesOutput) GetFallbackCascadeSelectedProductPriceQuantityOk() (*[][]SelectedPriceQuantity, bool)`

GetFallbackCascadeSelectedProductPriceQuantityOk returns a tuple with the FallbackCascadeSelectedProductPriceQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackCascadeSelectedProductPriceQuantity

`func (o *CheckoutPreferencesOutput) SetFallbackCascadeSelectedProductPriceQuantity(v [][]SelectedPriceQuantity)`

SetFallbackCascadeSelectedProductPriceQuantity sets FallbackCascadeSelectedProductPriceQuantity field to given value.

### HasFallbackCascadeSelectedProductPriceQuantity

`func (o *CheckoutPreferencesOutput) HasFallbackCascadeSelectedProductPriceQuantity() bool`

HasFallbackCascadeSelectedProductPriceQuantity returns a boolean if a field has been set.

### GetOverrideFeatureRolloutSettings

`func (o *CheckoutPreferencesOutput) GetOverrideFeatureRolloutSettings() map[string]interface{}`

GetOverrideFeatureRolloutSettings returns the OverrideFeatureRolloutSettings field if non-nil, zero value otherwise.

### GetOverrideFeatureRolloutSettingsOk

`func (o *CheckoutPreferencesOutput) GetOverrideFeatureRolloutSettingsOk() (*map[string]interface{}, bool)`

GetOverrideFeatureRolloutSettingsOk returns a tuple with the OverrideFeatureRolloutSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideFeatureRolloutSettings

`func (o *CheckoutPreferencesOutput) SetOverrideFeatureRolloutSettings(v map[string]interface{})`

SetOverrideFeatureRolloutSettings sets OverrideFeatureRolloutSettings field to given value.

### HasOverrideFeatureRolloutSettings

`func (o *CheckoutPreferencesOutput) HasOverrideFeatureRolloutSettings() bool`

HasOverrideFeatureRolloutSettings returns a boolean if a field has been set.

### SetOverrideFeatureRolloutSettingsNil

`func (o *CheckoutPreferencesOutput) SetOverrideFeatureRolloutSettingsNil(b bool)`

 SetOverrideFeatureRolloutSettingsNil sets the value for OverrideFeatureRolloutSettings to be an explicit nil

### UnsetOverrideFeatureRolloutSettings
`func (o *CheckoutPreferencesOutput) UnsetOverrideFeatureRolloutSettings()`

UnsetOverrideFeatureRolloutSettings ensures that no value is present for OverrideFeatureRolloutSettings, not even an explicit nil
### GetProcessorPreferences

`func (o *CheckoutPreferencesOutput) GetProcessorPreferences() CheckoutProcessorsPreferences`

GetProcessorPreferences returns the ProcessorPreferences field if non-nil, zero value otherwise.

### GetProcessorPreferencesOk

`func (o *CheckoutPreferencesOutput) GetProcessorPreferencesOk() (*CheckoutProcessorsPreferences, bool)`

GetProcessorPreferencesOk returns a tuple with the ProcessorPreferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessorPreferences

`func (o *CheckoutPreferencesOutput) SetProcessorPreferences(v CheckoutProcessorsPreferences)`

SetProcessorPreferences sets ProcessorPreferences field to given value.

### HasProcessorPreferences

`func (o *CheckoutPreferencesOutput) HasProcessorPreferences() bool`

HasProcessorPreferences returns a boolean if a field has been set.

### SetProcessorPreferencesNil

`func (o *CheckoutPreferencesOutput) SetProcessorPreferencesNil(b bool)`

 SetProcessorPreferencesNil sets the value for ProcessorPreferences to be an explicit nil

### UnsetProcessorPreferences
`func (o *CheckoutPreferencesOutput) UnsetProcessorPreferences()`

UnsetProcessorPreferences ensures that no value is present for ProcessorPreferences, not even an explicit nil
### GetUnused

`func (o *CheckoutPreferencesOutput) GetUnused() string`

GetUnused returns the Unused field if non-nil, zero value otherwise.

### GetUnusedOk

`func (o *CheckoutPreferencesOutput) GetUnusedOk() (*string, bool)`

GetUnusedOk returns a tuple with the Unused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused

`func (o *CheckoutPreferencesOutput) SetUnused(v string)`

SetUnused sets Unused field to given value.

### HasUnused

`func (o *CheckoutPreferencesOutput) HasUnused() bool`

HasUnused returns a boolean if a field has been set.

### SetUnusedNil

`func (o *CheckoutPreferencesOutput) SetUnusedNil(b bool)`

 SetUnusedNil sets the value for Unused to be an explicit nil

### UnsetUnused
`func (o *CheckoutPreferencesOutput) UnsetUnused()`

UnsetUnused ensures that no value is present for Unused, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


