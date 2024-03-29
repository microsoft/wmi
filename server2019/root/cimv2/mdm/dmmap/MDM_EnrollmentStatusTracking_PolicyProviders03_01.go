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

// MDM_EnrollmentStatusTracking_PolicyProviders03_01 struct
type MDM_EnrollmentStatusTracking_PolicyProviders03_01 struct {
	*cim.WmiInstance

	//
	InstanceID string

	//
	ParentID string

	//
	TrackingPoliciesCreated bool
}

func NewMDM_EnrollmentStatusTracking_PolicyProviders03_01Ex1(instance *cim.WmiInstance) (newInstance *MDM_EnrollmentStatusTracking_PolicyProviders03_01, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_EnrollmentStatusTracking_PolicyProviders03_01{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_EnrollmentStatusTracking_PolicyProviders03_01Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_EnrollmentStatusTracking_PolicyProviders03_01, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_EnrollmentStatusTracking_PolicyProviders03_01{
		WmiInstance: tmp,
	}
	return
}

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_EnrollmentStatusTracking_PolicyProviders03_01) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_EnrollmentStatusTracking_PolicyProviders03_01) GetPropertyInstanceID() (value string, err error) {
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
func (instance *MDM_EnrollmentStatusTracking_PolicyProviders03_01) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_EnrollmentStatusTracking_PolicyProviders03_01) GetPropertyParentID() (value string, err error) {
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

// SetTrackingPoliciesCreated sets the value of TrackingPoliciesCreated for the instance
func (instance *MDM_EnrollmentStatusTracking_PolicyProviders03_01) SetPropertyTrackingPoliciesCreated(value bool) (err error) {
	return instance.SetProperty("TrackingPoliciesCreated", value)
}

// GetTrackingPoliciesCreated gets the value of TrackingPoliciesCreated for the instance
func (instance *MDM_EnrollmentStatusTracking_PolicyProviders03_01) GetPropertyTrackingPoliciesCreated() (value bool, err error) {
	retValue, err := instance.GetProperty("TrackingPoliciesCreated")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}
