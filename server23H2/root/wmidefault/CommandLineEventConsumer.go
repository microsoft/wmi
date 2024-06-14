// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.DEFAULT
//////////////////////////////////////////////
package wmidefault

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// CommandLineEventConsumer struct
type CommandLineEventConsumer struct {
	*__EventConsumer

	//
	CommandLineTemplate string

	//
	CreateNewConsole bool

	//
	CreateNewProcessGroup bool

	//
	CreateSeparateWowVdm bool

	//
	CreateSharedWowVdm bool

	//
	DesktopName string

	//
	ExecutablePath string

	//
	FillAttribute uint32

	//
	ForceOffFeedback bool

	//
	ForceOnFeedback bool

	//
	KillTimeout uint32

	//
	Name string

	//
	Priority int32

	//
	RunInteractively bool

	//
	ShowWindowCommand uint32

	//
	UseDefaultErrorMode bool

	//
	WindowTitle string

	//
	WorkingDirectory string

	//
	XCoordinate uint32

	//
	XNumCharacters uint32

	//
	XSize uint32

	//
	YCoordinate uint32

	//
	YNumCharacters uint32

	//
	YSize uint32
}

func NewCommandLineEventConsumerEx1(instance *cim.WmiInstance) (newInstance *CommandLineEventConsumer, err error) {
	tmp, err := New__EventConsumerEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CommandLineEventConsumer{
		__EventConsumer: tmp,
	}
	return
}

func NewCommandLineEventConsumerEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CommandLineEventConsumer, err error) {
	tmp, err := New__EventConsumerEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CommandLineEventConsumer{
		__EventConsumer: tmp,
	}
	return
}

// SetCommandLineTemplate sets the value of CommandLineTemplate for the instance
func (instance *CommandLineEventConsumer) SetPropertyCommandLineTemplate(value string) (err error) {
	return instance.SetProperty("CommandLineTemplate", (value))
}

