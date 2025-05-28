# SetAcceptSuccessResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **int32** |  | [optional] 
**Data** | Pointer to [**SetAcceptSuccessResponseData**](SetAcceptSuccessResponseData.md) |  | [optional] 
**ApiMessageForDevelopers** | Pointer to **string** |  | [optional] 

## Methods

### NewSetAcceptSuccessResponse

`func NewSetAcceptSuccessResponse() *SetAcceptSuccessResponse`

NewSetAcceptSuccessResponse instantiates a new SetAcceptSuccessResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetAcceptSuccessResponseWithDefaults

`func NewSetAcceptSuccessResponseWithDefaults() *SetAcceptSuccessResponse`

NewSetAcceptSuccessResponseWithDefaults instantiates a new SetAcceptSuccessResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *SetAcceptSuccessResponse) GetError() int32`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *SetAcceptSuccessResponse) GetErrorOk() (*int32, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *SetAcceptSuccessResponse) SetError(v int32)`

SetError sets Error field to given value.

### HasError

`func (o *SetAcceptSuccessResponse) HasError() bool`

HasError returns a boolean if a field has been set.

### GetData

`func (o *SetAcceptSuccessResponse) GetData() SetAcceptSuccessResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SetAcceptSuccessResponse) GetDataOk() (*SetAcceptSuccessResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SetAcceptSuccessResponse) SetData(v SetAcceptSuccessResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *SetAcceptSuccessResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetApiMessageForDevelopers

`func (o *SetAcceptSuccessResponse) GetApiMessageForDevelopers() string`

GetApiMessageForDevelopers returns the ApiMessageForDevelopers field if non-nil, zero value otherwise.

### GetApiMessageForDevelopersOk

`func (o *SetAcceptSuccessResponse) GetApiMessageForDevelopersOk() (*string, bool)`

GetApiMessageForDevelopersOk returns a tuple with the ApiMessageForDevelopers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiMessageForDevelopers

`func (o *SetAcceptSuccessResponse) SetApiMessageForDevelopers(v string)`

SetApiMessageForDevelopers sets ApiMessageForDevelopers field to given value.

### HasApiMessageForDevelopers

`func (o *SetAcceptSuccessResponse) HasApiMessageForDevelopers() bool`

HasApiMessageForDevelopers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


