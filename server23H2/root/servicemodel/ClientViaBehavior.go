// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.ServiceModel
//////////////////////////////////////////////
package servicemodel

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// ClientViaBehavior struct
type ClientViaBehavior struct {
	*Behavior

	// The ViaUri.
	Uri string
}

func NewClientViaBehaviorEx1(instance *cim.WmiInstance) (newInstance *ClientViaBehavior, err error) {
	tmp, err := NewBehaviorEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ClientViaBehavior{
		Behavior: tmp,
	}
	return
}

func NewClientViaBehaviorEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ClientViaBehavior, err error) {
	tmp, err := NewBehaviorEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ClientViaBehavior{
		Behavior: tmp,
	}
	return
}

// SetUri sets the value of Uri for the instance
func (instance *ClientViaBehavior) SetPropertyUri(value string) (err error) {
	return instance.SetProperty("Uri", (value))
}

// GetUri gets the value of Uri for the instance
func (instance *ClientViaBehavior) GetPropertyUri() (value string, err error) {
	retValue, err := instance.GetProperty("Uri")
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
