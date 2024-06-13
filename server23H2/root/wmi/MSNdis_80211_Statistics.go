// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSNdis_80211_Statistics struct
type MSNdis_80211_Statistics struct {
	*MSNdis

	//
	ACKFailureCount uint64

	//
	Active bool

	//
	FailedCount uint64

	//
	FCSErrorCount uint64

	//
	FrameDuplicateCount uint64

	//
	InstanceName string

	//
	MulticastReceivedFrameCount uint64

	//
	MulticastTransmittedFrameCount uint64

	//
	MultipleRetryCount uint64

	//
	ReceivedFragmentCount uint64

	//
	RetryCount uint64

	//
	RTSFailureCount uint64

	//
	RTSSuccessCount uint64

	//
	StatisticsLength uint32

	//
	TransmittedFragmentCount uint64
}

func NewMSNdis_80211_StatisticsEx1(instance *cim.WmiInstance) (newInstance *MSNdis_80211_Statistics, err error) {
	tmp, err := NewMSNdisEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSNdis_80211_Statistics{
		MSNdis: tmp,
	}
	return
}

func NewMSNdis_80211_StatisticsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSNdis_80211_Statistics, err error) {
	tmp, err := NewMSNdisEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSNdis_80211_Statistics{
		MSNdis: tmp,
	}
	return
}

// SetACKFailureCount sets the value of ACKFailureCount for the instance
func (instance *MSNdis_80211_Statistics) SetPropertyACKFailureCount(value uint64) (err error) {
	return instance.SetProperty("ACKFailureCount", (value))
}

