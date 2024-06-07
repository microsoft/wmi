// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// Win32_PerfRawData_LdapPerformanceCounters_LdapClient struct
type Win32_PerfRawData_LdapPerformanceCounters_LdapClient struct { 
	*Win32_PerfRawData

	// 
	BindsDigestBindsPersec uint32

	// 
	BindsNegotiateBindsPersec uint32

	// 
	BindsNTLMBindsPersec uint32

	// 
	BindsSimpleBindsPersec uint32

	// 
	BindsTotalBindsPersec uint32

	// 
	ConnectionsNewConnectionsPersec uint32

	// 
	ConnectionsNewTCPConnectionsPersec uint32

	// 
	ConnectionsNewTLSConnectionsPersec uint32

	// 
	ConnectionsNewUDPConnectionsPersec uint32

	// 
	ConnectionsOpenConnections uint32

	// 
	OperationsAbandonsPersec uint32

	// 
	OperationsAddsPersec uint32

	// 
	OperationsDeletesPersec uint32

	// 
	OperationsModifyPersec uint32

	// 
	RequestsNewRequestsPersec uint32

	// 
	RequestsRequestCount uint32

	// 
	ResponsesAverageResponseTime uint32

	// 
	ResponsesAverageResponseTime_Base uint32

	// 
	ResponsesFailurePollingResponsesPersec uint32

	// 
	ResponsesFailureResponsesPersec uint32

	// 
	ResponsesPendingResponses uint32

	// 
	ResponsesSuccessfulPollingResponsesPersec uint32

	// 
	ResponsesSuccessfulResponsesPersec uint32

	// 
	SearchesBaseSearchesPersec uint32

	// 
	SearchesOnelevelSearchesPersec uint32

	// 
	SearchesSubtreeSearchesPersec uint32
}

	func NewWin32_PerfRawData_LdapPerformanceCounters_LdapClientEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfRawData_LdapPerformanceCounters_LdapClient, err error) {tmp, err := NewWin32_PerfRawDataEx1(instance)
		
	if err != nil { return }
	newInstance = &Win32_PerfRawData_LdapPerformanceCounters_LdapClient {
	Win32_PerfRawData: tmp,
	}
	return
	}
	

	func NewWin32_PerfRawData_LdapPerformanceCounters_LdapClientEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *Win32_PerfRawData_LdapPerformanceCounters_LdapClient, err error) {tmp, err := NewWin32_PerfRawDataEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &Win32_PerfRawData_LdapPerformanceCounters_LdapClient {
	Win32_PerfRawData: tmp,
	}
	return
	}
	

// SetBindsDigestBindsPersec sets the value of BindsDigestBindsPersec for the instance
func (instance *Win32_PerfRawData_LdapPerformanceCounters_LdapClient) SetPropertyBindsDigestBindsPersec(value uint32) (err error) { 
    return instance.SetProperty("BindsDigestBindsPersec", (value))
}

