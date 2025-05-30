# PreparePaymentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OctoShopId** | **int32** |  | 
**OctoSecret** | **string** |  | 
**ShopTransactionId** | **string** |  | 
**AutoCapture** | **bool** |  | [default to false]
**Test** | **bool** |  | 
**InitTime** | **string** |  | 
**UserData** | Pointer to [**UserData**](UserData.md) |  | [optional] 
**TotalSum** | **float64** |  | 
**Currency** | **string** |  | 
**Description** | **string** |  | 
**Basket** | Pointer to [**[]BasketItem**](BasketItem.md) |  | [optional] 
**PaymentMethods** | Pointer to [**[]PaymentMethod**](PaymentMethod.md) |  | [optional] 
**TspId** | Pointer to **int32** |  | [optional] 
**ReturnUrl** | **string** |  | 
**NotifyUrl** | **string** |  | 
**Language** | **string** |  | 
**Ttl** | Pointer to **int32** |  | [optional] 

## Methods

### NewPreparePaymentRequest

`func NewPreparePaymentRequest(octoShopId int32, octoSecret string, shopTransactionId string, autoCapture bool, test bool, initTime string, totalSum float64, currency string, description string, returnUrl string, notifyUrl string, language string, ) *PreparePaymentRequest`

NewPreparePaymentRequest instantiates a new PreparePaymentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPreparePaymentRequestWithDefaults

`func NewPreparePaymentRequestWithDefaults() *PreparePaymentRequest`

NewPreparePaymentRequestWithDefaults instantiates a new PreparePaymentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOctoShopId

`func (o *PreparePaymentRequest) GetOctoShopId() int32`

GetOctoShopId returns the OctoShopId field if non-nil, zero value otherwise.

### GetOctoShopIdOk

`func (o *PreparePaymentRequest) GetOctoShopIdOk() (*int32, bool)`

GetOctoShopIdOk returns a tuple with the OctoShopId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOctoShopId

`func (o *PreparePaymentRequest) SetOctoShopId(v int32)`

SetOctoShopId sets OctoShopId field to given value.


### GetOctoSecret

`func (o *PreparePaymentRequest) GetOctoSecret() string`

GetOctoSecret returns the OctoSecret field if non-nil, zero value otherwise.

### GetOctoSecretOk

`func (o *PreparePaymentRequest) GetOctoSecretOk() (*string, bool)`

GetOctoSecretOk returns a tuple with the OctoSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOctoSecret

`func (o *PreparePaymentRequest) SetOctoSecret(v string)`

SetOctoSecret sets OctoSecret field to given value.


### GetShopTransactionId

`func (o *PreparePaymentRequest) GetShopTransactionId() string`

GetShopTransactionId returns the ShopTransactionId field if non-nil, zero value otherwise.

### GetShopTransactionIdOk

`func (o *PreparePaymentRequest) GetShopTransactionIdOk() (*string, bool)`

GetShopTransactionIdOk returns a tuple with the ShopTransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopTransactionId

`func (o *PreparePaymentRequest) SetShopTransactionId(v string)`

SetShopTransactionId sets ShopTransactionId field to given value.


### GetAutoCapture

`func (o *PreparePaymentRequest) GetAutoCapture() bool`

GetAutoCapture returns the AutoCapture field if non-nil, zero value otherwise.

### GetAutoCaptureOk

`func (o *PreparePaymentRequest) GetAutoCaptureOk() (*bool, bool)`

GetAutoCaptureOk returns a tuple with the AutoCapture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoCapture

`func (o *PreparePaymentRequest) SetAutoCapture(v bool)`

SetAutoCapture sets AutoCapture field to given value.


### GetTest

`func (o *PreparePaymentRequest) GetTest() bool`

GetTest returns the Test field if non-nil, zero value otherwise.

### GetTestOk

`func (o *PreparePaymentRequest) GetTestOk() (*bool, bool)`

GetTestOk returns a tuple with the Test field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTest

`func (o *PreparePaymentRequest) SetTest(v bool)`

SetTest sets Test field to given value.


### GetInitTime

`func (o *PreparePaymentRequest) GetInitTime() string`

GetInitTime returns the InitTime field if non-nil, zero value otherwise.

### GetInitTimeOk

`func (o *PreparePaymentRequest) GetInitTimeOk() (*string, bool)`

GetInitTimeOk returns a tuple with the InitTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitTime

`func (o *PreparePaymentRequest) SetInitTime(v string)`

SetInitTime sets InitTime field to given value.


### GetUserData

