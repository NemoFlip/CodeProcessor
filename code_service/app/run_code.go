package app

import (
	"HomeWork1/internal/entity"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func RunCode(codeInfo entity.CodeRequest) []byte {
	var fileName string
	var dockerPath = "code_service/app/dockerfiles"
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
		log.Printf("failed to create the file: %s", err.Error())
		return nil
	}
	cmd := exec.Command("docker", "build", "-f", dockerPath, "-t", "code-http_server", "--force-rm", ".")
	if err = cmd.Run(); err != nil {
		log.Fatalf("Failed to build Docker image: %v", err.Error())
		return nil
	}

	cmd = exec.Command(
		"docker", "run", "--rm",
		"--memory=256m",
		"--ulimit", "cpu=1",
		"-e", "LANG="+codeInfo.Translator,
		"code-http_server",
	)

	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("failed to run the Docker container: %s\n", err.Error())
		return nil
	}
	log.Printf("Output:\n%s\n", string(output))
	return output
}
