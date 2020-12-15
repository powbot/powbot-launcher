package main

import (
	"errors"
	tm "github.com/buger/goterm"
	"os"
	"path"
	"path/filepath"
	"runtime"
)

func GetJREDownloadURL() string {
	if runtime.GOOS == "windows" {
		return "https://github.com/AdoptOpenJDK/openjdk11-binaries/releases/download/jdk-11.0.9.1%2B1/OpenJDK11U-jre_x86-32_windows_hotspot_11.0.9.1_1.zip"
	} else if runtime.GOOS == "linux" {
		return "https://github.com/AdoptOpenJDK/openjdk11-binaries/releases/download/jdk-11.0.9.1%2B1/OpenJDK11U-jre_x64_linux_hotspot_11.0.9.1_1.tar.gz"
	} else {
		return "https://github.com/AdoptOpenJDK/openjdk11-binaries/releases/download/jdk-11.0.9.1%2B1/OpenJDK11U-jre_x64_mac_hotspot_11.0.9.1_1.tar.gz"
	}
}

func GetBinaryName() string {
	if runtime.GOOS == "windows" {
		return "javaw.exe"
	} else if runtime.GOOS == "linux" {
		return "javaw"
	} else {
		return "javaw"
	}
}

func DownloadJRE(dest string) (string, error) {
	downloadURL := GetJREDownloadURL()
	return Download(downloadURL, dest)
}

func UnpackJRE(pkg string) error {
	extension := path.Ext(pkg)
	if extension == ".zip" {
		return Unzip(pkg, filepath.Dir(pkg))
	} else if extension == ".gz" {
		return Untar(pkg, filepath.Dir(pkg))
	} else {
		print("Unsupported extension " + extension)
	}
	return nil
}

func FindJava(powbotDirectory string) (string, error) {
	jreDirectory := filepath.FromSlash(powbotDirectory + "/jre/")
	if _, err := os.Stat(jreDirectory); err != nil {
		return "", err
	}

	var javaPath = ""

	err := filepath.Walk(jreDirectory, func(path string, info os.FileInfo, err error) error {
		if err == nil && info.Name() == GetBinaryName() {
			absolutePath, _ := filepath.Abs(path)
			javaPath = absolutePath
		}
		return nil
	})

	if err != nil {
		return "", err
	} else if len(javaPath) == 0 {
		return "", errors.New("java.exe not found in " + jreDirectory)
	}
	return javaPath, nil
}

func ObtainJRE(powbotDirectory string) (string, error) {
	jreDirectory := filepath.FromSlash(powbotDirectory + "/jre/")
	if err := CreateDirectory(jreDirectory); err != nil {
		tm.Println(tm.Color("\tCouldn't create directory for JRE at " + jreDirectory, tm.RED))
		tm.Flush()
		return "", err
	}

	jrePackage, err := DownloadJRE(jreDirectory)
	if err != nil {
		tm.Println(tm.Color("\tCouldn't download JRE to " + jreDirectory, tm.RED))
		tm.Flush()
		return "", err
	}

	err = UnpackJRE(jrePackage)
	if err != nil {
		tm.Println(tm.Color("\tCouldn't unpack JRE to " + jreDirectory, tm.RED))
		tm.Flush()
		return "", err
	}

	java, err := FindJava(powbotDirectory)
	if err != nil {
		tm.Println(tm.Color("\tCouldn't find the unpacked Java executable in " + jreDirectory, tm.RED))
		tm.Flush()
		return "", err
	}
	return java, nil
}

func EnsureJREPresent(powbotDirectory string) (string, error) {
	java, err := FindJava(powbotDirectory)
	if err != nil {
		tm.Println()
		tm.Println(tm.Color("\tBundled Java not found - downloading...", tm.YELLOW))
		tm.Flush()
		return ObtainJRE(powbotDirectory)
	}
	return java, nil
}
