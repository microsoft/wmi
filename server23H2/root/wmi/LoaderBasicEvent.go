// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// LoaderBasicEvent struct
type LoaderBasicEvent struct {
	*Image_V2
}

func NewLoaderBasicEventEx1(instance *cim.WmiInstance) (newInstance *LoaderBasicEvent, err error) {
	tmp, err := NewImage_V2Ex1(instance)

	if err != nil {
		return
	}
	newInstance = &LoaderBasicEvent{
		Image_V2: tmp,
	}
	return
}

func NewLoaderBasicEventEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *LoaderBasicEvent, err error) {
	tmp, err := NewImage_V2Ex6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &LoaderBasicEvent{
		Image_V2: tmp,
	}
	return
}
