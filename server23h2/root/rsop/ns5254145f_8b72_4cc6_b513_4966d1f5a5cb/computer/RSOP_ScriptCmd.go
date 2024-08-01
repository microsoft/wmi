// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.RSOP.NS5254145F_8B72_4CC6_B513_4966D1F5A5CB.Computer
//////////////////////////////////////////////
package computer

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// RSOP_ScriptCmd struct
type RSOP_ScriptCmd struct {
	*cim.WmiInstance

	//
	arguments string

	//
	executionTime string

	//
	script string
}

func NewRSOP_ScriptCmdEx1(instance *cim.WmiInstance) (newInstance *RSOP_ScriptCmd, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &RSOP_ScriptCmd{
		WmiInstance: tmp,
	}
	return
}

func NewRSOP_ScriptCmdEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RSOP_ScriptCmd, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RSOP_ScriptCmd{
		WmiInstance: tmp,
	}
	return
}

// Setarguments sets the value of arguments for the instance
func (instance *RSOP_ScriptCmd) SetPropertyarguments(value string) (err error) {
	return instance.SetProperty("arguments", (value))
}

// Getarguments gets the value of arguments for the instance
func (instance *RSOP_ScriptCmd) GetPropertyarguments() (value string, err error) {
	retValue, err := instance.GetProperty("arguments")
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

// SetexecutionTime sets the value of executionTime for the instance
func (instance *RSOP_ScriptCmd) SetPropertyexecutionTime(value string) (err error) {
	return instance.SetProperty("executionTime", (value))
}

// GetexecutionTime gets the value of executionTime for the instance
func (instance *RSOP_ScriptCmd) GetPropertyexecutionTime() (value string, err error) {
	retValue, err := instance.GetProperty("executionTime")
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

// Setscript sets the value of script for the instance
func (instance *RSOP_ScriptCmd) SetPropertyscript(value string) (err error) {
	return instance.SetProperty("script", (value))
}

// Getscript gets the value of script for the instance
func (instance *RSOP_ScriptCmd) GetPropertyscript() (value string, err error) {
	retValue, err := instance.GetProperty("script")
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
