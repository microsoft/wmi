// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.Microsoft.Windows.RemoteAccess
//////////////////////////////////////////////
package remoteaccess

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// RemoteAccessConnectionSummary struct
type RemoteAccessConnectionSummary struct {
	*cim.WmiInstance

	//
	TotalUniqueUsers uint32
}

func NewRemoteAccessConnectionSummaryEx1(instance *cim.WmiInstance) (newInstance *RemoteAccessConnectionSummary, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &RemoteAccessConnectionSummary{
		WmiInstance: tmp,
	}
	return
}

func NewRemoteAccessConnectionSummaryEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RemoteAccessConnectionSummary, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RemoteAccessConnectionSummary{
		WmiInstance: tmp,
	}
	return
}

// SetTotalUniqueUsers sets the value of TotalUniqueUsers for the instance
func (instance *RemoteAccessConnectionSummary) SetPropertyTotalUniqueUsers(value uint32) (err error) {
	return instance.SetProperty("TotalUniqueUsers", (value))
}

// GetTotalUniqueUsers gets the value of TotalUniqueUsers for the instance
func (instance *RemoteAccessConnectionSummary) GetPropertyTotalUniqueUsers() (value uint32, err error) {
	retValue, err := instance.GetProperty("TotalUniqueUsers")
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