// GetBindsDigestBindsPersec gets the value of BindsDigestBindsPersec for the instance
func (instance *Win32_PerfRawData_LdapPerformanceCounters_LdapClient) GetPropertyBindsDigestBindsPersec()(value uint32, err error) { 
    retValue, err := instance.GetProperty("BindsDigestBindsPersec")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetBindsNegotiateBindsPersec sets the value of BindsNegotiateBindsPersec for the instance
func (instance *Win32_PerfRawData_LdapPerformanceCounters_LdapClient) SetPropertyBindsNegotiateBindsPersec(value uint32) (err error) { 
    return instance.SetProperty("BindsNegotiateBindsPersec", (value))
}

// GetBindsNegotiateBindsPersec gets the value of BindsNegotiateBindsPersec for the instance
func (instance *Win32_PerfRawData_LdapPerformanceCounters_LdapClient) GetPropertyBindsNegotiateBindsPersec()(value uint32, err error) { 
    retValue, err := instance.GetProperty("BindsNegotiateBindsPersec")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetBindsNTLMBindsPersec sets the value of BindsNTLMBindsPersec for the instance
func (instance *Win32_PerfRawData_LdapPerformanceCounters_LdapClient) SetPropertyBindsNTLMBindsPersec(value uint32) (err error) { 
    return instance.SetProperty("BindsNTLMBindsPersec", (value))
}

// GetBindsNTLMBindsPersec gets the value of BindsNTLMBindsPersec for the instance
func (instance *Win32_PerfRawData_LdapPerformanceCounters_LdapClient) GetPropertyBindsNTLMBindsPersec()(value uint32, err error) { 
    retValue, err := instance.GetProperty("BindsNTLMBindsPersec")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetBindsSimpleBindsPersec sets the value of BindsSimpleBindsPersec for the instance
func (instance *Win32_PerfRawData_LdapPerformanceCounters_LdapClient) SetPropertyBindsSimpleBindsPersec(value uint32) (err error) { 
    return instance.SetProperty("BindsSimpleBindsPersec", (value))
}

// GetBindsSimpleBindsPersec gets the value of BindsSimpleBindsPersec for the instance
func (instance *Win32_PerfRawData_LdapPerformanceCounters_LdapClient) GetPropertyBindsSimpleBindsPersec()(value uint32, err error) { 
    retValue, err := instance.GetProperty("BindsSimpleBindsPersec")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetBindsTotalBindsPersec sets the value of BindsTotalBindsPersec for the instance
func (instance *Win32_PerfRawData_LdapPerformanceCounters_LdapClient) SetPropertyBindsTotalBindsPersec(value uint32) (err error) { 
    return instance.SetProperty("BindsTotalBindsPersec", (value))
}

// GetBindsTotalBindsPersec gets the value of BindsTotalBindsPersec for the instance
func (instance *Win32_PerfRawData_LdapPerformanceCounters_LdapClient) GetPropertyBindsTotalBindsPersec()(value uint32, err error) { 
    retValue, err := instance.GetProperty("BindsTotalBindsPersec")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetConnectionsNewConnectionsPersec sets the value of ConnectionsNewConnectionsPersec for the instance
func (instance *Win32_PerfRawData_LdapPerformanceCounters_LdapClient) SetPropertyConnectionsNewConnectionsPersec(value uint32) (err error) { 
    return instance.SetProperty("ConnectionsNewConnectionsPersec", (value))
}

// GetConnectionsNewConnectionsPersec gets the value of ConnectionsNewConnectionsPersec for the instance
func (instance *Win32_PerfRawData_LdapPerformanceCounters_LdapClient) GetPropertyConnectionsNewConnectionsPersec()(value uint32, err error) { 
    retValue, err := instance.GetProperty("ConnectionsNewConnectionsPersec")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetConnectionsNewTCPConnectionsPersec sets the value of ConnectionsNewTCPConnectionsPersec for the instance
func (instance *Win32_PerfRawData_LdapPerformanceCounters_LdapClient) SetPropertyConnectionsNewTCPConnectionsPersec(value uint32) (err error) { 
    return instance.SetProperty("ConnectionsNewTCPConnectionsPersec", (value))
}

// GetConnectionsNewTCPConnectionsPersec gets the value of ConnectionsNewTCPConnectionsPersec for the instance
func (instance *Win32_PerfRawData_LdapPerformanceCounters_LdapClient) GetPropertyConnectionsNewTCPConnectionsPersec()(value uint32, err error) { 
    retValue, err := instance.GetProperty("ConnectionsNewTCPConnectionsPersec")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetConnectionsNewTLSConnectionsPersec sets the value of ConnectionsNewTLSConnectionsPersec for the instance
func (instance *Win32_PerfRawData_LdapPerformanceCounters_LdapClient) SetPropertyConnectionsNewTLSConnectionsPersec(value uint32) (err error) { 
    return instance.SetProperty("ConnectionsNewTLSConnectionsPersec", (value))
}

// GetConnectionsNewTLSConnectionsPersec gets the value of ConnectionsNewTLSConnectionsPersec for the instance
func (instance *Win32_PerfRawData_LdapPerformanceCounters_LdapClient) GetPropertyConnectionsNewTLSConnectionsPersec()(value uint32, err error) { 
    retValue, err := instance.GetProperty("ConnectionsNewTLSConnectionsPersec")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetConnectionsNewUDPConnectionsPersec sets the value of ConnectionsNewUDPConnectionsPersec for the instance
func (instance *Win32_PerfRawData_LdapPerformanceCounters_LdapClient) SetPropertyConnectionsNewUDPConnectionsPersec(value uint32) (err error) { 
    return instance.SetProperty("ConnectionsNewUDPConnectionsPersec", (value))
}

// GetConnectionsNewUDPConnectionsPersec gets the value of ConnectionsNewUDPConnectionsPersec for the instance
func (instance *Win32_PerfRawData_LdapPerformanceCounters_LdapClient) GetPropertyConnectionsNewUDPConnectionsPersec()(value uint32, err error) { 
    retValue, err := instance.GetProperty("ConnectionsNewUDPConnectionsPersec")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetConnectionsOpenConnections sets the value of ConnectionsOpenConnections for the instance
func (instance *Win32_PerfRawData_LdapPerformanceCounters_LdapClient) SetPropertyConnectionsOpenConnections(value uint32) (err error) { 
    return instance.SetProperty("ConnectionsOpenConnections", (value))
}

// GetConnectionsOpenConnections gets the value of ConnectionsOpenConnections for the instance
func (instance *Win32_PerfRawData_LdapPerformanceCounters_LdapClient) GetPropertyConnectionsOpenConnections()(value uint32, err error) { 
    retValue, err := instance.GetProperty("ConnectionsOpenConnections")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetOperationsAbandonsPersec sets the value of OperationsAbandonsPersec for the instance
func (instance *Win32_PerfRawData_LdapPerformanceCounters_LdapClient) SetPropertyOperationsAbandonsPersec(value uint32) (err error) { 
    return instance.SetProperty("OperationsAbandonsPersec", (value))
}

// GetOperationsAbandonsPersec gets the value of OperationsAbandonsPersec for the instance
func (instance *Win32_PerfRawData_LdapPerformanceCounters_LdapClient) GetPropertyOperationsAbandonsPersec()(value uint32, err error) { 
    retValue, err := instance.GetProperty("OperationsAbandonsPersec")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetOperationsAddsPersec sets the value of OperationsAddsPersec for the instance
func (instance *Win32_PerfRawData_LdapPerformanceCounters_LdapClient) SetPropertyOperationsAddsPersec(value uint32) (err error) { 
    return instance.SetProperty("OperationsAddsPersec", (value))
}

// GetOperationsAddsPersec gets the value of OperationsAddsPersec for the instance
func (instance *Win32_PerfRawData_LdapPerformanceCounters_LdapClient) GetPropertyOperationsAddsPersec()(value uint32, err error) { 
    retValue, err := instance.GetProperty("OperationsAddsPersec")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetOperationsDeletesPersec sets the value of OperationsDeletesPersec for the instance
func (instance *Win32_PerfRawData_LdapPerformanceCounters_LdapClient) SetPropertyOperationsDeletesPersec(value uint32) (err error) { 
    return instance.SetProperty("OperationsDeletesPersec", (value))
}

// GetOperationsDeletesPersec gets the value of OperationsDeletesPersec for the instance
func (instance *Win32_PerfRawData_LdapPerformanceCounters_LdapClient) GetPropertyOperationsDeletesPersec()(value uint32, err error) { 
    retValue, err := instance.GetProperty("OperationsDeletesPersec")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetOperationsModifyPersec sets the value of OperationsModifyPersec for the instance
func (instance *Win32_PerfRawData_LdapPerformanceCounters_LdapClient) SetPropertyOperationsModifyPersec(value uint32) (err error) { 
    return instance.SetProperty("OperationsModifyPersec", (value))
}

// GetOperationsModifyPersec gets the value of OperationsModifyPersec for the instance
func (instance *Win32_PerfRawData_LdapPerformanceCounters_LdapClient) GetPropertyOperationsModifyPersec()(value uint32, err error) { 
    retValue, err := instance.GetProperty("OperationsModifyPersec")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetRequestsNewRequestsPersec sets the value of RequestsNewRequestsPersec for the instance
func (instance *Win32_PerfRawData_LdapPerformanceCounters_LdapClient) SetPropertyRequestsNewRequestsPersec(value uint32) (err error) { 
    return instance.SetProperty("RequestsNewRequestsPersec", (value))
}

// GetRequestsNewRequestsPersec gets the value of RequestsNewRequestsPersec for the instance
func (instance *Win32_PerfRawData_LdapPerformanceCounters_LdapClient) GetPropertyRequestsNewRequestsPersec()(value uint32, err error) { 
    retValue, err := instance.GetProperty("RequestsNewRequestsPersec")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetRequestsRequestCount sets the value of RequestsRequestCount for the instance
func (instance *Win32_PerfRawData_LdapPerformanceCounters_LdapClient) SetPropertyRequestsRequestCount(value uint32) (err error) { 
    return instance.SetProperty("RequestsRequestCount", (value))
}

// GetRequestsRequestCount gets the value of RequestsRequestCount for the instance
func (instance *Win32_PerfRawData_LdapPerformanceCounters_LdapClient) GetPropertyRequestsRequestCount()(value uint32, err error) { 
    retValue, err := instance.GetProperty("RequestsRequestCount")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetResponsesAverageResponseTime sets the value of ResponsesAverageResponseTime for the instance
func (instance *Win32_PerfRawData_LdapPerformanceCounters_LdapClient) SetPropertyResponsesAverageResponseTime(value uint32) (err error) { 
    return instance.SetProperty("ResponsesAverageResponseTime", (value))
}

// GetResponsesAverageResponseTime gets the value of ResponsesAverageResponseTime for the instance
func (instance *Win32_PerfRawData_LdapPerformanceCounters_LdapClient) GetPropertyResponsesAverageResponseTime()(value uint32, err error) { 
    retValue, err := instance.GetProperty("ResponsesAverageResponseTime")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetResponsesAverageResponseTime_Base sets the value of ResponsesAverageResponseTime_Base for the instance
func (instance *Win32_PerfRawData_LdapPerformanceCounters_LdapClient) SetPropertyResponsesAverageResponseTime_Base(value uint32) (err error) { 
    return instance.SetProperty("ResponsesAverageResponseTime_Base", (value))
}

// GetResponsesAverageResponseTime_Base gets the value of ResponsesAverageResponseTime_Base for the instance
func (instance *Win32_PerfRawData_LdapPerformanceCounters_LdapClient) GetPropertyResponsesAverageResponseTime_Base()(value uint32, err error) { 
    retValue, err := instance.GetProperty("ResponsesAverageResponseTime_Base")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetResponsesFailurePollingResponsesPersec sets the value of ResponsesFailurePollingResponsesPersec for the instance
func (instance *Win32_PerfRawData_LdapPerformanceCounters_LdapClient) SetPropertyResponsesFailurePollingResponsesPersec(value uint32) (err error) { 
    return instance.SetProperty("ResponsesFailurePollingResponsesPersec", (value))
}

// GetResponsesFailurePollingResponsesPersec gets the value of ResponsesFailurePollingResponsesPersec for the instance
func (instance *Win32_PerfRawData_LdapPerformanceCounters_LdapClient) GetPropertyResponsesFailurePollingResponsesPersec()(value uint32, err error) { 
    retValue, err := instance.GetProperty("ResponsesFailurePollingResponsesPersec")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetResponsesFailureResponsesPersec sets the value of ResponsesFailureResponsesPersec for the instance
func (instance *Win32_PerfRawData_LdapPerformanceCounters_LdapClient) SetPropertyResponsesFailureResponsesPersec(value uint32) (err error) { 
    return instance.SetProperty("ResponsesFailureResponsesPersec", (value))
}

// GetResponsesFailureResponsesPersec gets the value of ResponsesFailureResponsesPersec for the instance
func (instance *Win32_PerfRawData_LdapPerformanceCounters_LdapClient) GetPropertyResponsesFailureResponsesPersec()(value uint32, err error) { 
    retValue, err := instance.GetProperty("ResponsesFailureResponsesPersec")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetResponsesPendingResponses sets the value of ResponsesPendingResponses for the instance
func (instance *Win32_PerfRawData_LdapPerformanceCounters_LdapClient) SetPropertyResponsesPendingResponses(value uint32) (err error) { 
    return instance.SetProperty("ResponsesPendingResponses", (value))
}

// GetResponsesPendingResponses gets the value of ResponsesPendingResponses for the instance
func (instance *Win32_PerfRawData_LdapPerformanceCounters_LdapClient) GetPropertyResponsesPendingResponses()(value uint32, err error) { 
    retValue, err := instance.GetProperty("ResponsesPendingResponses")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetResponsesSuccessfulPollingResponsesPersec sets the value of ResponsesSuccessfulPollingResponsesPersec for the instance
func (instance *Win32_PerfRawData_LdapPerformanceCounters_LdapClient) SetPropertyResponsesSuccessfulPollingResponsesPersec(value uint32) (err error) { 
    return instance.SetProperty("ResponsesSuccessfulPollingResponsesPersec", (value))
}

// GetResponsesSuccessfulPollingResponsesPersec gets the value of ResponsesSuccessfulPollingResponsesPersec for the instance
func (instance *Win32_PerfRawData_LdapPerformanceCounters_LdapClient) GetPropertyResponsesSuccessfulPollingResponsesPersec()(value uint32, err error) { 
    retValue, err := instance.GetProperty("ResponsesSuccessfulPollingResponsesPersec")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetResponsesSuccessfulResponsesPersec sets the value of ResponsesSuccessfulResponsesPersec for the instance
func (instance *Win32_PerfRawData_LdapPerformanceCounters_LdapClient) SetPropertyResponsesSuccessfulResponsesPersec(value uint32) (err error) { 
    return instance.SetProperty("ResponsesSuccessfulResponsesPersec", (value))
}

// GetResponsesSuccessfulResponsesPersec gets the value of ResponsesSuccessfulResponsesPersec for the instance
func (instance *Win32_PerfRawData_LdapPerformanceCounters_LdapClient) GetPropertyResponsesSuccessfulResponsesPersec()(value uint32, err error) { 
    retValue, err := instance.GetProperty("ResponsesSuccessfulResponsesPersec")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetSearchesBaseSearchesPersec sets the value of SearchesBaseSearchesPersec for the instance
func (instance *Win32_PerfRawData_LdapPerformanceCounters_LdapClient) SetPropertySearchesBaseSearchesPersec(value uint32) (err error) { 
    return instance.SetProperty("SearchesBaseSearchesPersec", (value))
}

// GetSearchesBaseSearchesPersec gets the value of SearchesBaseSearchesPersec for the instance
func (instance *Win32_PerfRawData_LdapPerformanceCounters_LdapClient) GetPropertySearchesBaseSearchesPersec()(value uint32, err error) { 
    retValue, err := instance.GetProperty("SearchesBaseSearchesPersec")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetSearchesOnelevelSearchesPersec sets the value of SearchesOnelevelSearchesPersec for the instance
func (instance *Win32_PerfRawData_LdapPerformanceCounters_LdapClient) SetPropertySearchesOnelevelSearchesPersec(value uint32) (err error) { 
    return instance.SetProperty("SearchesOnelevelSearchesPersec", (value))
}

// GetSearchesOnelevelSearchesPersec gets the value of SearchesOnelevelSearchesPersec for the instance
func (instance *Win32_PerfRawData_LdapPerformanceCounters_LdapClient) GetPropertySearchesOnelevelSearchesPersec()(value uint32, err error) { 
    retValue, err := instance.GetProperty("SearchesOnelevelSearchesPersec")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetSearchesSubtreeSearchesPersec sets the value of SearchesSubtreeSearchesPersec for the instance
func (instance *Win32_PerfRawData_LdapPerformanceCounters_LdapClient) SetPropertySearchesSubtreeSearchesPersec(value uint32) (err error) { 
    return instance.SetProperty("SearchesSubtreeSearchesPersec", (value))
}

// GetSearchesSubtreeSearchesPersec gets the value of SearchesSubtreeSearchesPersec for the instance
func (instance *Win32_PerfRawData_LdapPerformanceCounters_LdapClient) GetPropertySearchesSubtreeSearchesPersec()(value uint32, err error) { 
    retValue, err := instance.GetProperty("SearchesSubtreeSearchesPersec")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

