// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 10/28/2026 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Win32_PerfFormattedData_ClusResPerfProvider_ClusterKeyValueStore struct
type Win32_PerfFormattedData_ClusResPerfProvider_ClusterKeyValueStore struct {
	*Win32_PerfFormattedData

	//
	BytesRead uint64

	//
	BytesReadPersec uint64

	//
	BytesWritten uint64

	//
	BytesWrittenPersec uint64

	//
	IOErrors uint64

	//
	IOErrorsPerSec uint64

	//
	ReadsSent uint64

	//
	ReadsSentPersec uint64

	//
	WritesSent uint64

	//
	WritesSentPersec uint64
}

func NewWin32_PerfFormattedData_ClusResPerfProvider_ClusterKeyValueStoreEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfFormattedData_ClusResPerfProvider_ClusterKeyValueStore, err error) {
	tmp, err := NewWin32_PerfFormattedDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfFormattedData_ClusResPerfProvider_ClusterKeyValueStore{
		Win32_PerfFormattedData: tmp,
	}
	return
}

func NewWin32_PerfFormattedData_ClusResPerfProvider_ClusterKeyValueStoreEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PerfFormattedData_ClusResPerfProvider_ClusterKeyValueStore, err error) {
	tmp, err := NewWin32_PerfFormattedDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfFormattedData_ClusResPerfProvider_ClusterKeyValueStore{
		Win32_PerfFormattedData: tmp,
	}
	return
}

// SetBytesRead sets the value of BytesRead for the instance
func (instance *Win32_PerfFormattedData_ClusResPerfProvider_ClusterKeyValueStore) SetPropertyBytesRead(value uint64) (err error) {
	return instance.SetProperty("BytesRead", (value))
}

// GetBytesRead gets the value of BytesRead for the instance
func (instance *Win32_PerfFormattedData_ClusResPerfProvider_ClusterKeyValueStore) GetPropertyBytesRead() (value uint64, err error) {
	retValue, err := instance.GetProperty("BytesRead")
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

// SetBytesReadPersec sets the value of BytesReadPersec for the instance
func (instance *Win32_PerfFormattedData_ClusResPerfProvider_ClusterKeyValueStore) SetPropertyBytesReadPersec(value uint64) (err error) {
	return instance.SetProperty("BytesReadPersec", (value))
}

// GetBytesReadPersec gets the value of BytesReadPersec for the instance
func (instance *Win32_PerfFormattedData_ClusResPerfProvider_ClusterKeyValueStore) GetPropertyBytesReadPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("BytesReadPersec")
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

// SetBytesWritten sets the value of BytesWritten for the instance
func (instance *Win32_PerfFormattedData_ClusResPerfProvider_ClusterKeyValueStore) SetPropertyBytesWritten(value uint64) (err error) {
	return instance.SetProperty("BytesWritten", (value))
}

// GetBytesWritten gets the value of BytesWritten for the instance
func (instance *Win32_PerfFormattedData_ClusResPerfProvider_ClusterKeyValueStore) GetPropertyBytesWritten() (value uint64, err error) {
	retValue, err := instance.GetProperty("BytesWritten")
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

// SetBytesWrittenPersec sets the value of BytesWrittenPersec for the instance
func (instance *Win32_PerfFormattedData_ClusResPerfProvider_ClusterKeyValueStore) SetPropertyBytesWrittenPersec(value uint64) (err error) {
	return instance.SetProperty("BytesWrittenPersec", (value))
}

// GetBytesWrittenPersec gets the value of BytesWrittenPersec for the instance
func (instance *Win32_PerfFormattedData_ClusResPerfProvider_ClusterKeyValueStore) GetPropertyBytesWrittenPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("BytesWrittenPersec")
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

// SetIOErrors sets the value of IOErrors for the instance
func (instance *Win32_PerfFormattedData_ClusResPerfProvider_ClusterKeyValueStore) SetPropertyIOErrors(value uint64) (err error) {
	return instance.SetProperty("IOErrors", (value))
}

// GetIOErrors gets the value of IOErrors for the instance
func (instance *Win32_PerfFormattedData_ClusResPerfProvider_ClusterKeyValueStore) GetPropertyIOErrors() (value uint64, err error) {
	retValue, err := instance.GetProperty("IOErrors")
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

// SetIOErrorsPerSec sets the value of IOErrorsPerSec for the instance
func (instance *Win32_PerfFormattedData_ClusResPerfProvider_ClusterKeyValueStore) SetPropertyIOErrorsPerSec(value uint64) (err error) {
	return instance.SetProperty("IOErrorsPerSec", (value))
}

// GetIOErrorsPerSec gets the value of IOErrorsPerSec for the instance
func (instance *Win32_PerfFormattedData_ClusResPerfProvider_ClusterKeyValueStore) GetPropertyIOErrorsPerSec() (value uint64, err error) {
	retValue, err := instance.GetProperty("IOErrorsPerSec")
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

// SetReadsSent sets the value of ReadsSent for the instance
func (instance *Win32_PerfFormattedData_ClusResPerfProvider_ClusterKeyValueStore) SetPropertyReadsSent(value uint64) (err error) {
	return instance.SetProperty("ReadsSent", (value))
}

// GetReadsSent gets the value of ReadsSent for the instance
func (instance *Win32_PerfFormattedData_ClusResPerfProvider_ClusterKeyValueStore) GetPropertyReadsSent() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReadsSent")
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

// SetReadsSentPersec sets the value of ReadsSentPersec for the instance
func (instance *Win32_PerfFormattedData_ClusResPerfProvider_ClusterKeyValueStore) SetPropertyReadsSentPersec(value uint64) (err error) {
	return instance.SetProperty("ReadsSentPersec", (value))
}

// GetReadsSentPersec gets the value of ReadsSentPersec for the instance
func (instance *Win32_PerfFormattedData_ClusResPerfProvider_ClusterKeyValueStore) GetPropertyReadsSentPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReadsSentPersec")
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

// SetWritesSent sets the value of WritesSent for the instance
func (instance *Win32_PerfFormattedData_ClusResPerfProvider_ClusterKeyValueStore) SetPropertyWritesSent(value uint64) (err error) {
	return instance.SetProperty("WritesSent", (value))
}

// GetWritesSent gets the value of WritesSent for the instance
func (instance *Win32_PerfFormattedData_ClusResPerfProvider_ClusterKeyValueStore) GetPropertyWritesSent() (value uint64, err error) {
	retValue, err := instance.GetProperty("WritesSent")
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

// SetWritesSentPersec sets the value of WritesSentPersec for the instance
func (instance *Win32_PerfFormattedData_ClusResPerfProvider_ClusterKeyValueStore) SetPropertyWritesSentPersec(value uint64) (err error) {
	return instance.SetProperty("WritesSentPersec", (value))
}

// GetWritesSentPersec gets the value of WritesSentPersec for the instance
func (instance *Win32_PerfFormattedData_ClusResPerfProvider_ClusterKeyValueStore) GetPropertyWritesSentPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("WritesSentPersec")
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
