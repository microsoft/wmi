// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

// This class implement a wrapper of the SWbemObject class (from an instance perspective).
// Documentation: https://docs.microsoft.com/en-us/windows/win32/wmisdk/swbemobject

package cim

import (
	"fmt"

	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"

	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
)

type UserAction uint32

const (
	None    UserAction = 0
	Async   UserAction = 1
	Wait    UserAction = 2
	Cancel  UserAction = 3
	Default UserAction = 4
)

// WmiInstance is a representation of a WMI instance
type WmiInstance struct {
	class       *WmiClass
	session     *WmiSession
	instance    *ole.IDispatch
	instanceVar *ole.VARIANT
}

// WmiInstanceCollection is a slice of WmiInstance
type WmiInstanceCollection []*WmiInstance

// Close all instances in a collection
func (wmic *WmiInstanceCollection) Close() {
	for _, i := range *wmic {
		i.Close()
	}
}

func CreateWmiInstance(instanceVar *ole.VARIANT, session *WmiSession) (*WmiInstance, error) {
	return &WmiInstance{
		instanceVar: instanceVar,
		instance:    instanceVar.ToIDispatch(),
		session:     session,
	}, nil
}

// GetInstance returns the latest Instance
func (c *WmiInstance) GetInstance() (*WmiInstance, error) {
	return c.session.GetInstance(c.InstancePath())
}
func (c *WmiInstance) GetSession() *WmiSession {
	return c.session
}

func (c *WmiInstance) GetIDispatch() *ole.IDispatch {
	return c.instance
}
func (c *WmiInstance) GetRawInstance() *ole.VARIANT {
	return c.instanceVar
}

func (c *WmiInstance) GetSystemProperty(name string) (*WmiProperty, error) {
	// Documentation: https://docs.microsoft.com/en-us/windows/win32/wmisdk/swbemobjectex-systemproperties-
	rawResult, err := oleutil.GetProperty(c.instance, "SystemProperties_")
	if err != nil {
		return nil, err
	}

	// SWbemObjectEx.SystemProperties_ returns
	// an SWbemPropertySet object that contains the collection
	// of sytem properties for the c class
	sWbemObjectExAsIDispatch := rawResult.ToIDispatch()
	defer rawResult.Clear()

	// Get the system property
	sWbemProperty, err := oleutil.CallMethod(sWbemObjectExAsIDispatch, "Item", name)
	if err != nil {
		return nil, err
	}

	property, err := CreateWmiProperty(sWbemProperty, c.session)
	if err != nil {
		return nil, err
	}

	return property, nil
}

// GetProperty gets the property of the instance specified by name and returns in value
func (c *WmiInstance) GetProperty(name string) (interface{}, error) {
	rawResult, err := oleutil.GetProperty(c.instance, name)
	if err != nil {
		return nil, err
	}

	defer rawResult.Clear()

	if rawResult.VT == 0x1 {
		return nil, err
	}

	return GetVariantValue(rawResult)
}

// SetProperty sets a value of property representation by name with value
func (c *WmiInstance) SetProperty(name string, value interface{}) error {
	rawResult, err := oleutil.PutProperty(c.instance, name, value)
	if err != nil {
		return err
	}

	defer rawResult.Clear()
	return nil
}

// ResetProperty resets a property
func (c *WmiInstance) ResetProperty(name string) error {
	return c.SetProperty(name, nil)
}

// GetClassName
func (c *WmiInstance) GetClassName() string {
	className, err := c.GetSystemProperty("__CLASS")
	if err != nil {
		panic("The class doesn't have a __CLASS member " + err.Error())
	}
	if className == nil {
		panic("The __CLASS member doesn't contain one element, while it was expected to be")
	}
	defer className.Close()

	return className.Value().(string)
}

// Class
func (c *WmiInstance) GetClass() *WmiClass {
	class, err := c.session.GetClass(c.GetClassName())
	if err != nil {
		panic("The class for this instance doesn't exist" + err.Error())
	}

	return class
}

// Class
func (c *WmiInstance) EmbeddedInstance() (string, error) {
	rawResult, err := oleutil.CallMethod(c.instance, "GetObjectText_")
	if err != nil {
		return "", err
	}
	defer rawResult.Clear()
	return rawResult.ToString(), err
}

