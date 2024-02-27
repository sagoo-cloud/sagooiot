package version

import "fmt"

var (
	BuildTime    string
	BuildVersion string
	CommitID     string
)

func ShowLogo(buildVersion, buildTime, commitID string) {
	BuildVersion = buildVersion
	BuildTime = buildTime
	CommitID = commitID
	//版本号
	fmt.Println("   _____                         \n  / ____|                        \n | (___   __ _  __ _  ___   ___  \n  \\___ \\ / _` |/ _` |/ _ \\ / _ \\ \n  ____) | (_| | (_| | (_) | (_) |\n |_____/ \\__,_|\\__, |\\___/ \\___/ \n                __/ |            \n               |___/             ")
	fmt.Println("Version   ：", buildVersion)
	fmt.Println("BuildTime ：", buildTime)
	fmt.Println("CommitID  ：", commitID)
	fmt.Println("")
	fmt.Println("Copyright:", "Liaoning Sagoo Cloud Technology Co.,Ltd")
}
