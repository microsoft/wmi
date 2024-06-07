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

// Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2ReceiveDatapathCounters struct
type Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2ReceiveDatapathCounters struct { 
	*Win32_PerfRawData

	// 
	Averagepacketcountperindicate uint64

	// 
	Consumedmaxreceives uint64

	// 
	CpuNumber uint64

	// 
	DpcWatchDogSingleDpcStarvation uint64

	// 
	DpcWatchDogTotalDpcStarvation uint64

	// 
	Dropsduetocompletionqueueerrors uint64

	// 
	Dropsduetoinvalidpacketsize uint64

	// 
	EcnMarkedPacketsIpv4 uint64

	// 
	EcnMarkedPacketsIpv6 uint64

	// 
	Interruptsonincorrectcpu uint64

	// 
	Numberofinterrupts uint64

	// 
	Numberofreceivebuffersposted uint64

	// 
	Numberoftrafficprofiletransitions uint64

	// 
	Packetsinlowresourcemode uint64

	// 
	Packetsprocessedininterruptmode uint64

	// 
	PacketsprocessedinNDISpollmode uint64

	// 
	Packetsprocessedinpollingmode uint64

	// 
	StridedWqes uint64
}

	func NewWin32_PerfRawData_Mlx5Provider_MellanoxWinOF2ReceiveDatapathCountersEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2ReceiveDatapathCounters, err error) {tmp, err := NewWin32_PerfRawDataEx1(instance)
		
	if err != nil { return }
	newInstance = &Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2ReceiveDatapathCounters {
	Win32_PerfRawData: tmp,
	}
	return
	}
	

	func NewWin32_PerfRawData_Mlx5Provider_MellanoxWinOF2ReceiveDatapathCountersEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2ReceiveDatapathCounters, err error) {tmp, err := NewWin32_PerfRawDataEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2ReceiveDatapathCounters {
	Win32_PerfRawData: tmp,
	}
	return
	}
	

// SetAveragepacketcountperindicate sets the value of Averagepacketcountperindicate for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2ReceiveDatapathCounters) SetPropertyAveragepacketcountperindicate(value uint64) (err error) { 
    return instance.SetProperty("Averagepacketcountperindicate", (value))
}

// GetAveragepacketcountperindicate gets the value of Averagepacketcountperindicate for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2ReceiveDatapathCounters) GetPropertyAveragepacketcountperindicate()(value uint64, err error) { 
    retValue, err := instance.GetProperty("Averagepacketcountperindicate")
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

// SetConsumedmaxreceives sets the value of Consumedmaxreceives for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2ReceiveDatapathCounters) SetPropertyConsumedmaxreceives(value uint64) (err error) { 
    return instance.SetProperty("Consumedmaxreceives", (value))
}

// GetConsumedmaxreceives gets the value of Consumedmaxreceives for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2ReceiveDatapathCounters) GetPropertyConsumedmaxreceives()(value uint64, err error) { 
    retValue, err := instance.GetProperty("Consumedmaxreceives")
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

// SetCpuNumber sets the value of CpuNumber for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2ReceiveDatapathCounters) SetPropertyCpuNumber(value uint64) (err error) { 
    return instance.SetProperty("CpuNumber", (value))
}

// GetCpuNumber gets the value of CpuNumber for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2ReceiveDatapathCounters) GetPropertyCpuNumber()(value uint64, err error) { 
    retValue, err := instance.GetProperty("CpuNumber")
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

// SetDpcWatchDogSingleDpcStarvation sets the value of DpcWatchDogSingleDpcStarvation for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2ReceiveDatapathCounters) SetPropertyDpcWatchDogSingleDpcStarvation(value uint64) (err error) { 
    return instance.SetProperty("DpcWatchDogSingleDpcStarvation", (value))
}

// GetDpcWatchDogSingleDpcStarvation gets the value of DpcWatchDogSingleDpcStarvation for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2ReceiveDatapathCounters) GetPropertyDpcWatchDogSingleDpcStarvation()(value uint64, err error) { 
    retValue, err := instance.GetProperty("DpcWatchDogSingleDpcStarvation")
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

