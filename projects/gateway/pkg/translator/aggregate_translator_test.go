package translator_test

import (
	"context"
	"errors"
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"google.golang.org/protobuf/types/known/durationpb"

	v1 "github.com/solo-io/gloo/projects/gateway/pkg/api/v1"
	. "github.com/solo-io/gloo/projects/gateway/pkg/translator"
	gloov1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	gloov1snap "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/gloosnapshot"
	"github.com/solo-io/gloo/test/samples"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	"github.com/solo-io/solo-kit/pkg/api/v2/reporter"
)

var _ = FDescribe("Aggregate translator", func() {
	var (
		ctx = context.TODO()

		snap    *gloov1snap.ApiSnapshot
		proxy   *gloov1.Proxy
		reports reporter.ResourceReports
		ns      = "namespace"
	)

	genProxyWithTranslatorOpts := func(opts Opts) {
		tx := NewDefaultTranslator(opts)
		proxy, reports = tx.Translate(ctx, "proxy-name", snap, snap.Gateways)
	}

	genProxyWithIsolatedVirtualHosts := func() {
		genProxyWithTranslatorOpts(Opts{
			WriteNamespace:                 ns,
			IsolateVirtualHostsBySslConfig: true,
		})
	}

	BeforeEach(func() {
		snap = samples.SimpleGlooSnapshot(ns)
	})

	It("Computes listener idempotently when provided different ssl configs", func() {
		gw1 := snap.Gateways[1]
		gw := gw1.GetHttpGateway()
		gw.VirtualServiceExpressions = nil
		gw.VirtualServiceSelector = nil
		gw.VirtualServices = append(gw.VirtualServices, &core.ResourceRef{
			Name:      "ssl-vs-0",
			Namespace: ns,
		}, &core.ResourceRef{
			Name:      "ssl-vs-1",
			Namespace: ns,
		}, &core.ResourceRef{
			Name:      "ssl-vs-2",
			Namespace: ns,
		}, &core.ResourceRef{
			Name:      "ssl-vs-3",
			Namespace: ns,
		}, &core.ResourceRef{
			Name:      "ssl-vs-4",
			Namespace: ns,
		}, &core.ResourceRef{
			Name:      "ssl-vs-5-empty",
			Namespace: ns,
		})
		snap.Gateways = v1.GatewayList{gw1}

		snap.VirtualServices = append(snap.VirtualServices, &v1.VirtualService{
			VirtualHost: &v1.VirtualHost{},
			SslConfig: &gloov1.SslConfig{
				SniDomains: []string{"sni-0"},
				// We have to add some other config since we merge configs where the only
				// difference is the SniDomains
				TransportSocketConnectTimeout: &durationpb.Duration{Seconds: 0},
			},
			DisplayName: "ssl-vs-0",
			Metadata: &core.Metadata{
				Name:      "ssl-vs-0",
				Namespace: ns,
			},
		}, &v1.VirtualService{
			VirtualHost: &v1.VirtualHost{},
			SslConfig: &gloov1.SslConfig{
				SniDomains: []string{"sni-1"},
				// We have to add some other config since we merge configs where the only
				// difference is the SniDomains
				TransportSocketConnectTimeout: &durationpb.Duration{Seconds: 1},
			},
			DisplayName: "ssl-vs-1",
			Metadata: &core.Metadata{
				Name:      "ssl-vs-1",
				Namespace: ns,
			},
		}, &v1.VirtualService{
			VirtualHost: &v1.VirtualHost{},
			SslConfig: &gloov1.SslConfig{
				SniDomains: []string{"sni-2"},
				// We have to add some other config since we merge configs where the only
				// difference is the SniDomains
				TransportSocketConnectTimeout: &durationpb.Duration{Seconds: 2},
			},
			DisplayName: "ssl-vs-2",
			Metadata: &core.Metadata{
				Name:      "ssl-vs-2",
				Namespace: ns,
			},
		}, &v1.VirtualService{
			VirtualHost: &v1.VirtualHost{},
			SslConfig: &gloov1.SslConfig{
				SniDomains: []string{"sni-3"},
				// We have to add some other config since we merge configs where the only
				// difference is the SniDomains
				TransportSocketConnectTimeout: &durationpb.Duration{Seconds: 3},
			},
			DisplayName: "ssl-vs-3",
			Metadata: &core.Metadata{
				Name:      "ssl-vs-3",
				Namespace: ns,
			},
		}, &v1.VirtualService{
			VirtualHost: &v1.VirtualHost{},
			SslConfig: &gloov1.SslConfig{
				SniDomains: []string{"sni-4"},
				// We have to add some other config since we merge configs where the only
				// difference is the SniDomains
				TransportSocketConnectTimeout: &durationpb.Duration{Seconds: 4},
			},
			DisplayName: "ssl-vs-4",
			Metadata: &core.Metadata{
				Name:      "ssl-vs-4",
				Namespace: ns,
			},
		},
			&v1.VirtualService{
				VirtualHost: &v1.VirtualHost{},
				SslConfig: &gloov1.SslConfig{
					SniDomains: []string{},
					// We have to add some other config since we merge configs where the only
					// difference is the SniDomains
					TransportSocketConnectTimeout: &durationpb.Duration{Seconds: 5},
				},
				DisplayName: "ssl-vs-5-empty",
				Metadata: &core.Metadata{
					Name:      "ssl-vs-5-empty",
					Namespace: ns,
				},
			})
		genProxyWithIsolatedVirtualHosts()
		proxyName := proxy.Metadata.Name
		aggregateTranslator := &AggregateTranslator{VirtualServiceTranslator: &VirtualServiceTranslator{}}

		firstSNIOrError := func(chain *gloov1.AggregateListener_HttpFilterChain) (sni string, err error) {
			sniDomains := chain.GetMatcher().GetSslConfig().GetSniDomains()
			if len(sniDomains) == 0 {
				return "", errors.New("no sni domains")
			}
			return sniDomains[0], nil
		}
		// run 100 times to ensure idempotency
		// not sure if 100 times is valid; in anecdotal testing it tended to fail in under 20
		for i := 0; i < 100; i++ {
			l := aggregateTranslator.ComputeListener(NewTranslatorParams(ctx, snap, reports), proxyName, snap.Gateways[0])
			Expect(l).NotTo(BeNil())
			Expect(l.GetAggregateListener())
			Expect(l.GetAggregateListener().HttpFilterChains).To(HaveLen(6), fmt.Sprintf("output filterchains %v", l.GetAggregateListener().HttpFilterChains))
			// since we sort on hashes, this is the ordered output of this config
			Expect(firstSNIOrError(l.GetAggregateListener().HttpFilterChains[0])).To(Equal("sni-1"))
			Expect(firstSNIOrError(l.GetAggregateListener().HttpFilterChains[1])).To(Equal("sni-4"))
			_, err := firstSNIOrError(l.GetAggregateListener().HttpFilterChains[2])
			Expect(err).To(HaveOccurred(), fmt.Sprintf("should have had no snidomain and error, output filterchains %v", l.GetAggregateListener().HttpFilterChains))

			Expect(firstSNIOrError(l.GetAggregateListener().HttpFilterChains[3])).To(Equal("sni-3"))
			Expect(firstSNIOrError(l.GetAggregateListener().HttpFilterChains[4])).To(Equal("sni-0"))
			Expect(firstSNIOrError(l.GetAggregateListener().HttpFilterChains[5])).To(Equal("sni-2"))
		}
	})
})
