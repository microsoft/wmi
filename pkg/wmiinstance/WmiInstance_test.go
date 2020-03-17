// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package cim

import (
	"strings"
	"testing"

	"github.com/go-ole/go-ole/oleutil"
)

func Test_WmiInstance(t *testing.T) {
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

	systemIdleProcessInstances, err := session.QueryInstances("SELECT * FROM WIN32_Process WHERE ProcessId = 0")
	if err != nil {
		t.Errorf("session.QueryInstances failed with error %v", err)
		return
	}
	defer CloseAllInstances(systemIdleProcessInstances)

	if len(systemIdleProcessInstances) != 1 {
		t.Errorf("session.QueryInstances returned (%v) processes where it should have returned 1", len(systemIdleProcessInstances))
		return
	}

	systemIdleProcess := systemIdleProcessInstances[0]

	systemProcessInstances, err := session.QueryInstances("SELECT * FROM WIN32_Process WHERE ProcessId = 4")
	if err != nil {
		t.Errorf("session.QueryInstances failed with error %v", err)
		return
	}
	defer CloseAllInstances(systemProcessInstances)

	if len(systemProcessInstances) != 1 {
		t.Errorf("session.QueryInstances returned (%v) processes where it should have returned 1", len(systemProcessInstances))
		return
	}

	systemProcess := systemProcessInstances[0]

	processName, err := systemIdleProcess.GetProperty("Name")
	if err != nil {
		t.Errorf("systemIdleProcess.GetProperty failed with error %v", err)
		return
	}

	newInstance, err := systemIdleProcess.GetInstance()
	if err != nil {
		t.Errorf("systemIdleProcess.GetInstance failed with error %v", err)
		return
	}
	defer newInstance.Close()

	newProcessName, err := newInstance.GetProperty("Name")
	if err != nil {
		t.Errorf("newInstance.GetProperty failed with error %v", err)
		return
	}

	if processName.(string) != newProcessName.(string) {
		t.Errorf("systemIdleProcess.GetInstance didn't return the same object [%v] != [%v]", newProcessName.(string), processName.(string))
		return
	}

	originalStatus, err := systemIdleProcess.GetProperty("Status")
	if err != nil {
		t.Errorf("systemIdleProcess.GetProperty failed with error %v", err)
		return
	}

	err = systemIdleProcess.SetProperty("Status", "hello")
	if err != nil {
		t.Errorf("systemIdleProcess.SetProperty failed with error %v", err)
		return
	}
	status, err := systemIdleProcess.GetProperty("Status")
	if err != nil {
		t.Errorf("systemIdleProcess.GetProperty failed with error %v", err)
		return
	}

	if status.(string) != "hello" {
		t.Errorf("systemIdleProcess.SetProperty didn't return the expected object [%v] != \"hello\"", status.(string))
		return
	}

	err = systemIdleProcess.ResetProperty("Status")
	if err != nil {
		t.Errorf("systemIdleProcess.ResetProperty failed with error %v", err)
		return
	}

	finalStatus, err := systemIdleProcess.GetProperty("Status")
	if err != nil {
		t.Errorf("systemIdleProcess.GetProperty failed with error %v", err)
		return
	}

	originalStatusStr := ""
	if originalStatus != nil {
		originalStatusStr = originalStatus.(string)
	}

	finalStatusStr := ""
	if finalStatus != nil {
		finalStatusStr = finalStatus.(string)
	}

	if originalStatusStr != finalStatusStr {
		t.Errorf("systemIdleProcess.ResetProperty didn't reset the object correctly. [%v] != [%v]", originalStatus.(string), finalStatus.(string))
		return
	}

	wmiClass := systemIdleProcess.GetClass()
	if wmiClass == nil {
		t.Errorf("systemIdleProcess.GetClass() couldn't find any object.")
		return
	}
	defer wmiClass.Close()

	className, err := wmiClass.GetSystemProperty("__CLASS")
	defer className.Close()

	if className.Value().(string) != "Win32_Process" {
		t.Errorf("systemIdleProcess.GetClass() didn't return the correct object. [%v] != \"Win32_Process\"", className)
		return
	}

	embeddedInstance, err := systemIdleProcess.EmbeddedInstance()
	if err != nil {
		t.Errorf("EmbeddedInstance() failed with error: %v", err)
		return
	}

	if !strings.Contains(embeddedInstance, "instance of Win32_Process") {
		t.Errorf("The returned embeddedInstance didn't return the correct content. [%s]", embeddedInstance)
		return
	}

	if !systemIdleProcess.Equals(systemIdleProcess) {
		t.Errorf("Error: Equals() considered the systemIdleProcess is not equal to itself")
		return
	}

	if systemIdleProcess.Equals(systemProcess) {
		t.Errorf("Error: Equals() wrongly considered the systemIdleProcess and the systemProcess are identical")
		return
	}

	err = systemIdleProcess.Refresh()
	if err != nil {
		t.Errorf("Error: Refresh() failed with error: %v", err)
		return
	}

	if !strings.Contains(systemIdleProcess.InstancePath(), "Win32_Process") {
		t.Errorf("Error: InstancePath() didn't return a correct path")
		return
	}

	if !strings.Contains(systemIdleProcess.RelativePath(), "Win32_Process") {
		t.Errorf("Error: RelativePath() didn't return a correct path")
		return
	}

	_, err = systemIdleProcess.InvokeMethod("GetOwnerSid")
	if err != nil {
		t.Errorf("systemIdleProcess.InvokeMethod failed with error %v", err)
		return
	}

	relatedInstances, err := systemIdleProcess.GetRelated("Win32_ComputerSystem")
	if err != nil {
		t.Errorf("systemIdleProcess.GetRelated failed with error %v", err)
		return
	}
	defer CloseAllInstances(relatedInstances)

	if len(relatedInstances) < 1 {
		t.Errorf("systemIdleProcess.GetRelated didn't find any related Win32_ComputerSystem instances of Win32_Process")
		return
	}

	for _, relatedInstance := range relatedInstances {
		if relatedInstance.GetClassName() != "Win32_ComputerSystem" {
			t.Errorf("systemIdleProcess.GetRelated didn't find any related instances of Win32_ComputerSystem")
			return
		}
	}

	referencingInstances, err := systemIdleProcess.EnumerateReferencingInstances("Win32_SystemProcesses", "")
	if err != nil {
		t.Errorf("systemIdleProcess.EnumerateReferencingInstances failed with error %v", err)
		return
	}
	defer CloseAllInstances(referencingInstances)

	if len(referencingInstances) < 1 {
		t.Errorf("systemIdleProcess.EnumerateReferencingInstances didn't find any Win32_SystemProcesses instances refering to Win32_Process")
		return
	}

	for _, referencingInstance := range referencingInstances {
		if referencingInstance.GetClassName() != "Win32_SystemProcesses" {
			t.Errorf("systemIdleProcess.EnumerateReferencingInstances didn't find any related instances of Win32_SystemProcesses")
			return
		}
	}
}

