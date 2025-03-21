/*
Copyright 2025.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controller

import (
	"context"
	"os"

	"github.com/ganawaj/go-vyos/vyos"
	"github.com/joho/godotenv"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	vyosv1 "vyos.kubebuilder.io/project/api/v1"
)

// FirewallRuleReconciler reconciles a FirewallRule object
type FirewallRuleReconciler struct {
	client.Client
	Scheme     *runtime.Scheme
	VyosClient *vyos.Client
}

// +kubebuilder:rbac:groups=vyos.vyos.kubebuilder.io,resources=firewallrules,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=vyos.vyos.kubebuilder.io,resources=firewallrules/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=vyos.vyos.kubebuilder.io,resources=firewallrules/finalizers,verbs=update

func NewVyosClient() (*vyos.Client, error) {
	_ = godotenv.Load(".env")
	token := os.Getenv("TOKEN")
	url := os.Getenv("VYOS_ADDR")

	return vyos.NewClient(nil).
		WithToken(token).
		WithURL(url).
		Insecure(), nil
}

func (r *FirewallRuleReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := log.FromContext(ctx)
	firewallRule := &vyosv1.FirewallRule{}
	if err := r.Get(ctx, req.NamespacedName, firewallRule); err != nil {
		log.Error(err, "unable to fetch FirewallRule")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	if r.VyosClient == nil {
		client, err := NewVyosClient()
		if err != nil {
			log.Error(err, "Unable to create VyOS client")
			return ctrl.Result{}, err
		}
		r.VyosClient = client
	}

	out, _, err := r.VyosClient.Show.Do(ctx, "conf")
	if err != nil {
		log.Error(err, "Error executing command")
		return ctrl.Result{}, err
	}

	log.Info("Successfully fetched configuration", "data", out.Data)
	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *FirewallRuleReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&vyosv1.FirewallRule{}).
		Named("firewallrule").
		Complete(r)
}
