// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.Hardware
//////////////////////////////////////////////
package hardware

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// CIM_RecordLog struct
type CIM_RecordLog struct {
	*CIM_Log

	//
	InstanceID string
}

func NewCIM_RecordLogEx1(instance *cim.WmiInstance) (newInstance *CIM_RecordLog, err error) {
	tmp, err := NewCIM_LogEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_RecordLog{
		CIM_Log: tmp,
	}
	return
}

func NewCIM_RecordLogEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_RecordLog, err error) {
	tmp, err := NewCIM_LogEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_RecordLog{
		CIM_Log: tmp,
	}
	return
}

// SetInstanceID sets the value of InstanceID for the instance
func (instance *CIM_RecordLog) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *CIM_RecordLog) GetPropertyInstanceID() (value string, err error) {
	retValue, err := instance.GetProperty("InstanceID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
