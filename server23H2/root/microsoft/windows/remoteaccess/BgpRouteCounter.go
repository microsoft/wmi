// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.Microsoft.Windows.RemoteAccess
//
// ////////////////////////////////////////////
package remoteaccess

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// BgpRouteCounter struct
type BgpRouteCounter struct {
	*cim.WmiInstance

	//
	UpdateReceivedCount uint64

	//
	UpdateSentCount uint64

	//
	WithdrawlReceivedCount uint64

	//
	WithdrawlSentCount uint64
}

func NewBgpRouteCounterEx1(instance *cim.WmiInstance) (newInstance *BgpRouteCounter, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &BgpRouteCounter{
		WmiInstance: tmp,
	}
	return
}

func NewBgpRouteCounterEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *BgpRouteCounter, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &BgpRouteCounter{
		WmiInstance: tmp,
	}
	return
}

// SetUpdateReceivedCount sets the value of UpdateReceivedCount for the instance
func (instance *BgpRouteCounter) SetPropertyUpdateReceivedCount(value uint64) (err error) {
	return instance.SetProperty("UpdateReceivedCount", (value))
}

// GetUpdateReceivedCount gets the value of UpdateReceivedCount for the instance
func (instance *BgpRouteCounter) GetPropertyUpdateReceivedCount() (value uint64, err error) {
	retValue, err := instance.GetProperty("UpdateReceivedCount")
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

// SetUpdateSentCount sets the value of UpdateSentCount for the instance
func (instance *BgpRouteCounter) SetPropertyUpdateSentCount(value uint64) (err error) {
	return instance.SetProperty("UpdateSentCount", (value))
}

// GetUpdateSentCount gets the value of UpdateSentCount for the instance
func (instance *BgpRouteCounter) GetPropertyUpdateSentCount() (value uint64, err error) {
	retValue, err := instance.GetProperty("UpdateSentCount")
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

// SetWithdrawlReceivedCount sets the value of WithdrawlReceivedCount for the instance
func (instance *BgpRouteCounter) SetPropertyWithdrawlReceivedCount(value uint64) (err error) {
	return instance.SetProperty("WithdrawlReceivedCount", (value))
}

// GetWithdrawlReceivedCount gets the value of WithdrawlReceivedCount for the instance
func (instance *BgpRouteCounter) GetPropertyWithdrawlReceivedCount() (value uint64, err error) {
	retValue, err := instance.GetProperty("WithdrawlReceivedCount")
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

// SetWithdrawlSentCount sets the value of WithdrawlSentCount for the instance
func (instance *BgpRouteCounter) SetPropertyWithdrawlSentCount(value uint64) (err error) {
	return instance.SetProperty("WithdrawlSentCount", (value))
}

// GetWithdrawlSentCount gets the value of WithdrawlSentCount for the instance
func (instance *BgpRouteCounter) GetPropertyWithdrawlSentCount() (value uint64, err error) {
	retValue, err := instance.GetProperty("WithdrawlSentCount")
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
