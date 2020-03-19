// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root
//////////////////////////////////////////////
package root

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// __CIMOMIdentification struct
type __CIMOMIdentification struct {
	*__SystemClass

	//
	SetupDateTime string

	//
	VersionCurrentlyRunning string

	//
	VersionUsedToCreateDB string

	//
	WorkingDirectory string
}

func New__CIMOMIdentificationEx1(instance *cim.WmiInstance) (newInstance *__CIMOMIdentification, err error) {
	tmp, err := New__SystemClassEx1(instance)

	if err != nil {
		return
	}
	newInstance = &__CIMOMIdentification{
		__SystemClass: tmp,
	}
	return
}

func New__CIMOMIdentificationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *__CIMOMIdentification, err error) {
	tmp, err := New__SystemClassEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &__CIMOMIdentification{
		__SystemClass: tmp,
	}
	return
}

// SetSetupDateTime sets the value of SetupDateTime for the instance
func (instance *__CIMOMIdentification) SetPropertySetupDateTime(value string) (err error) {
	return instance.SetProperty("SetupDateTime", value)
}

// GetSetupDateTime gets the value of SetupDateTime for the instance
func (instance *__CIMOMIdentification) GetPropertySetupDateTime() (value string, err error) {
	retValue, err := instance.GetProperty("SetupDateTime")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVersionCurrentlyRunning sets the value of VersionCurrentlyRunning for the instance
func (instance *__CIMOMIdentification) SetPropertyVersionCurrentlyRunning(value string) (err error) {
	return instance.SetProperty("VersionCurrentlyRunning", value)
}

// GetVersionCurrentlyRunning gets the value of VersionCurrentlyRunning for the instance
func (instance *__CIMOMIdentification) GetPropertyVersionCurrentlyRunning() (value string, err error) {
	retValue, err := instance.GetProperty("VersionCurrentlyRunning")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVersionUsedToCreateDB sets the value of VersionUsedToCreateDB for the instance
func (instance *__CIMOMIdentification) SetPropertyVersionUsedToCreateDB(value string) (err error) {
	return instance.SetProperty("VersionUsedToCreateDB", value)
}

// GetVersionUsedToCreateDB gets the value of VersionUsedToCreateDB for the instance
func (instance *__CIMOMIdentification) GetPropertyVersionUsedToCreateDB() (value string, err error) {
	retValue, err := instance.GetProperty("VersionUsedToCreateDB")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetWorkingDirectory sets the value of WorkingDirectory for the instance
func (instance *__CIMOMIdentification) SetPropertyWorkingDirectory(value string) (err error) {
	return instance.SetProperty("WorkingDirectory", value)
}

// GetWorkingDirectory gets the value of WorkingDirectory for the instance
func (instance *__CIMOMIdentification) GetPropertyWorkingDirectory() (value string, err error) {
	retValue, err := instance.GetProperty("WorkingDirectory")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
