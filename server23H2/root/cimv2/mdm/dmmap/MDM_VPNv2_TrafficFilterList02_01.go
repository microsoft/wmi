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

// MDM_VPNv2_TrafficFilterList02_01 struct
type MDM_VPNv2_TrafficFilterList02_01 struct {
	*cim.WmiInstance

	//
	Claims string

	//
	Direction string

	//
	InstanceID string

	//
	LocalAddressRanges string

	//
	LocalPortRanges string

	//
	ParentID string

	//
	Protocol int32

	//
	RemoteAddressRanges string

	//
	RemotePortRanges string

	//
	RoutingPolicyType string
}

func NewMDM_VPNv2_TrafficFilterList02_01Ex1(instance *cim.WmiInstance) (newInstance *MDM_VPNv2_TrafficFilterList02_01, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_VPNv2_TrafficFilterList02_01{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_VPNv2_TrafficFilterList02_01Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_VPNv2_TrafficFilterList02_01, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_VPNv2_TrafficFilterList02_01{
		WmiInstance: tmp,
	}
	return
}

// SetClaims sets the value of Claims for the instance
func (instance *MDM_VPNv2_TrafficFilterList02_01) SetPropertyClaims(value string) (err error) {
	return instance.SetProperty("Claims", (value))
}

// GetClaims gets the value of Claims for the instance
func (instance *MDM_VPNv2_TrafficFilterList02_01) GetPropertyClaims() (value string, err error) {
	retValue, err := instance.GetProperty("Claims")
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

// SetDirection sets the value of Direction for the instance
func (instance *MDM_VPNv2_TrafficFilterList02_01) SetPropertyDirection(value string) (err error) {
	return instance.SetProperty("Direction", (value))
}

// GetDirection gets the value of Direction for the instance
func (instance *MDM_VPNv2_TrafficFilterList02_01) GetPropertyDirection() (value string, err error) {
	retValue, err := instance.GetProperty("Direction")
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
func (instance *MDM_VPNv2_TrafficFilterList02_01) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", (value))
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_VPNv2_TrafficFilterList02_01) GetPropertyInstanceID() (value string, err error) {
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

// SetLocalAddressRanges sets the value of LocalAddressRanges for the instance
func (instance *MDM_VPNv2_TrafficFilterList02_01) SetPropertyLocalAddressRanges(value string) (err error) {
	return instance.SetProperty("LocalAddressRanges", (value))
}

// GetLocalAddressRanges gets the value of LocalAddressRanges for the instance
func (instance *MDM_VPNv2_TrafficFilterList02_01) GetPropertyLocalAddressRanges() (value string, err error) {
	retValue, err := instance.GetProperty("LocalAddressRanges")
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

// SetLocalPortRanges sets the value of LocalPortRanges for the instance
func (instance *MDM_VPNv2_TrafficFilterList02_01) SetPropertyLocalPortRanges(value string) (err error) {
	return instance.SetProperty("LocalPortRanges", (value))
}

// GetLocalPortRanges gets the value of LocalPortRanges for the instance
func (instance *MDM_VPNv2_TrafficFilterList02_01) GetPropertyLocalPortRanges() (value string, err error) {
	retValue, err := instance.GetProperty("LocalPortRanges")
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
func (instance *MDM_VPNv2_TrafficFilterList02_01) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", (value))
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_VPNv2_TrafficFilterList02_01) GetPropertyParentID() (value string, err error) {
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

// SetProtocol sets the value of Protocol for the instance
func (instance *MDM_VPNv2_TrafficFilterList02_01) SetPropertyProtocol(value int32) (err error) {
	return instance.SetProperty("Protocol", (value))
}

// GetProtocol gets the value of Protocol for the instance
func (instance *MDM_VPNv2_TrafficFilterList02_01) GetPropertyProtocol() (value int32, err error) {
	retValue, err := instance.GetProperty("Protocol")
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

// SetRemoteAddressRanges sets the value of RemoteAddressRanges for the instance
func (instance *MDM_VPNv2_TrafficFilterList02_01) SetPropertyRemoteAddressRanges(value string) (err error) {
	return instance.SetProperty("RemoteAddressRanges", (value))
}

// GetRemoteAddressRanges gets the value of RemoteAddressRanges for the instance
func (instance *MDM_VPNv2_TrafficFilterList02_01) GetPropertyRemoteAddressRanges() (value string, err error) {
	retValue, err := instance.GetProperty("RemoteAddressRanges")
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

// SetRemotePortRanges sets the value of RemotePortRanges for the instance
func (instance *MDM_VPNv2_TrafficFilterList02_01) SetPropertyRemotePortRanges(value string) (err error) {
	return instance.SetProperty("RemotePortRanges", (value))
}

// GetRemotePortRanges gets the value of RemotePortRanges for the instance
func (instance *MDM_VPNv2_TrafficFilterList02_01) GetPropertyRemotePortRanges() (value string, err error) {
	retValue, err := instance.GetProperty("RemotePortRanges")
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

// SetRoutingPolicyType sets the value of RoutingPolicyType for the instance
func (instance *MDM_VPNv2_TrafficFilterList02_01) SetPropertyRoutingPolicyType(value string) (err error) {
	return instance.SetProperty("RoutingPolicyType", (value))
}

// GetRoutingPolicyType gets the value of RoutingPolicyType for the instance
func (instance *MDM_VPNv2_TrafficFilterList02_01) GetPropertyRoutingPolicyType() (value string, err error) {
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
