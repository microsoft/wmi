// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.Microsoft.Windows.DeliveryOptimization
//////////////////////////////////////////////
package deliveryoptimization
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// MSFT_DOCurrentStatus struct
type MSFT_DOCurrentStatus struct { 
	*MSFT_DOBaseStatus

	// 
	CacheServerConnections uint32

	// 
	CacheSizeBytes uint64

	// 
	CdnConnections uint32

	// 
	CpuUsagePct float32

	// 
	DeviceProfile uint32

	// 
	DiskAvailableBytes uint64

	// 
	DiskTotalBytes uint64

	// 
	GroupConnections uint32

	// 
	InternetConnections uint32

	// 
	LanConnections uint32

	// 
	LinkLocalConnections uint32

	// 
	MemUsageKBytes uint64

	// 
	PeerInfoCount uint32

	// 
	Swarms uint32
}

	func NewMSFT_DOCurrentStatusEx1(instance *cim.WmiInstance) (newInstance *MSFT_DOCurrentStatus, err error) {tmp, err := NewMSFT_DOBaseStatusEx1(instance)
		
	if err != nil { return }
	newInstance = &MSFT_DOCurrentStatus {
	MSFT_DOBaseStatus: tmp,
	}
	return
	}
	

	func NewMSFT_DOCurrentStatusEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MSFT_DOCurrentStatus, err error) {tmp, err := NewMSFT_DOBaseStatusEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MSFT_DOCurrentStatus {
	MSFT_DOBaseStatus: tmp,
	}
	return
	}
	

// SetCacheServerConnections sets the value of CacheServerConnections for the instance
func (instance *MSFT_DOCurrentStatus) SetPropertyCacheServerConnections(value uint32) (err error) { 
    return instance.SetProperty("CacheServerConnections", (value))
}

