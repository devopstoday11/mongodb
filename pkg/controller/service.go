package controller

import (
	"fmt"

	"github.com/appscode/kutil"
	core_util "github.com/appscode/kutil/core/v1"
	api "github.com/kubedb/apimachinery/apis/kubedb/v1alpha1"
	"github.com/kubedb/apimachinery/pkg/eventer"
	core "k8s.io/api/core/v1"
	kerr "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	clientsetscheme "k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/reference"
	mona "kmodules.xyz/monitoring-agent-api/api/v1"
)

func (c *Controller) ensureService(mongodb *api.MongoDB) (kutil.VerbType, error) {
	// Check if service name exists
	if err := c.checkService(mongodb); err != nil {
		return kutil.VerbUnchanged, err
	}

	// create database Service
	vt, err := c.createService(mongodb)
	if err != nil {
		if ref, rerr := reference.GetReference(clientsetscheme.Scheme, mongodb); rerr == nil {
			c.recorder.Eventf(
				ref,
				core.EventTypeWarning,
				eventer.EventReasonFailedToCreate,
				"Failed to create Service. Reason: %v",
				err,
			)
		}
		return kutil.VerbUnchanged, err
	} else if vt != kutil.VerbUnchanged {
		if ref, rerr := reference.GetReference(clientsetscheme.Scheme, mongodb); rerr == nil {
			c.recorder.Eventf(
				ref,
				core.EventTypeNormal,
				eventer.EventReasonSuccessful,
				"Successfully %s Service",
				vt,
			)
		}
	}
	return vt, nil
}

func (c *Controller) checkService(mongodb *api.MongoDB) error {
	name := mongodb.OffshootName()
	service, err := c.Client.CoreV1().Services(mongodb.Namespace).Get(name, metav1.GetOptions{})
	if err != nil {
		if kerr.IsNotFound(err) {
			return nil
		}
		return err
	}

	if service.Spec.Selector[api.LabelDatabaseName] != name {
		return fmt.Errorf(`intended service "%v" already exists`, name)
	}

	return nil
}

func (c *Controller) createService(mongodb *api.MongoDB) (kutil.VerbType, error) {
	meta := metav1.ObjectMeta{
		Name:      mongodb.OffshootName(),
		Namespace: mongodb.Namespace,
	}

	ref, rerr := reference.GetReference(clientsetscheme.Scheme, mongodb)
	if rerr != nil {
		return kutil.VerbUnchanged, rerr
	}

	_, ok, err := core_util.CreateOrPatchService(c.Client, meta, func(in *core.Service) *core.Service {
		in.ObjectMeta = core_util.EnsureOwnerReference(in.ObjectMeta, ref)
		in.Labels = mongodb.OffshootLabels()
		in.Annotations = mongodb.Spec.ServiceTemplate.Annotations

		in.Spec.Selector = mongodb.OffshootSelectors()
		in.Spec.Ports = upsertServicePort(in, mongodb)

		if mongodb.Spec.ServiceTemplate.Spec.ClusterIP != "" {
			in.Spec.ClusterIP = mongodb.Spec.ServiceTemplate.Spec.ClusterIP
		}
		if mongodb.Spec.ServiceTemplate.Spec.Type != "" {
			in.Spec.Type = mongodb.Spec.ServiceTemplate.Spec.Type
		}
		in.Spec.ExternalIPs = mongodb.Spec.ServiceTemplate.Spec.ExternalIPs
		in.Spec.LoadBalancerIP = mongodb.Spec.ServiceTemplate.Spec.LoadBalancerIP
		in.Spec.LoadBalancerSourceRanges = mongodb.Spec.ServiceTemplate.Spec.LoadBalancerSourceRanges
		in.Spec.ExternalTrafficPolicy = mongodb.Spec.ServiceTemplate.Spec.ExternalTrafficPolicy
		in.Spec.HealthCheckNodePort = mongodb.Spec.ServiceTemplate.Spec.HealthCheckNodePort
		return in
	})
	return ok, err
}

func upsertServicePort(service *core.Service, mongodb *api.MongoDB) []core.ServicePort {
	desiredPorts := []core.ServicePort{
		{
			Name:       "db",
			Protocol:   core.ProtocolTCP,
			Port:       27017,
			TargetPort: intstr.FromString("db"),
		},
	}
	if mongodb.GetMonitoringVendor() == mona.VendorPrometheus {
		desiredPorts = append(desiredPorts, core.ServicePort{
			Name:       api.PrometheusExporterPortName,
			Protocol:   core.ProtocolTCP,
			Port:       mongodb.Spec.Monitor.Prometheus.Port,
			TargetPort: intstr.FromString(api.PrometheusExporterPortName),
		})
	}
	return core_util.MergeServicePorts(service.Spec.Ports, desiredPorts)
}
