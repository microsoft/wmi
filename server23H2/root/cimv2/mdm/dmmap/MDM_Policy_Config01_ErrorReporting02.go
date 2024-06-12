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

// MDM_Policy_Config01_ErrorReporting02 struct
type MDM_Policy_Config01_ErrorReporting02 struct {
	*cim.WmiInstance

	//
	CustomizeConsentSettings string

	//
	DisableWindowsErrorReporting string

	//
	DisplayErrorNotification string

	//
	DoNotSendAdditionalData string

	//
	InstanceID string

	//
	ParentID string

	//
	PreventCriticalErrorDisplay string
}

func NewMDM_Policy_Config01_ErrorReporting02Ex1(instance *cim.WmiInstance) (newInstance *MDM_Policy_Config01_ErrorReporting02, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_Policy_Config01_ErrorReporting02{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_Policy_Config01_ErrorReporting02Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_Policy_Config01_ErrorReporting02, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_Policy_Config01_ErrorReporting02{
		WmiInstance: tmp,
	}
	return
}

// SetCustomizeConsentSettings sets the value of CustomizeConsentSettings for the instance
func (instance *MDM_Policy_Config01_ErrorReporting02) SetPropertyCustomizeConsentSettings(value string) (err error) {
	return instance.SetProperty("CustomizeConsentSettings", (value))
}

// GetCustomizeConsentSettings gets the value of CustomizeConsentSettings for the instance
func (instance *MDM_Policy_Config01_ErrorReporting02) GetPropertyCustomizeConsentSettings() (value string, err error) {
	retValue, err := instance.GetProperty("CustomizeConsentSettings")
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

// SetDisableWindowsErrorReporting sets the value of DisableWindowsErrorReporting for the instance
func (instance *MDM_Policy_Config01_ErrorReporting02) SetPropertyDisableWindowsErrorReporting(value string) (err error) {
	return instance.SetProperty("DisableWindowsErrorReporting", (value))
}

// GetDisableWindowsErrorReporting gets the value of DisableWindowsErrorReporting for the instance
func (instance *MDM_Policy_Config01_ErrorReporting02) GetPropertyDisableWindowsErrorReporting() (value string, err error) {
	retValue, err := instance.GetProperty("DisableWindowsErrorReporting")
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

// SetDisplayErrorNotification sets the value of DisplayErrorNotification for the instance
func (instance *MDM_Policy_Config01_ErrorReporting02) SetPropertyDisplayErrorNotification(value string) (err error) {
	return instance.SetProperty("DisplayErrorNotification", (value))
}

// GetDisplayErrorNotification gets the value of DisplayErrorNotification for the instance
func (instance *MDM_Policy_Config01_ErrorReporting02) GetPropertyDisplayErrorNotification() (value string, err error) {
	retValue, err := instance.GetProperty("DisplayErrorNotification")
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

// SetDoNotSendAdditionalData sets the value of DoNotSendAdditionalData for the instance
func (instance *MDM_Policy_Config01_ErrorReporting02) SetPropertyDoNotSendAdditionalData(value string) (err error) {
	return instance.SetProperty("DoNotSendAdditionalData", (value))
}

// GetDoNotSendAdditionalData gets the value of DoNotSendAdditionalData for the instance
func (instance *MDM_Policy_Config01_ErrorReporting02) GetPropertyDoNotSendAdditionalData() (value string, err error) {
	retValue, err := instance.GetProperty("DoNotSendAdditionalData")
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

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_Policy_Config01_ErrorReporting02) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", (value))
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Policy_Config01_ErrorReporting02) GetPropertyInstanceID() (value string, err error) {
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
func (instance *MDM_Policy_Config01_ErrorReporting02) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", (value))
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Policy_Config01_ErrorReporting02) GetPropertyParentID() (value string, err error) {
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

// SetPreventCriticalErrorDisplay sets the value of PreventCriticalErrorDisplay for the instance
func (instance *MDM_Policy_Config01_ErrorReporting02) SetPropertyPreventCriticalErrorDisplay(value string) (err error) {
	return instance.SetProperty("PreventCriticalErrorDisplay", (value))
}

// GetPreventCriticalErrorDisplay gets the value of PreventCriticalErrorDisplay for the instance
func (instance *MDM_Policy_Config01_ErrorReporting02) GetPropertyPreventCriticalErrorDisplay() (value string, err error) {
	retValue, err := instance.GetProperty("PreventCriticalErrorDisplay")
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
