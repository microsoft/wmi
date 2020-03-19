// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Msvm_InteractiveSessionACE struct
type Msvm_InteractiveSessionACE struct {
	*cim.WmiInstance

	//
	AccessType uint16

	//
	Trustee string
}

func NewMsvm_InteractiveSessionACEEx1(instance *cim.WmiInstance) (newInstance *Msvm_InteractiveSessionACE, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &Msvm_InteractiveSessionACE{
		WmiInstance: tmp,
	}
	return
}

func NewMsvm_InteractiveSessionACEEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_InteractiveSessionACE, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_InteractiveSessionACE{
		WmiInstance: tmp,
	}
	return
}

// SetAccessType sets the value of AccessType for the instance
func (instance *Msvm_InteractiveSessionACE) SetPropertyAccessType(value uint16) (err error) {
	return instance.SetProperty("AccessType", value)
}

// GetAccessType gets the value of AccessType for the instance
func (instance *Msvm_InteractiveSessionACE) GetPropertyAccessType() (value uint16, err error) {
	retValue, err := instance.GetProperty("AccessType")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTrustee sets the value of Trustee for the instance
func (instance *Msvm_InteractiveSessionACE) SetPropertyTrustee(value string) (err error) {
	return instance.SetProperty("Trustee", value)
}

// GetTrustee gets the value of Trustee for the instance
func (instance *Msvm_InteractiveSessionACE) GetPropertyTrustee() (value string, err error) {
	retValue, err := instance.GetProperty("Trustee")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
