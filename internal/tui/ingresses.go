package tui

import (
	"github.com/AnatolyRugalev/kube-commander/internal/kube"
	"github.com/AnatolyRugalev/kube-commander/internal/widgets"
	netv1beta1 "k8s.io/api/networking/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"strings"
)

type IngressesTable struct {
	namespace string
}

func (it *IngressesTable) Namespace() string {
	return it.namespace
}

func (it *IngressesTable) GetActions() []*widgets.ListAction {
	return GetDefaultActions(it)
}

func (it *IngressesTable) DeleteDescription(idx int, row widgets.ListRow) string {
	return "Ingress " + row[0]
}

func (it *IngressesTable) Delete(idx int, row widgets.ListRow) error {
	return kube.GetClient().NetworkingV1beta1().Ingresses(it.namespace).Delete(row[0], metav1.NewDeleteOptions(0))
}

func (it *IngressesTable) TypeName() string {
	return "ingress"
}

func (it *IngressesTable) Name(row widgets.ListRow) string {
	return row[0]
}

func NewIngressesTable(namespace string) *widgets.DataTable {
	lt := widgets.NewDataTable(&IngressesTable{
		namespace: namespace,
	}, screen)
	lt.Title = "Ingresses"
	return lt
}

func (it *IngressesTable) GetHeaderRow() widgets.ListRow {
	return widgets.ListRow{"NAME", "HOSTS", "PORTS", "AGE"}
}

func (it *IngressesTable) LoadData() ([]widgets.ListRow, error) {
	Ingresses, err := kube.GetClient().NetworkingV1beta1().Ingresses(it.namespace).List(metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	var rows []widgets.ListRow
	for _, service := range Ingresses.Items {
		rows = append(rows, it.newRow(service))
	}
	return rows, nil
}

func (it *IngressesTable) newRow(svc netv1beta1.Ingress) widgets.ListRow {
	ports := []string{"80"}
	for _, rule := range svc.Spec.Rules {
		for _, tls := range svc.Spec.TLS {
			for _, host := range tls.Hosts {
				if host == rule.Host {
					ports = append(ports, "443")
				}
			}
		}
	}
	return widgets.ListRow{
		svc.Name,
		strings.Join(ports, ", "),
		Age(svc.CreationTimestamp.Time),
	}
}