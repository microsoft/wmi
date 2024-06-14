// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.directory.LDAP
//////////////////////////////////////////////
package ldap

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// ds_volume struct
type ds_volume struct {
	*ads_volume
}

func Newds_volumeEx1(instance *cim.WmiInstance) (newInstance *ds_volume, err error) {
	tmp, err := Newads_volumeEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_volume{
		ads_volume: tmp,
	}
	return
}

func Newds_volumeEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_volume, err error) {
	tmp, err := Newads_volumeEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_volume{
		ads_volume: tmp,
	}
	return
}
