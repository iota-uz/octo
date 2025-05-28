# PreparePaymentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **int32** |  | [optional] 
**ErrMessage** | Pointer to **string** |  | [optional] 
**ErrorMessage** | Pointer to **string** |  | [optional] 
**ApiMessageForDevelopers** | Pointer to **string** |  | [optional] 
**ShopTransactionId** | Pointer to **string** |  | [optional] 
**OctoPaymentUUID** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**OctoPayUrl** | Pointer to **string** |  | [optional] 
**RefundedSum** | Pointer to **float32** |  | [optional] 
**TotalSum** | Pointer to **float32** |  | [optional] 
**Data** | Pointer to [**PreparePaymentSuccessResponseData**](PreparePaymentSuccessResponseData.md) |  | [optional] 

## Methods

### NewPreparePaymentResponse

`func NewPreparePaymentResponse() *PreparePaymentResponse`

NewPreparePaymentResponse instantiates a new PreparePaymentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPreparePaymentResponseWithDefaults

`func NewPreparePaymentResponseWithDefaults() *PreparePaymentResponse`

NewPreparePaymentResponseWithDefaults instantiates a new PreparePaymentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *PreparePaymentResponse) GetError() int32`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *PreparePaymentResponse) GetErrorOk() (*int32, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *PreparePaymentResponse) SetError(v int32)`

SetError sets Error field to given value.

### HasError

`func (o *PreparePaymentResponse) HasError() bool`

HasError returns a boolean if a field has been set.

### GetErrMessage

`func (o *PreparePaymentResponse) GetErrMessage() string`

GetErrMessage returns the ErrMessage field if non-nil, zero value otherwise.

### GetErrMessageOk

`func (o *PreparePaymentResponse) GetErrMessageOk() (*string, bool)`

GetErrMessageOk returns a tuple with the ErrMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrMessage

`func (o *PreparePaymentResponse) SetErrMessage(v string)`

SetErrMessage sets ErrMessage field to given value.

### HasErrMessage

`func (o *PreparePaymentResponse) HasErrMessage() bool`

HasErrMessage returns a boolean if a field has been set.

### GetErrorMessage

`func (o *PreparePaymentResponse) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *PreparePaymentResponse) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *PreparePaymentResponse) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *PreparePaymentResponse) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetApiMessageForDevelopers

`func (o *PreparePaymentResponse) GetApiMessageForDevelopers() string`

GetApiMessageForDevelopers returns the ApiMessageForDevelopers field if non-nil, zero value otherwise.

### GetApiMessageForDevelopersOk

`func (o *PreparePaymentResponse) GetApiMessageForDevelopersOk() (*string, bool)`

GetApiMessageForDevelopersOk returns a tuple with the ApiMessageForDevelopers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiMessageForDevelopers

`func (o *PreparePaymentResponse) SetApiMessageForDevelopers(v string)`

SetApiMessageForDevelopers sets ApiMessageForDevelopers field to given value.

### HasApiMessageForDevelopers

`func (o *PreparePaymentResponse) HasApiMessageForDevelopers() bool`

HasApiMessageForDevelopers returns a boolean if a field has been set.

### GetShopTransactionId

`func (o *PreparePaymentResponse) GetShopTransactionId() string`

GetShopTransactionId returns the ShopTransactionId field if non-nil, zero value otherwise.

### GetShopTransactionIdOk

`func (o *PreparePaymentResponse) GetShopTransactionIdOk() (*string, bool)`

GetShopTransactionIdOk returns a tuple with the ShopTransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopTransactionId

`func (o *PreparePaymentResponse) SetShopTransactionId(v string)`

SetShopTransactionId sets ShopTransactionId field to given value.

### HasShopTransactionId

`func (o *PreparePaymentResponse) HasShopTransactionId() bool`

HasShopTransactionId returns a boolean if a field has been set.

### GetOctoPaymentUUID

`func (o *PreparePaymentResponse) GetOctoPaymentUUID() string`

GetOctoPaymentUUID returns the OctoPaymentUUID field if non-nil, zero value otherwise.

### GetOctoPaymentUUIDOk

`func (o *PreparePaymentResponse) GetOctoPaymentUUIDOk() (*string, bool)`

GetOctoPaymentUUIDOk returns a tuple with the OctoPaymentUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOctoPaymentUUID

`func (o *PreparePaymentResponse) SetOctoPaymentUUID(v string)`

SetOctoPaymentUUID sets OctoPaymentUUID field to given value.

### HasOctoPaymentUUID

`func (o *PreparePaymentResponse) HasOctoPaymentUUID() bool`

HasOctoPaymentUUID returns a boolean if a field has been set.

### GetStatus

`func (o *PreparePaymentResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PreparePaymentResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PreparePaymentResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PreparePaymentResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetOctoPayUrl

`func (o *PreparePaymentResponse) GetOctoPayUrl() string`

GetOctoPayUrl returns the OctoPayUrl field if non-nil, zero value otherwise.

### GetOctoPayUrlOk

`func (o *PreparePaymentResponse) GetOctoPayUrlOk() (*string, bool)`

GetOctoPayUrlOk returns a tuple with the OctoPayUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOctoPayUrl

`func (o *PreparePaymentResponse) SetOctoPayUrl(v string)`

SetOctoPayUrl sets OctoPayUrl field to given value.

### HasOctoPayUrl

`func (o *PreparePaymentResponse) HasOctoPayUrl() bool`

HasOctoPayUrl returns a boolean if a field has been set.

### GetRefundedSum

`func (o *PreparePaymentResponse) GetRefundedSum() float32`

GetRefundedSum returns the RefundedSum field if non-nil, zero value otherwise.

### GetRefundedSumOk

`func (o *PreparePaymentResponse) GetRefundedSumOk() (*float32, bool)`

GetRefundedSumOk returns a tuple with the RefundedSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundedSum

`func (o *PreparePaymentResponse) SetRefundedSum(v float32)`

SetRefundedSum sets RefundedSum field to given value.

### HasRefundedSum

`func (o *PreparePaymentResponse) HasRefundedSum() bool`

HasRefundedSum returns a boolean if a field has been set.

### GetTotalSum

`func (o *PreparePaymentResponse) GetTotalSum() float32`

GetTotalSum returns the TotalSum field if non-nil, zero value otherwise.

### GetTotalSumOk

`func (o *PreparePaymentResponse) GetTotalSumOk() (*float32, bool)`

GetTotalSumOk returns a tuple with the TotalSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSum

`func (o *PreparePaymentResponse) SetTotalSum(v float32)`

SetTotalSum sets TotalSum field to given value.

### HasTotalSum

`func (o *PreparePaymentResponse) HasTotalSum() bool`

HasTotalSum returns a boolean if a field has been set.

### GetData

`func (o *PreparePaymentResponse) GetData() PreparePaymentSuccessResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PreparePaymentResponse) GetDataOk() (*PreparePaymentSuccessResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PreparePaymentResponse) SetData(v PreparePaymentSuccessResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *PreparePaymentResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


