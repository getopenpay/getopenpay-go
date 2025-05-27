# UpdatePaymentRouteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the payment route. | 
**RouteConfiguration** | **map[string]interface{}** | The configuration object for the payment route. Contact the OpenPay team for more information. | 

## Methods

### NewUpdatePaymentRouteRequest

`func NewUpdatePaymentRouteRequest(name string, routeConfiguration map[string]interface{}, ) *UpdatePaymentRouteRequest`

NewUpdatePaymentRouteRequest instantiates a new UpdatePaymentRouteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePaymentRouteRequestWithDefaults

`func NewUpdatePaymentRouteRequestWithDefaults() *UpdatePaymentRouteRequest`

NewUpdatePaymentRouteRequestWithDefaults instantiates a new UpdatePaymentRouteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdatePaymentRouteRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdatePaymentRouteRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdatePaymentRouteRequest) SetName(v string)`

SetName sets Name field to given value.


### GetRouteConfiguration

`func (o *UpdatePaymentRouteRequest) GetRouteConfiguration() map[string]interface{}`

GetRouteConfiguration returns the RouteConfiguration field if non-nil, zero value otherwise.

### GetRouteConfigurationOk

`func (o *UpdatePaymentRouteRequest) GetRouteConfigurationOk() (*map[string]interface{}, bool)`

GetRouteConfigurationOk returns a tuple with the RouteConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteConfiguration

`func (o *UpdatePaymentRouteRequest) SetRouteConfiguration(v map[string]interface{})`

SetRouteConfiguration sets RouteConfiguration field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


