// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package cim

import (
	"os"
	"strings"
	"testing"
)

func Test_WmiClass(t *testing.T) {
	sessionManager := NewWmiSessionManager()

	session, err := sessionManager.GetLocalSession("ROOT\\CimV2")

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

	systemIdleProcessInstance := systemIdleProcessInstances[0]
	defer systemIdleProcessInstance.Close()

	property, err := systemIdleProcessInstance.GetSystemProperty("__CLASS")
	if err != nil {
		t.Errorf("systemIdleProcessInstance.GetSystemProperty failed with error %v", err)
		return
	}
	defer property.Close()

	if property.Value().(string) != "Win32_Process" {
		t.Errorf("systemIdleProcessInstance.GetSystemProperty(\"__CLASS\") returned a wrong item: %s", property.Value().(string))
		return
	}

	class := systemIdleProcessInstance.GetClass()
	defer class.Close()

	instance, err := class.MakeInstance()
	if err != nil {
		t.Errorf("systemIdleProcessInstance.MakeInstance failed with error %v", err)
		return
	}
	defer instance.Close()

	methods := class.GetMethodsNames()
	_, exists := FindStringInSlice(methods, "Terminate")
	if !exists {
		t.Errorf("class.GetMethods failed as it did not return the \"Terminate\" method.")
	}

	properties := class.GetPropertiesNames()
	_, exists = FindStringInSlice(properties, "CommandLine")
	if !exists {
		t.Errorf("class.GetProperties failed as it did not return the \"CommandLine\" property.")
	}

	qualifiers := class.GetQualifiersNames()
	_, exists = FindStringInSlice(qualifiers, "CreateBy")
	if !exists {
		t.Errorf("class.GetQualifiers failed as it did not return the \"CreateBy\" qualifier.")
	}

	derivation := class.GetDerivation()
	_, exists = FindStringInSlice(derivation, "CIM_Process")
	if !exists {
		t.Errorf("class.GetDerivation failed as it did not return the \"CIM_Process\" parent class.")
	}

	className := class.GetClassName()
	if className != "Win32_Process" {
		t.Errorf("class.GetClassName failed as it did not return the \"Win32_Process\" class name. Returned: %s", className)
	}

	superClassName := class.GetSuperClassName()
	if superClassName != "CIM_Process" {
		t.Errorf("class.GetSuperClassName failed as it did not return the \"CIM_Process\" super class name. Returned: %s", superClassName)
	}

	hostname, err := os.Hostname()
	if err != nil {
		t.Errorf("os.Hostname failed with error %v", err)
		return
	}

	serverName := class.GetServerName()
	if !strings.EqualFold(serverName, hostname) {
		t.Errorf("class.GetServerName failed as it did not return the correct host name. (Expected: %s, Returned: %s)", hostname, serverName)
	}

	namespace := class.GetNamespace()
	if !strings.EqualFold(namespace, "ROOT\\CimV2") {
		t.Errorf("class.GetNamespace failed as it did not return the \"ROOT\\CimV2\" namespace. Returned: %s", namespace)
	}

	superClass := class.GetSuperClass()
	defer superClass.Close()

	superClassName2 := superClass.GetClassName()
	if superClassName != superClassName2 {
		t.Errorf("superClass failed as it did not return the \"CIM_Process\" class.")
	}
}
