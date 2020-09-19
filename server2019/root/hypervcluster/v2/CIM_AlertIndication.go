// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.HyperVCluster.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// CIM_AlertIndication struct
type CIM_AlertIndication struct {
	*CIM_ProcessIndication

	// The format of the AlertingManagedElement property is interpretable based upon the value of this property. Values are defined as:
	///0 - Unknown. The format is unknown or not meaningfully interpretable by a CIM client application.
	///1 - Other. The format is defined by the value of the OtherAlertingElementFormat property.
	///2 - CIMObjectPath. The format is a CIMObjectPath, with format <NamespacePath>:<ClassName>.<Prop1>="<Value1>", <Prop2>="<Value2>", . . . specifying an instance in the CIM Schema.
	AlertingElementFormat AlertIndication_AlertingElementFormat

	// The identifying information of the entity (ie, the instance) for which this Indication is generated. The property contains the path of an instance, encoded as a string parameter - if the instance is modeled in the CIM Schema. If not a CIM instance, the property contains some identifying string that names the entity for which the Alert is generated. The path or identifying string is formatted per the AlertingElementFormat property.
	AlertingManagedElement string

	// Primary classification of the Indication. The following values are defined:
	///1 - Other. The Indication's OtherAlertType property conveys its classification. Use of "Other" in an enumeration is a standard CIM convention. It means that the current Indication does not fit into the categories described by this enumeration.
	///2 - Communications Alert. An Indication of this type is principally associated with the procedures and/or processes required to convey information from one point to another.
	///3 - Quality of Service Alert. An Indication of this type is principally associated with a degradation or errors in the performance or function of an entity.
	///4 - Processing Error. An Indication of this type is principally associated with a software or processing fault.
	///5 - Device Alert. An Indication of this type is principally associated with an equipment or hardware fault.
	///6 - Environmental Alert. An Indication of this type is principally associated with a condition relating to an enclosure in which the hardware resides, or other environmental considerations.
	///7 - Model Change. The Indication addresses changes in the Information Model. For example, it may embed a Lifecycle Indication to convey the specific model change being alerted.
	///8 - Security Alert. An Indication of this type is associated with security violations, detection of viruses, and similar issues.
	AlertType AlertIndication_AlertType

	// A short description of the Indication.
	Description string

	// An instrumentation or provider specific value that describes the underlying "real-world" event represented by the Indication. Two Indications with the same, non NULL EventID value are considered, by the creating entity, to represent the same event. The comparison of two EventID values is only defined for Alert Indications with identical, non NULL values of SystemCreateClassName, SystemName and ProviderName.
	EventID string

	// The time and date the underlying event was first detected. If specified, this property MUST be set to NULL if the creating entity is not capable of providing this information. This value is based on the notion of local date and time of the Managed System Element generating the Indication.
	EventTime string

	// The formatted message. This message is constructed by combining some or all of the dynamic elements specified in the MessageArguments property with the static elements uniquely identified by the MessageID in a message registry or other catalog associated with the OwningEntity.
	Message string

	// An array containing the dynamic content of the message.
	MessageArguments []string

	// A string that uniquely identifies, within the scope of the OwningEntity, the format of the Message.
	MessageID string

	// A string defining "Other" values for AlertingElementFormat. This value MUST be set to a non NULL value when AlertingElementFormat is set to a value of 1 ("Other"). For all other values of AlertingElementFormat, the value of this string must be set to NULL.
	OtherAlertingElementFormat string

	// A string describing the Alert type - used when the AlertType property is set to 1, "Other State Change".
	OtherAlertType string

	// A string that uniquely identifies the entity that owns the definition of the format of the Message described in this instance. OwningEntity MUST include a copyrighted, trademarked or otherwise unique name that is owned by the business entity or standards body defining the format.
	OwningEntity string

	// An enumerated value that describes the probable cause of the situation which resulted in the AlertIndication.
	ProbableCause AlertIndication_ProbableCause

	// Provides additional information related to the ProbableCause.
	ProbableCauseDescription string

	// The name of the Provider generating this Indication.
	ProviderName string

	// Free form descriptions of the recommended actions to take to resolve the cause of the notification.
	RecommendedActions []string

	// The scoping System's CreationClassName for the Provider generating this Indication.
	SystemCreationClassName string

	// The scoping System's Name for the Provider generating this Indication.
	SystemName string

	// Provides information on trending - trending up, down or no change.
	Trending AlertIndication_Trending
}

