# SetAcceptSuccessResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShopTransactionId** | Pointer to **string** |  | [optional] 
**OctoPaymentUUID** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**OctoPayUrl** | Pointer to **string** |  | [optional] 
**TransferSum** | Pointer to **float64** |  | [optional] 
**RefundedSum** | Pointer to **float64** |  | [optional] 
**TotalSum** | Pointer to **float64** |  | [optional] 
**PayedTime** | Pointer to **string** |  | [optional] 

## Methods

### NewSetAcceptSuccessResponseData

`func NewSetAcceptSuccessResponseData() *SetAcceptSuccessResponseData`

NewSetAcceptSuccessResponseData instantiates a new SetAcceptSuccessResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetAcceptSuccessResponseDataWithDefaults

`func NewSetAcceptSuccessResponseDataWithDefaults() *SetAcceptSuccessResponseData`

NewSetAcceptSuccessResponseDataWithDefaults instantiates a new SetAcceptSuccessResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShopTransactionId

`func (o *SetAcceptSuccessResponseData) GetShopTransactionId() string`

GetShopTransactionId returns the ShopTransactionId field if non-nil, zero value otherwise.

### GetShopTransactionIdOk

`func (o *SetAcceptSuccessResponseData) GetShopTransactionIdOk() (*string, bool)`

GetShopTransactionIdOk returns a tuple with the ShopTransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopTransactionId

`func (o *SetAcceptSuccessResponseData) SetShopTransactionId(v string)`

SetShopTransactionId sets ShopTransactionId field to given value.

### HasShopTransactionId

`func (o *SetAcceptSuccessResponseData) HasShopTransactionId() bool`

HasShopTransactionId returns a boolean if a field has been set.

### GetOctoPaymentUUID

`func (o *SetAcceptSuccessResponseData) GetOctoPaymentUUID() string`

GetOctoPaymentUUID returns the OctoPaymentUUID field if non-nil, zero value otherwise.

### GetOctoPaymentUUIDOk

`func (o *SetAcceptSuccessResponseData) GetOctoPaymentUUIDOk() (*string, bool)`

GetOctoPaymentUUIDOk returns a tuple with the OctoPaymentUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOctoPaymentUUID

`func (o *SetAcceptSuccessResponseData) SetOctoPaymentUUID(v string)`

SetOctoPaymentUUID sets OctoPaymentUUID field to given value.

### HasOctoPaymentUUID

`func (o *SetAcceptSuccessResponseData) HasOctoPaymentUUID() bool`

HasOctoPaymentUUID returns a boolean if a field has been set.

### GetStatus

`func (o *SetAcceptSuccessResponseData) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SetAcceptSuccessResponseData) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SetAcceptSuccessResponseData) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SetAcceptSuccessResponseData) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetOctoPayUrl

`func (o *SetAcceptSuccessResponseData) GetOctoPayUrl() string`

GetOctoPayUrl returns the OctoPayUrl field if non-nil, zero value otherwise.

### GetOctoPayUrlOk

`func (o *SetAcceptSuccessResponseData) GetOctoPayUrlOk() (*string, bool)`

GetOctoPayUrlOk returns a tuple with the OctoPayUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOctoPayUrl

`func (o *SetAcceptSuccessResponseData) SetOctoPayUrl(v string)`

SetOctoPayUrl sets OctoPayUrl field to given value.

### HasOctoPayUrl

`func (o *SetAcceptSuccessResponseData) HasOctoPayUrl() bool`

HasOctoPayUrl returns a boolean if a field has been set.

### GetTransferSum

`func (o *SetAcceptSuccessResponseData) GetTransferSum() float64`

GetTransferSum returns the TransferSum field if non-nil, zero value otherwise.

### GetTransferSumOk

`func (o *SetAcceptSuccessResponseData) GetTransferSumOk() (*float64, bool)`

GetTransferSumOk returns a tuple with the TransferSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferSum

`func (o *SetAcceptSuccessResponseData) SetTransferSum(v float64)`

SetTransferSum sets TransferSum field to given value.

### HasTransferSum

`func (o *SetAcceptSuccessResponseData) HasTransferSum() bool`

HasTransferSum returns a boolean if a field has been set.

### GetRefundedSum

`func (o *SetAcceptSuccessResponseData) GetRefundedSum() float64`

GetRefundedSum returns the RefundedSum field if non-nil, zero value otherwise.

### GetRefundedSumOk

`func (o *SetAcceptSuccessResponseData) GetRefundedSumOk() (*float64, bool)`

GetRefundedSumOk returns a tuple with the RefundedSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundedSum

`func (o *SetAcceptSuccessResponseData) SetRefundedSum(v float64)`

SetRefundedSum sets RefundedSum field to given value.

### HasRefundedSum

`func (o *SetAcceptSuccessResponseData) HasRefundedSum() bool`

HasRefundedSum returns a boolean if a field has been set.

### GetTotalSum

`func (o *SetAcceptSuccessResponseData) GetTotalSum() float64`

GetTotalSum returns the TotalSum field if non-nil, zero value otherwise.

### GetTotalSumOk

`func (o *SetAcceptSuccessResponseData) GetTotalSumOk() (*float64, bool)`

GetTotalSumOk returns a tuple with the TotalSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSum

`func (o *SetAcceptSuccessResponseData) SetTotalSum(v float64)`

SetTotalSum sets TotalSum field to given value.

### HasTotalSum

`func (o *SetAcceptSuccessResponseData) HasTotalSum() bool`

HasTotalSum returns a boolean if a field has been set.

### GetPayedTime

`func (o *SetAcceptSuccessResponseData) GetPayedTime() string`

GetPayedTime returns the PayedTime field if non-nil, zero value otherwise.

### GetPayedTimeOk

`func (o *SetAcceptSuccessResponseData) GetPayedTimeOk() (*string, bool)`

GetPayedTimeOk returns a tuple with the PayedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayedTime

`func (o *SetAcceptSuccessResponseData) SetPayedTime(v string)`

SetPayedTime sets PayedTime field to given value.

### HasPayedTime

`func (o *SetAcceptSuccessResponseData) HasPayedTime() bool`

HasPayedTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


