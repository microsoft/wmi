// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_NetServiceSlowStartup struct
type MSFT_NetServiceSlowStartup struct {
	*MSFT_SCMEventLogEvent

	//
	Service string

	//
	StartupTime uint32
}

func NewMSFT_NetServiceSlowStartupEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetServiceSlowStartup, err error) {
	tmp, err := NewMSFT_SCMEventLogEventEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetServiceSlowStartup{
		MSFT_SCMEventLogEvent: tmp,
	}
	return
}

func NewMSFT_NetServiceSlowStartupEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetServiceSlowStartup, err error) {
	tmp, err := NewMSFT_SCMEventLogEventEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetServiceSlowStartup{
		MSFT_SCMEventLogEvent: tmp,
	}
	return
}

// SetService sets the value of Service for the instance
func (instance *MSFT_NetServiceSlowStartup) SetPropertyService(value string) (err error) {
	return instance.SetProperty("Service", value)
}

// GetService gets the value of Service for the instance
func (instance *MSFT_NetServiceSlowStartup) GetPropertyService() (value string, err error) {
	retValue, err := instance.GetProperty("Service")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStartupTime sets the value of StartupTime for the instance
func (instance *MSFT_NetServiceSlowStartup) SetPropertyStartupTime(value uint32) (err error) {
	return instance.SetProperty("StartupTime", value)
}

// GetStartupTime gets the value of StartupTime for the instance
func (instance *MSFT_NetServiceSlowStartup) GetPropertyStartupTime() (value uint32, err error) {
	retValue, err := instance.GetProperty("StartupTime")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
