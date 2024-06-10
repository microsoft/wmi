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

// MDM_WindowsAdvancedThreatProtection_Configuration01 struct
type MDM_WindowsAdvancedThreatProtection_Configuration01 struct {
	*cim.WmiInstance

	//
	GroupIds string

	//
	InstanceID string

	//
	ParentID string

	//
	SampleSharing int32

	//
	TelemetryReportingFrequency int32
}

func NewMDM_WindowsAdvancedThreatProtection_Configuration01Ex1(instance *cim.WmiInstance) (newInstance *MDM_WindowsAdvancedThreatProtection_Configuration01, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_WindowsAdvancedThreatProtection_Configuration01{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_WindowsAdvancedThreatProtection_Configuration01Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_WindowsAdvancedThreatProtection_Configuration01, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_WindowsAdvancedThreatProtection_Configuration01{
		WmiInstance: tmp,
	}
	return
}

// SetGroupIds sets the value of GroupIds for the instance
func (instance *MDM_WindowsAdvancedThreatProtection_Configuration01) SetPropertyGroupIds(value string) (err error) {
	return instance.SetProperty("GroupIds", (value))
}

// GetGroupIds gets the value of GroupIds for the instance
func (instance *MDM_WindowsAdvancedThreatProtection_Configuration01) GetPropertyGroupIds() (value string, err error) {
	retValue, err := instance.GetProperty("GroupIds")
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
func (instance *MDM_WindowsAdvancedThreatProtection_Configuration01) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", (value))
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_WindowsAdvancedThreatProtection_Configuration01) GetPropertyInstanceID() (value string, err error) {
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
func (instance *MDM_WindowsAdvancedThreatProtection_Configuration01) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", (value))
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_WindowsAdvancedThreatProtection_Configuration01) GetPropertyParentID() (value string, err error) {
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

// SetSampleSharing sets the value of SampleSharing for the instance
func (instance *MDM_WindowsAdvancedThreatProtection_Configuration01) SetPropertySampleSharing(value int32) (err error) {
	return instance.SetProperty("SampleSharing", (value))
}

// GetSampleSharing gets the value of SampleSharing for the instance
func (instance *MDM_WindowsAdvancedThreatProtection_Configuration01) GetPropertySampleSharing() (value int32, err error) {
	retValue, err := instance.GetProperty("SampleSharing")
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

// SetTelemetryReportingFrequency sets the value of TelemetryReportingFrequency for the instance
func (instance *MDM_WindowsAdvancedThreatProtection_Configuration01) SetPropertyTelemetryReportingFrequency(value int32) (err error) {
	return instance.SetProperty("TelemetryReportingFrequency", (value))
}

// GetTelemetryReportingFrequency gets the value of TelemetryReportingFrequency for the instance
func (instance *MDM_WindowsAdvancedThreatProtection_Configuration01) GetPropertyTelemetryReportingFrequency() (value int32, err error) {
	retValue, err := instance.GetProperty("TelemetryReportingFrequency")
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
