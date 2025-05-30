# CallbackResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcceptStatus** | Pointer to **string** |  | [optional] 
**FinalAmount** | Pointer to **float64** |  | [optional] 

## Methods

### NewCallbackResponse

`func NewCallbackResponse() *CallbackResponse`

NewCallbackResponse instantiates a new CallbackResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallbackResponseWithDefaults

`func NewCallbackResponseWithDefaults() *CallbackResponse`

NewCallbackResponseWithDefaults instantiates a new CallbackResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcceptStatus

`func (o *CallbackResponse) GetAcceptStatus() string`

GetAcceptStatus returns the AcceptStatus field if non-nil, zero value otherwise.

### GetAcceptStatusOk

`func (o *CallbackResponse) GetAcceptStatusOk() (*string, bool)`

GetAcceptStatusOk returns a tuple with the AcceptStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptStatus

`func (o *CallbackResponse) SetAcceptStatus(v string)`

SetAcceptStatus sets AcceptStatus field to given value.

### HasAcceptStatus

`func (o *CallbackResponse) HasAcceptStatus() bool`

HasAcceptStatus returns a boolean if a field has been set.

### GetFinalAmount

`func (o *CallbackResponse) GetFinalAmount() float64`

GetFinalAmount returns the FinalAmount field if non-nil, zero value otherwise.

### GetFinalAmountOk

`func (o *CallbackResponse) GetFinalAmountOk() (*float64, bool)`

GetFinalAmountOk returns a tuple with the FinalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalAmount

`func (o *CallbackResponse) SetFinalAmount(v float64)`

SetFinalAmount sets FinalAmount field to given value.

### HasFinalAmount

`func (o *CallbackResponse) HasFinalAmount() bool`

HasFinalAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


