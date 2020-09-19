// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.MSCluster
//////////////////////////////////////////////
package mscluster

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSCluster_EventResourceStateChange struct
type MSCluster_EventResourceStateChange struct {
	*MSCluster_EventStateChange

	//
	EventGroup string

	//
	EventNode string
}

func NewMSCluster_EventResourceStateChangeEx1(instance *cim.WmiInstance) (newInstance *MSCluster_EventResourceStateChange, err error) {
	tmp, err := NewMSCluster_EventStateChangeEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSCluster_EventResourceStateChange{
		MSCluster_EventStateChange: tmp,
	}
	return
}

func NewMSCluster_EventResourceStateChangeEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSCluster_EventResourceStateChange, err error) {
	tmp, err := NewMSCluster_EventStateChangeEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSCluster_EventResourceStateChange{
		MSCluster_EventStateChange: tmp,
	}
	return
}

// SetEventGroup sets the value of EventGroup for the instance
func (instance *MSCluster_EventResourceStateChange) SetPropertyEventGroup(value string) (err error) {
	return instance.SetProperty("EventGroup", (value))
}

// GetEventGroup gets the value of EventGroup for the instance
func (instance *MSCluster_EventResourceStateChange) GetPropertyEventGroup() (value string, err error) {
	retValue, err := instance.GetProperty("EventGroup")
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

// SetEventNode sets the value of EventNode for the instance
func (instance *MSCluster_EventResourceStateChange) SetPropertyEventNode(value string) (err error) {
	return instance.SetProperty("EventNode", (value))
}

// GetEventNode gets the value of EventNode for the instance
func (instance *MSCluster_EventResourceStateChange) GetPropertyEventNode() (value string, err error) {
	retValue, err := instance.GetProperty("EventNode")
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
