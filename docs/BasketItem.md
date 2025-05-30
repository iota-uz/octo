# BasketItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PositionDesc** | **string** |  | 
**Count** | **int32** |  | 
**Price** | **float64** |  | 
**Spic** | Pointer to **string** |  | [optional] 

## Methods

### NewBasketItem

`func NewBasketItem(positionDesc string, count int32, price float64, ) *BasketItem`

NewBasketItem instantiates a new BasketItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBasketItemWithDefaults

`func NewBasketItemWithDefaults() *BasketItem`

NewBasketItemWithDefaults instantiates a new BasketItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPositionDesc

`func (o *BasketItem) GetPositionDesc() string`

GetPositionDesc returns the PositionDesc field if non-nil, zero value otherwise.

### GetPositionDescOk

`func (o *BasketItem) GetPositionDescOk() (*string, bool)`

GetPositionDescOk returns a tuple with the PositionDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionDesc

`func (o *BasketItem) SetPositionDesc(v string)`

SetPositionDesc sets PositionDesc field to given value.


### GetCount

`func (o *BasketItem) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *BasketItem) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *BasketItem) SetCount(v int32)`

SetCount sets Count field to given value.


### GetPrice

`func (o *BasketItem) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *BasketItem) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *BasketItem) SetPrice(v float64)`

SetPrice sets Price field to given value.


### GetSpic

`func (o *BasketItem) GetSpic() string`

GetSpic returns the Spic field if non-nil, zero value otherwise.

### GetSpicOk

`func (o *BasketItem) GetSpicOk() (*string, bool)`

GetSpicOk returns a tuple with the Spic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpic

`func (o *BasketItem) SetSpic(v string)`

SetSpic sets Spic field to given value.

### HasSpic

`func (o *BasketItem) HasSpic() bool`

HasSpic returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


