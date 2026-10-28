// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 10/28/2026 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// TeeDriverAmtProvStateNvarWmiInfo struct
type TeeDriverAmtProvStateNvarWmiInfo struct {
	*cim.WmiInstance

	//
	Active bool

	//
	AmtNvarIsValid bool

	// AMT Client Control Mode
	ClientControlMode TeeDriverAmtProvStateNvarWmiInfo_ClientControlMode

	//
	InstanceName string

	// AMT MEbx password changed/ not changed
	MebxPwdChanged uint32

	// AMT Provisioning State
	ProvisioningState TeeDriverAmtProvStateNvarWmiInfo_ProvisioningState

	// AMT ZTC enabled/ not enabled
	ZTCEnabled uint32
}

func NewTeeDriverAmtProvStateNvarWmiInfoEx1(instance *cim.WmiInstance) (newInstance *TeeDriverAmtProvStateNvarWmiInfo, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &TeeDriverAmtProvStateNvarWmiInfo{
		WmiInstance: tmp,
	}
	return
}

func NewTeeDriverAmtProvStateNvarWmiInfoEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *TeeDriverAmtProvStateNvarWmiInfo, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &TeeDriverAmtProvStateNvarWmiInfo{
		WmiInstance: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *TeeDriverAmtProvStateNvarWmiInfo) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *TeeDriverAmtProvStateNvarWmiInfo) GetPropertyActive() (value bool, err error) {
	retValue, err := instance.GetProperty("Active")
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

// SetAmtNvarIsValid sets the value of AmtNvarIsValid for the instance
func (instance *TeeDriverAmtProvStateNvarWmiInfo) SetPropertyAmtNvarIsValid(value bool) (err error) {
	return instance.SetProperty("AmtNvarIsValid", (value))
}

// GetAmtNvarIsValid gets the value of AmtNvarIsValid for the instance
func (instance *TeeDriverAmtProvStateNvarWmiInfo) GetPropertyAmtNvarIsValid() (value bool, err error) {
	retValue, err := instance.GetProperty("AmtNvarIsValid")
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

// SetClientControlMode sets the value of ClientControlMode for the instance
func (instance *TeeDriverAmtProvStateNvarWmiInfo) SetPropertyClientControlMode(value TeeDriverAmtProvStateNvarWmiInfo_ClientControlMode) (err error) {
	return instance.SetProperty("ClientControlMode", (value))
}

// GetClientControlMode gets the value of ClientControlMode for the instance
func (instance *TeeDriverAmtProvStateNvarWmiInfo) GetPropertyClientControlMode() (value TeeDriverAmtProvStateNvarWmiInfo_ClientControlMode, err error) {
	retValue, err := instance.GetProperty("ClientControlMode")
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

	value = TeeDriverAmtProvStateNvarWmiInfo_ClientControlMode(valuetmp)

	return
}

// SetInstanceName sets the value of InstanceName for the instance
func (instance *TeeDriverAmtProvStateNvarWmiInfo) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *TeeDriverAmtProvStateNvarWmiInfo) GetPropertyInstanceName() (value string, err error) {
	retValue, err := instance.GetProperty("InstanceName")
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

// SetMebxPwdChanged sets the value of MebxPwdChanged for the instance
func (instance *TeeDriverAmtProvStateNvarWmiInfo) SetPropertyMebxPwdChanged(value uint32) (err error) {
	return instance.SetProperty("MebxPwdChanged", (value))
}

// GetMebxPwdChanged gets the value of MebxPwdChanged for the instance
func (instance *TeeDriverAmtProvStateNvarWmiInfo) GetPropertyMebxPwdChanged() (value uint32, err error) {
	retValue, err := instance.GetProperty("MebxPwdChanged")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetProvisioningState sets the value of ProvisioningState for the instance
func (instance *TeeDriverAmtProvStateNvarWmiInfo) SetPropertyProvisioningState(value TeeDriverAmtProvStateNvarWmiInfo_ProvisioningState) (err error) {
	return instance.SetProperty("ProvisioningState", (value))
}

// GetProvisioningState gets the value of ProvisioningState for the instance
func (instance *TeeDriverAmtProvStateNvarWmiInfo) GetPropertyProvisioningState() (value TeeDriverAmtProvStateNvarWmiInfo_ProvisioningState, err error) {
	retValue, err := instance.GetProperty("ProvisioningState")
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

	value = TeeDriverAmtProvStateNvarWmiInfo_ProvisioningState(valuetmp)

	return
}

// SetZTCEnabled sets the value of ZTCEnabled for the instance
func (instance *TeeDriverAmtProvStateNvarWmiInfo) SetPropertyZTCEnabled(value uint32) (err error) {
	return instance.SetProperty("ZTCEnabled", (value))
}

// GetZTCEnabled gets the value of ZTCEnabled for the instance
func (instance *TeeDriverAmtProvStateNvarWmiInfo) GetPropertyZTCEnabled() (value uint32, err error) {
	retValue, err := instance.GetProperty("ZTCEnabled")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}
