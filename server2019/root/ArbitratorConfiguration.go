// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root
//////////////////////////////////////////////
package root

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// __ArbitratorConfiguration struct
type __ArbitratorConfiguration struct {
	*__SystemClass

	//
	OutstandingTasksPerUser uint32

	//
	OutstandingTasksTotal uint32

	//
	PermanentSubscriptionsPerUser uint32

	//
	PermanentSubscriptionsTotal uint32

	//
	PollingInstructionsPerUser uint32

	//
	PollingInstructionsTotal uint32

	//
	PollingMemoryPerUser uint32

	//
	PollingMemoryTotal uint32

	//
	QuotaRetryCount uint32

	//
	QuotaRetryWaitInterval uint32

	//
	TaskThreadsPerUser uint32

	//
	TaskThreadsTotal uint32

	//
	TemporarySubscriptionsPerUser uint32

	//
	TemporarySubscriptionsTotal uint32

	//
	TotalCacheDisk uint32

	//
	TotalCacheDiskPerTask uint32

	//
	TotalCacheDiskPerUser uint32

	//
	TotalCacheMemory uint32

	//
	TotalCacheMemoryPerTask uint32

	//
	TotalCacheMemoryPerUser uint32

	//
	TotalUsers uint32
}

func New__ArbitratorConfigurationEx1(instance *cim.WmiInstance) (newInstance *__ArbitratorConfiguration, err error) {
	tmp, err := New__SystemClassEx1(instance)

	if err != nil {
		return
	}
	newInstance = &__ArbitratorConfiguration{
		__SystemClass: tmp,
	}
	return
}

func New__ArbitratorConfigurationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *__ArbitratorConfiguration, err error) {
	tmp, err := New__SystemClassEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &__ArbitratorConfiguration{
		__SystemClass: tmp,
	}
	return
}

// SetOutstandingTasksPerUser sets the value of OutstandingTasksPerUser for the instance
func (instance *__ArbitratorConfiguration) SetPropertyOutstandingTasksPerUser(value uint32) (err error) {
	return instance.SetProperty("OutstandingTasksPerUser", value)
}

