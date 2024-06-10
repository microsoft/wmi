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

// MDM_Policy_Config01_Power02 struct
type MDM_Policy_Config01_Power02 struct {
	*cim.WmiInstance

	//
	AllowStandbyStatesWhenSleepingOnBattery string

	//
	AllowStandbyWhenSleepingPluggedIn string

	//
	DisplayOffTimeoutOnBattery string

	//
	DisplayOffTimeoutPluggedIn string

	//
	HibernateTimeoutOnBattery string

	//
	HibernateTimeoutPluggedIn string

	//
	InstanceID string

	//
	ParentID string

	//
	RequirePasswordWhenComputerWakesOnBattery string

	//
	RequirePasswordWhenComputerWakesPluggedIn string

	//
	StandbyTimeoutOnBattery string

	//
	StandbyTimeoutPluggedIn string
}

func NewMDM_Policy_Config01_Power02Ex1(instance *cim.WmiInstance) (newInstance *MDM_Policy_Config01_Power02, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_Policy_Config01_Power02{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_Policy_Config01_Power02Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_Policy_Config01_Power02, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_Policy_Config01_Power02{
		WmiInstance: tmp,
	}
	return
}

// SetAllowStandbyStatesWhenSleepingOnBattery sets the value of AllowStandbyStatesWhenSleepingOnBattery for the instance
func (instance *MDM_Policy_Config01_Power02) SetPropertyAllowStandbyStatesWhenSleepingOnBattery(value string) (err error) {
	return instance.SetProperty("AllowStandbyStatesWhenSleepingOnBattery", (value))
}

// GetAllowStandbyStatesWhenSleepingOnBattery gets the value of AllowStandbyStatesWhenSleepingOnBattery for the instance
func (instance *MDM_Policy_Config01_Power02) GetPropertyAllowStandbyStatesWhenSleepingOnBattery() (value string, err error) {
	retValue, err := instance.GetProperty("AllowStandbyStatesWhenSleepingOnBattery")
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

// SetAllowStandbyWhenSleepingPluggedIn sets the value of AllowStandbyWhenSleepingPluggedIn for the instance
func (instance *MDM_Policy_Config01_Power02) SetPropertyAllowStandbyWhenSleepingPluggedIn(value string) (err error) {
	return instance.SetProperty("AllowStandbyWhenSleepingPluggedIn", (value))
}

// GetAllowStandbyWhenSleepingPluggedIn gets the value of AllowStandbyWhenSleepingPluggedIn for the instance
func (instance *MDM_Policy_Config01_Power02) GetPropertyAllowStandbyWhenSleepingPluggedIn() (value string, err error) {
	retValue, err := instance.GetProperty("AllowStandbyWhenSleepingPluggedIn")
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

// SetDisplayOffTimeoutOnBattery sets the value of DisplayOffTimeoutOnBattery for the instance
func (instance *MDM_Policy_Config01_Power02) SetPropertyDisplayOffTimeoutOnBattery(value string) (err error) {
	return instance.SetProperty("DisplayOffTimeoutOnBattery", (value))
}

// GetDisplayOffTimeoutOnBattery gets the value of DisplayOffTimeoutOnBattery for the instance
func (instance *MDM_Policy_Config01_Power02) GetPropertyDisplayOffTimeoutOnBattery() (value string, err error) {
	retValue, err := instance.GetProperty("DisplayOffTimeoutOnBattery")
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

// SetDisplayOffTimeoutPluggedIn sets the value of DisplayOffTimeoutPluggedIn for the instance
func (instance *MDM_Policy_Config01_Power02) SetPropertyDisplayOffTimeoutPluggedIn(value string) (err error) {
	return instance.SetProperty("DisplayOffTimeoutPluggedIn", (value))
}

// GetDisplayOffTimeoutPluggedIn gets the value of DisplayOffTimeoutPluggedIn for the instance
func (instance *MDM_Policy_Config01_Power02) GetPropertyDisplayOffTimeoutPluggedIn() (value string, err error) {
	retValue, err := instance.GetProperty("DisplayOffTimeoutPluggedIn")
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

// SetHibernateTimeoutOnBattery sets the value of HibernateTimeoutOnBattery for the instance
func (instance *MDM_Policy_Config01_Power02) SetPropertyHibernateTimeoutOnBattery(value string) (err error) {
	return instance.SetProperty("HibernateTimeoutOnBattery", (value))
}

// GetHibernateTimeoutOnBattery gets the value of HibernateTimeoutOnBattery for the instance
func (instance *MDM_Policy_Config01_Power02) GetPropertyHibernateTimeoutOnBattery() (value string, err error) {
	retValue, err := instance.GetProperty("HibernateTimeoutOnBattery")
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

// SetHibernateTimeoutPluggedIn sets the value of HibernateTimeoutPluggedIn for the instance
func (instance *MDM_Policy_Config01_Power02) SetPropertyHibernateTimeoutPluggedIn(value string) (err error) {
	return instance.SetProperty("HibernateTimeoutPluggedIn", (value))
}

// GetHibernateTimeoutPluggedIn gets the value of HibernateTimeoutPluggedIn for the instance
func (instance *MDM_Policy_Config01_Power02) GetPropertyHibernateTimeoutPluggedIn() (value string, err error) {
	retValue, err := instance.GetProperty("HibernateTimeoutPluggedIn")
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
func (instance *MDM_Policy_Config01_Power02) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", (value))
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Policy_Config01_Power02) GetPropertyInstanceID() (value string, err error) {
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
func (instance *MDM_Policy_Config01_Power02) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", (value))
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Policy_Config01_Power02) GetPropertyParentID() (value string, err error) {
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

// SetRequirePasswordWhenComputerWakesOnBattery sets the value of RequirePasswordWhenComputerWakesOnBattery for the instance
func (instance *MDM_Policy_Config01_Power02) SetPropertyRequirePasswordWhenComputerWakesOnBattery(value string) (err error) {
	return instance.SetProperty("RequirePasswordWhenComputerWakesOnBattery", (value))
}

// GetRequirePasswordWhenComputerWakesOnBattery gets the value of RequirePasswordWhenComputerWakesOnBattery for the instance
func (instance *MDM_Policy_Config01_Power02) GetPropertyRequirePasswordWhenComputerWakesOnBattery() (value string, err error) {
	retValue, err := instance.GetProperty("RequirePasswordWhenComputerWakesOnBattery")
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

// SetRequirePasswordWhenComputerWakesPluggedIn sets the value of RequirePasswordWhenComputerWakesPluggedIn for the instance
func (instance *MDM_Policy_Config01_Power02) SetPropertyRequirePasswordWhenComputerWakesPluggedIn(value string) (err error) {
	return instance.SetProperty("RequirePasswordWhenComputerWakesPluggedIn", (value))
}

// GetRequirePasswordWhenComputerWakesPluggedIn gets the value of RequirePasswordWhenComputerWakesPluggedIn for the instance
func (instance *MDM_Policy_Config01_Power02) GetPropertyRequirePasswordWhenComputerWakesPluggedIn() (value string, err error) {
	retValue, err := instance.GetProperty("RequirePasswordWhenComputerWakesPluggedIn")
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

// SetStandbyTimeoutOnBattery sets the value of StandbyTimeoutOnBattery for the instance
func (instance *MDM_Policy_Config01_Power02) SetPropertyStandbyTimeoutOnBattery(value string) (err error) {
	return instance.SetProperty("StandbyTimeoutOnBattery", (value))
}

// GetStandbyTimeoutOnBattery gets the value of StandbyTimeoutOnBattery for the instance
func (instance *MDM_Policy_Config01_Power02) GetPropertyStandbyTimeoutOnBattery() (value string, err error) {
	retValue, err := instance.GetProperty("StandbyTimeoutOnBattery")
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

// SetStandbyTimeoutPluggedIn sets the value of StandbyTimeoutPluggedIn for the instance
func (instance *MDM_Policy_Config01_Power02) SetPropertyStandbyTimeoutPluggedIn(value string) (err error) {
	return instance.SetProperty("StandbyTimeoutPluggedIn", (value))
}

// GetStandbyTimeoutPluggedIn gets the value of StandbyTimeoutPluggedIn for the instance
func (instance *MDM_Policy_Config01_Power02) GetPropertyStandbyTimeoutPluggedIn() (value string, err error) {
	retValue, err := instance.GetProperty("StandbyTimeoutPluggedIn")
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
