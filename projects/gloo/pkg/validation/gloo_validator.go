package validation

import (
	"context"

	"github.com/solo-io/gloo/projects/gloo/pkg/api/grpc/validation"
	gloov1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/gloo/projects/gloo/pkg/api/v1/gloosnapshot"
	"github.com/solo-io/gloo/projects/gloo/pkg/plugins"
	"github.com/solo-io/gloo/projects/gloo/pkg/syncer"
	"github.com/solo-io/gloo/projects/gloo/pkg/syncer/sanitizer"
	"github.com/solo-io/go-utils/contextutils"
	"github.com/solo-io/solo-kit/pkg/api/v2/reporter"
	"k8s.io/apimachinery/pkg/runtime/schema"

	ratelimit "github.com/solo-io/gloo/projects/gloo/pkg/api/external/solo/ratelimit"
	gloo_translator "github.com/solo-io/gloo/projects/gloo/pkg/translator"
)

const GlooGroup = "gloo.solo.io"

var GvkToSupportedDeleteGlooResources = map[schema.GroupVersionKind]bool{
	gloov1.UpstreamGVK:           true,
	gloov1.SecretGVK:             true,
	ratelimit.RateLimitConfigGVK: true,
}

var GvkToSupportedGlooResources = map[schema.GroupVersionKind]bool{
	gloov1.UpstreamGVK:           true,
	ratelimit.RateLimitConfigGVK: true,
}

// GlooValidator is used to validate solo.io.gloo resources
type GlooValidator interface {
	Validate(ctx context.Context, proxy *gloov1.Proxy, snapshot *gloosnapshot.ApiSnapshot, delete bool) []*GlooValidationReport
}

type GlooValidatorConfig struct {
	GlooTranslator gloo_translator.Translator
	XdsSanitizer   sanitizer.XdsSanitizer
	Extensions     []syncer.TranslatorSyncerExtension
}

// NewGlooValidator will create a new GlooValidator
func NewGlooValidator(config GlooValidatorConfig) GlooValidator {
	return glooValidator{
		glooTranslator: config.GlooTranslator,
		xdsSanitizer:   config.XdsSanitizer,
		extensions:     config.Extensions,
	}
}

type glooValidator struct {
	glooTranslator gloo_translator.Translator
	xdsSanitizer   sanitizer.XdsSanitizer
	extensions     []syncer.TranslatorSyncerExtension
}

type GlooValidationReport struct {
	Proxy           *gloov1.Proxy
	ProxyReport     *validation.ProxyReport
	ResourceReports reporter.ResourceReports
}

func (gv glooValidator) Validate(ctx context.Context, proxy *gloov1.Proxy, snapshot *gloosnapshot.ApiSnapshot, delete bool) []*GlooValidationReport {
	ctx = contextutils.WithLogger(ctx, "proxy-validator")

	var validationReports []*GlooValidationReport
	var proxiesToValidate gloov1.ProxyList

	if proxy != nil {
		proxiesToValidate = gloov1.ProxyList{proxy}
	} else {
		// if no proxy was passed in, call translate for all proxies in snapshot
		proxiesToValidate = snapshot.Proxies
	}

	if len(proxiesToValidate) == 0 {
		// This can occur when a Gloo resource (Upstream), is modified before the ApiSnapshot
		// contains any Proxies. Orphaned resources are never invalid, but they may be accepted
		// even if they are semantically incorrect.
		// This log line is attempting to identify these situations
		contextutils.LoggerFrom(ctx).Warnf("found no proxies to validate, accepting update without translating Gloo resources")
		return validationReports
	}

	params := plugins.Params{
		Ctx:      ctx,
		Snapshot: snapshot,
	}
	// Validation with gateway occurs in /projects/gateway/pkg/validation/validator.go, where validation for the Gloo
	// resources occurs in the following for loop.
	for _, proxy := range proxiesToValidate {
		// so params has the ctx, snapshot, proxy
		xdsSnapshot, resourceReports, proxyReport := gv.glooTranslator.Translate(params, proxy)

		// Sanitize routes before sending report to gateway
		gv.xdsSanitizer.SanitizeSnapshot(ctx, snapshot, xdsSnapshot, resourceReports)
		routeErrorToWarnings(resourceReports, proxyReport)

		// TODO-JAKE we might want to set the proxies of the snapshot here so that they align with the proxiesToValidate list above...
		for _, ex := range gv.extensions {
			err := ex.Translate(ctx, snapshot, v1.ProxyList{proxy}, resourceReports)
			// TODO-JAKE not sure if we want to have an error here, because these errors are not really respective of the proxy resource
			if err != nil {
				resourceReports.AddError(proxy, err)
			}
		}

		validationReports = append(validationReports, &GlooValidationReport{
			Proxy:           proxy,
			ProxyReport:     proxyReport,
			ResourceReports: resourceReports,
		})
	}

	return validationReports
}