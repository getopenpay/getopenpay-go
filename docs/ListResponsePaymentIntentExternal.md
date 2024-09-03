# ListResponsePaymentIntentExternal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]PaymentIntentExternal**](PaymentIntentExternal.md) |  | 
**TotalObjects** | **int32** |  | 
**PageNumber** | **int32** |  | 
**PageSize** | **int32** |  | 

## Methods

### NewListResponsePaymentIntentExternal

`func NewListResponsePaymentIntentExternal(data []PaymentIntentExternal, totalObjects int32, pageNumber int32, pageSize int32, ) *ListResponsePaymentIntentExternal`

NewListResponsePaymentIntentExternal instantiates a new ListResponsePaymentIntentExternal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListResponsePaymentIntentExternalWithDefaults

`func NewListResponsePaymentIntentExternalWithDefaults() *ListResponsePaymentIntentExternal`

NewListResponsePaymentIntentExternalWithDefaults instantiates a new ListResponsePaymentIntentExternal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListResponsePaymentIntentExternal) GetData() []PaymentIntentExternal`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListResponsePaymentIntentExternal) GetDataOk() (*[]PaymentIntentExternal, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListResponsePaymentIntentExternal) SetData(v []PaymentIntentExternal)`

SetData sets Data field to given value.


### GetTotalObjects

`func (o *ListResponsePaymentIntentExternal) GetTotalObjects() int32`

GetTotalObjects returns the TotalObjects field if non-nil, zero value otherwise.

### GetTotalObjectsOk

`func (o *ListResponsePaymentIntentExternal) GetTotalObjectsOk() (*int32, bool)`

GetTotalObjectsOk returns a tuple with the TotalObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalObjects

`func (o *ListResponsePaymentIntentExternal) SetTotalObjects(v int32)`

SetTotalObjects sets TotalObjects field to given value.


### GetPageNumber

`func (o *ListResponsePaymentIntentExternal) GetPageNumber() int32`

GetPageNumber returns the PageNumber field if non-nil, zero value otherwise.

### GetPageNumberOk

`func (o *ListResponsePaymentIntentExternal) GetPageNumberOk() (*int32, bool)`

GetPageNumberOk returns a tuple with the PageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumber

`func (o *ListResponsePaymentIntentExternal) SetPageNumber(v int32)`

SetPageNumber sets PageNumber field to given value.


### GetPageSize

`func (o *ListResponsePaymentIntentExternal) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *ListResponsePaymentIntentExternal) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *ListResponsePaymentIntentExternal) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

