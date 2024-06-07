// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.Microsoft.Windows.SMB
//////////////////////////////////////////////
package smb
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/base/instance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// MSFT_SmbClientCertificateMapping struct
type MSFT_SmbClientCertificateMapping struct { 
	*cim.WmiInstance

	// 
	DisplayName string

	// 
	Flags SmbClientCertificateMapping_Flags

	// 
	IssuerName string

	// 
	MappingStatus uint32

	// 
	Namespace string

	// 
	RenewalChain string

	// 
	StoreName string

	// 
	Subject string

	// 
	Thumbprint string

	// 
	Type SmbClientCertificateMapping_Type
}

	func NewMSFT_SmbClientCertificateMappingEx1(instance *cim.WmiInstance) (newInstance *MSFT_SmbClientCertificateMapping, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &MSFT_SmbClientCertificateMapping {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewMSFT_SmbClientCertificateMappingEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MSFT_SmbClientCertificateMapping, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MSFT_SmbClientCertificateMapping {
	WmiInstance: tmp,
	}
	return
	}
	

// SetDisplayName sets the value of DisplayName for the instance
func (instance *MSFT_SmbClientCertificateMapping) SetPropertyDisplayName(value string) (err error) { 
    return instance.SetProperty("DisplayName", (value))
}

// GetDisplayName gets the value of DisplayName for the instance
func (instance *MSFT_SmbClientCertificateMapping) GetPropertyDisplayName()(value string, err error) { 
    retValue, err := instance.GetProperty("DisplayName")
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

// SetFlags sets the value of Flags for the instance
func (instance *MSFT_SmbClientCertificateMapping) SetPropertyFlags(value SmbClientCertificateMapping_Flags) (err error) { 
    return instance.SetProperty("Flags", (value))
}

// GetFlags gets the value of Flags for the instance
func (instance *MSFT_SmbClientCertificateMapping) GetPropertyFlags()(value SmbClientCertificateMapping_Flags, err error) { 
    retValue, err := instance.GetProperty("Flags")
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

    value = SmbClientCertificateMapping_Flags(valuetmp)

    return
}

// SetIssuerName sets the value of IssuerName for the instance
func (instance *MSFT_SmbClientCertificateMapping) SetPropertyIssuerName(value string) (err error) { 
    return instance.SetProperty("IssuerName", (value))
}

// GetIssuerName gets the value of IssuerName for the instance
func (instance *MSFT_SmbClientCertificateMapping) GetPropertyIssuerName()(value string, err error) { 
    retValue, err := instance.GetProperty("IssuerName")
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

// SetMappingStatus sets the value of MappingStatus for the instance
func (instance *MSFT_SmbClientCertificateMapping) SetPropertyMappingStatus(value uint32) (err error) { 
    return instance.SetProperty("MappingStatus", (value))
}

// GetMappingStatus gets the value of MappingStatus for the instance
func (instance *MSFT_SmbClientCertificateMapping) GetPropertyMappingStatus()(value uint32, err error) { 
    retValue, err := instance.GetProperty("MappingStatus")
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

// SetNamespace sets the value of Namespace for the instance
func (instance *MSFT_SmbClientCertificateMapping) SetPropertyNamespace(value string) (err error) { 
    return instance.SetProperty("Namespace", (value))
}

// GetNamespace gets the value of Namespace for the instance
func (instance *MSFT_SmbClientCertificateMapping) GetPropertyNamespace()(value string, err error) { 
    retValue, err := instance.GetProperty("Namespace")
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

// SetRenewalChain sets the value of RenewalChain for the instance
func (instance *MSFT_SmbClientCertificateMapping) SetPropertyRenewalChain(value string) (err error) { 
    return instance.SetProperty("RenewalChain", (value))
}

// GetRenewalChain gets the value of RenewalChain for the instance
func (instance *MSFT_SmbClientCertificateMapping) GetPropertyRenewalChain()(value string, err error) { 
    retValue, err := instance.GetProperty("RenewalChain")
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

// SetStoreName sets the value of StoreName for the instance
func (instance *MSFT_SmbClientCertificateMapping) SetPropertyStoreName(value string) (err error) { 
    return instance.SetProperty("StoreName", (value))
}

// GetStoreName gets the value of StoreName for the instance
func (instance *MSFT_SmbClientCertificateMapping) GetPropertyStoreName()(value string, err error) { 
    retValue, err := instance.GetProperty("StoreName")
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

// SetSubject sets the value of Subject for the instance
func (instance *MSFT_SmbClientCertificateMapping) SetPropertySubject(value string) (err error) { 
    return instance.SetProperty("Subject", (value))
}

// GetSubject gets the value of Subject for the instance
func (instance *MSFT_SmbClientCertificateMapping) GetPropertySubject()(value string, err error) { 
    retValue, err := instance.GetProperty("Subject")
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

// SetThumbprint sets the value of Thumbprint for the instance
func (instance *MSFT_SmbClientCertificateMapping) SetPropertyThumbprint(value string) (err error) { 
    return instance.SetProperty("Thumbprint", (value))
}

// GetThumbprint gets the value of Thumbprint for the instance
func (instance *MSFT_SmbClientCertificateMapping) GetPropertyThumbprint()(value string, err error) { 
    retValue, err := instance.GetProperty("Thumbprint")
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

// SetType sets the value of Type for the instance
func (instance *MSFT_SmbClientCertificateMapping) SetPropertyType(value SmbClientCertificateMapping_Type) (err error) { 
    return instance.SetProperty("Type", (value))
}

// GetType gets the value of Type for the instance
func (instance *MSFT_SmbClientCertificateMapping) GetPropertyType()(value SmbClientCertificateMapping_Type, err error) { 
    retValue, err := instance.GetProperty("Type")
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

    value = SmbClientCertificateMapping_Type(valuetmp)

    return
}

// 

// <param name="DisplayName" type="string "></param>
// <param name="Flags" type="uint32 "></param>
// <param name="IssuerName" type="string "></param>
// <param name="Namespace" type="string "></param>
// <param name="StoreName" type="string "></param>
// <param name="Subject" type="string "></param>
// <param name="Thumbprint" type="string "></param>
// <param name="Type" type="uint32 "></param>

// <param name="CreatedMapping" type="MSFT_SmbClientCertificateMapping "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_SmbClientCertificateMapping) AddClientCertificateMapping( /* IN */ Namespace string,
 /* IN */ IssuerName string,
 /* IN */ Subject string,
 /* IN */ Thumbprint string,
 /* IN */ DisplayName string,
 /* IN */ StoreName string,
 /* IN */ Type uint32,
 /* IN */ Flags uint32,
 /* OUT */ CreatedMapping MSFT_SmbClientCertificateMapping) (result uint32, err error) {retVal, err := instance.InvokeMethod("AddClientCertificateMapping" , Namespace, IssuerName, Subject, Thumbprint, DisplayName, StoreName, Type, Flags)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// 

// <param name="Flags" type="uint32 "></param>
// <param name="IssuerName" type="string "></param>
// <param name="Namespace" type="string "></param>
// <param name="StoreName" type="string "></param>
// <param name="Thumbprint" type="string "></param>

// <param name="CreatedMapping" type="MSFT_SmbClientCertificateMapping "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_SmbClientCertificateMapping) SetClientCertificateMapping( /* IN */ Namespace string,
 /* IN */ IssuerName string,
 /* IN */ Thumbprint string,
 /* IN */ StoreName string,
 /* IN */ Flags uint32,
 /* OUT */ CreatedMapping MSFT_SmbClientCertificateMapping) (result uint32, err error) {retVal, err := instance.InvokeMethod("SetClientCertificateMapping" , Namespace, IssuerName, Thumbprint, StoreName, Flags)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


