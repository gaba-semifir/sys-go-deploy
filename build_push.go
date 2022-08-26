package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

// Typage de notre instance Powershell
type PowerShell struct {
	powerShell string
}

// Fonction main qui comprend principalement l'affichage et l'appel à nos fonctions
func main() {
	fmt.Print("Bienvenue dans GoBuilderPusherToRegistry \n")
	fmt.Println("---")
	fmt.Println("")
	fmt.Println("---")
	fmt.Print("Veuillez indiquer le chemin du dossier contenant les dockerfiles\n")
	var path_dockerfile string
	fmt.Scanln(&path_dockerfile)
	fmt.Print("le path que vous avez indiqué est :\n\t")
	fmt.Printf("%s\n", path_dockerfile)
	fmt.Printf("Scanning path...\n")
	fmt.Println("---")
	fmt.Println("")
	fmt.Println("---")
	time.Sleep(2 * time.Second)
	fmt.Println("Les fichiers trouvés sont :")
	verification_file(path_dockerfile)
	fmt.Println("---")
	fmt.Println("")
	fmt.Println("---")
	fmt.Println("Veuillez entrer le chemin du registry")
	fmt.Println("Format : <nom>:<port>")
	var path_registry string
	fmt.Scanln(&path_registry)
	path_recursion(path_dockerfile, path_registry)
	fmt.Println("---")
	fmt.Println("")
	fmt.Println("---")

}

// La fonction nous permet de verifier que le chemin utilisateur est bon
func verification_file(path string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println("path ERROR")
		os.Exit(2)
	}
	for _, file := range files {
		fmt.Printf("Fichier " + file.Name() + " Dossier: " + strconv.FormatBool(file.IsDir()) + "\n")
	}
}

// la fonction nous permet de lister le chemin utilisateur et d'executer les commandes build push
func path_recursion(path_dockerfile string, path_registry string) {
	err := filepath.Walk(path_dockerfile, func(path_dockerfile string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if strings.HasSuffix(path_dockerfile, "Dockerfile") {
			fmt.Println("Début du build...")
			time.Sleep(2 * time.Second)
			fmt.Println("---")
			fmt.Println("")
			fmt.Println("---")
			fmt.Println(path_dockerfile)
			tag_docker := build(path_dockerfile, path_registry)
			fmt.Println("Début du push...")
			time.Sleep(2 * time.Second)
			fmt.Println("---")
			fmt.Println("")
			fmt.Println("---")
			fmt.Println(path_dockerfile)
			push(path_registry, tag_docker)
		}
		return nil
	})
	if err != nil {
		log.Println(err)
	}
}

// la fonction nous permet de build les images, et de recuperer le tag pour le push de l'image
func build(path_dockerfile string, path_registry string) string {
	var docker_build string
	path_split := strings.Split(path_dockerfile, "\\")
	tag := path_registry + "/" + path_split[len(path_split)-2]
	docker_build = "docker build -f " + path_dockerfile + " -t " + tag + " ."
	fmt.Print(docker_build)
	posh := getPowershellSession()
	stdOut, stdErr, err := posh.execute(docker_build)
	fmt.Println(stdOut)
	fmt.Println(stdErr)
	fmt.Println(err)
	return tag
}

// la fonction push nous permet de pousser l'image sur le registry rentré par l'utlisateur
func push(path_registry string, tag string) {
	var docker_push string
	docker_push = "docker push " + tag
	fmt.Printf(docker_push)
	posh := getPowershellSession()
	stdOut, stdErr, err := posh.execute(docker_push)
	fmt.Println(stdOut)
	fmt.Println(stdErr)
	fmt.Println(err)
}

// La fonction nous permet de recuperer une instance de Powershell
func getPowershellSession() *PowerShell {
	ps, _ := exec.LookPath("powershell.exe")
	return &PowerShell{
		powerShell: ps,
	}
}

// Méthode pour executer la commande dans notre instance Powershell
func (p *PowerShell) execute(args ...string) (stdOut string, stdErr string, err error) {
	args = append([]string{"-NonInteractive"}, args...)
	cmd := exec.Command(p.powerShell, args...)
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err = cmd.Run()
	stdOut, stdErr = stdout.String(), stderr.String()
	return
}
