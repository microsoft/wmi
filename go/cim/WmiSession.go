package cim

import (
	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
	"github.com/microsoft/wmicodegen/go/wmi"
)

// WmiSession struct to hold the current session information
type WmiSession struct {
	ServerName    string
	Namespace     string
	Username      string
	Password      string
	Domain        string
	Status        wmi.SessionStatus
	RawSession    *ole.VARIANT
	Session       *ole.IDispatch
	CimwmiService *ole.IDispatch
}

// CreateSessionEx creates a session based on credentials
func CreateSessionEx(cimwmi *ole.IDispatch, serverName, wmiNamespace string, credentials wmi.Credentials) (*WmiSession, error) {
	return CreateSession(cimwmi, serverName, wmiNamespace, credentials.UserName, credentials.Password, credentials.Domain)
}

// CreateSession creates a new session with the server and namespace
func CreateSession(cimwmi *ole.IDispatch, serverName, wmiNamespace, userName, password, domain string) (*WmiSession, error) {
	session := WmiSession{
		CimwmiService: cimwmi,
		ServerName:    serverName,
		Namespace:     wmiNamespace,
		Username:      userName,
		Password:      password,
		Domain:        domain,
		Status:        wmi.Created,
	}
	return &session, nil
}

// Connect the wmi session
func (c WmiSession) Connect() (bool, error) {
	var err error
	c.RawSession, err = oleutil.CallMethod(c.CimwmiService, "ConnectServer", nil)
	if err != nil {
		return false, err
	}
	c.Session = c.RawSession.ToIDispatch()
	c.Status = wmi.Connected
	return true, nil
}

// Close the wmi session
func (c WmiSession) Close() {
	c.RawSession.Clear()
	c.Status = wmi.Disconnected
}

// Dispose the wmi session
func (c WmiSession) Dispose() {
	if c.Status != wmi.Disposed {
		c.Close()
		c.Status = wmi.Disposed
	}
}

// TestConnection
func (c WmiSession) TestConnection() bool {
	panic("not implemented")
}

// GetClass
func (c WmiSession) GetClass(namespaceName, className string) (*wmi.Class, error) {
	panic("not implemented")
}

//GetInstance
func (c WmiSession) GetInstance(namespaceName string, instance *wmi.Instance) (*wmi.Instance, error) {
	panic("not implemented")
}

// Dispose
func (c WmiSession) EnumerateClasses(namespaceName, className string) (*[]wmi.Class, error) {
	panic("not implemented")
}

// Dispose
func (c WmiSession) EnumerateInstances(namespaceName, className string) (*[]wmi.Instance, error) {
	panic("not implemented")
}

// Dispose
func (c WmiSession) QueryInstances(namespaceName, queryDialect, queryExpression string) (*[]wmi.Instance, error) {
	rawResult, err := c.Session.CallMethod("ExecQuery", queryExpression)
	if err != nil {
		return nil, err
	}

	result := rawResult.ToIDispatch()
	defer rawResult.Clear()

	value, err := oleutil.GetProperty(result, "Count")
	if err != nil {
		return nil, err
	}

	defer value.Clear()
	count := int64(value.Val)

	return nil, nil
}

// Dispose
func (c WmiSession) EnumerateReferencingInstances(namespaceName string, sourceInstance wmi.Instance, associationClassName, sourceRole string) (*[]wmi.Instance, error) {
	panic("not implemented")
}

// Dispose
func (c WmiSession) ModifyInstance(namespaceName string, instance wmi.Instance) {
	panic("not implemented")
}

// Dispose
func (c WmiSession) DeleteInstance(namespaceName string, instance wmi.Instance) {
	panic("not implemented")
}

// InvokeMethod
func (c WmiSession) InvokeMethod(namespaceName string, instance *wmi.Instance, methodName string, methodParameters *[]wmi.MethodParameter) (wmi.MethodResult, error) {

}
