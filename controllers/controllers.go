package controllers

import (
	ctrl "sigs.k8s.io/controller-runtime"
)

var (
	PodConciler       *PodReconciler
	ConfigMapConciler *ConfigMapReconciler
)

func AddToManager(mgr ctrl.Manager) {
	PodConciler = &PodReconciler{Scheme: mgr.GetScheme(), Client: mgr.GetClient()}
	err := PodConciler.SetupWithManager(mgr)
	if err != nil {
		panic(err.Error())
	}

	ConfigMapConciler = &ConfigMapReconciler{Scheme: mgr.GetScheme(), Client: mgr.GetClient()}

	err = ConfigMapConciler.SetupWithManager(mgr)
	if err != nil {
		panic(err.Error())
	}
}
