// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2.mdm.dmmap
//////////////////////////////////////////////
package dmmap

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MDM_VPNv2_NativeProfile02 struct
type MDM_VPNv2_NativeProfile02 struct {
	cim.WmiInstance

	//
	DisableClassBasedDefaultRoute bool

	//
	InstanceID string

	//
	L2tpPsk string

	//
	NativeProtocolType string

	//
	ParentID string

	//
	RoutingPolicyType string

	//
	Servers string
}

// SetDisableClassBasedDefaultRoute sets the value of DisableClassBasedDefaultRoute for the instance
func (instance *MDM_VPNv2_NativeProfile02) SetPropertyDisableClassBasedDefaultRoute(value bool) (err error) {
	return instance.SetProperty("DisableClassBasedDefaultRoute", value)
}

// GetDisableClassBasedDefaultRoute gets the value of DisableClassBasedDefaultRoute for the instance
func (instance *MDM_VPNv2_NativeProfile02) GetPropertyDisableClassBasedDefaultRoute() (value bool, err error) {
	retValue, err := instance.GetProperty("DisableClassBasedDefaultRoute")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_VPNv2_NativeProfile02) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_VPNv2_NativeProfile02) GetPropertyInstanceID() (value string, err error) {
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

// SetL2tpPsk sets the value of L2tpPsk for the instance
func (instance *MDM_VPNv2_NativeProfile02) SetPropertyL2tpPsk(value string) (err error) {
	return instance.SetProperty("L2tpPsk", value)
}

// GetL2tpPsk gets the value of L2tpPsk for the instance
func (instance *MDM_VPNv2_NativeProfile02) GetPropertyL2tpPsk() (value string, err error) {
	retValue, err := instance.GetProperty("L2tpPsk")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNativeProtocolType sets the value of NativeProtocolType for the instance
func (instance *MDM_VPNv2_NativeProfile02) SetPropertyNativeProtocolType(value string) (err error) {
	return instance.SetProperty("NativeProtocolType", value)
}

// GetNativeProtocolType gets the value of NativeProtocolType for the instance
func (instance *MDM_VPNv2_NativeProfile02) GetPropertyNativeProtocolType() (value string, err error) {
	retValue, err := instance.GetProperty("NativeProtocolType")
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
func (instance *MDM_VPNv2_NativeProfile02) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_VPNv2_NativeProfile02) GetPropertyParentID() (value string, err error) {
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

// SetRoutingPolicyType sets the value of RoutingPolicyType for the instance
func (instance *MDM_VPNv2_NativeProfile02) SetPropertyRoutingPolicyType(value string) (err error) {
	return instance.SetProperty("RoutingPolicyType", value)
}

// GetRoutingPolicyType gets the value of RoutingPolicyType for the instance
func (instance *MDM_VPNv2_NativeProfile02) GetPropertyRoutingPolicyType() (value string, err error) {
	retValue, err := instance.GetProperty("RoutingPolicyType")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetServers sets the value of Servers for the instance
func (instance *MDM_VPNv2_NativeProfile02) SetPropertyServers(value string) (err error) {
	return instance.SetProperty("Servers", value)
}

// GetServers gets the value of Servers for the instance
func (instance *MDM_VPNv2_NativeProfile02) GetPropertyServers() (value string, err error) {
	retValue, err := instance.GetProperty("Servers")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
