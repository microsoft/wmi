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

// MSFT_NetLbfoTeam struct
type MSFT_NetLbfoTeam struct {
	*MSFT_NetImPlatTeam

	// 416
	LacpTimer uint32

	// 12
	LoadBalancingAlgorithm uint32

	// 13
	Status uint32

	// 11
	TeamingMode uint32
}

func NewMSFT_NetLbfoTeamEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetLbfoTeam, err error) {
	tmp, err := NewMSFT_NetImPlatTeamEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetLbfoTeam{
		MSFT_NetImPlatTeam: tmp,
	}
	return
}

func NewMSFT_NetLbfoTeamEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetLbfoTeam, err error) {
	tmp, err := NewMSFT_NetImPlatTeamEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetLbfoTeam{
		MSFT_NetImPlatTeam: tmp,
	}
	return
}

// SetLacpTimer sets the value of LacpTimer for the instance
func (instance *MSFT_NetLbfoTeam) SetPropertyLacpTimer(value uint32) (err error) {
	return instance.SetProperty("LacpTimer", (value))
}

// GetLacpTimer gets the value of LacpTimer for the instance
func (instance *MSFT_NetLbfoTeam) GetPropertyLacpTimer() (value uint32, err error) {
	retValue, err := instance.GetProperty("LacpTimer")
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

// SetLoadBalancingAlgorithm sets the value of LoadBalancingAlgorithm for the instance
func (instance *MSFT_NetLbfoTeam) SetPropertyLoadBalancingAlgorithm(value uint32) (err error) {
	return instance.SetProperty("LoadBalancingAlgorithm", (value))
}

// GetLoadBalancingAlgorithm gets the value of LoadBalancingAlgorithm for the instance
func (instance *MSFT_NetLbfoTeam) GetPropertyLoadBalancingAlgorithm() (value uint32, err error) {
	retValue, err := instance.GetProperty("LoadBalancingAlgorithm")
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

// SetStatus sets the value of Status for the instance
func (instance *MSFT_NetLbfoTeam) SetPropertyStatus(value uint32) (err error) {
	return instance.SetProperty("Status", (value))
}

// GetStatus gets the value of Status for the instance
func (instance *MSFT_NetLbfoTeam) GetPropertyStatus() (value uint32, err error) {
	retValue, err := instance.GetProperty("Status")
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

// SetTeamingMode sets the value of TeamingMode for the instance
func (instance *MSFT_NetLbfoTeam) SetPropertyTeamingMode(value uint32) (err error) {
	return instance.SetProperty("TeamingMode", (value))
}

// GetTeamingMode gets the value of TeamingMode for the instance
func (instance *MSFT_NetLbfoTeam) GetPropertyTeamingMode() (value uint32, err error) {
	retValue, err := instance.GetProperty("TeamingMode")
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

// 14

// <param name="Name" type="string "></param>
// <param name="NewName" type="string "></param>

// <param name="CmdletOutput" type="MSFT_NetLbfoTeam "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetLbfoTeam) Rename( /* IN */ Name string,
	/* IN */ NewName string,
	/* OUT */ CmdletOutput MSFT_NetLbfoTeam) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Rename", Name, NewName)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
