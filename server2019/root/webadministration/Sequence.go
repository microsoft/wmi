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

// Sequence struct
type Sequence struct { 
	*CollectionElement

	// 
	Sequence string
}

	func NewSequenceEx1(instance *cim.WmiInstance) (newInstance *Sequence, err error) {tmp, err := NewCollectionElementEx1(instance)
		
	if err != nil { return }
	newInstance = &Sequence {
	CollectionElement: tmp,
	}
	return
	}
	

	func NewSequenceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *Sequence, err error) {tmp, err := NewCollectionElementEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &Sequence {
	CollectionElement: tmp,
	}
	return
	}
	

// SetSequence sets the value of Sequence for the instance
func (instance *Sequence) SetPropertySequence(value string) (err error) { 
    return instance.SetProperty("Sequence", (value))
}

// GetSequence gets the value of Sequence for the instance
func (instance *Sequence) GetPropertySequence()(value string, err error) { 
    retValue, err := instance.GetProperty("Sequence")
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

