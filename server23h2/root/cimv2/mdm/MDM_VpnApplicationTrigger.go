// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.CIMV2.mdm
//////////////////////////////////////////////
package mdm

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MDM_VpnApplicationTrigger struct
type MDM_VpnApplicationTrigger struct {
	*cim.WmiInstance

	//
	ApplicationID string

	//
	TriggerEnabledInAllMDMProfiles bool
}

func NewMDM_VpnApplicationTriggerEx1(instance *cim.WmiInstance) (newInstance *MDM_VpnApplicationTrigger, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_VpnApplicationTrigger{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_VpnApplicationTriggerEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_VpnApplicationTrigger, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_VpnApplicationTrigger{
		WmiInstance: tmp,
	}
	return
}

// SetApplicationID sets the value of ApplicationID for the instance
func (instance *MDM_VpnApplicationTrigger) SetPropertyApplicationID(value string) (err error) {
	return instance.SetProperty("ApplicationID", (value))
}

// GetApplicationID gets the value of ApplicationID for the instance
func (instance *MDM_VpnApplicationTrigger) GetPropertyApplicationID() (value string, err error) {
	retValue, err := instance.GetProperty("ApplicationID")
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

// SetTriggerEnabledInAllMDMProfiles sets the value of TriggerEnabledInAllMDMProfiles for the instance
func (instance *MDM_VpnApplicationTrigger) SetPropertyTriggerEnabledInAllMDMProfiles(value bool) (err error) {
	return instance.SetProperty("TriggerEnabledInAllMDMProfiles", (value))
}

// GetTriggerEnabledInAllMDMProfiles gets the value of TriggerEnabledInAllMDMProfiles for the instance
func (instance *MDM_VpnApplicationTrigger) GetPropertyTriggerEnabledInAllMDMProfiles() (value bool, err error) {
	retValue, err := instance.GetProperty("TriggerEnabledInAllMDMProfiles")
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
