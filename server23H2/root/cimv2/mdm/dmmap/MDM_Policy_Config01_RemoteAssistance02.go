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

// MDM_Policy_Config01_RemoteAssistance02 struct
type MDM_Policy_Config01_RemoteAssistance02 struct {
	*cim.WmiInstance

	//
	CustomizeWarningMessages string

	//
	InstanceID string

	//
	ParentID string

	//
	SessionLogging string

	//
	SolicitedRemoteAssistance string

	//
	UnsolicitedRemoteAssistance string
}

func NewMDM_Policy_Config01_RemoteAssistance02Ex1(instance *cim.WmiInstance) (newInstance *MDM_Policy_Config01_RemoteAssistance02, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_Policy_Config01_RemoteAssistance02{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_Policy_Config01_RemoteAssistance02Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_Policy_Config01_RemoteAssistance02, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_Policy_Config01_RemoteAssistance02{
		WmiInstance: tmp,
	}
	return
}

// SetCustomizeWarningMessages sets the value of CustomizeWarningMessages for the instance
func (instance *MDM_Policy_Config01_RemoteAssistance02) SetPropertyCustomizeWarningMessages(value string) (err error) {
	return instance.SetProperty("CustomizeWarningMessages", (value))
}

// GetCustomizeWarningMessages gets the value of CustomizeWarningMessages for the instance
func (instance *MDM_Policy_Config01_RemoteAssistance02) GetPropertyCustomizeWarningMessages() (value string, err error) {
	retValue, err := instance.GetProperty("CustomizeWarningMessages")
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
func (instance *MDM_Policy_Config01_RemoteAssistance02) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", (value))
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Policy_Config01_RemoteAssistance02) GetPropertyInstanceID() (value string, err error) {
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
func (instance *MDM_Policy_Config01_RemoteAssistance02) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", (value))
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Policy_Config01_RemoteAssistance02) GetPropertyParentID() (value string, err error) {
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

// SetSessionLogging sets the value of SessionLogging for the instance
func (instance *MDM_Policy_Config01_RemoteAssistance02) SetPropertySessionLogging(value string) (err error) {
	return instance.SetProperty("SessionLogging", (value))
}

// GetSessionLogging gets the value of SessionLogging for the instance
func (instance *MDM_Policy_Config01_RemoteAssistance02) GetPropertySessionLogging() (value string, err error) {
	retValue, err := instance.GetProperty("SessionLogging")
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

// SetSolicitedRemoteAssistance sets the value of SolicitedRemoteAssistance for the instance
func (instance *MDM_Policy_Config01_RemoteAssistance02) SetPropertySolicitedRemoteAssistance(value string) (err error) {
	return instance.SetProperty("SolicitedRemoteAssistance", (value))
}

// GetSolicitedRemoteAssistance gets the value of SolicitedRemoteAssistance for the instance
func (instance *MDM_Policy_Config01_RemoteAssistance02) GetPropertySolicitedRemoteAssistance() (value string, err error) {
	retValue, err := instance.GetProperty("SolicitedRemoteAssistance")
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

// SetUnsolicitedRemoteAssistance sets the value of UnsolicitedRemoteAssistance for the instance
func (instance *MDM_Policy_Config01_RemoteAssistance02) SetPropertyUnsolicitedRemoteAssistance(value string) (err error) {
	return instance.SetProperty("UnsolicitedRemoteAssistance", (value))
}

// GetUnsolicitedRemoteAssistance gets the value of UnsolicitedRemoteAssistance for the instance
func (instance *MDM_Policy_Config01_RemoteAssistance02) GetPropertyUnsolicitedRemoteAssistance() (value string, err error) {
	retValue, err := instance.GetProperty("UnsolicitedRemoteAssistance")
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