// GetACKFailureCount gets the value of ACKFailureCount for the instance
func (instance *MSNdis_80211_Statistics) GetPropertyACKFailureCount() (value uint64, err error) {
	retValue, err := instance.GetProperty("ACKFailureCount")
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

// SetActive sets the value of Active for the instance
func (instance *MSNdis_80211_Statistics) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSNdis_80211_Statistics) GetPropertyActive() (value bool, err error) {
	retValue, err := instance.GetProperty("Active")
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

// SetFailedCount sets the value of FailedCount for the instance
func (instance *MSNdis_80211_Statistics) SetPropertyFailedCount(value uint64) (err error) {
	return instance.SetProperty("FailedCount", (value))
}

// GetFailedCount gets the value of FailedCount for the instance
func (instance *MSNdis_80211_Statistics) GetPropertyFailedCount() (value uint64, err error) {
	retValue, err := instance.GetProperty("FailedCount")
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

// SetFCSErrorCount sets the value of FCSErrorCount for the instance
func (instance *MSNdis_80211_Statistics) SetPropertyFCSErrorCount(value uint64) (err error) {
	return instance.SetProperty("FCSErrorCount", (value))
}

// GetFCSErrorCount gets the value of FCSErrorCount for the instance
func (instance *MSNdis_80211_Statistics) GetPropertyFCSErrorCount() (value uint64, err error) {
	retValue, err := instance.GetProperty("FCSErrorCount")
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

// SetFrameDuplicateCount sets the value of FrameDuplicateCount for the instance
func (instance *MSNdis_80211_Statistics) SetPropertyFrameDuplicateCount(value uint64) (err error) {
	return instance.SetProperty("FrameDuplicateCount", (value))
}

// GetFrameDuplicateCount gets the value of FrameDuplicateCount for the instance
func (instance *MSNdis_80211_Statistics) GetPropertyFrameDuplicateCount() (value uint64, err error) {
	retValue, err := instance.GetProperty("FrameDuplicateCount")
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

// SetInstanceName sets the value of InstanceName for the instance
func (instance *MSNdis_80211_Statistics) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSNdis_80211_Statistics) GetPropertyInstanceName() (value string, err error) {
	retValue, err := instance.GetProperty("InstanceName")
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

// SetMulticastReceivedFrameCount sets the value of MulticastReceivedFrameCount for the instance
func (instance *MSNdis_80211_Statistics) SetPropertyMulticastReceivedFrameCount(value uint64) (err error) {
	return instance.SetProperty("MulticastReceivedFrameCount", (value))
}

// GetMulticastReceivedFrameCount gets the value of MulticastReceivedFrameCount for the instance
func (instance *MSNdis_80211_Statistics) GetPropertyMulticastReceivedFrameCount() (value uint64, err error) {
	retValue, err := instance.GetProperty("MulticastReceivedFrameCount")
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

// SetMulticastTransmittedFrameCount sets the value of MulticastTransmittedFrameCount for the instance
func (instance *MSNdis_80211_Statistics) SetPropertyMulticastTransmittedFrameCount(value uint64) (err error) {
	return instance.SetProperty("MulticastTransmittedFrameCount", (value))
}

// GetMulticastTransmittedFrameCount gets the value of MulticastTransmittedFrameCount for the instance
func (instance *MSNdis_80211_Statistics) GetPropertyMulticastTransmittedFrameCount() (value uint64, err error) {
	retValue, err := instance.GetProperty("MulticastTransmittedFrameCount")
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

// SetMultipleRetryCount sets the value of MultipleRetryCount for the instance
func (instance *MSNdis_80211_Statistics) SetPropertyMultipleRetryCount(value uint64) (err error) {
	return instance.SetProperty("MultipleRetryCount", (value))
}

// GetMultipleRetryCount gets the value of MultipleRetryCount for the instance
func (instance *MSNdis_80211_Statistics) GetPropertyMultipleRetryCount() (value uint64, err error) {
	retValue, err := instance.GetProperty("MultipleRetryCount")
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

// SetReceivedFragmentCount sets the value of ReceivedFragmentCount for the instance
func (instance *MSNdis_80211_Statistics) SetPropertyReceivedFragmentCount(value uint64) (err error) {
	return instance.SetProperty("ReceivedFragmentCount", (value))
}

// GetReceivedFragmentCount gets the value of ReceivedFragmentCount for the instance
func (instance *MSNdis_80211_Statistics) GetPropertyReceivedFragmentCount() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReceivedFragmentCount")
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

// SetRetryCount sets the value of RetryCount for the instance
func (instance *MSNdis_80211_Statistics) SetPropertyRetryCount(value uint64) (err error) {
	return instance.SetProperty("RetryCount", (value))
}

// GetRetryCount gets the value of RetryCount for the instance
func (instance *MSNdis_80211_Statistics) GetPropertyRetryCount() (value uint64, err error) {
	retValue, err := instance.GetProperty("RetryCount")
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

// SetRTSFailureCount sets the value of RTSFailureCount for the instance
func (instance *MSNdis_80211_Statistics) SetPropertyRTSFailureCount(value uint64) (err error) {
	return instance.SetProperty("RTSFailureCount", (value))
}

// GetRTSFailureCount gets the value of RTSFailureCount for the instance
func (instance *MSNdis_80211_Statistics) GetPropertyRTSFailureCount() (value uint64, err error) {
	retValue, err := instance.GetProperty("RTSFailureCount")
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

// SetRTSSuccessCount sets the value of RTSSuccessCount for the instance
func (instance *MSNdis_80211_Statistics) SetPropertyRTSSuccessCount(value uint64) (err error) {
	return instance.SetProperty("RTSSuccessCount", (value))
}

// GetRTSSuccessCount gets the value of RTSSuccessCount for the instance
func (instance *MSNdis_80211_Statistics) GetPropertyRTSSuccessCount() (value uint64, err error) {
	retValue, err := instance.GetProperty("RTSSuccessCount")
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

// SetStatisticsLength sets the value of StatisticsLength for the instance
func (instance *MSNdis_80211_Statistics) SetPropertyStatisticsLength(value uint32) (err error) {
	return instance.SetProperty("StatisticsLength", (value))
}

// GetStatisticsLength gets the value of StatisticsLength for the instance
func (instance *MSNdis_80211_Statistics) GetPropertyStatisticsLength() (value uint32, err error) {
	retValue, err := instance.GetProperty("StatisticsLength")
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

// SetTransmittedFragmentCount sets the value of TransmittedFragmentCount for the instance
func (instance *MSNdis_80211_Statistics) SetPropertyTransmittedFragmentCount(value uint64) (err error) {
	return instance.SetProperty("TransmittedFragmentCount", (value))
}

// GetTransmittedFragmentCount gets the value of TransmittedFragmentCount for the instance
func (instance *MSNdis_80211_Statistics) GetPropertyTransmittedFragmentCount() (value uint64, err error) {
	retValue, err := instance.GetProperty("TransmittedFragmentCount")
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
