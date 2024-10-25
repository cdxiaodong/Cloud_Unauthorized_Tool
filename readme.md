
### README_EN.md

---

# clod_Unauthorized_tool

请查看 [README-zh.md](https://github.com/cdxiaodong/Cloud_Unauthorized_Tool/blob/master/readme-zh.md) 获取中文文档。

## Introduction

`clod_Unauthorized_tool` is a tool designed to detect and exploit unauthorized access in various services. It supports the following functionalities:

1. Detecting and exploiting unauthorized access to the Docker API to create privileged containers.
2. Exploiting the Kubernetes Dashboard to create containers and execute reverse shell commands.
3. Retrieving tokens from etcd and using them to take control of Kubernetes clusters.
4. Exploiting unauthorized access to the Kubernetes API Server to take control of namespaces.

---

Difference with kube-hunter:
In simple terms, kube-hunter can only perform proof-of-concept (POC) attacks, while cloud Unauthorized_tool can execute exploits.
For example:

It supports creating privilege containers through unauthorized access to the Docker API.
It can create containers through the k8s dashboard and establish a reverse shell.
It can take over the k8s cluster by acquiring a token from etcd and using it.
It can exploit unauthorized access to the API server and take control of namespaces.

## Installation

### Requirements

- Go 1.16 or higher

### Steps

1. Clone the repository:
   ```bash
   git clone https://github.com/cdxiaodong/clod_Unauthorized_tool.git
   cd clod_Unauthorized_tool
   ```

2. Build the project:
   ```bash
   go build -o clod_Unauthorized_tool
   ```

3. Run the tool:
   ```bash
   ./clod_Unauthorized_tool
   ```

---

## Example Usage

![image](https://github.com/user-attachments/assets/9aff0673-23e7-4444-808f-ff5de94d0bcf)

![image](https://github.com/user-attachments/assets/ced975f5-403f-450b-918e-88b7cb95edba)


## Contributing

Feel free to submit issues and pull requests.

---

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---


