// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

using System;
using Microsoft.WmiCodeGen.Common;
using Microsoft.WmiCodeGen.CSharp;
using Microsoft.WmiCodeGen.GO;

namespace Microsoft.WmiCodeGen.Tool
{
    class WmiClassGenenerator
    {
        enum Params
        {
            ns = 1,
            wmiclass = 2,
            unused = 4,
            outdir = 8,
            help = 16,
            generatesource = 32,
            recurse = 64,
            format = 128,
            debug = 256,
            related = 512,
        }

        enum Format
        {
            CS,
            PS,
            GO
        }

        private static Params m_modes;
        private static string m_classname;
        private static string[] m_namespace;
        private static string[] m_related;
        private static string m_outputDir = "output";
        private static Format m_fmt = Format.CS; // Default FormaT
        private static bool Parse(string[] args)
        {
            for (int i = 0; i < args.Length; i++)
            {
                Params param;
                if (!args[i].Contains("-"))
                {
                    throw new Exception("UnExpected Params " + args[i]);
                }
                if (Enum.TryParse<Params>(args[i].Split(new char[] { '-' })[1], out param))
                {
                    m_modes |= param;
                    switch (param)
                    {
                        case Params.ns:
                            if (++i < args.Length)
                            {
                                m_namespace = args[i].Split(new char[] { ',' });
                                for (int j = 0; j < m_namespace.Length; j++)
                                {
                                    m_namespace[j].Replace('\\', '/');
                                }
                            }
                            else
                            {
                                Logger.Info("Wmi Namespace names not specified");
                            }
                            break;
                        case Params.related:
                            if (++i < args.Length)
                                m_related = args[i].Split(new char[] { ',' });
                            else
                            {
                                Logger.Info("Related Module names not specified");
                            }
                            break;
                        case Params.wmiclass:
                            if (++i < args.Length)
                                m_classname = args[i];
                            break;
                        case Params.outdir:
                            if (++i < args.Length)
                                m_outputDir = args[i];
                            break;
                        case Params.generatesource:
                        case Params.recurse:
                        case Params.debug:
                            break;
                        case Params.format:
                            if (++i < args.Length)
                                m_fmt = (Format)Enum.Parse(typeof(Format), args[i]);
                            break;
                        case Params.help:
                        default:
                            return false;
                    }
                }
            }
            return true;
        }

        private static void Usage(string[] args)
        {
            Logger.Info("{0}", "WmiClassGen.exe");
            Logger.Info("\t -generatesource");
            Logger.Info("\t\tDescription : Generate Source");
            Logger.Info("\t\tExample     : -generatesource");
            Logger.Info("");
            Logger.Info("\t -ns");
            Logger.Info("\t\tDescription : Namespace Name");
            Logger.Info("\t\tExample     : -ns root/virtualization/v2");
            Logger.Info("\t\tExample     : -ns root/virtualization/v2,root/cimv2");
            Logger.Info("");
            Logger.Info("\t -recurse");
            Logger.Info("\t\tDescription : Recurse the namespace");
            Logger.Info("\t\tExample     : -recurse");
            Logger.Info("");
            Logger.Info("\t -format <Format Type>");
            Logger.Info("\t\t Values     : {0}", string.Join(", ", Enum.GetNames(typeof(Format))));
            Logger.Info("\t\tDescription : Format of the generated Source file");
            Logger.Info("\t\tExample     : -format CS ");
            Logger.Info("");
            Logger.Info("\t -wmiclass");
            Logger.Info("\t\tDescription : WMI Class Name");
            Logger.Info("\t\tExample     : -wmiclass Msvm_ComputerSystem");
            Logger.Info("\t\tOptional");
            Logger.Info("");
            Logger.Info("\t -outdir <output directory path>");
            Logger.Info("\t\tDescription : Output Directory Path");
            Logger.Info("\t\tExample     : -outdir test");
            Logger.Info("\t\tOptional. Default : output");
            Logger.Info("");
            Logger.Info("\t -help");
            Logger.Info("\t\tDescription : Display this text");
            Logger.Info("\t\tExample     : -help");
            Logger.Info("\t\tOptional");
            Logger.Info(""); 
        }

        static void Main(string[] args)
        {
            if (args.Length < 1)
            {
                Usage(args);
                return;
            }

            try
            {
                if (!Parse(args))
                {
                    Usage(args);
                    return;
                }
            }
            catch (Exception)
            {
                Usage(args);
            }

            if ((m_modes & Params.debug) != 0)
            {
                //Logger.Logger.CurrentLogLevel = LogLevel.Debug;
            }

            if ((m_modes & Params.outdir) == 0)
            {
                m_outputDir = ".";
            }

            if ((m_modes & Params.generatesource) != 0)
            {
                // Generate all classes in this Namespace, if no particular namespace is given
                if ((m_modes & Params.ns) != 0)
                {
                    bool recurse = (m_modes & Params.recurse) != 0;
                    switch (m_fmt)
                    {
                        case Format.CS:
                            Logger.Debug("Generating C# Source");
                            CSharpFormat csfmt = new CSharpFormat(m_namespace, m_outputDir, recurse);
                            csfmt.GenerateSources();
                            break;
                        case Format.PS:
                            break;
                        case Format.GO:
                            Logger.Debug("Generating C# Source");
                            GoFormat fmt = new GoFormat(m_namespace, m_outputDir, recurse);
                            fmt.GenerateSources();
                            break;
                        default:
                            throw new NotImplementedException();
                    }
                }
            }
        }
    }
}
