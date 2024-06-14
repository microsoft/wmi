// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.WebAdministration
//////////////////////////////////////////////
package webadministration

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// SqlCacheDependencyDatabaseSettings struct
type SqlCacheDependencyDatabaseSettings struct {
	*EmbeddedObject

	//
	Databases []SqlCacheDependencyDatabase
}

func NewSqlCacheDependencyDatabaseSettingsEx1(instance *cim.WmiInstance) (newInstance *SqlCacheDependencyDatabaseSettings, err error) {
	tmp, err := NewEmbeddedObjectEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SqlCacheDependencyDatabaseSettings{
		EmbeddedObject: tmp,
	}
	return
}

func NewSqlCacheDependencyDatabaseSettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SqlCacheDependencyDatabaseSettings, err error) {
	tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SqlCacheDependencyDatabaseSettings{
		EmbeddedObject: tmp,
	}
	return
}

// SetDatabases sets the value of Databases for the instance
func (instance *SqlCacheDependencyDatabaseSettings) SetPropertyDatabases(value []SqlCacheDependencyDatabase) (err error) {
	return instance.SetProperty("Databases", (value))
}

// GetDatabases gets the value of Databases for the instance
func (instance *SqlCacheDependencyDatabaseSettings) GetPropertyDatabases() (value []SqlCacheDependencyDatabase, err error) {
	retValue, err := instance.GetProperty("Databases")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(SqlCacheDependencyDatabase)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " SqlCacheDependencyDatabase is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, SqlCacheDependencyDatabase(valuetmp))
	}

	return
}
