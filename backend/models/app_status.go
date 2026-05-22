package models

type AppStatus struct {
	Name          string    `json:"name"`
	Namespace     string    `json:"namespace"`
	Stack         StackType `json:"stack"`
	Port          int32     `json:"port"`
	Replicas      int32     `json:"replicas"`
	ReadyReplicas int32     `json:"readyReplicas"`
	Status        string    `json:"status"`
	CPUUsage      string    `json:"cpu"`
	MemoryUsage   string    `json:"memory"`
	CreatedAt     string    `json:"createdAt"`
}