`func (o *PreparePaymentRequest) GetUserData() UserData`

GetUserData returns the UserData field if non-nil, zero value otherwise.

### GetUserDataOk

`func (o *PreparePaymentRequest) GetUserDataOk() (*UserData, bool)`

GetUserDataOk returns a tuple with the UserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserData

`func (o *PreparePaymentRequest) SetUserData(v UserData)`

SetUserData sets UserData field to given value.

### HasUserData

`func (o *PreparePaymentRequest) HasUserData() bool`

HasUserData returns a boolean if a field has been set.

### GetTotalSum

`func (o *PreparePaymentRequest) GetTotalSum() float64`

GetTotalSum returns the TotalSum field if non-nil, zero value otherwise.

### GetTotalSumOk

`func (o *PreparePaymentRequest) GetTotalSumOk() (*float64, bool)`

GetTotalSumOk returns a tuple with the TotalSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSum

`func (o *PreparePaymentRequest) SetTotalSum(v float64)`

SetTotalSum sets TotalSum field to given value.


### GetCurrency

`func (o *PreparePaymentRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PreparePaymentRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PreparePaymentRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetDescription

`func (o *PreparePaymentRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PreparePaymentRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PreparePaymentRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetBasket

`func (o *PreparePaymentRequest) GetBasket() []BasketItem`

GetBasket returns the Basket field if non-nil, zero value otherwise.

### GetBasketOk

`func (o *PreparePaymentRequest) GetBasketOk() (*[]BasketItem, bool)`

GetBasketOk returns a tuple with the Basket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasket

`func (o *PreparePaymentRequest) SetBasket(v []BasketItem)`

SetBasket sets Basket field to given value.

### HasBasket

`func (o *PreparePaymentRequest) HasBasket() bool`

HasBasket returns a boolean if a field has been set.

### GetPaymentMethods

`func (o *PreparePaymentRequest) GetPaymentMethods() []PaymentMethod`

GetPaymentMethods returns the PaymentMethods field if non-nil, zero value otherwise.

### GetPaymentMethodsOk

`func (o *PreparePaymentRequest) GetPaymentMethodsOk() (*[]PaymentMethod, bool)`

GetPaymentMethodsOk returns a tuple with the PaymentMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethods

`func (o *PreparePaymentRequest) SetPaymentMethods(v []PaymentMethod)`

SetPaymentMethods sets PaymentMethods field to given value.

### HasPaymentMethods

`func (o *PreparePaymentRequest) HasPaymentMethods() bool`

HasPaymentMethods returns a boolean if a field has been set.

### GetTspId

`func (o *PreparePaymentRequest) GetTspId() int32`

GetTspId returns the TspId field if non-nil, zero value otherwise.

### GetTspIdOk

`func (o *PreparePaymentRequest) GetTspIdOk() (*int32, bool)`

GetTspIdOk returns a tuple with the TspId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTspId

`func (o *PreparePaymentRequest) SetTspId(v int32)`

SetTspId sets TspId field to given value.

### HasTspId

`func (o *PreparePaymentRequest) HasTspId() bool`

HasTspId returns a boolean if a field has been set.

### GetReturnUrl

`func (o *PreparePaymentRequest) GetReturnUrl() string`

GetReturnUrl returns the ReturnUrl field if non-nil, zero value otherwise.

### GetReturnUrlOk

`func (o *PreparePaymentRequest) GetReturnUrlOk() (*string, bool)`

GetReturnUrlOk returns a tuple with the ReturnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnUrl

`func (o *PreparePaymentRequest) SetReturnUrl(v string)`

SetReturnUrl sets ReturnUrl field to given value.


### GetNotifyUrl

`func (o *PreparePaymentRequest) GetNotifyUrl() string`

GetNotifyUrl returns the NotifyUrl field if non-nil, zero value otherwise.

### GetNotifyUrlOk

`func (o *PreparePaymentRequest) GetNotifyUrlOk() (*string, bool)`

GetNotifyUrlOk returns a tuple with the NotifyUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyUrl

`func (o *PreparePaymentRequest) SetNotifyUrl(v string)`

SetNotifyUrl sets NotifyUrl field to given value.


### GetLanguage

`func (o *PreparePaymentRequest) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *PreparePaymentRequest) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *PreparePaymentRequest) SetLanguage(v string)`

SetLanguage sets Language field to given value.


### GetTtl

`func (o *PreparePaymentRequest) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *PreparePaymentRequest) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *PreparePaymentRequest) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *PreparePaymentRequest) HasTtl() bool`

HasTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


