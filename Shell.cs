using System;
using System.IO;
using System.Diagnostics;
using System.Collections.Generic;
using System.Collections.ObjectModel;

namespace powbot_launcher_v2
{
    class Shell
    {

        public static Boolean Execute(string process, string dir, bool background, List<string> args) {
            try
            {
                OperatingSystem os = Environment.OSVersion;
                PlatformID pid = os.Platform;
                switch (pid) 
                {
                    // case PlatformID.Win32NT:
                    // case PlatformID.Win32S:
                    // case PlatformID.Win32Windows:
                    // case PlatformID.WinCE:
                        // using (Process proc = Process.Start("/bin/bash", $"-c \"{cmd}\""))
                        // {
                        //     proc.WaitForExit();
                        //     return proc.ExitCode == 0;
                        // }
                    // default:
                    // Console.WriteLine(cmd);
                        // using (Process proc = Process.Start("/bin/bash", $"-c \"{cmd}\""))
                        // {
                        //     proc.WaitForExit();
                        //     proc.Close();
                        //     return proc.ExitCode == 0;
                        // }
                }

                Process p = new Process();
                p.StartInfo = new ProcessStartInfo(process);
                foreach (var arg in args)
                {
                    p.StartInfo.ArgumentList.Add(arg);
                }

                p.StartInfo.WorkingDirectory = dir;
                p.StartInfo.CreateNoWindow = background;
                p.Start();
            }
            catch(Exception e)
            {
                Console.WriteLine(e.ToString());
            }
            return false;
        }
    }
}
