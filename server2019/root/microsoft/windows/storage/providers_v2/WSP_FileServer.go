// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.Microsoft.Windows.Storage.Providers_v2
//////////////////////////////////////////////
package providers_v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// WSP_FileServer struct
type WSP_FileServer struct {
	*MSFT_FileServer
}

func NewWSP_FileServerEx1(instance *cim.WmiInstance) (newInstance *WSP_FileServer, err error) {
	tmp, err := NewMSFT_FileServerEx1(instance)

	if err != nil {
		return
	}
	newInstance = &WSP_FileServer{
		MSFT_FileServer: tmp,
	}
	return
}

func NewWSP_FileServerEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *WSP_FileServer, err error) {
	tmp, err := NewMSFT_FileServerEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &WSP_FileServer{
		MSFT_FileServer: tmp,
	}
	return
}

//

// <param name="FileSharingProtocols" type="uint16 []"></param>
// <param name="FriendlyName" type="string "></param>
// <param name="HostNames" type="string []"></param>
// <param name="StorageSubSystemId" type="string "></param>

// <param name="CreatedFileServer" type="MSFT_FileServer "></param>
// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *WSP_FileServer) CreateFileServer( /* IN */ StorageSubSystemId string,
	/* IN */ FriendlyName string,
	/* IN */ FileSharingProtocols []uint16,
	/* IN */ HostNames []string,
	/* OUT */ CreatedFileServer MSFT_FileServer,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("CreateFileServer", StorageSubSystemId, FriendlyName, FileSharingProtocols, HostNames)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
