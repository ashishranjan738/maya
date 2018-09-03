package command

import (
	"testing"

	"github.com/openebs/maya/pkg/apis/openebs.io/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestGetCASType(t *testing.T) {
	tests := map[string]struct {
		expectedOutput string
		Volume         VolumeInfo
	}{
		"Fetching CasType when CasType is Jiva": {
			Volume: VolumeInfo{
				Volume: v1alpha1.CASVolume{
					Spec: v1alpha1.CASVolumeSpec{
						CasType: "jiva",
					},
				},
			},
			expectedOutput: "jiva",
		},
		"Fetching CasType when CasType is cstor": {
			Volume: VolumeInfo{
				Volume: v1alpha1.CASVolume{
					Spec: v1alpha1.CASVolumeSpec{
						CasType: "cstor",
					},
				},
			},
			expectedOutput: "cstor",
		},
		"Fetching CasType when CasType is none": {
			Volume: VolumeInfo{
				Volume: v1alpha1.CASVolume{
					Spec: v1alpha1.CASVolumeSpec{
						CasType: "",
					},
				},
			},
			expectedOutput: "jiva",
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := tt.Volume.GetCASType()
			if got != tt.expectedOutput {
				t.Fatalf("Test: %v Expected: %v but got: %v", name, tt.expectedOutput, got)
			}
		})
	}
}

func TestGetClusterIP(t *testing.T) {
	tests := map[string]struct {
		expectedOutput string
		Volume         VolumeInfo
	}{
		"Fetching ClusterIP from openebs.io/cluster-ips": {
			Volume: VolumeInfo{
				Volume: v1alpha1.CASVolume{
					ObjectMeta: metav1.ObjectMeta{
						Annotations: map[string]string{
							"openebs.io/cluster-ips": "192.168.100.1",
						},
					},
				},
			},
			expectedOutput: "192.168.100.1",
		},
		"Fetching ClusterIP from vsm.openebs.io/cluster-ips": {
			Volume: VolumeInfo{
				Volume: v1alpha1.CASVolume{
					ObjectMeta: metav1.ObjectMeta{
						Annotations: map[string]string{
							"vsm.openebs.io/cluster-ips": "192.168.100.1",
						},
					},
				},
			},
			expectedOutput: "192.168.100.1",
		},
		"Fetching ClusterIP when both keys are present": {
			Volume: VolumeInfo{
				Volume: v1alpha1.CASVolume{
					ObjectMeta: metav1.ObjectMeta{
						Annotations: map[string]string{
							"vsm.openebs.io/cluster-ips": "192.168.100.1",
							"openebs.io/cluster-ips":     "192.168.100.2",
						},
					},
				},
			},
			expectedOutput: "192.168.100.2",
		},
		"Fetching ClusterIP when both keys are not present": {
			Volume: VolumeInfo{
				Volume: v1alpha1.CASVolume{
					ObjectMeta: metav1.ObjectMeta{
						Annotations: map[string]string{},
					},
				},
			},
			expectedOutput: "",
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := tt.Volume.GetClusterIP()
			if got != tt.expectedOutput {
				t.Fatalf("Test: %v Expected: %v but got: %v", name, tt.expectedOutput, got)
			}
		})
	}
}

func TestGetControllerStatus(t *testing.T) {
	tests := map[string]struct {
		expectedOutput string
		Volume         VolumeInfo
	}{
		"Fetching Controller status from openebs.io/controller-status": {
			Volume: VolumeInfo{
				Volume: v1alpha1.CASVolume{
					ObjectMeta: metav1.ObjectMeta{
						Annotations: map[string]string{
							"openebs.io/controller-status": "running",
						},
					},
				},
			},
			expectedOutput: "running",
		},
		"Fetching Controller status from vsm.openebs.io/controller-status": {
			Volume: VolumeInfo{
				Volume: v1alpha1.CASVolume{
					ObjectMeta: metav1.ObjectMeta{
						Annotations: map[string]string{
							"vsm.openebs.io/controller-status": "running",
						},
					},
				},
			},
			expectedOutput: "running",
		},
		"Fetching Controller status when both keys are present": {
			Volume: VolumeInfo{
				Volume: v1alpha1.CASVolume{
					ObjectMeta: metav1.ObjectMeta{
						Annotations: map[string]string{
							"vsm.openebs.io/controller-status": "running",
							"openebs.io/controller-status":     "evicted",
						},
					},
				},
			},
			expectedOutput: "evicted",
		},
		"Fetching Controller status when both keys are not present": {
			Volume: VolumeInfo{
				Volume: v1alpha1.CASVolume{
					ObjectMeta: metav1.ObjectMeta{
						Annotations: map[string]string{},
					},
				},
			},
			expectedOutput: "",
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := tt.Volume.GetControllerStatus()
			if got != tt.expectedOutput {
				t.Fatalf("Test: %v Expected: %v but got: %v", name, tt.expectedOutput, got)
			}
		})
	}
}

