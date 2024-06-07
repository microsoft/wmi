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

// Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2TransmitDatapathCounters struct
type Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2TransmitDatapathCounters struct { 
	*Win32_PerfFormattedData

	// 
	Averagepacketcountperindicate uint64

	// 
	CpuNumber uint64

	// 
	Dropsduetocompletionqueueerrors uint64

	// 
	Dropsduetoinvalidpacketsize uint64

	// 
	Interruptsonincorrectcpu uint64

	// 
	Numberofpacketscompleted uint64

	// 
	Numberofpacketsposted uint64

	// 
	Numberofpacketspostedinbypassmode uint64

	// 
	Numberoftrafficprofiletransitions uint64

	// 
	OScalltobuildSGLfailed uint64

	// 
	Packetsprocessedininterruptmode uint64

	// 
	PacketsprocessedinNDISpollmode uint64

	// 
	Packetsprocessedinpollingmode uint64

	// 
	Transmitcopypackets uint64

	// 
	Transmitringisfull uint64
}

	func NewWin32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2TransmitDatapathCountersEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2TransmitDatapathCounters, err error) {tmp, err := NewWin32_PerfFormattedDataEx1(instance)
		
	if err != nil { return }
	newInstance = &Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2TransmitDatapathCounters {
	Win32_PerfFormattedData: tmp,
	}
	return
	}
	

	func NewWin32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2TransmitDatapathCountersEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2TransmitDatapathCounters, err error) {tmp, err := NewWin32_PerfFormattedDataEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2TransmitDatapathCounters {
	Win32_PerfFormattedData: tmp,
	}
	return
	}
	

// SetAveragepacketcountperindicate sets the value of Averagepacketcountperindicate for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2TransmitDatapathCounters) SetPropertyAveragepacketcountperindicate(value uint64) (err error) { 
    return instance.SetProperty("Averagepacketcountperindicate", (value))
}

// GetAveragepacketcountperindicate gets the value of Averagepacketcountperindicate for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2TransmitDatapathCounters) GetPropertyAveragepacketcountperindicate()(value uint64, err error) { 
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

// SetCpuNumber sets the value of CpuNumber for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2TransmitDatapathCounters) SetPropertyCpuNumber(value uint64) (err error) { 
    return instance.SetProperty("CpuNumber", (value))
}

// GetCpuNumber gets the value of CpuNumber for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2TransmitDatapathCounters) GetPropertyCpuNumber()(value uint64, err error) { 
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

// SetDropsduetocompletionqueueerrors sets the value of Dropsduetocompletionqueueerrors for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2TransmitDatapathCounters) SetPropertyDropsduetocompletionqueueerrors(value uint64) (err error) { 
    return instance.SetProperty("Dropsduetocompletionqueueerrors", (value))
}

// GetDropsduetocompletionqueueerrors gets the value of Dropsduetocompletionqueueerrors for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2TransmitDatapathCounters) GetPropertyDropsduetocompletionqueueerrors()(value uint64, err error) { 
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
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2TransmitDatapathCounters) SetPropertyDropsduetoinvalidpacketsize(value uint64) (err error) { 
    return instance.SetProperty("Dropsduetoinvalidpacketsize", (value))
}

// GetDropsduetoinvalidpacketsize gets the value of Dropsduetoinvalidpacketsize for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2TransmitDatapathCounters) GetPropertyDropsduetoinvalidpacketsize()(value uint64, err error) { 
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

// SetInterruptsonincorrectcpu sets the value of Interruptsonincorrectcpu for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2TransmitDatapathCounters) SetPropertyInterruptsonincorrectcpu(value uint64) (err error) { 
    return instance.SetProperty("Interruptsonincorrectcpu", (value))
}

// GetInterruptsonincorrectcpu gets the value of Interruptsonincorrectcpu for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2TransmitDatapathCounters) GetPropertyInterruptsonincorrectcpu()(value uint64, err error) { 
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

// SetNumberofpacketscompleted sets the value of Numberofpacketscompleted for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2TransmitDatapathCounters) SetPropertyNumberofpacketscompleted(value uint64) (err error) { 
    return instance.SetProperty("Numberofpacketscompleted", (value))
}

// GetNumberofpacketscompleted gets the value of Numberofpacketscompleted for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2TransmitDatapathCounters) GetPropertyNumberofpacketscompleted()(value uint64, err error) { 
    retValue, err := instance.GetProperty("Numberofpacketscompleted")
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

