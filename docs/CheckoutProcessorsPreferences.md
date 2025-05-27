# CheckoutProcessorsPreferences

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilterProviders** | Pointer to [**[]PaymentProviderType**](PaymentProviderType.md) | If not empty, the enabled providers will be filtered by this list. | [optional] [default to []]
**IdsWhitelist** | Pointer to **[]string** | If not empty, only the processors with these ids will be accepted. | [optional] [default to []]
**NamesWhitelist** | Pointer to **[]string** | If not empty, only the processors with these names will be accepted. Valid values are: adyen, airwallex, authorize_net, checkout_com, braintree, stripe, foobar, pockyt, cybersource, loop, paypal, nmi. | [optional] [default to []]

## Methods

### NewCheckoutProcessorsPreferences

`func NewCheckoutProcessorsPreferences() *CheckoutProcessorsPreferences`

NewCheckoutProcessorsPreferences instantiates a new CheckoutProcessorsPreferences object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckoutProcessorsPreferencesWithDefaults

`func NewCheckoutProcessorsPreferencesWithDefaults() *CheckoutProcessorsPreferences`

NewCheckoutProcessorsPreferencesWithDefaults instantiates a new CheckoutProcessorsPreferences object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilterProviders

`func (o *CheckoutProcessorsPreferences) GetFilterProviders() []PaymentProviderType`

GetFilterProviders returns the FilterProviders field if non-nil, zero value otherwise.

### GetFilterProvidersOk

`func (o *CheckoutProcessorsPreferences) GetFilterProvidersOk() (*[]PaymentProviderType, bool)`

GetFilterProvidersOk returns a tuple with the FilterProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterProviders

`func (o *CheckoutProcessorsPreferences) SetFilterProviders(v []PaymentProviderType)`

SetFilterProviders sets FilterProviders field to given value.

### HasFilterProviders

`func (o *CheckoutProcessorsPreferences) HasFilterProviders() bool`

HasFilterProviders returns a boolean if a field has been set.

### GetIdsWhitelist

`func (o *CheckoutProcessorsPreferences) GetIdsWhitelist() []string`

GetIdsWhitelist returns the IdsWhitelist field if non-nil, zero value otherwise.

### GetIdsWhitelistOk

`func (o *CheckoutProcessorsPreferences) GetIdsWhitelistOk() (*[]string, bool)`

GetIdsWhitelistOk returns a tuple with the IdsWhitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdsWhitelist

`func (o *CheckoutProcessorsPreferences) SetIdsWhitelist(v []string)`

SetIdsWhitelist sets IdsWhitelist field to given value.

### HasIdsWhitelist

`func (o *CheckoutProcessorsPreferences) HasIdsWhitelist() bool`

HasIdsWhitelist returns a boolean if a field has been set.

### GetNamesWhitelist

`func (o *CheckoutProcessorsPreferences) GetNamesWhitelist() []string`

GetNamesWhitelist returns the NamesWhitelist field if non-nil, zero value otherwise.

### GetNamesWhitelistOk

`func (o *CheckoutProcessorsPreferences) GetNamesWhitelistOk() (*[]string, bool)`

GetNamesWhitelistOk returns a tuple with the NamesWhitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamesWhitelist

`func (o *CheckoutProcessorsPreferences) SetNamesWhitelist(v []string)`

SetNamesWhitelist sets NamesWhitelist field to given value.

### HasNamesWhitelist

`func (o *CheckoutProcessorsPreferences) HasNamesWhitelist() bool`

HasNamesWhitelist returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


