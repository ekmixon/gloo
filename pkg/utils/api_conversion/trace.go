package api_conversion

import (
	"fmt"

	v1 "github.com/census-instrumentation/opencensus-proto/gen-go/trace/v1"
	envoycore "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	envoytrace "github.com/envoyproxy/go-control-plane/envoy/config/trace/v3"
	envoytracegloo "github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/config/trace/v3"
)

// Converts between Envoy and Gloo/solokit versions of envoy protos
// This is required because go-control-plane dropped gogoproto in favor of goproto
// in v0.9.0, but solokit depends on gogoproto (and the generated deep equals it creates).
//
// we should work to remove that assumption from solokit and delete this code:
// https://github.com/solo-io/gloo/issues/1793

func ToEnvoyDatadogConfiguration(glooDatadogConfig *envoytracegloo.DatadogConfig, clusterName string) (*envoytrace.DatadogConfig, error) {
	envoyDatadogConfig := &envoytrace.DatadogConfig{
		CollectorCluster: clusterName,
		ServiceName:      glooDatadogConfig.GetServiceName().GetValue(),
	}
	return envoyDatadogConfig, nil
}

func ToEnvoyZipkinConfiguration(glooZipkinConfig *envoytracegloo.ZipkinConfig, clusterName string) (*envoytrace.ZipkinConfig, error) {
	envoyZipkinConfig := &envoytrace.ZipkinConfig{
		CollectorCluster:         clusterName,
		CollectorEndpoint:        glooZipkinConfig.GetCollectorEndpoint(),
		CollectorEndpointVersion: ToEnvoyZipkinCollectorEndpointVersion(glooZipkinConfig.GetCollectorEndpointVersion()),
		TraceId_128Bit:           glooZipkinConfig.GetTraceId_128Bit().GetValue(),
		SharedSpanContext:        glooZipkinConfig.GetSharedSpanContext(),
	}
	return envoyZipkinConfig, nil
}

func ToEnvoyOpenCensusConfiguration(glooOpenCensusConfig *envoytracegloo.OpenCensusConfig) (*envoytrace.OpenCensusConfig, error) {
	grpcService := glooOpenCensusConfig.GetOcagentGrpcService()
	grpcClusterName := grpcService.GetName()
	if grpcService.GetNamespace() != "" {
		grpcClusterName = fmt.Sprintf("%s_%s", grpcService.GetName(), grpcService.GetNamespace())
	}
	envoyGrpcService := envoycore.GrpcService{
		TargetSpecifier: &envoycore.GrpcService_EnvoyGrpc_{
			EnvoyGrpc: &envoycore.GrpcService_EnvoyGrpc{
				ClusterName: grpcClusterName,
			},
		},
	}

	envoyOpenCensusConfig := &envoytrace.OpenCensusConfig{
		TraceConfig: &v1.TraceConfig{
			Sampler:                  nil,
			MaxNumberOfAttributes:    glooOpenCensusConfig.GetTraceConfig().GetMaxNumberOfAttributes(),
			MaxNumberOfAnnotations:   glooOpenCensusConfig.GetTraceConfig().GetMaxNumberOfAnnotations(),
			MaxNumberOfMessageEvents: glooOpenCensusConfig.GetTraceConfig().GetMaxNumberOfMessageEvents(),
			MaxNumberOfLinks:         glooOpenCensusConfig.GetTraceConfig().GetMaxNumberOfLinks(),
		},
		OcagentExporterEnabled: glooOpenCensusConfig.GetOcagentExporterEnabled(),
		OcagentAddress:         glooOpenCensusConfig.GetOcagentAddress(),
		OcagentGrpcService:     &envoyGrpcService,
		IncomingTraceContext:   translateTraceContext(glooOpenCensusConfig.GetIncomingTraceContext()),
		OutgoingTraceContext:   translateTraceContext(glooOpenCensusConfig.GetOutgoingTraceContext()),
	}

	glooTraceConfig := glooOpenCensusConfig.GetTraceConfig()
	switch glooTraceConfig.GetSampler().(type) {
	case *envoytracegloo.TraceConfig_ConstantSampler:
		var decision v1.ConstantSampler_ConstantDecision
		switch glooTraceConfig.GetConstantSampler().GetDecision() {
		case envoytracegloo.ConstantSampler_ALWAYS_ON:
			decision = v1.ConstantSampler_ALWAYS_ON
		case envoytracegloo.ConstantSampler_ALWAYS_OFF:
			decision = v1.ConstantSampler_ALWAYS_OFF
		case envoytracegloo.ConstantSampler_ALWAYS_PARENT:
			decision = v1.ConstantSampler_ALWAYS_PARENT
		}
		envoyOpenCensusConfig.GetTraceConfig().Sampler = &v1.TraceConfig_ConstantSampler{
			ConstantSampler: &v1.ConstantSampler{
				Decision: decision,
			},
		}
	case *envoytracegloo.TraceConfig_ProbabilitySampler:
		envoyOpenCensusConfig.GetTraceConfig().Sampler = &v1.TraceConfig_ProbabilitySampler{
			ProbabilitySampler: &v1.ProbabilitySampler{
				SamplingProbability: glooTraceConfig.GetProbabilitySampler().GetSamplingProbability(),
			},
		}
	case *envoytracegloo.TraceConfig_RateLimitingSampler:
		envoyOpenCensusConfig.GetTraceConfig().Sampler = &v1.TraceConfig_RateLimitingSampler{RateLimitingSampler: &v1.RateLimitingSampler{
			Qps: glooTraceConfig.GetRateLimitingSampler().GetQps(),
		}}
	}

	return envoyOpenCensusConfig, nil
}

