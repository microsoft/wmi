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

// Win32_PerfFormattedData_QuicCounters_QUICPerformanceDiagnostics struct
type Win32_PerfFormattedData_QuicCounters_QUICPerformanceDiagnostics struct { 
	*Win32_PerfFormattedData

	// 
	QUICAppReceivedBytesPersec uint32

	// 
	QUICAppSentBytesPersec uint32

	// 
	QUICConnectionOperationsCompletedPersec uint32

	// 
	QUICConnectionOperationsQueued uint64

	// 
	QUICConnectionOperationsQueuedPersec uint32

	// 
	QUICConnectionsActive uint64

	// 
	QUICConnectionsConnected uint64

	// 
	QUICConnectionsCreated uint64

	// 
	QUICConnectionsCreatedPersec uint32

	// 
	QUICConnectionsNoALPN uint64

	// 
	QUICConnectionsNoALPNsPersec uint32

	// 
	QUICConnectionsProtocolError uint64

	// 
	QUICConnectionsProtocolErrorsPersec uint32

	// 
	QUICConnectionsQueued uint64

	// 
	QUICConnectionsRejected uint64

	// 
	QUICConnectionsRejectedPersec uint32

	// 
	QUICConnectionsResumed uint64

	// 
	QUICConnectionsResumedPersec uint32

	// 
	QUICHandshakesFailed uint64

	// 
	QUICHandshakesFailedPersec uint32

	// 
	QUICPacketDecryptionFailuresPersec uint32

	// 
	QUICPacketsDroppedPersec uint32

	// 
	QUICPacketsSuspectedLostPersec uint32

	// 
	QUICPathChallengesFailed uint64

	// 
	QUICPathChallengesSucceeded uint64

	// 
	QUICStatelessResetsSent uint64

	// 
	QUICStatelessRetriesSent uint64

	// 
	QUICStreamsActive uint64

	// 
	QUICUDPDatagramsReceivedPersec uint32

	// 
	QUICUDPDatagramsSentPersec uint32

	// 
	QUICUDPPayloadBytesReceivedPersec uint32

	// 
	QUICUDPPayloadBytesSentPersec uint32

	// 
	QUICUDPReceiveEventsPersec uint32

	// 
	QUICUDPSendCallsPersec uint32

	// 
	QUICWorkerOperationsPersec uint32

	// 
	QUICWorkerOperationsQueued uint64

	// 
	QUICWorkerOperationsQueuedPersec uint32
}

	func NewWin32_PerfFormattedData_QuicCounters_QUICPerformanceDiagnosticsEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfFormattedData_QuicCounters_QUICPerformanceDiagnostics, err error) {tmp, err := NewWin32_PerfFormattedDataEx1(instance)
		
	if err != nil { return }
	newInstance = &Win32_PerfFormattedData_QuicCounters_QUICPerformanceDiagnostics {
	Win32_PerfFormattedData: tmp,
	}
	return
	}
	

	func NewWin32_PerfFormattedData_QuicCounters_QUICPerformanceDiagnosticsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *Win32_PerfFormattedData_QuicCounters_QUICPerformanceDiagnostics, err error) {tmp, err := NewWin32_PerfFormattedDataEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &Win32_PerfFormattedData_QuicCounters_QUICPerformanceDiagnostics {
	Win32_PerfFormattedData: tmp,
	}
	return
	}
	

// SetQUICAppReceivedBytesPersec sets the value of QUICAppReceivedBytesPersec for the instance
func (instance *Win32_PerfFormattedData_QuicCounters_QUICPerformanceDiagnostics) SetPropertyQUICAppReceivedBytesPersec(value uint32) (err error) { 
    return instance.SetProperty("QUICAppReceivedBytesPersec", (value))
}

