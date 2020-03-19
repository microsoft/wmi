// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.InventoryLogging
//////////////////////////////////////////////
package inventorylogging

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Msft_MiCompareSuppression struct
type Msft_MiCompareSuppression struct {
	*cim.WmiInstance

	//
	SuppressionSignal []interface{}

	//
	Timestamp string
}

func NewMsft_MiCompareSuppressionEx1(instance *cim.WmiInstance) (newInstance *Msft_MiCompareSuppression, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &Msft_MiCompareSuppression{
		WmiInstance: tmp,
	}
	return
}

func NewMsft_MiCompareSuppressionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msft_MiCompareSuppression, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msft_MiCompareSuppression{
		WmiInstance: tmp,
	}
	return
}

// SetSuppressionSignal sets the value of SuppressionSignal for the instance
func (instance *Msft_MiCompareSuppression) SetPropertySuppressionSignal(value []interface{}) (err error) {
	return instance.SetProperty("SuppressionSignal", value)
}

// GetSuppressionSignal gets the value of SuppressionSignal for the instance
func (instance *Msft_MiCompareSuppression) GetPropertySuppressionSignal() (value []interface{}, err error) {
	retValue, err := instance.GetProperty("SuppressionSignal")
	if err != nil {
		return
	}
	value, ok := retValue.([]interface{})
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTimestamp sets the value of Timestamp for the instance
func (instance *Msft_MiCompareSuppression) SetPropertyTimestamp(value string) (err error) {
	return instance.SetProperty("Timestamp", value)
}

// GetTimestamp gets the value of Timestamp for the instance
func (instance *Msft_MiCompareSuppression) GetPropertyTimestamp() (value string, err error) {
	retValue, err := instance.GetProperty("Timestamp")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
