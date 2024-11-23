# CustomerTotalAmount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | [**CurrencyEnum**](CurrencyEnum.md) |  | 
**AmountAtom** | Pointer to **int32** | The total refunds/spent of the customer. | [optional] [default to 0]

## Methods

### NewCustomerTotalAmount

`func NewCustomerTotalAmount(currency CurrencyEnum, ) *CustomerTotalAmount`

NewCustomerTotalAmount instantiates a new CustomerTotalAmount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerTotalAmountWithDefaults

`func NewCustomerTotalAmountWithDefaults() *CustomerTotalAmount`

NewCustomerTotalAmountWithDefaults instantiates a new CustomerTotalAmount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *CustomerTotalAmount) GetCurrency() CurrencyEnum`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CustomerTotalAmount) GetCurrencyOk() (*CurrencyEnum, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CustomerTotalAmount) SetCurrency(v CurrencyEnum)`

SetCurrency sets Currency field to given value.


### GetAmountAtom

`func (o *CustomerTotalAmount) GetAmountAtom() int32`

GetAmountAtom returns the AmountAtom field if non-nil, zero value otherwise.

### GetAmountAtomOk

`func (o *CustomerTotalAmount) GetAmountAtomOk() (*int32, bool)`

GetAmountAtomOk returns a tuple with the AmountAtom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountAtom

`func (o *CustomerTotalAmount) SetAmountAtom(v int32)`

SetAmountAtom sets AmountAtom field to given value.

### HasAmountAtom

`func (o *CustomerTotalAmount) HasAmountAtom() bool`

HasAmountAtom returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


