# CheckoutPreferences

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FallbackCascadePreferences** | Pointer to [**[]FallbackConfigurationInput**](FallbackConfigurationInput.md) | This object contains a list of price groups that will be used to fall back to if the selected product price quantity fails to be created. Processor preferences here will override the processor preferences in the processor_preferences field. | [optional] [default to []]
**FallbackCascadeSelectedProductPriceQuantity** | Pointer to [**[][]SelectedPriceQuantity**]([]SelectedPriceQuantity.md) | This object contains a list of price groups that will be used to fall back to if the selected product price quantity fails to be created. | [optional] [default to []]
**OverrideFeatureRolloutSettings** | Pointer to **map[string]interface{}** |  | [optional] 
**ProcessorPreferences** | Pointer to [**NullableCheckoutProcessorsPreferences**](CheckoutProcessorsPreferences.md) |  | [optional] 

## Methods

### NewCheckoutPreferences

`func NewCheckoutPreferences() *CheckoutPreferences`

NewCheckoutPreferences instantiates a new CheckoutPreferences object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckoutPreferencesWithDefaults

`func NewCheckoutPreferencesWithDefaults() *CheckoutPreferences`

NewCheckoutPreferencesWithDefaults instantiates a new CheckoutPreferences object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFallbackCascadePreferences

`func (o *CheckoutPreferences) GetFallbackCascadePreferences() []FallbackConfigurationInput`

GetFallbackCascadePreferences returns the FallbackCascadePreferences field if non-nil, zero value otherwise.

### GetFallbackCascadePreferencesOk

`func (o *CheckoutPreferences) GetFallbackCascadePreferencesOk() (*[]FallbackConfigurationInput, bool)`

GetFallbackCascadePreferencesOk returns a tuple with the FallbackCascadePreferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackCascadePreferences

`func (o *CheckoutPreferences) SetFallbackCascadePreferences(v []FallbackConfigurationInput)`

SetFallbackCascadePreferences sets FallbackCascadePreferences field to given value.

### HasFallbackCascadePreferences

`func (o *CheckoutPreferences) HasFallbackCascadePreferences() bool`

HasFallbackCascadePreferences returns a boolean if a field has been set.

### GetFallbackCascadeSelectedProductPriceQuantity

`func (o *CheckoutPreferences) GetFallbackCascadeSelectedProductPriceQuantity() [][]SelectedPriceQuantity`

GetFallbackCascadeSelectedProductPriceQuantity returns the FallbackCascadeSelectedProductPriceQuantity field if non-nil, zero value otherwise.

### GetFallbackCascadeSelectedProductPriceQuantityOk

`func (o *CheckoutPreferences) GetFallbackCascadeSelectedProductPriceQuantityOk() (*[][]SelectedPriceQuantity, bool)`

GetFallbackCascadeSelectedProductPriceQuantityOk returns a tuple with the FallbackCascadeSelectedProductPriceQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackCascadeSelectedProductPriceQuantity

`func (o *CheckoutPreferences) SetFallbackCascadeSelectedProductPriceQuantity(v [][]SelectedPriceQuantity)`

SetFallbackCascadeSelectedProductPriceQuantity sets FallbackCascadeSelectedProductPriceQuantity field to given value.

### HasFallbackCascadeSelectedProductPriceQuantity

`func (o *CheckoutPreferences) HasFallbackCascadeSelectedProductPriceQuantity() bool`

HasFallbackCascadeSelectedProductPriceQuantity returns a boolean if a field has been set.

### GetOverrideFeatureRolloutSettings

`func (o *CheckoutPreferences) GetOverrideFeatureRolloutSettings() map[string]interface{}`

GetOverrideFeatureRolloutSettings returns the OverrideFeatureRolloutSettings field if non-nil, zero value otherwise.

### GetOverrideFeatureRolloutSettingsOk

`func (o *CheckoutPreferences) GetOverrideFeatureRolloutSettingsOk() (*map[string]interface{}, bool)`

GetOverrideFeatureRolloutSettingsOk returns a tuple with the OverrideFeatureRolloutSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideFeatureRolloutSettings

`func (o *CheckoutPreferences) SetOverrideFeatureRolloutSettings(v map[string]interface{})`

SetOverrideFeatureRolloutSettings sets OverrideFeatureRolloutSettings field to given value.

### HasOverrideFeatureRolloutSettings

`func (o *CheckoutPreferences) HasOverrideFeatureRolloutSettings() bool`

HasOverrideFeatureRolloutSettings returns a boolean if a field has been set.

### SetOverrideFeatureRolloutSettingsNil

`func (o *CheckoutPreferences) SetOverrideFeatureRolloutSettingsNil(b bool)`

 SetOverrideFeatureRolloutSettingsNil sets the value for OverrideFeatureRolloutSettings to be an explicit nil

### UnsetOverrideFeatureRolloutSettings
`func (o *CheckoutPreferences) UnsetOverrideFeatureRolloutSettings()`

UnsetOverrideFeatureRolloutSettings ensures that no value is present for OverrideFeatureRolloutSettings, not even an explicit nil
### GetProcessorPreferences

`func (o *CheckoutPreferences) GetProcessorPreferences() CheckoutProcessorsPreferences`

GetProcessorPreferences returns the ProcessorPreferences field if non-nil, zero value otherwise.

### GetProcessorPreferencesOk

`func (o *CheckoutPreferences) GetProcessorPreferencesOk() (*CheckoutProcessorsPreferences, bool)`

GetProcessorPreferencesOk returns a tuple with the ProcessorPreferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessorPreferences

`func (o *CheckoutPreferences) SetProcessorPreferences(v CheckoutProcessorsPreferences)`

SetProcessorPreferences sets ProcessorPreferences field to given value.

### HasProcessorPreferences

`func (o *CheckoutPreferences) HasProcessorPreferences() bool`

HasProcessorPreferences returns a boolean if a field has been set.

### SetProcessorPreferencesNil

`func (o *CheckoutPreferences) SetProcessorPreferencesNil(b bool)`

 SetProcessorPreferencesNil sets the value for ProcessorPreferences to be an explicit nil

### UnsetProcessorPreferences
`func (o *CheckoutPreferences) UnsetProcessorPreferences()`

UnsetProcessorPreferences ensures that no value is present for ProcessorPreferences, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


