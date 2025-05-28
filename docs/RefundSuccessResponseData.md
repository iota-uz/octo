# RefundSuccessResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OctoPaymentUUID** | Pointer to **string** |  | [optional] 
**RefundId** | Pointer to **string** |  | [optional] 
**RefundTime** | Pointer to **time.Time** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewRefundSuccessResponseData

`func NewRefundSuccessResponseData() *RefundSuccessResponseData`

NewRefundSuccessResponseData instantiates a new RefundSuccessResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRefundSuccessResponseDataWithDefaults

`func NewRefundSuccessResponseDataWithDefaults() *RefundSuccessResponseData`

NewRefundSuccessResponseDataWithDefaults instantiates a new RefundSuccessResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOctoPaymentUUID

`func (o *RefundSuccessResponseData) GetOctoPaymentUUID() string`

GetOctoPaymentUUID returns the OctoPaymentUUID field if non-nil, zero value otherwise.

### GetOctoPaymentUUIDOk

`func (o *RefundSuccessResponseData) GetOctoPaymentUUIDOk() (*string, bool)`

GetOctoPaymentUUIDOk returns a tuple with the OctoPaymentUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOctoPaymentUUID

`func (o *RefundSuccessResponseData) SetOctoPaymentUUID(v string)`

SetOctoPaymentUUID sets OctoPaymentUUID field to given value.

### HasOctoPaymentUUID

`func (o *RefundSuccessResponseData) HasOctoPaymentUUID() bool`

HasOctoPaymentUUID returns a boolean if a field has been set.

### GetRefundId

`func (o *RefundSuccessResponseData) GetRefundId() string`

GetRefundId returns the RefundId field if non-nil, zero value otherwise.

### GetRefundIdOk

`func (o *RefundSuccessResponseData) GetRefundIdOk() (*string, bool)`

GetRefundIdOk returns a tuple with the RefundId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundId

`func (o *RefundSuccessResponseData) SetRefundId(v string)`

SetRefundId sets RefundId field to given value.

### HasRefundId

`func (o *RefundSuccessResponseData) HasRefundId() bool`

HasRefundId returns a boolean if a field has been set.

### GetRefundTime

`func (o *RefundSuccessResponseData) GetRefundTime() time.Time`

GetRefundTime returns the RefundTime field if non-nil, zero value otherwise.

### GetRefundTimeOk

`func (o *RefundSuccessResponseData) GetRefundTimeOk() (*time.Time, bool)`

GetRefundTimeOk returns a tuple with the RefundTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundTime

`func (o *RefundSuccessResponseData) SetRefundTime(v time.Time)`

SetRefundTime sets RefundTime field to given value.

### HasRefundTime

`func (o *RefundSuccessResponseData) HasRefundTime() bool`

HasRefundTime returns a boolean if a field has been set.

### GetStatus

`func (o *RefundSuccessResponseData) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RefundSuccessResponseData) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RefundSuccessResponseData) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RefundSuccessResponseData) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


