// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.Microsoft.Windows.TaskScheduler
//////////////////////////////////////////////
package taskscheduler

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_TaskIdleTrigger struct
type MSFT_TaskIdleTrigger struct {
	*MSFT_TaskTrigger
}

func NewMSFT_TaskIdleTriggerEx1(instance *cim.WmiInstance) (newInstance *MSFT_TaskIdleTrigger, err error) {
	tmp, err := NewMSFT_TaskTriggerEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_TaskIdleTrigger{
		MSFT_TaskTrigger: tmp,
	}
	return
}

func NewMSFT_TaskIdleTriggerEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_TaskIdleTrigger, err error) {
	tmp, err := NewMSFT_TaskTriggerEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_TaskIdleTrigger{
		MSFT_TaskTrigger: tmp,
	}
	return
}