// GetCacheServerConnections gets the value of CacheServerConnections for the instance
func (instance *MSFT_DOCurrentStatus) GetPropertyCacheServerConnections()(value uint32, err error) { 
    retValue, err := instance.GetProperty("CacheServerConnections")
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

// SetCacheSizeBytes sets the value of CacheSizeBytes for the instance
func (instance *MSFT_DOCurrentStatus) SetPropertyCacheSizeBytes(value uint64) (err error) { 
    return instance.SetProperty("CacheSizeBytes", (value))
}

// GetCacheSizeBytes gets the value of CacheSizeBytes for the instance
func (instance *MSFT_DOCurrentStatus) GetPropertyCacheSizeBytes()(value uint64, err error) { 
    retValue, err := instance.GetProperty("CacheSizeBytes")
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

// SetCdnConnections sets the value of CdnConnections for the instance
func (instance *MSFT_DOCurrentStatus) SetPropertyCdnConnections(value uint32) (err error) { 
    return instance.SetProperty("CdnConnections", (value))
}

// GetCdnConnections gets the value of CdnConnections for the instance
func (instance *MSFT_DOCurrentStatus) GetPropertyCdnConnections()(value uint32, err error) { 
    retValue, err := instance.GetProperty("CdnConnections")
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

// SetCpuUsagePct sets the value of CpuUsagePct for the instance
func (instance *MSFT_DOCurrentStatus) SetPropertyCpuUsagePct(value float32) (err error) { 
    return instance.SetProperty("CpuUsagePct", (value))
}

// GetCpuUsagePct gets the value of CpuUsagePct for the instance
func (instance *MSFT_DOCurrentStatus) GetPropertyCpuUsagePct()(value float32, err error) { 
    retValue, err := instance.GetProperty("CpuUsagePct")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(float32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " float32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = float32(valuetmp)

    return
}

// SetDeviceProfile sets the value of DeviceProfile for the instance
func (instance *MSFT_DOCurrentStatus) SetPropertyDeviceProfile(value uint32) (err error) { 
    return instance.SetProperty("DeviceProfile", (value))
}

// GetDeviceProfile gets the value of DeviceProfile for the instance
func (instance *MSFT_DOCurrentStatus) GetPropertyDeviceProfile()(value uint32, err error) { 
    retValue, err := instance.GetProperty("DeviceProfile")
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

// SetDiskAvailableBytes sets the value of DiskAvailableBytes for the instance
func (instance *MSFT_DOCurrentStatus) SetPropertyDiskAvailableBytes(value uint64) (err error) { 
    return instance.SetProperty("DiskAvailableBytes", (value))
}

// GetDiskAvailableBytes gets the value of DiskAvailableBytes for the instance
func (instance *MSFT_DOCurrentStatus) GetPropertyDiskAvailableBytes()(value uint64, err error) { 
    retValue, err := instance.GetProperty("DiskAvailableBytes")
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

// SetDiskTotalBytes sets the value of DiskTotalBytes for the instance
func (instance *MSFT_DOCurrentStatus) SetPropertyDiskTotalBytes(value uint64) (err error) { 
    return instance.SetProperty("DiskTotalBytes", (value))
}

// GetDiskTotalBytes gets the value of DiskTotalBytes for the instance
func (instance *MSFT_DOCurrentStatus) GetPropertyDiskTotalBytes()(value uint64, err error) { 
    retValue, err := instance.GetProperty("DiskTotalBytes")
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

// SetGroupConnections sets the value of GroupConnections for the instance
func (instance *MSFT_DOCurrentStatus) SetPropertyGroupConnections(value uint32) (err error) { 
    return instance.SetProperty("GroupConnections", (value))
}

// GetGroupConnections gets the value of GroupConnections for the instance
func (instance *MSFT_DOCurrentStatus) GetPropertyGroupConnections()(value uint32, err error) { 
    retValue, err := instance.GetProperty("GroupConnections")
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

// SetInternetConnections sets the value of InternetConnections for the instance
func (instance *MSFT_DOCurrentStatus) SetPropertyInternetConnections(value uint32) (err error) { 
    return instance.SetProperty("InternetConnections", (value))
}

// GetInternetConnections gets the value of InternetConnections for the instance
func (instance *MSFT_DOCurrentStatus) GetPropertyInternetConnections()(value uint32, err error) { 
    retValue, err := instance.GetProperty("InternetConnections")
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

// SetLanConnections sets the value of LanConnections for the instance
func (instance *MSFT_DOCurrentStatus) SetPropertyLanConnections(value uint32) (err error) { 
    return instance.SetProperty("LanConnections", (value))
}

// GetLanConnections gets the value of LanConnections for the instance
func (instance *MSFT_DOCurrentStatus) GetPropertyLanConnections()(value uint32, err error) { 
    retValue, err := instance.GetProperty("LanConnections")
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

// SetLinkLocalConnections sets the value of LinkLocalConnections for the instance
func (instance *MSFT_DOCurrentStatus) SetPropertyLinkLocalConnections(value uint32) (err error) { 
    return instance.SetProperty("LinkLocalConnections", (value))
}

// GetLinkLocalConnections gets the value of LinkLocalConnections for the instance
func (instance *MSFT_DOCurrentStatus) GetPropertyLinkLocalConnections()(value uint32, err error) { 
    retValue, err := instance.GetProperty("LinkLocalConnections")
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

// SetMemUsageKBytes sets the value of MemUsageKBytes for the instance
func (instance *MSFT_DOCurrentStatus) SetPropertyMemUsageKBytes(value uint64) (err error) { 
    return instance.SetProperty("MemUsageKBytes", (value))
}

// GetMemUsageKBytes gets the value of MemUsageKBytes for the instance
func (instance *MSFT_DOCurrentStatus) GetPropertyMemUsageKBytes()(value uint64, err error) { 
    retValue, err := instance.GetProperty("MemUsageKBytes")
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

// SetPeerInfoCount sets the value of PeerInfoCount for the instance
func (instance *MSFT_DOCurrentStatus) SetPropertyPeerInfoCount(value uint32) (err error) { 
    return instance.SetProperty("PeerInfoCount", (value))
}

// GetPeerInfoCount gets the value of PeerInfoCount for the instance
func (instance *MSFT_DOCurrentStatus) GetPropertyPeerInfoCount()(value uint32, err error) { 
    retValue, err := instance.GetProperty("PeerInfoCount")
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

// SetSwarms sets the value of Swarms for the instance
func (instance *MSFT_DOCurrentStatus) SetPropertySwarms(value uint32) (err error) { 
    return instance.SetProperty("Swarms", (value))
}

// GetSwarms gets the value of Swarms for the instance
func (instance *MSFT_DOCurrentStatus) GetPropertySwarms()(value uint32, err error) { 
    retValue, err := instance.GetProperty("Swarms")
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

