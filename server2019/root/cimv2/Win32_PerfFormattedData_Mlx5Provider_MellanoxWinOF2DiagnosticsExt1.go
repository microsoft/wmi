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

// Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DiagnosticsExt1 struct
type Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DiagnosticsExt1 struct { 
	*Win32_PerfFormattedData

	// 
	ChecksumcalculatedbySWPerPacket uint64

	// 
	CQOverrun uint64

	// 
	RoCEAdaptiveRetransmission uint32

	// 
	RoCEadaptiveretransmissiontimeouts uint32

	// 
	RoCESlowRestart uint32

	// 
	RoCESlowRestartCNPs uint32

	// 
	RoCESlowRestartTransmission uint32
}

	func NewWin32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DiagnosticsExt1Ex1(instance *cim.WmiInstance) (newInstance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DiagnosticsExt1, err error) {tmp, err := NewWin32_PerfFormattedDataEx1(instance)
		
	if err != nil { return }
	newInstance = &Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DiagnosticsExt1 {
	Win32_PerfFormattedData: tmp,
	}
	return
	}
	

	func NewWin32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DiagnosticsExt1Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DiagnosticsExt1, err error) {tmp, err := NewWin32_PerfFormattedDataEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DiagnosticsExt1 {
	Win32_PerfFormattedData: tmp,
	}
	return
	}
	

// SetChecksumcalculatedbySWPerPacket sets the value of ChecksumcalculatedbySWPerPacket for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DiagnosticsExt1) SetPropertyChecksumcalculatedbySWPerPacket(value uint64) (err error) { 
    return instance.SetProperty("ChecksumcalculatedbySWPerPacket", (value))
}

// GetChecksumcalculatedbySWPerPacket gets the value of ChecksumcalculatedbySWPerPacket for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DiagnosticsExt1) GetPropertyChecksumcalculatedbySWPerPacket()(value uint64, err error) { 
    retValue, err := instance.GetProperty("ChecksumcalculatedbySWPerPacket")
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

// SetCQOverrun sets the value of CQOverrun for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DiagnosticsExt1) SetPropertyCQOverrun(value uint64) (err error) { 
    return instance.SetProperty("CQOverrun", (value))
}

// GetCQOverrun gets the value of CQOverrun for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DiagnosticsExt1) GetPropertyCQOverrun()(value uint64, err error) { 
    retValue, err := instance.GetProperty("CQOverrun")
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

// SetRoCEAdaptiveRetransmission sets the value of RoCEAdaptiveRetransmission for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DiagnosticsExt1) SetPropertyRoCEAdaptiveRetransmission(value uint32) (err error) { 
    return instance.SetProperty("RoCEAdaptiveRetransmission", (value))
}

// GetRoCEAdaptiveRetransmission gets the value of RoCEAdaptiveRetransmission for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DiagnosticsExt1) GetPropertyRoCEAdaptiveRetransmission()(value uint32, err error) { 
    retValue, err := instance.GetProperty("RoCEAdaptiveRetransmission")
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

// SetRoCEadaptiveretransmissiontimeouts sets the value of RoCEadaptiveretransmissiontimeouts for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DiagnosticsExt1) SetPropertyRoCEadaptiveretransmissiontimeouts(value uint32) (err error) { 
    return instance.SetProperty("RoCEadaptiveretransmissiontimeouts", (value))
}

// GetRoCEadaptiveretransmissiontimeouts gets the value of RoCEadaptiveretransmissiontimeouts for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DiagnosticsExt1) GetPropertyRoCEadaptiveretransmissiontimeouts()(value uint32, err error) { 
    retValue, err := instance.GetProperty("RoCEadaptiveretransmissiontimeouts")
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

// SetRoCESlowRestart sets the value of RoCESlowRestart for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DiagnosticsExt1) SetPropertyRoCESlowRestart(value uint32) (err error) { 
    return instance.SetProperty("RoCESlowRestart", (value))
}

// GetRoCESlowRestart gets the value of RoCESlowRestart for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DiagnosticsExt1) GetPropertyRoCESlowRestart()(value uint32, err error) { 
    retValue, err := instance.GetProperty("RoCESlowRestart")
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

// SetRoCESlowRestartCNPs sets the value of RoCESlowRestartCNPs for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DiagnosticsExt1) SetPropertyRoCESlowRestartCNPs(value uint32) (err error) { 
    return instance.SetProperty("RoCESlowRestartCNPs", (value))
}

// GetRoCESlowRestartCNPs gets the value of RoCESlowRestartCNPs for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DiagnosticsExt1) GetPropertyRoCESlowRestartCNPs()(value uint32, err error) { 
    retValue, err := instance.GetProperty("RoCESlowRestartCNPs")
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

// SetRoCESlowRestartTransmission sets the value of RoCESlowRestartTransmission for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DiagnosticsExt1) SetPropertyRoCESlowRestartTransmission(value uint32) (err error) { 
    return instance.SetProperty("RoCESlowRestartTransmission", (value))
}

// GetRoCESlowRestartTransmission gets the value of RoCESlowRestartTransmission for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DiagnosticsExt1) GetPropertyRoCESlowRestartTransmission()(value uint32, err error) { 
    retValue, err := instance.GetProperty("RoCESlowRestartTransmission")
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

