package wmi;

// WmiClass
type WmiClass interface {
	ClassName() string;
	SuperClassName() string;
	ServerName() string;
	Namespace() string;
	SuperClass() *WmiClass;
	Properties() []string;
	Qualifiers() []string;
	Methods() []string;
	MethodParameters(string) []string;
	InvokeMethod(string, []string, string) (error, string);
	Dispose();
}