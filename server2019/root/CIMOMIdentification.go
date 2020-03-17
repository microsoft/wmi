// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root
//////////////////////////////////////////////
package root

// __CIMOMIdentification struct
type __CIMOMIdentification struct {
	__SystemClass

	//
	SetupDateTime string

	//
	VersionCurrentlyRunning string

	//
	VersionUsedToCreateDB string

	//
	WorkingDirectory string
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