func translateTraceContext(glooTraceContexts []envoytracegloo.OpenCensusConfig_TraceContext) []envoytrace.OpenCensusConfig_TraceContext {
	result := make([]envoytrace.OpenCensusConfig_TraceContext, 0, len(glooTraceContexts))
	for i, glooTraceContext := range glooTraceContexts {
		var envoyTraceContext envoytrace.OpenCensusConfig_TraceContext
		switch glooTraceContext {
		case envoytracegloo.OpenCensusConfig_NONE:
			envoyTraceContext = envoytrace.OpenCensusConfig_NONE
		case envoytracegloo.OpenCensusConfig_TRACE_CONTEXT:
			envoyTraceContext = envoytrace.OpenCensusConfig_TRACE_CONTEXT
		case envoytracegloo.OpenCensusConfig_GRPC_TRACE_BIN:
			envoyTraceContext = envoytrace.OpenCensusConfig_GRPC_TRACE_BIN
		case envoytracegloo.OpenCensusConfig_CLOUD_TRACE_CONTEXT:
			envoyTraceContext = envoytrace.OpenCensusConfig_CLOUD_TRACE_CONTEXT
		case envoytracegloo.OpenCensusConfig_B3:
			envoyTraceContext = envoytrace.OpenCensusConfig_B3
		}
		result[i] = envoyTraceContext
	}
	return result
}

func ToEnvoyZipkinCollectorEndpointVersion(version envoytracegloo.ZipkinConfig_CollectorEndpointVersion) envoytrace.ZipkinConfig_CollectorEndpointVersion {
	switch str := version.String(); str {
	case envoytracegloo.ZipkinConfig_CollectorEndpointVersion_name[int32(envoytracegloo.ZipkinConfig_HTTP_JSON)]:
		return envoytrace.ZipkinConfig_HTTP_JSON
	case envoytracegloo.ZipkinConfig_CollectorEndpointVersion_name[int32(envoytracegloo.ZipkinConfig_HTTP_PROTO)]:
		return envoytrace.ZipkinConfig_HTTP_PROTO
	}
	return envoytrace.ZipkinConfig_HTTP_JSON
}
