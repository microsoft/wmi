// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

using System;
using System.Collections.Generic;
using System.Globalization;
using System.IO;
using System.Management;
using System.Text;

namespace Microsoft.WmiCodeGen
{
    public abstract class IFormat
    {
        Dictionary<string, WmiNamespace> m_nsList = new Dictionary<string, WmiNamespace>();
        string[] m_Namespaces;
        string m_outputDir;
        bool m_Recurse;

        #region Static Methods
        public static bool HasQualifier(QualifierDataCollection qDataCollection, string qName)
        {
            return GetQualifierValue(qDataCollection, qName) != null ? true : false;
        }
        public static bool TryGetQualifierValue(QualifierDataCollection qDataCollection, string qName, out object value)
        {
            value = GetQualifierValue(qDataCollection, qName);
            if (value != null) return true;
            return false;
        }
        public static Object GetQualifierValue(QualifierDataCollection qDataCollection, string qName)
        {
            foreach (QualifierData item in qDataCollection)
            {
                if (item.Name.Equals(qName, StringComparison.OrdinalIgnoreCase))
                {
                    return item.Value;
                }
            }
            return null;
        }
        public static bool IsAbstract(QualifierDataCollection qDataCollection)
        {
            return HasQualifier(qDataCollection, "abstract");
        }

        public static bool IsWritable(QualifierDataCollection qDataCollection)
        {
            return HasQualifier(qDataCollection, "write");
        }

        public static string GetDescription(QualifierDataCollection qDataCollection)
        {
            object description;
            if (TryGetQualifierValue(qDataCollection, "description", out description)
                || TryGetQualifierValue(qDataCollection, "Description", out description)
                )
            {
                if (!string.IsNullOrEmpty(description.ToString()))
                {
                    return description.ToString();
                }
            }
            return string.Empty;
        }

        public static string GetDescriptionText(QualifierDataCollection qDataCollection)
        {
            object description;
            if (TryGetQualifierValue(qDataCollection, "description", out description) ||
                TryGetQualifierValue(qDataCollection, "Description", out description))
            {
                if (!string.IsNullOrEmpty(description.ToString()))
                {
                    return string.Format(CultureInfo.InvariantCulture,
                        "\n[Description(@\"{0}\")]", description.ToString().Replace("\"", "'"));
                }
            }
            return string.Empty;
        }


        public static bool HasOutParams(MethodData mData)
        {
            if (mData.OutParameters != null)
            {
                foreach (var outParam in mData.OutParameters.Properties)
                {
                    if (outParam.Name != "ReturnValue")
                    {
                        return true;
                    }
                }
            }
            return false;
        }

        public static bool HasJobOutputParams(MethodData mData)
        {
            if (mData.OutParameters != null)
            {
                foreach (var outParam in mData.OutParameters.Properties)
                {
                    if (outParam.Name == "Job")
                    {
                        return true;
                    }
                }
            }
            return false;
        }

        #endregion

        #region Constructor
        public IFormat(string[] wmiNamespaces, string outDir, bool recurse)
        {
            m_Namespaces = wmiNamespaces;
            m_outputDir = outDir;
            m_Recurse = recurse;
        }
        #endregion

        protected abstract WmiNamespace GetWmiNamespace(string wmins);

        private void TraverseChildNamespaces(string wmiNamespace)
        {
            if (wmiNamespace.Equals("root/rsop/user", StringComparison.OrdinalIgnoreCase))
            {
                Logger.Info("Namespace {0} not supported by this tool", wmiNamespace);
                return;
            }
            Logger.Info("Namespace {0}", wmiNamespace);
            var wNamespace = GetWmiNamespace(wmiNamespace);
            m_nsList[wmiNamespace] = wNamespace;
            wNamespace.GenerateSources(m_outputDir);

            List<string> ExcludeFailbackNamespace = new List<string> { "ms_409", "ms_509", "ms_501" };

            if (m_Recurse)
            {
                using (ManagementClass cls = new ManagementClass(new ManagementScope(wmiNamespace), new ManagementPath("__namespace"), null))
                {
                    try
                    {
                        foreach (var instance in cls.GetInstances())
                        {
                            string instanceName = instance["Name"].ToString();
                            if (!ExcludeFailbackNamespace.Contains(instanceName.ToLowerInvariant())) // skip the failback namespaces
                                TraverseChildNamespaces(string.Format("{0}/{1}", wmiNamespace, instanceName));
                            else
                                Logger.Debug("Skipping {0}/{1}", wmiNamespace, instanceName);
                        }
                    }
                    catch (Exception e)
                    {
                        Logger.Exception(e, "TraverseChildNamespaces");
                    }
                }
            }
        }

        public void GenerateAllWmiClasses()
        {
            foreach (var item in m_Namespaces)
            {
                TraverseChildNamespaces(item);
            }
            GenerateDirFile(m_outputDir);
        }
        public void GenerateDirFile(string outDirectory)
        {
            foreach (KeyValuePair<string, WmiNamespace> kvp in m_nsList)
            {
                string sPath = Path.Combine(Environment.CurrentDirectory, outDirectory, kvp.Value.Name);
                while (sPath != null)
                {
                    string dirPath = Path.Combine(sPath, "dirs");
                    if (Directory.Exists(sPath) && !File.Exists(dirPath))
                    {
                        StringBuilder sb = new StringBuilder();
                        sb.AppendLine(@"DIRS= \");
                        foreach (var item in Directory.GetDirectories(sPath))
                        {
                            sb.AppendFormat(CultureInfo.InvariantCulture, "\t{0} \\\n", Path.GetFileName(item));
                        }
                        File.WriteAllText(dirPath, sb.ToString());
                        Logger.Info("Generated Dirs file at {0}", dirPath);
                    }
                    sPath = Path.GetDirectoryName(sPath);
                }
            }
        }
    }
}
