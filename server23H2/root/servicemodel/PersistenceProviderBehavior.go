// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.ServiceModel
//////////////////////////////////////////////
package servicemodel

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// PersistenceProviderBehavior struct
type PersistenceProviderBehavior struct {
	*Behavior

	// Specifies the interval after which persistence operations are considered timed out.
	PersistenceOperationTimeout string

	// Specifies the CLR type of the instance of persistence provider factory configured for durable services.
	PersistenceProviderFactoryType string
}

func NewPersistenceProviderBehaviorEx1(instance *cim.WmiInstance) (newInstance *PersistenceProviderBehavior, err error) {
	tmp, err := NewBehaviorEx1(instance)

	if err != nil {
		return
	}
	newInstance = &PersistenceProviderBehavior{
		Behavior: tmp,
	}
	return
}

func NewPersistenceProviderBehaviorEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *PersistenceProviderBehavior, err error) {
	tmp, err := NewBehaviorEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &PersistenceProviderBehavior{
		Behavior: tmp,
	}
	return
}

// SetPersistenceOperationTimeout sets the value of PersistenceOperationTimeout for the instance
func (instance *PersistenceProviderBehavior) SetPropertyPersistenceOperationTimeout(value string) (err error) {
	return instance.SetProperty("PersistenceOperationTimeout", (value))
}

// GetPersistenceOperationTimeout gets the value of PersistenceOperationTimeout for the instance
func (instance *PersistenceProviderBehavior) GetPropertyPersistenceOperationTimeout() (value string, err error) {
	retValue, err := instance.GetProperty("PersistenceOperationTimeout")
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

// SetPersistenceProviderFactoryType sets the value of PersistenceProviderFactoryType for the instance
func (instance *PersistenceProviderBehavior) SetPropertyPersistenceProviderFactoryType(value string) (err error) {
	return instance.SetProperty("PersistenceProviderFactoryType", (value))
}

// GetPersistenceProviderFactoryType gets the value of PersistenceProviderFactoryType for the instance
func (instance *PersistenceProviderBehavior) GetPropertyPersistenceProviderFactoryType() (value string, err error) {
	retValue, err := instance.GetProperty("PersistenceProviderFactoryType")
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
