// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.MSCluster
//
// ////////////////////////////////////////////
package mscluster

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSCluster_EventPropertyChange struct
type MSCluster_EventPropertyChange struct {
	*MSCluster_Event

	//
	EventProperty string
}

func NewMSCluster_EventPropertyChangeEx1(instance *cim.WmiInstance) (newInstance *MSCluster_EventPropertyChange, err error) {
	tmp, err := NewMSCluster_EventEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSCluster_EventPropertyChange{
		MSCluster_Event: tmp,
	}
	return
}

func NewMSCluster_EventPropertyChangeEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSCluster_EventPropertyChange, err error) {
	tmp, err := NewMSCluster_EventEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSCluster_EventPropertyChange{
		MSCluster_Event: tmp,
	}
	return
}

// SetEventProperty sets the value of EventProperty for the instance
func (instance *MSCluster_EventPropertyChange) SetPropertyEventProperty(value string) (err error) {
	return instance.SetProperty("EventProperty", (value))
}

// GetEventProperty gets the value of EventProperty for the instance
func (instance *MSCluster_EventPropertyChange) GetPropertyEventProperty() (value string, err error) {
	retValue, err := instance.GetProperty("EventProperty")
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
