package main

import (
	tm "github.com/buger/goterm"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func GetRemoteClientHash() (string, error) {
	resp, err := http.Get("https://powbot.org/game/current_client")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return strings.TrimSuffix(string(data), "\n"), nil
}

func ObtainClient(clientDirectory string, hash string) (string, error) {
	clientPath := filepath.FromSlash(path.Join(clientDirectory, "/PowBot.jar"))
	return DownloadToFile("https://powbot.org/game/"+hash+".jar", clientPath)
}

func FindClient(clientDirectory string) (string, error) {
	clientPath := filepath.FromSlash(path.Join(clientDirectory, "/PowBot.jar"))
	if _, err := os.Stat(clientPath); err != nil {
		return "", err
	}
	return clientPath, nil
}

func EnsureClientPresent(powbotDirectory string) (string, error) {
	clientDirectory := filepath.FromSlash(path.Join(powbotDirectory, "/client/"))
	if err := CreateDirectory(clientDirectory); err != nil {
		return "", err
	}

	remoteClientHash, err := GetRemoteClientHash()
	if err != nil {
		return "", err
	}

	client, err := FindClient(clientDirectory)
	if err != nil {
		tm.Println()
		tm.Println(tm.Color("\tClient not found - downloading...", tm.YELLOW))
		tm.Flush()
		client, err = ObtainClient(clientDirectory, remoteClientHash)
	}

	localHash, err := CalculateSHA1(client)
	if err != nil || localHash != remoteClientHash {
		tm.Println()
		tm.Println(tm.Color("\tClient is outdated - downloading...", tm.YELLOW))
		tm.Flush()
		return ObtainClient(clientDirectory, remoteClientHash)
	}
	return client, nil
}
