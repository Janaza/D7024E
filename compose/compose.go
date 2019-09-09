package compose
import (
"os"
"os/exec"
"log")
func main() {
cmd := exec.Command("docker", "swarm", "init")
cmd2 := exec.Command("docker", "stack", "deploy", "-c", "docker-compose.yml", "d7024e")
cmd.Stdout = os.Stdout
cmd.Stderr = os.Stderr
err := cmd.Run()
if err != nil {
log.Fatal(err)
}
cmd2.Stdout = os.Stdout
cmd2.Stderr = os.Stderr
err = cmd2.Run()
if err != nil {
        log.Fatal(err)
}
}
