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

// MDM_Policy_Result01_Autoplay02 struct
type MDM_Policy_Result01_Autoplay02 struct {
	*cim.WmiInstance

	//
	DisallowAutoplayForNonVolumeDevices string

	//
	InstanceID string

	//
	ParentID string

	//
	SetDefaultAutoRunBehavior string

	//
	TurnOffAutoPlay string
}

func NewMDM_Policy_Result01_Autoplay02Ex1(instance *cim.WmiInstance) (newInstance *MDM_Policy_Result01_Autoplay02, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_Policy_Result01_Autoplay02{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_Policy_Result01_Autoplay02Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_Policy_Result01_Autoplay02, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_Policy_Result01_Autoplay02{
		WmiInstance: tmp,
	}
	return
}

// SetDisallowAutoplayForNonVolumeDevices sets the value of DisallowAutoplayForNonVolumeDevices for the instance
func (instance *MDM_Policy_Result01_Autoplay02) SetPropertyDisallowAutoplayForNonVolumeDevices(value string) (err error) {
	return instance.SetProperty("DisallowAutoplayForNonVolumeDevices", (value))
}

// GetDisallowAutoplayForNonVolumeDevices gets the value of DisallowAutoplayForNonVolumeDevices for the instance
func (instance *MDM_Policy_Result01_Autoplay02) GetPropertyDisallowAutoplayForNonVolumeDevices() (value string, err error) {
	retValue, err := instance.GetProperty("DisallowAutoplayForNonVolumeDevices")
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
func (instance *MDM_Policy_Result01_Autoplay02) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", (value))
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Policy_Result01_Autoplay02) GetPropertyInstanceID() (value string, err error) {
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
func (instance *MDM_Policy_Result01_Autoplay02) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", (value))
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Policy_Result01_Autoplay02) GetPropertyParentID() (value string, err error) {
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

// SetSetDefaultAutoRunBehavior sets the value of SetDefaultAutoRunBehavior for the instance
func (instance *MDM_Policy_Result01_Autoplay02) SetPropertySetDefaultAutoRunBehavior(value string) (err error) {
	return instance.SetProperty("SetDefaultAutoRunBehavior", (value))
}

// GetSetDefaultAutoRunBehavior gets the value of SetDefaultAutoRunBehavior for the instance
func (instance *MDM_Policy_Result01_Autoplay02) GetPropertySetDefaultAutoRunBehavior() (value string, err error) {
	retValue, err := instance.GetProperty("SetDefaultAutoRunBehavior")
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

// SetTurnOffAutoPlay sets the value of TurnOffAutoPlay for the instance
func (instance *MDM_Policy_Result01_Autoplay02) SetPropertyTurnOffAutoPlay(value string) (err error) {
	return instance.SetProperty("TurnOffAutoPlay", (value))
}

// GetTurnOffAutoPlay gets the value of TurnOffAutoPlay for the instance
func (instance *MDM_Policy_Result01_Autoplay02) GetPropertyTurnOffAutoPlay() (value string, err error) {
	retValue, err := instance.GetProperty("TurnOffAutoPlay")
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
