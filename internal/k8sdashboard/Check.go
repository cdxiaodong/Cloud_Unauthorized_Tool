package k8sdashboard

import (
    "fmt"
    "io"
    "log"
    "net/http"
    "strings"
)

func Check(targetURL string) bool {
    apiURL := fmt.Sprintf("%s/api", targetURL)
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
        log.Println("可能存在未授权访问")
        return true
    } else {
        log.Println("完成/api路径请求，但是不存在未授权访问漏洞")
        return false
    }
}

func getDashboard(dashboardURL string) bool {
    resp, err := http.Get(dashboardURL)
    if err != nil {
        log.Printf("请求失败: %v", err)
        return false
    }
    defer resp.Body.Close()

    if resp.StatusCode == http.StatusOK {
        log.Println("成功访问k8s dashboard控制面板")
        return true
    }

    log.Printf("访问k8s dashboard控制面板失败，状态码: %d", resp.StatusCode)
    return false
}