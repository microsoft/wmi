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

// MDM_EnrollmentStatusTracking_TrackedResourceTypes04 struct
type MDM_EnrollmentStatusTracking_TrackedResourceTypes04 struct {
	*cim.WmiInstance

	//
	Apps bool

	//
	InstanceID string

	//
	ParentID string
}

func NewMDM_EnrollmentStatusTracking_TrackedResourceTypes04Ex1(instance *cim.WmiInstance) (newInstance *MDM_EnrollmentStatusTracking_TrackedResourceTypes04, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_EnrollmentStatusTracking_TrackedResourceTypes04{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_EnrollmentStatusTracking_TrackedResourceTypes04Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_EnrollmentStatusTracking_TrackedResourceTypes04, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_EnrollmentStatusTracking_TrackedResourceTypes04{
		WmiInstance: tmp,
	}
	return
}

// SetApps sets the value of Apps for the instance
func (instance *MDM_EnrollmentStatusTracking_TrackedResourceTypes04) SetPropertyApps(value bool) (err error) {
	return instance.SetProperty("Apps", (value))
}

// GetApps gets the value of Apps for the instance
func (instance *MDM_EnrollmentStatusTracking_TrackedResourceTypes04) GetPropertyApps() (value bool, err error) {
	retValue, err := instance.GetProperty("Apps")
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

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_EnrollmentStatusTracking_TrackedResourceTypes04) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", (value))
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_EnrollmentStatusTracking_TrackedResourceTypes04) GetPropertyInstanceID() (value string, err error) {
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
func (instance *MDM_EnrollmentStatusTracking_TrackedResourceTypes04) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", (value))
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_EnrollmentStatusTracking_TrackedResourceTypes04) GetPropertyParentID() (value string, err error) {
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