// SetClassName
func addClassKeyProperty(c *WmiClass, propertyName string, t *testing.T) error {
	rawResult, err := oleutil.GetProperty(c.class, "Properties_")
	if err != nil {
		t.Errorf("GetProperty \"Properties_\" failed with error %v", err)
		return err
	}

	propertiesIDispatch := rawResult.ToIDispatch()
	defer rawResult.Clear()

	addRawResult, err := oleutil.CallMethod(propertiesIDispatch, "Add", propertyName, 8)
	if err != nil {
		t.Errorf("CallMethod \"Add\" failed with error %v", err)
		return err
	}
	defer addRawResult.Clear()

	itemRawResult, err := oleutil.CallMethod(propertiesIDispatch, "Item", propertyName)
	if err != nil {
		t.Errorf("CallMethod \"Item\" failed with error %v", err)
		return err
	}
	defer itemRawResult.Clear()
	itemsIDispatch := itemRawResult.ToIDispatch()

	qualifiersRawResult, err := oleutil.GetProperty(itemsIDispatch, "Qualifiers_")
	if err != nil {
		t.Errorf("GetProperty \"Qualifiers_\" failed with error %v", err)
		return err
	}

	qualifiersIDispatch := qualifiersRawResult.ToIDispatch()
	defer qualifiersRawResult.Clear()

	addKeyRawResult, err := oleutil.CallMethod(qualifiersIDispatch, "Add", "key", true)
	if err != nil {
		t.Errorf("CallMethod \"Add\" failed with error %v", err)
		return err
	}
	defer addKeyRawResult.Clear()

	return nil
}

