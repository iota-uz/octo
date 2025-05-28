# PreparePaymentSuccessResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | **int32** |  | 
**Data** | [**PreparePaymentSuccessResponseData**](PreparePaymentSuccessResponseData.md) |  | 
**ApiMessageForDevelopers** | Pointer to **string** |  | [optional] 
**ShopTransactionId** | Pointer to **string** |  | [optional] 
**OctoPaymentUUID** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**OctoPayUrl** | Pointer to **string** |  | [optional] 
**RefundedSum** | Pointer to **float32** |  | [optional] 
**TotalSum** | Pointer to **float32** |  | [optional] 

## Methods

### NewPreparePaymentSuccessResponse

`func NewPreparePaymentSuccessResponse(error_ int32, data PreparePaymentSuccessResponseData, ) *PreparePaymentSuccessResponse`

NewPreparePaymentSuccessResponse instantiates a new PreparePaymentSuccessResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPreparePaymentSuccessResponseWithDefaults

`func NewPreparePaymentSuccessResponseWithDefaults() *PreparePaymentSuccessResponse`

NewPreparePaymentSuccessResponseWithDefaults instantiates a new PreparePaymentSuccessResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *PreparePaymentSuccessResponse) GetError() int32`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *PreparePaymentSuccessResponse) GetErrorOk() (*int32, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *PreparePaymentSuccessResponse) SetError(v int32)`

SetError sets Error field to given value.


### GetData

`func (o *PreparePaymentSuccessResponse) GetData() PreparePaymentSuccessResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PreparePaymentSuccessResponse) GetDataOk() (*PreparePaymentSuccessResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PreparePaymentSuccessResponse) SetData(v PreparePaymentSuccessResponseData)`

SetData sets Data field to given value.


### GetApiMessageForDevelopers

`func (o *PreparePaymentSuccessResponse) GetApiMessageForDevelopers() string`

GetApiMessageForDevelopers returns the ApiMessageForDevelopers field if non-nil, zero value otherwise.

### GetApiMessageForDevelopersOk

`func (o *PreparePaymentSuccessResponse) GetApiMessageForDevelopersOk() (*string, bool)`

GetApiMessageForDevelopersOk returns a tuple with the ApiMessageForDevelopers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiMessageForDevelopers

`func (o *PreparePaymentSuccessResponse) SetApiMessageForDevelopers(v string)`

SetApiMessageForDevelopers sets ApiMessageForDevelopers field to given value.

### HasApiMessageForDevelopers

`func (o *PreparePaymentSuccessResponse) HasApiMessageForDevelopers() bool`

HasApiMessageForDevelopers returns a boolean if a field has been set.

### GetShopTransactionId

`func (o *PreparePaymentSuccessResponse) GetShopTransactionId() string`

GetShopTransactionId returns the ShopTransactionId field if non-nil, zero value otherwise.

### GetShopTransactionIdOk

`func (o *PreparePaymentSuccessResponse) GetShopTransactionIdOk() (*string, bool)`

GetShopTransactionIdOk returns a tuple with the ShopTransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopTransactionId

`func (o *PreparePaymentSuccessResponse) SetShopTransactionId(v string)`

SetShopTransactionId sets ShopTransactionId field to given value.

### HasShopTransactionId

`func (o *PreparePaymentSuccessResponse) HasShopTransactionId() bool`

HasShopTransactionId returns a boolean if a field has been set.

### GetOctoPaymentUUID

`func (o *PreparePaymentSuccessResponse) GetOctoPaymentUUID() string`

GetOctoPaymentUUID returns the OctoPaymentUUID field if non-nil, zero value otherwise.

### GetOctoPaymentUUIDOk

`func (o *PreparePaymentSuccessResponse) GetOctoPaymentUUIDOk() (*string, bool)`

GetOctoPaymentUUIDOk returns a tuple with the OctoPaymentUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOctoPaymentUUID

`func (o *PreparePaymentSuccessResponse) SetOctoPaymentUUID(v string)`

SetOctoPaymentUUID sets OctoPaymentUUID field to given value.

### HasOctoPaymentUUID

`func (o *PreparePaymentSuccessResponse) HasOctoPaymentUUID() bool`

HasOctoPaymentUUID returns a boolean if a field has been set.

### GetStatus

`func (o *PreparePaymentSuccessResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PreparePaymentSuccessResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PreparePaymentSuccessResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PreparePaymentSuccessResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetOctoPayUrl

`func (o *PreparePaymentSuccessResponse) GetOctoPayUrl() string`

GetOctoPayUrl returns the OctoPayUrl field if non-nil, zero value otherwise.

### GetOctoPayUrlOk

`func (o *PreparePaymentSuccessResponse) GetOctoPayUrlOk() (*string, bool)`

GetOctoPayUrlOk returns a tuple with the OctoPayUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOctoPayUrl

`func (o *PreparePaymentSuccessResponse) SetOctoPayUrl(v string)`

SetOctoPayUrl sets OctoPayUrl field to given value.

### HasOctoPayUrl

`func (o *PreparePaymentSuccessResponse) HasOctoPayUrl() bool`

HasOctoPayUrl returns a boolean if a field has been set.

### GetRefundedSum

`func (o *PreparePaymentSuccessResponse) GetRefundedSum() float32`

GetRefundedSum returns the RefundedSum field if non-nil, zero value otherwise.

### GetRefundedSumOk

`func (o *PreparePaymentSuccessResponse) GetRefundedSumOk() (*float32, bool)`

GetRefundedSumOk returns a tuple with the RefundedSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundedSum

`func (o *PreparePaymentSuccessResponse) SetRefundedSum(v float32)`

SetRefundedSum sets RefundedSum field to given value.

### HasRefundedSum

`func (o *PreparePaymentSuccessResponse) HasRefundedSum() bool`

HasRefundedSum returns a boolean if a field has been set.

### GetTotalSum

`func (o *PreparePaymentSuccessResponse) GetTotalSum() float32`

GetTotalSum returns the TotalSum field if non-nil, zero value otherwise.

### GetTotalSumOk

`func (o *PreparePaymentSuccessResponse) GetTotalSumOk() (*float32, bool)`

GetTotalSumOk returns a tuple with the TotalSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSum

`func (o *PreparePaymentSuccessResponse) SetTotalSum(v float32)`

SetTotalSum sets TotalSum field to given value.

### HasTotalSum

`func (o *PreparePaymentSuccessResponse) HasTotalSum() bool`

HasTotalSum returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


