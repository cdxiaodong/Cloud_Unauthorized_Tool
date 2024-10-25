README.md
# 🌩️ Cloud_Unauthorized_tool
**English** | [简体中文](https://github.com/cdxiaodong/Cloud_Unauthorized_Tool/blob/master/readme-zh.md)
## Introduction
`Cloud_Unauthorized_tool` is a powerful utility crafted to identify and leverage unauthorized access across various services. It equips security enthusiasts and professionals with the means to detect and exploit vulnerabilities in cloud environments. The tool stands out with its support for multiple exploitation techniques, setting it apart from conventional vulnerability scanners.
## Key Features
- **Docker API Exploitation**: Detect and exploit unauthorized access to the Docker API, creating privileged containers.
- **Kubernetes Dashboard Exploits**: Exploit the Kubernetes Dashboard to spawn containers and execute reverse shell commands.
- **etcd Token Retrieval**: Extract tokens from etcd, facilitating control over Kubernetes clusters.
- **Kubernetes API Server Exploits**: Exploit unauthorized access to the Kubernetes API Server, enabling control over namespaces.
## Key Differentiator
Compared to kube-hunter, `Cloud_Unauthorized_tool` doesn't just stop at proof-of-concept (POC) attacks. It empowers users to execute real exploits, making it a formidable tool in the hands of security experts.
## Installation
### Requirements
- Go 1.16 or higher
### Steps
1. Clone the repository:
   ```bash
   git clone https://github.com/cdxiaodong/Cloud_Unauthorized_tool.git
   cd Cloud_Unauthorized_tool
   ```
2. Build the project:
   ```bash
   go build -o Cloud_Unauthorized_tool
   ```
3. Run the tool:
   ```bash
   ./Cloud_Unauthorized_tool
   ```
## Example Usage
![Example Usage](https://private-user-images.githubusercontent.com/84082748/379673789-9aff0673-23e7-4444-808f-ff5de94d0bcf.png?jwt=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJnaXRodWIuY29tIiwiYXVkIjoicmF3LmdpdGh1YnVzZXJjb250ZW50LmNvbSIsImtleSI6ImtleTUiLCJleHAiOjE3Mjk4MzI1MDcsIm5iZiI6MTcyOTgzMjIwNywicGF0aCI6Ii84NDA4Mjc0OC8zNzk2NzM3ODktOWFmZjA2NzMtMjNlNy00NDQ0LTgwOGYtZmY1ZGU5NGQwYmNmLnBuZz9YLUFtei1BbGdvcml0aG09QVdTNC1ITUFDLVNIQTI1NiZYLUFtei1DcmVkZW50aWFsPUFLSUFWQ09EWUxTQTUzUFFLNFpBJTJGMjAyNDEwMjUlMkZ1cy1lYXN0LTElMkZzMyUyRmF3czRfcmVxdWVzdCZYLUFtei1EYXRlPTIwMjQxMDI1VDA0NTY0N1omWC1BbXotRXhwaXJlcz0zMDAmWC1BbXotU2lnbmF0dXJlPTljZWE0OTVkNzZhOWJmNmMyODI5MDUxZTUyN2NmYTgzMWIwMDE2ZDg0OGZkMGQ2NTY0MGE3ZmE3OTc4OWY5ZDImWC1BbXotU2lnbmVkSGVhZGVycz1ob3N0In0.9ZtCawT3xa6xEJctzWhZrg56bj-vsSWj99_MGEGsb_A)
![Example Output](https://private-user-images.githubusercontent.com/84082748/379673843-ced975f5-403f-450b-918e-88b7cb95edba.png?jwt=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJnaXRodWIuY29tIiwiYXVkIjoicmF3LmdpdGh1YnVzZXJjb250ZW50LmNvbSIsImtleSI6ImtleTUiLCJleHAiOjE3Mjk4MzI1MDcsIm5iZiI6MTcyOTgzMjIwNywicGF0aCI6Ii84NDA4Mjc0OC8zNzk2NzM4NDMtY2VkOTc1ZjUtNDAzZi00NTBiLTkxOGUtODhiN2NiOTVlZGJhLnBuZz9YLUFtei1BbGdvcml0aG09QVdTNC1ITUFDLVNIQTI1NiZYLUFtei1DcmVkZW50aWFsPUFLSUFWQ09EWUxTQTUzUFFLNFpBJTJGMjAyNDEwMjUlMkZ1cy1lYXN0LTElMkZzMyUyRmF3czRfcmVxdWVzdCZYLUFtei1EYXRlPTIwMjQxMDI1VDA0NTY0N1omWC1BbXotRXhwaXJlcz0zMDAmWC1BbXotU2lnbmF0dXJlPTQ3MTQ5MWZhMjA3MmIxNmQ5MDQ5YjJmOTE2NjRiMWIzMDI5ODY5OTFiN2M3YTQxNjNlZDI3NDZiNWYyOTg3OWMmWC1BbXotU2lnbmVkSGVhZGVycz1ob3N0In0.tatmtcS2VNcu6YKdr5D__2UorY9paN7PNpQndXUaEnQ)
## Contributing
Contributions are welcome! Feel free to report issues or submit pull requests to enhance the tool's capabilities.
## License
This project is licensed under the MIT License. For more details, refer to the [LICENSE](LICENSE) file.
---
🗨️ Contacts:
Email: cdxiaodong@systemshell.org
Website: cdxiaodong.life