// GetQUICAppReceivedBytesPersec gets the value of QUICAppReceivedBytesPersec for the instance
func (instance *Win32_PerfFormattedData_QuicCounters_QUICPerformanceDiagnostics) GetPropertyQUICAppReceivedBytesPersec()(value uint32, err error) { 
    retValue, err := instance.GetProperty("QUICAppReceivedBytesPersec")
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

// SetQUICAppSentBytesPersec sets the value of QUICAppSentBytesPersec for the instance
func (instance *Win32_PerfFormattedData_QuicCounters_QUICPerformanceDiagnostics) SetPropertyQUICAppSentBytesPersec(value uint32) (err error) { 
    return instance.SetProperty("QUICAppSentBytesPersec", (value))
}

// GetQUICAppSentBytesPersec gets the value of QUICAppSentBytesPersec for the instance
func (instance *Win32_PerfFormattedData_QuicCounters_QUICPerformanceDiagnostics) GetPropertyQUICAppSentBytesPersec()(value uint32, err error) { 
    retValue, err := instance.GetProperty("QUICAppSentBytesPersec")
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

// SetQUICConnectionOperationsCompletedPersec sets the value of QUICConnectionOperationsCompletedPersec for the instance
func (instance *Win32_PerfFormattedData_QuicCounters_QUICPerformanceDiagnostics) SetPropertyQUICConnectionOperationsCompletedPersec(value uint32) (err error) { 
    return instance.SetProperty("QUICConnectionOperationsCompletedPersec", (value))
}

// GetQUICConnectionOperationsCompletedPersec gets the value of QUICConnectionOperationsCompletedPersec for the instance
func (instance *Win32_PerfFormattedData_QuicCounters_QUICPerformanceDiagnostics) GetPropertyQUICConnectionOperationsCompletedPersec()(value uint32, err error) { 
    retValue, err := instance.GetProperty("QUICConnectionOperationsCompletedPersec")
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

// SetQUICConnectionOperationsQueued sets the value of QUICConnectionOperationsQueued for the instance
func (instance *Win32_PerfFormattedData_QuicCounters_QUICPerformanceDiagnostics) SetPropertyQUICConnectionOperationsQueued(value uint64) (err error) { 
    return instance.SetProperty("QUICConnectionOperationsQueued", (value))
}

// GetQUICConnectionOperationsQueued gets the value of QUICConnectionOperationsQueued for the instance
func (instance *Win32_PerfFormattedData_QuicCounters_QUICPerformanceDiagnostics) GetPropertyQUICConnectionOperationsQueued()(value uint64, err error) { 
    retValue, err := instance.GetProperty("QUICConnectionOperationsQueued")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint64); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint64(valuetmp)

    return
}

// SetQUICConnectionOperationsQueuedPersec sets the value of QUICConnectionOperationsQueuedPersec for the instance
func (instance *Win32_PerfFormattedData_QuicCounters_QUICPerformanceDiagnostics) SetPropertyQUICConnectionOperationsQueuedPersec(value uint32) (err error) { 
    return instance.SetProperty("QUICConnectionOperationsQueuedPersec", (value))
}

// GetQUICConnectionOperationsQueuedPersec gets the value of QUICConnectionOperationsQueuedPersec for the instance
func (instance *Win32_PerfFormattedData_QuicCounters_QUICPerformanceDiagnostics) GetPropertyQUICConnectionOperationsQueuedPersec()(value uint32, err error) { 
    retValue, err := instance.GetProperty("QUICConnectionOperationsQueuedPersec")
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

// SetQUICConnectionsActive sets the value of QUICConnectionsActive for the instance
func (instance *Win32_PerfFormattedData_QuicCounters_QUICPerformanceDiagnostics) SetPropertyQUICConnectionsActive(value uint64) (err error) { 
    return instance.SetProperty("QUICConnectionsActive", (value))
}

// GetQUICConnectionsActive gets the value of QUICConnectionsActive for the instance
func (instance *Win32_PerfFormattedData_QuicCounters_QUICPerformanceDiagnostics) GetPropertyQUICConnectionsActive()(value uint64, err error) { 
    retValue, err := instance.GetProperty("QUICConnectionsActive")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint64); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint64(valuetmp)

    return
}

// SetQUICConnectionsConnected sets the value of QUICConnectionsConnected for the instance
func (instance *Win32_PerfFormattedData_QuicCounters_QUICPerformanceDiagnostics) SetPropertyQUICConnectionsConnected(value uint64) (err error) { 
    return instance.SetProperty("QUICConnectionsConnected", (value))
}

// GetQUICConnectionsConnected gets the value of QUICConnectionsConnected for the instance
func (instance *Win32_PerfFormattedData_QuicCounters_QUICPerformanceDiagnostics) GetPropertyQUICConnectionsConnected()(value uint64, err error) { 
    retValue, err := instance.GetProperty("QUICConnectionsConnected")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint64); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint64(valuetmp)

    return
}

// SetQUICConnectionsCreated sets the value of QUICConnectionsCreated for the instance
func (instance *Win32_PerfFormattedData_QuicCounters_QUICPerformanceDiagnostics) SetPropertyQUICConnectionsCreated(value uint64) (err error) { 
    return instance.SetProperty("QUICConnectionsCreated", (value))
}

// GetQUICConnectionsCreated gets the value of QUICConnectionsCreated for the instance
func (instance *Win32_PerfFormattedData_QuicCounters_QUICPerformanceDiagnostics) GetPropertyQUICConnectionsCreated()(value uint64, err error) { 
    retValue, err := instance.GetProperty("QUICConnectionsCreated")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint64); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint64(valuetmp)

    return
}

