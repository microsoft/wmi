// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.Microsoft.Windows.Storage
//////////////////////////////////////////////
package storage

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_TargetPortToTargetPortal struct
type MSFT_TargetPortToTargetPortal struct {
	*cim.WmiInstance

	//
	TargetPort MSFT_TargetPort

	//
	TargetPortal MSFT_TargetPortal
}

func NewMSFT_TargetPortToTargetPortalEx1(instance *cim.WmiInstance) (newInstance *MSFT_TargetPortToTargetPortal, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_TargetPortToTargetPortal{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_TargetPortToTargetPortalEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_TargetPortToTargetPortal, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_TargetPortToTargetPortal{
		WmiInstance: tmp,
	}
	return
}

// SetTargetPort sets the value of TargetPort for the instance
func (instance *MSFT_TargetPortToTargetPortal) SetPropertyTargetPort(value MSFT_TargetPort) (err error) {
	return instance.SetProperty("TargetPort", value)
}

// GetTargetPort gets the value of TargetPort for the instance
func (instance *MSFT_TargetPortToTargetPortal) GetPropertyTargetPort() (value MSFT_TargetPort, err error) {
	retValue, err := instance.GetProperty("TargetPort")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_TargetPort)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTargetPortal sets the value of TargetPortal for the instance
func (instance *MSFT_TargetPortToTargetPortal) SetPropertyTargetPortal(value MSFT_TargetPortal) (err error) {
	return instance.SetProperty("TargetPortal", value)
}

// GetTargetPortal gets the value of TargetPortal for the instance
func (instance *MSFT_TargetPortToTargetPortal) GetPropertyTargetPortal() (value MSFT_TargetPortal, err error) {
	retValue, err := instance.GetProperty("TargetPortal")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_TargetPortal)
	if !ok {
		// TODO: Set an error
	}
	return
}
