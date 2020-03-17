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

// MDM_Policy_Result01_FactoryComposer02 struct
type MDM_Policy_Result01_FactoryComposer02 struct {
	cim.WmiInstance

	//
	BackgroundImagePath string

	//
	InstanceID string

	//
	OEMVersion string

	//
	ParentID string

	//
	UserToSignIn string

	//
	UWPLaunchOnBoot string
}

// SetBackgroundImagePath sets the value of BackgroundImagePath for the instance
func (instance *MDM_Policy_Result01_FactoryComposer02) SetPropertyBackgroundImagePath(value string) (err error) {
	return instance.SetProperty("BackgroundImagePath", value)
}

// GetBackgroundImagePath gets the value of BackgroundImagePath for the instance
func (instance *MDM_Policy_Result01_FactoryComposer02) GetPropertyBackgroundImagePath() (value string, err error) {
	retValue, err := instance.GetProperty("BackgroundImagePath")
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
func (instance *MDM_Policy_Result01_FactoryComposer02) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Policy_Result01_FactoryComposer02) GetPropertyInstanceID() (value string, err error) {
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

// SetOEMVersion sets the value of OEMVersion for the instance
func (instance *MDM_Policy_Result01_FactoryComposer02) SetPropertyOEMVersion(value string) (err error) {
	return instance.SetProperty("OEMVersion", value)
}

// GetOEMVersion gets the value of OEMVersion for the instance
func (instance *MDM_Policy_Result01_FactoryComposer02) GetPropertyOEMVersion() (value string, err error) {
	retValue, err := instance.GetProperty("OEMVersion")
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
func (instance *MDM_Policy_Result01_FactoryComposer02) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Policy_Result01_FactoryComposer02) GetPropertyParentID() (value string, err error) {
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

// SetUserToSignIn sets the value of UserToSignIn for the instance
func (instance *MDM_Policy_Result01_FactoryComposer02) SetPropertyUserToSignIn(value string) (err error) {
	return instance.SetProperty("UserToSignIn", value)
}

// GetUserToSignIn gets the value of UserToSignIn for the instance
func (instance *MDM_Policy_Result01_FactoryComposer02) GetPropertyUserToSignIn() (value string, err error) {
	retValue, err := instance.GetProperty("UserToSignIn")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUWPLaunchOnBoot sets the value of UWPLaunchOnBoot for the instance
func (instance *MDM_Policy_Result01_FactoryComposer02) SetPropertyUWPLaunchOnBoot(value string) (err error) {
	return instance.SetProperty("UWPLaunchOnBoot", value)
}

// GetUWPLaunchOnBoot gets the value of UWPLaunchOnBoot for the instance
func (instance *MDM_Policy_Result01_FactoryComposer02) GetPropertyUWPLaunchOnBoot() (value string, err error) {
	retValue, err := instance.GetProperty("UWPLaunchOnBoot")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
