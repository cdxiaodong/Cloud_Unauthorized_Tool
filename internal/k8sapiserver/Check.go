package k8sapiserver

import (
    "fmt"
    "io"
    "log"
    "net/http"
    "strings"
)

func Check(apiaddr string) bool {
    apiURL := fmt.Sprintf("%s/api", apiaddr)
    log.Println(apiURL)

    resp, err := http.Get(apiURL)
    if err != nil {
        log.Printf("地址访问失败: %v", err)
        return false
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        log.Printf("读取响应失败: %v", err)
        return false
    }

    if strings.Contains(string(body), "/api") {
        log.Println("congrats, api-server request success.")
        return true
    } else {
        log.Println("完成apiserver请求，但是没有权限访问apiserver")
        return false
    }
}