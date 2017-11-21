package config

type PVCTemplateConfig struct {

}

type KubernetesConfig struct {
	APIServerHost string
	APIServerPort int
	APIVersion string
	Namespace string
}

func GetKubernetesConfig() *KubernetesConfig {
	kubernetes := &KubernetesConfig{
		APIServerHost: "localhost",
		APIServerPort: 8001,
		APIVersion: "1.7"
		PVCTemplate: pvcTemplate,
	}

	return kubernetes
}