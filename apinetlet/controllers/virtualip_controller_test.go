// Copyright 2022 OnMetal authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package controllers

import (
	onmetalapinetv1alpha1 "github.com/onmetal/onmetal-api-net/api/v1alpha1"
	commonv1alpha1 "github.com/onmetal/onmetal-api/apis/common/v1alpha1"
	networkingv1alpha1 "github.com/onmetal/onmetal-api/apis/networking/v1alpha1"
	"github.com/onmetal/onmetal-api/testutils"
	mcmeta "github.com/onmetal/poollet/multicluster/meta"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	. "sigs.k8s.io/controller-runtime/pkg/envtest/komega"
)

var _ = Describe("VirtualIPController", func() {
	ctx := testutils.SetupContext()
	ns := SetupTest(ctx)

	It("should allocate a public ip", func() {
		By("creating a virtual ip")
		virtualIP := &networkingv1alpha1.VirtualIP{
			ObjectMeta: metav1.ObjectMeta{
				Namespace:    ns.Name,
				GenerateName: "virtual-ip-",
			},
			Spec: networkingv1alpha1.VirtualIPSpec{
				Type:     networkingv1alpha1.VirtualIPTypePublic,
				IPFamily: corev1.IPv4Protocol,
			},
		}
		Expect(k8sClient.Create(ctx, virtualIP)).To(Succeed())

		By("waiting for the corresponding public ip to be created")
		publicIP := &onmetalapinetv1alpha1.PublicIP{
			ObjectMeta: metav1.ObjectMeta{
				Namespace: ns.Name,
				Name:      string(virtualIP.UID),
			},
		}
		Eventually(Get(publicIP)).Should(Succeed())

		By("inspecting the created public ip")
		Expect(publicIP.Spec).To(Equal(onmetalapinetv1alpha1.PublicIPSpec{
			IPFamilies: []corev1.IPFamily{corev1.IPv4Protocol},
			IPs:        nil,
			AllocatorRef: onmetalapinetv1alpha1.AllocatorRef{
				ClusterName: clusterName,
				Group:       networkingv1alpha1.SchemeGroupVersion.Group,
				Resource:    "virtualips",
				Namespace:   ns.Name,
				Name:        virtualIP.Name,
				UID:         virtualIP.UID,
			},
		}))

		By("asserting the virtual ip does not get an ip address")
		Consistently(Object(virtualIP)).Should(HaveField("Status.IP", BeNil()))

		By("setting the public ip to allocated")
		allocatedIP := commonv1alpha1.MustParseIP("10.0.0.1")
		publicIP.Status.IPs = []commonv1alpha1.IP{allocatedIP}
		Expect(k8sClient.Status().Update(ctx, publicIP)).To(Succeed())

		By("waiting for the virtual ip to reflect the allocated ip")
		Eventually(Object(virtualIP)).Should(HaveField("Status.IP", Equal(&allocatedIP)))

		By("deleting the virtual ip")
		Expect(k8sClient.Delete(ctx, virtualIP)).To(Succeed())

		By("waiting for it to be gone")
		Eventually(Get(virtualIP)).Should(Satisfy(apierrors.IsNotFound))

		By("asserting the corresponding public ip is gone as well")
		Expect(k8sClient.Get(ctx, client.ObjectKeyFromObject(publicIP), publicIP)).To(Satisfy(apierrors.IsNotFound))
	})

	It("should clean up dangling public ips", func() {
		By("creating a public ip")
		publicIP := &onmetalapinetv1alpha1.PublicIP{
			ObjectMeta: metav1.ObjectMeta{
				Namespace:    ns.Name,
				GenerateName: "public-ip-",
			},
			Spec: onmetalapinetv1alpha1.PublicIPSpec{
				IPFamilies: []corev1.IPFamily{corev1.IPv4Protocol},
				AllocatorRef: onmetalapinetv1alpha1.AllocatorRef{
					ClusterName: clusterName,
					Group:       networkingv1alpha1.SchemeGroupVersion.Group,
					Resource:    "virtualips",
					Namespace:   ns.Name,
					Name:        "some-name",
					UID:         "some-uuid",
				},
			},
		}
		Expect(k8sClient.Create(ctx, publicIP)).To(Succeed())

		By("waiting for the public ip to be gone")
		Eventually(Get(publicIP)).Should(Satisfy(apierrors.IsNotFound))
	})

	It("should only get the public ip if it's ancestor-managed", func() {
		By("creating a virtual ip with an ancestor")
		virtualIP := &networkingv1alpha1.VirtualIP{
			ObjectMeta: metav1.ObjectMeta{
				Namespace:    ns.Name,
				GenerateName: "virtual-ip-",
			},
			Spec: networkingv1alpha1.VirtualIPSpec{
				Type:     networkingv1alpha1.VirtualIPTypePublic,
				IPFamily: corev1.IPv4Protocol,
			},
		}
		const (
			rootAncestorClusterName = "root-ancestor-cluster"
			rootAncestorNamespace   = "root-ancestor-namespace"
			rootAncestorName        = "root-ancestor-name"
			rootAncestorUID         = "root-ancestor-uid"
		)
		mcmeta.SetAncestors(virtualIP, []mcmeta.Ancestor{
			{
				ClusterName: rootAncestorClusterName,
				Namespace:   rootAncestorNamespace,
				Name:        rootAncestorName,
				UID:         rootAncestorUID,
			},
		})
		Expect(k8sClient.Create(ctx, virtualIP)).Should(Succeed())

		By("asserting no public ip gets created")
		Consistently(Get(&onmetalapinetv1alpha1.PublicIP{
			ObjectMeta: metav1.ObjectMeta{
				Namespace: ns.Name,
				Name:      string(virtualIP.UID),
			},
		})).Should(Satisfy(apierrors.IsNotFound))

		By("creating a public ip with the matching name")
		publicIP := &onmetalapinetv1alpha1.PublicIP{
			ObjectMeta: metav1.ObjectMeta{
				Namespace: ns.Name,
				Name:      rootAncestorUID,
			},
			Spec: onmetalapinetv1alpha1.PublicIPSpec{
				IPFamilies: []corev1.IPFamily{corev1.IPv4Protocol},
				AllocatorRef: onmetalapinetv1alpha1.AllocatorRef{
					ClusterName: rootAncestorClusterName,
					Group:       networkingv1alpha1.SchemeGroupVersion.Group,
					Resource:    "virtualips",
					Namespace:   rootAncestorNamespace,
					Name:        rootAncestorName,
					UID:         rootAncestorUID,
				},
			},
		}
		Expect(k8sClient.Create(ctx, publicIP)).To(Succeed())

		By("patching the public ip to allocated")
		ip := commonv1alpha1.MustParseIP("10.0.0.1")
		publicIP.Status.IPs = []commonv1alpha1.IP{ip}
		Expect(k8sClient.Status().Update(ctx, publicIP)).To(Succeed())

		By("waiting for the virtual ip to reflect the allocated ip")
		Eventually(Object(virtualIP)).Should(HaveField("Status.IP", Equal(&ip)))

		By("deleting the virtual ip")
		Expect(k8sClient.Delete(ctx, virtualIP)).Should(Succeed())

		By("asserting the ancestor-managed public ip doesn't get deleted")
		Consistently(Get(publicIP)).Should(Succeed())
	})
})