// SetNumberofpacketsposted sets the value of Numberofpacketsposted for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2TransmitDatapathCounters) SetPropertyNumberofpacketsposted(value uint64) (err error) { 
    return instance.SetProperty("Numberofpacketsposted", (value))
}

// GetNumberofpacketsposted gets the value of Numberofpacketsposted for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2TransmitDatapathCounters) GetPropertyNumberofpacketsposted()(value uint64, err error) { 
    retValue, err := instance.GetProperty("Numberofpacketsposted")
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

// SetNumberofpacketspostedinbypassmode sets the value of Numberofpacketspostedinbypassmode for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2TransmitDatapathCounters) SetPropertyNumberofpacketspostedinbypassmode(value uint64) (err error) { 
    return instance.SetProperty("Numberofpacketspostedinbypassmode", (value))
}

// GetNumberofpacketspostedinbypassmode gets the value of Numberofpacketspostedinbypassmode for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2TransmitDatapathCounters) GetPropertyNumberofpacketspostedinbypassmode()(value uint64, err error) { 
    retValue, err := instance.GetProperty("Numberofpacketspostedinbypassmode")
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
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2TransmitDatapathCounters) SetPropertyNumberoftrafficprofiletransitions(value uint64) (err error) { 
    return instance.SetProperty("Numberoftrafficprofiletransitions", (value))
}

// GetNumberoftrafficprofiletransitions gets the value of Numberoftrafficprofiletransitions for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2TransmitDatapathCounters) GetPropertyNumberoftrafficprofiletransitions()(value uint64, err error) { 
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

// SetOScalltobuildSGLfailed sets the value of OScalltobuildSGLfailed for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2TransmitDatapathCounters) SetPropertyOScalltobuildSGLfailed(value uint64) (err error) { 
    return instance.SetProperty("OScalltobuildSGLfailed", (value))
}

// GetOScalltobuildSGLfailed gets the value of OScalltobuildSGLfailed for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2TransmitDatapathCounters) GetPropertyOScalltobuildSGLfailed()(value uint64, err error) { 
    retValue, err := instance.GetProperty("OScalltobuildSGLfailed")
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
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2TransmitDatapathCounters) SetPropertyPacketsprocessedininterruptmode(value uint64) (err error) { 
    return instance.SetProperty("Packetsprocessedininterruptmode", (value))
}

// GetPacketsprocessedininterruptmode gets the value of Packetsprocessedininterruptmode for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2TransmitDatapathCounters) GetPropertyPacketsprocessedininterruptmode()(value uint64, err error) { 
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
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2TransmitDatapathCounters) SetPropertyPacketsprocessedinNDISpollmode(value uint64) (err error) { 
    return instance.SetProperty("PacketsprocessedinNDISpollmode", (value))
}

// GetPacketsprocessedinNDISpollmode gets the value of PacketsprocessedinNDISpollmode for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2TransmitDatapathCounters) GetPropertyPacketsprocessedinNDISpollmode()(value uint64, err error) { 
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
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2TransmitDatapathCounters) SetPropertyPacketsprocessedinpollingmode(value uint64) (err error) { 
    return instance.SetProperty("Packetsprocessedinpollingmode", (value))
}

// GetPacketsprocessedinpollingmode gets the value of Packetsprocessedinpollingmode for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2TransmitDatapathCounters) GetPropertyPacketsprocessedinpollingmode()(value uint64, err error) { 
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

// SetTransmitcopypackets sets the value of Transmitcopypackets for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2TransmitDatapathCounters) SetPropertyTransmitcopypackets(value uint64) (err error) { 
    return instance.SetProperty("Transmitcopypackets", (value))
}

// GetTransmitcopypackets gets the value of Transmitcopypackets for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2TransmitDatapathCounters) GetPropertyTransmitcopypackets()(value uint64, err error) { 
    retValue, err := instance.GetProperty("Transmitcopypackets")
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

// SetTransmitringisfull sets the value of Transmitringisfull for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2TransmitDatapathCounters) SetPropertyTransmitringisfull(value uint64) (err error) { 
    return instance.SetProperty("Transmitringisfull", (value))
}

// GetTransmitringisfull gets the value of Transmitringisfull for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2TransmitDatapathCounters) GetPropertyTransmitringisfull()(value uint64, err error) { 
    retValue, err := instance.GetProperty("Transmitringisfull")
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

