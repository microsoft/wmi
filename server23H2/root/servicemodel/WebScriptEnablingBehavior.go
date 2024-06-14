// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.ServiceModel
//////////////////////////////////////////////
package servicemodel

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// WebScriptEnablingBehavior struct
type WebScriptEnablingBehavior struct {
	*WebHttpBehavior
}

func NewWebScriptEnablingBehaviorEx1(instance *cim.WmiInstance) (newInstance *WebScriptEnablingBehavior, err error) {
	tmp, err := NewWebHttpBehaviorEx1(instance)

	if err != nil {
		return
	}
	newInstance = &WebScriptEnablingBehavior{
		WebHttpBehavior: tmp,
	}
	return
}

func NewWebScriptEnablingBehaviorEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *WebScriptEnablingBehavior, err error) {
	tmp, err := NewWebHttpBehaviorEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &WebScriptEnablingBehavior{
		WebHttpBehavior: tmp,
	}
	return
}
