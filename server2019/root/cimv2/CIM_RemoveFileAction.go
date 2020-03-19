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

// CIM_RemoveFileAction struct
type CIM_RemoveFileAction struct {
	*CIM_FileAction

	//
	File string
}

func NewCIM_RemoveFileActionEx1(instance *cim.WmiInstance) (newInstance *CIM_RemoveFileAction, err error) {
	tmp, err := NewCIM_FileActionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_RemoveFileAction{
		CIM_FileAction: tmp,
	}
	return
}

func NewCIM_RemoveFileActionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_RemoveFileAction, err error) {
	tmp, err := NewCIM_FileActionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_RemoveFileAction{
		CIM_FileAction: tmp,
	}
	return
}

// SetFile sets the value of File for the instance
func (instance *CIM_RemoveFileAction) SetPropertyFile(value string) (err error) {
	return instance.SetProperty("File", value)
}

// GetFile gets the value of File for the instance
func (instance *CIM_RemoveFileAction) GetPropertyFile() (value string, err error) {
	retValue, err := instance.GetProperty("File")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
