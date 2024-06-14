// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.Microsoft.Windows.RemoteAccess
//////////////////////////////////////////////
package remoteaccess

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// BgpMessageStatistics struct
type BgpMessageStatistics struct {
	*cim.WmiInstance

	//
	LastReceived string

	//
	LastSent string

	//
	ReceivedCount uint64

	//
	SentCount uint64
}

func NewBgpMessageStatisticsEx1(instance *cim.WmiInstance) (newInstance *BgpMessageStatistics, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &BgpMessageStatistics{
		WmiInstance: tmp,
	}
	return
}

func NewBgpMessageStatisticsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *BgpMessageStatistics, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &BgpMessageStatistics{
		WmiInstance: tmp,
	}
	return
}

// SetLastReceived sets the value of LastReceived for the instance
func (instance *BgpMessageStatistics) SetPropertyLastReceived(value string) (err error) {
	return instance.SetProperty("LastReceived", (value))
}

// GetLastReceived gets the value of LastReceived for the instance
func (instance *BgpMessageStatistics) GetPropertyLastReceived() (value string, err error) {
	retValue, err := instance.GetProperty("LastReceived")
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

// SetLastSent sets the value of LastSent for the instance
func (instance *BgpMessageStatistics) SetPropertyLastSent(value string) (err error) {
	return instance.SetProperty("LastSent", (value))
}

// GetLastSent gets the value of LastSent for the instance
func (instance *BgpMessageStatistics) GetPropertyLastSent() (value string, err error) {
	retValue, err := instance.GetProperty("LastSent")
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

// SetReceivedCount sets the value of ReceivedCount for the instance
func (instance *BgpMessageStatistics) SetPropertyReceivedCount(value uint64) (err error) {
	return instance.SetProperty("ReceivedCount", (value))
}

// GetReceivedCount gets the value of ReceivedCount for the instance
func (instance *BgpMessageStatistics) GetPropertyReceivedCount() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReceivedCount")
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

// SetSentCount sets the value of SentCount for the instance
func (instance *BgpMessageStatistics) SetPropertySentCount(value uint64) (err error) {
	return instance.SetProperty("SentCount", (value))
}

// GetSentCount gets the value of SentCount for the instance
func (instance *BgpMessageStatistics) GetPropertySentCount() (value uint64, err error) {
	retValue, err := instance.GetProperty("SentCount")
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