// GetOutstandingTasksPerUser gets the value of OutstandingTasksPerUser for the instance
func (instance *__ArbitratorConfiguration) GetPropertyOutstandingTasksPerUser() (value uint32, err error) {
	retValue, err := instance.GetProperty("OutstandingTasksPerUser")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOutstandingTasksTotal sets the value of OutstandingTasksTotal for the instance
func (instance *__ArbitratorConfiguration) SetPropertyOutstandingTasksTotal(value uint32) (err error) {
	return instance.SetProperty("OutstandingTasksTotal", value)
}

// GetOutstandingTasksTotal gets the value of OutstandingTasksTotal for the instance
func (instance *__ArbitratorConfiguration) GetPropertyOutstandingTasksTotal() (value uint32, err error) {
	retValue, err := instance.GetProperty("OutstandingTasksTotal")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPermanentSubscriptionsPerUser sets the value of PermanentSubscriptionsPerUser for the instance
func (instance *__ArbitratorConfiguration) SetPropertyPermanentSubscriptionsPerUser(value uint32) (err error) {
	return instance.SetProperty("PermanentSubscriptionsPerUser", value)
}

// GetPermanentSubscriptionsPerUser gets the value of PermanentSubscriptionsPerUser for the instance
func (instance *__ArbitratorConfiguration) GetPropertyPermanentSubscriptionsPerUser() (value uint32, err error) {
	retValue, err := instance.GetProperty("PermanentSubscriptionsPerUser")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPermanentSubscriptionsTotal sets the value of PermanentSubscriptionsTotal for the instance
func (instance *__ArbitratorConfiguration) SetPropertyPermanentSubscriptionsTotal(value uint32) (err error) {
	return instance.SetProperty("PermanentSubscriptionsTotal", value)
}

// GetPermanentSubscriptionsTotal gets the value of PermanentSubscriptionsTotal for the instance
func (instance *__ArbitratorConfiguration) GetPropertyPermanentSubscriptionsTotal() (value uint32, err error) {
	retValue, err := instance.GetProperty("PermanentSubscriptionsTotal")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPollingInstructionsPerUser sets the value of PollingInstructionsPerUser for the instance
func (instance *__ArbitratorConfiguration) SetPropertyPollingInstructionsPerUser(value uint32) (err error) {
	return instance.SetProperty("PollingInstructionsPerUser", value)
}

// GetPollingInstructionsPerUser gets the value of PollingInstructionsPerUser for the instance
func (instance *__ArbitratorConfiguration) GetPropertyPollingInstructionsPerUser() (value uint32, err error) {
	retValue, err := instance.GetProperty("PollingInstructionsPerUser")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPollingInstructionsTotal sets the value of PollingInstructionsTotal for the instance
func (instance *__ArbitratorConfiguration) SetPropertyPollingInstructionsTotal(value uint32) (err error) {
	return instance.SetProperty("PollingInstructionsTotal", value)
}

// GetPollingInstructionsTotal gets the value of PollingInstructionsTotal for the instance
func (instance *__ArbitratorConfiguration) GetPropertyPollingInstructionsTotal() (value uint32, err error) {
	retValue, err := instance.GetProperty("PollingInstructionsTotal")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPollingMemoryPerUser sets the value of PollingMemoryPerUser for the instance
func (instance *__ArbitratorConfiguration) SetPropertyPollingMemoryPerUser(value uint32) (err error) {
	return instance.SetProperty("PollingMemoryPerUser", value)
}

// GetPollingMemoryPerUser gets the value of PollingMemoryPerUser for the instance
func (instance *__ArbitratorConfiguration) GetPropertyPollingMemoryPerUser() (value uint32, err error) {
	retValue, err := instance.GetProperty("PollingMemoryPerUser")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPollingMemoryTotal sets the value of PollingMemoryTotal for the instance
func (instance *__ArbitratorConfiguration) SetPropertyPollingMemoryTotal(value uint32) (err error) {
	return instance.SetProperty("PollingMemoryTotal", value)
}

// GetPollingMemoryTotal gets the value of PollingMemoryTotal for the instance
func (instance *__ArbitratorConfiguration) GetPropertyPollingMemoryTotal() (value uint32, err error) {
	retValue, err := instance.GetProperty("PollingMemoryTotal")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetQuotaRetryCount sets the value of QuotaRetryCount for the instance
func (instance *__ArbitratorConfiguration) SetPropertyQuotaRetryCount(value uint32) (err error) {
	return instance.SetProperty("QuotaRetryCount", value)
}

// GetQuotaRetryCount gets the value of QuotaRetryCount for the instance
func (instance *__ArbitratorConfiguration) GetPropertyQuotaRetryCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("QuotaRetryCount")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetQuotaRetryWaitInterval sets the value of QuotaRetryWaitInterval for the instance
func (instance *__ArbitratorConfiguration) SetPropertyQuotaRetryWaitInterval(value uint32) (err error) {
	return instance.SetProperty("QuotaRetryWaitInterval", value)
}

// GetQuotaRetryWaitInterval gets the value of QuotaRetryWaitInterval for the instance
func (instance *__ArbitratorConfiguration) GetPropertyQuotaRetryWaitInterval() (value uint32, err error) {
	retValue, err := instance.GetProperty("QuotaRetryWaitInterval")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTaskThreadsPerUser sets the value of TaskThreadsPerUser for the instance
func (instance *__ArbitratorConfiguration) SetPropertyTaskThreadsPerUser(value uint32) (err error) {
	return instance.SetProperty("TaskThreadsPerUser", value)
}

// GetTaskThreadsPerUser gets the value of TaskThreadsPerUser for the instance
func (instance *__ArbitratorConfiguration) GetPropertyTaskThreadsPerUser() (value uint32, err error) {
	retValue, err := instance.GetProperty("TaskThreadsPerUser")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTaskThreadsTotal sets the value of TaskThreadsTotal for the instance
func (instance *__ArbitratorConfiguration) SetPropertyTaskThreadsTotal(value uint32) (err error) {
	return instance.SetProperty("TaskThreadsTotal", value)
}

// GetTaskThreadsTotal gets the value of TaskThreadsTotal for the instance
func (instance *__ArbitratorConfiguration) GetPropertyTaskThreadsTotal() (value uint32, err error) {
	retValue, err := instance.GetProperty("TaskThreadsTotal")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTemporarySubscriptionsPerUser sets the value of TemporarySubscriptionsPerUser for the instance
func (instance *__ArbitratorConfiguration) SetPropertyTemporarySubscriptionsPerUser(value uint32) (err error) {
	return instance.SetProperty("TemporarySubscriptionsPerUser", value)
}

// GetTemporarySubscriptionsPerUser gets the value of TemporarySubscriptionsPerUser for the instance
func (instance *__ArbitratorConfiguration) GetPropertyTemporarySubscriptionsPerUser() (value uint32, err error) {
	retValue, err := instance.GetProperty("TemporarySubscriptionsPerUser")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTemporarySubscriptionsTotal sets the value of TemporarySubscriptionsTotal for the instance
func (instance *__ArbitratorConfiguration) SetPropertyTemporarySubscriptionsTotal(value uint32) (err error) {
	return instance.SetProperty("TemporarySubscriptionsTotal", value)
}

// GetTemporarySubscriptionsTotal gets the value of TemporarySubscriptionsTotal for the instance
func (instance *__ArbitratorConfiguration) GetPropertyTemporarySubscriptionsTotal() (value uint32, err error) {
	retValue, err := instance.GetProperty("TemporarySubscriptionsTotal")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTotalCacheDisk sets the value of TotalCacheDisk for the instance
func (instance *__ArbitratorConfiguration) SetPropertyTotalCacheDisk(value uint32) (err error) {
	return instance.SetProperty("TotalCacheDisk", value)
}

// GetTotalCacheDisk gets the value of TotalCacheDisk for the instance
func (instance *__ArbitratorConfiguration) GetPropertyTotalCacheDisk() (value uint32, err error) {
	retValue, err := instance.GetProperty("TotalCacheDisk")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTotalCacheDiskPerTask sets the value of TotalCacheDiskPerTask for the instance
func (instance *__ArbitratorConfiguration) SetPropertyTotalCacheDiskPerTask(value uint32) (err error) {
	return instance.SetProperty("TotalCacheDiskPerTask", value)
}

// GetTotalCacheDiskPerTask gets the value of TotalCacheDiskPerTask for the instance
func (instance *__ArbitratorConfiguration) GetPropertyTotalCacheDiskPerTask() (value uint32, err error) {
	retValue, err := instance.GetProperty("TotalCacheDiskPerTask")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTotalCacheDiskPerUser sets the value of TotalCacheDiskPerUser for the instance
func (instance *__ArbitratorConfiguration) SetPropertyTotalCacheDiskPerUser(value uint32) (err error) {
	return instance.SetProperty("TotalCacheDiskPerUser", value)
}

// GetTotalCacheDiskPerUser gets the value of TotalCacheDiskPerUser for the instance
func (instance *__ArbitratorConfiguration) GetPropertyTotalCacheDiskPerUser() (value uint32, err error) {
	retValue, err := instance.GetProperty("TotalCacheDiskPerUser")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTotalCacheMemory sets the value of TotalCacheMemory for the instance
func (instance *__ArbitratorConfiguration) SetPropertyTotalCacheMemory(value uint32) (err error) {
	return instance.SetProperty("TotalCacheMemory", value)
}

// GetTotalCacheMemory gets the value of TotalCacheMemory for the instance
func (instance *__ArbitratorConfiguration) GetPropertyTotalCacheMemory() (value uint32, err error) {
	retValue, err := instance.GetProperty("TotalCacheMemory")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTotalCacheMemoryPerTask sets the value of TotalCacheMemoryPerTask for the instance
func (instance *__ArbitratorConfiguration) SetPropertyTotalCacheMemoryPerTask(value uint32) (err error) {
	return instance.SetProperty("TotalCacheMemoryPerTask", value)
}

// GetTotalCacheMemoryPerTask gets the value of TotalCacheMemoryPerTask for the instance
func (instance *__ArbitratorConfiguration) GetPropertyTotalCacheMemoryPerTask() (value uint32, err error) {
	retValue, err := instance.GetProperty("TotalCacheMemoryPerTask")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTotalCacheMemoryPerUser sets the value of TotalCacheMemoryPerUser for the instance
func (instance *__ArbitratorConfiguration) SetPropertyTotalCacheMemoryPerUser(value uint32) (err error) {
	return instance.SetProperty("TotalCacheMemoryPerUser", value)
}

// GetTotalCacheMemoryPerUser gets the value of TotalCacheMemoryPerUser for the instance
func (instance *__ArbitratorConfiguration) GetPropertyTotalCacheMemoryPerUser() (value uint32, err error) {
	retValue, err := instance.GetProperty("TotalCacheMemoryPerUser")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTotalUsers sets the value of TotalUsers for the instance
func (instance *__ArbitratorConfiguration) SetPropertyTotalUsers(value uint32) (err error) {
	return instance.SetProperty("TotalUsers", value)
}

// GetTotalUsers gets the value of TotalUsers for the instance
func (instance *__ArbitratorConfiguration) GetPropertyTotalUsers() (value uint32, err error) {
	retValue, err := instance.GetProperty("TotalUsers")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