// SetDpcWatchDogTotalDpcStarvation sets the value of DpcWatchDogTotalDpcStarvation for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2ReceiveDatapathCounters) SetPropertyDpcWatchDogTotalDpcStarvation(value uint64) (err error) { 
    return instance.SetProperty("DpcWatchDogTotalDpcStarvation", (value))
}

// GetDpcWatchDogTotalDpcStarvation gets the value of DpcWatchDogTotalDpcStarvation for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2ReceiveDatapathCounters) GetPropertyDpcWatchDogTotalDpcStarvation()(value uint64, err error) { 
    retValue, err := instance.GetProperty("DpcWatchDogTotalDpcStarvation")
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

// SetDropsduetocompletionqueueerrors sets the value of Dropsduetocompletionqueueerrors for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2ReceiveDatapathCounters) SetPropertyDropsduetocompletionqueueerrors(value uint64) (err error) { 
    return instance.SetProperty("Dropsduetocompletionqueueerrors", (value))
}

// GetDropsduetocompletionqueueerrors gets the value of Dropsduetocompletionqueueerrors for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2ReceiveDatapathCounters) GetPropertyDropsduetocompletionqueueerrors()(value uint64, err error) { 
    retValue, err := instance.GetProperty("Dropsduetocompletionqueueerrors")
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

// SetDropsduetoinvalidpacketsize sets the value of Dropsduetoinvalidpacketsize for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2ReceiveDatapathCounters) SetPropertyDropsduetoinvalidpacketsize(value uint64) (err error) { 
    return instance.SetProperty("Dropsduetoinvalidpacketsize", (value))
}

// GetDropsduetoinvalidpacketsize gets the value of Dropsduetoinvalidpacketsize for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2ReceiveDatapathCounters) GetPropertyDropsduetoinvalidpacketsize()(value uint64, err error) { 
    retValue, err := instance.GetProperty("Dropsduetoinvalidpacketsize")
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

// SetEcnMarkedPacketsIpv4 sets the value of EcnMarkedPacketsIpv4 for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2ReceiveDatapathCounters) SetPropertyEcnMarkedPacketsIpv4(value uint64) (err error) { 
    return instance.SetProperty("EcnMarkedPacketsIpv4", (value))
}

// GetEcnMarkedPacketsIpv4 gets the value of EcnMarkedPacketsIpv4 for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2ReceiveDatapathCounters) GetPropertyEcnMarkedPacketsIpv4()(value uint64, err error) { 
    retValue, err := instance.GetProperty("EcnMarkedPacketsIpv4")
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

// SetEcnMarkedPacketsIpv6 sets the value of EcnMarkedPacketsIpv6 for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2ReceiveDatapathCounters) SetPropertyEcnMarkedPacketsIpv6(value uint64) (err error) { 
    return instance.SetProperty("EcnMarkedPacketsIpv6", (value))
}

// GetEcnMarkedPacketsIpv6 gets the value of EcnMarkedPacketsIpv6 for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2ReceiveDatapathCounters) GetPropertyEcnMarkedPacketsIpv6()(value uint64, err error) { 
    retValue, err := instance.GetProperty("EcnMarkedPacketsIpv6")
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

// SetInterruptsonincorrectcpu sets the value of Interruptsonincorrectcpu for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2ReceiveDatapathCounters) SetPropertyInterruptsonincorrectcpu(value uint64) (err error) { 
    return instance.SetProperty("Interruptsonincorrectcpu", (value))
}

// GetInterruptsonincorrectcpu gets the value of Interruptsonincorrectcpu for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2ReceiveDatapathCounters) GetPropertyInterruptsonincorrectcpu()(value uint64, err error) { 
    retValue, err := instance.GetProperty("Interruptsonincorrectcpu")
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

// SetNumberofinterrupts sets the value of Numberofinterrupts for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2ReceiveDatapathCounters) SetPropertyNumberofinterrupts(value uint64) (err error) { 
    return instance.SetProperty("Numberofinterrupts", (value))
}

// GetNumberofinterrupts gets the value of Numberofinterrupts for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2ReceiveDatapathCounters) GetPropertyNumberofinterrupts()(value uint64, err error) { 
    retValue, err := instance.GetProperty("Numberofinterrupts")
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

