package controllers

import (
	ctrl "sigs.k8s.io/controller-runtime"
)

func AddToManager(mgr ctrl.Manager, concilers ...interface{}) {

	for _, conciler := range concilers {
		switch conciler.(type) {
		case *PodReconciler:
			conciler.(*PodReconciler).Scheme = mgr.GetScheme()
			conciler.(*PodReconciler).Client = mgr.GetClient()
			err := conciler.(*PodReconciler).SetupWithManager(mgr)
			if err != nil {
				panic(err.Error())
			}

		case *ConfigMapReconciler:
			conciler.(*ConfigMapReconciler).Scheme = mgr.GetScheme()
			conciler.(*ConfigMapReconciler).Client = mgr.GetClient()
			err := conciler.(*ConfigMapReconciler).SetupWithManager(mgr)
			if err != nil {
				panic(err.Error())
			}
		}

	}
	//
	//PodConciler = &PodReconciler{Scheme: mgr.GetScheme(), Client: mgr.GetClient()}
	//err := PodConciler.SetupWithManager(mgr)
	//
	//ConfigMapConciler = &ConfigMapReconciler{Scheme: mgr.GetScheme(), Client: mgr.GetClient()}
	//
	//err = ConfigMapConciler.SetupWithManager(mgr)
	//if err != nil {
	//	panic(err.Error())
	//}
}
