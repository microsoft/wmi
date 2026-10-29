// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

using System;
using System.Collections.Generic;
using System.Globalization;
using System.Linq;
using System.Management;
using System.Text;
using Microsoft.wmi.Common;

namespace Microsoft.wmi.GO
{
    public class GOWmiClass : WmiClass
    {
        private string m_rootmodulepath;
        public GOWmiClass(GOWmiSource wSource, string rootmodulepath)
            : base(wSource, "cim.WmiInstance")
        {
            m_rootmodulepath = rootmodulepath;
            if (BaseLibDerivation == Derivation) {
                wSource.AddReference("github.com/microsoft/wmi/pkg/base/instance");
            }
            if (Properties.Count > 0) {
                wSource.AddReference("github.com/microsoft/wmi/pkg/errors");
                wSource.AddReference("reflect");
            }
        }

        public GOWmiClass(ManagementClass wmiClass, GOWmiSource wSource, string rootmodulepath)
            : base(wmiClass, wSource, "cim.WmiInstance")
        {
            m_rootmodulepath = rootmodulepath;
            if (BaseLibDerivation == Derivation) {
                wSource.AddReference("github.com/microsoft/wmi/pkg/base/instance");
            }
            if (Properties.Count > 0) {
                wSource.AddReference("github.com/microsoft/wmi/pkg/errors");
                wSource.AddReference("reflect");
            }
        }

        protected override WmiSource GetWmiSource(string sourceName, WmiModule wModule)
        {
            return new GOWmiSource(sourceName, wModule, m_rootmodulepath);
        }

        public void GenerateMakefile(string outputDir)
        {
            throw new NotImplementedException();
        }

        public override string GetSourceCode()
        {
            StringBuilder sb = new StringBuilder();
            // Create the struct
            sb.AppendFormat(CultureInfo.InvariantCulture,
                "\n// {0} struct", Name);
            sb.AppendFormat(CultureInfo.InvariantCulture,
                "\ntype {0} struct {{ \n", Name);
            sb.AppendFormat("\t*{0}\n", Derivation);
            foreach (var item in Properties)
            {
                sb.AppendLine(item.GetSourceCode().Replace("\n", "\n\t"));
            }

            sb.AppendLine("}");

            foreach (var item in Constructors)
            {
                sb.AppendLine(item.GetSourceCode().Replace("\n", "\n\t"));
            }

            // Create the setter and getter methods  for properties
            foreach (var item in Properties)
            {
                sb.AppendLine(item.Setter);
                sb.AppendLine(item.Getter);
            }

            // Create the methods for the instance
            foreach (var item in Methods)
            {
                sb.AppendLine(item.GetSourceCode());
            }

            foreach (var item in Related.GroupBy(r => r.Name))
            {
                bool multiple = false;
                if (item.Count() > 1)
                    multiple = true;

                sb.AppendLine((item.ElementAt(0) as GOWmiRelated).GetSourceCode(multiple).Replace("\n", "\n\t"));
            }

            return sb.ToString();
        }


        protected override WmiConstructor GetWmiConstructor(List<string> paramText, string derivation, string body)
        {
            GOWmiConstructor method = new GOWmiConstructor(this, derivation);
            method.Name = Name;
            foreach (var item in paramText)
            {
                GOWmiMethodParam mParam = new GOWmiMethodParam(method, item);
                method.Params.Add(mParam);
            }
            method.BodyText = body;
            return method;
        }

        protected override WmiMethod GetWmiMethod(MethodData mData, WmiClass wClass)
        {
            return new GOWmiMethod(mData, wClass as GOWmiClass);
        }

        protected override WmiProperty GetWmiProperty(PropertyData pData, WmiClass wClass)
        {
            return new GOWmiProperty(pData, wClass as GOWmiClass);
        }

        protected override WmiRelated GetWmiRelated(string name, WmiClass parent)
        {
            return new GOWmiRelated(name, parent as GOWmiClass);
        }

        protected override void AddReference(string className)
        {
           // Do nothing
        }

        protected override List<WmiConstructor> GetWmiConstructors()
        {
            List<WmiConstructor> wConstructors = new List<WmiConstructor>();
            string body = Derivation != "cim.WmiInstance" ? 
                String.Format(CultureInfo.InvariantCulture, "tmp, err := New{0}Ex{1}(instance)\n", Derivation, 1) : 
                "tmp, err := instance, nil\n";
            
            wConstructors.Add(GetWmiConstructor(
                new List<string> { "instance *cim.WmiInstance" }, Derivation, body));
            /*
#if false
            wConstructors.Add(GetWmiConstructor(
                    new List<string> 
                { 
                    "string hostName", 
                    "string wmiNamespace=\"" + Parent.Parent.Parent.Name + "\"",
                    "string className=\"" + Name + "\"", 
                },
                    "base(hostName, wmiNamespace, className) ",
                    " ")); 
#endif
            wConstructors.Add(GetWmiConstructor(
               new List<string> 
                { 
                    "hostName string", 
                },
               Derivation,
               " "));

            wConstructors.Add(GetWmiConstructor(
               new List<string> { " hostName string", "userName string ", "password string " },
                Derivation, " "));

            wConstructors.Add(GetWmiConstructor(
                new List<string>
                {
                    " hostName string",
                    "credentials *cred.Credential",
                },
               "",
               " "));

            wConstructors.Add(GetWmiConstructor(
                new List<string> { "string hostName", "string wmiNamespace", "string className", "string userName", "string password" },
                " ", " "));

            wConstructors.Add(GetWmiConstructor(
                new List<string> { "hostName string", "wmiNamespace string ", "query *query.WmiQuery ", "userName string ", "password string " },
                Derivation, " "));

            wConstructors.Add(GetWmiConstructor(
                new List<string> { " hostName string", "wmiNamespace string ", "className string ", "cred *credential.WmiCredential"},
                Derivation, " "));
                */

            body = (Derivation != "cim.WmiInstance" ? 
                String.Format(CultureInfo.InvariantCulture, "tmp, err := New{0}Ex6(hostName, wmiNamespace, userName, password, domainName, query)\n", Derivation) : 
                "tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)\n");
            
            wConstructors.Add(GetWmiConstructor(
                new List<string> { "hostName string", "wmiNamespace string", "userName string", "password string", "domainName string", "query *query.WmiQuery " },
                Derivation, body));
            return wConstructors;
        }
    }
}
