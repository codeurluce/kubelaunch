package analyzer

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/codeurluce/kubelaunch/backend/services"
)

func DetectNodePort(repoPath string) int {

	candidates := []string{
		"app.js",
		"server.js",
		"index.js",
		"main.js",
		"src/main.ts",
		"src/main.js",
	}

	for _, file := range candidates {

		fmt.Println("TRYING:", file)

		content, err := services.FetchFileContent(
			repoPath,
			file,
		)

		if err != nil {
			fmt.Println("Error:", err)
			fmt.Println("File not found:", file)
			continue
		}

		fmt.Println("FILE FOUND:", file)
		fmt.Println(content)

		// app.listen(3001)
		re := regexp.MustCompile(`listen\((\d+)`)

		match := re.FindStringSubmatch(content)

		if len(match) >= 2 {

			port, err := strconv.Atoi(match[1])

			if err == nil {

				println("=================================")
				println("PORT DETECTED:", port)
				println("FILE:", file)
				println("=================================")

				return port
			}
		}

		// const PORT = 3001
		rePort := regexp.MustCompile(`PORT\s*=\s*(\d+)`)

		portMatch := rePort.FindStringSubmatch(content)

		if len(portMatch) >= 2 {

			port, err := strconv.Atoi(portMatch[1])

			if err == nil {
				println("=================================")
				println("PORT DETECTED:", port)
				println("FILE:", file)
				println("=================================")
				return port
			}
		}

		reEnvPort := regexp.MustCompile(`\|\|\s*(\d+)`)
		envPortMatch := reEnvPort.FindStringSubmatch(content)

		if len(envPortMatch) >= 2 {

			port, err := strconv.Atoi(envPortMatch[1])

			if err == nil {
				println("=================================")
				println("PORT DETECTED:", port)
				println("FILE:", file)
				println("=================================")
				return port
			}
		}
	}
	println("NODE PORT DETECTOR RUNNING")
	println("REPO:", repoPath)
	println("FALLBACK PORT 3000")

	// fallback Node.js
	return 3000
}
