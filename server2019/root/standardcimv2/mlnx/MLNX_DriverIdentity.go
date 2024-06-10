// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.StandardCimv2.mlnx
//////////////////////////////////////////////
package mlnx

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MLNX_DriverIdentity struct
type MLNX_DriverIdentity struct {
	*CIM_SoftwareIdentity

	//
	InstallLocation string
}

func NewMLNX_DriverIdentityEx1(instance *cim.WmiInstance) (newInstance *MLNX_DriverIdentity, err error) {
	tmp, err := NewCIM_SoftwareIdentityEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MLNX_DriverIdentity{
		CIM_SoftwareIdentity: tmp,
	}
	return
}

func NewMLNX_DriverIdentityEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MLNX_DriverIdentity, err error) {
	tmp, err := NewCIM_SoftwareIdentityEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MLNX_DriverIdentity{
		CIM_SoftwareIdentity: tmp,
	}
	return
}

// SetInstallLocation sets the value of InstallLocation for the instance
func (instance *MLNX_DriverIdentity) SetPropertyInstallLocation(value string) (err error) {
	return instance.SetProperty("InstallLocation", (value))
}

// GetInstallLocation gets the value of InstallLocation for the instance
func (instance *MLNX_DriverIdentity) GetPropertyInstallLocation() (value string, err error) {
	retValue, err := instance.GetProperty("InstallLocation")
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
