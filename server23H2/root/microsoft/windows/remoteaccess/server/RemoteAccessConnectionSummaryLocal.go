// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.Microsoft.Windows.RemoteAccess.Server
//
// ////////////////////////////////////////////
package server

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// RemoteAccessConnectionSummaryLocal struct
type RemoteAccessConnectionSummaryLocal struct {
	*cim.WmiInstance

	//
	TotalUniqueUsers uint32
}

func NewRemoteAccessConnectionSummaryLocalEx1(instance *cim.WmiInstance) (newInstance *RemoteAccessConnectionSummaryLocal, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &RemoteAccessConnectionSummaryLocal{
		WmiInstance: tmp,
	}
	return
}

func NewRemoteAccessConnectionSummaryLocalEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RemoteAccessConnectionSummaryLocal, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RemoteAccessConnectionSummaryLocal{
		WmiInstance: tmp,
	}
	return
}

// SetTotalUniqueUsers sets the value of TotalUniqueUsers for the instance
func (instance *RemoteAccessConnectionSummaryLocal) SetPropertyTotalUniqueUsers(value uint32) (err error) {
	return instance.SetProperty("TotalUniqueUsers", (value))
}

// GetTotalUniqueUsers gets the value of TotalUniqueUsers for the instance
func (instance *RemoteAccessConnectionSummaryLocal) GetPropertyTotalUniqueUsers() (value uint32, err error) {
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
