// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.HyperVCluster.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// CIM_DVDDrive struct
type CIM_DVDDrive struct {
	*CIM_MediaAccessDevice

	// The CD and DVD formats that are supported by this Device. For example, the Drive may support "CD-ROM" and "DVD-RAM". In this case, the values 16 and 24 would be written to the array. This property's values align with those defined in PhysicalMedia.MediaType.
	FormatsSupported []DVDDrive_FormatsSupported
}

func NewCIM_DVDDriveEx1(instance *cim.WmiInstance) (newInstance *CIM_DVDDrive, err error) {
	tmp, err := NewCIM_MediaAccessDeviceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_DVDDrive{
		CIM_MediaAccessDevice: tmp,
	}
	return
}

func NewCIM_DVDDriveEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_DVDDrive, err error) {
	tmp, err := NewCIM_MediaAccessDeviceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_DVDDrive{
		CIM_MediaAccessDevice: tmp,
	}
	return
}

// SetFormatsSupported sets the value of FormatsSupported for the instance
func (instance *CIM_DVDDrive) SetPropertyFormatsSupported(value []DVDDrive_FormatsSupported) (err error) {
	return instance.SetProperty("FormatsSupported", (value))
}

// GetFormatsSupported gets the value of FormatsSupported for the instance
func (instance *CIM_DVDDrive) GetPropertyFormatsSupported() (value []DVDDrive_FormatsSupported, err error) {
	retValue, err := instance.GetProperty("FormatsSupported")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(int32)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, DVDDrive_FormatsSupported(valuetmp))
	}

	return
}
