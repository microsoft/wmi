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

// MSFC_DH_Chap_Parameters struct
type MSFC_DH_Chap_Parameters struct { 
	*cim.WmiInstance

	// 
	SecretEncoding Parameters_SecretEncoding

	// 
	SharedSecret []uint8

	// 
	SharedSecretLength uint32
}

	func NewMSFC_DH_Chap_ParametersEx1(instance *cim.WmiInstance) (newInstance *MSFC_DH_Chap_Parameters, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &MSFC_DH_Chap_Parameters {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewMSFC_DH_Chap_ParametersEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MSFC_DH_Chap_Parameters, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MSFC_DH_Chap_Parameters {
	WmiInstance: tmp,
	}
	return
	}
	

// SetSecretEncoding sets the value of SecretEncoding for the instance
func (instance *MSFC_DH_Chap_Parameters) SetPropertySecretEncoding(value Parameters_SecretEncoding) (err error) { 
    return instance.SetProperty("SecretEncoding", (value))
}

// GetSecretEncoding gets the value of SecretEncoding for the instance
func (instance *MSFC_DH_Chap_Parameters) GetPropertySecretEncoding()(value Parameters_SecretEncoding, err error) { 
    retValue, err := instance.GetProperty("SecretEncoding")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(int32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = Parameters_SecretEncoding(valuetmp)

    return
}

// SetSharedSecret sets the value of SharedSecret for the instance
func (instance *MSFC_DH_Chap_Parameters) SetPropertySharedSecret(value []uint8) (err error) { 
    return instance.SetProperty("SharedSecret", (value))
}

// GetSharedSecret gets the value of SharedSecret for the instance
func (instance *MSFC_DH_Chap_Parameters) GetPropertySharedSecret()(value []uint8, err error) { 
    retValue, err := instance.GetProperty("SharedSecret")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    for _, interfaceValue := range retValue.([]interface{}) {
        valuetmp, ok := interfaceValue.(uint8); 
        if !ok {
            err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
            return  
        }
        value = append(value, uint8(valuetmp))
    }

    return
}

// SetSharedSecretLength sets the value of SharedSecretLength for the instance
func (instance *MSFC_DH_Chap_Parameters) SetPropertySharedSecretLength(value uint32) (err error) { 
    return instance.SetProperty("SharedSecretLength", (value))
}

// GetSharedSecretLength gets the value of SharedSecretLength for the instance
func (instance *MSFC_DH_Chap_Parameters) GetPropertySharedSecretLength()(value uint32, err error) { 
    retValue, err := instance.GetProperty("SharedSecretLength")
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

