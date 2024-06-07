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
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// MSDiskDriver_Geometry struct
type MSDiskDriver_Geometry struct { 
	*MSDiskDriver

	// 
	Active bool

	// 
	BytesPerSector uint32

	// 
	Cylinders int64

	// 
	InstanceName string

	// 
	MediaType uint32

	// 
	SectorsPerTrack uint32

	// 
	TracksPerCylinder uint32
}

	func NewMSDiskDriver_GeometryEx1(instance *cim.WmiInstance) (newInstance *MSDiskDriver_Geometry, err error) {tmp, err := NewMSDiskDriverEx1(instance)
		
	if err != nil { return }
	newInstance = &MSDiskDriver_Geometry {
	MSDiskDriver: tmp,
	}
	return
	}
	

	func NewMSDiskDriver_GeometryEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MSDiskDriver_Geometry, err error) {tmp, err := NewMSDiskDriverEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MSDiskDriver_Geometry {
	MSDiskDriver: tmp,
	}
	return
	}
	

// SetActive sets the value of Active for the instance
func (instance *MSDiskDriver_Geometry) SetPropertyActive(value bool) (err error) { 
    return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSDiskDriver_Geometry) GetPropertyActive()(value bool, err error) { 
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

// SetBytesPerSector sets the value of BytesPerSector for the instance
func (instance *MSDiskDriver_Geometry) SetPropertyBytesPerSector(value uint32) (err error) { 
    return instance.SetProperty("BytesPerSector", (value))
}

// GetBytesPerSector gets the value of BytesPerSector for the instance
func (instance *MSDiskDriver_Geometry) GetPropertyBytesPerSector()(value uint32, err error) { 
    retValue, err := instance.GetProperty("BytesPerSector")
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

// SetCylinders sets the value of Cylinders for the instance
func (instance *MSDiskDriver_Geometry) SetPropertyCylinders(value int64) (err error) { 
    return instance.SetProperty("Cylinders", (value))
}

// GetCylinders gets the value of Cylinders for the instance
func (instance *MSDiskDriver_Geometry) GetPropertyCylinders()(value int64, err error) { 
    retValue, err := instance.GetProperty("Cylinders")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(int64); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " int64 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = int64(valuetmp)

    return
}

// SetInstanceName sets the value of InstanceName for the instance
func (instance *MSDiskDriver_Geometry) SetPropertyInstanceName(value string) (err error) { 
    return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSDiskDriver_Geometry) GetPropertyInstanceName()(value string, err error) { 
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

// SetMediaType sets the value of MediaType for the instance
func (instance *MSDiskDriver_Geometry) SetPropertyMediaType(value uint32) (err error) { 
    return instance.SetProperty("MediaType", (value))
}

// GetMediaType gets the value of MediaType for the instance
func (instance *MSDiskDriver_Geometry) GetPropertyMediaType()(value uint32, err error) { 
    retValue, err := instance.GetProperty("MediaType")
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

// SetSectorsPerTrack sets the value of SectorsPerTrack for the instance
func (instance *MSDiskDriver_Geometry) SetPropertySectorsPerTrack(value uint32) (err error) { 
    return instance.SetProperty("SectorsPerTrack", (value))
}

// GetSectorsPerTrack gets the value of SectorsPerTrack for the instance
func (instance *MSDiskDriver_Geometry) GetPropertySectorsPerTrack()(value uint32, err error) { 
    retValue, err := instance.GetProperty("SectorsPerTrack")
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

// SetTracksPerCylinder sets the value of TracksPerCylinder for the instance
func (instance *MSDiskDriver_Geometry) SetPropertyTracksPerCylinder(value uint32) (err error) { 
    return instance.SetProperty("TracksPerCylinder", (value))
}

// GetTracksPerCylinder gets the value of TracksPerCylinder for the instance
func (instance *MSDiskDriver_Geometry) GetPropertyTracksPerCylinder()(value uint32, err error) { 
    retValue, err := instance.GetProperty("TracksPerCylinder")
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

