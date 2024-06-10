// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.Microsoft.Windows.RemoteAccess
//
// ////////////////////////////////////////////
package remoteaccess

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// RemoteAccessHealthHeuristic struct
type RemoteAccessHealthHeuristic struct {
	*cim.WmiInstance

	//
	ErrorCause string

	//
	ErrorDesc string

	//
	ErrorResoln string

	//
	Id uint32

	//
	OperationStatus string

	//
	Status string
}

func NewRemoteAccessHealthHeuristicEx1(instance *cim.WmiInstance) (newInstance *RemoteAccessHealthHeuristic, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &RemoteAccessHealthHeuristic{
		WmiInstance: tmp,
	}
	return
}

func NewRemoteAccessHealthHeuristicEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RemoteAccessHealthHeuristic, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RemoteAccessHealthHeuristic{
		WmiInstance: tmp,
	}
	return
}

// SetErrorCause sets the value of ErrorCause for the instance
func (instance *RemoteAccessHealthHeuristic) SetPropertyErrorCause(value string) (err error) {
	return instance.SetProperty("ErrorCause", (value))
}

// GetErrorCause gets the value of ErrorCause for the instance
func (instance *RemoteAccessHealthHeuristic) GetPropertyErrorCause() (value string, err error) {
	retValue, err := instance.GetProperty("ErrorCause")
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

// SetErrorDesc sets the value of ErrorDesc for the instance
func (instance *RemoteAccessHealthHeuristic) SetPropertyErrorDesc(value string) (err error) {
	return instance.SetProperty("ErrorDesc", (value))
}

// GetErrorDesc gets the value of ErrorDesc for the instance
func (instance *RemoteAccessHealthHeuristic) GetPropertyErrorDesc() (value string, err error) {
	retValue, err := instance.GetProperty("ErrorDesc")
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

// SetErrorResoln sets the value of ErrorResoln for the instance
func (instance *RemoteAccessHealthHeuristic) SetPropertyErrorResoln(value string) (err error) {
	return instance.SetProperty("ErrorResoln", (value))
}

// GetErrorResoln gets the value of ErrorResoln for the instance
func (instance *RemoteAccessHealthHeuristic) GetPropertyErrorResoln() (value string, err error) {
	retValue, err := instance.GetProperty("ErrorResoln")
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

// SetId sets the value of Id for the instance
func (instance *RemoteAccessHealthHeuristic) SetPropertyId(value uint32) (err error) {
	return instance.SetProperty("Id", (value))
}

// GetId gets the value of Id for the instance
func (instance *RemoteAccessHealthHeuristic) GetPropertyId() (value uint32, err error) {
	retValue, err := instance.GetProperty("Id")
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

// SetOperationStatus sets the value of OperationStatus for the instance
func (instance *RemoteAccessHealthHeuristic) SetPropertyOperationStatus(value string) (err error) {
	return instance.SetProperty("OperationStatus", (value))
}

// GetOperationStatus gets the value of OperationStatus for the instance
func (instance *RemoteAccessHealthHeuristic) GetPropertyOperationStatus() (value string, err error) {
	retValue, err := instance.GetProperty("OperationStatus")
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

// SetStatus sets the value of Status for the instance
func (instance *RemoteAccessHealthHeuristic) SetPropertyStatus(value string) (err error) {
	return instance.SetProperty("Status", (value))
}

// GetStatus gets the value of Status for the instance
func (instance *RemoteAccessHealthHeuristic) GetPropertyStatus() (value string, err error) {
	retValue, err := instance.GetProperty("Status")
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
