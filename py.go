package main

import (
	"embed"
	"fmt"
	"log"
	"os"
	"os/exec"
)

var embeddedExe embed.FS

func main() {
	EnsurePythonInstalled()
	EnsurePythonModules()
	tmpFile, err := os.CreateTemp("", "script-*.exe")
	if err != nil {
		log.Fatalf("Error creando archivo temporal: %v", err)
	}
	defer os.Remove(tmpFile.Name())
	exeContent, err := embeddedExe.ReadFile("embedded_script")
	if err != nil {
		log.Fatalf("Error leyendo el ejecutable embebido: %v", err)
	}

	if _, err := tmpFile.Write(exeContent); err != nil {
		log.Fatalf("Error escribiendo en archivo temporal: %v", err)
	}
	tmpFile.Close()
	cmd := exec.Command(tmpFile.Name())
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatalf("Error ejecutando el .exe: %v", err)
	}

	fmt.Println("Ejecución completada correctamente.")
}

func EnsurePythonInstalled() {
	cmd := exec.Command("python", "--version")
	err := cmd.Run()
	if err != nil {
		log.Fatalf("Python no está instalado o no está en el PATH.")
	}
	fmt.Println("Python está instalado.")
}

func EnsurePythonModules() {
	modules := []string{
		"configparser", "os", "re", "threading", "shutil", "subprocess", "asyncio",
		"customtkinter", "tkinter", "Pillow", "telegram", "json", "uuid", "random",
		"string", "python-telegram-bot", "ctk",
	}

	for _, module := range modules {
		err := checkAndInstallModule(module)
		if err != nil {
			log.Fatalf("Error verificando el módulo '%s': %v", module, err)
		}
	}
}

func checkAndInstallModule(module string) error {
	cmd := exec.Command("python", "-c", fmt.Sprintf("import %s", module))
	err := cmd.Run()
	if err == nil {
		fmt.Printf("El módulo '%s' ya está instalado.\n", module)
		return nil
	}

	fmt.Printf("Instalando módulo '%s'...\n", module)
	cmd = exec.Command("python", "-m", "pip", "install", module)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error instalando módulo '%s': %s", module, string(output))
		return err
	}

	fmt.Printf("El módulo '%s' se instaló correctamente.\n", module)
	return nil
}
