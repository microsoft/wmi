// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.CIMV2.mdm.dmmap
//////////////////////////////////////////////
package dmmap

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MDM_Policy_User_Result01_Experience02 struct
type MDM_Policy_User_Result01_Experience02 struct {
	*cim.WmiInstance

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

func NewMDM_Policy_User_Result01_Experience02Ex1(instance *cim.WmiInstance) (newInstance *MDM_Policy_User_Result01_Experience02, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_Policy_User_Result01_Experience02{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_Policy_User_Result01_Experience02Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_Policy_User_Result01_Experience02, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_Policy_User_Result01_Experience02{
		WmiInstance: tmp,
	}
	return
}

// SetAllowTailoredExperiencesWithDiagnosticData sets the value of AllowTailoredExperiencesWithDiagnosticData for the instance
func (instance *MDM_Policy_User_Result01_Experience02) SetPropertyAllowTailoredExperiencesWithDiagnosticData(value int32) (err error) {
	return instance.SetProperty("AllowTailoredExperiencesWithDiagnosticData", (value))
}

// GetAllowTailoredExperiencesWithDiagnosticData gets the value of AllowTailoredExperiencesWithDiagnosticData for the instance
func (instance *MDM_Policy_User_Result01_Experience02) GetPropertyAllowTailoredExperiencesWithDiagnosticData() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowTailoredExperiencesWithDiagnosticData")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int32(valuetmp)

	return
}

// SetAllowThirdPartySuggestionsInWindowsSpotlight sets the value of AllowThirdPartySuggestionsInWindowsSpotlight for the instance
func (instance *MDM_Policy_User_Result01_Experience02) SetPropertyAllowThirdPartySuggestionsInWindowsSpotlight(value int32) (err error) {
	return instance.SetProperty("AllowThirdPartySuggestionsInWindowsSpotlight", (value))
}

// GetAllowThirdPartySuggestionsInWindowsSpotlight gets the value of AllowThirdPartySuggestionsInWindowsSpotlight for the instance
func (instance *MDM_Policy_User_Result01_Experience02) GetPropertyAllowThirdPartySuggestionsInWindowsSpotlight() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowThirdPartySuggestionsInWindowsSpotlight")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int32(valuetmp)

	return
}

// SetAllowWindowsSpotlight sets the value of AllowWindowsSpotlight for the instance
func (instance *MDM_Policy_User_Result01_Experience02) SetPropertyAllowWindowsSpotlight(value int32) (err error) {
	return instance.SetProperty("AllowWindowsSpotlight", (value))
}

// GetAllowWindowsSpotlight gets the value of AllowWindowsSpotlight for the instance
func (instance *MDM_Policy_User_Result01_Experience02) GetPropertyAllowWindowsSpotlight() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowWindowsSpotlight")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int32(valuetmp)

	return
}

// SetAllowWindowsSpotlightOnActionCenter sets the value of AllowWindowsSpotlightOnActionCenter for the instance
func (instance *MDM_Policy_User_Result01_Experience02) SetPropertyAllowWindowsSpotlightOnActionCenter(value int32) (err error) {
	return instance.SetProperty("AllowWindowsSpotlightOnActionCenter", (value))
}

// GetAllowWindowsSpotlightOnActionCenter gets the value of AllowWindowsSpotlightOnActionCenter for the instance
func (instance *MDM_Policy_User_Result01_Experience02) GetPropertyAllowWindowsSpotlightOnActionCenter() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowWindowsSpotlightOnActionCenter")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int32(valuetmp)

	return
}

// SetAllowWindowsSpotlightOnSettings sets the value of AllowWindowsSpotlightOnSettings for the instance
func (instance *MDM_Policy_User_Result01_Experience02) SetPropertyAllowWindowsSpotlightOnSettings(value int32) (err error) {
	return instance.SetProperty("AllowWindowsSpotlightOnSettings", (value))
}

// GetAllowWindowsSpotlightOnSettings gets the value of AllowWindowsSpotlightOnSettings for the instance
func (instance *MDM_Policy_User_Result01_Experience02) GetPropertyAllowWindowsSpotlightOnSettings() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowWindowsSpotlightOnSettings")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int32(valuetmp)

	return
}

// SetAllowWindowsSpotlightWindowsWelcomeExperience sets the value of AllowWindowsSpotlightWindowsWelcomeExperience for the instance
func (instance *MDM_Policy_User_Result01_Experience02) SetPropertyAllowWindowsSpotlightWindowsWelcomeExperience(value int32) (err error) {
	return instance.SetProperty("AllowWindowsSpotlightWindowsWelcomeExperience", (value))
}

// GetAllowWindowsSpotlightWindowsWelcomeExperience gets the value of AllowWindowsSpotlightWindowsWelcomeExperience for the instance
func (instance *MDM_Policy_User_Result01_Experience02) GetPropertyAllowWindowsSpotlightWindowsWelcomeExperience() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowWindowsSpotlightWindowsWelcomeExperience")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int32(valuetmp)

	return
}

// SetConfigureWindowsSpotlightOnLockScreen sets the value of ConfigureWindowsSpotlightOnLockScreen for the instance
func (instance *MDM_Policy_User_Result01_Experience02) SetPropertyConfigureWindowsSpotlightOnLockScreen(value int32) (err error) {
	return instance.SetProperty("ConfigureWindowsSpotlightOnLockScreen", (value))
}

// GetConfigureWindowsSpotlightOnLockScreen gets the value of ConfigureWindowsSpotlightOnLockScreen for the instance
func (instance *MDM_Policy_User_Result01_Experience02) GetPropertyConfigureWindowsSpotlightOnLockScreen() (value int32, err error) {
	retValue, err := instance.GetProperty("ConfigureWindowsSpotlightOnLockScreen")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int32(valuetmp)

	return
}

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_Policy_User_Result01_Experience02) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", (value))
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Policy_User_Result01_Experience02) GetPropertyInstanceID() (value string, err error) {
	retValue, err := instance.GetProperty("InstanceID")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetParentID sets the value of ParentID for the instance
func (instance *MDM_Policy_User_Result01_Experience02) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", (value))
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Policy_User_Result01_Experience02) GetPropertyParentID() (value string, err error) {
	retValue, err := instance.GetProperty("ParentID")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}