// SetQUICConnectionsCreatedPersec sets the value of QUICConnectionsCreatedPersec for the instance
func (instance *Win32_PerfFormattedData_QuicCounters_QUICPerformanceDiagnostics) SetPropertyQUICConnectionsCreatedPersec(value uint32) (err error) { 
    return instance.SetProperty("QUICConnectionsCreatedPersec", (value))
}

// GetQUICConnectionsCreatedPersec gets the value of QUICConnectionsCreatedPersec for the instance
func (instance *Win32_PerfFormattedData_QuicCounters_QUICPerformanceDiagnostics) GetPropertyQUICConnectionsCreatedPersec()(value uint32, err error) { 
    retValue, err := instance.GetProperty("QUICConnectionsCreatedPersec")
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

// SetQUICConnectionsNoALPN sets the value of QUICConnectionsNoALPN for the instance
func (instance *Win32_PerfFormattedData_QuicCounters_QUICPerformanceDiagnostics) SetPropertyQUICConnectionsNoALPN(value uint64) (err error) { 
    return instance.SetProperty("QUICConnectionsNoALPN", (value))
}

// GetQUICConnectionsNoALPN gets the value of QUICConnectionsNoALPN for the instance
func (instance *Win32_PerfFormattedData_QuicCounters_QUICPerformanceDiagnostics) GetPropertyQUICConnectionsNoALPN()(value uint64, err error) { 
    retValue, err := instance.GetProperty("QUICConnectionsNoALPN")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint64); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint64(valuetmp)

    return
}

// SetQUICConnectionsNoALPNsPersec sets the value of QUICConnectionsNoALPNsPersec for the instance
func (instance *Win32_PerfFormattedData_QuicCounters_QUICPerformanceDiagnostics) SetPropertyQUICConnectionsNoALPNsPersec(value uint32) (err error) { 
    return instance.SetProperty("QUICConnectionsNoALPNsPersec", (value))
}

