package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"time"

	"github.com/xlzd/gotp"
)

const secretFile = ".otp.secret"

func main() {
	homePath, _ := os.UserHomeDir()
	data, err := ioutil.ReadFile(path.Join(homePath, secretFile))
	if err != nil {
		fmt.Fprintln(os.Stderr, "load secret file failed", err)
		os.Exit(1)
	}

	tt := int(time.Now().Unix() / 60)
	secret := strings.TrimSpace(string(data))
	hotp := gotp.NewDefaultHOTP(secret)
	fmt.Println(hotp.At(tt))
}