// GetCommandLineTemplate gets the value of CommandLineTemplate for the instance
func (instance *CommandLineEventConsumer) GetPropertyCommandLineTemplate() (value string, err error) {
	retValue, err := instance.GetProperty("CommandLineTemplate")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetCreateNewConsole sets the value of CreateNewConsole for the instance
func (instance *CommandLineEventConsumer) SetPropertyCreateNewConsole(value bool) (err error) {
	return instance.SetProperty("CreateNewConsole", (value))
}

// GetCreateNewConsole gets the value of CreateNewConsole for the instance
func (instance *CommandLineEventConsumer) GetPropertyCreateNewConsole() (value bool, err error) {
	retValue, err := instance.GetProperty("CreateNewConsole")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetCreateNewProcessGroup sets the value of CreateNewProcessGroup for the instance
func (instance *CommandLineEventConsumer) SetPropertyCreateNewProcessGroup(value bool) (err error) {
	return instance.SetProperty("CreateNewProcessGroup", (value))
}

// GetCreateNewProcessGroup gets the value of CreateNewProcessGroup for the instance
func (instance *CommandLineEventConsumer) GetPropertyCreateNewProcessGroup() (value bool, err error) {
	retValue, err := instance.GetProperty("CreateNewProcessGroup")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetCreateSeparateWowVdm sets the value of CreateSeparateWowVdm for the instance
func (instance *CommandLineEventConsumer) SetPropertyCreateSeparateWowVdm(value bool) (err error) {
	return instance.SetProperty("CreateSeparateWowVdm", (value))
}

// GetCreateSeparateWowVdm gets the value of CreateSeparateWowVdm for the instance
func (instance *CommandLineEventConsumer) GetPropertyCreateSeparateWowVdm() (value bool, err error) {
	retValue, err := instance.GetProperty("CreateSeparateWowVdm")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetCreateSharedWowVdm sets the value of CreateSharedWowVdm for the instance
func (instance *CommandLineEventConsumer) SetPropertyCreateSharedWowVdm(value bool) (err error) {
	return instance.SetProperty("CreateSharedWowVdm", (value))
}

// GetCreateSharedWowVdm gets the value of CreateSharedWowVdm for the instance
func (instance *CommandLineEventConsumer) GetPropertyCreateSharedWowVdm() (value bool, err error) {
	retValue, err := instance.GetProperty("CreateSharedWowVdm")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetDesktopName sets the value of DesktopName for the instance
func (instance *CommandLineEventConsumer) SetPropertyDesktopName(value string) (err error) {
	return instance.SetProperty("DesktopName", (value))
}

// GetDesktopName gets the value of DesktopName for the instance
func (instance *CommandLineEventConsumer) GetPropertyDesktopName() (value string, err error) {
	retValue, err := instance.GetProperty("DesktopName")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetExecutablePath sets the value of ExecutablePath for the instance
func (instance *CommandLineEventConsumer) SetPropertyExecutablePath(value string) (err error) {
	return instance.SetProperty("ExecutablePath", (value))
}

// GetExecutablePath gets the value of ExecutablePath for the instance
func (instance *CommandLineEventConsumer) GetPropertyExecutablePath() (value string, err error) {
	retValue, err := instance.GetProperty("ExecutablePath")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetFillAttribute sets the value of FillAttribute for the instance
func (instance *CommandLineEventConsumer) SetPropertyFillAttribute(value uint32) (err error) {
	return instance.SetProperty("FillAttribute", (value))
}

// GetFillAttribute gets the value of FillAttribute for the instance
func (instance *CommandLineEventConsumer) GetPropertyFillAttribute() (value uint32, err error) {
	retValue, err := instance.GetProperty("FillAttribute")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetForceOffFeedback sets the value of ForceOffFeedback for the instance
func (instance *CommandLineEventConsumer) SetPropertyForceOffFeedback(value bool) (err error) {
	return instance.SetProperty("ForceOffFeedback", (value))
}

// GetForceOffFeedback gets the value of ForceOffFeedback for the instance
func (instance *CommandLineEventConsumer) GetPropertyForceOffFeedback() (value bool, err error) {
	retValue, err := instance.GetProperty("ForceOffFeedback")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetForceOnFeedback sets the value of ForceOnFeedback for the instance
func (instance *CommandLineEventConsumer) SetPropertyForceOnFeedback(value bool) (err error) {
	return instance.SetProperty("ForceOnFeedback", (value))
}

// GetForceOnFeedback gets the value of ForceOnFeedback for the instance
func (instance *CommandLineEventConsumer) GetPropertyForceOnFeedback() (value bool, err error) {
	retValue, err := instance.GetProperty("ForceOnFeedback")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetKillTimeout sets the value of KillTimeout for the instance
func (instance *CommandLineEventConsumer) SetPropertyKillTimeout(value uint32) (err error) {
	return instance.SetProperty("KillTimeout", (value))
}

// GetKillTimeout gets the value of KillTimeout for the instance
func (instance *CommandLineEventConsumer) GetPropertyKillTimeout() (value uint32, err error) {
	retValue, err := instance.GetProperty("KillTimeout")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetName sets the value of Name for the instance
func (instance *CommandLineEventConsumer) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *CommandLineEventConsumer) GetPropertyName() (value string, err error) {
	retValue, err := instance.GetProperty("Name")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetPriority sets the value of Priority for the instance
func (instance *CommandLineEventConsumer) SetPropertyPriority(value int32) (err error) {
	return instance.SetProperty("Priority", (value))
}

// GetPriority gets the value of Priority for the instance
func (instance *CommandLineEventConsumer) GetPropertyPriority() (value int32, err error) {
	retValue, err := instance.GetProperty("Priority")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int32(valuetmp)

	return
}

// SetRunInteractively sets the value of RunInteractively for the instance
func (instance *CommandLineEventConsumer) SetPropertyRunInteractively(value bool) (err error) {
	return instance.SetProperty("RunInteractively", (value))
}

// GetRunInteractively gets the value of RunInteractively for the instance
func (instance *CommandLineEventConsumer) GetPropertyRunInteractively() (value bool, err error) {
	retValue, err := instance.GetProperty("RunInteractively")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetShowWindowCommand sets the value of ShowWindowCommand for the instance
func (instance *CommandLineEventConsumer) SetPropertyShowWindowCommand(value uint32) (err error) {
	return instance.SetProperty("ShowWindowCommand", (value))
}

// GetShowWindowCommand gets the value of ShowWindowCommand for the instance
func (instance *CommandLineEventConsumer) GetPropertyShowWindowCommand() (value uint32, err error) {
	retValue, err := instance.GetProperty("ShowWindowCommand")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetUseDefaultErrorMode sets the value of UseDefaultErrorMode for the instance
func (instance *CommandLineEventConsumer) SetPropertyUseDefaultErrorMode(value bool) (err error) {
	return instance.SetProperty("UseDefaultErrorMode", (value))
}

// GetUseDefaultErrorMode gets the value of UseDefaultErrorMode for the instance
func (instance *CommandLineEventConsumer) GetPropertyUseDefaultErrorMode() (value bool, err error) {
	retValue, err := instance.GetProperty("UseDefaultErrorMode")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetWindowTitle sets the value of WindowTitle for the instance
func (instance *CommandLineEventConsumer) SetPropertyWindowTitle(value string) (err error) {
	return instance.SetProperty("WindowTitle", (value))
}

// GetWindowTitle gets the value of WindowTitle for the instance
func (instance *CommandLineEventConsumer) GetPropertyWindowTitle() (value string, err error) {
	retValue, err := instance.GetProperty("WindowTitle")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetWorkingDirectory sets the value of WorkingDirectory for the instance
func (instance *CommandLineEventConsumer) SetPropertyWorkingDirectory(value string) (err error) {
	return instance.SetProperty("WorkingDirectory", (value))
}

// GetWorkingDirectory gets the value of WorkingDirectory for the instance
func (instance *CommandLineEventConsumer) GetPropertyWorkingDirectory() (value string, err error) {
	retValue, err := instance.GetProperty("WorkingDirectory")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetXCoordinate sets the value of XCoordinate for the instance
func (instance *CommandLineEventConsumer) SetPropertyXCoordinate(value uint32) (err error) {
	return instance.SetProperty("XCoordinate", (value))
}

// GetXCoordinate gets the value of XCoordinate for the instance
func (instance *CommandLineEventConsumer) GetPropertyXCoordinate() (value uint32, err error) {
	retValue, err := instance.GetProperty("XCoordinate")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetXNumCharacters sets the value of XNumCharacters for the instance
func (instance *CommandLineEventConsumer) SetPropertyXNumCharacters(value uint32) (err error) {
	return instance.SetProperty("XNumCharacters", (value))
}

// GetXNumCharacters gets the value of XNumCharacters for the instance
func (instance *CommandLineEventConsumer) GetPropertyXNumCharacters() (value uint32, err error) {
	retValue, err := instance.GetProperty("XNumCharacters")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetXSize sets the value of XSize for the instance
func (instance *CommandLineEventConsumer) SetPropertyXSize(value uint32) (err error) {
	return instance.SetProperty("XSize", (value))
}

// GetXSize gets the value of XSize for the instance
func (instance *CommandLineEventConsumer) GetPropertyXSize() (value uint32, err error) {
	retValue, err := instance.GetProperty("XSize")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetYCoordinate sets the value of YCoordinate for the instance
func (instance *CommandLineEventConsumer) SetPropertyYCoordinate(value uint32) (err error) {
	return instance.SetProperty("YCoordinate", (value))
}

// GetYCoordinate gets the value of YCoordinate for the instance
func (instance *CommandLineEventConsumer) GetPropertyYCoordinate() (value uint32, err error) {
	retValue, err := instance.GetProperty("YCoordinate")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetYNumCharacters sets the value of YNumCharacters for the instance
func (instance *CommandLineEventConsumer) SetPropertyYNumCharacters(value uint32) (err error) {
	return instance.SetProperty("YNumCharacters", (value))
}

// GetYNumCharacters gets the value of YNumCharacters for the instance
func (instance *CommandLineEventConsumer) GetPropertyYNumCharacters() (value uint32, err error) {
	retValue, err := instance.GetProperty("YNumCharacters")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetYSize sets the value of YSize for the instance
func (instance *CommandLineEventConsumer) SetPropertyYSize(value uint32) (err error) {
	return instance.SetProperty("YSize", (value))
}

// GetYSize gets the value of YSize for the instance
func (instance *CommandLineEventConsumer) GetPropertyYSize() (value uint32, err error) {
	retValue, err := instance.GetProperty("YSize")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}