func NewCIM_AlertIndicationEx1(instance *cim.WmiInstance) (newInstance *CIM_AlertIndication, err error) {
	tmp, err := NewCIM_ProcessIndicationEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_AlertIndication{
		CIM_ProcessIndication: tmp,
	}
	return
}

func NewCIM_AlertIndicationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_AlertIndication, err error) {
	tmp, err := NewCIM_ProcessIndicationEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_AlertIndication{
		CIM_ProcessIndication: tmp,
	}
	return
}

// SetAlertingElementFormat sets the value of AlertingElementFormat for the instance
func (instance *CIM_AlertIndication) SetPropertyAlertingElementFormat(value AlertIndication_AlertingElementFormat) (err error) {
	return instance.SetProperty("AlertingElementFormat", (value))
}

// GetAlertingElementFormat gets the value of AlertingElementFormat for the instance
func (instance *CIM_AlertIndication) GetPropertyAlertingElementFormat() (value AlertIndication_AlertingElementFormat, err error) {
	retValue, err := instance.GetProperty("AlertingElementFormat")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = AlertIndication_AlertingElementFormat(valuetmp)

	return
}

// SetAlertingManagedElement sets the value of AlertingManagedElement for the instance
func (instance *CIM_AlertIndication) SetPropertyAlertingManagedElement(value string) (err error) {
	return instance.SetProperty("AlertingManagedElement", (value))
}