// Equals
func (c *WmiInstance) Equals(instance *WmiInstance) bool {
	rawResult, err := oleutil.CallMethod(c.instance, "CompareTo_", instance.instance)
	if err != nil {
		return false
	}
	defer rawResult.Clear()
	value, err := GetVariantValue(rawResult)
	if err != nil {
		return false
	}

	return value.(bool)
}

// Clone
func (c *WmiInstance) Clone() (*WmiInstance, error) {
	rawResult, err := oleutil.CallMethod(c.instance, "Clone_")
	if err != nil {
		return nil, err
	}
	return CreateWmiInstance(rawResult, c.session)
	// return c.session.GetInstance(c.InstancePath())
}

// Refresh
func (c *WmiInstance) Refresh() error {
	instance, err := c.session.GetInstance(c.InstancePath())
	if err != nil {
		return err
	}

	c.instance.Release()
	c.instance = instance.instance

	return nil
}

// Commit
func (c *WmiInstance) Commit() error {
	rawResult, err := oleutil.CallMethod(c.instance, "Put_")
	if err != nil {
		return err
	}
	defer rawResult.Clear()
	return nil

}

// Modify
func (c *WmiInstance) Modify() error {
	return c.Commit()
}

// Delete
func (c *WmiInstance) Delete() error {
	rawResult, err := oleutil.CallMethod(c.instance, "Delete_")
	if err != nil {
		return err
	}
	defer rawResult.Clear()
	return nil
}

// InstancePath
func (c *WmiInstance) InstancePath() string {
	path, err := c.GetSystemProperty("__PATH")
	if err != nil {
		panic("The instance doesn't have a path " + err.Error())
	}
	defer path.Close()

	return path.Value().(string)
}

// RelativePath
func (c *WmiInstance) RelativePath() string {
	path, err := c.GetSystemProperty("__RELPATH")
	if err != nil {
		panic("The instance doesn't have a path" + err.Error())
	}
	defer path.Close()

	return path.Value().(string)
}

// InvokeMethod
func (c *WmiInstance) InvokeMethod(methodName string, params ...interface{}) ([]interface{}, error) {
	rawResult, err := oleutil.CallMethod(c.instance, methodName, params...)
	if err != nil {
		return nil, err
	}
	defer rawResult.Clear()
	values, err := GetVariantValues(rawResult)
	return values, err
}

func (c *WmiInstance) GetWmiMethod(methodName string) (*WmiMethod, error) {
	return NewWmiMethod(methodName, c)
}

// InvokeMethodAsync
func (c *WmiInstance) InvokeMethodAsync(methodName string, action UserAction, percentComplete, timeoutSeconds uint32, params ...interface{}) ([]interface{}, error) {
	rawResult, err := oleutil.CallMethod(c.instance, methodName, params...)
	if err != nil {
		return nil, err
	}
	defer rawResult.Clear()
	return GetVariantValues(rawResult)
}

// InvokeMethodWithReturn invokes a method with return
func (c *WmiInstance) InvokeMethodWithReturn(methodName string, params ...interface{}) (int32, error) {
	results, err := c.InvokeMethod(methodName, params...)
	if err != nil {
		return 0, err
	}

	return results[0].(int32), nil
}

// GetAllRelatedWithQuery returns all related instances matching the query
func (c *WmiInstance) GetAllRelatedWithQuery(q *query.WmiQuery) ([]*WmiInstance, error) {
	winstances, err := c.GetAllRelated(q.ClassName)
	if err != nil {
		return nil, err
	}

	if !q.HasFilter() {
		return winstances, nil
	}

	// For now, only Equals is implemented
	filter := q.Filters[0]
	filteredCollection := []*WmiInstance{}
	for _, inst := range winstances {
		propVal, err := inst.GetProperty(filter.Name)
		if err != nil {
			inst.Close()
			continue
		}
		propString := fmt.Sprintf("%v", propVal)
		if propString == filter.Value {
			filteredCollection = append(filteredCollection, inst)
			continue
		}
		inst.Close()
	}
	//fmt.Printf("Query[%s]=>[%d] instances\n", q.String(), len(filteredCollection))
	return filteredCollection, nil
}

// GetAllRelated
func (c *WmiInstance) GetAllRelated(resultClassName string) ([]*WmiInstance, error) {
	return c.GetAssociated("", resultClassName, "", "")
}

