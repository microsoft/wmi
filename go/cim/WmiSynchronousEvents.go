// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

// This class wraps a typicaly SWbemSink object. Its implementation is based on the
// Documentation: https://docs.microsoft.com/en-us/windows/win32/wmisdk/receiving-synchronous-and-semisynchronous-event-notifications

package cim

import (
	ole "github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
)

const (
	wbemTimeoutInfinite int32 = -1
)

type WmiSynchronousEventsList struct {
	session     *WmiSession
	instance    *ole.IDispatch
	instanceVar *ole.VARIANT
}

func CreateWmiSynchronousEventsList(instanceVar *ole.VARIANT, session *WmiSession) *WmiSynchronousEventsList {
	return &WmiSynchronousEventsList{
		session:     session,
		instance:    instanceVar.ToIDispatch(),
		instanceVar: instanceVar,
	}
}

func (c *WmiSynchronousEventsList) Close() {
	c.instanceVar.Clear()
}

func (c *WmiSynchronousEventsList) WaitForNextEventUntil(timeout int32) (*WmiInstance, error) {
	eventInstance, err := oleutil.CallMethod(c.instance, "NextEvent", timeout)
	if err != nil {
		return nil, err
	}

	return CreateWmiInstance(eventInstance, c.session)
}

func (c *WmiSynchronousEventsList) WaitForNextEvent() (*WmiInstance, error) {
	return c.WaitForNextEventUntil(wbemTimeoutInfinite)
}
