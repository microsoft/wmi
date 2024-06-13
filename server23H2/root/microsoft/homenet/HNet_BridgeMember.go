// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.Microsoft.HomeNet
//////////////////////////////////////////////
package homenet

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// HNet_BridgeMember struct
type HNet_BridgeMember struct {
	*cim.WmiInstance

	//
	Bridge HNet_Connection

	//
	Member HNet_Connection
}

func NewHNet_BridgeMemberEx1(instance *cim.WmiInstance) (newInstance *HNet_BridgeMember, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &HNet_BridgeMember{
		WmiInstance: tmp,
	}
	return
}

func NewHNet_BridgeMemberEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *HNet_BridgeMember, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &HNet_BridgeMember{
		WmiInstance: tmp,
	}
	return
}

// SetBridge sets the value of Bridge for the instance
func (instance *HNet_BridgeMember) SetPropertyBridge(value HNet_Connection) (err error) {
	return instance.SetProperty("Bridge", (value))
}

// GetBridge gets the value of Bridge for the instance
func (instance *HNet_BridgeMember) GetPropertyBridge() (value HNet_Connection, err error) {
	retValue, err := instance.GetProperty("Bridge")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(HNet_Connection)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " HNet_Connection is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = HNet_Connection(valuetmp)

	return
}

// SetMember sets the value of Member for the instance
func (instance *HNet_BridgeMember) SetPropertyMember(value HNet_Connection) (err error) {
	return instance.SetProperty("Member", (value))
}

// GetMember gets the value of Member for the instance
func (instance *HNet_BridgeMember) GetPropertyMember() (value HNet_Connection, err error) {
	retValue, err := instance.GetProperty("Member")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(HNet_Connection)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " HNet_Connection is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = HNet_Connection(valuetmp)

	return
}
