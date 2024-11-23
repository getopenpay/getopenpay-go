# EventExternal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Object** | Pointer to [**ObjectName**](ObjectName.md) |  | [optional] [default to OBJECTNAME_EVENT]
**CreatedAt** | **time.Time** | DateTime at which the object was created, in &#39;ISO 8601&#39; format. | 
**UpdatedAt** | **time.Time** | DateTime at which the object was updated, in &#39;ISO 8601&#39; format. | 
**IsDeleted** | Pointer to **bool** | If true, indicates that this object has been deleted | [optional] [default to false]
**AccountId** | **NullableString** |  | 
**Type** | [**EventType**](EventType.md) |  | 
**Data** | **map[string]interface{}** |  | 
**DataPrevious** | **map[string]interface{}** |  | 
**RequestId** | **NullableString** |  | 
**RequestIdempotencyKey** | **NullableString** |  | 
**PendingWebhooks** | **int32** |  | 
**User** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewEventExternal

`func NewEventExternal(id string, createdAt time.Time, updatedAt time.Time, accountId NullableString, type_ EventType, data map[string]interface{}, dataPrevious map[string]interface{}, requestId NullableString, requestIdempotencyKey NullableString, pendingWebhooks int32, ) *EventExternal`

NewEventExternal instantiates a new EventExternal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventExternalWithDefaults

`func NewEventExternalWithDefaults() *EventExternal`

NewEventExternalWithDefaults instantiates a new EventExternal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EventExternal) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EventExternal) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EventExternal) SetId(v string)`

SetId sets Id field to given value.


### GetObject

`func (o *EventExternal) GetObject() ObjectName`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *EventExternal) GetObjectOk() (*ObjectName, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *EventExternal) SetObject(v ObjectName)`

SetObject sets Object field to given value.

### HasObject

`func (o *EventExternal) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetCreatedAt

`func (o *EventExternal) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EventExternal) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EventExternal) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *EventExternal) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *EventExternal) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *EventExternal) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetIsDeleted

`func (o *EventExternal) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *EventExternal) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *EventExternal) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *EventExternal) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetAccountId

`func (o *EventExternal) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *EventExternal) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *EventExternal) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### SetAccountIdNil

`func (o *EventExternal) SetAccountIdNil(b bool)`

 SetAccountIdNil sets the value for AccountId to be an explicit nil

### UnsetAccountId
`func (o *EventExternal) UnsetAccountId()`

UnsetAccountId ensures that no value is present for AccountId, not even an explicit nil
### GetType

`func (o *EventExternal) GetType() EventType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EventExternal) GetTypeOk() (*EventType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EventExternal) SetType(v EventType)`

SetType sets Type field to given value.


### GetData

`func (o *EventExternal) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *EventExternal) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *EventExternal) SetData(v map[string]interface{})`

SetData sets Data field to given value.


### GetDataPrevious

`func (o *EventExternal) GetDataPrevious() map[string]interface{}`

GetDataPrevious returns the DataPrevious field if non-nil, zero value otherwise.

### GetDataPreviousOk

`func (o *EventExternal) GetDataPreviousOk() (*map[string]interface{}, bool)`

GetDataPreviousOk returns a tuple with the DataPrevious field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataPrevious

`func (o *EventExternal) SetDataPrevious(v map[string]interface{})`

SetDataPrevious sets DataPrevious field to given value.


### SetDataPreviousNil

`func (o *EventExternal) SetDataPreviousNil(b bool)`

 SetDataPreviousNil sets the value for DataPrevious to be an explicit nil

### UnsetDataPrevious
`func (o *EventExternal) UnsetDataPrevious()`

UnsetDataPrevious ensures that no value is present for DataPrevious, not even an explicit nil
### GetRequestId

`func (o *EventExternal) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *EventExternal) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *EventExternal) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### SetRequestIdNil

`func (o *EventExternal) SetRequestIdNil(b bool)`

 SetRequestIdNil sets the value for RequestId to be an explicit nil

### UnsetRequestId
`func (o *EventExternal) UnsetRequestId()`

UnsetRequestId ensures that no value is present for RequestId, not even an explicit nil
### GetRequestIdempotencyKey

`func (o *EventExternal) GetRequestIdempotencyKey() string`

GetRequestIdempotencyKey returns the RequestIdempotencyKey field if non-nil, zero value otherwise.

### GetRequestIdempotencyKeyOk

`func (o *EventExternal) GetRequestIdempotencyKeyOk() (*string, bool)`

GetRequestIdempotencyKeyOk returns a tuple with the RequestIdempotencyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestIdempotencyKey

`func (o *EventExternal) SetRequestIdempotencyKey(v string)`

SetRequestIdempotencyKey sets RequestIdempotencyKey field to given value.


### SetRequestIdempotencyKeyNil

`func (o *EventExternal) SetRequestIdempotencyKeyNil(b bool)`

 SetRequestIdempotencyKeyNil sets the value for RequestIdempotencyKey to be an explicit nil

### UnsetRequestIdempotencyKey
`func (o *EventExternal) UnsetRequestIdempotencyKey()`

UnsetRequestIdempotencyKey ensures that no value is present for RequestIdempotencyKey, not even an explicit nil
### GetPendingWebhooks

`func (o *EventExternal) GetPendingWebhooks() int32`

GetPendingWebhooks returns the PendingWebhooks field if non-nil, zero value otherwise.

### GetPendingWebhooksOk

`func (o *EventExternal) GetPendingWebhooksOk() (*int32, bool)`

GetPendingWebhooksOk returns a tuple with the PendingWebhooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingWebhooks

`func (o *EventExternal) SetPendingWebhooks(v int32)`

SetPendingWebhooks sets PendingWebhooks field to given value.


### GetUser

`func (o *EventExternal) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *EventExternal) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *EventExternal) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *EventExternal) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *EventExternal) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *EventExternal) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


