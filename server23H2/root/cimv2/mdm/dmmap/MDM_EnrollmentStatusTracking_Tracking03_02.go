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

// MDM_EnrollmentStatusTracking_Tracking03_02 struct
type MDM_EnrollmentStatusTracking_Tracking03_02 struct {
	*cim.WmiInstance

	//
	AppId string

	//
	DisplayName string

	//
	ErrorHresult int32

	//
	InstallationState int32

	//
	InstanceID string

	//
	ParentID string

	//
	RebootRequired int32

	//
	TrackingUri string
}

func NewMDM_EnrollmentStatusTracking_Tracking03_02Ex1(instance *cim.WmiInstance) (newInstance *MDM_EnrollmentStatusTracking_Tracking03_02, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_EnrollmentStatusTracking_Tracking03_02{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_EnrollmentStatusTracking_Tracking03_02Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_EnrollmentStatusTracking_Tracking03_02, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_EnrollmentStatusTracking_Tracking03_02{
		WmiInstance: tmp,
	}
	return
}

// SetAppId sets the value of AppId for the instance
func (instance *MDM_EnrollmentStatusTracking_Tracking03_02) SetPropertyAppId(value string) (err error) {
	return instance.SetProperty("AppId", (value))
}

// GetAppId gets the value of AppId for the instance
func (instance *MDM_EnrollmentStatusTracking_Tracking03_02) GetPropertyAppId() (value string, err error) {
	retValue, err := instance.GetProperty("AppId")
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

// SetDisplayName sets the value of DisplayName for the instance
func (instance *MDM_EnrollmentStatusTracking_Tracking03_02) SetPropertyDisplayName(value string) (err error) {
	return instance.SetProperty("DisplayName", (value))
}

// GetDisplayName gets the value of DisplayName for the instance
func (instance *MDM_EnrollmentStatusTracking_Tracking03_02) GetPropertyDisplayName() (value string, err error) {
	retValue, err := instance.GetProperty("DisplayName")
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

// SetErrorHresult sets the value of ErrorHresult for the instance
func (instance *MDM_EnrollmentStatusTracking_Tracking03_02) SetPropertyErrorHresult(value int32) (err error) {
	return instance.SetProperty("ErrorHresult", (value))
}

// GetErrorHresult gets the value of ErrorHresult for the instance
func (instance *MDM_EnrollmentStatusTracking_Tracking03_02) GetPropertyErrorHresult() (value int32, err error) {
	retValue, err := instance.GetProperty("ErrorHresult")
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

// SetInstallationState sets the value of InstallationState for the instance
func (instance *MDM_EnrollmentStatusTracking_Tracking03_02) SetPropertyInstallationState(value int32) (err error) {
	return instance.SetProperty("InstallationState", (value))
}

// GetInstallationState gets the value of InstallationState for the instance
func (instance *MDM_EnrollmentStatusTracking_Tracking03_02) GetPropertyInstallationState() (value int32, err error) {
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
func (instance *MDM_EnrollmentStatusTracking_Tracking03_02) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", (value))
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_EnrollmentStatusTracking_Tracking03_02) GetPropertyInstanceID() (value string, err error) {
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
func (instance *MDM_EnrollmentStatusTracking_Tracking03_02) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", (value))
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_EnrollmentStatusTracking_Tracking03_02) GetPropertyParentID() (value string, err error) {
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

// SetRebootRequired sets the value of RebootRequired for the instance
func (instance *MDM_EnrollmentStatusTracking_Tracking03_02) SetPropertyRebootRequired(value int32) (err error) {
	return instance.SetProperty("RebootRequired", (value))
}

// GetRebootRequired gets the value of RebootRequired for the instance
func (instance *MDM_EnrollmentStatusTracking_Tracking03_02) GetPropertyRebootRequired() (value int32, err error) {
	retValue, err := instance.GetProperty("RebootRequired")
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

// SetTrackingUri sets the value of TrackingUri for the instance
func (instance *MDM_EnrollmentStatusTracking_Tracking03_02) SetPropertyTrackingUri(value string) (err error) {
	return instance.SetProperty("TrackingUri", (value))
}

// GetTrackingUri gets the value of TrackingUri for the instance
func (instance *MDM_EnrollmentStatusTracking_Tracking03_02) GetPropertyTrackingUri() (value string, err error) {
	retValue, err := instance.GetProperty("TrackingUri")
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
