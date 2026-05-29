package kubernetes

import (
	"fmt"

	"github.com/codeurluce/kubelaunch/backend/models"
)

func GenerateService(spec models.DeploymentSpec) string {

	return fmt.Sprintf(`
apiVersion: v1
kind: Service

metadata:
  name: %s-service
  namespace: %s

spec:
  type: %s

  selector:
    app: %s

  ports:
    - protocol: TCP
      port: %d
      targetPort: %d
`,
		spec.Name,
		spec.Namespace,
		spec.ServiceType,
		spec.Name,
		spec.Port,
		spec.Port,
	)
}
