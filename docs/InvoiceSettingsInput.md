# InvoiceSettingsInput

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

### NewInvoiceSettingsInput

`func NewInvoiceSettingsInput(id string, object ObjectName, createdAt time.Time, updatedAt time.Time, ) *InvoiceSettingsInput`

NewInvoiceSettingsInput instantiates a new InvoiceSettingsInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceSettingsInputWithDefaults

`func NewInvoiceSettingsInputWithDefaults() *InvoiceSettingsInput`

NewInvoiceSettingsInputWithDefaults instantiates a new InvoiceSettingsInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InvoiceSettingsInput) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InvoiceSettingsInput) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InvoiceSettingsInput) SetId(v string)`

SetId sets Id field to given value.


### GetObject

`func (o *InvoiceSettingsInput) GetObject() ObjectName`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *InvoiceSettingsInput) GetObjectOk() (*ObjectName, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *InvoiceSettingsInput) SetObject(v ObjectName)`

SetObject sets Object field to given value.


### GetCreatedAt

`func (o *InvoiceSettingsInput) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InvoiceSettingsInput) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InvoiceSettingsInput) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *InvoiceSettingsInput) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *InvoiceSettingsInput) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *InvoiceSettingsInput) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetIsDeleted

`func (o *InvoiceSettingsInput) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *InvoiceSettingsInput) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *InvoiceSettingsInput) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *InvoiceSettingsInput) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetEmailReceiptOnPaid

`func (o *InvoiceSettingsInput) GetEmailReceiptOnPaid() bool`

GetEmailReceiptOnPaid returns the EmailReceiptOnPaid field if non-nil, zero value otherwise.

### GetEmailReceiptOnPaidOk

`func (o *InvoiceSettingsInput) GetEmailReceiptOnPaidOk() (*bool, bool)`

GetEmailReceiptOnPaidOk returns a tuple with the EmailReceiptOnPaid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailReceiptOnPaid

`func (o *InvoiceSettingsInput) SetEmailReceiptOnPaid(v bool)`

SetEmailReceiptOnPaid sets EmailReceiptOnPaid field to given value.

### HasEmailReceiptOnPaid

`func (o *InvoiceSettingsInput) HasEmailReceiptOnPaid() bool`

HasEmailReceiptOnPaid returns a boolean if a field has been set.

### SetEmailReceiptOnPaidNil

`func (o *InvoiceSettingsInput) SetEmailReceiptOnPaidNil(b bool)`

 SetEmailReceiptOnPaidNil sets the value for EmailReceiptOnPaid to be an explicit nil

### UnsetEmailReceiptOnPaid
`func (o *InvoiceSettingsInput) UnsetEmailReceiptOnPaid()`

UnsetEmailReceiptOnPaid ensures that no value is present for EmailReceiptOnPaid, not even an explicit nil
### GetDefaultNetD

`func (o *InvoiceSettingsInput) GetDefaultNetD() int32`

GetDefaultNetD returns the DefaultNetD field if non-nil, zero value otherwise.

### GetDefaultNetDOk

`func (o *InvoiceSettingsInput) GetDefaultNetDOk() (*int32, bool)`

GetDefaultNetDOk returns a tuple with the DefaultNetD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultNetD

`func (o *InvoiceSettingsInput) SetDefaultNetD(v int32)`

SetDefaultNetD sets DefaultNetD field to given value.

### HasDefaultNetD

`func (o *InvoiceSettingsInput) HasDefaultNetD() bool`

HasDefaultNetD returns a boolean if a field has been set.

### SetDefaultNetDNil

`func (o *InvoiceSettingsInput) SetDefaultNetDNil(b bool)`

 SetDefaultNetDNil sets the value for DefaultNetD to be an explicit nil

### UnsetDefaultNetD
`func (o *InvoiceSettingsInput) UnsetDefaultNetD()`

UnsetDefaultNetD ensures that no value is present for DefaultNetD, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


