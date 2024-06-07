// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/base/instance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// MSFT_NetAdapterUsoEncapsulationTypes struct
type MSFT_NetAdapterUsoEncapsulationTypes struct { 
	*cim.WmiInstance

	// 
	NdisEncapsulationIeee802_3 bool

	// 
	NdisEncapsulationIeee802_3pAndq bool

	// 
	NdisEncapsulationIeee802_3PAndQInOob bool

	// 
	NdisEncapsulationIeeLlcSnapRouted bool

	// 
	NdisEncapsulationNotNull bool

	// 
	NdisEncapsulationNotSupported bool
}

	func NewMSFT_NetAdapterUsoEncapsulationTypesEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetAdapterUsoEncapsulationTypes, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &MSFT_NetAdapterUsoEncapsulationTypes {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewMSFT_NetAdapterUsoEncapsulationTypesEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MSFT_NetAdapterUsoEncapsulationTypes, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MSFT_NetAdapterUsoEncapsulationTypes {
	WmiInstance: tmp,
	}
	return
	}
	

// SetNdisEncapsulationIeee802_3 sets the value of NdisEncapsulationIeee802_3 for the instance
func (instance *MSFT_NetAdapterUsoEncapsulationTypes) SetPropertyNdisEncapsulationIeee802_3(value bool) (err error) { 
    return instance.SetProperty("NdisEncapsulationIeee802_3", (value))
}

// GetNdisEncapsulationIeee802_3 gets the value of NdisEncapsulationIeee802_3 for the instance
func (instance *MSFT_NetAdapterUsoEncapsulationTypes) GetPropertyNdisEncapsulationIeee802_3()(value bool, err error) { 
    retValue, err := instance.GetProperty("NdisEncapsulationIeee802_3")
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

// SetNdisEncapsulationIeee802_3pAndq sets the value of NdisEncapsulationIeee802_3pAndq for the instance
func (instance *MSFT_NetAdapterUsoEncapsulationTypes) SetPropertyNdisEncapsulationIeee802_3pAndq(value bool) (err error) { 
    return instance.SetProperty("NdisEncapsulationIeee802_3pAndq", (value))
}

// GetNdisEncapsulationIeee802_3pAndq gets the value of NdisEncapsulationIeee802_3pAndq for the instance
func (instance *MSFT_NetAdapterUsoEncapsulationTypes) GetPropertyNdisEncapsulationIeee802_3pAndq()(value bool, err error) { 
    retValue, err := instance.GetProperty("NdisEncapsulationIeee802_3pAndq")
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

// SetNdisEncapsulationIeee802_3PAndQInOob sets the value of NdisEncapsulationIeee802_3PAndQInOob for the instance
func (instance *MSFT_NetAdapterUsoEncapsulationTypes) SetPropertyNdisEncapsulationIeee802_3PAndQInOob(value bool) (err error) { 
    return instance.SetProperty("NdisEncapsulationIeee802_3PAndQInOob", (value))
}

// GetNdisEncapsulationIeee802_3PAndQInOob gets the value of NdisEncapsulationIeee802_3PAndQInOob for the instance
func (instance *MSFT_NetAdapterUsoEncapsulationTypes) GetPropertyNdisEncapsulationIeee802_3PAndQInOob()(value bool, err error) { 
    retValue, err := instance.GetProperty("NdisEncapsulationIeee802_3PAndQInOob")
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

// SetNdisEncapsulationIeeLlcSnapRouted sets the value of NdisEncapsulationIeeLlcSnapRouted for the instance
func (instance *MSFT_NetAdapterUsoEncapsulationTypes) SetPropertyNdisEncapsulationIeeLlcSnapRouted(value bool) (err error) { 
    return instance.SetProperty("NdisEncapsulationIeeLlcSnapRouted", (value))
}

// GetNdisEncapsulationIeeLlcSnapRouted gets the value of NdisEncapsulationIeeLlcSnapRouted for the instance
func (instance *MSFT_NetAdapterUsoEncapsulationTypes) GetPropertyNdisEncapsulationIeeLlcSnapRouted()(value bool, err error) { 
    retValue, err := instance.GetProperty("NdisEncapsulationIeeLlcSnapRouted")
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

// SetNdisEncapsulationNotNull sets the value of NdisEncapsulationNotNull for the instance
func (instance *MSFT_NetAdapterUsoEncapsulationTypes) SetPropertyNdisEncapsulationNotNull(value bool) (err error) { 
    return instance.SetProperty("NdisEncapsulationNotNull", (value))
}

// GetNdisEncapsulationNotNull gets the value of NdisEncapsulationNotNull for the instance
func (instance *MSFT_NetAdapterUsoEncapsulationTypes) GetPropertyNdisEncapsulationNotNull()(value bool, err error) { 
    retValue, err := instance.GetProperty("NdisEncapsulationNotNull")
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

// SetNdisEncapsulationNotSupported sets the value of NdisEncapsulationNotSupported for the instance
func (instance *MSFT_NetAdapterUsoEncapsulationTypes) SetPropertyNdisEncapsulationNotSupported(value bool) (err error) { 
    return instance.SetProperty("NdisEncapsulationNotSupported", (value))
}

// GetNdisEncapsulationNotSupported gets the value of NdisEncapsulationNotSupported for the instance
func (instance *MSFT_NetAdapterUsoEncapsulationTypes) GetPropertyNdisEncapsulationNotSupported()(value bool, err error) { 
    retValue, err := instance.GetProperty("NdisEncapsulationNotSupported")
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

