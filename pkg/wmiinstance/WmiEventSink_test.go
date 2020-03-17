// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package cim

import (
	"os/exec"
	"testing"
	"time"
)

type testCallbackContext struct {
	test      *testing.T
	completed chan bool
}

func onObjectReady(context interface{}, wmiInstances []*WmiInstance) {
	t := context.(*testCallbackContext)
	t.completed <- true
}

func onCompleted(context interface{}, wmiInstances []*WmiInstance) {
	t := context.(*testCallbackContext)
	t.completed <- true
}

func onProgress(context interface{}, wmiInstances []*WmiInstance) {
	t := context.(*testCallbackContext)
	t.completed <- true
}

func onObjectPut(context interface{}, wmiInstances []*WmiInstance) {
	t := context.(*testCallbackContext)
	t.completed <- true
}

func Test_WmiAsyncEvents(t *testing.T) {
	sessionManager := NewWmiSessionManager()
	defer sessionManager.Dispose()

	session, err := sessionManager.GetLocalSession("ROOT\\CimV2")

	if err != nil {
		t.Errorf("sessionManager.GetSession failed with error %v", err)
		return
	}

	connected, err := session.Connect()

	if !connected || err != nil {
		t.Errorf("session.Connect failed with error %v", err)
		return
	}
	defer session.Close()

	context := testCallbackContext{
		test:      t,
		completed: make(chan bool),
	}
	eventSink, err := CreateWmiEventSink(session, &context, onObjectReady, onCompleted, onProgress, onObjectPut)
	if err != nil {
		t.Errorf("CreateWmiEventSink failed with error '%v'", err)
		return
	}
	defer eventSink.Close()

	_, err = eventSink.Connect()
	if err != nil {
		t.Errorf("Connect failed with error '%v'", err)
		return
	}

	_, err = session.ExecNotificationQueryAsync(eventSink, "SELECT * FROM __InstanceCreationEvent WITHIN 0.1 WHERE TargetInstance ISA 'Win32_Process' and TargetInstance.Name = 'powershell.exe'")
	if err != nil {
		t.Errorf("CallMethod failed with error '%v'", err)
		return
	}

	cmd := exec.Command("powershell.exe", "-Command", "Sleep -Milliseconds 200")
	err = cmd.Start()
	if err != nil {
		t.Errorf("Could not start powershell on this computer")
		return
	}
	err = cmd.Wait()

	completed := false
	startTime := time.Now()
	timeout, _ := time.ParseDuration("1s")
	for time.Since(startTime) < timeout && !completed {
		select {
		case completed = <-context.completed:
			break
		default:
			// continue waiting
			time.Sleep(1 * time.Second)
		}
		for eventSink.PeekAndDispatchMessages() {
			// Continue pumping for message while they arrive
		}
	}

	if !completed {
		t.Errorf("The onObjectReady callback was never called")
		return
	}
}
