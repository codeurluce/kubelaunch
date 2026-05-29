package kubernetes

import (
	"fmt"
	"strings"

	"github.com/codeurluce/kubelaunch/backend/models"
)

func GenerateDeployment(spec models.DeploymentSpec) string {

	envVars := ""

	for key, value := range spec.EnvVars {
		envVars += fmt.Sprintf(`
            - name: %s
              value: "%s"`,
			key,
			value,
		)
	}

	return fmt.Sprintf(`
apiVersion: apps/v1
kind: Deployment
metadata:
  name: %s
  namespace: %s

spec:
  replicas: %d

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
          image: %s:%s

          imagePullPolicy: Never

          ports:
            - containerPort: %d

          env:%s
`,
		spec.Name,
		spec.Namespace,
		spec.Replicas,
		spec.Name,
		spec.Name,
		spec.Name,
		spec.Image,
		spec.Tag,
		spec.Port,
		indentEnv(envVars),
	)
}

func indentEnv(env string) string {

	if strings.TrimSpace(env) == "" {
		return " []"
	}

	return env
}
