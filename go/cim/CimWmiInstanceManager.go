package 
import (
	"github.com/microsoft/wmicodegen/go/wmi"
)

// CimWmiInstanceManager
type CimWmiInstanceManager struct {

}


// ServerName
func (c CimWmiClass) ServerName() string { 	panic("not implemented") }
// Namespace
func (c CimWmiClass) Namespace() string{ 	panic("not implemented") }
// Credentials
func (c CimWmiClass) Credentials() *WmiCredentials{ 	panic("not implemented") }
// EnumerateInstances
func (c CimWmiClass) EnumerateInstances( className string) (*[]WmiInstance, error){ 	panic("not implemented") }
// QueryInstances
func (c CimWmiClass) QueryInstances( query string)(*[]WmiInstance , error){ 	panic("not implemented") }
// CreateInstance
func (c CimWmiClass) QueryInstances(query WmiQuery) (*[]WmiInstance, error){ 	panic("not implemented") }

// ClassName
func (c CimWmiClass) CreateInstance(className string, propertyValues map[string]string ) (*WmiInstance, error){ 	panic("not implemented") }
// GetInstance
func (c CimWmiClass) GetInstance( className string, propertyValues map[string]string )(*WmiInstance, error){ 	panic("not implemented") }
// GetClass
func (c CimWmiClass) GetClass( className string)(*WmiClass, error){ 	panic("not implemented") }
// EnumerateClasses
func (c CimWmiClass) EnumerateClasses()(*[]WmiClass, error){ 	panic("not implemented") }
// GetInstancesFromPaths
func (c CimWmiClass) GetInstancesFromPaths(pathArray string[] )(*[]WmiInstance, error){ 	panic("not implemented") }