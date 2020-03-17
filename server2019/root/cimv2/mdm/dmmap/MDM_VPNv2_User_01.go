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

// MDM_VPNv2_User_01 struct
type MDM_VPNv2_User_01 struct {
	cim.WmiInstance

	//
	AlwaysOn bool

	//
	ByPassForLocal bool

	//
	DnsSuffix string

	//
	EdpModeId string

	//
	InstanceID string

	//
	ParentID string

	//
	ProfileXML string

	//
	RememberCredentials bool

	//
	TrustedNetworkDetection string
}

// SetAlwaysOn sets the value of AlwaysOn for the instance
func (instance *MDM_VPNv2_User_01) SetPropertyAlwaysOn(value bool) (err error) {
	return instance.SetProperty("AlwaysOn", value)
}

// GetAlwaysOn gets the value of AlwaysOn for the instance
func (instance *MDM_VPNv2_User_01) GetPropertyAlwaysOn() (value bool, err error) {
	retValue, err := instance.GetProperty("AlwaysOn")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetByPassForLocal sets the value of ByPassForLocal for the instance
func (instance *MDM_VPNv2_User_01) SetPropertyByPassForLocal(value bool) (err error) {
	return instance.SetProperty("ByPassForLocal", value)
}

// GetByPassForLocal gets the value of ByPassForLocal for the instance
func (instance *MDM_VPNv2_User_01) GetPropertyByPassForLocal() (value bool, err error) {
	retValue, err := instance.GetProperty("ByPassForLocal")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDnsSuffix sets the value of DnsSuffix for the instance
func (instance *MDM_VPNv2_User_01) SetPropertyDnsSuffix(value string) (err error) {
	return instance.SetProperty("DnsSuffix", value)
}

// GetDnsSuffix gets the value of DnsSuffix for the instance
func (instance *MDM_VPNv2_User_01) GetPropertyDnsSuffix() (value string, err error) {
	retValue, err := instance.GetProperty("DnsSuffix")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEdpModeId sets the value of EdpModeId for the instance
func (instance *MDM_VPNv2_User_01) SetPropertyEdpModeId(value string) (err error) {
	return instance.SetProperty("EdpModeId", value)
}

// GetEdpModeId gets the value of EdpModeId for the instance
func (instance *MDM_VPNv2_User_01) GetPropertyEdpModeId() (value string, err error) {
	retValue, err := instance.GetProperty("EdpModeId")
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
func (instance *MDM_VPNv2_User_01) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_VPNv2_User_01) GetPropertyInstanceID() (value string, err error) {
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
func (instance *MDM_VPNv2_User_01) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_VPNv2_User_01) GetPropertyParentID() (value string, err error) {
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

// SetProfileXML sets the value of ProfileXML for the instance
func (instance *MDM_VPNv2_User_01) SetPropertyProfileXML(value string) (err error) {
	return instance.SetProperty("ProfileXML", value)
}

// GetProfileXML gets the value of ProfileXML for the instance
func (instance *MDM_VPNv2_User_01) GetPropertyProfileXML() (value string, err error) {
	retValue, err := instance.GetProperty("ProfileXML")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRememberCredentials sets the value of RememberCredentials for the instance
func (instance *MDM_VPNv2_User_01) SetPropertyRememberCredentials(value bool) (err error) {
	return instance.SetProperty("RememberCredentials", value)
}

// GetRememberCredentials gets the value of RememberCredentials for the instance
func (instance *MDM_VPNv2_User_01) GetPropertyRememberCredentials() (value bool, err error) {
	retValue, err := instance.GetProperty("RememberCredentials")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTrustedNetworkDetection sets the value of TrustedNetworkDetection for the instance
func (instance *MDM_VPNv2_User_01) SetPropertyTrustedNetworkDetection(value string) (err error) {
	return instance.SetProperty("TrustedNetworkDetection", value)
}

// GetTrustedNetworkDetection gets the value of TrustedNetworkDetection for the instance
func (instance *MDM_VPNv2_User_01) GetPropertyTrustedNetworkDetection() (value string, err error) {
	retValue, err := instance.GetProperty("TrustedNetworkDetection")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
