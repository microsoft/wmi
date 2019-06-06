package wmi;

type WmiMethodDeclaration struct {
	Name string;
	Parameters *[]WmiMethodParameter
	Qualifiers *[]WmiQualifier
}