### README_CN.md

---

# clod_Unauthorized_tool

## 简介

`clod_Unauthorized_tool` 是一个用于检测和利用各种服务中未经授权访问的工具。它支持以下功能：

1. 检测并利用 Docker API 的未授权访问来创建特权容器。
2. 利用 Kubernetes Dashboard 创建容器并执行反弹 shell 命令。
3. 从 etcd 获取 token 并使用 token 接管 Kubernetes 集群。
4. 利用 Kubernetes API Server 的未授权访问来接管命名空间。

---

## 安装

### 要求

- Go 1.16 或更高版本

### 安装步骤

1. 克隆仓库：
   ```bash
   git clone https://github.com/cdxiaodong/clod_Unauthorized_tool.git
   cd clod_Unauthorized_tool
   ```

2. 构建项目：
   ```bash
   go build -o clod_Unauthorized_tool
   ```

3. 运行工具：
   ```bash
   ./clod_Unauthorized_tool
   ```

---

## 使用示例
![image](https://github.com/user-attachments/assets/9aff0673-23e7-4444-808f-ff5de94d0bcf)

![image](https://github.com/user-attachments/assets/ced975f5-403f-450b-918e-88b7cb95edba)


## 贡献

欢迎提交问题和拉取请求。

---

## 许可证

本项目采用 MIT 许可证。详情请参阅 [LICENSE](LICENSE) 文件。

---
