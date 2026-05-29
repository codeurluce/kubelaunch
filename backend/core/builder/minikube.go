package builder

import "os/exec"

func LoadImageToMinikube(image string) error {
	cmd := exec.Command(
		"minikube",
		"image",
		"load",
		image,
	)
	return cmd.Run()
}
