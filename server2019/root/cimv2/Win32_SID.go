// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Win32_SID struct
type Win32_SID struct {
	*cim.WmiInstance

	//
	AccountName string

	//
	BinaryRepresentation []uint8

	//
	ReferencedDomainName string

	//
	SID string

	//
	SidLength uint32
}

func NewWin32_SIDEx1(instance *cim.WmiInstance) (newInstance *Win32_SID, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &Win32_SID{
		WmiInstance: tmp,
	}
	return
}

func NewWin32_SIDEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_SID, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_SID{
		WmiInstance: tmp,
	}
	return
}

// SetAccountName sets the value of AccountName for the instance
func (instance *Win32_SID) SetPropertyAccountName(value string) (err error) {
	return instance.SetProperty("AccountName", value)
}

// GetAccountName gets the value of AccountName for the instance
func (instance *Win32_SID) GetPropertyAccountName() (value string, err error) {
	retValue, err := instance.GetProperty("AccountName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetBinaryRepresentation sets the value of BinaryRepresentation for the instance
func (instance *Win32_SID) SetPropertyBinaryRepresentation(value []uint8) (err error) {
	return instance.SetProperty("BinaryRepresentation", value)
}

// GetBinaryRepresentation gets the value of BinaryRepresentation for the instance
func (instance *Win32_SID) GetPropertyBinaryRepresentation() (value []uint8, err error) {
	retValue, err := instance.GetProperty("BinaryRepresentation")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetReferencedDomainName sets the value of ReferencedDomainName for the instance
func (instance *Win32_SID) SetPropertyReferencedDomainName(value string) (err error) {
	return instance.SetProperty("ReferencedDomainName", value)
}

// GetReferencedDomainName gets the value of ReferencedDomainName for the instance
func (instance *Win32_SID) GetPropertyReferencedDomainName() (value string, err error) {
	retValue, err := instance.GetProperty("ReferencedDomainName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSID sets the value of SID for the instance
func (instance *Win32_SID) SetPropertySID(value string) (err error) {
	return instance.SetProperty("SID", value)
}

// GetSID gets the value of SID for the instance
func (instance *Win32_SID) GetPropertySID() (value string, err error) {
	retValue, err := instance.GetProperty("SID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSidLength sets the value of SidLength for the instance
func (instance *Win32_SID) SetPropertySidLength(value uint32) (err error) {
	return instance.SetProperty("SidLength", value)
}

// GetSidLength gets the value of SidLength for the instance
func (instance *Win32_SID) GetPropertySidLength() (value uint32, err error) {
	retValue, err := instance.GetProperty("SidLength")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