// SetNumberofreceivebuffersposted sets the value of Numberofreceivebuffersposted for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2ReceiveDatapathCounters) SetPropertyNumberofreceivebuffersposted(value uint64) (err error) { 
    return instance.SetProperty("Numberofreceivebuffersposted", (value))
}

// GetNumberofreceivebuffersposted gets the value of Numberofreceivebuffersposted for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2ReceiveDatapathCounters) GetPropertyNumberofreceivebuffersposted()(value uint64, err error) { 
    retValue, err := instance.GetProperty("Numberofreceivebuffersposted")
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

// SetNumberoftrafficprofiletransitions sets the value of Numberoftrafficprofiletransitions for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2ReceiveDatapathCounters) SetPropertyNumberoftrafficprofiletransitions(value uint64) (err error) { 
    return instance.SetProperty("Numberoftrafficprofiletransitions", (value))
}

// GetNumberoftrafficprofiletransitions gets the value of Numberoftrafficprofiletransitions for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2ReceiveDatapathCounters) GetPropertyNumberoftrafficprofiletransitions()(value uint64, err error) { 
    retValue, err := instance.GetProperty("Numberoftrafficprofiletransitions")
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

// SetPacketsinlowresourcemode sets the value of Packetsinlowresourcemode for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2ReceiveDatapathCounters) SetPropertyPacketsinlowresourcemode(value uint64) (err error) { 
    return instance.SetProperty("Packetsinlowresourcemode", (value))
}

// GetPacketsinlowresourcemode gets the value of Packetsinlowresourcemode for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2ReceiveDatapathCounters) GetPropertyPacketsinlowresourcemode()(value uint64, err error) { 
    retValue, err := instance.GetProperty("Packetsinlowresourcemode")
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

// SetPacketsprocessedininterruptmode sets the value of Packetsprocessedininterruptmode for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2ReceiveDatapathCounters) SetPropertyPacketsprocessedininterruptmode(value uint64) (err error) { 
    return instance.SetProperty("Packetsprocessedininterruptmode", (value))
}

// GetPacketsprocessedininterruptmode gets the value of Packetsprocessedininterruptmode for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2ReceiveDatapathCounters) GetPropertyPacketsprocessedininterruptmode()(value uint64, err error) { 
    retValue, err := instance.GetProperty("Packetsprocessedininterruptmode")
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

// SetPacketsprocessedinNDISpollmode sets the value of PacketsprocessedinNDISpollmode for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2ReceiveDatapathCounters) SetPropertyPacketsprocessedinNDISpollmode(value uint64) (err error) { 
    return instance.SetProperty("PacketsprocessedinNDISpollmode", (value))
}

// GetPacketsprocessedinNDISpollmode gets the value of PacketsprocessedinNDISpollmode for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2ReceiveDatapathCounters) GetPropertyPacketsprocessedinNDISpollmode()(value uint64, err error) { 
    retValue, err := instance.GetProperty("PacketsprocessedinNDISpollmode")
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

// SetPacketsprocessedinpollingmode sets the value of Packetsprocessedinpollingmode for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2ReceiveDatapathCounters) SetPropertyPacketsprocessedinpollingmode(value uint64) (err error) { 
    return instance.SetProperty("Packetsprocessedinpollingmode", (value))
}

// GetPacketsprocessedinpollingmode gets the value of Packetsprocessedinpollingmode for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2ReceiveDatapathCounters) GetPropertyPacketsprocessedinpollingmode()(value uint64, err error) { 
    retValue, err := instance.GetProperty("Packetsprocessedinpollingmode")
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

// SetStridedWqes sets the value of StridedWqes for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2ReceiveDatapathCounters) SetPropertyStridedWqes(value uint64) (err error) { 
    return instance.SetProperty("StridedWqes", (value))
}

// GetStridedWqes gets the value of StridedWqes for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2ReceiveDatapathCounters) GetPropertyStridedWqes()(value uint64, err error) { 
    retValue, err := instance.GetProperty("StridedWqes")
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

