# SetAcceptResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **int32** |  | [optional] 
**ErrMessage** | Pointer to **string** |  | [optional] 
**ErrorMessage** | Pointer to **string** |  | [optional] 
**ApiMessageForDevelopers** | Pointer to **string** |  | [optional] 
**Data** | Pointer to [**SetAcceptSuccessResponseData**](SetAcceptSuccessResponseData.md) |  | [optional] 

## Methods

### NewSetAcceptResponse

`func NewSetAcceptResponse() *SetAcceptResponse`

NewSetAcceptResponse instantiates a new SetAcceptResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetAcceptResponseWithDefaults

`func NewSetAcceptResponseWithDefaults() *SetAcceptResponse`

NewSetAcceptResponseWithDefaults instantiates a new SetAcceptResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *SetAcceptResponse) GetError() int32`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *SetAcceptResponse) GetErrorOk() (*int32, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *SetAcceptResponse) SetError(v int32)`

SetError sets Error field to given value.

### HasError

`func (o *SetAcceptResponse) HasError() bool`

HasError returns a boolean if a field has been set.

### GetErrMessage

`func (o *SetAcceptResponse) GetErrMessage() string`

GetErrMessage returns the ErrMessage field if non-nil, zero value otherwise.

### GetErrMessageOk

`func (o *SetAcceptResponse) GetErrMessageOk() (*string, bool)`

GetErrMessageOk returns a tuple with the ErrMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrMessage

`func (o *SetAcceptResponse) SetErrMessage(v string)`

SetErrMessage sets ErrMessage field to given value.

### HasErrMessage

`func (o *SetAcceptResponse) HasErrMessage() bool`

HasErrMessage returns a boolean if a field has been set.

### GetErrorMessage

`func (o *SetAcceptResponse) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *SetAcceptResponse) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *SetAcceptResponse) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *SetAcceptResponse) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetApiMessageForDevelopers

`func (o *SetAcceptResponse) GetApiMessageForDevelopers() string`

GetApiMessageForDevelopers returns the ApiMessageForDevelopers field if non-nil, zero value otherwise.

### GetApiMessageForDevelopersOk

`func (o *SetAcceptResponse) GetApiMessageForDevelopersOk() (*string, bool)`

GetApiMessageForDevelopersOk returns a tuple with the ApiMessageForDevelopers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiMessageForDevelopers

`func (o *SetAcceptResponse) SetApiMessageForDevelopers(v string)`

SetApiMessageForDevelopers sets ApiMessageForDevelopers field to given value.

### HasApiMessageForDevelopers

`func (o *SetAcceptResponse) HasApiMessageForDevelopers() bool`

HasApiMessageForDevelopers returns a boolean if a field has been set.

### GetData

`func (o *SetAcceptResponse) GetData() SetAcceptSuccessResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SetAcceptResponse) GetDataOk() (*SetAcceptSuccessResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SetAcceptResponse) SetData(v SetAcceptSuccessResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *SetAcceptResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


