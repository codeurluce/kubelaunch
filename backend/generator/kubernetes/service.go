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
spec:
  selector:
    app: %s
  ports:
    - protocol: TCP
      port: 80
      targetPort: %d
  type: ClusterIP
`,
		spec.Name,
		spec.Name,
		spec.Port,
	)
}
