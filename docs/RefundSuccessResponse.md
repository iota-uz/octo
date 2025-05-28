# RefundSuccessResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **int32** |  | [optional] 
**ErrMessage** | Pointer to **string** |  | [optional] 
**Data** | Pointer to [**RefundSuccessResponseData**](RefundSuccessResponseData.md) |  | [optional] 
**ApiMessageForDevelopers** | Pointer to **string** |  | [optional] 

## Methods

### NewRefundSuccessResponse

`func NewRefundSuccessResponse() *RefundSuccessResponse`

NewRefundSuccessResponse instantiates a new RefundSuccessResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRefundSuccessResponseWithDefaults

`func NewRefundSuccessResponseWithDefaults() *RefundSuccessResponse`

NewRefundSuccessResponseWithDefaults instantiates a new RefundSuccessResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *RefundSuccessResponse) GetError() int32`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *RefundSuccessResponse) GetErrorOk() (*int32, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *RefundSuccessResponse) SetError(v int32)`

SetError sets Error field to given value.

### HasError

`func (o *RefundSuccessResponse) HasError() bool`

HasError returns a boolean if a field has been set.

### GetErrMessage

`func (o *RefundSuccessResponse) GetErrMessage() string`

GetErrMessage returns the ErrMessage field if non-nil, zero value otherwise.

### GetErrMessageOk

`func (o *RefundSuccessResponse) GetErrMessageOk() (*string, bool)`

GetErrMessageOk returns a tuple with the ErrMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrMessage

`func (o *RefundSuccessResponse) SetErrMessage(v string)`

SetErrMessage sets ErrMessage field to given value.

### HasErrMessage

`func (o *RefundSuccessResponse) HasErrMessage() bool`

HasErrMessage returns a boolean if a field has been set.

### GetData

`func (o *RefundSuccessResponse) GetData() RefundSuccessResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *RefundSuccessResponse) GetDataOk() (*RefundSuccessResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *RefundSuccessResponse) SetData(v RefundSuccessResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *RefundSuccessResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetApiMessageForDevelopers

`func (o *RefundSuccessResponse) GetApiMessageForDevelopers() string`

GetApiMessageForDevelopers returns the ApiMessageForDevelopers field if non-nil, zero value otherwise.

### GetApiMessageForDevelopersOk

`func (o *RefundSuccessResponse) GetApiMessageForDevelopersOk() (*string, bool)`

GetApiMessageForDevelopersOk returns a tuple with the ApiMessageForDevelopers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiMessageForDevelopers

`func (o *RefundSuccessResponse) SetApiMessageForDevelopers(v string)`

SetApiMessageForDevelopers sets ApiMessageForDevelopers field to given value.

### HasApiMessageForDevelopers

`func (o *RefundSuccessResponse) HasApiMessageForDevelopers() bool`

HasApiMessageForDevelopers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


