package cli

import (
	"fmt"
	"os/exec"
	"runtime"
)

func OpenBrowser(url string) (error) {
	var browserCommand *exec.Cmd
	switch runtime.GOOS {
	case "linux":
		browserCommand = exec.Command("xdg-open", url)
	case "windows":
		browserCommand = exec.Command("rundll32", "url.dll,FileProtocolHandler", url)
	case "darwin":
		browserCommand = exec.Command("open", url)
	default:
		return fmt.Errorf("unsupported operating system: %v", runtime.GOOS)
	}
	err := browserCommand.Run()
	return err
}