// GetQUICConnectionsNoALPNsPersec gets the value of QUICConnectionsNoALPNsPersec for the instance
func (instance *Win32_PerfFormattedData_QuicCounters_QUICPerformanceDiagnostics) GetPropertyQUICConnectionsNoALPNsPersec()(value uint32, err error) { 
    retValue, err := instance.GetProperty("QUICConnectionsNoALPNsPersec")
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

// SetQUICConnectionsProtocolError sets the value of QUICConnectionsProtocolError for the instance
func (instance *Win32_PerfFormattedData_QuicCounters_QUICPerformanceDiagnostics) SetPropertyQUICConnectionsProtocolError(value uint64) (err error) { 
    return instance.SetProperty("QUICConnectionsProtocolError", (value))
}

// GetQUICConnectionsProtocolError gets the value of QUICConnectionsProtocolError for the instance
func (instance *Win32_PerfFormattedData_QuicCounters_QUICPerformanceDiagnostics) GetPropertyQUICConnectionsProtocolError()(value uint64, err error) { 
    retValue, err := instance.GetProperty("QUICConnectionsProtocolError")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint64); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint64(valuetmp)

    return
}

// SetQUICConnectionsProtocolErrorsPersec sets the value of QUICConnectionsProtocolErrorsPersec for the instance
func (instance *Win32_PerfFormattedData_QuicCounters_QUICPerformanceDiagnostics) SetPropertyQUICConnectionsProtocolErrorsPersec(value uint32) (err error) { 
    return instance.SetProperty("QUICConnectionsProtocolErrorsPersec", (value))
}

// GetQUICConnectionsProtocolErrorsPersec gets the value of QUICConnectionsProtocolErrorsPersec for the instance
func (instance *Win32_PerfFormattedData_QuicCounters_QUICPerformanceDiagnostics) GetPropertyQUICConnectionsProtocolErrorsPersec()(value uint32, err error) { 
    retValue, err := instance.GetProperty("QUICConnectionsProtocolErrorsPersec")
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

// SetQUICConnectionsQueued sets the value of QUICConnectionsQueued for the instance
func (instance *Win32_PerfFormattedData_QuicCounters_QUICPerformanceDiagnostics) SetPropertyQUICConnectionsQueued(value uint64) (err error) { 
    return instance.SetProperty("QUICConnectionsQueued", (value))
}

// GetQUICConnectionsQueued gets the value of QUICConnectionsQueued for the instance
func (instance *Win32_PerfFormattedData_QuicCounters_QUICPerformanceDiagnostics) GetPropertyQUICConnectionsQueued()(value uint64, err error) { 
    retValue, err := instance.GetProperty("QUICConnectionsQueued")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint64); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint64(valuetmp)

    return
}

// SetQUICConnectionsRejected sets the value of QUICConnectionsRejected for the instance
func (instance *Win32_PerfFormattedData_QuicCounters_QUICPerformanceDiagnostics) SetPropertyQUICConnectionsRejected(value uint64) (err error) { 
    return instance.SetProperty("QUICConnectionsRejected", (value))
}

// GetQUICConnectionsRejected gets the value of QUICConnectionsRejected for the instance
func (instance *Win32_PerfFormattedData_QuicCounters_QUICPerformanceDiagnostics) GetPropertyQUICConnectionsRejected()(value uint64, err error) { 
    retValue, err := instance.GetProperty("QUICConnectionsRejected")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint64); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint64(valuetmp)

    return
}

// SetQUICConnectionsRejectedPersec sets the value of QUICConnectionsRejectedPersec for the instance
func (instance *Win32_PerfFormattedData_QuicCounters_QUICPerformanceDiagnostics) SetPropertyQUICConnectionsRejectedPersec(value uint32) (err error) { 
    return instance.SetProperty("QUICConnectionsRejectedPersec", (value))
}

// GetQUICConnectionsRejectedPersec gets the value of QUICConnectionsRejectedPersec for the instance
func (instance *Win32_PerfFormattedData_QuicCounters_QUICPerformanceDiagnostics) GetPropertyQUICConnectionsRejectedPersec()(value uint32, err error) { 
    retValue, err := instance.GetProperty("QUICConnectionsRejectedPersec")
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

// SetQUICConnectionsResumed sets the value of QUICConnectionsResumed for the instance
func (instance *Win32_PerfFormattedData_QuicCounters_QUICPerformanceDiagnostics) SetPropertyQUICConnectionsResumed(value uint64) (err error) { 
    return instance.SetProperty("QUICConnectionsResumed", (value))
}

// GetQUICConnectionsResumed gets the value of QUICConnectionsResumed for the instance
func (instance *Win32_PerfFormattedData_QuicCounters_QUICPerformanceDiagnostics) GetPropertyQUICConnectionsResumed()(value uint64, err error) { 
    retValue, err := instance.GetProperty("QUICConnectionsResumed")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint64); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint64(valuetmp)

    return
}

