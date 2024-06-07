// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.Microsoft.Windows.RemoteAccess
//////////////////////////////////////////////
package remoteaccess
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/base/instance"
)

// PS_DAEntryPointDC struct
type PS_DAEntryPointDC struct { 
	*cim.WmiInstance
}

	func NewPS_DAEntryPointDCEx1(instance *cim.WmiInstance) (newInstance *PS_DAEntryPointDC, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &PS_DAEntryPointDC {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewPS_DAEntryPointDCEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *PS_DAEntryPointDC, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &PS_DAEntryPointDC {
	WmiInstance: tmp,
	}
	return
	}
	

// 

// <param name="ComputerName" type="string "></param>
// <param name="EntryPointName" type="string "></param>

// <param name="cmdletOutput" type="DADomainController []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_DAEntryPointDC) Get( /* IN */ ComputerName string,
 /* IN */ EntryPointName string,
 /* OUT */ cmdletOutput []DADomainController) (result uint32, err error) {retVal, err := instance.InvokeMethod("Get" , ComputerName, EntryPointName)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// 

// <param name="ComputerName" type="string "></param>
// <param name="ExistingDC" type="string "></param>
// <param name="Force" type="bool "></param>
// <param name="NewDC" type="string "></param>
// <param name="PassThru" type="bool "></param>

// <param name="cmdletOutput" type="DADomainController []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_DAEntryPointDC) SetByDCName( /* IN */ ComputerName string,
 /* IN */ NewDC string,
 /* IN */ Force bool,
 /* IN */ PassThru bool,
 /* IN */ ExistingDC string,
 /* OUT */ cmdletOutput []DADomainController) (result uint32, err error) {retVal, err := instance.InvokeMethod("SetByDCName" , ComputerName, NewDC, Force, PassThru, ExistingDC)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// 

// <param name="ComputerName" type="string "></param>
// <param name="EntryPointName" type="string "></param>
// <param name="Force" type="bool "></param>
// <param name="NewDC" type="string "></param>
// <param name="PassThru" type="bool "></param>

// <param name="cmdletOutput" type="DADomainController []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_DAEntryPointDC) SetByEntryPoint( /* IN */ ComputerName string,
 /* IN */ NewDC string,
 /* IN */ Force bool,
 /* IN */ PassThru bool,
 /* IN */ EntryPointName string,
 /* OUT */ cmdletOutput []DADomainController) (result uint32, err error) {retVal, err := instance.InvokeMethod("SetByEntryPoint" , ComputerName, NewDC, Force, PassThru, EntryPointName)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


