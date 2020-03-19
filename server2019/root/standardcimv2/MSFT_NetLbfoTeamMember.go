// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_NetLbfoTeamMember struct
type MSFT_NetLbfoTeamMember struct {
	*MSFT_NetImPlatAdapter

	// 396
	AdministrativeMode uint32

	// 397
	OperationalMode uint32
}

func NewMSFT_NetLbfoTeamMemberEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetLbfoTeamMember, err error) {
	tmp, err := NewMSFT_NetImPlatAdapterEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetLbfoTeamMember{
		MSFT_NetImPlatAdapter: tmp,
	}
	return
}

func NewMSFT_NetLbfoTeamMemberEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetLbfoTeamMember, err error) {
	tmp, err := NewMSFT_NetImPlatAdapterEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetLbfoTeamMember{
		MSFT_NetImPlatAdapter: tmp,
	}
	return
}

// SetAdministrativeMode sets the value of AdministrativeMode for the instance
func (instance *MSFT_NetLbfoTeamMember) SetPropertyAdministrativeMode(value uint32) (err error) {
	return instance.SetProperty("AdministrativeMode", value)
}

// GetAdministrativeMode gets the value of AdministrativeMode for the instance
func (instance *MSFT_NetLbfoTeamMember) GetPropertyAdministrativeMode() (value uint32, err error) {
	retValue, err := instance.GetProperty("AdministrativeMode")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOperationalMode sets the value of OperationalMode for the instance
func (instance *MSFT_NetLbfoTeamMember) SetPropertyOperationalMode(value uint32) (err error) {
	return instance.SetProperty("OperationalMode", value)
}

// GetOperationalMode gets the value of OperationalMode for the instance
func (instance *MSFT_NetLbfoTeamMember) GetPropertyOperationalMode() (value uint32, err error) {
	retValue, err := instance.GetProperty("OperationalMode")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
