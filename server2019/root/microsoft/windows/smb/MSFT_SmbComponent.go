// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.Microsoft.Windows.SMB
//////////////////////////////////////////////
package smb
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/base/instance"
)

// MSFT_SmbComponent struct
type MSFT_SmbComponent struct { 
	*cim.WmiInstance
}

	func NewMSFT_SmbComponentEx1(instance *cim.WmiInstance) (newInstance *MSFT_SmbComponent, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &MSFT_SmbComponent {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewMSFT_SmbComponentEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MSFT_SmbComponent, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MSFT_SmbComponent {
	WmiInstance: tmp,
	}
	return
	}
	

// 

// <param name="Name" type="string []"></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_SmbComponent) RemoveSmbComponent( /* IN */ Name []string) (result uint32, err error) {retVal, err := instance.InvokeMethodWithReturn("RemoveSmbComponent" , Name);
	if err != nil { return }
	result = uint32(retVal)
	return
	
}


// 

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_SmbComponent) ClientUnloadSmbDirect() (result uint32, err error) {retVal, err := instance.InvokeMethodWithReturn("ClientUnloadSmbDirect" );
	if err != nil { return }
	result = uint32(retVal)
	return
	
}


// 

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_SmbComponent) ClientReloadSmbDirect() (result uint32, err error) {retVal, err := instance.InvokeMethodWithReturn("ClientReloadSmbDirect" );
	if err != nil { return }
	result = uint32(retVal)
	return
	
}


// 

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_SmbComponent) ServerUnloadSmbDirect() (result uint32, err error) {retVal, err := instance.InvokeMethodWithReturn("ServerUnloadSmbDirect" );
	if err != nil { return }
	result = uint32(retVal)
	return
	
}


// 

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_SmbComponent) ServerReloadSmbDirect() (result uint32, err error) {retVal, err := instance.InvokeMethodWithReturn("ServerReloadSmbDirect" );
	if err != nil { return }
	result = uint32(retVal)
	return
	
}


