// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.Microsoft.Windows.Storage.Providers_v2
//////////////////////////////////////////////
package providers_v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// WSP_FileServerToVolume struct
type WSP_FileServerToVolume struct {
	*MSFT_FileServerToVolume
}

func NewWSP_FileServerToVolumeEx1(instance *cim.WmiInstance) (newInstance *WSP_FileServerToVolume, err error) {
	tmp, err := NewMSFT_FileServerToVolumeEx1(instance)

	if err != nil {
		return
	}
	newInstance = &WSP_FileServerToVolume{
		MSFT_FileServerToVolume: tmp,
	}
	return
}

func NewWSP_FileServerToVolumeEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *WSP_FileServerToVolume, err error) {
	tmp, err := NewMSFT_FileServerToVolumeEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &WSP_FileServerToVolume{
		MSFT_FileServerToVolume: tmp,
	}
	return
}
