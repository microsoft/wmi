// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2.mdm.dmmap
//////////////////////////////////////////////
package dmmap

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MDM_DeviceStatus_CellularIdentities01_01 struct
type MDM_DeviceStatus_CellularIdentities01_01 struct {
	cim.WmiInstance

	//
	CommercializationOperator string

	//
	ICCID string

	//
	IMSI string

	//
	InstanceID string

	//
	ParentID string

	//
	PhoneNumber string

	//
	RoamingCompliance bool

	//
	RoamingStatus bool
}

// SetCommercializationOperator sets the value of CommercializationOperator for the instance
func (instance *MDM_DeviceStatus_CellularIdentities01_01) SetPropertyCommercializationOperator(value string) (err error) {
	return instance.SetProperty("CommercializationOperator", value)
}

// GetCommercializationOperator gets the value of CommercializationOperator for the instance
func (instance *MDM_DeviceStatus_CellularIdentities01_01) GetPropertyCommercializationOperator() (value string, err error) {
	retValue, err := instance.GetProperty("CommercializationOperator")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetICCID sets the value of ICCID for the instance
func (instance *MDM_DeviceStatus_CellularIdentities01_01) SetPropertyICCID(value string) (err error) {
	return instance.SetProperty("ICCID", value)
}

// GetICCID gets the value of ICCID for the instance
func (instance *MDM_DeviceStatus_CellularIdentities01_01) GetPropertyICCID() (value string, err error) {
	retValue, err := instance.GetProperty("ICCID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIMSI sets the value of IMSI for the instance
func (instance *MDM_DeviceStatus_CellularIdentities01_01) SetPropertyIMSI(value string) (err error) {
	return instance.SetProperty("IMSI", value)
}

// GetIMSI gets the value of IMSI for the instance
func (instance *MDM_DeviceStatus_CellularIdentities01_01) GetPropertyIMSI() (value string, err error) {
	retValue, err := instance.GetProperty("IMSI")
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
func (instance *MDM_DeviceStatus_CellularIdentities01_01) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_DeviceStatus_CellularIdentities01_01) GetPropertyInstanceID() (value string, err error) {
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
func (instance *MDM_DeviceStatus_CellularIdentities01_01) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_DeviceStatus_CellularIdentities01_01) GetPropertyParentID() (value string, err error) {
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

// SetPhoneNumber sets the value of PhoneNumber for the instance
func (instance *MDM_DeviceStatus_CellularIdentities01_01) SetPropertyPhoneNumber(value string) (err error) {
	return instance.SetProperty("PhoneNumber", value)
}

// GetPhoneNumber gets the value of PhoneNumber for the instance
func (instance *MDM_DeviceStatus_CellularIdentities01_01) GetPropertyPhoneNumber() (value string, err error) {
	retValue, err := instance.GetProperty("PhoneNumber")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRoamingCompliance sets the value of RoamingCompliance for the instance
func (instance *MDM_DeviceStatus_CellularIdentities01_01) SetPropertyRoamingCompliance(value bool) (err error) {
	return instance.SetProperty("RoamingCompliance", value)
}

// GetRoamingCompliance gets the value of RoamingCompliance for the instance
func (instance *MDM_DeviceStatus_CellularIdentities01_01) GetPropertyRoamingCompliance() (value bool, err error) {
	retValue, err := instance.GetProperty("RoamingCompliance")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRoamingStatus sets the value of RoamingStatus for the instance
func (instance *MDM_DeviceStatus_CellularIdentities01_01) SetPropertyRoamingStatus(value bool) (err error) {
	return instance.SetProperty("RoamingStatus", value)
}

// GetRoamingStatus gets the value of RoamingStatus for the instance
func (instance *MDM_DeviceStatus_CellularIdentities01_01) GetPropertyRoamingStatus() (value bool, err error) {
	retValue, err := instance.GetProperty("RoamingStatus")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}
