# CallbackRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OctoSecret** | **string** |  | 
**OctoPaymentUUID** | **string** |  | 
**AcceptStatus** | **string** |  | 
**FinalAmount** | Pointer to **float32** |  | [optional] 

## Methods

### NewCallbackRequest

`func NewCallbackRequest(octoSecret string, octoPaymentUUID string, acceptStatus string, ) *CallbackRequest`

NewCallbackRequest instantiates a new CallbackRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallbackRequestWithDefaults

`func NewCallbackRequestWithDefaults() *CallbackRequest`

NewCallbackRequestWithDefaults instantiates a new CallbackRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOctoSecret

`func (o *CallbackRequest) GetOctoSecret() string`

GetOctoSecret returns the OctoSecret field if non-nil, zero value otherwise.

### GetOctoSecretOk

`func (o *CallbackRequest) GetOctoSecretOk() (*string, bool)`

GetOctoSecretOk returns a tuple with the OctoSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOctoSecret

`func (o *CallbackRequest) SetOctoSecret(v string)`

SetOctoSecret sets OctoSecret field to given value.


### GetOctoPaymentUUID

`func (o *CallbackRequest) GetOctoPaymentUUID() string`

GetOctoPaymentUUID returns the OctoPaymentUUID field if non-nil, zero value otherwise.

### GetOctoPaymentUUIDOk

`func (o *CallbackRequest) GetOctoPaymentUUIDOk() (*string, bool)`

GetOctoPaymentUUIDOk returns a tuple with the OctoPaymentUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOctoPaymentUUID

`func (o *CallbackRequest) SetOctoPaymentUUID(v string)`

SetOctoPaymentUUID sets OctoPaymentUUID field to given value.


### GetAcceptStatus

`func (o *CallbackRequest) GetAcceptStatus() string`

GetAcceptStatus returns the AcceptStatus field if non-nil, zero value otherwise.

### GetAcceptStatusOk

`func (o *CallbackRequest) GetAcceptStatusOk() (*string, bool)`

GetAcceptStatusOk returns a tuple with the AcceptStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptStatus

`func (o *CallbackRequest) SetAcceptStatus(v string)`

SetAcceptStatus sets AcceptStatus field to given value.


### GetFinalAmount

`func (o *CallbackRequest) GetFinalAmount() float32`

GetFinalAmount returns the FinalAmount field if non-nil, zero value otherwise.

### GetFinalAmountOk

`func (o *CallbackRequest) GetFinalAmountOk() (*float32, bool)`

GetFinalAmountOk returns a tuple with the FinalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalAmount

`func (o *CallbackRequest) SetFinalAmount(v float32)`

SetFinalAmount sets FinalAmount field to given value.

### HasFinalAmount

`func (o *CallbackRequest) HasFinalAmount() bool`

HasFinalAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


