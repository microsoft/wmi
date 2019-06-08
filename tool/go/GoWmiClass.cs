// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

using System;
using System.Collections.Generic;
using System.Globalization;
using System.Linq;
using System.Management;
using System.Text;
using Microsoft.WmiCodeGen.Common;

namespace Microsoft.WmiCodeGen.GO
{
    public class GOWmiClass : WmiClass
    {
        public GOWmiClass(GOWmiSource wSource)
            : base(wSource)
        {
        }

        public GOWmiClass(ManagementClass wmiClass, GOWmiSource wSource)
            : base(wmiClass, wSource)
        {

        }

        protected override WmiSource GetWmiSource(string sourceName, WmiModule wModule)
        {
            return new GOWmiSource(sourceName, wModule);
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
                "\ntype {0} struct { \n", Name);
            sb.AppendFormat("\t{0}\n", Derivation);
            foreach (var item in Properties)
            {
                sb.AppendLine(item.GetSourceCode().Replace("\n", "\n\t"));
            }

            sb.AppendLine("}");

            // Create the setter and getter methods  for properties
            foreach (var item in Properties)
            {
                sb.AppendLine(item.Setter.Replace("\n", "\n\t"));
                sb.AppendLine(item.Getter.Replace("\n", "\n\t"));
            }

            // Create the methods for the instance
            foreach (var item in Methods)
            {
                sb.AppendLine(item.GetSourceCode().Replace("\n", "\n\t"));
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
