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

// WebProxyScriptSettings struct
type WebProxyScriptSettings struct {
	*EmbeddedObject

	//
	DownloadTimeout string
}

func NewWebProxyScriptSettingsEx1(instance *cim.WmiInstance) (newInstance *WebProxyScriptSettings, err error) {
	tmp, err := NewEmbeddedObjectEx1(instance)

	if err != nil {
		return
	}
	newInstance = &WebProxyScriptSettings{
		EmbeddedObject: tmp,
	}
	return
}

func NewWebProxyScriptSettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *WebProxyScriptSettings, err error) {
	tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &WebProxyScriptSettings{
		EmbeddedObject: tmp,
	}
	return
}

// SetDownloadTimeout sets the value of DownloadTimeout for the instance
func (instance *WebProxyScriptSettings) SetPropertyDownloadTimeout(value string) (err error) {
	return instance.SetProperty("DownloadTimeout", (value))
}

// GetDownloadTimeout gets the value of DownloadTimeout for the instance
func (instance *WebProxyScriptSettings) GetPropertyDownloadTimeout() (value string, err error) {
	retValue, err := instance.GetProperty("DownloadTimeout")
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
