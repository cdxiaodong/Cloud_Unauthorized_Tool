package docker

import (
	"fmt"
	"net/http"
	"strings"
	"io"
)

func Check(target string) bool {
	removeProto := strings.Replace(target, "http://", "", -1)
	removeProto = strings.Replace(removeProto, "https://", "", -1)
	removeProto = strings.Replace(removeProto, "/version", "", -1)
	removeProto = strings.Replace(removeProto, ":2375", "", -1)

	url := fmt.Sprintf("http://%s:2375/version", removeProto)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error checking %s: %v\n", url, err)
		return false
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response from %s: %v\n", url, err)
		return false
	}

	responseText := string(body)
	if strings.Contains(responseText, "Docker Engine") || strings.Contains(responseText, "GoVersion") || strings.Contains(responseText, "BuildTime") {
		fmt.Printf("[+] Potential -> %s\n", url)
		return true
	} else {
		fmt.Printf("[-] No vulnerability detected -> %s\n", url)
		return false
	}
}