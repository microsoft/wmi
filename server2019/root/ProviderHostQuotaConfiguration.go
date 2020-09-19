// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root
//////////////////////////////////////////////
package root

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// __ProviderHostQuotaConfiguration struct
type __ProviderHostQuotaConfiguration struct {
	*__SystemClass

	//
	HandlesPerHost uint32

	//
	MemoryAllHosts uint64

	//
	MemoryPerHost uint64

	//
	ProcessLimitAllHosts uint32

	//
	ThreadsPerHost uint32
}

func New__ProviderHostQuotaConfigurationEx1(instance *cim.WmiInstance) (newInstance *__ProviderHostQuotaConfiguration, err error) {
	tmp, err := New__SystemClassEx1(instance)

	if err != nil {
		return
	}
	newInstance = &__ProviderHostQuotaConfiguration{
		__SystemClass: tmp,
	}
	return
}

func New__ProviderHostQuotaConfigurationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *__ProviderHostQuotaConfiguration, err error) {
	tmp, err := New__SystemClassEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &__ProviderHostQuotaConfiguration{
		__SystemClass: tmp,
	}
	return
}

// SetHandlesPerHost sets the value of HandlesPerHost for the instance
func (instance *__ProviderHostQuotaConfiguration) SetPropertyHandlesPerHost(value uint32) (err error) {
	return instance.SetProperty("HandlesPerHost", (value))
}

// GetHandlesPerHost gets the value of HandlesPerHost for the instance
func (instance *__ProviderHostQuotaConfiguration) GetPropertyHandlesPerHost() (value uint32, err error) {
	retValue, err := instance.GetProperty("HandlesPerHost")
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

// SetMemoryAllHosts sets the value of MemoryAllHosts for the instance
func (instance *__ProviderHostQuotaConfiguration) SetPropertyMemoryAllHosts(value uint64) (err error) {
	return instance.SetProperty("MemoryAllHosts", (value))
}

// GetMemoryAllHosts gets the value of MemoryAllHosts for the instance
func (instance *__ProviderHostQuotaConfiguration) GetPropertyMemoryAllHosts() (value uint64, err error) {
	retValue, err := instance.GetProperty("MemoryAllHosts")
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

// SetMemoryPerHost sets the value of MemoryPerHost for the instance
func (instance *__ProviderHostQuotaConfiguration) SetPropertyMemoryPerHost(value uint64) (err error) {
	return instance.SetProperty("MemoryPerHost", (value))
}

// GetMemoryPerHost gets the value of MemoryPerHost for the instance
func (instance *__ProviderHostQuotaConfiguration) GetPropertyMemoryPerHost() (value uint64, err error) {
	retValue, err := instance.GetProperty("MemoryPerHost")
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

// SetProcessLimitAllHosts sets the value of ProcessLimitAllHosts for the instance
func (instance *__ProviderHostQuotaConfiguration) SetPropertyProcessLimitAllHosts(value uint32) (err error) {
	return instance.SetProperty("ProcessLimitAllHosts", (value))
}

// GetProcessLimitAllHosts gets the value of ProcessLimitAllHosts for the instance
func (instance *__ProviderHostQuotaConfiguration) GetPropertyProcessLimitAllHosts() (value uint32, err error) {
	retValue, err := instance.GetProperty("ProcessLimitAllHosts")
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

// SetThreadsPerHost sets the value of ThreadsPerHost for the instance
func (instance *__ProviderHostQuotaConfiguration) SetPropertyThreadsPerHost(value uint32) (err error) {
	return instance.SetProperty("ThreadsPerHost", (value))
}

// GetThreadsPerHost gets the value of ThreadsPerHost for the instance
func (instance *__ProviderHostQuotaConfiguration) GetPropertyThreadsPerHost() (value uint32, err error) {
	retValue, err := instance.GetProperty("ThreadsPerHost")
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
