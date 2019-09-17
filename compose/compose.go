package main

import (
	"log"
	"os"
	"os/exec"
)

func initSwarm() error {
	cmd := exec.Command("docker", "swarm", "init")
	cmd2 := exec.Command("docker", "stack", "deploy", "-c", "docker-compose.yml", "d7024e")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return err
	}
	cmd2.Stdout = os.Stdout
	cmd2.Stderr = os.Stderr
	err = cmd2.Run()
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
