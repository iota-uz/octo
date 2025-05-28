# SetAcceptRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OctoShopId** | **int32** |  | 
**OctoSecret** | **string** |  | 
**OctoPaymentUUID** | **string** |  | 
**AcceptStatus** | **string** |  | 
**FinalAmount** | **float32** |  | 

## Methods

### NewSetAcceptRequest

`func NewSetAcceptRequest(octoShopId int32, octoSecret string, octoPaymentUUID string, acceptStatus string, finalAmount float32, ) *SetAcceptRequest`

NewSetAcceptRequest instantiates a new SetAcceptRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetAcceptRequestWithDefaults

`func NewSetAcceptRequestWithDefaults() *SetAcceptRequest`

NewSetAcceptRequestWithDefaults instantiates a new SetAcceptRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOctoShopId

`func (o *SetAcceptRequest) GetOctoShopId() int32`

GetOctoShopId returns the OctoShopId field if non-nil, zero value otherwise.

### GetOctoShopIdOk

`func (o *SetAcceptRequest) GetOctoShopIdOk() (*int32, bool)`

GetOctoShopIdOk returns a tuple with the OctoShopId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOctoShopId

`func (o *SetAcceptRequest) SetOctoShopId(v int32)`

SetOctoShopId sets OctoShopId field to given value.


### GetOctoSecret

`func (o *SetAcceptRequest) GetOctoSecret() string`

GetOctoSecret returns the OctoSecret field if non-nil, zero value otherwise.

### GetOctoSecretOk

`func (o *SetAcceptRequest) GetOctoSecretOk() (*string, bool)`

GetOctoSecretOk returns a tuple with the OctoSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOctoSecret

`func (o *SetAcceptRequest) SetOctoSecret(v string)`

SetOctoSecret sets OctoSecret field to given value.


### GetOctoPaymentUUID

`func (o *SetAcceptRequest) GetOctoPaymentUUID() string`

GetOctoPaymentUUID returns the OctoPaymentUUID field if non-nil, zero value otherwise.

### GetOctoPaymentUUIDOk

`func (o *SetAcceptRequest) GetOctoPaymentUUIDOk() (*string, bool)`

GetOctoPaymentUUIDOk returns a tuple with the OctoPaymentUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOctoPaymentUUID

`func (o *SetAcceptRequest) SetOctoPaymentUUID(v string)`

SetOctoPaymentUUID sets OctoPaymentUUID field to given value.


### GetAcceptStatus

`func (o *SetAcceptRequest) GetAcceptStatus() string`

GetAcceptStatus returns the AcceptStatus field if non-nil, zero value otherwise.

### GetAcceptStatusOk

`func (o *SetAcceptRequest) GetAcceptStatusOk() (*string, bool)`

GetAcceptStatusOk returns a tuple with the AcceptStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptStatus

`func (o *SetAcceptRequest) SetAcceptStatus(v string)`

SetAcceptStatus sets AcceptStatus field to given value.


### GetFinalAmount

`func (o *SetAcceptRequest) GetFinalAmount() float32`

GetFinalAmount returns the FinalAmount field if non-nil, zero value otherwise.

### GetFinalAmountOk

`func (o *SetAcceptRequest) GetFinalAmountOk() (*float32, bool)`

GetFinalAmountOk returns a tuple with the FinalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalAmount

`func (o *SetAcceptRequest) SetFinalAmount(v float32)`

SetFinalAmount sets FinalAmount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


