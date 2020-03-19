// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.Microsoft.SecurityClient
//////////////////////////////////////////////
package securityclient

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// ProtectionTechnologyStatus struct
type ProtectionTechnologyStatus struct {
	*SerializableToXml

	// Is protection technology enabled
	Enabled bool

	// Protection technology name
	Name string

	// Protection technology version (major, minor, build, revision)
	Version string
}

func NewProtectionTechnologyStatusEx1(instance *cim.WmiInstance) (newInstance *ProtectionTechnologyStatus, err error) {
	tmp, err := NewSerializableToXmlEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ProtectionTechnologyStatus{
		SerializableToXml: tmp,
	}
	return
}

func NewProtectionTechnologyStatusEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ProtectionTechnologyStatus, err error) {
	tmp, err := NewSerializableToXmlEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ProtectionTechnologyStatus{
		SerializableToXml: tmp,
	}
	return
}

// SetEnabled sets the value of Enabled for the instance
func (instance *ProtectionTechnologyStatus) SetPropertyEnabled(value bool) (err error) {
	return instance.SetProperty("Enabled", value)
}

// GetEnabled gets the value of Enabled for the instance
func (instance *ProtectionTechnologyStatus) GetPropertyEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("Enabled")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetName sets the value of Name for the instance
func (instance *ProtectionTechnologyStatus) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", value)
}

// GetName gets the value of Name for the instance
func (instance *ProtectionTechnologyStatus) GetPropertyName() (value string, err error) {
	retValue, err := instance.GetProperty("Name")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVersion sets the value of Version for the instance
func (instance *ProtectionTechnologyStatus) SetPropertyVersion(value string) (err error) {
	return instance.SetProperty("Version", value)
}

// GetVersion gets the value of Version for the instance
func (instance *ProtectionTechnologyStatus) GetPropertyVersion() (value string, err error) {
	retValue, err := instance.GetProperty("Version")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
