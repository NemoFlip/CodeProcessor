package consumer

import (
	"HomeWork1/entity"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func RunCode(codeInfo entity.CodeRequest) []byte {
	var fileName string
	switch codeInfo.Translator {
	case "python3":
		fileName = "pyCode.py"
	case "clang":
		fileName = "cCode.c"
	case "c++":
		fileName = "gcc.cpp"
	}
	err := os.WriteFile(fileName, []byte(codeInfo.Code), 0777)
	if err != nil {
		fmt.Printf("failed to create the file: %s", err.Error())
		return nil
	}
	cmd := exec.Command("docker", "build", "-f", "rabbitmq/consumer/Dockerfile", "-t", "code-app", ".")
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
	return output
}