func TestGetIQN(t *testing.T) {
	tests := map[string]struct {
		expectedOutput string
		Volume         VolumeInfo
	}{
		"Fetching IQN": {
			Volume: VolumeInfo{
				Volume: v1alpha1.CASVolume{
					ObjectMeta: metav1.ObjectMeta{
						Annotations: map[string]string{},
					},
					Spec: v1alpha1.CASVolumeSpec{
						Iqn: "iqn.2016-09.com.openebs.cstor:default-testclaim7",
					},
				},
			},
			expectedOutput: "iqn.2016-09.com.openebs.cstor:default-testclaim7",
		},
		"Fetching Controller when iqn is present in openebs.io annotations": {
			Volume: VolumeInfo{
				Volume: v1alpha1.CASVolume{
					ObjectMeta: metav1.ObjectMeta{
						Annotations: map[string]string{
							"openebs.io/iqn": "iqn.2016-09.com.openebs.cstor:default-testclaim7",
						},
					},
				},
			},
			expectedOutput: "iqn.2016-09.com.openebs.cstor:default-testclaim7",
		},
		"Fetching Controller when iqn is present in vsm.openebs.io annotations": {
			Volume: VolumeInfo{
				Volume: v1alpha1.CASVolume{
					ObjectMeta: metav1.ObjectMeta{
						Annotations: map[string]string{
							"vsm.openebs.io/iqn": "iqn.2016-09.com.openebs.cstor:default-testclaim7",
						},
					},
				},
			},
			expectedOutput: "iqn.2016-09.com.openebs.cstor:default-testclaim7",
		},
		"Fetching Controller when iqn is not present": {
			Volume: VolumeInfo{
				Volume: v1alpha1.CASVolume{
					ObjectMeta: metav1.ObjectMeta{},
				},
			},
			expectedOutput: "",
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := tt.Volume.GetIQN()
			if got != tt.expectedOutput {
				t.Fatalf("Test: %v Expected: %v but got: %v", name, tt.expectedOutput, got)
			}
		})
	}
}

func TestGetVolumeName(t *testing.T) {
	tests := map[string]struct {
		expectedOutput string
		Volume         VolumeInfo
	}{
		"Fetching VolumeInfo": {
			Volume: VolumeInfo{
				Volume: v1alpha1.CASVolume{
					ObjectMeta: metav1.ObjectMeta{
						Annotations: map[string]string{},
						Name:        "default-testclaim",
					},
				},
			},
			expectedOutput: "default-testclaim",
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := tt.Volume.GetVolumeName()
			if got != tt.expectedOutput {
				t.Fatalf("Test: %v Expected: %v but got: %v", name, tt.expectedOutput, got)
			}
		})
	}
}

func TestGetTargetPortal(t *testing.T) {
	tests := map[string]struct {
		expectedOutput string
		Volume         VolumeInfo
	}{
		"Fetching TargetPortal": {
			Volume: VolumeInfo{
				Volume: v1alpha1.CASVolume{
					ObjectMeta: metav1.ObjectMeta{
						Annotations: map[string]string{},
					},
					Spec: v1alpha1.CASVolumeSpec{
						TargetPortal: "10.63.247.173:3260",
					},
				},
			},
			expectedOutput: "10.63.247.173:3260",
		},
		"Fetching TargetPortal when it is present in openebs.io annotations": {
			Volume: VolumeInfo{
				Volume: v1alpha1.CASVolume{
					ObjectMeta: metav1.ObjectMeta{
						Annotations: map[string]string{
							"openebs.io/targetportals": "10.35.244.116:3260",
						},
					},
				},
			},
			expectedOutput: "10.35.244.116:3260",
		},
		"Fetching TargetPortal when it is present in vsm.openebs.io annotations": {
			Volume: VolumeInfo{
				Volume: v1alpha1.CASVolume{
					ObjectMeta: metav1.ObjectMeta{
						Annotations: map[string]string{
							"vsm.openebs.io/targetportals": "10.35.244.116:3260",
						},
					},
				},
			},
			expectedOutput: "10.35.244.116:3260",
		},
		"Fetching TargetPortal when it is not present": {
			Volume: VolumeInfo{
				Volume: v1alpha1.CASVolume{
					ObjectMeta: metav1.ObjectMeta{},
				},
			},
			expectedOutput: "",
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := tt.Volume.GetTargetPortal()
			if got != tt.expectedOutput {
				t.Fatalf("Test: %v Expected: %v but got: %v", name, tt.expectedOutput, got)
			}
		})
	}
}

