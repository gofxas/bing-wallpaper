package main

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"syscall"
	"time"
	"unsafe"
	"golang.org/x/sys/windows/registry"
)

const (
	bingAPI      = "https://www.bing.com/HPImageArchive.aspx?format=js&idx=0&n=1"
	spiSetWallpaper = 0x0014
	spifUpdateIniFile = 0x01
)

var (
	user32           = syscall.NewLazyDLL("user32.dll")
	systemParameters = user32.NewProc("SystemParametersInfoW")
)

func main() {
	// 获取图片URL
	url := getBingURL()
	if url == "" {
		return
	}

	// 下载图片
	imgPath := downloadImage(url)
	if imgPath == "" {
		return
	}

	// 设置壁纸
	setWallpaper(imgPath)

	// 设置开机启动
	setAutoStart()
}

func getBingURL() string {
	resp, err := http.Get(bingAPI)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()

	var result struct {
		Images []struct {
			URL string `json:"url"`
		} `json:"images"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return ""
	}

	if len(result.Images) == 0 {
		return ""
	}

	return "https://www.bing.com" + result.Images[0].URL
}

func downloadImage(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()

	// 获取可执行文件所在目录
	exePath, err := os.Executable()
	if err != nil {
		return ""
	}
	appDir := filepath.Dir(exePath)
	
	// 创建 wallpaper 子目录
	saveDir := filepath.Join(appDir, "wallpaper")
	if err := os.MkdirAll(saveDir, 0755); err != nil {
		return ""
	}

	// 生成带日期的文件名
	filename := time.Now().Format("20060102") + ".jpg"
	filePath := filepath.Join(saveDir, filename)

	// 保存文件
	out, err := os.Create(filePath)
	if err != nil {
		return ""
	}
	defer out.Close()

	if _, err := io.Copy(out, resp.Body); err != nil {
		return ""
	}

	return filePath
}

func setWallpaper(imgPath string) {
	pathUTF16, _ := syscall.UTF16PtrFromString(imgPath)
	systemParameters.Call(
		uintptr(spiSetWallpaper),
		uintptr(0),
		uintptr(unsafe.Pointer(pathUTF16)),
		uintptr(spifUpdateIniFile),
	)
}

func setAutoStart() {
	exePath, _ := os.Executable()
	key, _, _ := registry.CreateKey(
		registry.CURRENT_USER,
		`Software\Microsoft\Windows\CurrentVersion\Run`,
		registry.ALL_ACCESS,
	)
	defer key.Close()

	key.SetStringValue("BingWallpaper", exePath)
}