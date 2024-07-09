// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package cim

import (
	"os/exec"
	"testing"
)

func Test_WmiSynchronousEvents(t *testing.T) {
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

	instances, err := session.ExecNotificationQuery("SELECT * FROM __InstanceCreationEvent WITHIN 0.1 WHERE TargetInstance ISA 'Win32_Process' and TargetInstance.Name = 'powershell.exe'")
	if err != nil {
		t.Errorf("ExecNotificationQuery failed with error '%v'", err)
		return
	}
	defer instances.Close()

	cmd := exec.Command("powershell.exe", "-Command", "Sleep -Milliseconds 200")
	err = cmd.Start()
	if err != nil {
		t.Errorf("Could not start powershell on this computer")
		return
	}
	err = cmd.Wait()
	if err != nil {
		t.Errorf("Could not wait for powershell on this computer")
		return
	}

	_, err = instances.WaitForNextEvent()
	if err != nil {
		t.Errorf("WaitForNextEvent failed with error '%v'", err)
		return
	}
}
