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

// MDM_VPNv2_APNBinding02 struct
type MDM_VPNv2_APNBinding02 struct {
	*cim.WmiInstance

	//
	AccessPointName string

	//
	AuthenticationType string

	//
	InstanceID string

	//
	IsCompressionEnabled bool

	//
	ParentID string

	//
	Password string

	//
	ProviderId string

	//
	UserName string
}

func NewMDM_VPNv2_APNBinding02Ex1(instance *cim.WmiInstance) (newInstance *MDM_VPNv2_APNBinding02, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_VPNv2_APNBinding02{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_VPNv2_APNBinding02Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_VPNv2_APNBinding02, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_VPNv2_APNBinding02{
		WmiInstance: tmp,
	}
	return
}

// SetAccessPointName sets the value of AccessPointName for the instance
func (instance *MDM_VPNv2_APNBinding02) SetPropertyAccessPointName(value string) (err error) {
	return instance.SetProperty("AccessPointName", (value))
}

// GetAccessPointName gets the value of AccessPointName for the instance
func (instance *MDM_VPNv2_APNBinding02) GetPropertyAccessPointName() (value string, err error) {
	retValue, err := instance.GetProperty("AccessPointName")
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

// SetAuthenticationType sets the value of AuthenticationType for the instance
func (instance *MDM_VPNv2_APNBinding02) SetPropertyAuthenticationType(value string) (err error) {
	return instance.SetProperty("AuthenticationType", (value))
}

// GetAuthenticationType gets the value of AuthenticationType for the instance
func (instance *MDM_VPNv2_APNBinding02) GetPropertyAuthenticationType() (value string, err error) {
	retValue, err := instance.GetProperty("AuthenticationType")
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
func (instance *MDM_VPNv2_APNBinding02) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", (value))
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_VPNv2_APNBinding02) GetPropertyInstanceID() (value string, err error) {
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

// SetIsCompressionEnabled sets the value of IsCompressionEnabled for the instance
func (instance *MDM_VPNv2_APNBinding02) SetPropertyIsCompressionEnabled(value bool) (err error) {
	return instance.SetProperty("IsCompressionEnabled", (value))
}

// GetIsCompressionEnabled gets the value of IsCompressionEnabled for the instance
func (instance *MDM_VPNv2_APNBinding02) GetPropertyIsCompressionEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("IsCompressionEnabled")
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
func (instance *MDM_VPNv2_APNBinding02) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", (value))
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_VPNv2_APNBinding02) GetPropertyParentID() (value string, err error) {
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

// SetPassword sets the value of Password for the instance
func (instance *MDM_VPNv2_APNBinding02) SetPropertyPassword(value string) (err error) {
	return instance.SetProperty("Password", (value))
}

// GetPassword gets the value of Password for the instance
func (instance *MDM_VPNv2_APNBinding02) GetPropertyPassword() (value string, err error) {
	retValue, err := instance.GetProperty("Password")
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

// SetProviderId sets the value of ProviderId for the instance
func (instance *MDM_VPNv2_APNBinding02) SetPropertyProviderId(value string) (err error) {
	return instance.SetProperty("ProviderId", (value))
}

// GetProviderId gets the value of ProviderId for the instance
func (instance *MDM_VPNv2_APNBinding02) GetPropertyProviderId() (value string, err error) {
	retValue, err := instance.GetProperty("ProviderId")
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

// SetUserName sets the value of UserName for the instance
func (instance *MDM_VPNv2_APNBinding02) SetPropertyUserName(value string) (err error) {
	return instance.SetProperty("UserName", (value))
}

// GetUserName gets the value of UserName for the instance
func (instance *MDM_VPNv2_APNBinding02) GetPropertyUserName() (value string, err error) {
	retValue, err := instance.GetProperty("UserName")
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
