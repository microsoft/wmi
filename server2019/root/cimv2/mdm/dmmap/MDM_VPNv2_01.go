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

// MDM_VPNv2_01 struct
type MDM_VPNv2_01 struct {
	*cim.WmiInstance

	//
	AlwaysOn bool

	//
	ByPassForLocal bool

	//
	DnsSuffix string

	//
	EdpModeId string

	//
	InstanceID string

	//
	LockDown bool

	//
	ParentID string

	//
	ProfileXML string

	//
	RememberCredentials bool

	//
	TrustedNetworkDetection string
}

func NewMDM_VPNv2_01Ex1(instance *cim.WmiInstance) (newInstance *MDM_VPNv2_01, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_VPNv2_01{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_VPNv2_01Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_VPNv2_01, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_VPNv2_01{
		WmiInstance: tmp,
	}
	return
}

// SetAlwaysOn sets the value of AlwaysOn for the instance
func (instance *MDM_VPNv2_01) SetPropertyAlwaysOn(value bool) (err error) {
	return instance.SetProperty("AlwaysOn", (value))
}

// GetAlwaysOn gets the value of AlwaysOn for the instance
func (instance *MDM_VPNv2_01) GetPropertyAlwaysOn() (value bool, err error) {
	retValue, err := instance.GetProperty("AlwaysOn")
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

// SetByPassForLocal sets the value of ByPassForLocal for the instance
func (instance *MDM_VPNv2_01) SetPropertyByPassForLocal(value bool) (err error) {
	return instance.SetProperty("ByPassForLocal", (value))
}

// GetByPassForLocal gets the value of ByPassForLocal for the instance
func (instance *MDM_VPNv2_01) GetPropertyByPassForLocal() (value bool, err error) {
	retValue, err := instance.GetProperty("ByPassForLocal")
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

// SetDnsSuffix sets the value of DnsSuffix for the instance
func (instance *MDM_VPNv2_01) SetPropertyDnsSuffix(value string) (err error) {
	return instance.SetProperty("DnsSuffix", (value))
}

// GetDnsSuffix gets the value of DnsSuffix for the instance
func (instance *MDM_VPNv2_01) GetPropertyDnsSuffix() (value string, err error) {
	retValue, err := instance.GetProperty("DnsSuffix")
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

// SetEdpModeId sets the value of EdpModeId for the instance
func (instance *MDM_VPNv2_01) SetPropertyEdpModeId(value string) (err error) {
	return instance.SetProperty("EdpModeId", (value))
}

// GetEdpModeId gets the value of EdpModeId for the instance
func (instance *MDM_VPNv2_01) GetPropertyEdpModeId() (value string, err error) {
	retValue, err := instance.GetProperty("EdpModeId")
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
func (instance *MDM_VPNv2_01) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", (value))
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_VPNv2_01) GetPropertyInstanceID() (value string, err error) {
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

// SetLockDown sets the value of LockDown for the instance
func (instance *MDM_VPNv2_01) SetPropertyLockDown(value bool) (err error) {
	return instance.SetProperty("LockDown", (value))
}

// GetLockDown gets the value of LockDown for the instance
func (instance *MDM_VPNv2_01) GetPropertyLockDown() (value bool, err error) {
	retValue, err := instance.GetProperty("LockDown")
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

// SetParentID sets the value of ParentID for the instance
func (instance *MDM_VPNv2_01) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", (value))
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_VPNv2_01) GetPropertyParentID() (value string, err error) {
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

// SetProfileXML sets the value of ProfileXML for the instance
func (instance *MDM_VPNv2_01) SetPropertyProfileXML(value string) (err error) {
	return instance.SetProperty("ProfileXML", (value))
}

// GetProfileXML gets the value of ProfileXML for the instance
func (instance *MDM_VPNv2_01) GetPropertyProfileXML() (value string, err error) {
	retValue, err := instance.GetProperty("ProfileXML")
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

// SetRememberCredentials sets the value of RememberCredentials for the instance
func (instance *MDM_VPNv2_01) SetPropertyRememberCredentials(value bool) (err error) {
	return instance.SetProperty("RememberCredentials", (value))
}

// GetRememberCredentials gets the value of RememberCredentials for the instance
func (instance *MDM_VPNv2_01) GetPropertyRememberCredentials() (value bool, err error) {
	retValue, err := instance.GetProperty("RememberCredentials")
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

// SetTrustedNetworkDetection sets the value of TrustedNetworkDetection for the instance
func (instance *MDM_VPNv2_01) SetPropertyTrustedNetworkDetection(value string) (err error) {
	return instance.SetProperty("TrustedNetworkDetection", (value))
}

// GetTrustedNetworkDetection gets the value of TrustedNetworkDetection for the instance
func (instance *MDM_VPNv2_01) GetPropertyTrustedNetworkDetection() (value string, err error) {
	retValue, err := instance.GetProperty("TrustedNetworkDetection")
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
