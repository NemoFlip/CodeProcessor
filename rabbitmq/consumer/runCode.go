package consumer

import (
	"HomeWork1/entity"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func RunCode(codeInfo entity.CodeRequest) []byte {
	var fileName string
	var dockerPath = "rabbitmq/consumer"
	//var imageName string
	switch codeInfo.Translator {
	case "python3":
		fileName = "pyCode.py"
		dockerPath = filepath.Join(dockerPath, "pyDockerfile")
		//imageName =
	case "clang":
		fileName = "cCode.c"
		dockerPath = filepath.Join(dockerPath, "cDockerfile")
	case "c++":
		fileName = "gcc.cpp"
		dockerPath = filepath.Join(dockerPath, "cppDockerfile")
	}
	err := os.WriteFile(fileName, []byte(codeInfo.Code), 0777)
	if err != nil {
		fmt.Printf("failed to create the file: %s", err.Error())
		return nil
	}
	cmd := exec.Command("docker", "build", "-f", dockerPath, "-t", "code-app", ".")
	if err = cmd.Run(); err != nil {
		log.Fatalf("Failed to build Docker image: %v", err.Error())
		return nil
	}

	cmd = exec.Command("docker", "run", "--rm", "-e", "LANG="+codeInfo.Translator, "code-app")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("failed to run the Docker container: %s\n", err.Error())
		return nil
	}
	// Вывод результата
	fmt.Printf("Output:\n%s\n", string(output))
	exec.Command("docker", "rmi", "code-app")
	return output
}
