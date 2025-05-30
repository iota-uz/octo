# RefundRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OctoShopId** | **int32** |  | 
**ShopRefundId** | **string** |  | 
**OctoSecret** | **string** |  | 
**OctoPaymentUUID** | **string** |  | 
**Amount** | **float64** |  | 

## Methods

### NewRefundRequest

`func NewRefundRequest(octoShopId int32, shopRefundId string, octoSecret string, octoPaymentUUID string, amount float64, ) *RefundRequest`

NewRefundRequest instantiates a new RefundRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRefundRequestWithDefaults

`func NewRefundRequestWithDefaults() *RefundRequest`

NewRefundRequestWithDefaults instantiates a new RefundRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOctoShopId

`func (o *RefundRequest) GetOctoShopId() int32`

GetOctoShopId returns the OctoShopId field if non-nil, zero value otherwise.

### GetOctoShopIdOk

`func (o *RefundRequest) GetOctoShopIdOk() (*int32, bool)`

GetOctoShopIdOk returns a tuple with the OctoShopId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOctoShopId

`func (o *RefundRequest) SetOctoShopId(v int32)`

SetOctoShopId sets OctoShopId field to given value.


### GetShopRefundId

`func (o *RefundRequest) GetShopRefundId() string`

GetShopRefundId returns the ShopRefundId field if non-nil, zero value otherwise.

### GetShopRefundIdOk

`func (o *RefundRequest) GetShopRefundIdOk() (*string, bool)`

GetShopRefundIdOk returns a tuple with the ShopRefundId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopRefundId

`func (o *RefundRequest) SetShopRefundId(v string)`

SetShopRefundId sets ShopRefundId field to given value.


### GetOctoSecret

`func (o *RefundRequest) GetOctoSecret() string`

GetOctoSecret returns the OctoSecret field if non-nil, zero value otherwise.

### GetOctoSecretOk

`func (o *RefundRequest) GetOctoSecretOk() (*string, bool)`

GetOctoSecretOk returns a tuple with the OctoSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOctoSecret

`func (o *RefundRequest) SetOctoSecret(v string)`

SetOctoSecret sets OctoSecret field to given value.


### GetOctoPaymentUUID

`func (o *RefundRequest) GetOctoPaymentUUID() string`

GetOctoPaymentUUID returns the OctoPaymentUUID field if non-nil, zero value otherwise.

### GetOctoPaymentUUIDOk

`func (o *RefundRequest) GetOctoPaymentUUIDOk() (*string, bool)`

GetOctoPaymentUUIDOk returns a tuple with the OctoPaymentUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOctoPaymentUUID

`func (o *RefundRequest) SetOctoPaymentUUID(v string)`

SetOctoPaymentUUID sets OctoPaymentUUID field to given value.


### GetAmount

`func (o *RefundRequest) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *RefundRequest) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *RefundRequest) SetAmount(v float64)`

SetAmount sets Amount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


