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

// MDM_SecureAssessment struct
type MDM_SecureAssessment struct {
	cim.WmiInstance

	//
	AllowScreenMonitoring bool

	//
	AllowTextSuggestions bool

	//
	InstanceID string

	//
	LaunchURI string

	//
	ParentID string

	//
	RequirePrinting bool

	//
	TesterAccount string
}

// SetAllowScreenMonitoring sets the value of AllowScreenMonitoring for the instance
func (instance *MDM_SecureAssessment) SetPropertyAllowScreenMonitoring(value bool) (err error) {
	return instance.SetProperty("AllowScreenMonitoring", value)
}

// GetAllowScreenMonitoring gets the value of AllowScreenMonitoring for the instance
func (instance *MDM_SecureAssessment) GetPropertyAllowScreenMonitoring() (value bool, err error) {
	retValue, err := instance.GetProperty("AllowScreenMonitoring")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAllowTextSuggestions sets the value of AllowTextSuggestions for the instance
func (instance *MDM_SecureAssessment) SetPropertyAllowTextSuggestions(value bool) (err error) {
	return instance.SetProperty("AllowTextSuggestions", value)
}

// GetAllowTextSuggestions gets the value of AllowTextSuggestions for the instance
func (instance *MDM_SecureAssessment) GetPropertyAllowTextSuggestions() (value bool, err error) {
	retValue, err := instance.GetProperty("AllowTextSuggestions")
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
func (instance *MDM_SecureAssessment) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_SecureAssessment) GetPropertyInstanceID() (value string, err error) {
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

// SetLaunchURI sets the value of LaunchURI for the instance
func (instance *MDM_SecureAssessment) SetPropertyLaunchURI(value string) (err error) {
	return instance.SetProperty("LaunchURI", value)
}

// GetLaunchURI gets the value of LaunchURI for the instance
func (instance *MDM_SecureAssessment) GetPropertyLaunchURI() (value string, err error) {
	retValue, err := instance.GetProperty("LaunchURI")
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
func (instance *MDM_SecureAssessment) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_SecureAssessment) GetPropertyParentID() (value string, err error) {
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

// SetRequirePrinting sets the value of RequirePrinting for the instance
func (instance *MDM_SecureAssessment) SetPropertyRequirePrinting(value bool) (err error) {
	return instance.SetProperty("RequirePrinting", value)
}

// GetRequirePrinting gets the value of RequirePrinting for the instance
func (instance *MDM_SecureAssessment) GetPropertyRequirePrinting() (value bool, err error) {
	retValue, err := instance.GetProperty("RequirePrinting")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTesterAccount sets the value of TesterAccount for the instance
func (instance *MDM_SecureAssessment) SetPropertyTesterAccount(value string) (err error) {
	return instance.SetProperty("TesterAccount", value)
}

// GetTesterAccount gets the value of TesterAccount for the instance
func (instance *MDM_SecureAssessment) GetPropertyTesterAccount() (value string, err error) {
	retValue, err := instance.GetProperty("TesterAccount")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
