package main

import (
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func main() {
	// KnownType external  有版本资源,外部版本
	coreGV := schema.GroupVersion{Group: "", Version: "v1"}
	extensionsGV := schema.GroupVersion{Group: "extensions", Version: "v1beta1"}
	// KnownType internal  有版本的资源,内部版本
	coreInternalGV := schema.GroupVersion{Group: "", Version: runtime.APIVersionInternal} //runtime.APIVersionInternal,内部版本,开发使用,kubectl拿到的是用户可以使用的外部版本

	// UnversionedType  无版本类型
	Unversioned := schema.GroupVersion{Group: "", Version: "v1"}
	//实例化一个新的注册表
	scheme := runtime.NewScheme()
	scheme.AddKnownTypes(coreGV, &corev1.Pod{})
	scheme.AddKnownTypes(extensionsGV, &appsv1.DaemonSet{})
	scheme.AddKnownTypes(coreInternalGV, &corev1.Pod{})
	scheme.AddUnversionedTypes(Unversioned, &metav1.Status{}) //status是无版本类型的资源
}
