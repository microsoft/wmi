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

// MDM_Win32CompatibilityAppraiser_UniversalTelemetryClient01 struct
type MDM_Win32CompatibilityAppraiser_UniversalTelemetryClient01 struct {
	*cim.WmiInstance

	//
	InstanceID string

	//
	ParentID string

	//
	UtcConnectionReport string
}

func NewMDM_Win32CompatibilityAppraiser_UniversalTelemetryClient01Ex1(instance *cim.WmiInstance) (newInstance *MDM_Win32CompatibilityAppraiser_UniversalTelemetryClient01, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_Win32CompatibilityAppraiser_UniversalTelemetryClient01{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_Win32CompatibilityAppraiser_UniversalTelemetryClient01Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_Win32CompatibilityAppraiser_UniversalTelemetryClient01, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_Win32CompatibilityAppraiser_UniversalTelemetryClient01{
		WmiInstance: tmp,
	}
	return
}

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_Win32CompatibilityAppraiser_UniversalTelemetryClient01) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Win32CompatibilityAppraiser_UniversalTelemetryClient01) GetPropertyInstanceID() (value string, err error) {
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
func (instance *MDM_Win32CompatibilityAppraiser_UniversalTelemetryClient01) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Win32CompatibilityAppraiser_UniversalTelemetryClient01) GetPropertyParentID() (value string, err error) {
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

// SetUtcConnectionReport sets the value of UtcConnectionReport for the instance
func (instance *MDM_Win32CompatibilityAppraiser_UniversalTelemetryClient01) SetPropertyUtcConnectionReport(value string) (err error) {
	return instance.SetProperty("UtcConnectionReport", value)
}

// GetUtcConnectionReport gets the value of UtcConnectionReport for the instance
func (instance *MDM_Win32CompatibilityAppraiser_UniversalTelemetryClient01) GetPropertyUtcConnectionReport() (value string, err error) {
	retValue, err := instance.GetProperty("UtcConnectionReport")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
