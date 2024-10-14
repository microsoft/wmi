// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.Microsoft.Windows.TaskScheduler
//////////////////////////////////////////////
package taskscheduler

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_ScheduledTask struct
type MSFT_ScheduledTask struct {
	*cim.WmiInstance

	//
	Actions []MSFT_TaskAction

	//
	Author string

	//
	Date string

	//
	Description string

	//
	Documentation string

	//
	Principal MSFT_TaskPrincipal

	//
	SecurityDescriptor string

	//
	Settings MSFT_TaskSettings

	//
	Source string

	//
	State ScheduledTask_State

	//
	TaskName string

	//
	TaskPath string

	//
	Triggers []MSFT_TaskTrigger

	//
	URI string

	//
	Version string
}

func NewMSFT_ScheduledTaskEx1(instance *cim.WmiInstance) (newInstance *MSFT_ScheduledTask, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_ScheduledTask{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_ScheduledTaskEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_ScheduledTask, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_ScheduledTask{
		WmiInstance: tmp,
	}
	return
}

// SetActions sets the value of Actions for the instance
func (instance *MSFT_ScheduledTask) SetPropertyActions(value []MSFT_TaskAction) (err error) {
	return instance.SetProperty("Actions", (value))
}

// GetActions gets the value of Actions for the instance
func (instance *MSFT_ScheduledTask) GetPropertyActions() (value []MSFT_TaskAction, err error) {
	retValue, err := instance.GetProperty("Actions")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(MSFT_TaskAction)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " MSFT_TaskAction is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, MSFT_TaskAction(valuetmp))
	}

	return
}

// SetAuthor sets the value of Author for the instance
func (instance *MSFT_ScheduledTask) SetPropertyAuthor(value string) (err error) {
	return instance.SetProperty("Author", (value))
}

