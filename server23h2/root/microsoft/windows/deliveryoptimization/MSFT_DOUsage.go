// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.Microsoft.Windows.DeliveryOptimization
//////////////////////////////////////////////
package deliveryoptimization

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_DOUsage struct
type MSFT_DOUsage struct {
	*MSFT_DOBaseStatus

	//
	LinkBps uint32

	//
	LinkUsageBps uint32

	//
	MonthlyGroupBytes uint64

	//
	MonthlyInternetBytes uint64

	//
	MonthlyLanBytes uint64

	//
	MonthlyLinkLocalBytes uint64
}

func NewMSFT_DOUsageEx1(instance *cim.WmiInstance) (newInstance *MSFT_DOUsage, err error) {
	tmp, err := NewMSFT_DOBaseStatusEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_DOUsage{
		MSFT_DOBaseStatus: tmp,
	}
	return
}

func NewMSFT_DOUsageEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_DOUsage, err error) {
	tmp, err := NewMSFT_DOBaseStatusEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_DOUsage{
		MSFT_DOBaseStatus: tmp,
	}
	return
}

// SetLinkBps sets the value of LinkBps for the instance
func (instance *MSFT_DOUsage) SetPropertyLinkBps(value uint32) (err error) {
	return instance.SetProperty("LinkBps", (value))
}

// GetLinkBps gets the value of LinkBps for the instance
func (instance *MSFT_DOUsage) GetPropertyLinkBps() (value uint32, err error) {
	retValue, err := instance.GetProperty("LinkBps")
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

// SetLinkUsageBps sets the value of LinkUsageBps for the instance
func (instance *MSFT_DOUsage) SetPropertyLinkUsageBps(value uint32) (err error) {
	return instance.SetProperty("LinkUsageBps", (value))
}

// GetLinkUsageBps gets the value of LinkUsageBps for the instance
func (instance *MSFT_DOUsage) GetPropertyLinkUsageBps() (value uint32, err error) {
	retValue, err := instance.GetProperty("LinkUsageBps")
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

// SetMonthlyGroupBytes sets the value of MonthlyGroupBytes for the instance
func (instance *MSFT_DOUsage) SetPropertyMonthlyGroupBytes(value uint64) (err error) {
	return instance.SetProperty("MonthlyGroupBytes", (value))
}

// GetMonthlyGroupBytes gets the value of MonthlyGroupBytes for the instance
func (instance *MSFT_DOUsage) GetPropertyMonthlyGroupBytes() (value uint64, err error) {
	retValue, err := instance.GetProperty("MonthlyGroupBytes")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetMonthlyInternetBytes sets the value of MonthlyInternetBytes for the instance
func (instance *MSFT_DOUsage) SetPropertyMonthlyInternetBytes(value uint64) (err error) {
	return instance.SetProperty("MonthlyInternetBytes", (value))
}

// GetMonthlyInternetBytes gets the value of MonthlyInternetBytes for the instance
func (instance *MSFT_DOUsage) GetPropertyMonthlyInternetBytes() (value uint64, err error) {
	retValue, err := instance.GetProperty("MonthlyInternetBytes")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetMonthlyLanBytes sets the value of MonthlyLanBytes for the instance
func (instance *MSFT_DOUsage) SetPropertyMonthlyLanBytes(value uint64) (err error) {
	return instance.SetProperty("MonthlyLanBytes", (value))
}

// GetMonthlyLanBytes gets the value of MonthlyLanBytes for the instance
func (instance *MSFT_DOUsage) GetPropertyMonthlyLanBytes() (value uint64, err error) {
	retValue, err := instance.GetProperty("MonthlyLanBytes")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetMonthlyLinkLocalBytes sets the value of MonthlyLinkLocalBytes for the instance
func (instance *MSFT_DOUsage) SetPropertyMonthlyLinkLocalBytes(value uint64) (err error) {
	return instance.SetProperty("MonthlyLinkLocalBytes", (value))
}

// GetMonthlyLinkLocalBytes gets the value of MonthlyLinkLocalBytes for the instance
func (instance *MSFT_DOUsage) GetPropertyMonthlyLinkLocalBytes() (value uint64, err error) {
	retValue, err := instance.GetProperty("MonthlyLinkLocalBytes")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}
