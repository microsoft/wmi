// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Msvm_VLANEndpoint struct
type Msvm_VLANEndpoint struct {
	*CIM_VLANEndpoint

	//
	SupportedEndpointModes []uint16
}

func NewMsvm_VLANEndpointEx1(instance *cim.WmiInstance) (newInstance *Msvm_VLANEndpoint, err error) {
	tmp, err := NewCIM_VLANEndpointEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_VLANEndpoint{
		CIM_VLANEndpoint: tmp,
	}
	return
}

func NewMsvm_VLANEndpointEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_VLANEndpoint, err error) {
	tmp, err := NewCIM_VLANEndpointEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_VLANEndpoint{
		CIM_VLANEndpoint: tmp,
	}
	return
}

// SetSupportedEndpointModes sets the value of SupportedEndpointModes for the instance
func (instance *Msvm_VLANEndpoint) SetPropertySupportedEndpointModes(value []uint16) (err error) {
	return instance.SetProperty("SupportedEndpointModes", value)
}

// GetSupportedEndpointModes gets the value of SupportedEndpointModes for the instance
func (instance *Msvm_VLANEndpoint) GetPropertySupportedEndpointModes() (value []uint16, err error) {
	retValue, err := instance.GetProperty("SupportedEndpointModes")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}
func (instance *Msvm_VLANEndpoint) GetRelatedLANEndpoint() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_LANEndpoint")
}
