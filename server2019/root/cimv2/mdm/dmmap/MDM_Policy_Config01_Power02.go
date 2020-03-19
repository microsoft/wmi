// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.CIMV2.mdm.dmmap
//////////////////////////////////////////////
package dmmap

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
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
	EnergySaverBatteryThresholdOnBattery int32

	//
	EnergySaverBatteryThresholdPluggedIn int32

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
	SelectLidCloseActionOnBattery int32

	//
	SelectLidCloseActionPluggedIn int32

	//
	SelectPowerButtonActionOnBattery int32

	//
	SelectPowerButtonActionPluggedIn int32

	//
	SelectSleepButtonActionOnBattery int32

	//
	SelectSleepButtonActionPluggedIn int32

	//
	StandbyTimeoutOnBattery string

	//
	StandbyTimeoutPluggedIn string

	//
	TurnOffHybridSleepOnBattery int32

	//
	TurnOffHybridSleepPluggedIn int32

	//
	UnattendedSleepTimeoutOnBattery int32

	//
	UnattendedSleepTimeoutPluggedIn int32
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
	return instance.SetProperty("AllowStandbyStatesWhenSleepingOnBattery", value)
}

// GetAllowStandbyStatesWhenSleepingOnBattery gets the value of AllowStandbyStatesWhenSleepingOnBattery for the instance
func (instance *MDM_Policy_Config01_Power02) GetPropertyAllowStandbyStatesWhenSleepingOnBattery() (value string, err error) {
	retValue, err := instance.GetProperty("AllowStandbyStatesWhenSleepingOnBattery")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAllowStandbyWhenSleepingPluggedIn sets the value of AllowStandbyWhenSleepingPluggedIn for the instance
func (instance *MDM_Policy_Config01_Power02) SetPropertyAllowStandbyWhenSleepingPluggedIn(value string) (err error) {
	return instance.SetProperty("AllowStandbyWhenSleepingPluggedIn", value)
}

// GetAllowStandbyWhenSleepingPluggedIn gets the value of AllowStandbyWhenSleepingPluggedIn for the instance
func (instance *MDM_Policy_Config01_Power02) GetPropertyAllowStandbyWhenSleepingPluggedIn() (value string, err error) {
	retValue, err := instance.GetProperty("AllowStandbyWhenSleepingPluggedIn")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDisplayOffTimeoutOnBattery sets the value of DisplayOffTimeoutOnBattery for the instance
func (instance *MDM_Policy_Config01_Power02) SetPropertyDisplayOffTimeoutOnBattery(value string) (err error) {
	return instance.SetProperty("DisplayOffTimeoutOnBattery", value)
}

// GetDisplayOffTimeoutOnBattery gets the value of DisplayOffTimeoutOnBattery for the instance
func (instance *MDM_Policy_Config01_Power02) GetPropertyDisplayOffTimeoutOnBattery() (value string, err error) {
	retValue, err := instance.GetProperty("DisplayOffTimeoutOnBattery")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDisplayOffTimeoutPluggedIn sets the value of DisplayOffTimeoutPluggedIn for the instance
func (instance *MDM_Policy_Config01_Power02) SetPropertyDisplayOffTimeoutPluggedIn(value string) (err error) {
	return instance.SetProperty("DisplayOffTimeoutPluggedIn", value)
}

// GetDisplayOffTimeoutPluggedIn gets the value of DisplayOffTimeoutPluggedIn for the instance
func (instance *MDM_Policy_Config01_Power02) GetPropertyDisplayOffTimeoutPluggedIn() (value string, err error) {
	retValue, err := instance.GetProperty("DisplayOffTimeoutPluggedIn")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEnergySaverBatteryThresholdOnBattery sets the value of EnergySaverBatteryThresholdOnBattery for the instance
func (instance *MDM_Policy_Config01_Power02) SetPropertyEnergySaverBatteryThresholdOnBattery(value int32) (err error) {
	return instance.SetProperty("EnergySaverBatteryThresholdOnBattery", value)
}

// GetEnergySaverBatteryThresholdOnBattery gets the value of EnergySaverBatteryThresholdOnBattery for the instance
func (instance *MDM_Policy_Config01_Power02) GetPropertyEnergySaverBatteryThresholdOnBattery() (value int32, err error) {
	retValue, err := instance.GetProperty("EnergySaverBatteryThresholdOnBattery")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEnergySaverBatteryThresholdPluggedIn sets the value of EnergySaverBatteryThresholdPluggedIn for the instance
func (instance *MDM_Policy_Config01_Power02) SetPropertyEnergySaverBatteryThresholdPluggedIn(value int32) (err error) {
	return instance.SetProperty("EnergySaverBatteryThresholdPluggedIn", value)
}

// GetEnergySaverBatteryThresholdPluggedIn gets the value of EnergySaverBatteryThresholdPluggedIn for the instance
func (instance *MDM_Policy_Config01_Power02) GetPropertyEnergySaverBatteryThresholdPluggedIn() (value int32, err error) {
	retValue, err := instance.GetProperty("EnergySaverBatteryThresholdPluggedIn")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetHibernateTimeoutOnBattery sets the value of HibernateTimeoutOnBattery for the instance
func (instance *MDM_Policy_Config01_Power02) SetPropertyHibernateTimeoutOnBattery(value string) (err error) {
	return instance.SetProperty("HibernateTimeoutOnBattery", value)
}

// GetHibernateTimeoutOnBattery gets the value of HibernateTimeoutOnBattery for the instance
func (instance *MDM_Policy_Config01_Power02) GetPropertyHibernateTimeoutOnBattery() (value string, err error) {
	retValue, err := instance.GetProperty("HibernateTimeoutOnBattery")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetHibernateTimeoutPluggedIn sets the value of HibernateTimeoutPluggedIn for the instance
func (instance *MDM_Policy_Config01_Power02) SetPropertyHibernateTimeoutPluggedIn(value string) (err error) {
	return instance.SetProperty("HibernateTimeoutPluggedIn", value)
}

// GetHibernateTimeoutPluggedIn gets the value of HibernateTimeoutPluggedIn for the instance
func (instance *MDM_Policy_Config01_Power02) GetPropertyHibernateTimeoutPluggedIn() (value string, err error) {
	retValue, err := instance.GetProperty("HibernateTimeoutPluggedIn")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_Policy_Config01_Power02) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Policy_Config01_Power02) GetPropertyInstanceID() (value string, err error) {
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

// SetParentID sets the value of ParentID for the instance
func (instance *MDM_Policy_Config01_Power02) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Policy_Config01_Power02) GetPropertyParentID() (value string, err error) {
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

// SetRequirePasswordWhenComputerWakesOnBattery sets the value of RequirePasswordWhenComputerWakesOnBattery for the instance
func (instance *MDM_Policy_Config01_Power02) SetPropertyRequirePasswordWhenComputerWakesOnBattery(value string) (err error) {
	return instance.SetProperty("RequirePasswordWhenComputerWakesOnBattery", value)
}

// GetRequirePasswordWhenComputerWakesOnBattery gets the value of RequirePasswordWhenComputerWakesOnBattery for the instance
func (instance *MDM_Policy_Config01_Power02) GetPropertyRequirePasswordWhenComputerWakesOnBattery() (value string, err error) {
	retValue, err := instance.GetProperty("RequirePasswordWhenComputerWakesOnBattery")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRequirePasswordWhenComputerWakesPluggedIn sets the value of RequirePasswordWhenComputerWakesPluggedIn for the instance
func (instance *MDM_Policy_Config01_Power02) SetPropertyRequirePasswordWhenComputerWakesPluggedIn(value string) (err error) {
	return instance.SetProperty("RequirePasswordWhenComputerWakesPluggedIn", value)
}

// GetRequirePasswordWhenComputerWakesPluggedIn gets the value of RequirePasswordWhenComputerWakesPluggedIn for the instance
func (instance *MDM_Policy_Config01_Power02) GetPropertyRequirePasswordWhenComputerWakesPluggedIn() (value string, err error) {
	retValue, err := instance.GetProperty("RequirePasswordWhenComputerWakesPluggedIn")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSelectLidCloseActionOnBattery sets the value of SelectLidCloseActionOnBattery for the instance
func (instance *MDM_Policy_Config01_Power02) SetPropertySelectLidCloseActionOnBattery(value int32) (err error) {
	return instance.SetProperty("SelectLidCloseActionOnBattery", value)
}

// GetSelectLidCloseActionOnBattery gets the value of SelectLidCloseActionOnBattery for the instance
func (instance *MDM_Policy_Config01_Power02) GetPropertySelectLidCloseActionOnBattery() (value int32, err error) {
	retValue, err := instance.GetProperty("SelectLidCloseActionOnBattery")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSelectLidCloseActionPluggedIn sets the value of SelectLidCloseActionPluggedIn for the instance
func (instance *MDM_Policy_Config01_Power02) SetPropertySelectLidCloseActionPluggedIn(value int32) (err error) {
	return instance.SetProperty("SelectLidCloseActionPluggedIn", value)
}

// GetSelectLidCloseActionPluggedIn gets the value of SelectLidCloseActionPluggedIn for the instance
func (instance *MDM_Policy_Config01_Power02) GetPropertySelectLidCloseActionPluggedIn() (value int32, err error) {
	retValue, err := instance.GetProperty("SelectLidCloseActionPluggedIn")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSelectPowerButtonActionOnBattery sets the value of SelectPowerButtonActionOnBattery for the instance
func (instance *MDM_Policy_Config01_Power02) SetPropertySelectPowerButtonActionOnBattery(value int32) (err error) {
	return instance.SetProperty("SelectPowerButtonActionOnBattery", value)
}

// GetSelectPowerButtonActionOnBattery gets the value of SelectPowerButtonActionOnBattery for the instance
func (instance *MDM_Policy_Config01_Power02) GetPropertySelectPowerButtonActionOnBattery() (value int32, err error) {
	retValue, err := instance.GetProperty("SelectPowerButtonActionOnBattery")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSelectPowerButtonActionPluggedIn sets the value of SelectPowerButtonActionPluggedIn for the instance
func (instance *MDM_Policy_Config01_Power02) SetPropertySelectPowerButtonActionPluggedIn(value int32) (err error) {
	return instance.SetProperty("SelectPowerButtonActionPluggedIn", value)
}

// GetSelectPowerButtonActionPluggedIn gets the value of SelectPowerButtonActionPluggedIn for the instance
func (instance *MDM_Policy_Config01_Power02) GetPropertySelectPowerButtonActionPluggedIn() (value int32, err error) {
	retValue, err := instance.GetProperty("SelectPowerButtonActionPluggedIn")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSelectSleepButtonActionOnBattery sets the value of SelectSleepButtonActionOnBattery for the instance
func (instance *MDM_Policy_Config01_Power02) SetPropertySelectSleepButtonActionOnBattery(value int32) (err error) {
	return instance.SetProperty("SelectSleepButtonActionOnBattery", value)
}

// GetSelectSleepButtonActionOnBattery gets the value of SelectSleepButtonActionOnBattery for the instance
func (instance *MDM_Policy_Config01_Power02) GetPropertySelectSleepButtonActionOnBattery() (value int32, err error) {
	retValue, err := instance.GetProperty("SelectSleepButtonActionOnBattery")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSelectSleepButtonActionPluggedIn sets the value of SelectSleepButtonActionPluggedIn for the instance
func (instance *MDM_Policy_Config01_Power02) SetPropertySelectSleepButtonActionPluggedIn(value int32) (err error) {
	return instance.SetProperty("SelectSleepButtonActionPluggedIn", value)
}

// GetSelectSleepButtonActionPluggedIn gets the value of SelectSleepButtonActionPluggedIn for the instance
func (instance *MDM_Policy_Config01_Power02) GetPropertySelectSleepButtonActionPluggedIn() (value int32, err error) {
	retValue, err := instance.GetProperty("SelectSleepButtonActionPluggedIn")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStandbyTimeoutOnBattery sets the value of StandbyTimeoutOnBattery for the instance
func (instance *MDM_Policy_Config01_Power02) SetPropertyStandbyTimeoutOnBattery(value string) (err error) {
	return instance.SetProperty("StandbyTimeoutOnBattery", value)
}

// GetStandbyTimeoutOnBattery gets the value of StandbyTimeoutOnBattery for the instance
func (instance *MDM_Policy_Config01_Power02) GetPropertyStandbyTimeoutOnBattery() (value string, err error) {
	retValue, err := instance.GetProperty("StandbyTimeoutOnBattery")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStandbyTimeoutPluggedIn sets the value of StandbyTimeoutPluggedIn for the instance
func (instance *MDM_Policy_Config01_Power02) SetPropertyStandbyTimeoutPluggedIn(value string) (err error) {
	return instance.SetProperty("StandbyTimeoutPluggedIn", value)
}

// GetStandbyTimeoutPluggedIn gets the value of StandbyTimeoutPluggedIn for the instance
func (instance *MDM_Policy_Config01_Power02) GetPropertyStandbyTimeoutPluggedIn() (value string, err error) {
	retValue, err := instance.GetProperty("StandbyTimeoutPluggedIn")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTurnOffHybridSleepOnBattery sets the value of TurnOffHybridSleepOnBattery for the instance
func (instance *MDM_Policy_Config01_Power02) SetPropertyTurnOffHybridSleepOnBattery(value int32) (err error) {
	return instance.SetProperty("TurnOffHybridSleepOnBattery", value)
}

// GetTurnOffHybridSleepOnBattery gets the value of TurnOffHybridSleepOnBattery for the instance
func (instance *MDM_Policy_Config01_Power02) GetPropertyTurnOffHybridSleepOnBattery() (value int32, err error) {
	retValue, err := instance.GetProperty("TurnOffHybridSleepOnBattery")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTurnOffHybridSleepPluggedIn sets the value of TurnOffHybridSleepPluggedIn for the instance
func (instance *MDM_Policy_Config01_Power02) SetPropertyTurnOffHybridSleepPluggedIn(value int32) (err error) {
	return instance.SetProperty("TurnOffHybridSleepPluggedIn", value)
}

// GetTurnOffHybridSleepPluggedIn gets the value of TurnOffHybridSleepPluggedIn for the instance
func (instance *MDM_Policy_Config01_Power02) GetPropertyTurnOffHybridSleepPluggedIn() (value int32, err error) {
	retValue, err := instance.GetProperty("TurnOffHybridSleepPluggedIn")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUnattendedSleepTimeoutOnBattery sets the value of UnattendedSleepTimeoutOnBattery for the instance
func (instance *MDM_Policy_Config01_Power02) SetPropertyUnattendedSleepTimeoutOnBattery(value int32) (err error) {
	return instance.SetProperty("UnattendedSleepTimeoutOnBattery", value)
}

// GetUnattendedSleepTimeoutOnBattery gets the value of UnattendedSleepTimeoutOnBattery for the instance
func (instance *MDM_Policy_Config01_Power02) GetPropertyUnattendedSleepTimeoutOnBattery() (value int32, err error) {
	retValue, err := instance.GetProperty("UnattendedSleepTimeoutOnBattery")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUnattendedSleepTimeoutPluggedIn sets the value of UnattendedSleepTimeoutPluggedIn for the instance
func (instance *MDM_Policy_Config01_Power02) SetPropertyUnattendedSleepTimeoutPluggedIn(value int32) (err error) {
	return instance.SetProperty("UnattendedSleepTimeoutPluggedIn", value)
}

// GetUnattendedSleepTimeoutPluggedIn gets the value of UnattendedSleepTimeoutPluggedIn for the instance
func (instance *MDM_Policy_Config01_Power02) GetPropertyUnattendedSleepTimeoutPluggedIn() (value int32, err error) {
	retValue, err := instance.GetProperty("UnattendedSleepTimeoutPluggedIn")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}
