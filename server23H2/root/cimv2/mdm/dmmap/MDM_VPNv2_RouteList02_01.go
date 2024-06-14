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

// MDM_VPNv2_RouteList02_01 struct
type MDM_VPNv2_RouteList02_01 struct {
	*cim.WmiInstance

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

func NewMDM_VPNv2_RouteList02_01Ex1(instance *cim.WmiInstance) (newInstance *MDM_VPNv2_RouteList02_01, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_VPNv2_RouteList02_01{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_VPNv2_RouteList02_01Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_VPNv2_RouteList02_01, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_VPNv2_RouteList02_01{
		WmiInstance: tmp,
	}
	return
}

// SetAddress sets the value of Address for the instance
func (instance *MDM_VPNv2_RouteList02_01) SetPropertyAddress(value string) (err error) {
	return instance.SetProperty("Address", (value))
}

// GetAddress gets the value of Address for the instance
func (instance *MDM_VPNv2_RouteList02_01) GetPropertyAddress() (value string, err error) {
	retValue, err := instance.GetProperty("Address")
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

// SetExclusionRoute sets the value of ExclusionRoute for the instance
func (instance *MDM_VPNv2_RouteList02_01) SetPropertyExclusionRoute(value bool) (err error) {
	return instance.SetProperty("ExclusionRoute", (value))
}

// GetExclusionRoute gets the value of ExclusionRoute for the instance
func (instance *MDM_VPNv2_RouteList02_01) GetPropertyExclusionRoute() (value bool, err error) {
	retValue, err := instance.GetProperty("ExclusionRoute")
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
func (instance *MDM_VPNv2_RouteList02_01) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", (value))
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_VPNv2_RouteList02_01) GetPropertyInstanceID() (value string, err error) {
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

// SetMetric sets the value of Metric for the instance
func (instance *MDM_VPNv2_RouteList02_01) SetPropertyMetric(value int32) (err error) {
	return instance.SetProperty("Metric", (value))
}

// GetMetric gets the value of Metric for the instance
func (instance *MDM_VPNv2_RouteList02_01) GetPropertyMetric() (value int32, err error) {
	retValue, err := instance.GetProperty("Metric")
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
func (instance *MDM_VPNv2_RouteList02_01) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", (value))
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_VPNv2_RouteList02_01) GetPropertyParentID() (value string, err error) {
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

// SetPrefixSize sets the value of PrefixSize for the instance
func (instance *MDM_VPNv2_RouteList02_01) SetPropertyPrefixSize(value int32) (err error) {
	return instance.SetProperty("PrefixSize", (value))
}

// GetPrefixSize gets the value of PrefixSize for the instance
func (instance *MDM_VPNv2_RouteList02_01) GetPropertyPrefixSize() (value int32, err error) {
	retValue, err := instance.GetProperty("PrefixSize")
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
