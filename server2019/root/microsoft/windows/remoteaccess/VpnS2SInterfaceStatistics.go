// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.Microsoft.Windows.RemoteAccess
//////////////////////////////////////////////
package remoteaccess
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/base/instance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// VpnS2SInterfaceStatistics struct
type VpnS2SInterfaceStatistics struct { 
	*cim.WmiInstance

	// 
	AlignmentErrors uint32

	// 
	BufferOverrunErrors uint32

	// 
	BytesReceived uint64

	// 
	BytesTransmitted uint64

	// 
	ConnectionUpTime uint32

	// 
	CrcErrors uint32

	// 
	CumulativeBytesReceived uint64

	// 
	CumulativeBytesTransmitted uint64

	// 
	FramesReceived uint32

	// 
	FramesTransmitted uint32

	// 
	FramingErrors uint32

	// 
	HardwareOverrunErrors uint32

	// 
	LastClearedTime string

	// 
	Name string

	// 
	RoutingDomain string

	// 
	RxRateKbps uint32

	// 
	RxRateLimitedPacketsDropped uint32

	// 
	RxTotalPacketsDropped uint32

	// 
	TimeoutErrors uint32

	// 
	TxRateKbps uint32

	// 
	TxRateLimitedPacketsDropped uint32

	// 
	TxTotalPacketsDropped uint32
}

	func NewVpnS2SInterfaceStatisticsEx1(instance *cim.WmiInstance) (newInstance *VpnS2SInterfaceStatistics, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &VpnS2SInterfaceStatistics {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewVpnS2SInterfaceStatisticsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *VpnS2SInterfaceStatistics, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &VpnS2SInterfaceStatistics {
	WmiInstance: tmp,
	}
	return
	}
	

// SetAlignmentErrors sets the value of AlignmentErrors for the instance
func (instance *VpnS2SInterfaceStatistics) SetPropertyAlignmentErrors(value uint32) (err error) { 
    return instance.SetProperty("AlignmentErrors", (value))
}

// GetAlignmentErrors gets the value of AlignmentErrors for the instance
func (instance *VpnS2SInterfaceStatistics) GetPropertyAlignmentErrors()(value uint32, err error) { 
    retValue, err := instance.GetProperty("AlignmentErrors")
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

// SetBufferOverrunErrors sets the value of BufferOverrunErrors for the instance
func (instance *VpnS2SInterfaceStatistics) SetPropertyBufferOverrunErrors(value uint32) (err error) { 
    return instance.SetProperty("BufferOverrunErrors", (value))
}

// GetBufferOverrunErrors gets the value of BufferOverrunErrors for the instance
func (instance *VpnS2SInterfaceStatistics) GetPropertyBufferOverrunErrors()(value uint32, err error) { 
    retValue, err := instance.GetProperty("BufferOverrunErrors")
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

// SetBytesReceived sets the value of BytesReceived for the instance
func (instance *VpnS2SInterfaceStatistics) SetPropertyBytesReceived(value uint64) (err error) { 
    return instance.SetProperty("BytesReceived", (value))
}

// GetBytesReceived gets the value of BytesReceived for the instance
func (instance *VpnS2SInterfaceStatistics) GetPropertyBytesReceived()(value uint64, err error) { 
    retValue, err := instance.GetProperty("BytesReceived")
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

// SetBytesTransmitted sets the value of BytesTransmitted for the instance
func (instance *VpnS2SInterfaceStatistics) SetPropertyBytesTransmitted(value uint64) (err error) { 
    return instance.SetProperty("BytesTransmitted", (value))
}

// GetBytesTransmitted gets the value of BytesTransmitted for the instance
func (instance *VpnS2SInterfaceStatistics) GetPropertyBytesTransmitted()(value uint64, err error) { 
    retValue, err := instance.GetProperty("BytesTransmitted")
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

// SetConnectionUpTime sets the value of ConnectionUpTime for the instance
func (instance *VpnS2SInterfaceStatistics) SetPropertyConnectionUpTime(value uint32) (err error) { 
    return instance.SetProperty("ConnectionUpTime", (value))
}

// GetConnectionUpTime gets the value of ConnectionUpTime for the instance
func (instance *VpnS2SInterfaceStatistics) GetPropertyConnectionUpTime()(value uint32, err error) { 
    retValue, err := instance.GetProperty("ConnectionUpTime")
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

// SetCrcErrors sets the value of CrcErrors for the instance
func (instance *VpnS2SInterfaceStatistics) SetPropertyCrcErrors(value uint32) (err error) { 
    return instance.SetProperty("CrcErrors", (value))
}

// GetCrcErrors gets the value of CrcErrors for the instance
func (instance *VpnS2SInterfaceStatistics) GetPropertyCrcErrors()(value uint32, err error) { 
    retValue, err := instance.GetProperty("CrcErrors")
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

// SetCumulativeBytesReceived sets the value of CumulativeBytesReceived for the instance
func (instance *VpnS2SInterfaceStatistics) SetPropertyCumulativeBytesReceived(value uint64) (err error) { 
    return instance.SetProperty("CumulativeBytesReceived", (value))
}

// GetCumulativeBytesReceived gets the value of CumulativeBytesReceived for the instance
func (instance *VpnS2SInterfaceStatistics) GetPropertyCumulativeBytesReceived()(value uint64, err error) { 
    retValue, err := instance.GetProperty("CumulativeBytesReceived")
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

// SetCumulativeBytesTransmitted sets the value of CumulativeBytesTransmitted for the instance
func (instance *VpnS2SInterfaceStatistics) SetPropertyCumulativeBytesTransmitted(value uint64) (err error) { 
    return instance.SetProperty("CumulativeBytesTransmitted", (value))
}

// GetCumulativeBytesTransmitted gets the value of CumulativeBytesTransmitted for the instance
func (instance *VpnS2SInterfaceStatistics) GetPropertyCumulativeBytesTransmitted()(value uint64, err error) { 
    retValue, err := instance.GetProperty("CumulativeBytesTransmitted")
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

// SetFramesReceived sets the value of FramesReceived for the instance
func (instance *VpnS2SInterfaceStatistics) SetPropertyFramesReceived(value uint32) (err error) { 
    return instance.SetProperty("FramesReceived", (value))
}

// GetFramesReceived gets the value of FramesReceived for the instance
func (instance *VpnS2SInterfaceStatistics) GetPropertyFramesReceived()(value uint32, err error) { 
    retValue, err := instance.GetProperty("FramesReceived")
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

// SetFramesTransmitted sets the value of FramesTransmitted for the instance
func (instance *VpnS2SInterfaceStatistics) SetPropertyFramesTransmitted(value uint32) (err error) { 
    return instance.SetProperty("FramesTransmitted", (value))
}

// GetFramesTransmitted gets the value of FramesTransmitted for the instance
func (instance *VpnS2SInterfaceStatistics) GetPropertyFramesTransmitted()(value uint32, err error) { 
    retValue, err := instance.GetProperty("FramesTransmitted")
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

// SetFramingErrors sets the value of FramingErrors for the instance
func (instance *VpnS2SInterfaceStatistics) SetPropertyFramingErrors(value uint32) (err error) { 
    return instance.SetProperty("FramingErrors", (value))
}

// GetFramingErrors gets the value of FramingErrors for the instance
func (instance *VpnS2SInterfaceStatistics) GetPropertyFramingErrors()(value uint32, err error) { 
    retValue, err := instance.GetProperty("FramingErrors")
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

// SetHardwareOverrunErrors sets the value of HardwareOverrunErrors for the instance
func (instance *VpnS2SInterfaceStatistics) SetPropertyHardwareOverrunErrors(value uint32) (err error) { 
    return instance.SetProperty("HardwareOverrunErrors", (value))
}

// GetHardwareOverrunErrors gets the value of HardwareOverrunErrors for the instance
func (instance *VpnS2SInterfaceStatistics) GetPropertyHardwareOverrunErrors()(value uint32, err error) { 
    retValue, err := instance.GetProperty("HardwareOverrunErrors")
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

// SetLastClearedTime sets the value of LastClearedTime for the instance
func (instance *VpnS2SInterfaceStatistics) SetPropertyLastClearedTime(value string) (err error) { 
    return instance.SetProperty("LastClearedTime", (value))
}

// GetLastClearedTime gets the value of LastClearedTime for the instance
func (instance *VpnS2SInterfaceStatistics) GetPropertyLastClearedTime()(value string, err error) { 
    retValue, err := instance.GetProperty("LastClearedTime")
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

// SetName sets the value of Name for the instance
func (instance *VpnS2SInterfaceStatistics) SetPropertyName(value string) (err error) { 
    return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *VpnS2SInterfaceStatistics) GetPropertyName()(value string, err error) { 
    retValue, err := instance.GetProperty("Name")
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

// SetRoutingDomain sets the value of RoutingDomain for the instance
func (instance *VpnS2SInterfaceStatistics) SetPropertyRoutingDomain(value string) (err error) { 
    return instance.SetProperty("RoutingDomain", (value))
}

// GetRoutingDomain gets the value of RoutingDomain for the instance
func (instance *VpnS2SInterfaceStatistics) GetPropertyRoutingDomain()(value string, err error) { 
    retValue, err := instance.GetProperty("RoutingDomain")
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

// SetRxRateKbps sets the value of RxRateKbps for the instance
func (instance *VpnS2SInterfaceStatistics) SetPropertyRxRateKbps(value uint32) (err error) { 
    return instance.SetProperty("RxRateKbps", (value))
}

// GetRxRateKbps gets the value of RxRateKbps for the instance
func (instance *VpnS2SInterfaceStatistics) GetPropertyRxRateKbps()(value uint32, err error) { 
    retValue, err := instance.GetProperty("RxRateKbps")
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

// SetRxRateLimitedPacketsDropped sets the value of RxRateLimitedPacketsDropped for the instance
func (instance *VpnS2SInterfaceStatistics) SetPropertyRxRateLimitedPacketsDropped(value uint32) (err error) { 
    return instance.SetProperty("RxRateLimitedPacketsDropped", (value))
}

// GetRxRateLimitedPacketsDropped gets the value of RxRateLimitedPacketsDropped for the instance
func (instance *VpnS2SInterfaceStatistics) GetPropertyRxRateLimitedPacketsDropped()(value uint32, err error) { 
    retValue, err := instance.GetProperty("RxRateLimitedPacketsDropped")
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

// SetRxTotalPacketsDropped sets the value of RxTotalPacketsDropped for the instance
func (instance *VpnS2SInterfaceStatistics) SetPropertyRxTotalPacketsDropped(value uint32) (err error) { 
    return instance.SetProperty("RxTotalPacketsDropped", (value))
}

// GetRxTotalPacketsDropped gets the value of RxTotalPacketsDropped for the instance
func (instance *VpnS2SInterfaceStatistics) GetPropertyRxTotalPacketsDropped()(value uint32, err error) { 
    retValue, err := instance.GetProperty("RxTotalPacketsDropped")
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

// SetTimeoutErrors sets the value of TimeoutErrors for the instance
func (instance *VpnS2SInterfaceStatistics) SetPropertyTimeoutErrors(value uint32) (err error) { 
    return instance.SetProperty("TimeoutErrors", (value))
}

// GetTimeoutErrors gets the value of TimeoutErrors for the instance
func (instance *VpnS2SInterfaceStatistics) GetPropertyTimeoutErrors()(value uint32, err error) { 
    retValue, err := instance.GetProperty("TimeoutErrors")
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

// SetTxRateKbps sets the value of TxRateKbps for the instance
func (instance *VpnS2SInterfaceStatistics) SetPropertyTxRateKbps(value uint32) (err error) { 
    return instance.SetProperty("TxRateKbps", (value))
}

// GetTxRateKbps gets the value of TxRateKbps for the instance
func (instance *VpnS2SInterfaceStatistics) GetPropertyTxRateKbps()(value uint32, err error) { 
    retValue, err := instance.GetProperty("TxRateKbps")
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

// SetTxRateLimitedPacketsDropped sets the value of TxRateLimitedPacketsDropped for the instance
func (instance *VpnS2SInterfaceStatistics) SetPropertyTxRateLimitedPacketsDropped(value uint32) (err error) { 
    return instance.SetProperty("TxRateLimitedPacketsDropped", (value))
}

// GetTxRateLimitedPacketsDropped gets the value of TxRateLimitedPacketsDropped for the instance
func (instance *VpnS2SInterfaceStatistics) GetPropertyTxRateLimitedPacketsDropped()(value uint32, err error) { 
    retValue, err := instance.GetProperty("TxRateLimitedPacketsDropped")
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

// SetTxTotalPacketsDropped sets the value of TxTotalPacketsDropped for the instance
func (instance *VpnS2SInterfaceStatistics) SetPropertyTxTotalPacketsDropped(value uint32) (err error) { 
    return instance.SetProperty("TxTotalPacketsDropped", (value))
}

// GetTxTotalPacketsDropped gets the value of TxTotalPacketsDropped for the instance
func (instance *VpnS2SInterfaceStatistics) GetPropertyTxTotalPacketsDropped()(value uint32, err error) { 
    retValue, err := instance.GetProperty("TxTotalPacketsDropped")
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

