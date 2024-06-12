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

// Msft_MiQuery struct
type Msft_MiQuery struct {
	*Msft_MiStream

	//
	Dialect string

	//
	Expression string

	//
	NamespaceName string
}

func NewMsft_MiQueryEx1(instance *cim.WmiInstance) (newInstance *Msft_MiQuery, err error) {
	tmp, err := NewMsft_MiStreamEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msft_MiQuery{
		Msft_MiStream: tmp,
	}
	return
}

func NewMsft_MiQueryEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msft_MiQuery, err error) {
	tmp, err := NewMsft_MiStreamEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msft_MiQuery{
		Msft_MiStream: tmp,
	}
	return
}

// SetDialect sets the value of Dialect for the instance
func (instance *Msft_MiQuery) SetPropertyDialect(value string) (err error) {
	return instance.SetProperty("Dialect", (value))
}

// GetDialect gets the value of Dialect for the instance
func (instance *Msft_MiQuery) GetPropertyDialect() (value string, err error) {
	retValue, err := instance.GetProperty("Dialect")
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

// SetExpression sets the value of Expression for the instance
func (instance *Msft_MiQuery) SetPropertyExpression(value string) (err error) {
	return instance.SetProperty("Expression", (value))
}

// GetExpression gets the value of Expression for the instance
func (instance *Msft_MiQuery) GetPropertyExpression() (value string, err error) {
	retValue, err := instance.GetProperty("Expression")
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

// SetNamespaceName sets the value of NamespaceName for the instance
func (instance *Msft_MiQuery) SetPropertyNamespaceName(value string) (err error) {
	return instance.SetProperty("NamespaceName", (value))
}

// GetNamespaceName gets the value of NamespaceName for the instance
func (instance *Msft_MiQuery) GetPropertyNamespaceName() (value string, err error) {
	retValue, err := instance.GetProperty("NamespaceName")
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