// SetQUICConnectionsResumedPersec sets the value of QUICConnectionsResumedPersec for the instance
func (instance *Win32_PerfFormattedData_QuicCounters_QUICPerformanceDiagnostics) SetPropertyQUICConnectionsResumedPersec(value uint32) (err error) { 
    return instance.SetProperty("QUICConnectionsResumedPersec", (value))
}

// GetQUICConnectionsResumedPersec gets the value of QUICConnectionsResumedPersec for the instance
func (instance *Win32_PerfFormattedData_QuicCounters_QUICPerformanceDiagnostics) GetPropertyQUICConnectionsResumedPersec()(value uint32, err error) { 
    retValue, err := instance.GetProperty("QUICConnectionsResumedPersec")
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

// SetQUICHandshakesFailed sets the value of QUICHandshakesFailed for the instance
func (instance *Win32_PerfFormattedData_QuicCounters_QUICPerformanceDiagnostics) SetPropertyQUICHandshakesFailed(value uint64) (err error) { 
    return instance.SetProperty("QUICHandshakesFailed", (value))
}

// GetQUICHandshakesFailed gets the value of QUICHandshakesFailed for the instance
func (instance *Win32_PerfFormattedData_QuicCounters_QUICPerformanceDiagnostics) GetPropertyQUICHandshakesFailed()(value uint64, err error) { 
    retValue, err := instance.GetProperty("QUICHandshakesFailed")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint64); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint64(valuetmp)

    return
}

// SetQUICHandshakesFailedPersec sets the value of QUICHandshakesFailedPersec for the instance
func (instance *Win32_PerfFormattedData_QuicCounters_QUICPerformanceDiagnostics) SetPropertyQUICHandshakesFailedPersec(value uint32) (err error) { 
    return instance.SetProperty("QUICHandshakesFailedPersec", (value))
}

