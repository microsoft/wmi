// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/base/instance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// MS_SM_FabricAndDomainManagementMethods struct
type MS_SM_FabricAndDomainManagementMethods struct { 
	*cim.WmiInstance

	// 
	Active bool

	// 
	InstanceName string
}

	func NewMS_SM_FabricAndDomainManagementMethodsEx1(instance *cim.WmiInstance) (newInstance *MS_SM_FabricAndDomainManagementMethods, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &MS_SM_FabricAndDomainManagementMethods {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewMS_SM_FabricAndDomainManagementMethodsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MS_SM_FabricAndDomainManagementMethods, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MS_SM_FabricAndDomainManagementMethods {
	WmiInstance: tmp,
	}
	return
	}
	

// SetActive sets the value of Active for the instance
func (instance *MS_SM_FabricAndDomainManagementMethods) SetPropertyActive(value bool) (err error) { 
    return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MS_SM_FabricAndDomainManagementMethods) GetPropertyActive()(value bool, err error) { 
    retValue, err := instance.GetProperty("Active")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(bool); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = bool(valuetmp)

    return
}

// SetInstanceName sets the value of InstanceName for the instance
func (instance *MS_SM_FabricAndDomainManagementMethods) SetPropertyInstanceName(value string) (err error) { 
    return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MS_SM_FabricAndDomainManagementMethods) GetPropertyInstanceName()(value string, err error) { 
    retValue, err := instance.GetProperty("InstanceName")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(string); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = string(valuetmp)

    return
}

// 

// <param name="DestFCID" type="uint32 "></param>
// <param name="DestWWN" type="uint8 []"></param>
// <param name="HbaPortWWN" type="uint8 []"></param>
// <param name="ReqBuffer" type="uint8 []"></param>
// <param name="ReqBufferSize" type="uint32 "></param>

// <param name="HBAStatus" type="uint32 "></param>
func (instance *MS_SM_FabricAndDomainManagementMethods) SM_SendTEST( /* IN */ HbaPortWWN []uint8,
 /* IN */ DestWWN []uint8,
 /* IN */ DestFCID uint32,
 /* IN */ ReqBufferSize uint32,
 /* IN */ ReqBuffer []uint8,
 /* OUT */ HBAStatus uint32) (err error) {_, err = instance.InvokeMethod("SM_SendTEST" , HbaPortWWN, DestWWN, DestFCID, ReqBufferSize, ReqBuffer)
	if err != nil { return }
	return
	
}


// 

// <param name="DestFCID" type="uint32 "></param>
// <param name="DestWWN" type="uint8 []"></param>
// <param name="HbaPortWWN" type="uint8 []"></param>
// <param name="InRespBufferMaxSize" type="uint32 "></param>
// <param name="ReqBuffer" type="uint8 []"></param>
// <param name="ReqBufferSize" type="uint32 "></param>

// <param name="HBAStatus" type="uint32 "></param>
// <param name="OutRespBufferSize" type="uint32 "></param>
// <param name="RespBuffer" type="uint8 []"></param>
func (instance *MS_SM_FabricAndDomainManagementMethods) SM_SendECHO( /* IN */ HbaPortWWN []uint8,
 /* IN */ DestWWN []uint8,
 /* IN */ DestFCID uint32,
 /* IN */ InRespBufferMaxSize uint32,
 /* IN */ ReqBufferSize uint32,
 /* IN */ ReqBuffer []uint8,
 /* OUT */ HBAStatus uint32,
 /* OUT */ OutRespBufferSize uint32,
 /* OUT */ RespBuffer []uint8) (err error) {_, err = instance.InvokeMethod("SM_SendECHO" , HbaPortWWN, DestWWN, DestFCID, InRespBufferMaxSize, ReqBufferSize, ReqBuffer)
	if err != nil { return }
	return
	
}


// 

// <param name="DestPortWWN" type="uint8 []"></param>
// <param name="DomainPortWWN" type="uint8 []"></param>
// <param name="HbaPortWWN" type="uint8 []"></param>
// <param name="InRespBufferMaxSize" type="uint32 "></param>
// <param name="ReqBuffer" type="uint8 []"></param>
// <param name="ReqBufferSize" type="uint32 "></param>

// <param name="HBAStatus" type="uint32 "></param>
// <param name="OutRespBufferSize" type="uint32 "></param>
// <param name="RespBuffer" type="uint8 []"></param>
// <param name="TotalRespBufferSize" type="uint32 "></param>
func (instance *MS_SM_FabricAndDomainManagementMethods) SM_SendSMPPassThru( /* IN */ HbaPortWWN []uint8,
 /* IN */ DestPortWWN []uint8,
 /* IN */ DomainPortWWN []uint8,
 /* IN */ InRespBufferMaxSize uint32,
 /* IN */ ReqBufferSize uint32,
 /* IN */ ReqBuffer []uint8,
 /* OUT */ HBAStatus uint32,
 /* OUT */ TotalRespBufferSize uint32,
 /* OUT */ OutRespBufferSize uint32,
 /* OUT */ RespBuffer []uint8) (err error) {_, err = instance.InvokeMethod("SM_SendSMPPassThru" , HbaPortWWN, DestPortWWN, DomainPortWWN, InRespBufferMaxSize, ReqBufferSize, ReqBuffer)
	if err != nil { return }
	return
	
}


// 

// <param name="HbaPortWWN" type="uint8 []"></param>
// <param name="InRespBufferMaxSize" type="uint32 "></param>
// <param name="ReqBuffer" type="uint8 []"></param>
// <param name="ReqBufferSize" type="uint32 "></param>

// <param name="HBAStatus" type="uint32 "></param>
// <param name="OutRespBufferSize" type="uint32 "></param>
// <param name="RespBuffer" type="uint8 []"></param>
// <param name="TotalRespBufferSize" type="uint32 "></param>
func (instance *MS_SM_FabricAndDomainManagementMethods) SM_SendCTPassThru( /* IN */ HbaPortWWN []uint8,
 /* IN */ InRespBufferMaxSize uint32,
 /* IN */ ReqBufferSize uint32,
 /* IN */ ReqBuffer []uint8,
 /* OUT */ HBAStatus uint32,
 /* OUT */ TotalRespBufferSize uint32,
 /* OUT */ OutRespBufferSize uint32,
 /* OUT */ RespBuffer []uint8) (err error) {_, err = instance.InvokeMethod("SM_SendCTPassThru" , HbaPortWWN, InRespBufferMaxSize, ReqBufferSize, ReqBuffer)
	if err != nil { return }
	return
	
}


// 

// <param name="HBAStatus" type="uint32 "></param>
// <param name="MgmtInfo" type="HBAFC3MgmtInfo "></param>
func (instance *MS_SM_FabricAndDomainManagementMethods) SM_GetRNIDMgmtInfo( /* OUT */ HBAStatus uint32,
 /* OUT */ MgmtInfo HBAFC3MgmtInfo) (err error) {_, err = instance.InvokeMethod("SM_GetRNIDMgmtInfo" )
	if err != nil { return }
	return
	
}


// 

// <param name="MgmtInfo" type="HBAFC3MgmtInfo "></param>

// <param name="HBAStatus" type="uint32 "></param>
func (instance *MS_SM_FabricAndDomainManagementMethods) SM_SetRNIDMgmtInfo( /* IN */ MgmtInfo HBAFC3MgmtInfo,
 /* OUT */ HBAStatus uint32) (err error) {_, err = instance.InvokeMethod("SM_SetRNIDMgmtInfo" , MgmtInfo)
	if err != nil { return }
	return
	
}


// 

// <param name="DestFCID" type="uint32 "></param>
// <param name="DestWWN" type="uint8 []"></param>
// <param name="HbaPortWWN" type="uint8 []"></param>
// <param name="InRespBufferMaxSize" type="uint32 "></param>
// <param name="NodeIdDataFormat" type="uint32 "></param>

// <param name="HBAStatus" type="uint32 "></param>
// <param name="OutRespBufferSize" type="uint32 "></param>
// <param name="RespBuffer" type="uint8 []"></param>
// <param name="TotalRespBufferSize" type="uint32 "></param>
func (instance *MS_SM_FabricAndDomainManagementMethods) SM_SendRNID( /* IN */ HbaPortWWN []uint8,
 /* IN */ DestWWN []uint8,
 /* IN */ DestFCID uint32,
 /* IN */ NodeIdDataFormat uint32,
 /* IN */ InRespBufferMaxSize uint32,
 /* OUT */ HBAStatus uint32,
 /* OUT */ TotalRespBufferSize uint32,
 /* OUT */ OutRespBufferSize uint32,
 /* OUT */ RespBuffer []uint8) (err error) {_, err = instance.InvokeMethod("SM_SendRNID" , HbaPortWWN, DestWWN, DestFCID, NodeIdDataFormat, InRespBufferMaxSize)
	if err != nil { return }
	return
	
}


// 

// <param name="AgentDomain" type="uint32 "></param>
// <param name="AgentWWN" type="uint8 []"></param>
// <param name="HbaPortWWN" type="uint8 []"></param>
// <param name="InRespBufferMaxSize" type="uint32 "></param>
// <param name="PortIndex" type="uint32 "></param>

// <param name="HBAStatus" type="uint32 "></param>
// <param name="OutRespBufferSize" type="uint32 "></param>
// <param name="RespBuffer" type="uint8 []"></param>
// <param name="TotalRespBufferSize" type="uint32 "></param>
func (instance *MS_SM_FabricAndDomainManagementMethods) SM_SendRPL( /* IN */ HbaPortWWN []uint8,
 /* IN */ AgentWWN []uint8,
 /* IN */ AgentDomain uint32,
 /* IN */ PortIndex uint32,
 /* IN */ InRespBufferMaxSize uint32,
 /* OUT */ HBAStatus uint32,
 /* OUT */ TotalRespBufferSize uint32,
 /* OUT */ OutRespBufferSize uint32,
 /* OUT */ RespBuffer []uint8) (err error) {_, err = instance.InvokeMethod("SM_SendRPL" , HbaPortWWN, AgentWWN, AgentDomain, PortIndex, InRespBufferMaxSize)
	if err != nil { return }
	return
	
}


// 

// <param name="AgentDomain" type="uint32 "></param>
// <param name="AgentWWN" type="uint8 []"></param>
// <param name="HbaPortWWN" type="uint8 []"></param>
// <param name="InRespBufferMaxSize" type="uint32 "></param>
// <param name="ObjectPortNumber" type="uint32 "></param>
// <param name="ObjectWWN" type="uint8 []"></param>

// <param name="HBAStatus" type="uint32 "></param>
// <param name="OutRespBufferSize" type="uint32 "></param>
// <param name="RespBuffer" type="uint8 []"></param>
// <param name="TotalRespBufferSize" type="uint32 "></param>
func (instance *MS_SM_FabricAndDomainManagementMethods) SM_SendRPS( /* IN */ HbaPortWWN []uint8,
 /* IN */ AgentWWN []uint8,
 /* IN */ ObjectWWN []uint8,
 /* IN */ AgentDomain uint32,
 /* IN */ ObjectPortNumber uint32,
 /* IN */ InRespBufferMaxSize uint32,
 /* OUT */ HBAStatus uint32,
 /* OUT */ TotalRespBufferSize uint32,
 /* OUT */ OutRespBufferSize uint32,
 /* OUT */ RespBuffer []uint8) (err error) {_, err = instance.InvokeMethod("SM_SendRPS" , HbaPortWWN, AgentWWN, ObjectWWN, AgentDomain, ObjectPortNumber, InRespBufferMaxSize)
	if err != nil { return }
	return
	
}


// 

// <param name="Domain" type="uint32 "></param>
// <param name="HbaPortWWN" type="uint8 []"></param>
// <param name="InRespBufferMaxSize" type="uint32 "></param>
// <param name="WWN" type="uint8 []"></param>

// <param name="HBAStatus" type="uint32 "></param>
// <param name="OutRespBufferSize" type="uint32 "></param>
// <param name="RespBuffer" type="uint8 []"></param>
// <param name="TotalRespBufferSize" type="uint32 "></param>
func (instance *MS_SM_FabricAndDomainManagementMethods) SM_SendSRL( /* IN */ HbaPortWWN []uint8,
 /* IN */ WWN []uint8,
 /* IN */ Domain uint32,
 /* IN */ InRespBufferMaxSize uint32,
 /* OUT */ HBAStatus uint32,
 /* OUT */ TotalRespBufferSize uint32,
 /* OUT */ OutRespBufferSize uint32,
 /* OUT */ RespBuffer []uint8) (err error) {_, err = instance.InvokeMethod("SM_SendSRL" , HbaPortWWN, WWN, Domain, InRespBufferMaxSize)
	if err != nil { return }
	return
	
}


// 

// <param name="DestWWN" type="uint8 []"></param>
// <param name="Function" type="uint8 "></param>
// <param name="InRespBufferMaxSize" type="uint32 "></param>
// <param name="SourceWWN" type="uint8 []"></param>
// <param name="Type" type="uint8 "></param>

// <param name="HBAStatus" type="uint32 "></param>
// <param name="OutRespBufferSize" type="uint32 "></param>
// <param name="RespBuffer" type="uint8 []"></param>
// <param name="TotalRespBufferSize" type="uint32 "></param>
func (instance *MS_SM_FabricAndDomainManagementMethods) SM_SendLIRR( /* IN */ SourceWWN []uint8,
 /* IN */ DestWWN []uint8,
 /* IN */ Function uint8,
 /* IN */ Type uint8,
 /* IN */ InRespBufferMaxSize uint32,
 /* OUT */ HBAStatus uint32,
 /* OUT */ TotalRespBufferSize uint32,
 /* OUT */ OutRespBufferSize uint32,
 /* OUT */ RespBuffer []uint8) (err error) {_, err = instance.InvokeMethod("SM_SendLIRR" , SourceWWN, DestWWN, Function, Type, InRespBufferMaxSize)
	if err != nil { return }
	return
	
}


// 

// <param name="DestWWN" type="uint8 []"></param>
// <param name="HbaPortWWN" type="uint8 []"></param>
// <param name="InRespBufferMaxSize" type="uint32 "></param>

// <param name="HBAStatus" type="uint32 "></param>
// <param name="OutRespBufferSize" type="uint32 "></param>
// <param name="RespBuffer" type="uint8 []"></param>
// <param name="TotalRespBufferSize" type="uint32 "></param>
func (instance *MS_SM_FabricAndDomainManagementMethods) SM_SendRLS( /* IN */ HbaPortWWN []uint8,
 /* IN */ DestWWN []uint8,
 /* IN */ InRespBufferMaxSize uint32,
 /* OUT */ HBAStatus uint32,
 /* OUT */ TotalRespBufferSize uint32,
 /* OUT */ OutRespBufferSize uint32,
 /* OUT */ RespBuffer []uint8) (err error) {_, err = instance.InvokeMethod("SM_SendRLS" , HbaPortWWN, DestWWN, InRespBufferMaxSize)
	if err != nil { return }
	return
	
}


