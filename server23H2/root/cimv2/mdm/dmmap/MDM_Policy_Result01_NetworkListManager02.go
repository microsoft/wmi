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

// MDM_Policy_Result01_NetworkListManager02 struct
type MDM_Policy_Result01_NetworkListManager02 struct {
	*cim.WmiInstance

	//
	AllowedTlsAuthenticationEndpoints string

	//
	InstanceID string

	//
	ParentID string
}

func NewMDM_Policy_Result01_NetworkListManager02Ex1(instance *cim.WmiInstance) (newInstance *MDM_Policy_Result01_NetworkListManager02, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_Policy_Result01_NetworkListManager02{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_Policy_Result01_NetworkListManager02Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_Policy_Result01_NetworkListManager02, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_Policy_Result01_NetworkListManager02{
		WmiInstance: tmp,
	}
	return
}

// SetAllowedTlsAuthenticationEndpoints sets the value of AllowedTlsAuthenticationEndpoints for the instance
func (instance *MDM_Policy_Result01_NetworkListManager02) SetPropertyAllowedTlsAuthenticationEndpoints(value string) (err error) {
	return instance.SetProperty("AllowedTlsAuthenticationEndpoints", (value))
}

// GetAllowedTlsAuthenticationEndpoints gets the value of AllowedTlsAuthenticationEndpoints for the instance
func (instance *MDM_Policy_Result01_NetworkListManager02) GetPropertyAllowedTlsAuthenticationEndpoints() (value string, err error) {
	retValue, err := instance.GetProperty("AllowedTlsAuthenticationEndpoints")
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
func (instance *MDM_Policy_Result01_NetworkListManager02) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", (value))
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Policy_Result01_NetworkListManager02) GetPropertyInstanceID() (value string, err error) {
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
func (instance *MDM_Policy_Result01_NetworkListManager02) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", (value))
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Policy_Result01_NetworkListManager02) GetPropertyParentID() (value string, err error) {
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