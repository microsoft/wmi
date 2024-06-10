// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.WMI
//
// ////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSNdis_NetworkDirectVersion struct
type MSNdis_NetworkDirectVersion struct {
	*MSNdis

	//
	MajorVersionNumber uint16

	//
	MinorVersionNumber uint16
}

func NewMSNdis_NetworkDirectVersionEx1(instance *cim.WmiInstance) (newInstance *MSNdis_NetworkDirectVersion, err error) {
	tmp, err := NewMSNdisEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSNdis_NetworkDirectVersion{
		MSNdis: tmp,
	}
	return
}

func NewMSNdis_NetworkDirectVersionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSNdis_NetworkDirectVersion, err error) {
	tmp, err := NewMSNdisEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSNdis_NetworkDirectVersion{
		MSNdis: tmp,
	}
	return
}

// SetMajorVersionNumber sets the value of MajorVersionNumber for the instance
func (instance *MSNdis_NetworkDirectVersion) SetPropertyMajorVersionNumber(value uint16) (err error) {
	return instance.SetProperty("MajorVersionNumber", (value))
}

// GetMajorVersionNumber gets the value of MajorVersionNumber for the instance
func (instance *MSNdis_NetworkDirectVersion) GetPropertyMajorVersionNumber() (value uint16, err error) {
	retValue, err := instance.GetProperty("MajorVersionNumber")
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

// SetMinorVersionNumber sets the value of MinorVersionNumber for the instance
func (instance *MSNdis_NetworkDirectVersion) SetPropertyMinorVersionNumber(value uint16) (err error) {
	return instance.SetProperty("MinorVersionNumber", (value))
}

// GetMinorVersionNumber gets the value of MinorVersionNumber for the instance
func (instance *MSNdis_NetworkDirectVersion) GetPropertyMinorVersionNumber() (value uint16, err error) {
	retValue, err := instance.GetProperty("MinorVersionNumber")
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
