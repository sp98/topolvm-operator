package monitor

import (
	"bytes"
	"fmt"
	"github.com/alauda/topolvm-operator/pkg/cluster"
	"github.com/alauda/topolvm-operator/pkg/operator/k8sutil"
	"github.com/pkg/errors"
	monitoringv1 "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sYAML "k8s.io/apimachinery/pkg/util/yaml"
)

const (
	serviceMonitorName           = "topolvm-service-monitor"
	interval                     = "30s"
	path                         = "/metrics"
	port                         = "metrics"
	prometheusRuleClusterLabel   = "alert.cpaas.io/cluster"
	prometheusRuleNameSpaceLabel = "alert.cpaas.io/namespace"
)

func EnableServiceMonitor(ref *metav1.OwnerReference) error {
	serviceMonitor := monitoringv1.ServiceMonitor{
		ObjectMeta: metav1.ObjectMeta{
			Name:            serviceMonitorName,
			Namespace:       cluster.NameSpace,
			OwnerReferences: []metav1.OwnerReference{*ref},
			Labels: map[string]string{
				"prometheus": "kube-prometheus",
			},
		},
		Spec: monitoringv1.ServiceMonitorSpec{
			Endpoints: []monitoringv1.Endpoint{
				{Interval: interval, Path: path, Port: port},
			},
			NamespaceSelector: monitoringv1.NamespaceSelector{
				MatchNames: []string{cluster.NameSpace},
			},
			Selector: metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app.kubernetes.io/compose": "metrics",
				},
			},
		},
	}
	if _, err := k8sutil.CreateOrUpdateServiceMonitor(&serviceMonitor); err != nil {
		return errors.Wrap(err, "service monitor could not be enabled")
	}
	return nil
}

func CreateOrUpdatePrometheusRule(ref *metav1.OwnerReference, k8sClusterName string) error {
	var rule monitoringv1.PrometheusRule
	err := k8sYAML.NewYAMLOrJSONDecoder(bytes.NewBufferString(PrometheusRule), 1000).Decode(&rule)
	if err != nil {
		return fmt.Errorf("prometheusRules could not be decoded. %v", err)
	}
	rule.OwnerReferences = []metav1.OwnerReference{*ref}
	rule.Namespace = cluster.NameSpace
	addK8sClusterLable(&rule, k8sClusterName)
	_, err = k8sutil.CreateOrUpdatePrometheusRule(&rule)
	return err
}

func addK8sClusterLable(rule *monitoringv1.PrometheusRule, k8sClusterName string) {

	rule.Labels[prometheusRuleClusterLabel] = k8sClusterName
	rule.Labels[prometheusRuleNameSpaceLabel] = cluster.NameSpace
	for i, _ := range rule.Spec.Groups {
		for j, _ := range rule.Spec.Groups[i].Rules {
			rule.Spec.Groups[i].Rules[j].Labels["alert_cluster"] = k8sClusterName
		}
	}
}