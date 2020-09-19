// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_NetImPlatAdapter struct
type MSFT_NetImPlatAdapter struct {
	*CIM_EnabledLogicalElement

	//
	FailureReason uint32

	//
	InterfaceDescription string

	//
	NumberOfFailures uint32

	//
	ReceiveLinkSpeed uint64

	//
	Team string

	//
	TransmitLinkSpeed uint64
}

func NewMSFT_NetImPlatAdapterEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetImPlatAdapter, err error) {
	tmp, err := NewCIM_EnabledLogicalElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetImPlatAdapter{
		CIM_EnabledLogicalElement: tmp,
	}
	return
}

func NewMSFT_NetImPlatAdapterEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetImPlatAdapter, err error) {
	tmp, err := NewCIM_EnabledLogicalElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetImPlatAdapter{
		CIM_EnabledLogicalElement: tmp,
	}
	return
}

// SetFailureReason sets the value of FailureReason for the instance
func (instance *MSFT_NetImPlatAdapter) SetPropertyFailureReason(value uint32) (err error) {
	return instance.SetProperty("FailureReason", (value))
}

// GetFailureReason gets the value of FailureReason for the instance
func (instance *MSFT_NetImPlatAdapter) GetPropertyFailureReason() (value uint32, err error) {
	retValue, err := instance.GetProperty("FailureReason")
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

// SetInterfaceDescription sets the value of InterfaceDescription for the instance
func (instance *MSFT_NetImPlatAdapter) SetPropertyInterfaceDescription(value string) (err error) {
	return instance.SetProperty("InterfaceDescription", (value))
}

// GetInterfaceDescription gets the value of InterfaceDescription for the instance
func (instance *MSFT_NetImPlatAdapter) GetPropertyInterfaceDescription() (value string, err error) {
	retValue, err := instance.GetProperty("InterfaceDescription")
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

// SetNumberOfFailures sets the value of NumberOfFailures for the instance
func (instance *MSFT_NetImPlatAdapter) SetPropertyNumberOfFailures(value uint32) (err error) {
	return instance.SetProperty("NumberOfFailures", (value))
}

// GetNumberOfFailures gets the value of NumberOfFailures for the instance
func (instance *MSFT_NetImPlatAdapter) GetPropertyNumberOfFailures() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumberOfFailures")
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

// SetReceiveLinkSpeed sets the value of ReceiveLinkSpeed for the instance
func (instance *MSFT_NetImPlatAdapter) SetPropertyReceiveLinkSpeed(value uint64) (err error) {
	return instance.SetProperty("ReceiveLinkSpeed", (value))
}

// GetReceiveLinkSpeed gets the value of ReceiveLinkSpeed for the instance
func (instance *MSFT_NetImPlatAdapter) GetPropertyReceiveLinkSpeed() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReceiveLinkSpeed")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetTeam sets the value of Team for the instance
func (instance *MSFT_NetImPlatAdapter) SetPropertyTeam(value string) (err error) {
	return instance.SetProperty("Team", (value))
}

// GetTeam gets the value of Team for the instance
func (instance *MSFT_NetImPlatAdapter) GetPropertyTeam() (value string, err error) {
	retValue, err := instance.GetProperty("Team")
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

// SetTransmitLinkSpeed sets the value of TransmitLinkSpeed for the instance
func (instance *MSFT_NetImPlatAdapter) SetPropertyTransmitLinkSpeed(value uint64) (err error) {
	return instance.SetProperty("TransmitLinkSpeed", (value))
}

// GetTransmitLinkSpeed gets the value of TransmitLinkSpeed for the instance
func (instance *MSFT_NetImPlatAdapter) GetPropertyTransmitLinkSpeed() (value uint64, err error) {
	retValue, err := instance.GetProperty("TransmitLinkSpeed")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}
