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

// MDM_Policy_Result01_TimeLanguageSettings02 struct
type MDM_Policy_Result01_TimeLanguageSettings02 struct {
	*cim.WmiInstance

	//
	BlockCleanupOfUnusedPreinstalledLangPacks int32

	//
	ConfigureTimeZone string

	//
	InstanceID string

	//
	MachineUILanguageOverwrite int32

	//
	ParentID string

	//
	RestrictLanguagePacksAndFeaturesInstall int32
}

func NewMDM_Policy_Result01_TimeLanguageSettings02Ex1(instance *cim.WmiInstance) (newInstance *MDM_Policy_Result01_TimeLanguageSettings02, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_Policy_Result01_TimeLanguageSettings02{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_Policy_Result01_TimeLanguageSettings02Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_Policy_Result01_TimeLanguageSettings02, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_Policy_Result01_TimeLanguageSettings02{
		WmiInstance: tmp,
	}
	return
}

// SetBlockCleanupOfUnusedPreinstalledLangPacks sets the value of BlockCleanupOfUnusedPreinstalledLangPacks for the instance
func (instance *MDM_Policy_Result01_TimeLanguageSettings02) SetPropertyBlockCleanupOfUnusedPreinstalledLangPacks(value int32) (err error) {
	return instance.SetProperty("BlockCleanupOfUnusedPreinstalledLangPacks", (value))
}

// GetBlockCleanupOfUnusedPreinstalledLangPacks gets the value of BlockCleanupOfUnusedPreinstalledLangPacks for the instance
func (instance *MDM_Policy_Result01_TimeLanguageSettings02) GetPropertyBlockCleanupOfUnusedPreinstalledLangPacks() (value int32, err error) {
	retValue, err := instance.GetProperty("BlockCleanupOfUnusedPreinstalledLangPacks")
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

// SetConfigureTimeZone sets the value of ConfigureTimeZone for the instance
func (instance *MDM_Policy_Result01_TimeLanguageSettings02) SetPropertyConfigureTimeZone(value string) (err error) {
	return instance.SetProperty("ConfigureTimeZone", (value))
}

// GetConfigureTimeZone gets the value of ConfigureTimeZone for the instance
func (instance *MDM_Policy_Result01_TimeLanguageSettings02) GetPropertyConfigureTimeZone() (value string, err error) {
	retValue, err := instance.GetProperty("ConfigureTimeZone")
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
func (instance *MDM_Policy_Result01_TimeLanguageSettings02) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", (value))
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Policy_Result01_TimeLanguageSettings02) GetPropertyInstanceID() (value string, err error) {
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

// SetMachineUILanguageOverwrite sets the value of MachineUILanguageOverwrite for the instance
func (instance *MDM_Policy_Result01_TimeLanguageSettings02) SetPropertyMachineUILanguageOverwrite(value int32) (err error) {
	return instance.SetProperty("MachineUILanguageOverwrite", (value))
}

// GetMachineUILanguageOverwrite gets the value of MachineUILanguageOverwrite for the instance
func (instance *MDM_Policy_Result01_TimeLanguageSettings02) GetPropertyMachineUILanguageOverwrite() (value int32, err error) {
	retValue, err := instance.GetProperty("MachineUILanguageOverwrite")
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
func (instance *MDM_Policy_Result01_TimeLanguageSettings02) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", (value))
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Policy_Result01_TimeLanguageSettings02) GetPropertyParentID() (value string, err error) {
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

// SetRestrictLanguagePacksAndFeaturesInstall sets the value of RestrictLanguagePacksAndFeaturesInstall for the instance
func (instance *MDM_Policy_Result01_TimeLanguageSettings02) SetPropertyRestrictLanguagePacksAndFeaturesInstall(value int32) (err error) {
	return instance.SetProperty("RestrictLanguagePacksAndFeaturesInstall", (value))
}

// GetRestrictLanguagePacksAndFeaturesInstall gets the value of RestrictLanguagePacksAndFeaturesInstall for the instance
func (instance *MDM_Policy_Result01_TimeLanguageSettings02) GetPropertyRestrictLanguagePacksAndFeaturesInstall() (value int32, err error) {
	retValue, err := instance.GetProperty("RestrictLanguagePacksAndFeaturesInstall")
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
