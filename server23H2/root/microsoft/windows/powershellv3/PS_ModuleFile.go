// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.Microsoft.Windows.Powershellv3
//////////////////////////////////////////////
package powershellv3

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// PS_ModuleFile struct
type PS_ModuleFile struct {
	*CIM_ManagedElement

	//
	FileData []uint8

	//
	FileName string
}

func NewPS_ModuleFileEx1(instance *cim.WmiInstance) (newInstance *PS_ModuleFile, err error) {
	tmp, err := NewCIM_ManagedElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &PS_ModuleFile{
		CIM_ManagedElement: tmp,
	}
	return
}

func NewPS_ModuleFileEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *PS_ModuleFile, err error) {
	tmp, err := NewCIM_ManagedElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &PS_ModuleFile{
		CIM_ManagedElement: tmp,
	}
	return
}

// SetFileData sets the value of FileData for the instance
func (instance *PS_ModuleFile) SetPropertyFileData(value []uint8) (err error) {
	return instance.SetProperty("FileData", (value))
}

// GetFileData gets the value of FileData for the instance
func (instance *PS_ModuleFile) GetPropertyFileData() (value []uint8, err error) {
	retValue, err := instance.GetProperty("FileData")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint8)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint8(valuetmp))
	}

	return
}

// SetFileName sets the value of FileName for the instance
func (instance *PS_ModuleFile) SetPropertyFileName(value string) (err error) {
	return instance.SetProperty("FileName", (value))
}

// GetFileName gets the value of FileName for the instance
func (instance *PS_ModuleFile) GetPropertyFileName() (value string, err error) {
	retValue, err := instance.GetProperty("FileName")
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
