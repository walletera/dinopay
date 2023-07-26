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
	"fmt"
)

// PaymentStatus Status of the payment
type PaymentStatus string

// List of paymentStatus
const (
	PENDING PaymentStatus = "pending"
	CONFIRMED PaymentStatus = "confirmed"
	REJECTED PaymentStatus = "rejected"
)

func (v *PaymentStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PaymentStatus(value)
	for _, existing := range []PaymentStatus{ "pending", "confirmed", "rejected",   } {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PaymentStatus", value)
}

// Ptr returns reference to paymentStatus value
func (v PaymentStatus) Ptr() *PaymentStatus {
	return &v
}

type NullablePaymentStatus struct {
	value *PaymentStatus
	isSet bool
}

func (v NullablePaymentStatus) Get() *PaymentStatus {
	return v.value
}

func (v *NullablePaymentStatus) Set(val *PaymentStatus) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentStatus) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentStatus(val *PaymentStatus) *NullablePaymentStatus {
	return &NullablePaymentStatus{value: val, isSet: true}
}

func (v NullablePaymentStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

