# up0x

一款集合常见提权方法的提权工具

## 项目简介

up0x 是一个面向渗透测试人员与安全研究者的 Linux 提权工具，集合了多种常见提权检测与利用方法，涵盖了 SUID 配置错误、CVE 漏洞利用（如 Sudo、特定年份的高危漏洞）等功能。

## 功能特性

- **SUID 检测与利用**：自动扫描系统中的 SUID 文件，并尝试利用常见的提权方法（如 GTFOBins）进行提权。
- **Sudo 漏洞检测与利用**：自动检测系统 Sudo 版本，判断是否存在已知高危漏洞（如 CVE-2025-32463），并尝试进行漏洞利用。
- **CVE 漏洞集成**：内置多种 CVE 漏洞利用模块，便于一键检测和利用。

## 使用方法

### 编译

本项目使用 Go 语言开发，需先安装 Go 环境。

```bash
git clone https://github.com/g0ubu1i/up0x.git
cd up0x
go build -o up0x main.go
```

### 运行

```bash
./up0x
```

运行后，工具将自动检测可能存在的提权漏洞

### 输出示例

```
██    ██ ██████   ██████  ██   ██ 
██    ██ ██   ██ ██  ████  ██ ██   
██    ██ ██████  ██ ██ ██   ███   
██    ██ ██      ████  ██  ██ ██  
 ██████  ██       ██████  ██   ██ 
  author: g0ubu1i

[+] up0x start work
[+] check SUID ...
find SUID: /usr/bin/vim
try get root: /usr/bin/vim [-c :py import os; os.execl("/bin/sh", "sh", "-pc", "reset; exec sh -p")]
...
[+] check Sudo version
[+] Sudo version: 1.9.15
[+] CVE-2025-32463 vulnerability found
[+] Compiled malicious shared library.
[+] Created stage directory: /tmp/sudowoot.stage.xxxxx
```

## 支持的提权方式

- SUID 程序利用（覆盖常见 GTFOBins，如 python3, vim, find, awk, bash, cp, less, more, nano, ed, env, expect, ftp, gdb, git, lua, perl, php, ruby, scp, tar, timeout, rsync, zsh, chmod 等）
- 挂载 Docker Socket
- Sudo 漏洞（如 CVE-2025-32463 等）
- 可扩展更多 CVE 漏洞和自定义利用方式

## 免责声明

本工具仅供学习与安全测试使用，禁止用于非法用途。因使用本工具产生的任何后果与作者无关。

## 致谢

- [GTFOBins](https://gtfobins.github.io/)
- 安全社区与所有开源贡献者

---

> 作者：g0ubu1i  
> 仓库地址：https://github.com/g0ubu1i/up0x