//go:generate goversioninfo -icon=resources/icon.ico -manifest=resources/powbot-launcher.manifest
package main

import (
	"bufio"
	"crypto/sha1"
	"fmt"
	tm "github.com/buger/goterm"
	"github.com/fatih/color"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
	"time"
)

func CreateDirectory(directory string) (err error) {
	if _, err := os.Stat(directory); err != nil {
		if os.IsNotExist(err) {
			if err := os.MkdirAll(directory, 0700); err != nil {
				return err
			}
		} else {
			return err
		}
	}
	return nil
}

func DownloadToFile(downloadURL string, destFile string) (string, error) {
	resp, err := http.Get(downloadURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	outputPath := filepath.FromSlash(destFile)
	out, err := os.Create(outputPath)
	if err != nil {
		return "", err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return outputPath, err
}

func Download(downloadURL string, dest string) (string, error) {
	outputPath := filepath.FromSlash(path.Join(dest, path.Base(downloadURL)))
	return DownloadToFile(downloadURL, outputPath)
}

func CalculateSHA1(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := sha1.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}

func main() {
	tm.Output = bufio.NewWriter(color.Output)
	tm.Clear()

	tm.MoveCursor(1, 1)
	tm.Println(tm.Color(tm.Bold("PowBot is starting..."), tm.RED))
	tm.Print(tm.Color("Creating home directory...", tm.BLUE))
	tm.Flush()
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	powbotDirectory := filepath.FromSlash(path.Join(homeDir, "/.powbot/"))
	if err := CreateDirectory(powbotDirectory); err != nil {
		tm.Println(tm.Color("\tFailed.", tm.RED))
		tm.Flush()
		log.Fatal(err)
	} else {
		tm.Println(tm.Color("\tDone.", tm.GREEN))
		tm.Flush()
	}

	tm.Print(tm.Color("Checking Java installation...", tm.BLUE))
	tm.Flush()
	java, err := EnsureJREPresent(powbotDirectory)
	if err != nil {
		tm.Println(tm.Color("\tFailed.", tm.RED))
		tm.Flush()
		log.Fatal(err)
	} else {
		tm.Println(tm.Color("\tDone.", tm.GREEN))
		tm.Flush()
	}

	tm.Print(tm.Color("Checking Client version...", tm.BLUE))
	client, err := EnsureClientPresent(powbotDirectory)
	if err != nil {
		tm.Println(tm.Color("\tFailed.", tm.RED))
		tm.Flush()
		log.Fatal(err)
	} else {
		tm.Println(tm.Color("\tDone.", tm.GREEN))
		tm.Flush()
	}

	tm.Println(tm.Color(tm.Bold("Launching the client - this window will close in 5 seconds."), tm.GREEN))
	tm.Flush()

	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/D", powbotDirectory, "/c", "start "+java+" -jar "+client)
		cmd.Dir = powbotDirectory
		err = cmd.Start()
		if err != nil {
			log.Fatal(err)
		}
		cmd.Process.Release()
	} else {
		cmd := exec.Command("/bin/sh", "-c", java+" -jar "+client)
		cmd.Dir = powbotDirectory
		err = cmd.Start()
		if err != nil {
			log.Fatal(err)
		}
		cmd.Process.Release()
	}

	time.Sleep(5 * time.Second)
}
