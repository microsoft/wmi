// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.Microsoft.Windows.ClusterUpdate
//
// ////////////////////////////////////////////
package clusterupdate

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_CAURun_Report_ID struct
type MSFT_CAURun_Report_ID struct {
	*cim.WmiInstance

	//
	Timestamp string
}

func NewMSFT_CAURun_Report_IDEx1(instance *cim.WmiInstance) (newInstance *MSFT_CAURun_Report_ID, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_CAURun_Report_ID{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_CAURun_Report_IDEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_CAURun_Report_ID, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_CAURun_Report_ID{
		WmiInstance: tmp,
	}
	return
}

// SetTimestamp sets the value of Timestamp for the instance
func (instance *MSFT_CAURun_Report_ID) SetPropertyTimestamp(value string) (err error) {
	return instance.SetProperty("Timestamp", (value))
}

// GetTimestamp gets the value of Timestamp for the instance
func (instance *MSFT_CAURun_Report_ID) GetPropertyTimestamp() (value string, err error) {
	retValue, err := instance.GetProperty("Timestamp")
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
