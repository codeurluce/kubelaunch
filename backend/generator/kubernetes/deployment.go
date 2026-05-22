package kubernetes

import (
	"fmt"

	"github.com/codeurluce/kubelaunch/backend/models"
)

func GenerateDeployment(spec models.DeploymentSpec) string {

	return fmt.Sprintf(`
apiVersion: apps/v1
kind: Deployment
metadata:
  name: %s
spec:
  replicas: 1
  selector:
    matchLabels:
      app: %s
  template:
    metadata:
      labels:
        app: %s
    spec:
      containers:
        - name: %s
          image: %s:latest
          ports:
            - containerPort: %d
`,
		spec.Name,
		spec.Name,
		spec.Name,
		spec.Name,
		spec.Image,
		spec.Port,
	)
}
