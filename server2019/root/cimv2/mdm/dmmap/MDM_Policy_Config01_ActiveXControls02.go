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

// MDM_Policy_Config01_ActiveXControls02 struct
type MDM_Policy_Config01_ActiveXControls02 struct {
	cim.WmiInstance

	//
	ApprovedInstallationSites string

	//
	InstanceID string

	//
	ParentID string
}

// SetApprovedInstallationSites sets the value of ApprovedInstallationSites for the instance
func (instance *MDM_Policy_Config01_ActiveXControls02) SetPropertyApprovedInstallationSites(value string) (err error) {
	return instance.SetProperty("ApprovedInstallationSites", value)
}

// GetApprovedInstallationSites gets the value of ApprovedInstallationSites for the instance
func (instance *MDM_Policy_Config01_ActiveXControls02) GetPropertyApprovedInstallationSites() (value string, err error) {
	retValue, err := instance.GetProperty("ApprovedInstallationSites")
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
func (instance *MDM_Policy_Config01_ActiveXControls02) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Policy_Config01_ActiveXControls02) GetPropertyInstanceID() (value string, err error) {
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
func (instance *MDM_Policy_Config01_ActiveXControls02) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Policy_Config01_ActiveXControls02) GetPropertyParentID() (value string, err error) {
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
