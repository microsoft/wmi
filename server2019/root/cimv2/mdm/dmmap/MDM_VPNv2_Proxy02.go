// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.CIMV2.mdm.dmmap
//////////////////////////////////////////////
package dmmap

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MDM_VPNv2_Proxy02 struct
type MDM_VPNv2_Proxy02 struct {
	*cim.WmiInstance

	//
	AutoConfigUrl string

	//
	InstanceID string

	//
	ParentID string
}

func NewMDM_VPNv2_Proxy02Ex1(instance *cim.WmiInstance) (newInstance *MDM_VPNv2_Proxy02, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_VPNv2_Proxy02{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_VPNv2_Proxy02Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_VPNv2_Proxy02, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_VPNv2_Proxy02{
		WmiInstance: tmp,
	}
	return
}

// SetAutoConfigUrl sets the value of AutoConfigUrl for the instance
func (instance *MDM_VPNv2_Proxy02) SetPropertyAutoConfigUrl(value string) (err error) {
	return instance.SetProperty("AutoConfigUrl", value)
}

// GetAutoConfigUrl gets the value of AutoConfigUrl for the instance
func (instance *MDM_VPNv2_Proxy02) GetPropertyAutoConfigUrl() (value string, err error) {
	retValue, err := instance.GetProperty("AutoConfigUrl")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_VPNv2_Proxy02) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_VPNv2_Proxy02) GetPropertyInstanceID() (value string, err error) {
	retValue, err := instance.GetProperty("InstanceID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetParentID sets the value of ParentID for the instance
func (instance *MDM_VPNv2_Proxy02) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_VPNv2_Proxy02) GetPropertyParentID() (value string, err error) {
	retValue, err := instance.GetProperty("ParentID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
