package app

import (
	"HomeWork1/internal/entity"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func RunCode(codeInfo entity.CodeRequest) []byte {
	var fileName string
	var dockerPath = "code_service/app"
	switch codeInfo.Translator {
	case "python3":
		fileName = "pyCode.py"
		dockerPath = filepath.Join(dockerPath, "pyDockerfile")
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
	cmd := exec.Command("docker", "build", "-f", dockerPath, "-t", "code-app", "--force-rm", ".")
	if err = cmd.Run(); err != nil {
		log.Fatalf("Failed to build Docker image: %v", err.Error())
		return nil
	}

	cmd = exec.Command(
		"docker", "run", "--rm",
		"--memory=256m",
		"--ulimit", "cpu=1",
		"-e", "LANG="+codeInfo.Translator,
		"code-app",
	)

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("failed to run the Docker container: %s\n", err.Error())
		return nil
	}
	fmt.Printf("Output:\n%s\n", string(output))
	return output
}
