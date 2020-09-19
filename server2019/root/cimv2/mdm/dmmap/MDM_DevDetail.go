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

// MDM_DevDetail struct
type MDM_DevDetail struct {
	*cim.WmiInstance

	//
	DevTyp string

	//
	FwV string

	//
	HwV string

	//
	InstanceID string

	//
	LrgObj bool

	//
	OEM string

	//
	ParentID string

	//
	SwV string
}

func NewMDM_DevDetailEx1(instance *cim.WmiInstance) (newInstance *MDM_DevDetail, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_DevDetail{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_DevDetailEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_DevDetail, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_DevDetail{
		WmiInstance: tmp,
	}
	return
}

// SetDevTyp sets the value of DevTyp for the instance
func (instance *MDM_DevDetail) SetPropertyDevTyp(value string) (err error) {
	return instance.SetProperty("DevTyp", (value))
}

// GetDevTyp gets the value of DevTyp for the instance
func (instance *MDM_DevDetail) GetPropertyDevTyp() (value string, err error) {
	retValue, err := instance.GetProperty("DevTyp")
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

// SetFwV sets the value of FwV for the instance
func (instance *MDM_DevDetail) SetPropertyFwV(value string) (err error) {
	return instance.SetProperty("FwV", (value))
}

// GetFwV gets the value of FwV for the instance
func (instance *MDM_DevDetail) GetPropertyFwV() (value string, err error) {
	retValue, err := instance.GetProperty("FwV")
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

// SetHwV sets the value of HwV for the instance
func (instance *MDM_DevDetail) SetPropertyHwV(value string) (err error) {
	return instance.SetProperty("HwV", (value))
}

// GetHwV gets the value of HwV for the instance
func (instance *MDM_DevDetail) GetPropertyHwV() (value string, err error) {
	retValue, err := instance.GetProperty("HwV")
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
func (instance *MDM_DevDetail) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", (value))
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_DevDetail) GetPropertyInstanceID() (value string, err error) {
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

// SetLrgObj sets the value of LrgObj for the instance
func (instance *MDM_DevDetail) SetPropertyLrgObj(value bool) (err error) {
	return instance.SetProperty("LrgObj", (value))
}

// GetLrgObj gets the value of LrgObj for the instance
func (instance *MDM_DevDetail) GetPropertyLrgObj() (value bool, err error) {
	retValue, err := instance.GetProperty("LrgObj")
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

// SetOEM sets the value of OEM for the instance
func (instance *MDM_DevDetail) SetPropertyOEM(value string) (err error) {
	return instance.SetProperty("OEM", (value))
}

// GetOEM gets the value of OEM for the instance
func (instance *MDM_DevDetail) GetPropertyOEM() (value string, err error) {
	retValue, err := instance.GetProperty("OEM")
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
func (instance *MDM_DevDetail) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", (value))
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_DevDetail) GetPropertyParentID() (value string, err error) {
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

// SetSwV sets the value of SwV for the instance
func (instance *MDM_DevDetail) SetPropertySwV(value string) (err error) {
	return instance.SetProperty("SwV", (value))
}

// GetSwV gets the value of SwV for the instance
func (instance *MDM_DevDetail) GetPropertySwV() (value string, err error) {
	retValue, err := instance.GetProperty("SwV")
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