// GetAuthor gets the value of Author for the instance
func (instance *MSFT_ScheduledTask) GetPropertyAuthor() (value string, err error) {
	retValue, err := instance.GetProperty("Author")
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

// SetDate sets the value of Date for the instance
func (instance *MSFT_ScheduledTask) SetPropertyDate(value string) (err error) {
	return instance.SetProperty("Date", (value))
}

// GetDate gets the value of Date for the instance
func (instance *MSFT_ScheduledTask) GetPropertyDate() (value string, err error) {
	retValue, err := instance.GetProperty("Date")
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

// SetDescription sets the value of Description for the instance
func (instance *MSFT_ScheduledTask) SetPropertyDescription(value string) (err error) {
	return instance.SetProperty("Description", (value))
}

// GetDescription gets the value of Description for the instance
func (instance *MSFT_ScheduledTask) GetPropertyDescription() (value string, err error) {
	retValue, err := instance.GetProperty("Description")
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

// SetDocumentation sets the value of Documentation for the instance
func (instance *MSFT_ScheduledTask) SetPropertyDocumentation(value string) (err error) {
	return instance.SetProperty("Documentation", (value))
}

// GetDocumentation gets the value of Documentation for the instance
func (instance *MSFT_ScheduledTask) GetPropertyDocumentation() (value string, err error) {
	retValue, err := instance.GetProperty("Documentation")
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

// SetPrincipal sets the value of Principal for the instance
func (instance *MSFT_ScheduledTask) SetPropertyPrincipal(value MSFT_TaskPrincipal) (err error) {
	return instance.SetProperty("Principal", (value))
}

// GetPrincipal gets the value of Principal for the instance
func (instance *MSFT_ScheduledTask) GetPropertyPrincipal() (value MSFT_TaskPrincipal, err error) {
	retValue, err := instance.GetProperty("Principal")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(MSFT_TaskPrincipal)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " MSFT_TaskPrincipal is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = MSFT_TaskPrincipal(valuetmp)

	return
}

// SetSecurityDescriptor sets the value of SecurityDescriptor for the instance
func (instance *MSFT_ScheduledTask) SetPropertySecurityDescriptor(value string) (err error) {
	return instance.SetProperty("SecurityDescriptor", (value))
}

// GetSecurityDescriptor gets the value of SecurityDescriptor for the instance
func (instance *MSFT_ScheduledTask) GetPropertySecurityDescriptor() (value string, err error) {
	retValue, err := instance.GetProperty("SecurityDescriptor")
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

// SetSettings sets the value of Settings for the instance
func (instance *MSFT_ScheduledTask) SetPropertySettings(value MSFT_TaskSettings) (err error) {
	return instance.SetProperty("Settings", (value))
}

// GetSettings gets the value of Settings for the instance
func (instance *MSFT_ScheduledTask) GetPropertySettings() (value MSFT_TaskSettings, err error) {
	retValue, err := instance.GetProperty("Settings")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(MSFT_TaskSettings)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " MSFT_TaskSettings is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = MSFT_TaskSettings(valuetmp)

	return
}

// SetSource sets the value of Source for the instance
func (instance *MSFT_ScheduledTask) SetPropertySource(value string) (err error) {
	return instance.SetProperty("Source", (value))
}

// GetSource gets the value of Source for the instance
func (instance *MSFT_ScheduledTask) GetPropertySource() (value string, err error) {
	retValue, err := instance.GetProperty("Source")
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

// SetState sets the value of State for the instance
func (instance *MSFT_ScheduledTask) SetPropertyState(value ScheduledTask_State) (err error) {
	return instance.SetProperty("State", (value))
}

// GetState gets the value of State for the instance
func (instance *MSFT_ScheduledTask) GetPropertyState() (value ScheduledTask_State, err error) {
	retValue, err := instance.GetProperty("State")
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

	value = ScheduledTask_State(valuetmp)

	return
}

// SetTaskName sets the value of TaskName for the instance
func (instance *MSFT_ScheduledTask) SetPropertyTaskName(value string) (err error) {
	return instance.SetProperty("TaskName", (value))
}

// GetTaskName gets the value of TaskName for the instance
func (instance *MSFT_ScheduledTask) GetPropertyTaskName() (value string, err error) {
	retValue, err := instance.GetProperty("TaskName")
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

// SetTaskPath sets the value of TaskPath for the instance
func (instance *MSFT_ScheduledTask) SetPropertyTaskPath(value string) (err error) {
	return instance.SetProperty("TaskPath", (value))
}

// GetTaskPath gets the value of TaskPath for the instance
func (instance *MSFT_ScheduledTask) GetPropertyTaskPath() (value string, err error) {
	retValue, err := instance.GetProperty("TaskPath")
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

// SetTriggers sets the value of Triggers for the instance
func (instance *MSFT_ScheduledTask) SetPropertyTriggers(value []MSFT_TaskTrigger) (err error) {
	return instance.SetProperty("Triggers", (value))
}

// GetTriggers gets the value of Triggers for the instance
func (instance *MSFT_ScheduledTask) GetPropertyTriggers() (value []MSFT_TaskTrigger, err error) {
	retValue, err := instance.GetProperty("Triggers")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(MSFT_TaskTrigger)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " MSFT_TaskTrigger is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, MSFT_TaskTrigger(valuetmp))
	}

	return
}

// SetURI sets the value of URI for the instance
func (instance *MSFT_ScheduledTask) SetPropertyURI(value string) (err error) {
	return instance.SetProperty("URI", (value))
}

// GetURI gets the value of URI for the instance
func (instance *MSFT_ScheduledTask) GetPropertyURI() (value string, err error) {
	retValue, err := instance.GetProperty("URI")
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

// SetVersion sets the value of Version for the instance
func (instance *MSFT_ScheduledTask) SetPropertyVersion(value string) (err error) {
	return instance.SetProperty("Version", (value))
}

// GetVersion gets the value of Version for the instance
func (instance *MSFT_ScheduledTask) GetPropertyVersion() (value string, err error) {
	retValue, err := instance.GetProperty("Version")
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
