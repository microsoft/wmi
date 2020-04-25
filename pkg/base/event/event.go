// +build windows
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package event

import (
	"github.com/microsoft/wmi/pkg/base/session"
	wmi "github.com/microsoft/wmi/pkg/wmiinstance"
)

type CallbackContext struct {
	callbackData string
	callback     func(*wmi.WmiInstance, string)
}

// NewCallbackContext
func NewCallbackContext(cb func(*wmi.WmiInstance, string), data string) *CallbackContext {
	return &CallbackContext{
		callback:     cb,
		callbackData: data,
	}
}

// Execute the callback
func (cb *CallbackContext) Execute() {
	if cb.callback == nil {
		return
	}
	cb.callback(nil, cb.callbackData)
}

func onInstanceReady(ctx interface{}, wmiInstances []*wmi.WmiInstance) {
	context := ctx.(*CallbackContext)
	if context == nil {
		return
	}

	context.Execute()
}

func onCompleted(ctx interface{}, wmiInstances []*wmi.WmiInstance) {
	// Not used
}

func onProgress(ctx interface{}, wmiInstances []*wmi.WmiInstance) {
	// Not used
}

func onInstancePut(ctx interface{}, wmiInstances []*wmi.WmiInstance) {
	// Not used
}

// RegisterWmiCallback
func RegisterWmiCallback(context *CallbackContext, wmiNamespace, hostName, queryString string) (eventSink *wmi.WmiEventSink, err error) {
	wmiSession, err := session.GetSession(wmiNamespace, hostName, "", "", "")
	if err != nil {
		return nil, err
	}

	eventSink, err = wmi.CreateWmiEventSink(wmiSession, context, onInstanceReady, onCompleted, onProgress, onInstancePut)
	if err != nil {
		return
	}
	defer func() {
		if err != nil {
			eventSink.Close()
			eventSink = nil
		}
	}()

	_, err = eventSink.Connect()
	if err != nil {
		return
	}

	_, err = wmiSession.ExecNotificationQueryAsync(eventSink, queryString)
	if err != nil {
		return
	}

	return
}
