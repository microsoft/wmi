// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.msdtc
//////////////////////////////////////////////
package msdtc

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// DtcTransactionInfo struct
type DtcTransactionInfo struct {
	cim.WmiInstance

	//
	Description string

	//
	IsolationLevel uint32

	//
	Parent string

	//
	State string

	//
	TransactionId string
}

// SetDescription sets the value of Description for the instance
func (instance *DtcTransactionInfo) SetPropertyDescription(value string) (err error) {
	return instance.SetProperty("Description", value)
}

// GetDescription gets the value of Description for the instance
func (instance *DtcTransactionInfo) GetPropertyDescription() (value string, err error) {
	retValue, err := instance.GetProperty("Description")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIsolationLevel sets the value of IsolationLevel for the instance
func (instance *DtcTransactionInfo) SetPropertyIsolationLevel(value uint32) (err error) {
	return instance.SetProperty("IsolationLevel", value)
}

// GetIsolationLevel gets the value of IsolationLevel for the instance
func (instance *DtcTransactionInfo) GetPropertyIsolationLevel() (value uint32, err error) {
	retValue, err := instance.GetProperty("IsolationLevel")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetParent sets the value of Parent for the instance
func (instance *DtcTransactionInfo) SetPropertyParent(value string) (err error) {
	return instance.SetProperty("Parent", value)
}

// GetParent gets the value of Parent for the instance
func (instance *DtcTransactionInfo) GetPropertyParent() (value string, err error) {
	retValue, err := instance.GetProperty("Parent")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetState sets the value of State for the instance
func (instance *DtcTransactionInfo) SetPropertyState(value string) (err error) {
	return instance.SetProperty("State", value)
}

// GetState gets the value of State for the instance
func (instance *DtcTransactionInfo) GetPropertyState() (value string, err error) {
	retValue, err := instance.GetProperty("State")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTransactionId sets the value of TransactionId for the instance
func (instance *DtcTransactionInfo) SetPropertyTransactionId(value string) (err error) {
	return instance.SetProperty("TransactionId", value)
}

// GetTransactionId gets the value of TransactionId for the instance
func (instance *DtcTransactionInfo) GetPropertyTransactionId() (value string, err error) {
	retValue, err := instance.GetProperty("TransactionId")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
