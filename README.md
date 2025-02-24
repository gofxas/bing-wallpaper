# Bing 每日壁纸工具 🖼️

[![Go Version](https://img.shields.io/github/go-mod/go-version/username/bing-wallpaper)](https://golang.org/)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![Build Status](https://img.shields.io/github/actions/workflow/status/username/bing-wallpaper/build.yml)](https://github.com/username/bing-wallpaper/actions)

自动获取 Bing 每日壁纸并设置为 Windows 桌面的轻量级工具，开箱即用的单文件解决方案。

👉 [下载最新版](https://github.com/username/bing-wallpaper/releases/latest)

![](demo.gif)

## 功能特性 ✨

- 🕑 **每日自动更新** - 开机后自动获取当日壁纸
- 🖼️ **壁纸本地缓存** - 保存到 `wallpaper` 子目录（按日期命名）
- 🔄 **开机自启动** - 自动注册到系统启动项
- 📂 **便携模式** - 所有文件存储在当前目录
- 📝 **错误日志** - 自动生成 error.log
- 🧹 **自动清理** - 保留最近30天的壁纸

## 快速开始 🚀

### 安装方式

#### 方式一：编译安装（需要 Go 1.20+）
```bash
git clone https://github.com/username/bing-wallpaper.git
cd bing-wallpaper
go build -ldflags="-H windowsgui" -o bing-wallpaper.exe