# AppModelsCustomersInvoiceSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier for the object. | 
**Object** | [**ObjectName**](ObjectName.md) |  | 
**CreatedAt** | **time.Time** | DateTime at which the object was created, in &#39;ISO 8601&#39; format. | 
**UpdatedAt** | **time.Time** | DateTime at which the object was updated, in &#39;ISO 8601&#39; format. | 
**IsDeleted** | Pointer to **bool** | If true, indicates that this object has been deleted | [optional] [default to false]
**EmailReceiptOnPaid** | Pointer to **NullableBool** |  | [optional] 
**DefaultNetD** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewAppModelsCustomersInvoiceSettings

`func NewAppModelsCustomersInvoiceSettings(id string, object ObjectName, createdAt time.Time, updatedAt time.Time, ) *AppModelsCustomersInvoiceSettings`

NewAppModelsCustomersInvoiceSettings instantiates a new AppModelsCustomersInvoiceSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppModelsCustomersInvoiceSettingsWithDefaults

`func NewAppModelsCustomersInvoiceSettingsWithDefaults() *AppModelsCustomersInvoiceSettings`

NewAppModelsCustomersInvoiceSettingsWithDefaults instantiates a new AppModelsCustomersInvoiceSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AppModelsCustomersInvoiceSettings) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppModelsCustomersInvoiceSettings) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppModelsCustomersInvoiceSettings) SetId(v string)`

SetId sets Id field to given value.


### GetObject

`func (o *AppModelsCustomersInvoiceSettings) GetObject() ObjectName`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *AppModelsCustomersInvoiceSettings) GetObjectOk() (*ObjectName, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *AppModelsCustomersInvoiceSettings) SetObject(v ObjectName)`

SetObject sets Object field to given value.


### GetCreatedAt

`func (o *AppModelsCustomersInvoiceSettings) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AppModelsCustomersInvoiceSettings) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AppModelsCustomersInvoiceSettings) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *AppModelsCustomersInvoiceSettings) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AppModelsCustomersInvoiceSettings) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AppModelsCustomersInvoiceSettings) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetIsDeleted

`func (o *AppModelsCustomersInvoiceSettings) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *AppModelsCustomersInvoiceSettings) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *AppModelsCustomersInvoiceSettings) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *AppModelsCustomersInvoiceSettings) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetEmailReceiptOnPaid

`func (o *AppModelsCustomersInvoiceSettings) GetEmailReceiptOnPaid() bool`

GetEmailReceiptOnPaid returns the EmailReceiptOnPaid field if non-nil, zero value otherwise.

### GetEmailReceiptOnPaidOk

`func (o *AppModelsCustomersInvoiceSettings) GetEmailReceiptOnPaidOk() (*bool, bool)`

GetEmailReceiptOnPaidOk returns a tuple with the EmailReceiptOnPaid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailReceiptOnPaid

`func (o *AppModelsCustomersInvoiceSettings) SetEmailReceiptOnPaid(v bool)`

SetEmailReceiptOnPaid sets EmailReceiptOnPaid field to given value.

### HasEmailReceiptOnPaid

`func (o *AppModelsCustomersInvoiceSettings) HasEmailReceiptOnPaid() bool`

HasEmailReceiptOnPaid returns a boolean if a field has been set.

### SetEmailReceiptOnPaidNil

`func (o *AppModelsCustomersInvoiceSettings) SetEmailReceiptOnPaidNil(b bool)`

 SetEmailReceiptOnPaidNil sets the value for EmailReceiptOnPaid to be an explicit nil

### UnsetEmailReceiptOnPaid
`func (o *AppModelsCustomersInvoiceSettings) UnsetEmailReceiptOnPaid()`

UnsetEmailReceiptOnPaid ensures that no value is present for EmailReceiptOnPaid, not even an explicit nil
### GetDefaultNetD

`func (o *AppModelsCustomersInvoiceSettings) GetDefaultNetD() int32`

GetDefaultNetD returns the DefaultNetD field if non-nil, zero value otherwise.

### GetDefaultNetDOk

`func (o *AppModelsCustomersInvoiceSettings) GetDefaultNetDOk() (*int32, bool)`

GetDefaultNetDOk returns a tuple with the DefaultNetD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultNetD

`func (o *AppModelsCustomersInvoiceSettings) SetDefaultNetD(v int32)`

SetDefaultNetD sets DefaultNetD field to given value.

### HasDefaultNetD

`func (o *AppModelsCustomersInvoiceSettings) HasDefaultNetD() bool`

HasDefaultNetD returns a boolean if a field has been set.

### SetDefaultNetDNil

`func (o *AppModelsCustomersInvoiceSettings) SetDefaultNetDNil(b bool)`

 SetDefaultNetDNil sets the value for DefaultNetD to be an explicit nil

### UnsetDefaultNetD
`func (o *AppModelsCustomersInvoiceSettings) UnsetDefaultNetD()`

UnsetDefaultNetD ensures that no value is present for DefaultNetD, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


