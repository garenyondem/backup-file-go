# Backup File (Go)
Copies source file to destination directory and archives it with a date as a name.


## Usage
```bash
go build main.go
```
Then pass source and destination arguments respectively to binary file.

### macOS and Linux

```bash
./main /Users/garenyondem/desktop/preciousFile.txt /Users/BackupFolder
```
### Windows
```bash
main.exe "C:\Users\%username%\Desktop\preciousFile.txt" "E:\BackupFolder"
```