// GetAlertingManagedElement gets the value of AlertingManagedElement for the instance
func (instance *CIM_AlertIndication) GetPropertyAlertingManagedElement() (value string, err error) {
	retValue, err := instance.GetProperty("AlertingManagedElement")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetAlertType sets the value of AlertType for the instance
func (instance *CIM_AlertIndication) SetPropertyAlertType(value AlertIndication_AlertType) (err error) {
	return instance.SetProperty("AlertType", (value))
}

// GetAlertType gets the value of AlertType for the instance
func (instance *CIM_AlertIndication) GetPropertyAlertType() (value AlertIndication_AlertType, err error) {
	retValue, err := instance.GetProperty("AlertType")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = AlertIndication_AlertType(valuetmp)

	return
}

// SetDescription sets the value of Description for the instance
func (instance *CIM_AlertIndication) SetPropertyDescription(value string) (err error) {
	return instance.SetProperty("Description", (value))
}

// GetDescription gets the value of Description for the instance
func (instance *CIM_AlertIndication) GetPropertyDescription() (value string, err error) {
	retValue, err := instance.GetProperty("Description")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetEventID sets the value of EventID for the instance
func (instance *CIM_AlertIndication) SetPropertyEventID(value string) (err error) {
	return instance.SetProperty("EventID", (value))
}

// GetEventID gets the value of EventID for the instance
func (instance *CIM_AlertIndication) GetPropertyEventID() (value string, err error) {
	retValue, err := instance.GetProperty("EventID")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetEventTime sets the value of EventTime for the instance
func (instance *CIM_AlertIndication) SetPropertyEventTime(value string) (err error) {
	return instance.SetProperty("EventTime", (value))
}

// GetEventTime gets the value of EventTime for the instance
func (instance *CIM_AlertIndication) GetPropertyEventTime() (value string, err error) {
	retValue, err := instance.GetProperty("EventTime")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetMessage sets the value of Message for the instance
func (instance *CIM_AlertIndication) SetPropertyMessage(value string) (err error) {
	return instance.SetProperty("Message", (value))
}

// GetMessage gets the value of Message for the instance
func (instance *CIM_AlertIndication) GetPropertyMessage() (value string, err error) {
	retValue, err := instance.GetProperty("Message")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetMessageArguments sets the value of MessageArguments for the instance
func (instance *CIM_AlertIndication) SetPropertyMessageArguments(value []string) (err error) {
	return instance.SetProperty("MessageArguments", (value))
}

// GetMessageArguments gets the value of MessageArguments for the instance
func (instance *CIM_AlertIndication) GetPropertyMessageArguments() (value []string, err error) {
	retValue, err := instance.GetProperty("MessageArguments")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(string)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, string(valuetmp))
	}

	return
}

// SetMessageID sets the value of MessageID for the instance
func (instance *CIM_AlertIndication) SetPropertyMessageID(value string) (err error) {
	return instance.SetProperty("MessageID", (value))
}

// GetMessageID gets the value of MessageID for the instance
func (instance *CIM_AlertIndication) GetPropertyMessageID() (value string, err error) {
	retValue, err := instance.GetProperty("MessageID")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetOtherAlertingElementFormat sets the value of OtherAlertingElementFormat for the instance
func (instance *CIM_AlertIndication) SetPropertyOtherAlertingElementFormat(value string) (err error) {
	return instance.SetProperty("OtherAlertingElementFormat", (value))
}

// GetOtherAlertingElementFormat gets the value of OtherAlertingElementFormat for the instance
func (instance *CIM_AlertIndication) GetPropertyOtherAlertingElementFormat() (value string, err error) {
	retValue, err := instance.GetProperty("OtherAlertingElementFormat")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetOtherAlertType sets the value of OtherAlertType for the instance
func (instance *CIM_AlertIndication) SetPropertyOtherAlertType(value string) (err error) {
	return instance.SetProperty("OtherAlertType", (value))
}

// GetOtherAlertType gets the value of OtherAlertType for the instance
func (instance *CIM_AlertIndication) GetPropertyOtherAlertType() (value string, err error) {
	retValue, err := instance.GetProperty("OtherAlertType")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetOwningEntity sets the value of OwningEntity for the instance
func (instance *CIM_AlertIndication) SetPropertyOwningEntity(value string) (err error) {
	return instance.SetProperty("OwningEntity", (value))
}

// GetOwningEntity gets the value of OwningEntity for the instance
func (instance *CIM_AlertIndication) GetPropertyOwningEntity() (value string, err error) {
	retValue, err := instance.GetProperty("OwningEntity")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetProbableCause sets the value of ProbableCause for the instance
func (instance *CIM_AlertIndication) SetPropertyProbableCause(value AlertIndication_ProbableCause) (err error) {
	return instance.SetProperty("ProbableCause", (value))
}

// GetProbableCause gets the value of ProbableCause for the instance
func (instance *CIM_AlertIndication) GetPropertyProbableCause() (value AlertIndication_ProbableCause, err error) {
	retValue, err := instance.GetProperty("ProbableCause")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = AlertIndication_ProbableCause(valuetmp)

	return
}

// SetProbableCauseDescription sets the value of ProbableCauseDescription for the instance
func (instance *CIM_AlertIndication) SetPropertyProbableCauseDescription(value string) (err error) {
	return instance.SetProperty("ProbableCauseDescription", (value))
}

// GetProbableCauseDescription gets the value of ProbableCauseDescription for the instance
func (instance *CIM_AlertIndication) GetPropertyProbableCauseDescription() (value string, err error) {
	retValue, err := instance.GetProperty("ProbableCauseDescription")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetProviderName sets the value of ProviderName for the instance
func (instance *CIM_AlertIndication) SetPropertyProviderName(value string) (err error) {
	return instance.SetProperty("ProviderName", (value))
}

// GetProviderName gets the value of ProviderName for the instance
func (instance *CIM_AlertIndication) GetPropertyProviderName() (value string, err error) {
	retValue, err := instance.GetProperty("ProviderName")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetRecommendedActions sets the value of RecommendedActions for the instance
func (instance *CIM_AlertIndication) SetPropertyRecommendedActions(value []string) (err error) {
	return instance.SetProperty("RecommendedActions", (value))
}

// GetRecommendedActions gets the value of RecommendedActions for the instance
func (instance *CIM_AlertIndication) GetPropertyRecommendedActions() (value []string, err error) {
	retValue, err := instance.GetProperty("RecommendedActions")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(string)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, string(valuetmp))
	}

	return
}

// SetSystemCreationClassName sets the value of SystemCreationClassName for the instance
func (instance *CIM_AlertIndication) SetPropertySystemCreationClassName(value string) (err error) {
	return instance.SetProperty("SystemCreationClassName", (value))
}

// GetSystemCreationClassName gets the value of SystemCreationClassName for the instance
func (instance *CIM_AlertIndication) GetPropertySystemCreationClassName() (value string, err error) {
	retValue, err := instance.GetProperty("SystemCreationClassName")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetSystemName sets the value of SystemName for the instance
func (instance *CIM_AlertIndication) SetPropertySystemName(value string) (err error) {
	return instance.SetProperty("SystemName", (value))
}

// GetSystemName gets the value of SystemName for the instance
func (instance *CIM_AlertIndication) GetPropertySystemName() (value string, err error) {
	retValue, err := instance.GetProperty("SystemName")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetTrending sets the value of Trending for the instance
func (instance *CIM_AlertIndication) SetPropertyTrending(value AlertIndication_Trending) (err error) {
	return instance.SetProperty("Trending", (value))
}

// GetTrending gets the value of Trending for the instance
func (instance *CIM_AlertIndication) GetPropertyTrending() (value AlertIndication_Trending, err error) {
	retValue, err := instance.GetProperty("Trending")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = AlertIndication_Trending(valuetmp)

	return
}
