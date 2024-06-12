// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.InventoryLogging
//////////////////////////////////////////////
package inventorylogging

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Msft_MiMerge struct
type Msft_MiMerge struct {
	*Msft_MiStream

	//
	Inputs []Msft_MiStream
}

func NewMsft_MiMergeEx1(instance *cim.WmiInstance) (newInstance *Msft_MiMerge, err error) {
	tmp, err := NewMsft_MiStreamEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msft_MiMerge{
		Msft_MiStream: tmp,
	}
	return
}

func NewMsft_MiMergeEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msft_MiMerge, err error) {
	tmp, err := NewMsft_MiStreamEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msft_MiMerge{
		Msft_MiStream: tmp,
	}
	return
}

// SetInputs sets the value of Inputs for the instance
func (instance *Msft_MiMerge) SetPropertyInputs(value []Msft_MiStream) (err error) {
	return instance.SetProperty("Inputs", (value))
}

// GetInputs gets the value of Inputs for the instance
func (instance *Msft_MiMerge) GetPropertyInputs() (value []Msft_MiStream, err error) {
	retValue, err := instance.GetProperty("Inputs")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(Msft_MiStream)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " Msft_MiStream is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, Msft_MiStream(valuetmp))
	}

	return
}