func TestGetVolumeSize(t *testing.T) {
	tests := map[string]struct {
		expectedOutput string
		Volume         VolumeInfo
	}{
		"Fetching VolumeSize": {
			Volume: VolumeInfo{
				Volume: v1alpha1.CASVolume{
					ObjectMeta: metav1.ObjectMeta{
						Annotations: map[string]string{},
					},
					Spec: v1alpha1.CASVolumeSpec{
						Capacity: "5G",
					},
				},
			},
			expectedOutput: "5G",
		},
		"Fetching VolumeSize when it is present in openebs.io annotations": {
			Volume: VolumeInfo{
				Volume: v1alpha1.CASVolume{
					ObjectMeta: metav1.ObjectMeta{
						Annotations: map[string]string{
							"openebs.io/volume-size": "5G",
						},
					},
				},
			},
			expectedOutput: "5G",
		},
		"Fetching Volume Size it is present in vsm.openebs.io annotations": {
			Volume: VolumeInfo{
				Volume: v1alpha1.CASVolume{
					ObjectMeta: metav1.ObjectMeta{
						Annotations: map[string]string{
							"vsm.openebs.io/volume-size": "5G",
						},
					},
				},
			},
			expectedOutput: "5G",
		},
		"Fetching Volume Size it is not present": {
			Volume: VolumeInfo{
				Volume: v1alpha1.CASVolume{
					ObjectMeta: metav1.ObjectMeta{},
				},
			},
			expectedOutput: "",
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := tt.Volume.GetVolumeSize()
			if got != tt.expectedOutput {
				t.Fatalf("Test: %v Expected: %v but got: %v", name, tt.expectedOutput, got)
			}
		})
	}
}

func TestGetReplicaCount(t *testing.T) {
	tests := map[string]struct {
		expectedOutput string
		Volume         VolumeInfo
	}{
		"Fetching ReplicaCount": {
			Volume: VolumeInfo{
				Volume: v1alpha1.CASVolume{
					ObjectMeta: metav1.ObjectMeta{
						Annotations: map[string]string{},
					},
					Spec: v1alpha1.CASVolumeSpec{
						Replicas: "3",
					},
				},
			},
			expectedOutput: "3",
		},
		"Fetching ReplicaCount  when it is present in openebs.io annotations": {
			Volume: VolumeInfo{
				Volume: v1alpha1.CASVolume{
					ObjectMeta: metav1.ObjectMeta{
						Annotations: map[string]string{
							"openebs.io/replica-count": "3",
						},
					},
				},
			},
			expectedOutput: "3",
		},
		"Fetching ReplicaCount when it is present in vsm.openebs.io annotations": {
			Volume: VolumeInfo{
				Volume: v1alpha1.CASVolume{
					ObjectMeta: metav1.ObjectMeta{
						Annotations: map[string]string{
							"vsm.openebs.io/replica-count": "3",
						},
					},
				},
			},
			expectedOutput: "3",
		},
		"Fetching ReplicaCount when it is not present": {
			Volume: VolumeInfo{
				Volume: v1alpha1.CASVolume{
					ObjectMeta: metav1.ObjectMeta{},
				},
			},
			expectedOutput: "",
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := tt.Volume.GetReplicaCount()
			if got != tt.expectedOutput {
				t.Fatalf("Test: %v Expected: %v but got: %v", name, tt.expectedOutput, got)
			}
		})
	}
}

