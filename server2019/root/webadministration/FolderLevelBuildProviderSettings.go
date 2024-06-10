// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.WebAdministration
//
// ////////////////////////////////////////////
package webadministration

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// FolderLevelBuildProviderSettings struct
type FolderLevelBuildProviderSettings struct {
	*EmbeddedObject

	//
	FolderLevelBuildProviders []FolderLevelBuildProvider
}

func NewFolderLevelBuildProviderSettingsEx1(instance *cim.WmiInstance) (newInstance *FolderLevelBuildProviderSettings, err error) {
	tmp, err := NewEmbeddedObjectEx1(instance)

	if err != nil {
		return
	}
	newInstance = &FolderLevelBuildProviderSettings{
		EmbeddedObject: tmp,
	}
	return
}

func NewFolderLevelBuildProviderSettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *FolderLevelBuildProviderSettings, err error) {
	tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &FolderLevelBuildProviderSettings{
		EmbeddedObject: tmp,
	}
	return
}

// SetFolderLevelBuildProviders sets the value of FolderLevelBuildProviders for the instance
func (instance *FolderLevelBuildProviderSettings) SetPropertyFolderLevelBuildProviders(value []FolderLevelBuildProvider) (err error) {
	return instance.SetProperty("FolderLevelBuildProviders", (value))
}

// GetFolderLevelBuildProviders gets the value of FolderLevelBuildProviders for the instance
func (instance *FolderLevelBuildProviderSettings) GetPropertyFolderLevelBuildProviders() (value []FolderLevelBuildProvider, err error) {
	retValue, err := instance.GetProperty("FolderLevelBuildProviders")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(FolderLevelBuildProvider)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " FolderLevelBuildProvider is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, FolderLevelBuildProvider(valuetmp))
	}

	return
}
