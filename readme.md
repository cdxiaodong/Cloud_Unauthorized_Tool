
### README_EN.md

---

# clod_Unauthorized_tool

请查看 [README-zh.md](README-zh.md) 获取中文文档。

## Introduction

`clod_Unauthorized_tool` is a tool designed to detect and exploit unauthorized access in various services. It supports the following functionalities:

1. Detecting and exploiting unauthorized access to the Docker API to create privileged containers.
2. Exploiting the Kubernetes Dashboard to create containers and execute reverse shell commands.
3. Retrieving tokens from etcd and using them to take control of Kubernetes clusters.
4. Exploiting unauthorized access to the Kubernetes API Server to take control of namespaces.

---

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




## Contributing

Feel free to submit issues and pull requests.

---

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---


