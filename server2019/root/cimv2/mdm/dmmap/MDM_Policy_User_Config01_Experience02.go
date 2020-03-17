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

// MDM_Policy_User_Config01_Experience02 struct
type MDM_Policy_User_Config01_Experience02 struct {
	cim.WmiInstance

	//
	AllowTailoredExperiencesWithDiagnosticData int32

	//
	AllowThirdPartySuggestionsInWindowsSpotlight int32

	//
	AllowWindowsSpotlight int32

	//
	AllowWindowsSpotlightOnActionCenter int32

	//
	AllowWindowsSpotlightOnSettings int32

	//
	AllowWindowsSpotlightWindowsWelcomeExperience int32

	//
	ConfigureWindowsSpotlightOnLockScreen int32

	//
	InstanceID string

	//
	ParentID string
}

// SetAllowTailoredExperiencesWithDiagnosticData sets the value of AllowTailoredExperiencesWithDiagnosticData for the instance
func (instance *MDM_Policy_User_Config01_Experience02) SetPropertyAllowTailoredExperiencesWithDiagnosticData(value int32) (err error) {
	return instance.SetProperty("AllowTailoredExperiencesWithDiagnosticData", value)
}

// GetAllowTailoredExperiencesWithDiagnosticData gets the value of AllowTailoredExperiencesWithDiagnosticData for the instance
func (instance *MDM_Policy_User_Config01_Experience02) GetPropertyAllowTailoredExperiencesWithDiagnosticData() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowTailoredExperiencesWithDiagnosticData")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAllowThirdPartySuggestionsInWindowsSpotlight sets the value of AllowThirdPartySuggestionsInWindowsSpotlight for the instance
func (instance *MDM_Policy_User_Config01_Experience02) SetPropertyAllowThirdPartySuggestionsInWindowsSpotlight(value int32) (err error) {
	return instance.SetProperty("AllowThirdPartySuggestionsInWindowsSpotlight", value)
}

// GetAllowThirdPartySuggestionsInWindowsSpotlight gets the value of AllowThirdPartySuggestionsInWindowsSpotlight for the instance
func (instance *MDM_Policy_User_Config01_Experience02) GetPropertyAllowThirdPartySuggestionsInWindowsSpotlight() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowThirdPartySuggestionsInWindowsSpotlight")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAllowWindowsSpotlight sets the value of AllowWindowsSpotlight for the instance
func (instance *MDM_Policy_User_Config01_Experience02) SetPropertyAllowWindowsSpotlight(value int32) (err error) {
	return instance.SetProperty("AllowWindowsSpotlight", value)
}

// GetAllowWindowsSpotlight gets the value of AllowWindowsSpotlight for the instance
func (instance *MDM_Policy_User_Config01_Experience02) GetPropertyAllowWindowsSpotlight() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowWindowsSpotlight")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAllowWindowsSpotlightOnActionCenter sets the value of AllowWindowsSpotlightOnActionCenter for the instance
func (instance *MDM_Policy_User_Config01_Experience02) SetPropertyAllowWindowsSpotlightOnActionCenter(value int32) (err error) {
	return instance.SetProperty("AllowWindowsSpotlightOnActionCenter", value)
}

// GetAllowWindowsSpotlightOnActionCenter gets the value of AllowWindowsSpotlightOnActionCenter for the instance
func (instance *MDM_Policy_User_Config01_Experience02) GetPropertyAllowWindowsSpotlightOnActionCenter() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowWindowsSpotlightOnActionCenter")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAllowWindowsSpotlightOnSettings sets the value of AllowWindowsSpotlightOnSettings for the instance
func (instance *MDM_Policy_User_Config01_Experience02) SetPropertyAllowWindowsSpotlightOnSettings(value int32) (err error) {
	return instance.SetProperty("AllowWindowsSpotlightOnSettings", value)
}

// GetAllowWindowsSpotlightOnSettings gets the value of AllowWindowsSpotlightOnSettings for the instance
func (instance *MDM_Policy_User_Config01_Experience02) GetPropertyAllowWindowsSpotlightOnSettings() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowWindowsSpotlightOnSettings")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAllowWindowsSpotlightWindowsWelcomeExperience sets the value of AllowWindowsSpotlightWindowsWelcomeExperience for the instance
func (instance *MDM_Policy_User_Config01_Experience02) SetPropertyAllowWindowsSpotlightWindowsWelcomeExperience(value int32) (err error) {
	return instance.SetProperty("AllowWindowsSpotlightWindowsWelcomeExperience", value)
}

// GetAllowWindowsSpotlightWindowsWelcomeExperience gets the value of AllowWindowsSpotlightWindowsWelcomeExperience for the instance
func (instance *MDM_Policy_User_Config01_Experience02) GetPropertyAllowWindowsSpotlightWindowsWelcomeExperience() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowWindowsSpotlightWindowsWelcomeExperience")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetConfigureWindowsSpotlightOnLockScreen sets the value of ConfigureWindowsSpotlightOnLockScreen for the instance
func (instance *MDM_Policy_User_Config01_Experience02) SetPropertyConfigureWindowsSpotlightOnLockScreen(value int32) (err error) {
	return instance.SetProperty("ConfigureWindowsSpotlightOnLockScreen", value)
}

// GetConfigureWindowsSpotlightOnLockScreen gets the value of ConfigureWindowsSpotlightOnLockScreen for the instance
func (instance *MDM_Policy_User_Config01_Experience02) GetPropertyConfigureWindowsSpotlightOnLockScreen() (value int32, err error) {
	retValue, err := instance.GetProperty("ConfigureWindowsSpotlightOnLockScreen")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_Policy_User_Config01_Experience02) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Policy_User_Config01_Experience02) GetPropertyInstanceID() (value string, err error) {
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
func (instance *MDM_Policy_User_Config01_Experience02) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Policy_User_Config01_Experience02) GetPropertyParentID() (value string, err error) {
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
