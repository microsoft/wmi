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

// MLNX_SoftwareIdentity struct
type MLNX_SoftwareIdentity struct {
	*CIM_SoftwareIdentity

	//
	InstallLocation string
}

func NewMLNX_SoftwareIdentityEx1(instance *cim.WmiInstance) (newInstance *MLNX_SoftwareIdentity, err error) {
	tmp, err := NewCIM_SoftwareIdentityEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MLNX_SoftwareIdentity{
		CIM_SoftwareIdentity: tmp,
	}
	return
}

func NewMLNX_SoftwareIdentityEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MLNX_SoftwareIdentity, err error) {
	tmp, err := NewCIM_SoftwareIdentityEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MLNX_SoftwareIdentity{
		CIM_SoftwareIdentity: tmp,
	}
	return
}

// SetInstallLocation sets the value of InstallLocation for the instance
func (instance *MLNX_SoftwareIdentity) SetPropertyInstallLocation(value string) (err error) {
	return instance.SetProperty("InstallLocation", value)
}

// GetInstallLocation gets the value of InstallLocation for the instance
func (instance *MLNX_SoftwareIdentity) GetPropertyInstallLocation() (value string, err error) {
	retValue, err := instance.GetProperty("InstallLocation")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
