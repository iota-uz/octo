# NotificationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShopTransactionId** | **string** | Transaction ID on merchant&#39;s side | 
**OctoPaymentUUID** | **string** | Octo payment UUID | 
**Status** | **string** | Current transaction status | 
**Signature** | **string** | Cryptographic signature | 
**HashKey** | **string** | Hash key for validation | 
**TotalSum** | Pointer to **float64** | Total amount of the payment | [optional] 
**TransferSum** | Pointer to **float64** | Amount transferred after commissions | [optional] 
**RefundedSum** | Pointer to **float64** | Refunded amount if any | [optional] 
**CardCountry** | Pointer to **string** | Country where the card was issued | [optional] 
**MaskedPan** | Pointer to **string** | Masked card number | [optional] 
**Rrn** | Pointer to **string** | Retrieval Reference Number | [optional] 
**RiskLevel** | Pointer to **int32** | Risk level of the transaction | [optional] 
**PayedTime** | Pointer to **time.Time** | Timestamp of successful payment | [optional] 
**CardType** | Pointer to **string** | Type of the card used | [optional] 
**IsPhysicalCard** | Pointer to **bool** | Whether the card is physical | [optional] 

## Methods

### NewNotificationRequest

`func NewNotificationRequest(shopTransactionId string, octoPaymentUUID string, status string, signature string, hashKey string, ) *NotificationRequest`

NewNotificationRequest instantiates a new NotificationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationRequestWithDefaults

`func NewNotificationRequestWithDefaults() *NotificationRequest`

NewNotificationRequestWithDefaults instantiates a new NotificationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShopTransactionId

`func (o *NotificationRequest) GetShopTransactionId() string`

GetShopTransactionId returns the ShopTransactionId field if non-nil, zero value otherwise.

### GetShopTransactionIdOk

`func (o *NotificationRequest) GetShopTransactionIdOk() (*string, bool)`

GetShopTransactionIdOk returns a tuple with the ShopTransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopTransactionId

`func (o *NotificationRequest) SetShopTransactionId(v string)`

SetShopTransactionId sets ShopTransactionId field to given value.


### GetOctoPaymentUUID

`func (o *NotificationRequest) GetOctoPaymentUUID() string`

GetOctoPaymentUUID returns the OctoPaymentUUID field if non-nil, zero value otherwise.

### GetOctoPaymentUUIDOk

`func (o *NotificationRequest) GetOctoPaymentUUIDOk() (*string, bool)`

GetOctoPaymentUUIDOk returns a tuple with the OctoPaymentUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOctoPaymentUUID

`func (o *NotificationRequest) SetOctoPaymentUUID(v string)`

SetOctoPaymentUUID sets OctoPaymentUUID field to given value.


### GetStatus

