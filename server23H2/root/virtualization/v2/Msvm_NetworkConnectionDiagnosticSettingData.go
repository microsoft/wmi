// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Msvm_NetworkConnectionDiagnosticSettingData struct
type Msvm_NetworkConnectionDiagnosticSettingData struct {
	*CIM_SettingData

	//
	IsolationId uint32

	//
	IsSender bool

	//
	PayloadSize uint32

	//
	ReceiverIP string

	//
	ReceiverMac string

	//
	SenderIP string

	//
	SequenceNumber uint32
}

func NewMsvm_NetworkConnectionDiagnosticSettingDataEx1(instance *cim.WmiInstance) (newInstance *Msvm_NetworkConnectionDiagnosticSettingData, err error) {
	tmp, err := NewCIM_SettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_NetworkConnectionDiagnosticSettingData{
		CIM_SettingData: tmp,
	}
	return
}

func NewMsvm_NetworkConnectionDiagnosticSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_NetworkConnectionDiagnosticSettingData, err error) {
	tmp, err := NewCIM_SettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_NetworkConnectionDiagnosticSettingData{
		CIM_SettingData: tmp,
	}
	return
}

// SetIsolationId sets the value of IsolationId for the instance
func (instance *Msvm_NetworkConnectionDiagnosticSettingData) SetPropertyIsolationId(value uint32) (err error) {
	return instance.SetProperty("IsolationId", (value))
}

// GetIsolationId gets the value of IsolationId for the instance
func (instance *Msvm_NetworkConnectionDiagnosticSettingData) GetPropertyIsolationId() (value uint32, err error) {
	retValue, err := instance.GetProperty("IsolationId")
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

// SetIsSender sets the value of IsSender for the instance
func (instance *Msvm_NetworkConnectionDiagnosticSettingData) SetPropertyIsSender(value bool) (err error) {
	return instance.SetProperty("IsSender", (value))
}

// GetIsSender gets the value of IsSender for the instance
func (instance *Msvm_NetworkConnectionDiagnosticSettingData) GetPropertyIsSender() (value bool, err error) {
	retValue, err := instance.GetProperty("IsSender")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetPayloadSize sets the value of PayloadSize for the instance
func (instance *Msvm_NetworkConnectionDiagnosticSettingData) SetPropertyPayloadSize(value uint32) (err error) {
	return instance.SetProperty("PayloadSize", (value))
}

// GetPayloadSize gets the value of PayloadSize for the instance
func (instance *Msvm_NetworkConnectionDiagnosticSettingData) GetPropertyPayloadSize() (value uint32, err error) {
	retValue, err := instance.GetProperty("PayloadSize")
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

// SetReceiverIP sets the value of ReceiverIP for the instance
func (instance *Msvm_NetworkConnectionDiagnosticSettingData) SetPropertyReceiverIP(value string) (err error) {
	return instance.SetProperty("ReceiverIP", (value))
}

// GetReceiverIP gets the value of ReceiverIP for the instance
func (instance *Msvm_NetworkConnectionDiagnosticSettingData) GetPropertyReceiverIP() (value string, err error) {
	retValue, err := instance.GetProperty("ReceiverIP")
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

// SetReceiverMac sets the value of ReceiverMac for the instance
func (instance *Msvm_NetworkConnectionDiagnosticSettingData) SetPropertyReceiverMac(value string) (err error) {
	return instance.SetProperty("ReceiverMac", (value))
}

// GetReceiverMac gets the value of ReceiverMac for the instance
func (instance *Msvm_NetworkConnectionDiagnosticSettingData) GetPropertyReceiverMac() (value string, err error) {
	retValue, err := instance.GetProperty("ReceiverMac")
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

// SetSenderIP sets the value of SenderIP for the instance
func (instance *Msvm_NetworkConnectionDiagnosticSettingData) SetPropertySenderIP(value string) (err error) {
	return instance.SetProperty("SenderIP", (value))
}

// GetSenderIP gets the value of SenderIP for the instance
func (instance *Msvm_NetworkConnectionDiagnosticSettingData) GetPropertySenderIP() (value string, err error) {
	retValue, err := instance.GetProperty("SenderIP")
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

// SetSequenceNumber sets the value of SequenceNumber for the instance
func (instance *Msvm_NetworkConnectionDiagnosticSettingData) SetPropertySequenceNumber(value uint32) (err error) {
	return instance.SetProperty("SequenceNumber", (value))
}

// GetSequenceNumber gets the value of SequenceNumber for the instance
func (instance *Msvm_NetworkConnectionDiagnosticSettingData) GetPropertySequenceNumber() (value uint32, err error) {
	retValue, err := instance.GetProperty("SequenceNumber")
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
