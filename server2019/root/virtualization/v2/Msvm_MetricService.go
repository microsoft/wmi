// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Msvm_MetricService struct
type Msvm_MetricService struct { 
	*CIM_MetricService
}

	func NewMsvm_MetricServiceEx1(instance *cim.WmiInstance) (newInstance *Msvm_MetricService, err error) {tmp, err := NewCIM_MetricServiceEx1(instance)
		
	if err != nil { return }
	newInstance = &Msvm_MetricService {
	CIM_MetricService: tmp,
	}
	return
	}
	

	func NewMsvm_MetricServiceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *Msvm_MetricService, err error) {tmp, err := NewCIM_MetricServiceEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &Msvm_MetricService {
	CIM_MetricService: tmp,
	}
	return
	}
	

// 

// <param name="SettingData" type="string "></param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_MetricService) ModifyServiceSettings( /* IN */ SettingData string,
 /* OUT */ Job CIM_ConcreteJob,
/*Custom IN*/  Action cim.UserAction,
/*Custon IN*/  PercentComplete uint32,
/*Custon IN*/  Timeout uint32) (result uint32, err error) {retVal, err := instance.InvokeMethodAsync("ModifyServiceSettings", Action, PercentComplete, Timeout , SettingData)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}

func  (instance* Msvm_MetricService) GetRelatedComputerSystem() (value *cim.WmiInstance, err error) {
		 return instance.GetRelated("Msvm_ComputerSystem"); 
	}
	
func  (instance* Msvm_MetricService) GetRelatedAggregationMetricDefinition() (value []*cim.WmiInstance, err error) {
		 return instance.GetAllRelated("Msvm_AggregationMetricDefinition"); 
	}
	
func  (instance* Msvm_MetricService) GetRelatedBaseMetricDefinition() (value []*cim.WmiInstance, err error) {
		 return instance.GetAllRelated("Msvm_BaseMetricDefinition"); 
	}
	
func  (instance* Msvm_MetricService) GetRelatedMetricServiceSettingData() (value *cim.WmiInstance, err error) {
		 return instance.GetRelated("Msvm_MetricServiceSettingData"); 
	}
	
func  (instance* Msvm_MetricService) GetRelatedMetricServiceCapabilities() (value *cim.WmiInstance, err error) {
		 return instance.GetRelated("Msvm_MetricServiceCapabilities"); 
	}
	

