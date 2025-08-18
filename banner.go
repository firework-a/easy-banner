// Package banner provides Spring Boot-like banner printing
package banner

import (
	"embed"
	"fmt"
	"io"
	"os"
	"text/template"
	"time"
)

//go:embed assets/default_banner.txt
var defaultBanner embed.FS

const (
	defaultBannerFile = "assets/default_banner.txt"
	customBannerFile  = "banner.txt"
)

// 包级配置
var (
	Enabled           = true // 全局开关
	Output  io.Writer = os.Stdout
	Color             = true // 启用彩色输出
)

// Print 打印Banner
// 如果传入data不为nil，则渲染动态banner；否则渲染静态banner
func Print(data map[string]string) error {
	if !Enabled {
		return nil
	}

	bannerContent, err := getBannerContent()
	if err != nil {
		return err
	}

	// 如果传入了数据，则作为模板渲染
	if data != nil {
		tmpl, err := template.New("banner").Parse(string(bannerContent))
		if err != nil {
			return fmt.Errorf("template parse error: %w", err)
		}
		return tmpl.Execute(Output, prepareData(data))
	}

	// 静态banner直接输出
	_, err = fmt.Fprintln(Output, string(bannerContent))
	return err
}

// SetConfig 设置配置参数
func SetConfig(cfg Config) {
	if cfg.Enabled != nil {
		Enabled = *cfg.Enabled
	}
	if cfg.Color != nil {
		Color = *cfg.Color
	}
	if cfg.Output != nil {
		Output = cfg.Output
	}
}

// Config 配置结构体
type Config struct {
	Enabled *bool
	Color   *bool
	Output  io.Writer
}

// getBannerContent 获取banner内容
func getBannerContent() ([]byte, error) {
	// 优先尝试加载用户自定义banner
	if customBanner, err := os.ReadFile(customBannerFile); err == nil {
		return customBanner, nil
	}
	// 回退到默认banner
	return defaultBanner.ReadFile(defaultBannerFile)
}

// prepareData 准备模板数据
func prepareData(data map[string]string) map[string]interface{} {
	result := make(map[string]interface{})
	for k, v := range data {
		result[k] = v
	}

	// 设置默认值
	setDefaultValue(result, "Version", "0.0.0")
	setDefaultValue(result, "BuildTime", time.Now().Format("2006-01-02 15:04:05"))
	setDefaultValue(result, "Author", "Unknown")

	// 处理颜色
	setColorCodes(result)
	return result
}

func setDefaultValue(m map[string]interface{}, key, defaultValue string) {
	if _, ok := m[key]; !ok {
		m[key] = defaultValue
	}
}

func setColorCodes(m map[string]interface{}) {
	useColor := Color && os.Getenv("NO_COLOR") == ""

	codes := map[string]string{
		"Reset":   "",
		"Bold":    "",
		"Red":     "",
		"Green":   "",
		"Yellow":  "",
		"Blue":    "",
		"Magenta": "",
		"Cyan":    "",
	}

	if useColor {
		codes = map[string]string{
			"Reset":   "\x1b[0m",
			"Bold":    "\x1b[1m",
			"Red":     "\x1b[31m",
			"Green":   "\x1b[32m",
			"Yellow":  "\x1b[33m",
			"Blue":    "\x1b[34m",
			"Magenta": "\x1b[35m",
			"Cyan":    "\x1b[36m",
		}
	}

	for k, v := range codes {
		m[k] = v
	}
}
