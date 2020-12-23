using System;

namespace powbot_launcher_v2
{
    class HomeFolder
    {

        public static string GetDirectory() {
            string homeFolder = System.Environment.GetFolderPath(System.Environment.SpecialFolder.Personal);
            return System.IO.Path.Combine(homeFolder, ".powbot");
        }
        public static void Create()
        {
            string powbotHomeFolder = GetDirectory();
            if (!System.IO.Directory.Exists(powbotHomeFolder)) {
                Console.WriteLine($"Creating powbot home directory at {powbotHomeFolder}");
                System.IO.Directory.CreateDirectory(powbotHomeFolder);
            }
        }
    }
}
