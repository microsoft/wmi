// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSiSCSIInitiator_SendTargetPortalClass struct
type MSiSCSIInitiator_SendTargetPortalClass struct {
	*cim.WmiInstance

	//
	InitiatorName string

	//
	InitiatorPortNumber uint32

	//
	LoginOptions MSiSCSIInitiator_TargetLoginOptions

	//
	PortalAddress string

	//
	PortalIdentifierString string

	//
	PortalPort uint16

	//
	PortalSymbolicName string

	//
	SecurityFlags uint64
}

func NewMSiSCSIInitiator_SendTargetPortalClassEx1(instance *cim.WmiInstance) (newInstance *MSiSCSIInitiator_SendTargetPortalClass, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSiSCSIInitiator_SendTargetPortalClass{
		WmiInstance: tmp,
	}
	return
}

func NewMSiSCSIInitiator_SendTargetPortalClassEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSiSCSIInitiator_SendTargetPortalClass, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSiSCSIInitiator_SendTargetPortalClass{
		WmiInstance: tmp,
	}
	return
}

// SetInitiatorName sets the value of InitiatorName for the instance
func (instance *MSiSCSIInitiator_SendTargetPortalClass) SetPropertyInitiatorName(value string) (err error) {
	return instance.SetProperty("InitiatorName", (value))
}

// GetInitiatorName gets the value of InitiatorName for the instance
func (instance *MSiSCSIInitiator_SendTargetPortalClass) GetPropertyInitiatorName() (value string, err error) {
	retValue, err := instance.GetProperty("InitiatorName")
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

// SetInitiatorPortNumber sets the value of InitiatorPortNumber for the instance
func (instance *MSiSCSIInitiator_SendTargetPortalClass) SetPropertyInitiatorPortNumber(value uint32) (err error) {
	return instance.SetProperty("InitiatorPortNumber", (value))
}

// GetInitiatorPortNumber gets the value of InitiatorPortNumber for the instance
func (instance *MSiSCSIInitiator_SendTargetPortalClass) GetPropertyInitiatorPortNumber() (value uint32, err error) {
	retValue, err := instance.GetProperty("InitiatorPortNumber")
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

// SetLoginOptions sets the value of LoginOptions for the instance
func (instance *MSiSCSIInitiator_SendTargetPortalClass) SetPropertyLoginOptions(value MSiSCSIInitiator_TargetLoginOptions) (err error) {
	return instance.SetProperty("LoginOptions", (value))
}

// GetLoginOptions gets the value of LoginOptions for the instance
func (instance *MSiSCSIInitiator_SendTargetPortalClass) GetPropertyLoginOptions() (value MSiSCSIInitiator_TargetLoginOptions, err error) {
	retValue, err := instance.GetProperty("LoginOptions")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(MSiSCSIInitiator_TargetLoginOptions)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " MSiSCSIInitiator_TargetLoginOptions is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = MSiSCSIInitiator_TargetLoginOptions(valuetmp)

	return
}

// SetPortalAddress sets the value of PortalAddress for the instance
func (instance *MSiSCSIInitiator_SendTargetPortalClass) SetPropertyPortalAddress(value string) (err error) {
	return instance.SetProperty("PortalAddress", (value))
}

// GetPortalAddress gets the value of PortalAddress for the instance
func (instance *MSiSCSIInitiator_SendTargetPortalClass) GetPropertyPortalAddress() (value string, err error) {
	retValue, err := instance.GetProperty("PortalAddress")
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

// SetPortalIdentifierString sets the value of PortalIdentifierString for the instance
func (instance *MSiSCSIInitiator_SendTargetPortalClass) SetPropertyPortalIdentifierString(value string) (err error) {
	return instance.SetProperty("PortalIdentifierString", (value))
}

// GetPortalIdentifierString gets the value of PortalIdentifierString for the instance
func (instance *MSiSCSIInitiator_SendTargetPortalClass) GetPropertyPortalIdentifierString() (value string, err error) {
	retValue, err := instance.GetProperty("PortalIdentifierString")
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

// SetPortalPort sets the value of PortalPort for the instance
func (instance *MSiSCSIInitiator_SendTargetPortalClass) SetPropertyPortalPort(value uint16) (err error) {
	return instance.SetProperty("PortalPort", (value))
}

// GetPortalPort gets the value of PortalPort for the instance
func (instance *MSiSCSIInitiator_SendTargetPortalClass) GetPropertyPortalPort() (value uint16, err error) {
	retValue, err := instance.GetProperty("PortalPort")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}

// SetPortalSymbolicName sets the value of PortalSymbolicName for the instance
func (instance *MSiSCSIInitiator_SendTargetPortalClass) SetPropertyPortalSymbolicName(value string) (err error) {
	return instance.SetProperty("PortalSymbolicName", (value))
}

// GetPortalSymbolicName gets the value of PortalSymbolicName for the instance
func (instance *MSiSCSIInitiator_SendTargetPortalClass) GetPropertyPortalSymbolicName() (value string, err error) {
	retValue, err := instance.GetProperty("PortalSymbolicName")
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

// SetSecurityFlags sets the value of SecurityFlags for the instance
func (instance *MSiSCSIInitiator_SendTargetPortalClass) SetPropertySecurityFlags(value uint64) (err error) {
	return instance.SetProperty("SecurityFlags", (value))
}

// GetSecurityFlags gets the value of SecurityFlags for the instance
func (instance *MSiSCSIInitiator_SendTargetPortalClass) GetPropertySecurityFlags() (value uint64, err error) {
	retValue, err := instance.GetProperty("SecurityFlags")
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

//

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSiSCSIInitiator_SendTargetPortalClass) Refresh() (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Refresh")
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
