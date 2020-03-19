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

// MLNX_FirmwareIdentity struct
type MLNX_FirmwareIdentity struct {
	*CIM_SoftwareIdentity

	//
	Location string
}

func NewMLNX_FirmwareIdentityEx1(instance *cim.WmiInstance) (newInstance *MLNX_FirmwareIdentity, err error) {
	tmp, err := NewCIM_SoftwareIdentityEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MLNX_FirmwareIdentity{
		CIM_SoftwareIdentity: tmp,
	}
	return
}

func NewMLNX_FirmwareIdentityEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MLNX_FirmwareIdentity, err error) {
	tmp, err := NewCIM_SoftwareIdentityEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MLNX_FirmwareIdentity{
		CIM_SoftwareIdentity: tmp,
	}
	return
}

// SetLocation sets the value of Location for the instance
func (instance *MLNX_FirmwareIdentity) SetPropertyLocation(value string) (err error) {
	return instance.SetProperty("Location", value)
}

// GetLocation gets the value of Location for the instance
func (instance *MLNX_FirmwareIdentity) GetPropertyLocation() (value string, err error) {
	retValue, err := instance.GetProperty("Location")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
