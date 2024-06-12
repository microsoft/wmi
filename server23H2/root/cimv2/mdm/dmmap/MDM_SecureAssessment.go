// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
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

// MDM_SecureAssessment struct
type MDM_SecureAssessment struct {
	*cim.WmiInstance

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

func NewMDM_SecureAssessmentEx1(instance *cim.WmiInstance) (newInstance *MDM_SecureAssessment, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_SecureAssessment{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_SecureAssessmentEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_SecureAssessment, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_SecureAssessment{
		WmiInstance: tmp,
	}
	return
}

// SetAllowScreenMonitoring sets the value of AllowScreenMonitoring for the instance
func (instance *MDM_SecureAssessment) SetPropertyAllowScreenMonitoring(value bool) (err error) {
	return instance.SetProperty("AllowScreenMonitoring", (value))
}

// GetAllowScreenMonitoring gets the value of AllowScreenMonitoring for the instance
func (instance *MDM_SecureAssessment) GetPropertyAllowScreenMonitoring() (value bool, err error) {
	retValue, err := instance.GetProperty("AllowScreenMonitoring")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetAllowTextSuggestions sets the value of AllowTextSuggestions for the instance
func (instance *MDM_SecureAssessment) SetPropertyAllowTextSuggestions(value bool) (err error) {
	return instance.SetProperty("AllowTextSuggestions", (value))
}

// GetAllowTextSuggestions gets the value of AllowTextSuggestions for the instance
func (instance *MDM_SecureAssessment) GetPropertyAllowTextSuggestions() (value bool, err error) {
	retValue, err := instance.GetProperty("AllowTextSuggestions")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_SecureAssessment) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", (value))
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_SecureAssessment) GetPropertyInstanceID() (value string, err error) {
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

// SetLaunchURI sets the value of LaunchURI for the instance
func (instance *MDM_SecureAssessment) SetPropertyLaunchURI(value string) (err error) {
	return instance.SetProperty("LaunchURI", (value))
}

// GetLaunchURI gets the value of LaunchURI for the instance
func (instance *MDM_SecureAssessment) GetPropertyLaunchURI() (value string, err error) {
	retValue, err := instance.GetProperty("LaunchURI")
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
func (instance *MDM_SecureAssessment) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", (value))
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_SecureAssessment) GetPropertyParentID() (value string, err error) {
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

// SetRequirePrinting sets the value of RequirePrinting for the instance
func (instance *MDM_SecureAssessment) SetPropertyRequirePrinting(value bool) (err error) {
	return instance.SetProperty("RequirePrinting", (value))
}

// GetRequirePrinting gets the value of RequirePrinting for the instance
func (instance *MDM_SecureAssessment) GetPropertyRequirePrinting() (value bool, err error) {
	retValue, err := instance.GetProperty("RequirePrinting")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetTesterAccount sets the value of TesterAccount for the instance
func (instance *MDM_SecureAssessment) SetPropertyTesterAccount(value string) (err error) {
	return instance.SetProperty("TesterAccount", (value))
}

// GetTesterAccount gets the value of TesterAccount for the instance
func (instance *MDM_SecureAssessment) GetPropertyTesterAccount() (value string, err error) {
	retValue, err := instance.GetProperty("TesterAccount")
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
