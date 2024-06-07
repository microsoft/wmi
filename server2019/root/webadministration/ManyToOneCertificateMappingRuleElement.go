// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.WebAdministration
//////////////////////////////////////////////
package webadministration
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// ManyToOneCertificateMappingRuleElement struct
type ManyToOneCertificateMappingRuleElement struct { 
	*CollectionElement

	// 
	CertificateField int32

	// 
	CertificateSubField string

	// 
	CompareCaseSensitive bool

	// 
	MatchCriteria string
}

	func NewManyToOneCertificateMappingRuleElementEx1(instance *cim.WmiInstance) (newInstance *ManyToOneCertificateMappingRuleElement, err error) {tmp, err := NewCollectionElementEx1(instance)
		
	if err != nil { return }
	newInstance = &ManyToOneCertificateMappingRuleElement {
	CollectionElement: tmp,
	}
	return
	}
	

	func NewManyToOneCertificateMappingRuleElementEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *ManyToOneCertificateMappingRuleElement, err error) {tmp, err := NewCollectionElementEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &ManyToOneCertificateMappingRuleElement {
	CollectionElement: tmp,
	}
	return
	}
	

// SetCertificateField sets the value of CertificateField for the instance
func (instance *ManyToOneCertificateMappingRuleElement) SetPropertyCertificateField(value int32) (err error) { 
    return instance.SetProperty("CertificateField", (value))
}

// GetCertificateField gets the value of CertificateField for the instance
func (instance *ManyToOneCertificateMappingRuleElement) GetPropertyCertificateField()(value int32, err error) { 
    retValue, err := instance.GetProperty("CertificateField")
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

    value = int32(valuetmp)

    return
}

// SetCertificateSubField sets the value of CertificateSubField for the instance
func (instance *ManyToOneCertificateMappingRuleElement) SetPropertyCertificateSubField(value string) (err error) { 
    return instance.SetProperty("CertificateSubField", (value))
}

// GetCertificateSubField gets the value of CertificateSubField for the instance
func (instance *ManyToOneCertificateMappingRuleElement) GetPropertyCertificateSubField()(value string, err error) { 
    retValue, err := instance.GetProperty("CertificateSubField")
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

// SetCompareCaseSensitive sets the value of CompareCaseSensitive for the instance
func (instance *ManyToOneCertificateMappingRuleElement) SetPropertyCompareCaseSensitive(value bool) (err error) { 
    return instance.SetProperty("CompareCaseSensitive", (value))
}

// GetCompareCaseSensitive gets the value of CompareCaseSensitive for the instance
func (instance *ManyToOneCertificateMappingRuleElement) GetPropertyCompareCaseSensitive()(value bool, err error) { 
    retValue, err := instance.GetProperty("CompareCaseSensitive")
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

// SetMatchCriteria sets the value of MatchCriteria for the instance
func (instance *ManyToOneCertificateMappingRuleElement) SetPropertyMatchCriteria(value string) (err error) { 
    return instance.SetProperty("MatchCriteria", (value))
}

// GetMatchCriteria gets the value of MatchCriteria for the instance
func (instance *ManyToOneCertificateMappingRuleElement) GetPropertyMatchCriteria()(value string, err error) { 
    retValue, err := instance.GetProperty("MatchCriteria")
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

