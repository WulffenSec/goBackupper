package main

import (
	"fmt"
	"github.com/fatih/color"
	"os"
	"os/exec"
	"strings"
)

const BannerHi string = `           ___          _
  __ _ ___| _ ) __ _ __| |___  _ _ __ _ __  ___ _ _
 / _' / _ \ _ \/ _' / _| / / || | '_ \ '_ \/ -_) '_|
 \__, \___/___/\__,_\__|_\_\\_,_| .__/ .__/\___|_|
 |___/                          |_|  |_|`

const BannerLo string = "                Github: WulffenSec | Version: 1.2"

func main() {
    // Check Args
	if len(os.Args) < 3 || os.Args[1] == "-h" || os.Args[1] == "--help" {
		fmt.Println("Backups recursively the source directory and its subdirectories to the target directory.")
		fmt.Println("Usage:\n$", os.Args[0], "path-to-source-directory path-to-target-directory -slient")
		os.Exit(1)
	}
    
    // No Banner
    silent := false
    if len(os.Args) > 3 {
        if os.Args[3] == "-silent" {
                silent = true
            }
        }

    // Banner
    if silent == false {
        red := color.New(color.FgRed).SprintFunc()
        yellow := color.New(color.FgYellow).SprintFunc()
        fmt.Printf("%s\n%s\n", red(BannerHi), yellow(BannerLo))
        fmt.Printf("Running %s, this may %s.\n", yellow("diff"), yellow("take a while"))
    }
	for true {
		// Loops until no differences are found.
		results := makeBackup()
		if results == true {
			break
		}
	}
}

func makeBackup() bool {
    // Colors
	red := color.New(color.FgRed).SprintFunc()
	blue := color.New(color.FgBlue).SprintFunc()
	cyan := color.New(color.FgCyan).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()

	sourceDir := os.Args[1]
	targetDir := os.Args[2]
    _, err := os.Open(targetDir)
    if err != nil {
        os.Mkdir(targetDir, 0755)
    }

	diff, err := exec.Command("diff", "-qr", sourceDir, targetDir).CombinedOutput()
	if err == nil {
		fmt.Printf("%s No differences found.\n", green("[O]"))
		return true
	} else if err.Error() == "exit status 2" {
		fmt.Println(strings.Split(string(diff), "\n")[0])
		os.Exit(2)
	}
	for _, line := range strings.Split(string(diff), "\n") {
		if line == "" {
			continue
		}
		if strings.HasPrefix(line, "Only in "+sourceDir) {
			// Backup files present only in source.
			sourceFile := strings.Split(line, "Only in ")[1]
			sourceFile = strings.Replace(sourceFile, "/: ", "/", -1)
			sourceFile = strings.Replace(sourceFile, ": ", "/", -1)
			targetFile := strings.Split(line, "Only in ")[1]
			targetFile = strings.Replace(targetFile, sourceDir, targetDir, -1)
			targetFile = strings.Replace(targetFile, "/: ", "/", -1)
			targetFile = strings.Replace(targetFile, ": ", "/", -1)
			file := strings.Split(line, "Only in ")[1]
			file = strings.Split(file, ": ")[1]
			fmt.Printf("%s Writing %s to backup location.\n", cyan("[!]"), cyan("\""+file+"\""))
			cmd := exec.Command("cp", "-r", sourceFile, targetFile)
			err := cmd.Run()
			if err != nil {
				fmt.Printf("Error: %s\n", err)
				os.Exit(1)
			}
		} else if strings.HasPrefix(line, "Only in "+targetDir) {
			// Delete files persent only in target.
			targetFile := strings.Split(line, "Only in ")[1]
			targetFile = strings.Replace(targetFile, "/: ", "/", -1)
			targetFile = strings.Replace(targetFile, ": ", "/", -1)
			file := strings.Split(line, "Only in ")[1]
			file = strings.Split(file, ": ")[1]
			fmt.Printf("%s Deleting %s from backup location.\n", red("[X]"), red("\""+file+"\""))
			cmd := exec.Command("rm", "-rf", targetFile)
			err := cmd.Run()
			if err != nil {
				fmt.Printf("Error: %s\n", err)
				os.Exit(1)
			}

		} else if strings.HasPrefix(line, "Files ") {
			// Backup files present in both source and target.
			bothFiles := strings.Split(line, " and ")
			sourceFile := strings.Split(bothFiles[0], "Files ")[1]
			targetFile := strings.Split(bothFiles[1], " differ")[0]
			file := strings.Split(bothFiles[0], "Files ")[1]
			file = file[strings.LastIndex(file, "/")+1:]
			fmt.Printf("%s Rewriting %s to backup location.\n", blue("[!]"), blue("\""+file+"\""))
			cmd := exec.Command("cp", "-r", sourceFile, targetFile)
			err := cmd.Run()
			if err != nil {
				fmt.Printf("Error: %s\n", err)
				os.Exit(1)
			}
		}
	}
	return false
}
