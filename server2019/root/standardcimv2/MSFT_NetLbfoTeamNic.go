// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.StandardCimv2
//
// ////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_NetLbfoTeamNic struct
type MSFT_NetLbfoTeamNic struct {
	*MSFT_NetImPlatAdapter

	// 401
	Default bool

	// 400
	Primary bool

	// 399
	VlanID uint32
}

func NewMSFT_NetLbfoTeamNicEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetLbfoTeamNic, err error) {
	tmp, err := NewMSFT_NetImPlatAdapterEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetLbfoTeamNic{
		MSFT_NetImPlatAdapter: tmp,
	}
	return
}

func NewMSFT_NetLbfoTeamNicEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetLbfoTeamNic, err error) {
	tmp, err := NewMSFT_NetImPlatAdapterEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetLbfoTeamNic{
		MSFT_NetImPlatAdapter: tmp,
	}
	return
}

// SetDefault sets the value of Default for the instance
func (instance *MSFT_NetLbfoTeamNic) SetPropertyDefault(value bool) (err error) {
	return instance.SetProperty("Default", (value))
}

// GetDefault gets the value of Default for the instance
func (instance *MSFT_NetLbfoTeamNic) GetPropertyDefault() (value bool, err error) {
	retValue, err := instance.GetProperty("Default")
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

// SetPrimary sets the value of Primary for the instance
func (instance *MSFT_NetLbfoTeamNic) SetPropertyPrimary(value bool) (err error) {
	return instance.SetProperty("Primary", (value))
}

// GetPrimary gets the value of Primary for the instance
func (instance *MSFT_NetLbfoTeamNic) GetPropertyPrimary() (value bool, err error) {
	retValue, err := instance.GetProperty("Primary")
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

// SetVlanID sets the value of VlanID for the instance
func (instance *MSFT_NetLbfoTeamNic) SetPropertyVlanID(value uint32) (err error) {
	return instance.SetProperty("VlanID", (value))
}

// GetVlanID gets the value of VlanID for the instance
func (instance *MSFT_NetLbfoTeamNic) GetPropertyVlanID() (value uint32, err error) {
	retValue, err := instance.GetProperty("VlanID")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}
