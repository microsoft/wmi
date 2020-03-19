// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.Microsoft.Windows.RemoteAccess
//////////////////////////////////////////////
package remoteaccess

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// __AbsoluteTimerInstruction struct
type __AbsoluteTimerInstruction struct {
	*__TimerInstruction

	//
	EventDateTime string
}

func New__AbsoluteTimerInstructionEx1(instance *cim.WmiInstance) (newInstance *__AbsoluteTimerInstruction, err error) {
	tmp, err := New__TimerInstructionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &__AbsoluteTimerInstruction{
		__TimerInstruction: tmp,
	}
	return
}

func New__AbsoluteTimerInstructionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *__AbsoluteTimerInstruction, err error) {
	tmp, err := New__TimerInstructionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &__AbsoluteTimerInstruction{
		__TimerInstruction: tmp,
	}
	return
}

// SetEventDateTime sets the value of EventDateTime for the instance
func (instance *__AbsoluteTimerInstruction) SetPropertyEventDateTime(value string) (err error) {
	return instance.SetProperty("EventDateTime", value)
}

// GetEventDateTime gets the value of EventDateTime for the instance
func (instance *__AbsoluteTimerInstruction) GetPropertyEventDateTime() (value string, err error) {
	retValue, err := instance.GetProperty("EventDateTime")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
