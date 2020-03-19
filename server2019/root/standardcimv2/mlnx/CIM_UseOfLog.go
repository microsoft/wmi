// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.StandardCimv2.mlnx
//////////////////////////////////////////////
package mlnx

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// CIM_UseOfLog struct
type CIM_UseOfLog struct {
	*CIM_Dependency

	//
	RecordedData string
}

func NewCIM_UseOfLogEx1(instance *cim.WmiInstance) (newInstance *CIM_UseOfLog, err error) {
	tmp, err := NewCIM_DependencyEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_UseOfLog{
		CIM_Dependency: tmp,
	}
	return
}

func NewCIM_UseOfLogEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_UseOfLog, err error) {
	tmp, err := NewCIM_DependencyEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_UseOfLog{
		CIM_Dependency: tmp,
	}
	return
}

// SetRecordedData sets the value of RecordedData for the instance
func (instance *CIM_UseOfLog) SetPropertyRecordedData(value string) (err error) {
	return instance.SetProperty("RecordedData", value)
}

// GetRecordedData gets the value of RecordedData for the instance
func (instance *CIM_UseOfLog) GetPropertyRecordedData() (value string, err error) {
	retValue, err := instance.GetProperty("RecordedData")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
