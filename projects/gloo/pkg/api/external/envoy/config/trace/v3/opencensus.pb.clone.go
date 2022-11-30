// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/external/envoy/config/trace/v3/opencensus.proto

package v3

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/solo-io/protoc-gen-ext/pkg/clone"
	"google.golang.org/protobuf/proto"

	github_com_solo_io_solo_kit_pkg_api_v1_resources_core "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
)

// ensure the imports are used
var (
	_ = errors.New("")
	_ = fmt.Print
	_ = binary.LittleEndian
	_ = bytes.Compare
	_ = strings.Compare
	_ = clone.Cloner(nil)
	_ = proto.Message(nil)
)

// Clone function
func (m *OpenCensusConfig) Clone() proto.Message {
	var target *OpenCensusConfig
	if m == nil {
		return target
	}
	target = &OpenCensusConfig{}

	if h, ok := interface{}(m.GetTraceConfig()).(clone.Cloner); ok {
		target.TraceConfig = h.Clone().(*TraceConfig)
	} else {
		target.TraceConfig = proto.Clone(m.GetTraceConfig()).(*TraceConfig)
	}

	target.OcagentExporterEnabled = m.GetOcagentExporterEnabled()

	target.OcagentAddress = m.GetOcagentAddress()

	if h, ok := interface{}(m.GetOcagentGrpcService()).(clone.Cloner); ok {
		target.OcagentGrpcService = h.Clone().(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.ResourceRef)
	} else {
		target.OcagentGrpcService = proto.Clone(m.GetOcagentGrpcService()).(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.ResourceRef)
	}

	if m.GetIncomingTraceContext() != nil {
		target.IncomingTraceContext = make([]OpenCensusConfig_TraceContext, len(m.GetIncomingTraceContext()))
		for idx, v := range m.GetIncomingTraceContext() {

			target.IncomingTraceContext[idx] = v

		}
	}

	if m.GetOutgoingTraceContext() != nil {
		target.OutgoingTraceContext = make([]OpenCensusConfig_TraceContext, len(m.GetOutgoingTraceContext()))
		for idx, v := range m.GetOutgoingTraceContext() {

			target.OutgoingTraceContext[idx] = v

		}
	}

	return target
}

// Clone function
func (m *TraceConfig) Clone() proto.Message {
	var target *TraceConfig
	if m == nil {
		return target
	}
	target = &TraceConfig{}

	target.MaxNumberOfAttributes = m.GetMaxNumberOfAttributes()

	target.MaxNumberOfAnnotations = m.GetMaxNumberOfAnnotations()

	target.MaxNumberOfMessageEvents = m.GetMaxNumberOfMessageEvents()

	target.MaxNumberOfLinks = m.GetMaxNumberOfLinks()

	switch m.Sampler.(type) {

	case *TraceConfig_ProbabilitySampler:

		if h, ok := interface{}(m.GetProbabilitySampler()).(clone.Cloner); ok {
			target.Sampler = &TraceConfig_ProbabilitySampler{
				ProbabilitySampler: h.Clone().(*ProbabilitySampler),
			}
		} else {
			target.Sampler = &TraceConfig_ProbabilitySampler{
				ProbabilitySampler: proto.Clone(m.GetProbabilitySampler()).(*ProbabilitySampler),
			}
		}

	case *TraceConfig_ConstantSampler:

		if h, ok := interface{}(m.GetConstantSampler()).(clone.Cloner); ok {
			target.Sampler = &TraceConfig_ConstantSampler{
				ConstantSampler: h.Clone().(*ConstantSampler),
			}
		} else {
			target.Sampler = &TraceConfig_ConstantSampler{
				ConstantSampler: proto.Clone(m.GetConstantSampler()).(*ConstantSampler),
			}
		}

	case *TraceConfig_RateLimitingSampler:

		if h, ok := interface{}(m.GetRateLimitingSampler()).(clone.Cloner); ok {
			target.Sampler = &TraceConfig_RateLimitingSampler{
				RateLimitingSampler: h.Clone().(*RateLimitingSampler),
			}
		} else {
			target.Sampler = &TraceConfig_RateLimitingSampler{
				RateLimitingSampler: proto.Clone(m.GetRateLimitingSampler()).(*RateLimitingSampler),
			}
		}

	}

	return target
}

// Clone function
func (m *ProbabilitySampler) Clone() proto.Message {
	var target *ProbabilitySampler
	if m == nil {
		return target
	}
	target = &ProbabilitySampler{}

	target.SamplingProbability = m.GetSamplingProbability()

	return target
}

// Clone function
func (m *ConstantSampler) Clone() proto.Message {
	var target *ConstantSampler
	if m == nil {
		return target
	}
	target = &ConstantSampler{}

	target.Decision = m.GetDecision()

	return target
}

// Clone function
func (m *RateLimitingSampler) Clone() proto.Message {
	var target *RateLimitingSampler
	if m == nil {
		return target
	}
	target = &RateLimitingSampler{}

	target.Qps = m.GetQps()

	return target
}
