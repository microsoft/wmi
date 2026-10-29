// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

using System;
using System.Collections.Generic;
using System.Globalization;
using System.Linq;
using System.Management;
using System.Text;

using Microsoft.wmi.Common;

namespace Microsoft.wmi.CSharp
{
    public class CSWmiClass : WmiClass
    {
        public CSWmiClass(CSWmiSource wSource)
            : base(wSource)
        {
        }

        public CSWmiClass(ManagementClass wmiClass, CSWmiSource wSource)
            : base(wmiClass, wSource)
        {

        }

        protected override WmiSource GetWmiSource(string sourceName, WmiModule wModule)
        {
            return new CSWmiSource(sourceName, wModule);
        }

        public void GenerateMakefile(string outputDir)
        {
            throw new NotImplementedException();
        }

        public override string GetSourceCode()
        {
            StringBuilder sb = new StringBuilder();
            sb.AppendFormat(CultureInfo.InvariantCulture,
                "\npublic {0} class {1} : {2}\n", Abstract ? "/* abstract */" : "", Name, Derivation);
            sb.AppendLine("{");

            foreach (var item in Constructors)
            {
                sb.AppendLine(item.GetSourceCode().Replace("\n", "\n\t"));
            }

            foreach (var item in Properties)
            {
                sb.AppendLine(item.GetSourceCode().Replace("\n", "\n\t"));
            }

            foreach (var item in Methods)
            {
                sb.AppendLine(item.GetSourceCode().Replace("\n", "\n\t"));
            }

            foreach (var item in Related.GroupBy(r => r.Name))
            {
                bool multiple = false;
                if (item.Count() > 1)
                    multiple = true;

                sb.AppendLine((item.ElementAt(0) as CSWmiRelated).GetSourceCode(multiple).Replace("\n", "\n\t"));
            }

#if false
            if (Name.Equals("CIM_ConcreteJob"))
            {
                sb.Append(CIM_ConcreteJobText).Replace("\n", "\n\t");
            } 
#endif

            sb.AppendLine("}");
            return sb.ToString();
        }


        protected override WmiConstructor GetWmiConstructor(List<string> paramText, string derivation, string body)
        {
            CSWmiConstructor method = new CSWmiConstructor(this, derivation);
            method.Name = Name;
            foreach (var item in paramText)
            {
                CSWmiMethodParam mParam = new CSWmiMethodParam(method, item);
                method.Params.Add(mParam);
            }
            method.BodyText = body;
            return method;
        }

        protected override WmiMethod GetWmiMethod(MethodData mData, WmiClass wClass)
        {
            return new CSWmiMethod(mData, wClass as CSWmiClass);
        }

        protected override WmiProperty GetWmiProperty(PropertyData pData, WmiClass wClass)
        {
            return new CSWmiProperty(pData, wClass as CSWmiClass);
        }

        protected override WmiRelated GetWmiRelated(string name, WmiClass parent)
        {
            return new CSWmiRelated(name, parent as CSWmiClass);
        }

        protected override void AddReference(string className)
        {
            if (!Parent.Parent.Parent.HasSource(className) &&
                    !Parent.CheckClass(className, Parent.Parent.Parent.Name))
            {
                Logger.Debug("AddReference: Class not found in the current Namespace. Start searching from root namespace");
                string reference = Parent.GetReference(className, "root");
                if (!string.IsNullOrEmpty(reference))
                {
                    Logger.Debug("WmiClass AddReference {0}", reference);
                    Parent.AddReference(reference);
                    Parent.Parent.Parent.AddReference(reference);
                }
                else
                {
                    Logger.Debug("No reference found anywhere. Create a Dummy Source - {0}", className);
                    Parent.Parent.Parent.AddSource(GetWmiSource(className,
                        Parent.Parent.Parent.AddModule(CSWmiSource.GetModuleName(className))));
                }
            }
        }
        protected override List<WmiConstructor> GetWmiConstructors()
        {
            List<WmiConstructor> wConstructors = new List<WmiConstructor>();
            wConstructors.Add(GetWmiConstructor(
                new List<string> { "IWmiInstance instance" }, "base(instance) ", " "));
            wConstructors.Add(GetWmiConstructor(
                new List<string> { "WmiInstancePath instancePath" }, "base(instancePath) ", " "));

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
                    "string hostName", 
                },
               string.Format(CultureInfo.InvariantCulture,
               "this(hostName, default(string), default(string)) "),
               " "));

            wConstructors.Add(GetWmiConstructor(
               new List<string> { "string hostName", "string userName", "string password" },
                "this (hostName, new WmiCredentials (userName, password)) ", " "));

            wConstructors.Add(GetWmiConstructor(
                new List<string>
                {
                    "string hostName",
                    "WmiCredentials credentials",
                },
               string.Format(CultureInfo.InvariantCulture,
               "this(hostName, \"{0}\", \"{1}\", credentials) ", Parent.Parent.Parent.Name, Name),
               " "));

            wConstructors.Add(GetWmiConstructor(
                new List<string> { "string hostName", "string wmiNamespace", "string className", "string userName", "string password" },
                "this (hostName, wmiNamespace, className, new WmiCredentials (userName, password)) ", " "));

            wConstructors.Add(GetWmiConstructor(
                new List<string> { "string hostName", "string wmiNamespace", "IWmiQuery query", "string userName", "string password" },
                "this (hostName, wmiNamespace, query, new WmiCredentials(userName, password)) ", " "));

            wConstructors.Add(GetWmiConstructor(
                new List<string> { "string hostName", "string wmiNamespace", "string className", "WmiCredentials credentials"},
                "base(hostName, wmiNamespace, className, credentials) ", " "));

            wConstructors.Add(GetWmiConstructor(
                new List<string> { "string hostName", "string wmiNamespace", "IWmiQuery query", "WmiCredentials credentials" },
                "base(hostName, wmiNamespace, query, credentials) ", " "));

            return wConstructors;
        }
    }
}
