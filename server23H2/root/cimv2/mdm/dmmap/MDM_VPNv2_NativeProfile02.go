// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
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

// MDM_VPNv2_NativeProfile02 struct
type MDM_VPNv2_NativeProfile02 struct {
	*cim.WmiInstance

	//
	DisableClassBasedDefaultRoute bool

	//
	InstanceID string

	//
	L2tpPsk string

	//
	NativeProtocolType string

	//
	ParentID string

	//
	PlumbIKEv2TSAsRoutes bool

	//
	RoutingPolicyType string

	//
	Servers string
}

func NewMDM_VPNv2_NativeProfile02Ex1(instance *cim.WmiInstance) (newInstance *MDM_VPNv2_NativeProfile02, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_VPNv2_NativeProfile02{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_VPNv2_NativeProfile02Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_VPNv2_NativeProfile02, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_VPNv2_NativeProfile02{
		WmiInstance: tmp,
	}
	return
}

// SetDisableClassBasedDefaultRoute sets the value of DisableClassBasedDefaultRoute for the instance
func (instance *MDM_VPNv2_NativeProfile02) SetPropertyDisableClassBasedDefaultRoute(value bool) (err error) {
	return instance.SetProperty("DisableClassBasedDefaultRoute", (value))
}

// GetDisableClassBasedDefaultRoute gets the value of DisableClassBasedDefaultRoute for the instance
func (instance *MDM_VPNv2_NativeProfile02) GetPropertyDisableClassBasedDefaultRoute() (value bool, err error) {
	retValue, err := instance.GetProperty("DisableClassBasedDefaultRoute")
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
func (instance *MDM_VPNv2_NativeProfile02) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", (value))
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_VPNv2_NativeProfile02) GetPropertyInstanceID() (value string, err error) {
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

// SetL2tpPsk sets the value of L2tpPsk for the instance
func (instance *MDM_VPNv2_NativeProfile02) SetPropertyL2tpPsk(value string) (err error) {
	return instance.SetProperty("L2tpPsk", (value))
}

// GetL2tpPsk gets the value of L2tpPsk for the instance
func (instance *MDM_VPNv2_NativeProfile02) GetPropertyL2tpPsk() (value string, err error) {
	retValue, err := instance.GetProperty("L2tpPsk")
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

// SetNativeProtocolType sets the value of NativeProtocolType for the instance
func (instance *MDM_VPNv2_NativeProfile02) SetPropertyNativeProtocolType(value string) (err error) {
	return instance.SetProperty("NativeProtocolType", (value))
}

// GetNativeProtocolType gets the value of NativeProtocolType for the instance
func (instance *MDM_VPNv2_NativeProfile02) GetPropertyNativeProtocolType() (value string, err error) {
	retValue, err := instance.GetProperty("NativeProtocolType")
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
func (instance *MDM_VPNv2_NativeProfile02) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", (value))
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_VPNv2_NativeProfile02) GetPropertyParentID() (value string, err error) {
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

// SetPlumbIKEv2TSAsRoutes sets the value of PlumbIKEv2TSAsRoutes for the instance
func (instance *MDM_VPNv2_NativeProfile02) SetPropertyPlumbIKEv2TSAsRoutes(value bool) (err error) {
	return instance.SetProperty("PlumbIKEv2TSAsRoutes", (value))
}

// GetPlumbIKEv2TSAsRoutes gets the value of PlumbIKEv2TSAsRoutes for the instance
func (instance *MDM_VPNv2_NativeProfile02) GetPropertyPlumbIKEv2TSAsRoutes() (value bool, err error) {
	retValue, err := instance.GetProperty("PlumbIKEv2TSAsRoutes")
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

// SetRoutingPolicyType sets the value of RoutingPolicyType for the instance
func (instance *MDM_VPNv2_NativeProfile02) SetPropertyRoutingPolicyType(value string) (err error) {
	return instance.SetProperty("RoutingPolicyType", (value))
}

// GetRoutingPolicyType gets the value of RoutingPolicyType for the instance
func (instance *MDM_VPNv2_NativeProfile02) GetPropertyRoutingPolicyType() (value string, err error) {
	retValue, err := instance.GetProperty("RoutingPolicyType")
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

// SetServers sets the value of Servers for the instance
func (instance *MDM_VPNv2_NativeProfile02) SetPropertyServers(value string) (err error) {
	return instance.SetProperty("Servers", (value))
}

// GetServers gets the value of Servers for the instance
func (instance *MDM_VPNv2_NativeProfile02) GetPropertyServers() (value string, err error) {
	retValue, err := instance.GetProperty("Servers")
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