// GetQUICHandshakesFailedPersec gets the value of QUICHandshakesFailedPersec for the instance
func (instance *Win32_PerfFormattedData_QuicCounters_QUICPerformanceDiagnostics) GetPropertyQUICHandshakesFailedPersec()(value uint32, err error) { 
    retValue, err := instance.GetProperty("QUICHandshakesFailedPersec")
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

// SetQUICPacketDecryptionFailuresPersec sets the value of QUICPacketDecryptionFailuresPersec for the instance
func (instance *Win32_PerfFormattedData_QuicCounters_QUICPerformanceDiagnostics) SetPropertyQUICPacketDecryptionFailuresPersec(value uint32) (err error) { 
    return instance.SetProperty("QUICPacketDecryptionFailuresPersec", (value))
}

// GetQUICPacketDecryptionFailuresPersec gets the value of QUICPacketDecryptionFailuresPersec for the instance
func (instance *Win32_PerfFormattedData_QuicCounters_QUICPerformanceDiagnostics) GetPropertyQUICPacketDecryptionFailuresPersec()(value uint32, err error) { 
    retValue, err := instance.GetProperty("QUICPacketDecryptionFailuresPersec")
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

// SetQUICPacketsDroppedPersec sets the value of QUICPacketsDroppedPersec for the instance
func (instance *Win32_PerfFormattedData_QuicCounters_QUICPerformanceDiagnostics) SetPropertyQUICPacketsDroppedPersec(value uint32) (err error) { 
    return instance.SetProperty("QUICPacketsDroppedPersec", (value))
}

// GetQUICPacketsDroppedPersec gets the value of QUICPacketsDroppedPersec for the instance
func (instance *Win32_PerfFormattedData_QuicCounters_QUICPerformanceDiagnostics) GetPropertyQUICPacketsDroppedPersec()(value uint32, err error) { 
    retValue, err := instance.GetProperty("QUICPacketsDroppedPersec")
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

// SetQUICPacketsSuspectedLostPersec sets the value of QUICPacketsSuspectedLostPersec for the instance
func (instance *Win32_PerfFormattedData_QuicCounters_QUICPerformanceDiagnostics) SetPropertyQUICPacketsSuspectedLostPersec(value uint32) (err error) { 
    return instance.SetProperty("QUICPacketsSuspectedLostPersec", (value))
}

// GetQUICPacketsSuspectedLostPersec gets the value of QUICPacketsSuspectedLostPersec for the instance
func (instance *Win32_PerfFormattedData_QuicCounters_QUICPerformanceDiagnostics) GetPropertyQUICPacketsSuspectedLostPersec()(value uint32, err error) { 
    retValue, err := instance.GetProperty("QUICPacketsSuspectedLostPersec")
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

// SetQUICPathChallengesFailed sets the value of QUICPathChallengesFailed for the instance
func (instance *Win32_PerfFormattedData_QuicCounters_QUICPerformanceDiagnostics) SetPropertyQUICPathChallengesFailed(value uint64) (err error) { 
    return instance.SetProperty("QUICPathChallengesFailed", (value))
}

// GetQUICPathChallengesFailed gets the value of QUICPathChallengesFailed for the instance
func (instance *Win32_PerfFormattedData_QuicCounters_QUICPerformanceDiagnostics) GetPropertyQUICPathChallengesFailed()(value uint64, err error) { 
    retValue, err := instance.GetProperty("QUICPathChallengesFailed")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint64); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint64(valuetmp)

    return
}

// SetQUICPathChallengesSucceeded sets the value of QUICPathChallengesSucceeded for the instance
func (instance *Win32_PerfFormattedData_QuicCounters_QUICPerformanceDiagnostics) SetPropertyQUICPathChallengesSucceeded(value uint64) (err error) { 
    return instance.SetProperty("QUICPathChallengesSucceeded", (value))
}

// GetQUICPathChallengesSucceeded gets the value of QUICPathChallengesSucceeded for the instance
func (instance *Win32_PerfFormattedData_QuicCounters_QUICPerformanceDiagnostics) GetPropertyQUICPathChallengesSucceeded()(value uint64, err error) { 
    retValue, err := instance.GetProperty("QUICPathChallengesSucceeded")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint64); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint64(valuetmp)

    return
}

// SetQUICStatelessResetsSent sets the value of QUICStatelessResetsSent for the instance
func (instance *Win32_PerfFormattedData_QuicCounters_QUICPerformanceDiagnostics) SetPropertyQUICStatelessResetsSent(value uint64) (err error) { 
    return instance.SetProperty("QUICStatelessResetsSent", (value))
}

// GetQUICStatelessResetsSent gets the value of QUICStatelessResetsSent for the instance
func (instance *Win32_PerfFormattedData_QuicCounters_QUICPerformanceDiagnostics) GetPropertyQUICStatelessResetsSent()(value uint64, err error) { 
    retValue, err := instance.GetProperty("QUICStatelessResetsSent")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint64); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint64(valuetmp)

    return
}

// SetQUICStatelessRetriesSent sets the value of QUICStatelessRetriesSent for the instance
func (instance *Win32_PerfFormattedData_QuicCounters_QUICPerformanceDiagnostics) SetPropertyQUICStatelessRetriesSent(value uint64) (err error) { 
    return instance.SetProperty("QUICStatelessRetriesSent", (value))
}

// GetQUICStatelessRetriesSent gets the value of QUICStatelessRetriesSent for the instance
func (instance *Win32_PerfFormattedData_QuicCounters_QUICPerformanceDiagnostics) GetPropertyQUICStatelessRetriesSent()(value uint64, err error) { 
    retValue, err := instance.GetProperty("QUICStatelessRetriesSent")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint64); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint64(valuetmp)

    return
}

