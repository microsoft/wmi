// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.StandardCimv2.mlnx
//////////////////////////////////////////////
package mlnx

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// CIM_SoftwareInstallationService struct
type CIM_SoftwareInstallationService struct {
	*CIM_Service
}

func NewCIM_SoftwareInstallationServiceEx1(instance *cim.WmiInstance) (newInstance *CIM_SoftwareInstallationService, err error) {
	tmp, err := NewCIM_ServiceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_SoftwareInstallationService{
		CIM_Service: tmp,
	}
	return
}

func NewCIM_SoftwareInstallationServiceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_SoftwareInstallationService, err error) {
	tmp, err := NewCIM_ServiceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_SoftwareInstallationService{
		CIM_Service: tmp,
	}
	return
}

//

// <param name="Collection" type="CIM_Collection "></param>
// <param name="Source" type="CIM_SoftwareIdentity "></param>
// <param name="Target" type="CIM_ManagedElement "></param>

// <param name="InstallCharacteristics" type="SoftwareInstallationService_InstallCharacteristics []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *CIM_SoftwareInstallationService) CheckSoftwareIdentity( /* IN */ Source CIM_SoftwareIdentity,
	/* IN */ Target CIM_ManagedElement,
	/* IN */ Collection CIM_Collection,
	/* OUT */ InstallCharacteristics []SoftwareInstallationService_InstallCharacteristics) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("CheckSoftwareIdentity", Source, Target, Collection)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="Collection" type="CIM_Collection "></param>
// <param name="InstallOptions" type="SoftwareInstallationService_InstallOptions []"></param>
// <param name="InstallOptionsValues" type="string []"></param>
// <param name="Source" type="CIM_SoftwareIdentity "></param>
// <param name="Target" type="CIM_ManagedElement "></param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *CIM_SoftwareInstallationService) InstallFromSoftwareIdentity( /* OUT */ Job CIM_ConcreteJob,
	/* OPTIONAL IN */ InstallOptions []SoftwareInstallationService_InstallOptions,
	/* OPTIONAL IN */ InstallOptionsValues []string,
	/* OPTIONAL IN */ Source CIM_SoftwareIdentity,
	/* OPTIONAL IN */ Target CIM_ManagedElement,
	/* OPTIONAL IN */ Collection CIM_Collection,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("InstallFromSoftwareIdentity", Action, PercentComplete, Timeout, InstallOptions, InstallOptionsValues, Source, Target, Collection)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="InstallOptions" type="SoftwareInstallationService_InstallOptions []"></param>
// <param name="InstallOptionsValues" type="string []"></param>
// <param name="Target" type="CIM_ManagedElement "></param>
// <param name="URI" type="string "></param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *CIM_SoftwareInstallationService) InstallFromURI( /* OUT */ Job CIM_ConcreteJob,
	/* OPTIONAL IN */ URI string,
	/* OPTIONAL IN */ Target CIM_ManagedElement,
	/* OPTIONAL IN */ InstallOptions []SoftwareInstallationService_InstallOptions,
	/* OPTIONAL IN */ InstallOptionsValues []string,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("InstallFromURI", Action, PercentComplete, Timeout, URI, Target, InstallOptions, InstallOptionsValues)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
