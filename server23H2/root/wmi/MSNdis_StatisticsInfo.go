// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSNdis_StatisticsInfo struct
type MSNdis_StatisticsInfo struct {
	*MSNdis

	//
	Header MSNdis_ObjectHeader

	//
	ifHCInBroadcastOctets uint64

	//
	ifHCInBroadcastPkts uint64

	//
	ifHCInMulticastOctets uint64

	//
	ifHCInMulticastPkts uint64

	//
	ifHCInOctets uint64

	//
	ifHCInUcastOctets uint64

	//
	ifHCInUcastPkts uint64

	//
	ifHCOutBroadcastOctets uint64

	//
	ifHCOutBroadcastPkts uint64

	//
	ifHCOutMulticastOctets uint64

	//
	ifHCOutMulticastPkts uint64

	//
	ifHCOutOctets uint64

	//
	ifHCOutUcastOctets uint64

	//
	ifHCOutUcastPkts uint64

	//
	ifInDiscards uint64

	//
	ifInErrors uint64

	//
	ifOutDiscards uint64

	//
	ifOutErrors uint64

	//
	SupportedStatistics uint32
}

func NewMSNdis_StatisticsInfoEx1(instance *cim.WmiInstance) (newInstance *MSNdis_StatisticsInfo, err error) {
	tmp, err := NewMSNdisEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSNdis_StatisticsInfo{
		MSNdis: tmp,
	}
	return
}

func NewMSNdis_StatisticsInfoEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSNdis_StatisticsInfo, err error) {
	tmp, err := NewMSNdisEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSNdis_StatisticsInfo{
		MSNdis: tmp,
	}
	return
}

// SetHeader sets the value of Header for the instance
func (instance *MSNdis_StatisticsInfo) SetPropertyHeader(value MSNdis_ObjectHeader) (err error) {
	return instance.SetProperty("Header", (value))
}

// GetHeader gets the value of Header for the instance
func (instance *MSNdis_StatisticsInfo) GetPropertyHeader() (value MSNdis_ObjectHeader, err error) {
	retValue, err := instance.GetProperty("Header")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(MSNdis_ObjectHeader)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " MSNdis_ObjectHeader is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = MSNdis_ObjectHeader(valuetmp)

	return
}

// SetifHCInBroadcastOctets sets the value of ifHCInBroadcastOctets for the instance
func (instance *MSNdis_StatisticsInfo) SetPropertyifHCInBroadcastOctets(value uint64) (err error) {
	return instance.SetProperty("ifHCInBroadcastOctets", (value))
}

