# EasyBanner

[![Go Version](https://img.shields.io/github/go-mod/go-version/firework-a/easy-banner?style=flat-square)](https://github.com/firework-a/easy-banner)
[![Latest Release](https://img.shields.io/github/v/release/firework-a/easy-banner?style=flat-square)](https://github.com/firework-a/easy-banner/releases)
[![License](https://img.shields.io/github/license/firework-a/easy-banner?style=flat-square)](LICENSE)
[![Stars](https://img.shields.io/github/stars/firework-a/easy-banner?style=flat-square)](https://github.com/firework-a/easy-banner/stargazers)
[![Go Report](https://goreportcard.com/badge/github.com/firework-a/easy-banner?style=flat-square)](https://goreportcard.com/report/github.com/firework-a/easy-banner)

便捷的为 `Go` 应用添加控制台 `Banner`

终端打印 Banner，支持彩色Banner

## 特性

- **动态参数注入** - 支持版本号、构建时间等动态内容
- **终端颜色适配** - 自动检测终端色彩支持
- **轻量无依赖** - 纯标准库实现，无需第三方依赖

## 安装

```bash
# 获取最新版本
go get github.com/firework-a/easy-banner@latest

# 指定版本
go get github.com/firework-a/easy-banner@v1.0.0
```

## 快速开始

### 基础用法

#### 在项目根目录创建`banner.txt`文件

```go
package main

import "github.com/firework-a/easy-banner"

func main() {
	// 打印静态easy-banner
	_ = banner.Print(nil)
}
```

### 带动态参数

#### 在代码中：
```go
package main

import (
	"time"

	"github.com/firework-a/easy-banner"
)

func main() {
	err := banner.Print(map[string]string{
		"Version":   "1.0.0",
		"BuildTime": time.Now().Format("2006-01-02"),
		"Author":    "Your Name",
	})
	if err != nil {
		return
	}
}
```

#### 在txt模板中：
```text
  _________  ____  _______
/ ___/ __ \/ __ \/ ___/ /   Version: {{.Version}}
/ /__/ /_/ / / / / /__/ /   Build: {{.BuildTime}}
\___/\____/_/ /_/\___/_/    Author: {{.Author}}
```

#### 彩色banner
```text
{{.Bold}}{{.Cyan}}_________  ____  _______{{.Reset}}
{{.Bold}}{{.Cyan}} / ___/ __ \/ __ \/ ___/ /{{.Reset}}  {{.Green}}Version: {{.Yellow}}{{.Version}}{{.Reset}}
{{.Bold}}{{.Cyan}}/ /__/ /_/ / / / / /__/ /{{.Reset}}   {{.Green}}Build: {{.Yellow}}{{.BuildTime}}{{.Reset}}
{{.Bold}}{{.Cyan}}\___/\____/_/ /_/\___/_{{.Reset}}    {{.Green}}Author: {{.Yellow}}{{.Author}}{{.Reset}}
```

## 配置参数

### 环境变量

| 变量名        | 默认值    | 说明             |
|------------|--------|----------------|
| `NO_COLOR` | -      | 禁用所有颜色输出(遵循标准) |

### 配置

```go
// 在 main() 前配置
banner.Enabled = true  // 全局开关
banner.Color = false    // 禁用颜色
banner.Writer = os.Stderr // 输出到标准错误
```

## 贡献

欢迎提交 Issue 和 PR，请遵循以下流程：
1. Fork 仓库
2. 创建特性分支 (`git checkout -b feature/xxx`)
3. 提交更改 (`git commit -am 'Add some feature'`)
4. 推送到分支 (`git push origin feature/xxx`)
5. 创建 Pull Request

## 许可证

MIT License © [firework-a](https://github.com/firework-a)

## Star History

<a href="https://www.star-history.com/?repos=firework-a/easy-banner&type=date&legend=top-left">
 <picture>
   <source media="(prefers-color-scheme: dark)" srcset="https://api.star-history.com/image?repos=firework-a/easy-banner&type=date&theme=dark&legend=top-left" />
   <source media="(prefers-color-scheme: light)" srcset="https://api.star-history.com/image?repos=firework-a/easy-banner&type=date&legend=top-left" />
   <img alt="Star History Chart" src="https://api.star-history.com/image?repos=firework-a/easy-banner&type=date&legend=top-left" />
 </picture>
</a>
