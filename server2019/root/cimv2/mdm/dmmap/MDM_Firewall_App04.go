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

// MDM_Firewall_App04 struct
type MDM_Firewall_App04 struct {
	cim.WmiInstance

	//
	FilePath string

	//
	Fqbn string

	//
	InstanceID string

	//
	PackageFamilyName string

	//
	ParentID string

	//
	ServiceName string
}

// SetFilePath sets the value of FilePath for the instance
func (instance *MDM_Firewall_App04) SetPropertyFilePath(value string) (err error) {
	return instance.SetProperty("FilePath", value)
}

// GetFilePath gets the value of FilePath for the instance
func (instance *MDM_Firewall_App04) GetPropertyFilePath() (value string, err error) {
	retValue, err := instance.GetProperty("FilePath")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFqbn sets the value of Fqbn for the instance
func (instance *MDM_Firewall_App04) SetPropertyFqbn(value string) (err error) {
	return instance.SetProperty("Fqbn", value)
}

// GetFqbn gets the value of Fqbn for the instance
func (instance *MDM_Firewall_App04) GetPropertyFqbn() (value string, err error) {
	retValue, err := instance.GetProperty("Fqbn")
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
func (instance *MDM_Firewall_App04) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Firewall_App04) GetPropertyInstanceID() (value string, err error) {
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

// SetPackageFamilyName sets the value of PackageFamilyName for the instance
func (instance *MDM_Firewall_App04) SetPropertyPackageFamilyName(value string) (err error) {
	return instance.SetProperty("PackageFamilyName", value)
}

// GetPackageFamilyName gets the value of PackageFamilyName for the instance
func (instance *MDM_Firewall_App04) GetPropertyPackageFamilyName() (value string, err error) {
	retValue, err := instance.GetProperty("PackageFamilyName")
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
func (instance *MDM_Firewall_App04) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Firewall_App04) GetPropertyParentID() (value string, err error) {
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

// SetServiceName sets the value of ServiceName for the instance
func (instance *MDM_Firewall_App04) SetPropertyServiceName(value string) (err error) {
	return instance.SetProperty("ServiceName", value)
}

// GetServiceName gets the value of ServiceName for the instance
func (instance *MDM_Firewall_App04) GetPropertyServiceName() (value string, err error) {
	retValue, err := instance.GetProperty("ServiceName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