// GetRelated
func (c *WmiInstance) GetRelated(resultClassName string) (*WmiInstance, error) {
	result, err := c.GetAllRelated(resultClassName)
	if err != nil {
		return nil, err
	}

	if len(result) == 0 {
		return nil, errors.Wrapf(errors.NotFound, "No Related Items were received for [%s]", resultClassName)
	}
	return result[0], nil
}

// GetRelatedEx
func (c *WmiInstance) GetRelatedEx(associatedClassName, resultClassName, resultRole, sourceRole string) ([]*WmiInstance, error) {
	return c.GetAssociated(associatedClassName, resultClassName, resultRole, sourceRole)
}

// GetAssociated
func (c *WmiInstance) GetAssociated(associatedClassName, resultClassName, resultRole, sourceRole string) ([]*WmiInstance, error) {
	// Documentation here: https://docs.microsoft.com/en-us/windows/win32/wmisdk/swbemobject-associators-
	rawResult, err := oleutil.CallMethod(c.instance, "Associators_",
		associatedClassName,
		resultClassName,
		resultRole,
		sourceRole,
	)
	if err != nil {
		return nil, err
	}

	result := rawResult.ToIDispatch()
	defer rawResult.Clear()

	// Doc: https://docs.microsoft.com/en-us/previous-versions/windows/desktop/automat/dispid-constants
	enum_property, err := result.GetProperty("_NewEnum")
	if err != nil {
		return nil, err
	}
	defer enum_property.Clear()

	// https://docs.microsoft.com/en-us/windows/win32/api/oaidl/nn-oaidl-ienumvariant
	enum, err := enum_property.ToIUnknown().IEnumVARIANT(ole.IID_IEnumVariant)
	if err != nil {
		return nil, err
	}
	if enum == nil {
		return nil, fmt.Errorf("Enum is nil")
	}

	defer enum.Release()

	wmiInstances := []*WmiInstance{}
	for tmp, length, err := enum.Next(1); length > 0; tmp, length, err = enum.Next(1) {
		if err != nil {
			return nil, err
		}

		wmiInstance, err := CreateWmiInstance(&tmp, c.session)
		if err != nil {
			return nil, err
		}

		wmiInstances = append(wmiInstances, wmiInstance)
	}

	return wmiInstances, nil
}

// EnumerateReferencingInstances
func (c *WmiInstance) EnumerateReferencingInstances(resultClassName, sourceRole string) ([]*WmiInstance, error) {
	//Documentation: https://docs.microsoft.com/en-us/windows/win32/wmisdk/swbemobject-references-
	rawResult, err := oleutil.CallMethod(c.instance, "References_", resultClassName, sourceRole)
	if err != nil {
		return nil, err
	}

	result := rawResult.ToIDispatch()
	defer rawResult.Clear()

	// Doc: https://docs.microsoft.com/en-us/previous-versions/windows/desktop/automat/dispid-constants
	enum_property, err := result.GetProperty("_NewEnum")
	if err != nil {
		return nil, err
	}
	defer enum_property.Clear()

	// https://docs.microsoft.com/en-us/windows/win32/api/oaidl/nn-oaidl-ienumvariant
	enum, err := enum_property.ToIUnknown().IEnumVARIANT(ole.IID_IEnumVariant)
	if err != nil {
		return nil, err
	}
	if enum == nil {
		return nil, fmt.Errorf("Enum is nil")
	}

	defer enum.Release()

	wmiInstances := []*WmiInstance{}
	for tmp, length, err := enum.Next(1); length > 0; tmp, length, err = enum.Next(1) {
		if err != nil {
			return nil, err
		}

		wmiInstance, err := CreateWmiInstance(&tmp, c.session)
		if err != nil {
			return nil, err
		}

		wmiInstances = append(wmiInstances, wmiInstance)
	}

	return wmiInstances, nil
}

// CloseAllInstances
func CloseAllInstances(instances []*WmiInstance) {
	for _, instance := range instances {
		instance.Close()
	}
}

// Dispose
func (c *WmiInstance) Close() error {
	if c.instance != nil {
		c.instance.Release()
	}
	if c.instanceVar != nil {
		return c.instanceVar.Clear()
	}
	return nil
}
