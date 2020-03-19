// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Win32_NamedJobObject struct
type Win32_NamedJobObject struct {
	*CIM_CollectionOfMSEs

	//
	BasicUIRestrictions uint32
}

func NewWin32_NamedJobObjectEx1(instance *cim.WmiInstance) (newInstance *Win32_NamedJobObject, err error) {
	tmp, err := NewCIM_CollectionOfMSEsEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_NamedJobObject{
		CIM_CollectionOfMSEs: tmp,
	}
	return
}

func NewWin32_NamedJobObjectEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_NamedJobObject, err error) {
	tmp, err := NewCIM_CollectionOfMSEsEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_NamedJobObject{
		CIM_CollectionOfMSEs: tmp,
	}
	return
}

// SetBasicUIRestrictions sets the value of BasicUIRestrictions for the instance
func (instance *Win32_NamedJobObject) SetPropertyBasicUIRestrictions(value uint32) (err error) {
	return instance.SetProperty("BasicUIRestrictions", value)
}

// GetBasicUIRestrictions gets the value of BasicUIRestrictions for the instance
func (instance *Win32_NamedJobObject) GetPropertyBasicUIRestrictions() (value uint32, err error) {
	retValue, err := instance.GetProperty("BasicUIRestrictions")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
