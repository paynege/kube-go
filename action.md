1. IfAppInstalled(appName string) (bool, error)
2. IfFileExisted(file string) (bool, error)
3. ServiceStatus(serviceName string) (string, error)
4. Chmod(file string, arg string) error
5. Copy(rsc string, dist string) error
6. DaemonReload() error
7. IfDirectoryExisted(directory string) (bool, error)
8. Mkdir(directory string) error
9. Move(rsc string, dist string) error
10. ServiceRestart(serviceName string) error
11. ServiceEnable(serviceName string) error
12. ServiceStart(serviceName string) error