`func (o *NotificationRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NotificationRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NotificationRequest) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetSignature

`func (o *NotificationRequest) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *NotificationRequest) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *NotificationRequest) SetSignature(v string)`

SetSignature sets Signature field to given value.


### GetHashKey

`func (o *NotificationRequest) GetHashKey() string`

GetHashKey returns the HashKey field if non-nil, zero value otherwise.

### GetHashKeyOk

`func (o *NotificationRequest) GetHashKeyOk() (*string, bool)`

GetHashKeyOk returns a tuple with the HashKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashKey

`func (o *NotificationRequest) SetHashKey(v string)`

SetHashKey sets HashKey field to given value.


### GetTotalSum

`func (o *NotificationRequest) GetTotalSum() float64`

GetTotalSum returns the TotalSum field if non-nil, zero value otherwise.

### GetTotalSumOk

`func (o *NotificationRequest) GetTotalSumOk() (*float64, bool)`

GetTotalSumOk returns a tuple with the TotalSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSum

`func (o *NotificationRequest) SetTotalSum(v float64)`

SetTotalSum sets TotalSum field to given value.

### HasTotalSum

`func (o *NotificationRequest) HasTotalSum() bool`

HasTotalSum returns a boolean if a field has been set.

### GetTransferSum

`func (o *NotificationRequest) GetTransferSum() float64`

GetTransferSum returns the TransferSum field if non-nil, zero value otherwise.

### GetTransferSumOk

`func (o *NotificationRequest) GetTransferSumOk() (*float64, bool)`

GetTransferSumOk returns a tuple with the TransferSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferSum

`func (o *NotificationRequest) SetTransferSum(v float64)`

SetTransferSum sets TransferSum field to given value.

### HasTransferSum

`func (o *NotificationRequest) HasTransferSum() bool`

HasTransferSum returns a boolean if a field has been set.

### GetRefundedSum

`func (o *NotificationRequest) GetRefundedSum() float64`

GetRefundedSum returns the RefundedSum field if non-nil, zero value otherwise.

### GetRefundedSumOk

`func (o *NotificationRequest) GetRefundedSumOk() (*float64, bool)`

GetRefundedSumOk returns a tuple with the RefundedSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundedSum

`func (o *NotificationRequest) SetRefundedSum(v float64)`

SetRefundedSum sets RefundedSum field to given value.

### HasRefundedSum

`func (o *NotificationRequest) HasRefundedSum() bool`

HasRefundedSum returns a boolean if a field has been set.

### GetCardCountry

`func (o *NotificationRequest) GetCardCountry() string`

GetCardCountry returns the CardCountry field if non-nil, zero value otherwise.

### GetCardCountryOk

`func (o *NotificationRequest) GetCardCountryOk() (*string, bool)`

GetCardCountryOk returns a tuple with the CardCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardCountry

`func (o *NotificationRequest) SetCardCountry(v string)`

SetCardCountry sets CardCountry field to given value.

### HasCardCountry

`func (o *NotificationRequest) HasCardCountry() bool`

HasCardCountry returns a boolean if a field has been set.

### GetMaskedPan

`func (o *NotificationRequest) GetMaskedPan() string`

GetMaskedPan returns the MaskedPan field if non-nil, zero value otherwise.

### GetMaskedPanOk

`func (o *NotificationRequest) GetMaskedPanOk() (*string, bool)`

GetMaskedPanOk returns a tuple with the MaskedPan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaskedPan

`func (o *NotificationRequest) SetMaskedPan(v string)`

SetMaskedPan sets MaskedPan field to given value.

### HasMaskedPan

`func (o *NotificationRequest) HasMaskedPan() bool`

HasMaskedPan returns a boolean if a field has been set.

### GetRrn

`func (o *NotificationRequest) GetRrn() string`

GetRrn returns the Rrn field if non-nil, zero value otherwise.

### GetRrnOk

`func (o *NotificationRequest) GetRrnOk() (*string, bool)`

GetRrnOk returns a tuple with the Rrn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRrn

`func (o *NotificationRequest) SetRrn(v string)`

SetRrn sets Rrn field to given value.

### HasRrn

`func (o *NotificationRequest) HasRrn() bool`

HasRrn returns a boolean if a field has been set.

### GetRiskLevel

`func (o *NotificationRequest) GetRiskLevel() int32`

GetRiskLevel returns the RiskLevel field if non-nil, zero value otherwise.

### GetRiskLevelOk

`func (o *NotificationRequest) GetRiskLevelOk() (*int32, bool)`

GetRiskLevelOk returns a tuple with the RiskLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskLevel

`func (o *NotificationRequest) SetRiskLevel(v int32)`

SetRiskLevel sets RiskLevel field to given value.

### HasRiskLevel

`func (o *NotificationRequest) HasRiskLevel() bool`

HasRiskLevel returns a boolean if a field has been set.

### GetPayedTime

`func (o *NotificationRequest) GetPayedTime() time.Time`

GetPayedTime returns the PayedTime field if non-nil, zero value otherwise.

### GetPayedTimeOk

`func (o *NotificationRequest) GetPayedTimeOk() (*time.Time, bool)`

GetPayedTimeOk returns a tuple with the PayedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayedTime

`func (o *NotificationRequest) SetPayedTime(v time.Time)`

SetPayedTime sets PayedTime field to given value.

### HasPayedTime

`func (o *NotificationRequest) HasPayedTime() bool`

HasPayedTime returns a boolean if a field has been set.

### GetCardType

`func (o *NotificationRequest) GetCardType() string`

GetCardType returns the CardType field if non-nil, zero value otherwise.

### GetCardTypeOk

`func (o *NotificationRequest) GetCardTypeOk() (*string, bool)`

GetCardTypeOk returns a tuple with the CardType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardType

`func (o *NotificationRequest) SetCardType(v string)`

SetCardType sets CardType field to given value.

### HasCardType

`func (o *NotificationRequest) HasCardType() bool`

HasCardType returns a boolean if a field has been set.

### GetIsPhysicalCard

`func (o *NotificationRequest) GetIsPhysicalCard() bool`

GetIsPhysicalCard returns the IsPhysicalCard field if non-nil, zero value otherwise.

### GetIsPhysicalCardOk

`func (o *NotificationRequest) GetIsPhysicalCardOk() (*bool, bool)`

GetIsPhysicalCardOk returns a tuple with the IsPhysicalCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPhysicalCard

`func (o *NotificationRequest) SetIsPhysicalCard(v bool)`

SetIsPhysicalCard sets IsPhysicalCard field to given value.

### HasIsPhysicalCard

`func (o *NotificationRequest) HasIsPhysicalCard() bool`

HasIsPhysicalCard returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


