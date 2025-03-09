package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const Csi = "\\x1b["

const SetBlue = Csi + "96m"
const SetMuted = Csi + "90m"
const SetNoStyle = Csi + "0m"

const ArrowIcon = "\\uf178 "
const FolderIcon = "\\uf07b "
const DriveIcon = "\\uf0a0 "

const HomeIcon = "\\uf015 "
const DevIcon = "\\ueeff "
const DocIcon = "\\uf016 "
const DownloadsIcon = "\\uf063 "
const DesktopIcon = "\\uf108 "
const UsersIcon = "\\uedca "
const LibraryIcon = "\\uf19c "
const ApplicationsIcon = "\\uf0be "

func setPrompt(prompt string) {
	escaped := strings.ReplaceAll(prompt, "'", "'\"'\"'")
	fmt.Printf("PS1=$(printf '%s')\n", escaped)
}

func iconForDir(dir string) string {
	homedir, err := os.UserHomeDir()

	if err != nil {
		setPrompt(ArrowIcon + " ")
		os.Exit(0)
	}

	if dir == homedir {
		return HomeIcon
	}

	if filepath.Base(dir) == "Developer" && filepath.Dir(dir) == homedir {
		return DevIcon
	}

	if filepath.Base(dir) == "Documents" && filepath.Dir(dir) == homedir {
		return DocIcon
	}

	if filepath.Base(dir) == "Downloads" && filepath.Dir(dir) == homedir {
		return DownloadsIcon
	}

	if filepath.Base(dir) == "Desktop" && filepath.Dir(dir) == homedir {
		return DesktopIcon
	}

	switch dir {
	case "/Users":
		return UsersIcon
	case "/Library":
		return LibraryIcon
	case "/Applications":
		return ApplicationsIcon
	case "/":
		return DriveIcon
	}

	return FolderIcon
}

func printSetPs1() {
	cwd, err := os.Getwd()

	if err != nil {
		setPrompt(ArrowIcon + " ")
		os.Exit(0)
	}

	homedir, err := os.UserHomeDir()

	if err != nil {
		setPrompt(ArrowIcon + " ")
		os.Exit(0)
	}

	dirname := filepath.Base(cwd)
	if cwd == homedir {
		dirname = "~"
	}

	setPrompt(SetBlue + iconForDir(cwd) + " " + dirname + SetNoStyle + " " + SetMuted + ArrowIcon + SetNoStyle + " ")
}

func printSetup() {
	executablePath, err := os.Executable()

	if err != nil {
		fmt.Println("echo 'error: could not determine grompt executable path'")
		os.Exit(0)
	}

	fmt.Printf("precmd () {\n")
	fmt.Printf("eval \"$(%s set-prompt)\"\n", executablePath)
	fmt.Printf("}\n")
}

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("grompt -> Ian's opinionated shell prompt.")
		fmt.Println("          See https://github.com/iwotastic/grompt for setup instructions.")
		fmt.Println("")
	}else if len(os.Args) == 2 && os.Args[1] == "setup-precmd" {
		printSetup()
	}else if len(os.Args) == 2 && os.Args[1] == "set-prompt" {
		printSetPs1()
	}else{
		fmt.Println("error: invalid command.")
		os.Exit(1)
	}
}
