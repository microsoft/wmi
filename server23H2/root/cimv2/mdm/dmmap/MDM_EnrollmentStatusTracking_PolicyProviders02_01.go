// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.CIMV2.mdm.dmmap
//
// ////////////////////////////////////////////
package dmmap

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MDM_EnrollmentStatusTracking_PolicyProviders02_01 struct
type MDM_EnrollmentStatusTracking_PolicyProviders02_01 struct {
	*cim.WmiInstance

	//
	InstallationState int32

	//
	InstanceID string

	//
	LastError int32

	//
	ParentID string

	//
	Timeout int32
}

func NewMDM_EnrollmentStatusTracking_PolicyProviders02_01Ex1(instance *cim.WmiInstance) (newInstance *MDM_EnrollmentStatusTracking_PolicyProviders02_01, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_EnrollmentStatusTracking_PolicyProviders02_01{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_EnrollmentStatusTracking_PolicyProviders02_01Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_EnrollmentStatusTracking_PolicyProviders02_01, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_EnrollmentStatusTracking_PolicyProviders02_01{
		WmiInstance: tmp,
	}
	return
}

// SetInstallationState sets the value of InstallationState for the instance
func (instance *MDM_EnrollmentStatusTracking_PolicyProviders02_01) SetPropertyInstallationState(value int32) (err error) {
	return instance.SetProperty("InstallationState", (value))
}

// GetInstallationState gets the value of InstallationState for the instance
func (instance *MDM_EnrollmentStatusTracking_PolicyProviders02_01) GetPropertyInstallationState() (value int32, err error) {
	retValue, err := instance.GetProperty("InstallationState")
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
func (instance *MDM_EnrollmentStatusTracking_PolicyProviders02_01) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", (value))
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_EnrollmentStatusTracking_PolicyProviders02_01) GetPropertyInstanceID() (value string, err error) {
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

// SetLastError sets the value of LastError for the instance
func (instance *MDM_EnrollmentStatusTracking_PolicyProviders02_01) SetPropertyLastError(value int32) (err error) {
	return instance.SetProperty("LastError", (value))
}

// GetLastError gets the value of LastError for the instance
func (instance *MDM_EnrollmentStatusTracking_PolicyProviders02_01) GetPropertyLastError() (value int32, err error) {
	retValue, err := instance.GetProperty("LastError")
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

// SetParentID sets the value of ParentID for the instance
func (instance *MDM_EnrollmentStatusTracking_PolicyProviders02_01) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", (value))
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_EnrollmentStatusTracking_PolicyProviders02_01) GetPropertyParentID() (value string, err error) {
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

// SetTimeout sets the value of Timeout for the instance
func (instance *MDM_EnrollmentStatusTracking_PolicyProviders02_01) SetPropertyTimeout(value int32) (err error) {
	return instance.SetProperty("Timeout", (value))
}

// GetTimeout gets the value of Timeout for the instance
func (instance *MDM_EnrollmentStatusTracking_PolicyProviders02_01) GetPropertyTimeout() (value int32, err error) {
	retValue, err := instance.GetProperty("Timeout")
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
