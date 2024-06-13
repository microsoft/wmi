// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
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

// MDM_DeviceStatus_CellularIdentities01_01 struct
type MDM_DeviceStatus_CellularIdentities01_01 struct {
	*cim.WmiInstance

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

func NewMDM_DeviceStatus_CellularIdentities01_01Ex1(instance *cim.WmiInstance) (newInstance *MDM_DeviceStatus_CellularIdentities01_01, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_DeviceStatus_CellularIdentities01_01{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_DeviceStatus_CellularIdentities01_01Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_DeviceStatus_CellularIdentities01_01, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_DeviceStatus_CellularIdentities01_01{
		WmiInstance: tmp,
	}
	return
}

// SetCommercializationOperator sets the value of CommercializationOperator for the instance
func (instance *MDM_DeviceStatus_CellularIdentities01_01) SetPropertyCommercializationOperator(value string) (err error) {
	return instance.SetProperty("CommercializationOperator", (value))
}

// GetCommercializationOperator gets the value of CommercializationOperator for the instance
func (instance *MDM_DeviceStatus_CellularIdentities01_01) GetPropertyCommercializationOperator() (value string, err error) {
	retValue, err := instance.GetProperty("CommercializationOperator")
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

// SetICCID sets the value of ICCID for the instance
func (instance *MDM_DeviceStatus_CellularIdentities01_01) SetPropertyICCID(value string) (err error) {
	return instance.SetProperty("ICCID", (value))
}

// GetICCID gets the value of ICCID for the instance
func (instance *MDM_DeviceStatus_CellularIdentities01_01) GetPropertyICCID() (value string, err error) {
	retValue, err := instance.GetProperty("ICCID")
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

// SetIMSI sets the value of IMSI for the instance
func (instance *MDM_DeviceStatus_CellularIdentities01_01) SetPropertyIMSI(value string) (err error) {
	return instance.SetProperty("IMSI", (value))
}

// GetIMSI gets the value of IMSI for the instance
func (instance *MDM_DeviceStatus_CellularIdentities01_01) GetPropertyIMSI() (value string, err error) {
	retValue, err := instance.GetProperty("IMSI")
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
func (instance *MDM_DeviceStatus_CellularIdentities01_01) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", (value))
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_DeviceStatus_CellularIdentities01_01) GetPropertyInstanceID() (value string, err error) {
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
func (instance *MDM_DeviceStatus_CellularIdentities01_01) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", (value))
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_DeviceStatus_CellularIdentities01_01) GetPropertyParentID() (value string, err error) {
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

// SetPhoneNumber sets the value of PhoneNumber for the instance
func (instance *MDM_DeviceStatus_CellularIdentities01_01) SetPropertyPhoneNumber(value string) (err error) {
	return instance.SetProperty("PhoneNumber", (value))
}

// GetPhoneNumber gets the value of PhoneNumber for the instance
func (instance *MDM_DeviceStatus_CellularIdentities01_01) GetPropertyPhoneNumber() (value string, err error) {
	retValue, err := instance.GetProperty("PhoneNumber")
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

// SetRoamingCompliance sets the value of RoamingCompliance for the instance
func (instance *MDM_DeviceStatus_CellularIdentities01_01) SetPropertyRoamingCompliance(value bool) (err error) {
	return instance.SetProperty("RoamingCompliance", (value))
}

// GetRoamingCompliance gets the value of RoamingCompliance for the instance
func (instance *MDM_DeviceStatus_CellularIdentities01_01) GetPropertyRoamingCompliance() (value bool, err error) {
	retValue, err := instance.GetProperty("RoamingCompliance")
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

// SetRoamingStatus sets the value of RoamingStatus for the instance
func (instance *MDM_DeviceStatus_CellularIdentities01_01) SetPropertyRoamingStatus(value bool) (err error) {
	return instance.SetProperty("RoamingStatus", (value))
}

// GetRoamingStatus gets the value of RoamingStatus for the instance
func (instance *MDM_DeviceStatus_CellularIdentities01_01) GetPropertyRoamingStatus() (value bool, err error) {
	retValue, err := instance.GetProperty("RoamingStatus")
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
