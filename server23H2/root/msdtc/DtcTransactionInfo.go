// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.msdtc
//////////////////////////////////////////////
package msdtc

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// DtcTransactionInfo struct
type DtcTransactionInfo struct {
	*cim.WmiInstance

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

func NewDtcTransactionInfoEx1(instance *cim.WmiInstance) (newInstance *DtcTransactionInfo, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &DtcTransactionInfo{
		WmiInstance: tmp,
	}
	return
}

func NewDtcTransactionInfoEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *DtcTransactionInfo, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &DtcTransactionInfo{
		WmiInstance: tmp,
	}
	return
}

// SetDescription sets the value of Description for the instance
func (instance *DtcTransactionInfo) SetPropertyDescription(value string) (err error) {
	return instance.SetProperty("Description", (value))
}

// GetDescription gets the value of Description for the instance
func (instance *DtcTransactionInfo) GetPropertyDescription() (value string, err error) {
	retValue, err := instance.GetProperty("Description")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetIsolationLevel sets the value of IsolationLevel for the instance
func (instance *DtcTransactionInfo) SetPropertyIsolationLevel(value uint32) (err error) {
	return instance.SetProperty("IsolationLevel", (value))
}

// GetIsolationLevel gets the value of IsolationLevel for the instance
func (instance *DtcTransactionInfo) GetPropertyIsolationLevel() (value uint32, err error) {
	retValue, err := instance.GetProperty("IsolationLevel")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetParent sets the value of Parent for the instance
func (instance *DtcTransactionInfo) SetPropertyParent(value string) (err error) {
	return instance.SetProperty("Parent", (value))
}

// GetParent gets the value of Parent for the instance
func (instance *DtcTransactionInfo) GetPropertyParent() (value string, err error) {
	retValue, err := instance.GetProperty("Parent")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetState sets the value of State for the instance
func (instance *DtcTransactionInfo) SetPropertyState(value string) (err error) {
	return instance.SetProperty("State", (value))
}

// GetState gets the value of State for the instance
func (instance *DtcTransactionInfo) GetPropertyState() (value string, err error) {
	retValue, err := instance.GetProperty("State")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetTransactionId sets the value of TransactionId for the instance
func (instance *DtcTransactionInfo) SetPropertyTransactionId(value string) (err error) {
	return instance.SetProperty("TransactionId", (value))
}

// GetTransactionId gets the value of TransactionId for the instance
func (instance *DtcTransactionInfo) GetPropertyTransactionId() (value string, err error) {
	retValue, err := instance.GetProperty("TransactionId")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}
