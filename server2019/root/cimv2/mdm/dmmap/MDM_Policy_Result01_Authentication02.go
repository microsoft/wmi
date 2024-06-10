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

// MDM_Policy_Result01_Authentication02 struct
type MDM_Policy_Result01_Authentication02 struct {
	*cim.WmiInstance

	//
	AllowAadPasswordReset int32

	//
	AllowFastReconnect int32

	//
	AllowSecondaryAuthenticationDevice int32

	//
	EnableFastFirstSignIn int32

	//
	EnableWebSignIn int32

	//
	InstanceID string

	//
	ParentID string

	//
	PreferredAadTenantDomainName string
}

func NewMDM_Policy_Result01_Authentication02Ex1(instance *cim.WmiInstance) (newInstance *MDM_Policy_Result01_Authentication02, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_Policy_Result01_Authentication02{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_Policy_Result01_Authentication02Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_Policy_Result01_Authentication02, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_Policy_Result01_Authentication02{
		WmiInstance: tmp,
	}
	return
}

// SetAllowAadPasswordReset sets the value of AllowAadPasswordReset for the instance
func (instance *MDM_Policy_Result01_Authentication02) SetPropertyAllowAadPasswordReset(value int32) (err error) {
	return instance.SetProperty("AllowAadPasswordReset", (value))
}

// GetAllowAadPasswordReset gets the value of AllowAadPasswordReset for the instance
func (instance *MDM_Policy_Result01_Authentication02) GetPropertyAllowAadPasswordReset() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowAadPasswordReset")
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

// SetAllowFastReconnect sets the value of AllowFastReconnect for the instance
func (instance *MDM_Policy_Result01_Authentication02) SetPropertyAllowFastReconnect(value int32) (err error) {
	return instance.SetProperty("AllowFastReconnect", (value))
}

// GetAllowFastReconnect gets the value of AllowFastReconnect for the instance
func (instance *MDM_Policy_Result01_Authentication02) GetPropertyAllowFastReconnect() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowFastReconnect")
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

// SetAllowSecondaryAuthenticationDevice sets the value of AllowSecondaryAuthenticationDevice for the instance
func (instance *MDM_Policy_Result01_Authentication02) SetPropertyAllowSecondaryAuthenticationDevice(value int32) (err error) {
	return instance.SetProperty("AllowSecondaryAuthenticationDevice", (value))
}

// GetAllowSecondaryAuthenticationDevice gets the value of AllowSecondaryAuthenticationDevice for the instance
func (instance *MDM_Policy_Result01_Authentication02) GetPropertyAllowSecondaryAuthenticationDevice() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowSecondaryAuthenticationDevice")
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

// SetEnableFastFirstSignIn sets the value of EnableFastFirstSignIn for the instance
func (instance *MDM_Policy_Result01_Authentication02) SetPropertyEnableFastFirstSignIn(value int32) (err error) {
	return instance.SetProperty("EnableFastFirstSignIn", (value))
}

// GetEnableFastFirstSignIn gets the value of EnableFastFirstSignIn for the instance
func (instance *MDM_Policy_Result01_Authentication02) GetPropertyEnableFastFirstSignIn() (value int32, err error) {
	retValue, err := instance.GetProperty("EnableFastFirstSignIn")
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

// SetEnableWebSignIn sets the value of EnableWebSignIn for the instance
func (instance *MDM_Policy_Result01_Authentication02) SetPropertyEnableWebSignIn(value int32) (err error) {
	return instance.SetProperty("EnableWebSignIn", (value))
}

// GetEnableWebSignIn gets the value of EnableWebSignIn for the instance
func (instance *MDM_Policy_Result01_Authentication02) GetPropertyEnableWebSignIn() (value int32, err error) {
	retValue, err := instance.GetProperty("EnableWebSignIn")
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
func (instance *MDM_Policy_Result01_Authentication02) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", (value))
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Policy_Result01_Authentication02) GetPropertyInstanceID() (value string, err error) {
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
func (instance *MDM_Policy_Result01_Authentication02) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", (value))
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Policy_Result01_Authentication02) GetPropertyParentID() (value string, err error) {
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

// SetPreferredAadTenantDomainName sets the value of PreferredAadTenantDomainName for the instance
func (instance *MDM_Policy_Result01_Authentication02) SetPropertyPreferredAadTenantDomainName(value string) (err error) {
	return instance.SetProperty("PreferredAadTenantDomainName", (value))
}

// GetPreferredAadTenantDomainName gets the value of PreferredAadTenantDomainName for the instance
func (instance *MDM_Policy_Result01_Authentication02) GetPropertyPreferredAadTenantDomainName() (value string, err error) {
	retValue, err := instance.GetProperty("PreferredAadTenantDomainName")
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
