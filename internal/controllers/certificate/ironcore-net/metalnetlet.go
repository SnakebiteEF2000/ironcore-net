// SPDX-FileCopyrightText: 2023 SAP SE or an SAP affiliate company and IronCore contributors
// SPDX-License-Identifier: Apache-2.0

package ironcorenet

import (
	"crypto/x509"
	"fmt"
	"strings"

	ironcorenetv1alpha1 "github.com/ironcore-dev/ironcore-net/api/core/v1alpha1"
	"github.com/ironcore-dev/ironcore-net/internal/controllers/certificate/generic"
	"golang.org/x/exp/slices"
	authv1 "k8s.io/api/authorization/v1"
	certificatesv1 "k8s.io/api/certificates/v1"
	"k8s.io/apimachinery/pkg/util/sets"
)

var (
	MetalnetletRequiredUsages = sets.New[certificatesv1.KeyUsage](
		certificatesv1.UsageDigitalSignature,
		certificatesv1.UsageKeyEncipherment,
		certificatesv1.UsageClientAuth,
	)
)

func IsMetalnetletClientCert(csr *certificatesv1.CertificateSigningRequest, x509cr *x509.CertificateRequest) bool {
	if csr.Spec.SignerName != certificatesv1.KubeAPIServerClientSignerName {
		return false
	}

	return ValidateMetalnetletClientCSR(x509cr, sets.New(csr.Spec.Usages...)) == nil
}

func ValidateMetalnetletClientCSR(req *x509.CertificateRequest, usages sets.Set[certificatesv1.KeyUsage]) error {
	if !slices.Equal([]string{ironcorenetv1alpha1.MetalnetletsGroup}, req.Subject.Organization) {
		return fmt.Errorf("organization is not %s", ironcorenetv1alpha1.MetalnetletsGroup)
	}

	if len(req.DNSNames) > 0 {
		return fmt.Errorf("dns subject alternative names are not allowed")
	}
	if len(req.EmailAddresses) > 0 {
		return fmt.Errorf("email subject alternative names are not allowed")
	}
	if len(req.IPAddresses) > 0 {
		return fmt.Errorf("ip subject alternative names are not allowed")
	}
	if len(req.URIs) > 0 {
		return fmt.Errorf("uri subject alternative names are not allowed")
	}

	if !strings.HasPrefix(req.Subject.CommonName, ironcorenetv1alpha1.MetalnetletUserNamePrefix) {
		return fmt.Errorf("subject common name does not begin with %s", ironcorenetv1alpha1.MetalnetletUserNamePrefix)
	}

	if !MetalnetletRequiredUsages.Equal(usages) {
		return fmt.Errorf("usages did not match %v", sets.List(MetalnetletRequiredUsages))
	}

	return nil
}

var (
	MetalnetletRecognizer = generic.NewCertificateSigningRequestRecognizer(
		IsMetalnetletClientCert,
		authv1.ResourceAttributes{
			Group:       certificatesv1.GroupName,
			Resource:    "certificatesigningrequests",
			Verb:        "create",
			Subresource: "metalnetletclient",
		},
		"Auto approving metalnetlet client certificate after SubjectAccessReview.",
	)
)

func init() {
	Recognizers = append(Recognizers, MetalnetletRecognizer)
}
