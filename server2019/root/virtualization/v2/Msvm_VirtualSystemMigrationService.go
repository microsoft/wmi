// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.virtualization.v2
//
// ////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Msvm_VirtualSystemMigrationService struct
type Msvm_VirtualSystemMigrationService struct {
	*CIM_VirtualSystemMigrationService

	//
	ActiveStorageMigrationCount uint32

	//
	ActiveVirtualSystemMigrationCount uint32

	//
	MigrationServiceListenerIPAddressList []string
}

func NewMsvm_VirtualSystemMigrationServiceEx1(instance *cim.WmiInstance) (newInstance *Msvm_VirtualSystemMigrationService, err error) {
	tmp, err := NewCIM_VirtualSystemMigrationServiceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_VirtualSystemMigrationService{
		CIM_VirtualSystemMigrationService: tmp,
	}
	return
}

func NewMsvm_VirtualSystemMigrationServiceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_VirtualSystemMigrationService, err error) {
	tmp, err := NewCIM_VirtualSystemMigrationServiceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_VirtualSystemMigrationService{
		CIM_VirtualSystemMigrationService: tmp,
	}
	return
}

// SetActiveStorageMigrationCount sets the value of ActiveStorageMigrationCount for the instance
func (instance *Msvm_VirtualSystemMigrationService) SetPropertyActiveStorageMigrationCount(value uint32) (err error) {
	return instance.SetProperty("ActiveStorageMigrationCount", (value))
}

// GetActiveStorageMigrationCount gets the value of ActiveStorageMigrationCount for the instance
func (instance *Msvm_VirtualSystemMigrationService) GetPropertyActiveStorageMigrationCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("ActiveStorageMigrationCount")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetActiveVirtualSystemMigrationCount sets the value of ActiveVirtualSystemMigrationCount for the instance
func (instance *Msvm_VirtualSystemMigrationService) SetPropertyActiveVirtualSystemMigrationCount(value uint32) (err error) {
	return instance.SetProperty("ActiveVirtualSystemMigrationCount", (value))
}

// GetActiveVirtualSystemMigrationCount gets the value of ActiveVirtualSystemMigrationCount for the instance
func (instance *Msvm_VirtualSystemMigrationService) GetPropertyActiveVirtualSystemMigrationCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("ActiveVirtualSystemMigrationCount")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetMigrationServiceListenerIPAddressList sets the value of MigrationServiceListenerIPAddressList for the instance
func (instance *Msvm_VirtualSystemMigrationService) SetPropertyMigrationServiceListenerIPAddressList(value []string) (err error) {
	return instance.SetProperty("MigrationServiceListenerIPAddressList", (value))
}

// GetMigrationServiceListenerIPAddressList gets the value of MigrationServiceListenerIPAddressList for the instance
func (instance *Msvm_VirtualSystemMigrationService) GetPropertyMigrationServiceListenerIPAddressList() (value []string, err error) {
	retValue, err := instance.GetProperty("MigrationServiceListenerIPAddressList")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(string)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, string(valuetmp))
	}

	return
}

//

// <param name="ComputerSystem" type="CIM_ComputerSystem "></param>
// <param name="DestinationHost" type="string "></param>
// <param name="MigrationSettingData" type="string "></param>
// <param name="NewResourceSettingData" type="string []"></param>
// <param name="NewSystemSettingData" type="string "></param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_VirtualSystemMigrationService) CheckVirtualSystemIsMigratable( /* IN */ ComputerSystem CIM_ComputerSystem,
	/* IN */ DestinationHost string,
	/* IN */ MigrationSettingData string,
	/* IN */ NewSystemSettingData string,
	/* IN */ NewResourceSettingData []string,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("CheckVirtualSystemIsMigratable", Action, PercentComplete, Timeout, ComputerSystem, DestinationHost, MigrationSettingData, NewSystemSettingData, NewResourceSettingData)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ServiceSettingData" type="string "></param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_VirtualSystemMigrationService) ModifyServiceSettings( /* IN */ ServiceSettingData string,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("ModifyServiceSettings", Action, PercentComplete, Timeout, ServiceSettingData)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="NetworkSettings" type="string []"></param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_VirtualSystemMigrationService) AddNetworkSettings( /* IN */ NetworkSettings []string,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("AddNetworkSettings", Action, PercentComplete, Timeout, NetworkSettings)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="NetworkSettings" type="Msvm_VirtualSystemMigrationNetworkSettingData []"></param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_VirtualSystemMigrationService) RemoveNetworkSettings( /* IN */ NetworkSettings []Msvm_VirtualSystemMigrationNetworkSettingData,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("RemoveNetworkSettings", Action, PercentComplete, Timeout, NetworkSettings)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="NetworkSettings" type="string []"></param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_VirtualSystemMigrationService) ModifyNetworkSettings( /* IN */ NetworkSettings []string,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("ModifyNetworkSettings", Action, PercentComplete, Timeout, NetworkSettings)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ComputerSystem" type="CIM_ComputerSystem "></param>

// <param name="CompatibilityInfo" type="uint8 []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_VirtualSystemMigrationService) GetSystemCompatibilityInfo( /* IN */ ComputerSystem CIM_ComputerSystem,
	/* OUT */ CompatibilityInfo []uint8) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetSystemCompatibilityInfo", ComputerSystem)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="CompatibilityInfo" type="uint8 []"></param>

// <param name="Reasons" type="string []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_VirtualSystemMigrationService) CheckSystemCompatibilityInfo( /* IN */ CompatibilityInfo []uint8,
	/* OUT */ Reasons []string) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("CheckSystemCompatibilityInfo", CompatibilityInfo)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ComputerSystem" type="CIM_ComputerSystem "></param>

// <param name="CompatibilityVectors" type="Msvm_CompatibilityVector []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_VirtualSystemMigrationService) GetSystemCompatibilityVectors( /* IN */ ComputerSystem CIM_ComputerSystem,
	/* OUT */ CompatibilityVectors []Msvm_CompatibilityVector) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetSystemCompatibilityVectors", ComputerSystem)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="Options" type="string "></param>

// <param name="CompatibilityVectors" type="Msvm_CompatibilityVector []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_VirtualSystemMigrationService) GetProcessorFeatureLimits( /* IN */ Options string,
	/* OUT */ CompatibilityVectors []Msvm_CompatibilityVector) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetProcessorFeatureLimits", Options)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

func (instance *Msvm_VirtualSystemMigrationService) GetRelatedComputerSystem() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_ComputerSystem")
}

func (instance *Msvm_VirtualSystemMigrationService) GetRelatedVirtualSystemMigrationServiceSettingData() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_VirtualSystemMigrationServiceSettingData")
}

func (instance *Msvm_VirtualSystemMigrationService) GetRelatedVirtualSystemMigrationCapabilities() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_VirtualSystemMigrationCapabilities")
}