// SetQUICStreamsActive sets the value of QUICStreamsActive for the instance
func (instance *Win32_PerfFormattedData_QuicCounters_QUICPerformanceDiagnostics) SetPropertyQUICStreamsActive(value uint64) (err error) { 
    return instance.SetProperty("QUICStreamsActive", (value))
}

// GetQUICStreamsActive gets the value of QUICStreamsActive for the instance
func (instance *Win32_PerfFormattedData_QuicCounters_QUICPerformanceDiagnostics) GetPropertyQUICStreamsActive()(value uint64, err error) { 
    retValue, err := instance.GetProperty("QUICStreamsActive")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint64); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint64(valuetmp)

    return
}

// SetQUICUDPDatagramsReceivedPersec sets the value of QUICUDPDatagramsReceivedPersec for the instance
func (instance *Win32_PerfFormattedData_QuicCounters_QUICPerformanceDiagnostics) SetPropertyQUICUDPDatagramsReceivedPersec(value uint32) (err error) { 
    return instance.SetProperty("QUICUDPDatagramsReceivedPersec", (value))
}

// GetQUICUDPDatagramsReceivedPersec gets the value of QUICUDPDatagramsReceivedPersec for the instance
func (instance *Win32_PerfFormattedData_QuicCounters_QUICPerformanceDiagnostics) GetPropertyQUICUDPDatagramsReceivedPersec()(value uint32, err error) { 
    retValue, err := instance.GetProperty("QUICUDPDatagramsReceivedPersec")
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

// SetQUICUDPDatagramsSentPersec sets the value of QUICUDPDatagramsSentPersec for the instance
func (instance *Win32_PerfFormattedData_QuicCounters_QUICPerformanceDiagnostics) SetPropertyQUICUDPDatagramsSentPersec(value uint32) (err error) { 
    return instance.SetProperty("QUICUDPDatagramsSentPersec", (value))
}

// GetQUICUDPDatagramsSentPersec gets the value of QUICUDPDatagramsSentPersec for the instance
func (instance *Win32_PerfFormattedData_QuicCounters_QUICPerformanceDiagnostics) GetPropertyQUICUDPDatagramsSentPersec()(value uint32, err error) { 
    retValue, err := instance.GetProperty("QUICUDPDatagramsSentPersec")
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

// SetQUICUDPPayloadBytesReceivedPersec sets the value of QUICUDPPayloadBytesReceivedPersec for the instance
func (instance *Win32_PerfFormattedData_QuicCounters_QUICPerformanceDiagnostics) SetPropertyQUICUDPPayloadBytesReceivedPersec(value uint32) (err error) { 
    return instance.SetProperty("QUICUDPPayloadBytesReceivedPersec", (value))
}

// GetQUICUDPPayloadBytesReceivedPersec gets the value of QUICUDPPayloadBytesReceivedPersec for the instance
func (instance *Win32_PerfFormattedData_QuicCounters_QUICPerformanceDiagnostics) GetPropertyQUICUDPPayloadBytesReceivedPersec()(value uint32, err error) { 
    retValue, err := instance.GetProperty("QUICUDPPayloadBytesReceivedPersec")
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

// SetQUICUDPPayloadBytesSentPersec sets the value of QUICUDPPayloadBytesSentPersec for the instance
func (instance *Win32_PerfFormattedData_QuicCounters_QUICPerformanceDiagnostics) SetPropertyQUICUDPPayloadBytesSentPersec(value uint32) (err error) { 
    return instance.SetProperty("QUICUDPPayloadBytesSentPersec", (value))
}

// GetQUICUDPPayloadBytesSentPersec gets the value of QUICUDPPayloadBytesSentPersec for the instance
func (instance *Win32_PerfFormattedData_QuicCounters_QUICPerformanceDiagnostics) GetPropertyQUICUDPPayloadBytesSentPersec()(value uint32, err error) { 
    retValue, err := instance.GetProperty("QUICUDPPayloadBytesSentPersec")
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

// SetQUICUDPReceiveEventsPersec sets the value of QUICUDPReceiveEventsPersec for the instance
func (instance *Win32_PerfFormattedData_QuicCounters_QUICPerformanceDiagnostics) SetPropertyQUICUDPReceiveEventsPersec(value uint32) (err error) { 
    return instance.SetProperty("QUICUDPReceiveEventsPersec", (value))
}

// GetQUICUDPReceiveEventsPersec gets the value of QUICUDPReceiveEventsPersec for the instance
func (instance *Win32_PerfFormattedData_QuicCounters_QUICPerformanceDiagnostics) GetPropertyQUICUDPReceiveEventsPersec()(value uint32, err error) { 
    retValue, err := instance.GetProperty("QUICUDPReceiveEventsPersec")
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

// SetQUICUDPSendCallsPersec sets the value of QUICUDPSendCallsPersec for the instance
func (instance *Win32_PerfFormattedData_QuicCounters_QUICPerformanceDiagnostics) SetPropertyQUICUDPSendCallsPersec(value uint32) (err error) { 
    return instance.SetProperty("QUICUDPSendCallsPersec", (value))
}

// GetQUICUDPSendCallsPersec gets the value of QUICUDPSendCallsPersec for the instance
func (instance *Win32_PerfFormattedData_QuicCounters_QUICPerformanceDiagnostics) GetPropertyQUICUDPSendCallsPersec()(value uint32, err error) { 
    retValue, err := instance.GetProperty("QUICUDPSendCallsPersec")
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

// SetQUICWorkerOperationsPersec sets the value of QUICWorkerOperationsPersec for the instance
func (instance *Win32_PerfFormattedData_QuicCounters_QUICPerformanceDiagnostics) SetPropertyQUICWorkerOperationsPersec(value uint32) (err error) { 
    return instance.SetProperty("QUICWorkerOperationsPersec", (value))
}

// GetQUICWorkerOperationsPersec gets the value of QUICWorkerOperationsPersec for the instance
func (instance *Win32_PerfFormattedData_QuicCounters_QUICPerformanceDiagnostics) GetPropertyQUICWorkerOperationsPersec()(value uint32, err error) { 
    retValue, err := instance.GetProperty("QUICWorkerOperationsPersec")
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

// SetQUICWorkerOperationsQueued sets the value of QUICWorkerOperationsQueued for the instance
func (instance *Win32_PerfFormattedData_QuicCounters_QUICPerformanceDiagnostics) SetPropertyQUICWorkerOperationsQueued(value uint64) (err error) { 
    return instance.SetProperty("QUICWorkerOperationsQueued", (value))
}

// GetQUICWorkerOperationsQueued gets the value of QUICWorkerOperationsQueued for the instance
func (instance *Win32_PerfFormattedData_QuicCounters_QUICPerformanceDiagnostics) GetPropertyQUICWorkerOperationsQueued()(value uint64, err error) { 
    retValue, err := instance.GetProperty("QUICWorkerOperationsQueued")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint64); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint64(valuetmp)

    return
}

// SetQUICWorkerOperationsQueuedPersec sets the value of QUICWorkerOperationsQueuedPersec for the instance
func (instance *Win32_PerfFormattedData_QuicCounters_QUICPerformanceDiagnostics) SetPropertyQUICWorkerOperationsQueuedPersec(value uint32) (err error) { 
    return instance.SetProperty("QUICWorkerOperationsQueuedPersec", (value))
}

// GetQUICWorkerOperationsQueuedPersec gets the value of QUICWorkerOperationsQueuedPersec for the instance
func (instance *Win32_PerfFormattedData_QuicCounters_QUICPerformanceDiagnostics) GetPropertyQUICWorkerOperationsQueuedPersec()(value uint32, err error) { 
    retValue, err := instance.GetProperty("QUICWorkerOperationsQueuedPersec")
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

