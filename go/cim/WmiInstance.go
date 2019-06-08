package cim

import (
	"reflect"

	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
	"github.com/microsoft/wmicodegen/go/wmi"
)

type WmiInstance struct {
	class    *WmiClass
	session  *WmiSession
	instance *ole.IDispatch
}

type WmiInstanceCollection []WmiInstance

// GetInstance
func (c WmiInstance) GetInstance() string {
	panic("not implemented")

}

// GetProperty gets the property of the instance specified by name and returns in value
func (c WmiInstance) GetProperty(name string, value interface{}) error {
	rawResult, err := oleutil.GetProperty(c.instance, name)
	if err != nil {
		return err
	}

	defer rawResult.Clear()

	if rawResult.VT == 0x1 {
		return err
	}

	getPropertyValue(rawResult, value)

	return err

}

// SetProperty
func (c WmiInstance) SetProperty(name string, value interface{}) error {
	rawResult, err := oleutil.PutProperty(c.instance, name)
	if err != nil {
		return err
	}

	defer rawResult.Clear()
	return err

}

// ResetProperty
func (c WmiInstance) ResetProperty(name string) error {
	return c.SetProperty(name, nil)
}

// Class
func (c WmiInstance) Class() wmi.Class {
	return c.class
}

// Class
func (c WmiInstance) EmbeddedInstance() (string, error) {
	rawResult, err := oleutil.CallMethod(c.instance, "GetObjectText")
	if err != nil {
		return "", err
	}
	defer rawResult.Clear()
	return rawResult.ToString(), err
}

// InstanceManager
func (c WmiInstance) InstanceManager() *wmi.InstanceManager {
	panic("not implemented")

}

// Equals
func (c WmiInstance) Equals(*wmi.Instance) bool {
	panic("not implemented")

}

// Refresh
func (c WmiInstance) Refresh() error {
	panic("not implemented")

}

// Commit
func (c WmiInstance) Commit() error {
	rawResult, err := oleutil.CallMethod(c.instance, "Put")
	if err != nil {
		return err
	}
	defer rawResult.Clear()
	return nil

}

// Modify
func (c WmiInstance) Modify() error {
	return c.Commit()
}

// Delete
func (c WmiInstance) Delete() error {
	rawResult, err := oleutil.CallMethod(c.instance, "Delete")
	if err != nil {
		return err
	}
	defer rawResult.Clear()
	return nil
}

// InstancePath
func (c WmiInstance) InstancePath() (string, error) {
	panic("not implemented")

}

// InvokeMethod
func (c WmiInstance) InvokeMethod(methodName string, methodParameters *[]wmi.MethodParameter) (*wmi.MethodResult, error) {
	rawResult, err := oleutil.CallMethod(c.instance, "ExecMethod")
	if err != nil {
		return nil, err
	}
	defer rawResult.Clear()
	return nil, err
}

// InvokeMethodWithReturn invokes a method with return
func (c WmiInstance) InvokeMethodWithReturn(methodName string, methodParameters *[]wmi.MethodParameter, returnValue interface{}) (uint32, error) {

	result, err := c.InvokeMethod(methodName, methodParameters)
	if err != nil {
		return 0, err
	}
	returnValue = result.ReturnValue.Value

	switch reflect.Value(returnValue).Kind {
	case reflect.Bool:
		return uint32(returnValue), nil
	case reflect.Uint, reflect.Int:
		return uint(returnValue), nil
	}

	return 0, nil
}

// GetRelated
func (c WmiInstance) GetRelated(resultClassName string) *[]wmi.Instance {
	panic("not implemented")

}

// Rel
func (c WmiInstance) GetRelatedEx(resultClassName, associatedClassName, resultRole, sourceRole string) *[]wmi.Instance {
	panic("not implemented")

}

// GetAssociated
func (c WmiInstance) GetAssociated(resultClassName, associatedClassName, resultRole, sourceRole string) *[]wmi.Instance {

	rawResult, err := oleutil.CallMethod(c.instance, "Associators")
	if err != nil {
		return nil, err
	}
	defer rawResult.Clear()
	return nil, err
}

// EnumerateReferencingInstances
func (c WmiInstance) EnumerateReferencingInstances(associatedClassName, sourceRole string) *[]wmi.Instance {
	rawResult, err := oleutil.CallMethod(c.instance, "References")
	if err != nil {
		return nil, err
	}
	defer rawResult.Clear()
	return nil, err
}

func getPropertyValue(rawValue *ole.VARIANT, value interface{}) error {

	v := reflect.ValueOf(value)

	switch v.Kind() {
	case reflect.Slice, reflect.Array:
		value := make([]interface{}, v.Len())
		for i := 0; i < v.Len(); i++ {
			value[i] = v.Index(i).Interface()
			// FixMe
		}

	default:
		value = rawValue.Value
	}

	return nil
}
