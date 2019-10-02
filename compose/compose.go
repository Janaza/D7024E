package main

import (
	"log"
	"os"
	"os/exec"
	"time"
)

func initSwarm() error {
	cmd1 := exec.Command("go", "build", "-o", "../main/main", "../main/main.go")
	cmd2 := exec.Command("docker", "swarm", "leave", "--force")
	cmd3 := exec.Command("docker", "swarm", "init")
	cmd4 := exec.Command("docker", "stack", "deploy", "-c", "docker-compose.yml", "d7024e")

	cmd1.Stdout = os.Stdout
	cmd1.Stderr = os.Stderr
	err := cmd1.Run()
	if err != nil {
		return err
	}
	time.Sleep(5 * time.Second)

	cmd2.Stdout = os.Stdout
	cmd2.Stderr = os.Stderr
	err = cmd3.Run()
	if err != nil {
		return err
	}

	cmd3.Stdout = os.Stdout
	cmd3.Stderr = os.Stderr
	err = cmd3.Run()
	if err != nil {
		return err
	}
	cmd4.Stdout = os.Stdout
	cmd4.Stderr = os.Stderr
	err = cmd4.Run()
	if err != nil {
		return err
	}
	return nil
}
func main() {
	err := initSwarm()

	if err != nil {
		log.Fatal(err)
	}

	/*time.Sleep(10 * time.Second)
	cmd3 := exec.Command("docker", "inspect", "d7024e_kademlia_network")
	out, err := cmd3.Output()
	if err != nil {
		log.Fatal(err)
	}
	file, err := os.Create("../nodes.JSON")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = io.WriteString(file, fmt.Sprintf(string(out)))
	if err != nil {
		log.Fatal(err)
	}*/
}