func TestGetReplicaStatus(t *testing.T) {
	tests := map[string]struct {
		expectedOutput string
		Volume         VolumeInfo
	}{
		"Fetching ReplicaStatus from openebs.io/replica-status": {
			Volume: VolumeInfo{
				Volume: v1alpha1.CASVolume{
					ObjectMeta: metav1.ObjectMeta{
						Annotations: map[string]string{
							"openebs.io/replica-status": "running, running, running",
						},
					},
				},
			},
			expectedOutput: "running, running, running",
		},
		"Fetching ReplicaStatus from vsm.openebs.io/replica-status": {
			Volume: VolumeInfo{
				Volume: v1alpha1.CASVolume{
					ObjectMeta: metav1.ObjectMeta{
						Annotations: map[string]string{
							"vsm.openebs.io/replica-status": "running, running, running",
						},
					},
				},
			},
			expectedOutput: "running, running, running",
		},
		"Fetching Replica status when both keys are present": {
			Volume: VolumeInfo{
				Volume: v1alpha1.CASVolume{
					ObjectMeta: metav1.ObjectMeta{
						Annotations: map[string]string{
							"openebs.io/replica-status":     "running, running, running",
							"vsm.openebs.io/replica-status": "running, running, running",
						},
					},
				},
			},
			expectedOutput: "running, running, running",
		},
		"Fetching ReplicaStatus when no key is present": {
			Volume: VolumeInfo{
				Volume: v1alpha1.CASVolume{
					ObjectMeta: metav1.ObjectMeta{
						Annotations: map[string]string{},
					},
				},
			},
			expectedOutput: "",
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := tt.Volume.GetReplicaStatus()
			if got != tt.expectedOutput {
				t.Fatalf("Test: %v Expected: %v but got: %v", name, tt.expectedOutput, got)
			}
		})
	}
}

func TestGetReplicaIP(t *testing.T) {
	tests := map[string]struct {
		expectedOutput string
		Volume         VolumeInfo
	}{
		"Fetching ReplicaIP from openebs.io/replica-ips": {
			Volume: VolumeInfo{
				Volume: v1alpha1.CASVolume{
					ObjectMeta: metav1.ObjectMeta{
						Annotations: map[string]string{
							"openebs.io/replica-ips": "10.60.0.11, 10.60.1.16, 10.60.2.10",
						},
					},
				},
			},
			expectedOutput: "10.60.0.11, 10.60.1.16, 10.60.2.10",
		},
		"Fetching ReplicaIP from vsm.openebs.io/replica-ips": {
			Volume: VolumeInfo{
				Volume: v1alpha1.CASVolume{
					ObjectMeta: metav1.ObjectMeta{
						Annotations: map[string]string{
							"vsm.openebs.io/replica-ips": "10.60.0.11, 10.60.1.16, 10.60.2.10",
						},
					},
				},
			},
			expectedOutput: "10.60.0.11, 10.60.1.16, 10.60.2.10",
		},

		"Fetching ReplicaIP when both keys are present": {
			Volume: VolumeInfo{
				Volume: v1alpha1.CASVolume{
					ObjectMeta: metav1.ObjectMeta{
						Annotations: map[string]string{
							"vsm.openebs.io/replica-ips": "10.60.0.11, 10.60.1.16, 10.60.2.10",
							"openebs.io/replica-ips":     "10.60.0.11, 10.60.1.16, 10.60.2.10",
						},
					},
				},
			},
			expectedOutput: "10.60.0.11, 10.60.1.16, 10.60.2.10",
		},
		"Fetching ReplicaIP when no key is present": {
			Volume: VolumeInfo{
				Volume: v1alpha1.CASVolume{
					ObjectMeta: metav1.ObjectMeta{
						Annotations: map[string]string{},
					},
				},
			},
			expectedOutput: "",
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := tt.Volume.GetReplicaIP()
			if got != tt.expectedOutput {
				t.Fatalf("Test: %v Expected: %v but got: %v", name, tt.expectedOutput, got)
			}
		})
	}
}

func TestGetStoragePool(t *testing.T) {
	tests := map[string]struct {
		expectedOutput string
		Volume         VolumeInfo
	}{
		"Fetching StoragePool from openebs.io/openebs.io/pool-names": {
			Volume: VolumeInfo{
				Volume: v1alpha1.CASVolume{
					ObjectMeta: metav1.ObjectMeta{
						Annotations: map[string]string{
							"openebs.io/pool-names": "cstor-sparse-pool-g7e8,cstor-sparse-pool-l9dp,cstor-sparse-pool-yq8t",
						},
					},
				},
			},
			expectedOutput: "cstor-sparse-pool-g7e8,cstor-sparse-pool-l9dp,cstor-sparse-pool-yq8t",
		},
		"Fetching StoragePool when no key is present": {
			Volume: VolumeInfo{
				Volume: v1alpha1.CASVolume{
					ObjectMeta: metav1.ObjectMeta{
						Annotations: map[string]string{},
					},
				},
			},
			expectedOutput: "",
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := tt.Volume.GetStoragePool()
			if got != tt.expectedOutput {
				t.Fatalf("Test: %v Expected: %v but got: %v", name, tt.expectedOutput, got)
			}
		})
	}
}

