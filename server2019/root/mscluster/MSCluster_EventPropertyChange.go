// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.MSCluster
//////////////////////////////////////////////
package mscluster

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
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
	return instance.SetProperty("EventProperty", value)
}

// GetEventProperty gets the value of EventProperty for the instance
func (instance *MSCluster_EventPropertyChange) GetPropertyEventProperty() (value string, err error) {
	retValue, err := instance.GetProperty("EventProperty")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
