# RefundResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **int32** |  | [optional] 
**ErrMessage** | Pointer to **string** |  | [optional] 
**ErrorMessage** | Pointer to **string** |  | [optional] 
**ApiMessageForDevelopers** | Pointer to **string** |  | [optional] 
**Data** | Pointer to [**RefundSuccessResponseData**](RefundSuccessResponseData.md) |  | [optional] 

## Methods

### NewRefundResponse

`func NewRefundResponse() *RefundResponse`

NewRefundResponse instantiates a new RefundResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRefundResponseWithDefaults

`func NewRefundResponseWithDefaults() *RefundResponse`

NewRefundResponseWithDefaults instantiates a new RefundResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *RefundResponse) GetError() int32`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *RefundResponse) GetErrorOk() (*int32, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *RefundResponse) SetError(v int32)`

SetError sets Error field to given value.

### HasError

`func (o *RefundResponse) HasError() bool`

HasError returns a boolean if a field has been set.

### GetErrMessage

`func (o *RefundResponse) GetErrMessage() string`

GetErrMessage returns the ErrMessage field if non-nil, zero value otherwise.

### GetErrMessageOk

`func (o *RefundResponse) GetErrMessageOk() (*string, bool)`

GetErrMessageOk returns a tuple with the ErrMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrMessage

`func (o *RefundResponse) SetErrMessage(v string)`

SetErrMessage sets ErrMessage field to given value.

### HasErrMessage

`func (o *RefundResponse) HasErrMessage() bool`

HasErrMessage returns a boolean if a field has been set.

### GetErrorMessage

`func (o *RefundResponse) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *RefundResponse) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *RefundResponse) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *RefundResponse) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetApiMessageForDevelopers

`func (o *RefundResponse) GetApiMessageForDevelopers() string`

GetApiMessageForDevelopers returns the ApiMessageForDevelopers field if non-nil, zero value otherwise.

### GetApiMessageForDevelopersOk

`func (o *RefundResponse) GetApiMessageForDevelopersOk() (*string, bool)`

GetApiMessageForDevelopersOk returns a tuple with the ApiMessageForDevelopers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiMessageForDevelopers

`func (o *RefundResponse) SetApiMessageForDevelopers(v string)`

SetApiMessageForDevelopers sets ApiMessageForDevelopers field to given value.

### HasApiMessageForDevelopers

`func (o *RefundResponse) HasApiMessageForDevelopers() bool`

HasApiMessageForDevelopers returns a boolean if a field has been set.

### GetData

`func (o *RefundResponse) GetData() RefundSuccessResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *RefundResponse) GetDataOk() (*RefundSuccessResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *RefundResponse) SetData(v RefundSuccessResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *RefundResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


