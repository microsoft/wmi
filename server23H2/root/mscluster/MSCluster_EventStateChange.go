// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.MSCluster
//////////////////////////////////////////////
package mscluster

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSCluster_EventStateChange struct
type MSCluster_EventStateChange struct {
	*MSCluster_Event

	//
	EventNewState uint32
}

func NewMSCluster_EventStateChangeEx1(instance *cim.WmiInstance) (newInstance *MSCluster_EventStateChange, err error) {
	tmp, err := NewMSCluster_EventEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSCluster_EventStateChange{
		MSCluster_Event: tmp,
	}
	return
}

func NewMSCluster_EventStateChangeEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSCluster_EventStateChange, err error) {
	tmp, err := NewMSCluster_EventEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSCluster_EventStateChange{
		MSCluster_Event: tmp,
	}
	return
}

// SetEventNewState sets the value of EventNewState for the instance
func (instance *MSCluster_EventStateChange) SetPropertyEventNewState(value uint32) (err error) {
	return instance.SetProperty("EventNewState", (value))
}

// GetEventNewState gets the value of EventNewState for the instance
func (instance *MSCluster_EventStateChange) GetPropertyEventNewState() (value uint32, err error) {
	retValue, err := instance.GetProperty("EventNewState")
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
