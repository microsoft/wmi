// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.Microsoft.Windows.ManagementTools
//////////////////////////////////////////////
package managementtools
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// MSFT_MTDiskVolume struct
type MSFT_MTDiskVolume struct { 
	*CIM_ManagedElement

	// 
	FormattedSize uint64

	// 
	PageFile bool

	// 
	SystemDisk bool

	// 
	VolumePath string
}

	func NewMSFT_MTDiskVolumeEx1(instance *cim.WmiInstance) (newInstance *MSFT_MTDiskVolume, err error) {tmp, err := NewCIM_ManagedElementEx1(instance)
		
	if err != nil { return }
	newInstance = &MSFT_MTDiskVolume {
	CIM_ManagedElement: tmp,
	}
	return
	}
	

	func NewMSFT_MTDiskVolumeEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MSFT_MTDiskVolume, err error) {tmp, err := NewCIM_ManagedElementEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MSFT_MTDiskVolume {
	CIM_ManagedElement: tmp,
	}
	return
	}
	

// SetFormattedSize sets the value of FormattedSize for the instance
func (instance *MSFT_MTDiskVolume) SetPropertyFormattedSize(value uint64) (err error) { 
    return instance.SetProperty("FormattedSize", (value))
}

// GetFormattedSize gets the value of FormattedSize for the instance
func (instance *MSFT_MTDiskVolume) GetPropertyFormattedSize()(value uint64, err error) { 
    retValue, err := instance.GetProperty("FormattedSize")
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

// SetPageFile sets the value of PageFile for the instance
func (instance *MSFT_MTDiskVolume) SetPropertyPageFile(value bool) (err error) { 
    return instance.SetProperty("PageFile", (value))
}

// GetPageFile gets the value of PageFile for the instance
func (instance *MSFT_MTDiskVolume) GetPropertyPageFile()(value bool, err error) { 
    retValue, err := instance.GetProperty("PageFile")
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

// SetSystemDisk sets the value of SystemDisk for the instance
func (instance *MSFT_MTDiskVolume) SetPropertySystemDisk(value bool) (err error) { 
    return instance.SetProperty("SystemDisk", (value))
}

// GetSystemDisk gets the value of SystemDisk for the instance
func (instance *MSFT_MTDiskVolume) GetPropertySystemDisk()(value bool, err error) { 
    retValue, err := instance.GetProperty("SystemDisk")
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

// SetVolumePath sets the value of VolumePath for the instance
func (instance *MSFT_MTDiskVolume) SetPropertyVolumePath(value string) (err error) { 
    return instance.SetProperty("VolumePath", (value))
}

// GetVolumePath gets the value of VolumePath for the instance
func (instance *MSFT_MTDiskVolume) GetPropertyVolumePath()(value string, err error) { 
    retValue, err := instance.GetProperty("VolumePath")
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

