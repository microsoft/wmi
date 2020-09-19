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

// MDM_PassportForWork_Policies02 struct
type MDM_PassportForWork_Policies02 struct {
	*cim.WmiInstance

	//
	EnablePinRecovery bool

	//
	InstanceID string

	//
	ParentID string

	//
	RequireSecurityDevice bool

	//
	UseHelloCertificatesAsSmartCardCertificates bool

	//
	UsePassportForWork bool
}

func NewMDM_PassportForWork_Policies02Ex1(instance *cim.WmiInstance) (newInstance *MDM_PassportForWork_Policies02, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_PassportForWork_Policies02{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_PassportForWork_Policies02Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_PassportForWork_Policies02, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_PassportForWork_Policies02{
		WmiInstance: tmp,
	}
	return
}

// SetEnablePinRecovery sets the value of EnablePinRecovery for the instance
func (instance *MDM_PassportForWork_Policies02) SetPropertyEnablePinRecovery(value bool) (err error) {
	return instance.SetProperty("EnablePinRecovery", (value))
}

// GetEnablePinRecovery gets the value of EnablePinRecovery for the instance
func (instance *MDM_PassportForWork_Policies02) GetPropertyEnablePinRecovery() (value bool, err error) {
	retValue, err := instance.GetProperty("EnablePinRecovery")
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
func (instance *MDM_PassportForWork_Policies02) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", (value))
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_PassportForWork_Policies02) GetPropertyInstanceID() (value string, err error) {
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
func (instance *MDM_PassportForWork_Policies02) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", (value))
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_PassportForWork_Policies02) GetPropertyParentID() (value string, err error) {
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

// SetRequireSecurityDevice sets the value of RequireSecurityDevice for the instance
func (instance *MDM_PassportForWork_Policies02) SetPropertyRequireSecurityDevice(value bool) (err error) {
	return instance.SetProperty("RequireSecurityDevice", (value))
}

// GetRequireSecurityDevice gets the value of RequireSecurityDevice for the instance
func (instance *MDM_PassportForWork_Policies02) GetPropertyRequireSecurityDevice() (value bool, err error) {
	retValue, err := instance.GetProperty("RequireSecurityDevice")
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

// SetUseHelloCertificatesAsSmartCardCertificates sets the value of UseHelloCertificatesAsSmartCardCertificates for the instance
func (instance *MDM_PassportForWork_Policies02) SetPropertyUseHelloCertificatesAsSmartCardCertificates(value bool) (err error) {
	return instance.SetProperty("UseHelloCertificatesAsSmartCardCertificates", (value))
}

// GetUseHelloCertificatesAsSmartCardCertificates gets the value of UseHelloCertificatesAsSmartCardCertificates for the instance
func (instance *MDM_PassportForWork_Policies02) GetPropertyUseHelloCertificatesAsSmartCardCertificates() (value bool, err error) {
	retValue, err := instance.GetProperty("UseHelloCertificatesAsSmartCardCertificates")
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

// SetUsePassportForWork sets the value of UsePassportForWork for the instance
func (instance *MDM_PassportForWork_Policies02) SetPropertyUsePassportForWork(value bool) (err error) {
	return instance.SetProperty("UsePassportForWork", (value))
}

// GetUsePassportForWork gets the value of UsePassportForWork for the instance
func (instance *MDM_PassportForWork_Policies02) GetPropertyUsePassportForWork() (value bool, err error) {
	retValue, err := instance.GetProperty("UsePassportForWork")
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
