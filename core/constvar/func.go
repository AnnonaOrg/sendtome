package constvar

import (
	"fmt"
	"time"
)

func APPName() string {
	return APP_NAME
}
func APPVersion() string {
	return "v" + APP_VERSION
}

func APPAbout() string {
	text := APPName() + " " + APPVersion() + "\nPowered by @umfaka"
	return text
}

func APPDesc() string {
	return fmt.Sprintf(
		"支持频道: @annonaOrg 交流群组: @annonaChat \nCopyright ©2018-%d annonaOrg Team. All Rights Reserved",
		time.Now().Year(),
	)
}
func APPDesc404() string {
	return fmt.Sprintf(
		"支持频道: @annonaOrg 交流群组: @annonaChat \nCopyright ©2018-%d annonaOrg Team. All Rights Reserved)(Error API route.)",
		time.Now().Year(),
	)
}

func APPDescEx() string {
	return fmt.Sprintf(
		"支持频道: @annonaOrg 交流群组: @annonaChat \nCopyright ©2018-%d annonaOrg Team. All Rights Reserved",
		APPVersion(), time.Now().Year(),
	)
}