// GetifHCInBroadcastOctets gets the value of ifHCInBroadcastOctets for the instance
func (instance *MSNdis_StatisticsInfo) GetPropertyifHCInBroadcastOctets() (value uint64, err error) {
	retValue, err := instance.GetProperty("ifHCInBroadcastOctets")
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

// SetifHCInBroadcastPkts sets the value of ifHCInBroadcastPkts for the instance
func (instance *MSNdis_StatisticsInfo) SetPropertyifHCInBroadcastPkts(value uint64) (err error) {
	return instance.SetProperty("ifHCInBroadcastPkts", (value))
}

// GetifHCInBroadcastPkts gets the value of ifHCInBroadcastPkts for the instance
func (instance *MSNdis_StatisticsInfo) GetPropertyifHCInBroadcastPkts() (value uint64, err error) {
	retValue, err := instance.GetProperty("ifHCInBroadcastPkts")
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

// SetifHCInMulticastOctets sets the value of ifHCInMulticastOctets for the instance
func (instance *MSNdis_StatisticsInfo) SetPropertyifHCInMulticastOctets(value uint64) (err error) {
	return instance.SetProperty("ifHCInMulticastOctets", (value))
}

// GetifHCInMulticastOctets gets the value of ifHCInMulticastOctets for the instance
func (instance *MSNdis_StatisticsInfo) GetPropertyifHCInMulticastOctets() (value uint64, err error) {
	retValue, err := instance.GetProperty("ifHCInMulticastOctets")
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

// SetifHCInMulticastPkts sets the value of ifHCInMulticastPkts for the instance
func (instance *MSNdis_StatisticsInfo) SetPropertyifHCInMulticastPkts(value uint64) (err error) {
	return instance.SetProperty("ifHCInMulticastPkts", (value))
}

// GetifHCInMulticastPkts gets the value of ifHCInMulticastPkts for the instance
func (instance *MSNdis_StatisticsInfo) GetPropertyifHCInMulticastPkts() (value uint64, err error) {
	retValue, err := instance.GetProperty("ifHCInMulticastPkts")
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

// SetifHCInOctets sets the value of ifHCInOctets for the instance
func (instance *MSNdis_StatisticsInfo) SetPropertyifHCInOctets(value uint64) (err error) {
	return instance.SetProperty("ifHCInOctets", (value))
}

// GetifHCInOctets gets the value of ifHCInOctets for the instance
func (instance *MSNdis_StatisticsInfo) GetPropertyifHCInOctets() (value uint64, err error) {
	retValue, err := instance.GetProperty("ifHCInOctets")
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

// SetifHCInUcastOctets sets the value of ifHCInUcastOctets for the instance
func (instance *MSNdis_StatisticsInfo) SetPropertyifHCInUcastOctets(value uint64) (err error) {
	return instance.SetProperty("ifHCInUcastOctets", (value))
}

// GetifHCInUcastOctets gets the value of ifHCInUcastOctets for the instance
func (instance *MSNdis_StatisticsInfo) GetPropertyifHCInUcastOctets() (value uint64, err error) {
	retValue, err := instance.GetProperty("ifHCInUcastOctets")
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

// SetifHCInUcastPkts sets the value of ifHCInUcastPkts for the instance
func (instance *MSNdis_StatisticsInfo) SetPropertyifHCInUcastPkts(value uint64) (err error) {
	return instance.SetProperty("ifHCInUcastPkts", (value))
}

// GetifHCInUcastPkts gets the value of ifHCInUcastPkts for the instance
func (instance *MSNdis_StatisticsInfo) GetPropertyifHCInUcastPkts() (value uint64, err error) {
	retValue, err := instance.GetProperty("ifHCInUcastPkts")
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

// SetifHCOutBroadcastOctets sets the value of ifHCOutBroadcastOctets for the instance
func (instance *MSNdis_StatisticsInfo) SetPropertyifHCOutBroadcastOctets(value uint64) (err error) {
	return instance.SetProperty("ifHCOutBroadcastOctets", (value))
}

// GetifHCOutBroadcastOctets gets the value of ifHCOutBroadcastOctets for the instance
func (instance *MSNdis_StatisticsInfo) GetPropertyifHCOutBroadcastOctets() (value uint64, err error) {
	retValue, err := instance.GetProperty("ifHCOutBroadcastOctets")
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

// SetifHCOutBroadcastPkts sets the value of ifHCOutBroadcastPkts for the instance
func (instance *MSNdis_StatisticsInfo) SetPropertyifHCOutBroadcastPkts(value uint64) (err error) {
	return instance.SetProperty("ifHCOutBroadcastPkts", (value))
}

// GetifHCOutBroadcastPkts gets the value of ifHCOutBroadcastPkts for the instance
func (instance *MSNdis_StatisticsInfo) GetPropertyifHCOutBroadcastPkts() (value uint64, err error) {
	retValue, err := instance.GetProperty("ifHCOutBroadcastPkts")
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

// SetifHCOutMulticastOctets sets the value of ifHCOutMulticastOctets for the instance
func (instance *MSNdis_StatisticsInfo) SetPropertyifHCOutMulticastOctets(value uint64) (err error) {
	return instance.SetProperty("ifHCOutMulticastOctets", (value))
}

// GetifHCOutMulticastOctets gets the value of ifHCOutMulticastOctets for the instance
func (instance *MSNdis_StatisticsInfo) GetPropertyifHCOutMulticastOctets() (value uint64, err error) {
	retValue, err := instance.GetProperty("ifHCOutMulticastOctets")
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

// SetifHCOutMulticastPkts sets the value of ifHCOutMulticastPkts for the instance
func (instance *MSNdis_StatisticsInfo) SetPropertyifHCOutMulticastPkts(value uint64) (err error) {
	return instance.SetProperty("ifHCOutMulticastPkts", (value))
}

// GetifHCOutMulticastPkts gets the value of ifHCOutMulticastPkts for the instance
func (instance *MSNdis_StatisticsInfo) GetPropertyifHCOutMulticastPkts() (value uint64, err error) {
	retValue, err := instance.GetProperty("ifHCOutMulticastPkts")
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

// SetifHCOutOctets sets the value of ifHCOutOctets for the instance
func (instance *MSNdis_StatisticsInfo) SetPropertyifHCOutOctets(value uint64) (err error) {
	return instance.SetProperty("ifHCOutOctets", (value))
}

// GetifHCOutOctets gets the value of ifHCOutOctets for the instance
func (instance *MSNdis_StatisticsInfo) GetPropertyifHCOutOctets() (value uint64, err error) {
	retValue, err := instance.GetProperty("ifHCOutOctets")
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

// SetifHCOutUcastOctets sets the value of ifHCOutUcastOctets for the instance
func (instance *MSNdis_StatisticsInfo) SetPropertyifHCOutUcastOctets(value uint64) (err error) {
	return instance.SetProperty("ifHCOutUcastOctets", (value))
}

// GetifHCOutUcastOctets gets the value of ifHCOutUcastOctets for the instance
func (instance *MSNdis_StatisticsInfo) GetPropertyifHCOutUcastOctets() (value uint64, err error) {
	retValue, err := instance.GetProperty("ifHCOutUcastOctets")
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

// SetifHCOutUcastPkts sets the value of ifHCOutUcastPkts for the instance
func (instance *MSNdis_StatisticsInfo) SetPropertyifHCOutUcastPkts(value uint64) (err error) {
	return instance.SetProperty("ifHCOutUcastPkts", (value))
}

// GetifHCOutUcastPkts gets the value of ifHCOutUcastPkts for the instance
func (instance *MSNdis_StatisticsInfo) GetPropertyifHCOutUcastPkts() (value uint64, err error) {
	retValue, err := instance.GetProperty("ifHCOutUcastPkts")
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

// SetifInDiscards sets the value of ifInDiscards for the instance
func (instance *MSNdis_StatisticsInfo) SetPropertyifInDiscards(value uint64) (err error) {
	return instance.SetProperty("ifInDiscards", (value))
}

// GetifInDiscards gets the value of ifInDiscards for the instance
func (instance *MSNdis_StatisticsInfo) GetPropertyifInDiscards() (value uint64, err error) {
	retValue, err := instance.GetProperty("ifInDiscards")
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

// SetifInErrors sets the value of ifInErrors for the instance
func (instance *MSNdis_StatisticsInfo) SetPropertyifInErrors(value uint64) (err error) {
	return instance.SetProperty("ifInErrors", (value))
}

// GetifInErrors gets the value of ifInErrors for the instance
func (instance *MSNdis_StatisticsInfo) GetPropertyifInErrors() (value uint64, err error) {
	retValue, err := instance.GetProperty("ifInErrors")
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

// SetifOutDiscards sets the value of ifOutDiscards for the instance
func (instance *MSNdis_StatisticsInfo) SetPropertyifOutDiscards(value uint64) (err error) {
	return instance.SetProperty("ifOutDiscards", (value))
}

// GetifOutDiscards gets the value of ifOutDiscards for the instance
func (instance *MSNdis_StatisticsInfo) GetPropertyifOutDiscards() (value uint64, err error) {
	retValue, err := instance.GetProperty("ifOutDiscards")
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

// SetifOutErrors sets the value of ifOutErrors for the instance
func (instance *MSNdis_StatisticsInfo) SetPropertyifOutErrors(value uint64) (err error) {
	return instance.SetProperty("ifOutErrors", (value))
}

// GetifOutErrors gets the value of ifOutErrors for the instance
func (instance *MSNdis_StatisticsInfo) GetPropertyifOutErrors() (value uint64, err error) {
	retValue, err := instance.GetProperty("ifOutErrors")
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

// SetSupportedStatistics sets the value of SupportedStatistics for the instance
func (instance *MSNdis_StatisticsInfo) SetPropertySupportedStatistics(value uint32) (err error) {
	return instance.SetProperty("SupportedStatistics", (value))
}

// GetSupportedStatistics gets the value of SupportedStatistics for the instance
func (instance *MSNdis_StatisticsInfo) GetPropertySupportedStatistics() (value uint32, err error) {
	retValue, err := instance.GetProperty("SupportedStatistics")
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
