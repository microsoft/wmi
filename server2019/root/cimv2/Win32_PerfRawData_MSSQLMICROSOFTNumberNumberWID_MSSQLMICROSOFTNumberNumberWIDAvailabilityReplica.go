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

// Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDAvailabilityReplica struct
type Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDAvailabilityReplica struct { 
	*Win32_PerfRawData

	// 
	BytesReceivedfromReplicaPersec uint64

	// 
	BytesSenttoReplicaPersec uint64

	// 
	BytesSenttoTransportPersec uint64

	// 
	FlowControlPersec uint64

	// 
	FlowControlTimemsPersec uint64

	// 
	ReceivesfromReplicaPersec uint64

	// 
	ResentMessagesPersec uint64

	// 
	SendstoReplicaPersec uint64

	// 
	SendstoTransportPersec uint64
}

	func NewWin32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDAvailabilityReplicaEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDAvailabilityReplica, err error) {tmp, err := NewWin32_PerfRawDataEx1(instance)
		
	if err != nil { return }
	newInstance = &Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDAvailabilityReplica {
	Win32_PerfRawData: tmp,
	}
	return
	}
	

	func NewWin32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDAvailabilityReplicaEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDAvailabilityReplica, err error) {tmp, err := NewWin32_PerfRawDataEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDAvailabilityReplica {
	Win32_PerfRawData: tmp,
	}
	return
	}
	

// SetBytesReceivedfromReplicaPersec sets the value of BytesReceivedfromReplicaPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDAvailabilityReplica) SetPropertyBytesReceivedfromReplicaPersec(value uint64) (err error) { 
    return instance.SetProperty("BytesReceivedfromReplicaPersec", (value))
}

// GetBytesReceivedfromReplicaPersec gets the value of BytesReceivedfromReplicaPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDAvailabilityReplica) GetPropertyBytesReceivedfromReplicaPersec()(value uint64, err error) { 
    retValue, err := instance.GetProperty("BytesReceivedfromReplicaPersec")
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

// SetBytesSenttoReplicaPersec sets the value of BytesSenttoReplicaPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDAvailabilityReplica) SetPropertyBytesSenttoReplicaPersec(value uint64) (err error) { 
    return instance.SetProperty("BytesSenttoReplicaPersec", (value))
}

// GetBytesSenttoReplicaPersec gets the value of BytesSenttoReplicaPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDAvailabilityReplica) GetPropertyBytesSenttoReplicaPersec()(value uint64, err error) { 
    retValue, err := instance.GetProperty("BytesSenttoReplicaPersec")
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

// SetBytesSenttoTransportPersec sets the value of BytesSenttoTransportPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDAvailabilityReplica) SetPropertyBytesSenttoTransportPersec(value uint64) (err error) { 
    return instance.SetProperty("BytesSenttoTransportPersec", (value))
}

// GetBytesSenttoTransportPersec gets the value of BytesSenttoTransportPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDAvailabilityReplica) GetPropertyBytesSenttoTransportPersec()(value uint64, err error) { 
    retValue, err := instance.GetProperty("BytesSenttoTransportPersec")
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

// SetFlowControlPersec sets the value of FlowControlPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDAvailabilityReplica) SetPropertyFlowControlPersec(value uint64) (err error) { 
    return instance.SetProperty("FlowControlPersec", (value))
}

// GetFlowControlPersec gets the value of FlowControlPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDAvailabilityReplica) GetPropertyFlowControlPersec()(value uint64, err error) { 
    retValue, err := instance.GetProperty("FlowControlPersec")
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

// SetFlowControlTimemsPersec sets the value of FlowControlTimemsPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDAvailabilityReplica) SetPropertyFlowControlTimemsPersec(value uint64) (err error) { 
    return instance.SetProperty("FlowControlTimemsPersec", (value))
}

// GetFlowControlTimemsPersec gets the value of FlowControlTimemsPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDAvailabilityReplica) GetPropertyFlowControlTimemsPersec()(value uint64, err error) { 
    retValue, err := instance.GetProperty("FlowControlTimemsPersec")
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

// SetReceivesfromReplicaPersec sets the value of ReceivesfromReplicaPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDAvailabilityReplica) SetPropertyReceivesfromReplicaPersec(value uint64) (err error) { 
    return instance.SetProperty("ReceivesfromReplicaPersec", (value))
}

// GetReceivesfromReplicaPersec gets the value of ReceivesfromReplicaPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDAvailabilityReplica) GetPropertyReceivesfromReplicaPersec()(value uint64, err error) { 
    retValue, err := instance.GetProperty("ReceivesfromReplicaPersec")
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

// SetResentMessagesPersec sets the value of ResentMessagesPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDAvailabilityReplica) SetPropertyResentMessagesPersec(value uint64) (err error) { 
    return instance.SetProperty("ResentMessagesPersec", (value))
}

// GetResentMessagesPersec gets the value of ResentMessagesPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDAvailabilityReplica) GetPropertyResentMessagesPersec()(value uint64, err error) { 
    retValue, err := instance.GetProperty("ResentMessagesPersec")
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

// SetSendstoReplicaPersec sets the value of SendstoReplicaPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDAvailabilityReplica) SetPropertySendstoReplicaPersec(value uint64) (err error) { 
    return instance.SetProperty("SendstoReplicaPersec", (value))
}

// GetSendstoReplicaPersec gets the value of SendstoReplicaPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDAvailabilityReplica) GetPropertySendstoReplicaPersec()(value uint64, err error) { 
    retValue, err := instance.GetProperty("SendstoReplicaPersec")
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

// SetSendstoTransportPersec sets the value of SendstoTransportPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDAvailabilityReplica) SetPropertySendstoTransportPersec(value uint64) (err error) { 
    return instance.SetProperty("SendstoTransportPersec", (value))
}

// GetSendstoTransportPersec gets the value of SendstoTransportPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDAvailabilityReplica) GetPropertySendstoTransportPersec()(value uint64, err error) { 
    retValue, err := instance.GetProperty("SendstoTransportPersec")
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

