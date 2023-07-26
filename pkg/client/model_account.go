/*
 * DinoPay API
 *
 * This API allows sending and receiving payment through the DinoPay platform.
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// Account A DinoPay account
type Account struct {
	// Name of the account owner of the account
	AccountHolder string `json:"accountHolder"`
	// Number that identifies the account
	AccountNumber string `json:"accountNumber"`
}

// NewAccount instantiates a new Account object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccount(accountHolder string, accountNumber string, ) *Account {
	this := Account{}
	this.AccountHolder = accountHolder
	this.AccountNumber = accountNumber
	return &this
}

// NewAccountWithDefaults instantiates a new Account object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountWithDefaults() *Account {
	this := Account{}
	return &this
}

// GetAccountHolder returns the AccountHolder field value
func (o *Account) GetAccountHolder() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.AccountHolder
}

// GetAccountHolderOk returns a tuple with the AccountHolder field value
// and a boolean to check if the value has been set.
func (o *Account) GetAccountHolderOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccountHolder, true
}

// SetAccountHolder sets field value
func (o *Account) SetAccountHolder(v string) {
	o.AccountHolder = v
}

// GetAccountNumber returns the AccountNumber field value
func (o *Account) GetAccountNumber() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.AccountNumber
}

// GetAccountNumberOk returns a tuple with the AccountNumber field value
// and a boolean to check if the value has been set.
func (o *Account) GetAccountNumberOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccountNumber, true
}

// SetAccountNumber sets field value
func (o *Account) SetAccountNumber(v string) {
	o.AccountNumber = v
}

func (o Account) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["accountHolder"] = o.AccountHolder
	}
	if true {
		toSerialize["accountNumber"] = o.AccountNumber
	}
	return json.Marshal(toSerialize)
}

type NullableAccount struct {
	value *Account
	isSet bool
}

func (v NullableAccount) Get() *Account {
	return v.value
}

func (v *NullableAccount) Set(val *Account) {
	v.value = val
	v.isSet = true
}

func (v NullableAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccount(val *Account) *NullableAccount {
	return &NullableAccount{value: val, isSet: true}
}

func (v NullableAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