// This test is similar to the one illustrated here: https://docs.microsoft.com/en-us/windows/win32/wmisdk/swbemobject
func Test_NewObjects(t *testing.T) {
	sessionManager := NewWmiSessionManager()

	session, err := sessionManager.GetLocalSession("ROOT\\default")

	if err != nil {
		t.Errorf("sessionManager.GetSession failed with error %v", err)
		return
	}

	defer sessionManager.Dispose()

	connected, err := session.Connect()

	if !connected || err != nil {
		t.Errorf("session.Connect failed with error %v", err)
		return
	}

	class, err := session.CreateNewClass()
	if err != nil {
		t.Errorf("session.CreateNewClass failed with error %v", err)
		return
	}
	defer class.Close()

	err = class.SetClassName("newClass")
	if err != nil {
		t.Errorf("class.SetClassName failed with error %v", err)
		return
	}
	status := class.GetClassName()

	if status != "newClass" {
		t.Errorf("class.SetClassName didn't return the expected object [%v] != \"hello\"", status)
		return
	}

	err = addClassKeyProperty(class, "MyProperty", t)
	if err != nil {
		t.Errorf("Error: addClassKeyProperty() failed with error: %v", err)
		return
	}

	// If commit fails here, make sure you allow the current user to write into the ROOT\default WMI namespace
	// i.e. right click on the start menu -> choose "Computer management" -> select "Services and applications" -> "WMI Control" ->
	// right click on "WMI control" -> select "Properties" -> select "Security" tab -> browse to "ROOT\Default" and select the node and
	// left click on the "Security" button -> add your user alias and make sure you have full write permissions.
	err = class.Commit()
	if err != nil {
		t.Errorf("Error: class.Commit() failed with error: %v", err)
		return
	}

	err = class.Modify()
	if err != nil {
		t.Errorf("Error: class.Modify() failed with error: %v", err)
		return
	}

	freshClass, err := session.GetClass(class.GetClassName())
	if err != nil {
		t.Errorf("session.GetClass failed with error %v", err)
		return

	}
	defer freshClass.Close()

	instance, err := freshClass.MakeInstance()
	if err != nil {
		t.Errorf("class.MakeInstance failed with error %v", err)
		return
	}
	defer instance.Close()

	err = instance.SetProperty("MyProperty", "hello")
	if err != nil {
		t.Errorf("systemIdleProcess.SetProperty failed with error %v", err)
		return
	}

	err = instance.Commit()
	if err != nil {
		t.Errorf("Error: instance.Commit() failed with error: %v", err)
		return
	}

	err = instance.Modify()
	if err != nil {
		t.Errorf("Error: instance.Modify() failed with error: %v", err)
		return
	}

	err = instance.Delete()
	if err != nil {
		t.Errorf("Error: instance.Delete() failed with error: %v", err)
		return
	}
}

// func Test_ClusterWMI(t *testing.T) {
// 	sessionManager := NewWmiSessionManager()
// 	defer sessionManager.Dispose()

// 	session, err := sessionManager.GetSession("Root\\MSCluster", "", "", "", "")
// 	if err != nil {
// 		t.Errorf("sessionManager.GetSession failed with error %v", err)
// 		return
// 	}

// 	connected, err := session.Connect()

// 	if !connected || err != nil {
// 		t.Errorf("session.Connect failed with error %v", err)
// 		return
// 	}
// 	defer session.Dispose()

// 	clusterInstances, err := session.QueryInstances("SELECT * FROM MSCluster_Cluster")
// 	if err != nil {
// 		t.Errorf("session.QueryInstances failed with error %v", err)
// 		return
// 	}
// 	defer CloseAllInstances(clusterInstances)

// 	if len(clusterInstances) != 1 {
// 		t.Errorf("session.QueryInstances returned (%v) processes where it should have returned 1", len(clusterInstances))
// 		return
// 	}

// 	clusterInstance := clusterInstances[0]

// 	clusterName, err := clusterInstance.GetProperty("Name");
// 	if err != nil {
// 		t.Errorf("clusterInstance.GetProperty failed with error %v", err)
// 		return
// 	} else {
// 		t.Logf("Cluster name: %s", clusterName.(string))
// 	}

// 	nodeInstances, err := clusterInstance.GetRelated("MSCluster_Node")
// 	if err != nil {
// 		t.Errorf("session.GetRelated failed with error %v", err)
// 		return
// 	}
// 	defer CloseAllInstances(nodeInstances)

// 	t.Logf("------------------------------------------------------")

// 	for _, nodeInstance := range nodeInstances {
// 		nodeName, err := nodeInstance.GetProperty("Name");
// 		if err != nil {
// 			t.Errorf("nodeInstance.GetProperty failed with error %v", err)
// 			return
// 		} else {
// 			t.Logf("Node name: %s", nodeName.(string))
// 		}
// 	}

// 	t.Logf("------------------------------------------------------")
// 	t.Logf("Clustering ClusterTestVM")

// 	_, err = clusterInstance.InvokeMethod("AddVirtualMachine", "ClusterTestVM")
// 	if err != nil {
// 		t.Errorf("InvokeMethod \"AddVirtualMachine\" failed with error %v", err)
// 		return
// 	}

// }
