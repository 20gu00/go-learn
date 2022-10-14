package main

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"log"
	"time"
)

func main() {
	config, err := clientcmd.BuildConfigFromFlags("", "/root/. kube/config")
	if err != nil {
		panic(err)
	}
	//⾸先通过 kubernetes.NewForConfig 创建 clientset 对象， Informer 需要通过 ClientSet 与 Kubernetes API Server 进⾏交互。另外，创建
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	//stopCh 对象，该对象⽤于在程序进程退出之前通知 Informer 提前退出，因为 Informer 是⼀个持久运⾏的 goroutine 。
	stopCh := make(chan struct{})
	defer close(stopCh)
	//数实例化了 SharedInformer 对象
	//第1个参数 clientset 是⽤于与 Kubernetes
	//API Server 交互的客户端，第2个参数 time.Minute ⽤于设置多久进⾏⼀次 resync （重新同步）， resync 会周期性地执⾏List操作，
	//将所有的资源存放在 Informer Store 中，如果该参数为0，则禁⽤ resync 功能。
	sharedInformers := informers.NewSharedInformerFactory(clientset, time.Minute)
	//通过 sharedInformers.Core（）.V1（）.Pods（）.Informer 可以得到具体 Pod 资源的 informer 对象
	informer := sharedInformers.Core().V1().Pods().Informer()

	//通过 informer.AddEventHandler 函数可以为 Pod 资源添加资源事件回调⽅法，⽀持3种资源事件回调⽅法，分别介绍如下。
	//AddFunc ：当创建 Pod 资源对象时触发的事件回调⽅法。
	//UpdateFunc ：当更新 Pod 资源对象时触发的事件回调⽅法。
	//DeleteFunc ：当删除 Pod 资源对象时触发的事件回调⽅法。
	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			mObj := obj.(v1.Object)
			log.Printf("New Pod Added to Store: %s", mObj.GetName())
		},
		UpdateFunc: func(old0bj, newObj interface{}) {
			oObj := old0bj.(v1.Object)
			nObj := newObj.(v1.Object)
			log.Printf("%s Pod Updated to %s", nObj.GetName(), nObj.GetName())
		},
		DeleteFunc: func(obj interface{}) {
			mObj := obj.(v1.Object)
			log.Printf("Pod Deleted from Store: s", mObj.GetName())
		},
	})
	//通过 informer.Run 函数运⾏当前的 Informer ，内部为 Pod 资源类型创建 Informer
	informer.Run(stopCh)
}
