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

// MDM_Policy_Config01_Accounts02 struct
type MDM_Policy_Config01_Accounts02 struct {
	*cim.WmiInstance

	//
	AllowAddingNonMicrosoftAccountsManually int32

	//
	AllowMicrosoftAccountConnection int32

	//
	AllowMicrosoftAccountSignInAssistant int32

	//
	DomainNamesForEmailSync string

	//
	InstanceID string

	//
	ParentID string
}

func NewMDM_Policy_Config01_Accounts02Ex1(instance *cim.WmiInstance) (newInstance *MDM_Policy_Config01_Accounts02, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_Policy_Config01_Accounts02{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_Policy_Config01_Accounts02Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_Policy_Config01_Accounts02, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_Policy_Config01_Accounts02{
		WmiInstance: tmp,
	}
	return
}

// SetAllowAddingNonMicrosoftAccountsManually sets the value of AllowAddingNonMicrosoftAccountsManually for the instance
func (instance *MDM_Policy_Config01_Accounts02) SetPropertyAllowAddingNonMicrosoftAccountsManually(value int32) (err error) {
	return instance.SetProperty("AllowAddingNonMicrosoftAccountsManually", (value))
}

// GetAllowAddingNonMicrosoftAccountsManually gets the value of AllowAddingNonMicrosoftAccountsManually for the instance
func (instance *MDM_Policy_Config01_Accounts02) GetPropertyAllowAddingNonMicrosoftAccountsManually() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowAddingNonMicrosoftAccountsManually")
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

// SetAllowMicrosoftAccountConnection sets the value of AllowMicrosoftAccountConnection for the instance
func (instance *MDM_Policy_Config01_Accounts02) SetPropertyAllowMicrosoftAccountConnection(value int32) (err error) {
	return instance.SetProperty("AllowMicrosoftAccountConnection", (value))
}

// GetAllowMicrosoftAccountConnection gets the value of AllowMicrosoftAccountConnection for the instance
func (instance *MDM_Policy_Config01_Accounts02) GetPropertyAllowMicrosoftAccountConnection() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowMicrosoftAccountConnection")
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

// SetAllowMicrosoftAccountSignInAssistant sets the value of AllowMicrosoftAccountSignInAssistant for the instance
func (instance *MDM_Policy_Config01_Accounts02) SetPropertyAllowMicrosoftAccountSignInAssistant(value int32) (err error) {
	return instance.SetProperty("AllowMicrosoftAccountSignInAssistant", (value))
}

// GetAllowMicrosoftAccountSignInAssistant gets the value of AllowMicrosoftAccountSignInAssistant for the instance
func (instance *MDM_Policy_Config01_Accounts02) GetPropertyAllowMicrosoftAccountSignInAssistant() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowMicrosoftAccountSignInAssistant")
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

// SetDomainNamesForEmailSync sets the value of DomainNamesForEmailSync for the instance
func (instance *MDM_Policy_Config01_Accounts02) SetPropertyDomainNamesForEmailSync(value string) (err error) {
	return instance.SetProperty("DomainNamesForEmailSync", (value))
}

// GetDomainNamesForEmailSync gets the value of DomainNamesForEmailSync for the instance
func (instance *MDM_Policy_Config01_Accounts02) GetPropertyDomainNamesForEmailSync() (value string, err error) {
	retValue, err := instance.GetProperty("DomainNamesForEmailSync")
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
func (instance *MDM_Policy_Config01_Accounts02) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", (value))
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Policy_Config01_Accounts02) GetPropertyInstanceID() (value string, err error) {
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
func (instance *MDM_Policy_Config01_Accounts02) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", (value))
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Policy_Config01_Accounts02) GetPropertyParentID() (value string, err error) {
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