func TestGetCVRName(t *testing.T) {
	tests := map[string]struct {
		expectedOutput string
		Volume         VolumeInfo
	}{
		"Fetching CVRName from openebs.io/cvr-names": {
			Volume: VolumeInfo{
				Volume: v1alpha1.CASVolume{
					ObjectMeta: metav1.ObjectMeta{
						Annotations: map[string]string{
							"openebs.io/cvr-names": "default-cstor-volume-3227802448-cstor-sparse-pool-g7e8,default-cstor-volume-3227802448-cstor-sparse-pool-l9dp,default-cstor-volume-3227802448-cstor-sparse-pool-yq8t",
						},
					},
				},
			},
			expectedOutput: "default-cstor-volume-3227802448-cstor-sparse-pool-g7e8,default-cstor-volume-3227802448-cstor-sparse-pool-l9dp,default-cstor-volume-3227802448-cstor-sparse-pool-yq8t",
		},
		"Fetching CVRName when no key is present": {
			Volume: VolumeInfo{
				Volume: v1alpha1.CASVolume{
					ObjectMeta: metav1.ObjectMeta{
						Annotations: map[string]string{},
					},
				},
			},
			expectedOutput: "",
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := tt.Volume.GetCVRName()
			if got != tt.expectedOutput {
				t.Fatalf("Test: %v Expected: %v but got: %v", name, tt.expectedOutput, got)
			}
		})
	}
}

func TestGetNodeName(t *testing.T) {
	tests := map[string]struct {
		expectedOutput string
		Volume         VolumeInfo
	}{
		"Fetching Node Name from openebs.io/node-names": {
			Volume: VolumeInfo{
				Volume: v1alpha1.CASVolume{
					ObjectMeta: metav1.ObjectMeta{
						Annotations: map[string]string{
							"openebs.io/node-names": "gke-ashish-dev-default-pool-1fe155b7-rvqd,gke-ashish-dev-default-pool-1fe155b7-qv7v,gke-ashish-dev-default-pool-1fe155b7-w75t",
						},
					},
				},
			},
			expectedOutput: "gke-ashish-dev-default-pool-1fe155b7-rvqd,gke-ashish-dev-default-pool-1fe155b7-qv7v,gke-ashish-dev-default-pool-1fe155b7-w75t",
		},
		"Fetching Node Name when no key is present": {
			Volume: VolumeInfo{
				Volume: v1alpha1.CASVolume{
					ObjectMeta: metav1.ObjectMeta{
						Annotations: map[string]string{},
					},
				},
			},
			expectedOutput: "",
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := tt.Volume.GetNodeName()
			if got != tt.expectedOutput {
				t.Fatalf("Test: %v Expected: %v but got: %v", name, tt.expectedOutput, got)
			}
		})
	}
}

func TestGetVolumeNamespace(t *testing.T) {
	tests := map[string]struct {
		expectedOutput string
		Volume         VolumeInfo
	}{
		"Fetching volume namespace when present in response": {
			Volume: VolumeInfo{
				Volume: v1alpha1.CASVolume{
					ObjectMeta: metav1.ObjectMeta{
						Namespace: "default",
					},
				},
			},
			expectedOutput: "default",
		},
		"Fetching volume namespace when not present in response": {
			Volume: VolumeInfo{
				Volume: v1alpha1.CASVolume{
					ObjectMeta: metav1.ObjectMeta{},
				},
			},
			expectedOutput: "N/A",
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := tt.Volume.GetVolumeNamespace()
			if got != tt.expectedOutput {
				t.Fatalf("Test: %v Expected: %v but got: %v", name, tt.expectedOutput, got)
			}
		})
	}
}

func TestGetVolumeStatus(t *testing.T) {
	tests := map[string]struct {
		expectedOutput string
		Volume         VolumeInfo
	}{
		"Fetching volume status when reason is empty": {
			Volume: VolumeInfo{
				Volume: v1alpha1.CASVolume{},
			},
			expectedOutput: "Running",
		},
		"Fetching volume status when reason is not empty": {
			Volume: VolumeInfo{
				Volume: v1alpha1.CASVolume{
					Status: v1alpha1.CASVolumeStatus{
						Reason: "Some Error",
					},
				},
			},
			expectedOutput: "Some Error",
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := tt.Volume.GetVolumeStatus()
			if got != tt.expectedOutput {
				t.Fatalf("Test: %v Expected: %v but got: %v", name, tt.expectedOutput, got)
			}
		})
	}
}