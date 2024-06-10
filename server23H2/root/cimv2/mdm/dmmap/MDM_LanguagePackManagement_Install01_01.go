// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.CIMV2.mdm.dmmap
//
// ////////////////////////////////////////////
package dmmap

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MDM_LanguagePackManagement_Install01_01 struct
type MDM_LanguagePackManagement_Install01_01 struct {
	*cim.WmiInstance

	//
	CopyToDeviceInternationalSettings bool

	//
	EnableLanguageFeatureInstallations bool

	//
	ErrorCode int32

	//
	InstanceID string

	//
	ParentID string

	//
	Status int32
}

func NewMDM_LanguagePackManagement_Install01_01Ex1(instance *cim.WmiInstance) (newInstance *MDM_LanguagePackManagement_Install01_01, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_LanguagePackManagement_Install01_01{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_LanguagePackManagement_Install01_01Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_LanguagePackManagement_Install01_01, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_LanguagePackManagement_Install01_01{
		WmiInstance: tmp,
	}
	return
}

// SetCopyToDeviceInternationalSettings sets the value of CopyToDeviceInternationalSettings for the instance
func (instance *MDM_LanguagePackManagement_Install01_01) SetPropertyCopyToDeviceInternationalSettings(value bool) (err error) {
	return instance.SetProperty("CopyToDeviceInternationalSettings", (value))
}

// GetCopyToDeviceInternationalSettings gets the value of CopyToDeviceInternationalSettings for the instance
func (instance *MDM_LanguagePackManagement_Install01_01) GetPropertyCopyToDeviceInternationalSettings() (value bool, err error) {
	retValue, err := instance.GetProperty("CopyToDeviceInternationalSettings")
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

// SetEnableLanguageFeatureInstallations sets the value of EnableLanguageFeatureInstallations for the instance
func (instance *MDM_LanguagePackManagement_Install01_01) SetPropertyEnableLanguageFeatureInstallations(value bool) (err error) {
	return instance.SetProperty("EnableLanguageFeatureInstallations", (value))
}

// GetEnableLanguageFeatureInstallations gets the value of EnableLanguageFeatureInstallations for the instance
func (instance *MDM_LanguagePackManagement_Install01_01) GetPropertyEnableLanguageFeatureInstallations() (value bool, err error) {
	retValue, err := instance.GetProperty("EnableLanguageFeatureInstallations")
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

// SetErrorCode sets the value of ErrorCode for the instance
func (instance *MDM_LanguagePackManagement_Install01_01) SetPropertyErrorCode(value int32) (err error) {
	return instance.SetProperty("ErrorCode", (value))
}

// GetErrorCode gets the value of ErrorCode for the instance
func (instance *MDM_LanguagePackManagement_Install01_01) GetPropertyErrorCode() (value int32, err error) {
	retValue, err := instance.GetProperty("ErrorCode")
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
func (instance *MDM_LanguagePackManagement_Install01_01) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", (value))
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_LanguagePackManagement_Install01_01) GetPropertyInstanceID() (value string, err error) {
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
func (instance *MDM_LanguagePackManagement_Install01_01) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", (value))
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_LanguagePackManagement_Install01_01) GetPropertyParentID() (value string, err error) {
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

// SetStatus sets the value of Status for the instance
func (instance *MDM_LanguagePackManagement_Install01_01) SetPropertyStatus(value int32) (err error) {
	return instance.SetProperty("Status", (value))
}

// GetStatus gets the value of Status for the instance
func (instance *MDM_LanguagePackManagement_Install01_01) GetPropertyStatus() (value int32, err error) {
	retValue, err := instance.GetProperty("Status")
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

//

// <param name="param" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MDM_LanguagePackManagement_Install01_01) StartInstallationMethod( /* IN */ param string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("StartInstallationMethod", param)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
