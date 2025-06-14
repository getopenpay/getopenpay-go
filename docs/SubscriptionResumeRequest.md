# SubscriptionResumeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsPreview** | Pointer to **bool** | Whether the request is in preview mode (subscriptions won&#39;t actually be resumed) | [optional] [default to false]
**ProrationBehavior** | Pointer to [**ProrationEnum**](ProrationEnum.md) | Determines how to handle prorations when the subscription is resumed. | [optional] 

## Methods

### NewSubscriptionResumeRequest

`func NewSubscriptionResumeRequest() *SubscriptionResumeRequest`

NewSubscriptionResumeRequest instantiates a new SubscriptionResumeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionResumeRequestWithDefaults

`func NewSubscriptionResumeRequestWithDefaults() *SubscriptionResumeRequest`

NewSubscriptionResumeRequestWithDefaults instantiates a new SubscriptionResumeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsPreview

`func (o *SubscriptionResumeRequest) GetIsPreview() bool`

GetIsPreview returns the IsPreview field if non-nil, zero value otherwise.

### GetIsPreviewOk

`func (o *SubscriptionResumeRequest) GetIsPreviewOk() (*bool, bool)`

GetIsPreviewOk returns a tuple with the IsPreview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPreview

`func (o *SubscriptionResumeRequest) SetIsPreview(v bool)`

SetIsPreview sets IsPreview field to given value.

### HasIsPreview

`func (o *SubscriptionResumeRequest) HasIsPreview() bool`

HasIsPreview returns a boolean if a field has been set.

### GetProrationBehavior

`func (o *SubscriptionResumeRequest) GetProrationBehavior() ProrationEnum`

GetProrationBehavior returns the ProrationBehavior field if non-nil, zero value otherwise.

### GetProrationBehaviorOk

`func (o *SubscriptionResumeRequest) GetProrationBehaviorOk() (*ProrationEnum, bool)`

GetProrationBehaviorOk returns a tuple with the ProrationBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProrationBehavior

`func (o *SubscriptionResumeRequest) SetProrationBehavior(v ProrationEnum)`

SetProrationBehavior sets ProrationBehavior field to given value.

### HasProrationBehavior

`func (o *SubscriptionResumeRequest) HasProrationBehavior() bool`

HasProrationBehavior returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


