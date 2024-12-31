package k8s

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"path"
	"strconv"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func CreatePods(client *kubernetes.Clientset, count int) error {
	for i := 0; i < count; i++ {
		randint := rand.Intn(50)
		podName := "pod-" + strconv.Itoa(i+1)

		pod := &corev1.Pod{
			ObjectMeta: metav1.ObjectMeta{
				Name:   podName,
				Labels: map[string]string{"value": strconv.Itoa(randint), "app": "binary-tree"},
			},
			Spec: corev1.PodSpec{
				Containers: []corev1.Container{
					{
						Name:  "nginx-container",
						Image: "nginx",
						Ports: []corev1.ContainerPort{
							{
								ContainerPort: 80,
							},
						},
					},
				},
			},
		}

		_, err := client.CoreV1().Pods("default").Create(context.TODO(), pod, metav1.CreateOptions{})
		if err != nil {
			return fmt.Errorf("failed to create pod %s: %v", podName, err)
		}

		fmt.Printf("Pod %s created with label value: %d\n", podName, randint)
	}

	return nil
}

// GetPods retrieves all pods with the label app=binary-tree
func GetPods(client *kubernetes.Clientset) ([]corev1.Pod, error) {
	ctx := context.Background()
	// List pods with the label app=binary-tree
	podList, err := client.CoreV1().Pods("default").List(ctx, metav1.ListOptions{
		LabelSelector: "app=binary-tree",
	})
	if err != nil {
		return nil, fmt.Errorf("failed to list pods: %v", err)
	}

	for _, pod := range podList.Items {
		fmt.Println("Pod Name:", pod.Name)
	}
	return podList.Items, nil
}

func NewK8sClient() (*kubernetes.Clientset, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	config, err := clientcmd.BuildConfigFromFlags("", path.Join(home, ".kube/config"))
	if err != nil {
		panic(err.Error())
	}

	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	return client, nil
}

// CleanUpPods deletes all pods with the label app=binary-tree
func CleanUpPods(client *kubernetes.Clientset) error {
	ctx := context.Background()
	pods, err := GetPods(client)
	if err != nil {
		return fmt.Errorf("failed to get pods: %v", err)
	}

	for _, pod := range pods {
		err := client.CoreV1().Pods("default").Delete(ctx, pod.Name, metav1.DeleteOptions{})
		if err != nil {
			return fmt.Errorf("failed to delete pod %s: %v", pod.Name, err)
		}
		fmt.Printf("Pod %s deleted successfully\n", pod.Name)
	}

	return nil
}
