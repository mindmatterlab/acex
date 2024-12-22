// usercenter is the user center of the acex cloud platform.
package main

import (
	"github.com/mindmatterlab/acex/cmd/acex-usercenter/app"
)

func main() {
	// Creating a new instance of the usercenter application and running it
	app.NewApp("acex-usercenter").Run()
}
