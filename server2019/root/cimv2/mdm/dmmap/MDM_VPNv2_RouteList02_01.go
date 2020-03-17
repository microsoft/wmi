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

// MDM_VPNv2_RouteList02_01 struct
type MDM_VPNv2_RouteList02_01 struct {
	cim.WmiInstance

	//
	Address string

	//
	ExclusionRoute bool

	//
	InstanceID string

	//
	Metric int32

	//
	ParentID string

	//
	PrefixSize int32
}

// SetAddress sets the value of Address for the instance
func (instance *MDM_VPNv2_RouteList02_01) SetPropertyAddress(value string) (err error) {
	return instance.SetProperty("Address", value)
}

// GetAddress gets the value of Address for the instance
func (instance *MDM_VPNv2_RouteList02_01) GetPropertyAddress() (value string, err error) {
	retValue, err := instance.GetProperty("Address")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetExclusionRoute sets the value of ExclusionRoute for the instance
func (instance *MDM_VPNv2_RouteList02_01) SetPropertyExclusionRoute(value bool) (err error) {
	return instance.SetProperty("ExclusionRoute", value)
}

// GetExclusionRoute gets the value of ExclusionRoute for the instance
func (instance *MDM_VPNv2_RouteList02_01) GetPropertyExclusionRoute() (value bool, err error) {
	retValue, err := instance.GetProperty("ExclusionRoute")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_VPNv2_RouteList02_01) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_VPNv2_RouteList02_01) GetPropertyInstanceID() (value string, err error) {
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

// SetMetric sets the value of Metric for the instance
func (instance *MDM_VPNv2_RouteList02_01) SetPropertyMetric(value int32) (err error) {
	return instance.SetProperty("Metric", value)
}

// GetMetric gets the value of Metric for the instance
func (instance *MDM_VPNv2_RouteList02_01) GetPropertyMetric() (value int32, err error) {
	retValue, err := instance.GetProperty("Metric")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetParentID sets the value of ParentID for the instance
func (instance *MDM_VPNv2_RouteList02_01) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_VPNv2_RouteList02_01) GetPropertyParentID() (value string, err error) {
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

// SetPrefixSize sets the value of PrefixSize for the instance
func (instance *MDM_VPNv2_RouteList02_01) SetPropertyPrefixSize(value int32) (err error) {
	return instance.SetProperty("PrefixSize", value)
}

// GetPrefixSize gets the value of PrefixSize for the instance
func (instance *MDM_VPNv2_RouteList02_01) GetPropertyPrefixSize() (value int32, err error) {
	retValue, err := instance.GetProperty("PrefixSize")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}
