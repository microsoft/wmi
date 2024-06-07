// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.ServiceModel
//////////////////////////////////////////////
package servicemodel
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// ReliableSessionBindingElement struct
type ReliableSessionBindingElement struct { 
	*BindingElement

	// The interval of time that a destination waits before sending an acknowledgement to the message source on reliable channels that are created by the factory.
	AcknowledgementInterval string

	// A boolean value that specifies if flow control is enabled.
	FlowControlEnabled bool

	// Specifies the maximum duration the channel is going to allow the other communicating party not to send any messages before faulting the channel.
	InactivityTimeout string

	// The maximum number of channels that can wait to be accepted on the listener.
	MaxPendingChannels int32

	// The maximum number of times a reliable channel attempts to retransmit a message it has not received an acknowledgement for, by calling Send on its underlying channel.
	MaxRetryCount int32

	// The maximum transfer window size for the reliable session.
	MaxTransferWindowSize int32

	// A Boolean value that specifies whether messages are guaranteed to arrive in the order they were sent.
	Ordered bool

	// The WS-ReliableMessaging protocol version used in the reliable session.
	ReliableMessagingVersion string
}

	func NewReliableSessionBindingElementEx1(instance *cim.WmiInstance) (newInstance *ReliableSessionBindingElement, err error) {tmp, err := NewBindingElementEx1(instance)
		
	if err != nil { return }
	newInstance = &ReliableSessionBindingElement {
	BindingElement: tmp,
	}
	return
	}
	

	func NewReliableSessionBindingElementEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *ReliableSessionBindingElement, err error) {tmp, err := NewBindingElementEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &ReliableSessionBindingElement {
	BindingElement: tmp,
	}
	return
	}
	

// SetAcknowledgementInterval sets the value of AcknowledgementInterval for the instance
func (instance *ReliableSessionBindingElement) SetPropertyAcknowledgementInterval(value string) (err error) { 
    return instance.SetProperty("AcknowledgementInterval", (value))
}

// GetAcknowledgementInterval gets the value of AcknowledgementInterval for the instance
func (instance *ReliableSessionBindingElement) GetPropertyAcknowledgementInterval()(value string, err error) { 
    retValue, err := instance.GetProperty("AcknowledgementInterval")
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

// SetFlowControlEnabled sets the value of FlowControlEnabled for the instance
func (instance *ReliableSessionBindingElement) SetPropertyFlowControlEnabled(value bool) (err error) { 
    return instance.SetProperty("FlowControlEnabled", (value))
}

// GetFlowControlEnabled gets the value of FlowControlEnabled for the instance
func (instance *ReliableSessionBindingElement) GetPropertyFlowControlEnabled()(value bool, err error) { 
    retValue, err := instance.GetProperty("FlowControlEnabled")
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

// SetInactivityTimeout sets the value of InactivityTimeout for the instance
func (instance *ReliableSessionBindingElement) SetPropertyInactivityTimeout(value string) (err error) { 
    return instance.SetProperty("InactivityTimeout", (value))
}

// GetInactivityTimeout gets the value of InactivityTimeout for the instance
func (instance *ReliableSessionBindingElement) GetPropertyInactivityTimeout()(value string, err error) { 
    retValue, err := instance.GetProperty("InactivityTimeout")
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

// SetMaxPendingChannels sets the value of MaxPendingChannels for the instance
func (instance *ReliableSessionBindingElement) SetPropertyMaxPendingChannels(value int32) (err error) { 
    return instance.SetProperty("MaxPendingChannels", (value))
}

// GetMaxPendingChannels gets the value of MaxPendingChannels for the instance
func (instance *ReliableSessionBindingElement) GetPropertyMaxPendingChannels()(value int32, err error) { 
    retValue, err := instance.GetProperty("MaxPendingChannels")
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

// SetMaxRetryCount sets the value of MaxRetryCount for the instance
func (instance *ReliableSessionBindingElement) SetPropertyMaxRetryCount(value int32) (err error) { 
    return instance.SetProperty("MaxRetryCount", (value))
}

// GetMaxRetryCount gets the value of MaxRetryCount for the instance
func (instance *ReliableSessionBindingElement) GetPropertyMaxRetryCount()(value int32, err error) { 
    retValue, err := instance.GetProperty("MaxRetryCount")
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

// SetMaxTransferWindowSize sets the value of MaxTransferWindowSize for the instance
func (instance *ReliableSessionBindingElement) SetPropertyMaxTransferWindowSize(value int32) (err error) { 
    return instance.SetProperty("MaxTransferWindowSize", (value))
}

// GetMaxTransferWindowSize gets the value of MaxTransferWindowSize for the instance
func (instance *ReliableSessionBindingElement) GetPropertyMaxTransferWindowSize()(value int32, err error) { 
    retValue, err := instance.GetProperty("MaxTransferWindowSize")
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

// SetOrdered sets the value of Ordered for the instance
func (instance *ReliableSessionBindingElement) SetPropertyOrdered(value bool) (err error) { 
    return instance.SetProperty("Ordered", (value))
}

// GetOrdered gets the value of Ordered for the instance
func (instance *ReliableSessionBindingElement) GetPropertyOrdered()(value bool, err error) { 
    retValue, err := instance.GetProperty("Ordered")
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

// SetReliableMessagingVersion sets the value of ReliableMessagingVersion for the instance
func (instance *ReliableSessionBindingElement) SetPropertyReliableMessagingVersion(value string) (err error) { 
    return instance.SetProperty("ReliableMessagingVersion", (value))
}

// GetReliableMessagingVersion gets the value of ReliableMessagingVersion for the instance
func (instance *ReliableSessionBindingElement) GetPropertyReliableMessagingVersion()(value string, err error) { 
    retValue, err := instance.GetProperty("ReliableMessagingVersion")
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

