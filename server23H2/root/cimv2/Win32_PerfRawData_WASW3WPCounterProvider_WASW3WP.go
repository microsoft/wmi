// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Win32_PerfRawData_WASW3WPCounterProvider_WASW3WP struct
type Win32_PerfRawData_WASW3WPCounterProvider_WASW3WP struct {
	*Win32_PerfRawData

	//
	ActiveListenerChannels uint32

	//
	ActiveProtocolHandlers uint32

	//
	HealthPingReplyLatency uint64

	//
	TotalHealthPings uint32

	//
	TotalMessagesSenttoWAS uint32

	//
	TotalRequestsServed uint32

	//
	TotalRuntimeStatusQueries uint32

	//
	TotalWASMessagesReceived uint32
}

func NewWin32_PerfRawData_WASW3WPCounterProvider_WASW3WPEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfRawData_WASW3WPCounterProvider_WASW3WP, err error) {
	tmp, err := NewWin32_PerfRawDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfRawData_WASW3WPCounterProvider_WASW3WP{
		Win32_PerfRawData: tmp,
	}
	return
}

func NewWin32_PerfRawData_WASW3WPCounterProvider_WASW3WPEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PerfRawData_WASW3WPCounterProvider_WASW3WP, err error) {
	tmp, err := NewWin32_PerfRawDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfRawData_WASW3WPCounterProvider_WASW3WP{
		Win32_PerfRawData: tmp,
	}
	return
}

// SetActiveListenerChannels sets the value of ActiveListenerChannels for the instance
func (instance *Win32_PerfRawData_WASW3WPCounterProvider_WASW3WP) SetPropertyActiveListenerChannels(value uint32) (err error) {
	return instance.SetProperty("ActiveListenerChannels", (value))
}

// GetActiveListenerChannels gets the value of ActiveListenerChannels for the instance
func (instance *Win32_PerfRawData_WASW3WPCounterProvider_WASW3WP) GetPropertyActiveListenerChannels() (value uint32, err error) {
	retValue, err := instance.GetProperty("ActiveListenerChannels")
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

// SetActiveProtocolHandlers sets the value of ActiveProtocolHandlers for the instance
func (instance *Win32_PerfRawData_WASW3WPCounterProvider_WASW3WP) SetPropertyActiveProtocolHandlers(value uint32) (err error) {
	return instance.SetProperty("ActiveProtocolHandlers", (value))
}

// GetActiveProtocolHandlers gets the value of ActiveProtocolHandlers for the instance
func (instance *Win32_PerfRawData_WASW3WPCounterProvider_WASW3WP) GetPropertyActiveProtocolHandlers() (value uint32, err error) {
	retValue, err := instance.GetProperty("ActiveProtocolHandlers")
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

// SetHealthPingReplyLatency sets the value of HealthPingReplyLatency for the instance
func (instance *Win32_PerfRawData_WASW3WPCounterProvider_WASW3WP) SetPropertyHealthPingReplyLatency(value uint64) (err error) {
	return instance.SetProperty("HealthPingReplyLatency", (value))
}

// GetHealthPingReplyLatency gets the value of HealthPingReplyLatency for the instance
func (instance *Win32_PerfRawData_WASW3WPCounterProvider_WASW3WP) GetPropertyHealthPingReplyLatency() (value uint64, err error) {
	retValue, err := instance.GetProperty("HealthPingReplyLatency")
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

// SetTotalHealthPings sets the value of TotalHealthPings for the instance
func (instance *Win32_PerfRawData_WASW3WPCounterProvider_WASW3WP) SetPropertyTotalHealthPings(value uint32) (err error) {
	return instance.SetProperty("TotalHealthPings", (value))
}

// GetTotalHealthPings gets the value of TotalHealthPings for the instance
func (instance *Win32_PerfRawData_WASW3WPCounterProvider_WASW3WP) GetPropertyTotalHealthPings() (value uint32, err error) {
	retValue, err := instance.GetProperty("TotalHealthPings")
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

// SetTotalMessagesSenttoWAS sets the value of TotalMessagesSenttoWAS for the instance
func (instance *Win32_PerfRawData_WASW3WPCounterProvider_WASW3WP) SetPropertyTotalMessagesSenttoWAS(value uint32) (err error) {
	return instance.SetProperty("TotalMessagesSenttoWAS", (value))
}

// GetTotalMessagesSenttoWAS gets the value of TotalMessagesSenttoWAS for the instance
func (instance *Win32_PerfRawData_WASW3WPCounterProvider_WASW3WP) GetPropertyTotalMessagesSenttoWAS() (value uint32, err error) {
	retValue, err := instance.GetProperty("TotalMessagesSenttoWAS")
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

// SetTotalRequestsServed sets the value of TotalRequestsServed for the instance
func (instance *Win32_PerfRawData_WASW3WPCounterProvider_WASW3WP) SetPropertyTotalRequestsServed(value uint32) (err error) {
	return instance.SetProperty("TotalRequestsServed", (value))
}

// GetTotalRequestsServed gets the value of TotalRequestsServed for the instance
func (instance *Win32_PerfRawData_WASW3WPCounterProvider_WASW3WP) GetPropertyTotalRequestsServed() (value uint32, err error) {
	retValue, err := instance.GetProperty("TotalRequestsServed")
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

// SetTotalRuntimeStatusQueries sets the value of TotalRuntimeStatusQueries for the instance
func (instance *Win32_PerfRawData_WASW3WPCounterProvider_WASW3WP) SetPropertyTotalRuntimeStatusQueries(value uint32) (err error) {
	return instance.SetProperty("TotalRuntimeStatusQueries", (value))
}

// GetTotalRuntimeStatusQueries gets the value of TotalRuntimeStatusQueries for the instance
func (instance *Win32_PerfRawData_WASW3WPCounterProvider_WASW3WP) GetPropertyTotalRuntimeStatusQueries() (value uint32, err error) {
	retValue, err := instance.GetProperty("TotalRuntimeStatusQueries")
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

// SetTotalWASMessagesReceived sets the value of TotalWASMessagesReceived for the instance
func (instance *Win32_PerfRawData_WASW3WPCounterProvider_WASW3WP) SetPropertyTotalWASMessagesReceived(value uint32) (err error) {
	return instance.SetProperty("TotalWASMessagesReceived", (value))
}

// GetTotalWASMessagesReceived gets the value of TotalWASMessagesReceived for the instance
func (instance *Win32_PerfRawData_WASW3WPCounterProvider_WASW3WP) GetPropertyTotalWASMessagesReceived() (value uint32, err error) {
	retValue, err := instance.GetProperty("TotalWASMessagesReceived")
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
