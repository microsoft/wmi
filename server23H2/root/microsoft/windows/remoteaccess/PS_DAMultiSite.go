// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.Microsoft.Windows.RemoteAccess
//
// ////////////////////////////////////////////
package remoteaccess

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// PS_DAMultiSite struct
type PS_DAMultiSite struct {
	*cim.WmiInstance
}

func NewPS_DAMultiSiteEx1(instance *cim.WmiInstance) (newInstance *PS_DAMultiSite, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &PS_DAMultiSite{
		WmiInstance: tmp,
	}
	return
}

func NewPS_DAMultiSiteEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *PS_DAMultiSite, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &PS_DAMultiSite{
		WmiInstance: tmp,
	}
	return
}

//

// <param name="ComputerName" type="string "></param>
// <param name="EntryPointName" type="string "></param>
// <param name="Force" type="bool "></param>
// <param name="GslbFqdn" type="string "></param>
// <param name="GslbIP" type="string "></param>
// <param name="ManualEntryPointSelectionAllowed" type="string "></param>
// <param name="Name" type="string "></param>
// <param name="PassThru" type="bool "></param>

// <param name="cmdletOutput" type="DAMultiSite "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_DAMultiSite) Enable( /* IN */ ComputerName string,
	/* IN */ EntryPointName string,
	/* IN */ GslbFqdn string,
	/* IN */ ManualEntryPointSelectionAllowed string,
	/* IN */ Name string,
	/* IN */ GslbIP string,
	/* IN */ Force bool,
	/* IN */ PassThru bool,
	/* OUT */ cmdletOutput DAMultiSite) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Enable", ComputerName, EntryPointName, GslbFqdn, ManualEntryPointSelectionAllowed, Name, GslbIP, Force, PassThru)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ComputerName" type="string "></param>
// <param name="GslbFqdn" type="string "></param>
// <param name="ManualEntryPointSelectionAllowed" type="string "></param>
// <param name="Name" type="string "></param>
// <param name="PassThru" type="bool "></param>

// <param name="cmdletOutput" type="DAMultiSite "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_DAMultiSite) Set( /* IN */ ComputerName string,
	/* IN */ GslbFqdn string,
	/* IN */ ManualEntryPointSelectionAllowed string,
	/* IN */ Name string,
	/* IN */ PassThru bool,
	/* OUT */ cmdletOutput DAMultiSite) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Set", ComputerName, GslbFqdn, ManualEntryPointSelectionAllowed, Name, PassThru)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ComputerName" type="string "></param>
// <param name="Force" type="bool "></param>
// <param name="PassThru" type="bool "></param>

// <param name="cmdletOutput" type="DAMultiSite "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_DAMultiSite) Disable( /* IN */ ComputerName string,
	/* IN */ Force bool,
	/* IN */ PassThru bool,
	/* OUT */ cmdletOutput DAMultiSite) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Disable", ComputerName, Force, PassThru)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ComputerName" type="string "></param>

// <param name="cmdletOutput" type="DAMultiSite "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_DAMultiSite) Get( /* IN */ ComputerName string,
	/* OUT */ cmdletOutput DAMultiSite) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Get", ComputerName)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
