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

// MDM_MultiSIM_Slots02_01 struct
type MDM_MultiSIM_Slots02_01 struct {
	*cim.WmiInstance

	//
	Identifier int32

	//
	InstanceID string

	//
	IsEmbedded bool

	//
	IsSelected bool

	//
	ParentID string

	//
	State int32
}

func NewMDM_MultiSIM_Slots02_01Ex1(instance *cim.WmiInstance) (newInstance *MDM_MultiSIM_Slots02_01, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_MultiSIM_Slots02_01{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_MultiSIM_Slots02_01Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_MultiSIM_Slots02_01, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_MultiSIM_Slots02_01{
		WmiInstance: tmp,
	}
	return
}

// SetIdentifier sets the value of Identifier for the instance
func (instance *MDM_MultiSIM_Slots02_01) SetPropertyIdentifier(value int32) (err error) {
	return instance.SetProperty("Identifier", (value))
}

// GetIdentifier gets the value of Identifier for the instance
func (instance *MDM_MultiSIM_Slots02_01) GetPropertyIdentifier() (value int32, err error) {
	retValue, err := instance.GetProperty("Identifier")
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

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_MultiSIM_Slots02_01) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", (value))
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_MultiSIM_Slots02_01) GetPropertyInstanceID() (value string, err error) {
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

// SetIsEmbedded sets the value of IsEmbedded for the instance
func (instance *MDM_MultiSIM_Slots02_01) SetPropertyIsEmbedded(value bool) (err error) {
	return instance.SetProperty("IsEmbedded", (value))
}

// GetIsEmbedded gets the value of IsEmbedded for the instance
func (instance *MDM_MultiSIM_Slots02_01) GetPropertyIsEmbedded() (value bool, err error) {
	retValue, err := instance.GetProperty("IsEmbedded")
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

// SetIsSelected sets the value of IsSelected for the instance
func (instance *MDM_MultiSIM_Slots02_01) SetPropertyIsSelected(value bool) (err error) {
	return instance.SetProperty("IsSelected", (value))
}

// GetIsSelected gets the value of IsSelected for the instance
func (instance *MDM_MultiSIM_Slots02_01) GetPropertyIsSelected() (value bool, err error) {
	retValue, err := instance.GetProperty("IsSelected")
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
func (instance *MDM_MultiSIM_Slots02_01) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", (value))
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_MultiSIM_Slots02_01) GetPropertyParentID() (value string, err error) {
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

// SetState sets the value of State for the instance
func (instance *MDM_MultiSIM_Slots02_01) SetPropertyState(value int32) (err error) {
	return instance.SetProperty("State", (value))
}

// GetState gets the value of State for the instance
func (instance *MDM_MultiSIM_Slots02_01) GetPropertyState() (value int32, err error) {
	retValue, err := instance.GetProperty("State")
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
