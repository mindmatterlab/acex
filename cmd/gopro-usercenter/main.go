// usercenter is the user center of the gopro cloud platform.
package main

import (
	// 在 go.uber.org/automaxprocs 包中，有一个 init 方法会自动执行。
	// 尝试从系统环境（如 cgroup 或 Linux 控制组的 CPU 限制）获取 CPU 配额，然后调整 GOMAXPROCS 的值。如果无法检测到 CPU 配额，就保持默认的 GOMAXPROCS（通常是 CPU 核心数）。
	// 使用场景：在容器环境中特别有用，比如 Kubernetes 中的容器可能被限制只能使用 2 个 CPU 核心，但是如果不设置 GOMAXPROCS，Go 程序会认为它可以使用主机所有的 CPU 核心。
	// 本机 Mac 开发环境通常不需要关心这个设置
	// Importing the package to automatically set GOMAXPROCS.
	_ "go.uber.org/automaxprocs/maxprocs"

	"github.com/mindmatterlab/go-pro/cmd/gopro-usercenter/app"
)

func main() {
	// Creating a new instance of the usercenter application and running it
	app.NewApp("gopro-usercenter").Run()
}
