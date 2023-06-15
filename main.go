package main

import (
	b "Auth/Auth"
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

var clear map[string]func()

//* KeyAuth Application Details *//
var name = ""
var ownerid = ""
var version = "1.0"

//* API SET UP VALUES ^^^^ *//
var username = "" //* Keep Clear
var password = "" //* Keep Clear
var key = ""      //* Keep Clear

func main() {
	b.Api(name, ownerid, version) // Important to set up the API Details

	reader := bufio.NewReader(os.Stdin)

	ClearConsole()

	fmt.Println("\n\n Connecting..")
	b.Init()

	fmt.Println("\n App Data:")
	fmt.Println(" Number of users:", b.NumUsers)
	fmt.Println(" Number of online users:", b.NumOnlineUsers)
	fmt.Println(" Number of keys:", b.NumKeys)
	fmt.Println(" Application Version:", version)
	fmt.Println(" Customer Panel Link:", b.CustomerPanelLink)

	time.Sleep(1 * time.Second) //* Rate Limit Wait
	fmt.Println("\n Current Session Validation Status:", b.Check())
	time.Sleep(1 * time.Second) //* Rate Limit Wait
	fmt.Println(" Blacklisted?:", b.CheckBlack())
	time.Sleep(1 * time.Second) //* Rate Limit Wait

	fmt.Println("\n [1] Login\n [2] Register\n [3] Upgrade\n [4] License key only\n\n Choose option: ")

	char, _, err := reader.ReadRune()

	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	switch char {
	case '1':
		reader := bufio.NewReader(os.Stdin)

		fmt.Println("\n\n Enter Username: ")
		username, _ := reader.ReadString('\n')
		username = strings.TrimSuffix(username, "\n")

		fmt.Println("\n\n Enter Password: ")
		password, _ := reader.ReadString('\n')
		password = strings.TrimSuffix(password, "\n")

		b.Login(username, password)

		break
	case '2':
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("\n\n Enter Username: ")
		username, _ := reader.ReadString('\n')
		username = strings.TrimSuffix(username, "\n")

		fmt.Println("\n\n Enter Password: ")
		password, _ := reader.ReadString('\n')
		password = strings.TrimSuffix(password, "\n")

		fmt.Println("\n\n Enter License: ")
		key, _ := reader.ReadString('\n')
		key = strings.TrimSuffix(key, "\n")

		b.Register(username, password, key)

		break
	case '3':
		reader := bufio.NewReader(os.Stdin)

		fmt.Println("\n\n Enter Username: ")
		username, _ := reader.ReadString('\n')
		username = strings.TrimSuffix(username, "\n")

		fmt.Println("\n\n Enter License: ")
		key, _ := reader.ReadString('\n')
		key = strings.TrimSuffix(key, "\n")

		b.Upgrade(username, key)

		break
	case '4':
		reader := bufio.NewReader(os.Stdin)

		fmt.Println("\n\n Enter License: ")
		key, _ := reader.ReadString('\n')
		key = strings.TrimSuffix(key, "\n")

		b.License(key)

		break
	default:
		fmt.Println("\n\n Invalid Selection")
		os.Exit(0)
		break
	}

	fmt.Println("\n User data:")
	fmt.Println(" Username:", b.Username)
	fmt.Println(" IP address:", b.Ip)
	fmt.Println(" Hardware-Id:", b.Hwid)
	fmt.Println(" Created at:", b.Createdate)
	fmt.Println(" Last login at:", b.Lastlogin)
	fmt.Println(" Subscription:", b.Subscription)

	fmt.Println("\n Current Session Validation Status:", b.Check())

	/* --> Extra Functions <--
	* User Variables *
	b.SetVar("VariableName", "VariableData") // Set up User Variable
	b.GetVar("VariableName") // Get User Variable

	* Get Public Variables * - https://keyauth.cc/dashboard/app/variables/

	b.Var("VariableName") // Get Public Variable

	Example:
	var publicVariable = b.Var("VariableName")
	fmt.Println("Variable Content: " + publicVariable)

	* Webhooks * - https://keyauth.cc/dashboard/app/webhooks/
	b.Webhook("WebhookName", "WebhookData") // Send Webhook

	Example:

	var WbData = b.Webhook("Webhook ID", "?type=test")
	fmt.Println("Webhook Data: " + WbData)

	* Logs * - https://keyauth.cc/dashboard/app/settings/
	b.Log("Message") // Send Log to Webhook of your choice ^^
	*/

	fmt.Println("\n Closing in ten seconds...")
	time.Sleep(10 * time.Second)
	os.Exit(0)
}

func init() {

	clear = make(map[string]func())
	clear["linux"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func ClearConsole() {
	value, ok := clear[runtime.GOOS]
	if ok {
		value()
	} else {
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}
