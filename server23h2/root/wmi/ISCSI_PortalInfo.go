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

// ISCSI_PortalInfo struct
type ISCSI_PortalInfo struct {
	*cim.WmiInstance

	// An integer used to uniquely identify a paticular port
	Index uint32

	// The portal's network address
	IPAddr ISCSI_IP_Address

	// The portal's socket number
	Port uint32

	// The portal's aggregation tag
	PortalTag uint16

	// **typedef** The type of portal (Initiator or Target)
	///
	PortalType PortalInfo_PortalType

	// The portal's transport protocol
	Protocol PortalInfo_Protocol

	//
	Reserved1 uint8

	//
	Reserved2 uint8
}

func NewISCSI_PortalInfoEx1(instance *cim.WmiInstance) (newInstance *ISCSI_PortalInfo, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &ISCSI_PortalInfo{
		WmiInstance: tmp,
	}
	return
}

func NewISCSI_PortalInfoEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ISCSI_PortalInfo, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ISCSI_PortalInfo{
		WmiInstance: tmp,
	}
	return
}

// SetIndex sets the value of Index for the instance
func (instance *ISCSI_PortalInfo) SetPropertyIndex(value uint32) (err error) {
	return instance.SetProperty("Index", (value))
}

// GetIndex gets the value of Index for the instance
func (instance *ISCSI_PortalInfo) GetPropertyIndex() (value uint32, err error) {
	retValue, err := instance.GetProperty("Index")
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

// SetIPAddr sets the value of IPAddr for the instance
func (instance *ISCSI_PortalInfo) SetPropertyIPAddr(value ISCSI_IP_Address) (err error) {
	return instance.SetProperty("IPAddr", (value))
}

// GetIPAddr gets the value of IPAddr for the instance
func (instance *ISCSI_PortalInfo) GetPropertyIPAddr() (value ISCSI_IP_Address, err error) {
	retValue, err := instance.GetProperty("IPAddr")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(ISCSI_IP_Address)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " ISCSI_IP_Address is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = ISCSI_IP_Address(valuetmp)

	return
}

// SetPort sets the value of Port for the instance
func (instance *ISCSI_PortalInfo) SetPropertyPort(value uint32) (err error) {
	return instance.SetProperty("Port", (value))
}

// GetPort gets the value of Port for the instance
func (instance *ISCSI_PortalInfo) GetPropertyPort() (value uint32, err error) {
	retValue, err := instance.GetProperty("Port")
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

// SetPortalTag sets the value of PortalTag for the instance
func (instance *ISCSI_PortalInfo) SetPropertyPortalTag(value uint16) (err error) {
	return instance.SetProperty("PortalTag", (value))
}

// GetPortalTag gets the value of PortalTag for the instance
func (instance *ISCSI_PortalInfo) GetPropertyPortalTag() (value uint16, err error) {
	retValue, err := instance.GetProperty("PortalTag")
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

// SetPortalType sets the value of PortalType for the instance
func (instance *ISCSI_PortalInfo) SetPropertyPortalType(value PortalInfo_PortalType) (err error) {
	return instance.SetProperty("PortalType", (value))
}

// GetPortalType gets the value of PortalType for the instance
func (instance *ISCSI_PortalInfo) GetPropertyPortalType() (value PortalInfo_PortalType, err error) {
	retValue, err := instance.GetProperty("PortalType")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = PortalInfo_PortalType(valuetmp)

	return
}

// SetProtocol sets the value of Protocol for the instance
func (instance *ISCSI_PortalInfo) SetPropertyProtocol(value PortalInfo_Protocol) (err error) {
	return instance.SetProperty("Protocol", (value))
}

// GetProtocol gets the value of Protocol for the instance
func (instance *ISCSI_PortalInfo) GetPropertyProtocol() (value PortalInfo_Protocol, err error) {
	retValue, err := instance.GetProperty("Protocol")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = PortalInfo_Protocol(valuetmp)

	return
}

// SetReserved1 sets the value of Reserved1 for the instance
func (instance *ISCSI_PortalInfo) SetPropertyReserved1(value uint8) (err error) {
	return instance.SetProperty("Reserved1", (value))
}

// GetReserved1 gets the value of Reserved1 for the instance
func (instance *ISCSI_PortalInfo) GetPropertyReserved1() (value uint8, err error) {
	retValue, err := instance.GetProperty("Reserved1")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint8(valuetmp)

	return
}

// SetReserved2 sets the value of Reserved2 for the instance
func (instance *ISCSI_PortalInfo) SetPropertyReserved2(value uint8) (err error) {
	return instance.SetProperty("Reserved2", (value))
}

// GetReserved2 gets the value of Reserved2 for the instance
func (instance *ISCSI_PortalInfo) GetPropertyReserved2() (value uint8, err error) {
	retValue, err := instance.GetProperty("Reserved2")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint8(valuetmp)

	return
}
