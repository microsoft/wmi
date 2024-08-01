// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
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

// MDM_Reporting_EnterpriseDataProtection01_RetrieveByTimeRange02 struct
type MDM_Reporting_EnterpriseDataProtection01_RetrieveByTimeRange02 struct {
	*cim.WmiInstance

	//
	InstanceID string

	//
	Logs string

	//
	ParentID string

	//
	StartTime string

	//
	StopTime string

	//
	Type int32
}

func NewMDM_Reporting_EnterpriseDataProtection01_RetrieveByTimeRange02Ex1(instance *cim.WmiInstance) (newInstance *MDM_Reporting_EnterpriseDataProtection01_RetrieveByTimeRange02, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_Reporting_EnterpriseDataProtection01_RetrieveByTimeRange02{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_Reporting_EnterpriseDataProtection01_RetrieveByTimeRange02Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_Reporting_EnterpriseDataProtection01_RetrieveByTimeRange02, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_Reporting_EnterpriseDataProtection01_RetrieveByTimeRange02{
		WmiInstance: tmp,
	}
	return
}

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_Reporting_EnterpriseDataProtection01_RetrieveByTimeRange02) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", (value))
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Reporting_EnterpriseDataProtection01_RetrieveByTimeRange02) GetPropertyInstanceID() (value string, err error) {
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

// SetLogs sets the value of Logs for the instance
func (instance *MDM_Reporting_EnterpriseDataProtection01_RetrieveByTimeRange02) SetPropertyLogs(value string) (err error) {
	return instance.SetProperty("Logs", (value))
}

// GetLogs gets the value of Logs for the instance
func (instance *MDM_Reporting_EnterpriseDataProtection01_RetrieveByTimeRange02) GetPropertyLogs() (value string, err error) {
	retValue, err := instance.GetProperty("Logs")
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
func (instance *MDM_Reporting_EnterpriseDataProtection01_RetrieveByTimeRange02) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", (value))
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Reporting_EnterpriseDataProtection01_RetrieveByTimeRange02) GetPropertyParentID() (value string, err error) {
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

// SetStartTime sets the value of StartTime for the instance
func (instance *MDM_Reporting_EnterpriseDataProtection01_RetrieveByTimeRange02) SetPropertyStartTime(value string) (err error) {
	return instance.SetProperty("StartTime", (value))
}

// GetStartTime gets the value of StartTime for the instance
func (instance *MDM_Reporting_EnterpriseDataProtection01_RetrieveByTimeRange02) GetPropertyStartTime() (value string, err error) {
	retValue, err := instance.GetProperty("StartTime")
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

// SetStopTime sets the value of StopTime for the instance
func (instance *MDM_Reporting_EnterpriseDataProtection01_RetrieveByTimeRange02) SetPropertyStopTime(value string) (err error) {
	return instance.SetProperty("StopTime", (value))
}

// GetStopTime gets the value of StopTime for the instance
func (instance *MDM_Reporting_EnterpriseDataProtection01_RetrieveByTimeRange02) GetPropertyStopTime() (value string, err error) {
	retValue, err := instance.GetProperty("StopTime")
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

// SetType sets the value of Type for the instance
func (instance *MDM_Reporting_EnterpriseDataProtection01_RetrieveByTimeRange02) SetPropertyType(value int32) (err error) {
	return instance.SetProperty("Type", (value))
}

// GetType gets the value of Type for the instance
func (instance *MDM_Reporting_EnterpriseDataProtection01_RetrieveByTimeRange02) GetPropertyType() (value int32, err error) {
	retValue, err := instance.GetProperty("Type")
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
