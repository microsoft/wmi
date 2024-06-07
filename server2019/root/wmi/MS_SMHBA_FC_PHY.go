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
 "github.com/microsoft/wmi/pkg/base/instance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// MS_SMHBA_FC_PHY struct
type MS_SMHBA_FC_PHY struct { 
	*cim.WmiInstance

	// 
	MaxFrameSize uint32

	// 
	PhySpeed uint32

	// 
	PhySupportSpeed uint32

	// 
	PhyType uint8
}

	func NewMS_SMHBA_FC_PHYEx1(instance *cim.WmiInstance) (newInstance *MS_SMHBA_FC_PHY, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &MS_SMHBA_FC_PHY {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewMS_SMHBA_FC_PHYEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MS_SMHBA_FC_PHY, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MS_SMHBA_FC_PHY {
	WmiInstance: tmp,
	}
	return
	}
	

// SetMaxFrameSize sets the value of MaxFrameSize for the instance
func (instance *MS_SMHBA_FC_PHY) SetPropertyMaxFrameSize(value uint32) (err error) { 
    return instance.SetProperty("MaxFrameSize", (value))
}

// GetMaxFrameSize gets the value of MaxFrameSize for the instance
func (instance *MS_SMHBA_FC_PHY) GetPropertyMaxFrameSize()(value uint32, err error) { 
    retValue, err := instance.GetProperty("MaxFrameSize")
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

// SetPhySpeed sets the value of PhySpeed for the instance
func (instance *MS_SMHBA_FC_PHY) SetPropertyPhySpeed(value uint32) (err error) { 
    return instance.SetProperty("PhySpeed", (value))
}

// GetPhySpeed gets the value of PhySpeed for the instance
func (instance *MS_SMHBA_FC_PHY) GetPropertyPhySpeed()(value uint32, err error) { 
    retValue, err := instance.GetProperty("PhySpeed")
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

// SetPhySupportSpeed sets the value of PhySupportSpeed for the instance
func (instance *MS_SMHBA_FC_PHY) SetPropertyPhySupportSpeed(value uint32) (err error) { 
    return instance.SetProperty("PhySupportSpeed", (value))
}

// GetPhySupportSpeed gets the value of PhySupportSpeed for the instance
func (instance *MS_SMHBA_FC_PHY) GetPropertyPhySupportSpeed()(value uint32, err error) { 
    retValue, err := instance.GetProperty("PhySupportSpeed")
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

// SetPhyType sets the value of PhyType for the instance
func (instance *MS_SMHBA_FC_PHY) SetPropertyPhyType(value uint8) (err error) { 
    return instance.SetProperty("PhyType", (value))
}

// GetPhyType gets the value of PhyType for the instance
func (instance *MS_SMHBA_FC_PHY) GetPropertyPhyType()(value uint8, err error) { 
    retValue, err := instance.GetProperty("PhyType")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint8); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint8(valuetmp)

    return
}

