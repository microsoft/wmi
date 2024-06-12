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

// MDM_Policy_Config01_KioskBrowser02 struct
type MDM_Policy_Config01_KioskBrowser02 struct {
	*cim.WmiInstance

	//
	BlockedUrlExceptions string

	//
	BlockedUrls string

	//
	DefaultURL string

	//
	EnableEndSessionButton int32

	//
	EnableHomeButton int32

	//
	EnableNavigationButtons int32

	//
	InstanceID string

	//
	ParentID string

	//
	RestartOnIdleTime int32
}

func NewMDM_Policy_Config01_KioskBrowser02Ex1(instance *cim.WmiInstance) (newInstance *MDM_Policy_Config01_KioskBrowser02, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_Policy_Config01_KioskBrowser02{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_Policy_Config01_KioskBrowser02Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_Policy_Config01_KioskBrowser02, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_Policy_Config01_KioskBrowser02{
		WmiInstance: tmp,
	}
	return
}

// SetBlockedUrlExceptions sets the value of BlockedUrlExceptions for the instance
func (instance *MDM_Policy_Config01_KioskBrowser02) SetPropertyBlockedUrlExceptions(value string) (err error) {
	return instance.SetProperty("BlockedUrlExceptions", (value))
}

// GetBlockedUrlExceptions gets the value of BlockedUrlExceptions for the instance
func (instance *MDM_Policy_Config01_KioskBrowser02) GetPropertyBlockedUrlExceptions() (value string, err error) {
	retValue, err := instance.GetProperty("BlockedUrlExceptions")
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

// SetBlockedUrls sets the value of BlockedUrls for the instance
func (instance *MDM_Policy_Config01_KioskBrowser02) SetPropertyBlockedUrls(value string) (err error) {
	return instance.SetProperty("BlockedUrls", (value))
}

// GetBlockedUrls gets the value of BlockedUrls for the instance
func (instance *MDM_Policy_Config01_KioskBrowser02) GetPropertyBlockedUrls() (value string, err error) {
	retValue, err := instance.GetProperty("BlockedUrls")
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

// SetDefaultURL sets the value of DefaultURL for the instance
func (instance *MDM_Policy_Config01_KioskBrowser02) SetPropertyDefaultURL(value string) (err error) {
	return instance.SetProperty("DefaultURL", (value))
}

// GetDefaultURL gets the value of DefaultURL for the instance
func (instance *MDM_Policy_Config01_KioskBrowser02) GetPropertyDefaultURL() (value string, err error) {
	retValue, err := instance.GetProperty("DefaultURL")
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

// SetEnableEndSessionButton sets the value of EnableEndSessionButton for the instance
func (instance *MDM_Policy_Config01_KioskBrowser02) SetPropertyEnableEndSessionButton(value int32) (err error) {
	return instance.SetProperty("EnableEndSessionButton", (value))
}

// GetEnableEndSessionButton gets the value of EnableEndSessionButton for the instance
func (instance *MDM_Policy_Config01_KioskBrowser02) GetPropertyEnableEndSessionButton() (value int32, err error) {
	retValue, err := instance.GetProperty("EnableEndSessionButton")
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

// SetEnableHomeButton sets the value of EnableHomeButton for the instance
func (instance *MDM_Policy_Config01_KioskBrowser02) SetPropertyEnableHomeButton(value int32) (err error) {
	return instance.SetProperty("EnableHomeButton", (value))
}

// GetEnableHomeButton gets the value of EnableHomeButton for the instance
func (instance *MDM_Policy_Config01_KioskBrowser02) GetPropertyEnableHomeButton() (value int32, err error) {
	retValue, err := instance.GetProperty("EnableHomeButton")
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

// SetEnableNavigationButtons sets the value of EnableNavigationButtons for the instance
func (instance *MDM_Policy_Config01_KioskBrowser02) SetPropertyEnableNavigationButtons(value int32) (err error) {
	return instance.SetProperty("EnableNavigationButtons", (value))
}

// GetEnableNavigationButtons gets the value of EnableNavigationButtons for the instance
func (instance *MDM_Policy_Config01_KioskBrowser02) GetPropertyEnableNavigationButtons() (value int32, err error) {
	retValue, err := instance.GetProperty("EnableNavigationButtons")
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
func (instance *MDM_Policy_Config01_KioskBrowser02) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", (value))
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Policy_Config01_KioskBrowser02) GetPropertyInstanceID() (value string, err error) {
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
func (instance *MDM_Policy_Config01_KioskBrowser02) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", (value))
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Policy_Config01_KioskBrowser02) GetPropertyParentID() (value string, err error) {
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

// SetRestartOnIdleTime sets the value of RestartOnIdleTime for the instance
func (instance *MDM_Policy_Config01_KioskBrowser02) SetPropertyRestartOnIdleTime(value int32) (err error) {
	return instance.SetProperty("RestartOnIdleTime", (value))
}

// GetRestartOnIdleTime gets the value of RestartOnIdleTime for the instance
func (instance *MDM_Policy_Config01_KioskBrowser02) GetPropertyRestartOnIdleTime() (value int32, err error) {
	retValue, err := instance.GetProperty("RestartOnIdleTime")
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
