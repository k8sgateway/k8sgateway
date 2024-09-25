// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/enterprise/options/ai/ai.proto

package ai

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/solo-io/protoc-gen-ext/pkg/clone"
	"google.golang.org/protobuf/proto"

	github_com_solo_io_solo_kit_pkg_api_v1_resources_core "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"

	google_golang_org_protobuf_types_known_structpb "google.golang.org/protobuf/types/known/structpb"
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
func (m *SingleAuthToken) Clone() proto.Message {
	var target *SingleAuthToken
	if m == nil {
		return target
	}
	target = &SingleAuthToken{}

	switch m.AuthTokenSource.(type) {

	case *SingleAuthToken_Inline:

		target.AuthTokenSource = &SingleAuthToken_Inline{
			Inline: m.GetInline(),
		}

	case *SingleAuthToken_SecretRef:

		if h, ok := interface{}(m.GetSecretRef()).(clone.Cloner); ok {
			target.AuthTokenSource = &SingleAuthToken_SecretRef{
				SecretRef: h.Clone().(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.ResourceRef),
			}
		} else {
			target.AuthTokenSource = &SingleAuthToken_SecretRef{
				SecretRef: proto.Clone(m.GetSecretRef()).(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.ResourceRef),
			}
		}

	}

	return target
}

// Clone function
func (m *UpstreamSpec) Clone() proto.Message {
	var target *UpstreamSpec
	if m == nil {
		return target
	}
	target = &UpstreamSpec{}

	switch m.Llm.(type) {

	case *UpstreamSpec_Openai:

		if h, ok := interface{}(m.GetOpenai()).(clone.Cloner); ok {
			target.Llm = &UpstreamSpec_Openai{
				Openai: h.Clone().(*UpstreamSpec_OpenAI),
			}
		} else {
			target.Llm = &UpstreamSpec_Openai{
				Openai: proto.Clone(m.GetOpenai()).(*UpstreamSpec_OpenAI),
			}
		}

	case *UpstreamSpec_Mistral_:

		if h, ok := interface{}(m.GetMistral()).(clone.Cloner); ok {
			target.Llm = &UpstreamSpec_Mistral_{
				Mistral: h.Clone().(*UpstreamSpec_Mistral),
			}
		} else {
			target.Llm = &UpstreamSpec_Mistral_{
				Mistral: proto.Clone(m.GetMistral()).(*UpstreamSpec_Mistral),
			}
		}

	case *UpstreamSpec_Anthropic_:

		if h, ok := interface{}(m.GetAnthropic()).(clone.Cloner); ok {
			target.Llm = &UpstreamSpec_Anthropic_{
				Anthropic: h.Clone().(*UpstreamSpec_Anthropic),
			}
		} else {
			target.Llm = &UpstreamSpec_Anthropic_{
				Anthropic: proto.Clone(m.GetAnthropic()).(*UpstreamSpec_Anthropic),
			}
		}

	case *UpstreamSpec_AzureOpenai:

		if h, ok := interface{}(m.GetAzureOpenai()).(clone.Cloner); ok {
			target.Llm = &UpstreamSpec_AzureOpenai{
				AzureOpenai: h.Clone().(*UpstreamSpec_AzureOpenAI),
			}
		} else {
			target.Llm = &UpstreamSpec_AzureOpenai{
				AzureOpenai: proto.Clone(m.GetAzureOpenai()).(*UpstreamSpec_AzureOpenAI),
			}
		}

	case *UpstreamSpec_Multi:

		if h, ok := interface{}(m.GetMulti()).(clone.Cloner); ok {
			target.Llm = &UpstreamSpec_Multi{
				Multi: h.Clone().(*UpstreamSpec_MultiPool),
			}
		} else {
			target.Llm = &UpstreamSpec_Multi{
				Multi: proto.Clone(m.GetMulti()).(*UpstreamSpec_MultiPool),
			}
		}

	}

	return target
}

// Clone function
func (m *RouteSettings) Clone() proto.Message {
	var target *RouteSettings
	if m == nil {
		return target
	}
	target = &RouteSettings{}

	if h, ok := interface{}(m.GetPromptEnrichment()).(clone.Cloner); ok {
		target.PromptEnrichment = h.Clone().(*AIPromptEnrichment)
	} else {
		target.PromptEnrichment = proto.Clone(m.GetPromptEnrichment()).(*AIPromptEnrichment)
	}

	if h, ok := interface{}(m.GetPromptGuard()).(clone.Cloner); ok {
		target.PromptGuard = h.Clone().(*AIPromptGaurd)
	} else {
		target.PromptGuard = proto.Clone(m.GetPromptGuard()).(*AIPromptGaurd)
	}

	if h, ok := interface{}(m.GetRag()).(clone.Cloner); ok {
		target.Rag = h.Clone().(*RAG)
	} else {
		target.Rag = proto.Clone(m.GetRag()).(*RAG)
	}

	if h, ok := interface{}(m.GetSemanticCache()).(clone.Cloner); ok {
		target.SemanticCache = h.Clone().(*SemanticCache)
	} else {
		target.SemanticCache = proto.Clone(m.GetSemanticCache()).(*SemanticCache)
	}

	if m.GetDefaults() != nil {
		target.Defaults = make([]*FieldDefault, len(m.GetDefaults()))
		for idx, v := range m.GetDefaults() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Defaults[idx] = h.Clone().(*FieldDefault)
			} else {
				target.Defaults[idx] = proto.Clone(v).(*FieldDefault)
			}

		}
	}

	target.RouteType = m.GetRouteType()

	return target
}

// Clone function
func (m *FieldDefault) Clone() proto.Message {
	var target *FieldDefault
	if m == nil {
		return target
	}
	target = &FieldDefault{}

	target.Field = m.GetField()

	if h, ok := interface{}(m.GetValue()).(clone.Cloner); ok {
		target.Value = h.Clone().(*google_golang_org_protobuf_types_known_structpb.Value)
	} else {
		target.Value = proto.Clone(m.GetValue()).(*google_golang_org_protobuf_types_known_structpb.Value)
	}

	target.Override = m.GetOverride()

	return target
}

// Clone function
func (m *Postgres) Clone() proto.Message {
	var target *Postgres
	if m == nil {
		return target
	}
	target = &Postgres{}

	target.ConnectionString = m.GetConnectionString()

	target.CollectionName = m.GetCollectionName()

	return target
}

// Clone function
func (m *Embedding) Clone() proto.Message {
	var target *Embedding
	if m == nil {
		return target
	}
	target = &Embedding{}

	switch m.Embedding.(type) {

	case *Embedding_Openai:

		if h, ok := interface{}(m.GetOpenai()).(clone.Cloner); ok {
			target.Embedding = &Embedding_Openai{
				Openai: h.Clone().(*Embedding_OpenAI),
			}
		} else {
			target.Embedding = &Embedding_Openai{
				Openai: proto.Clone(m.GetOpenai()).(*Embedding_OpenAI),
			}
		}

	case *Embedding_AzureOpenai:

		if h, ok := interface{}(m.GetAzureOpenai()).(clone.Cloner); ok {
			target.Embedding = &Embedding_AzureOpenai{
				AzureOpenai: h.Clone().(*Embedding_AzureOpenAI),
			}
		} else {
			target.Embedding = &Embedding_AzureOpenai{
				AzureOpenai: proto.Clone(m.GetAzureOpenai()).(*Embedding_AzureOpenAI),
			}
		}

	}

	return target
}

// Clone function
func (m *SemanticCache) Clone() proto.Message {
	var target *SemanticCache
	if m == nil {
		return target
	}
	target = &SemanticCache{}

	if h, ok := interface{}(m.GetDatastore()).(clone.Cloner); ok {
		target.Datastore = h.Clone().(*SemanticCache_DataStore)
	} else {
		target.Datastore = proto.Clone(m.GetDatastore()).(*SemanticCache_DataStore)
	}

	if h, ok := interface{}(m.GetEmbedding()).(clone.Cloner); ok {
		target.Embedding = h.Clone().(*Embedding)
	} else {
		target.Embedding = proto.Clone(m.GetEmbedding()).(*Embedding)
	}

	target.Ttl = m.GetTtl()

	target.Mode = m.GetMode()

	return target
}

// Clone function
func (m *RAG) Clone() proto.Message {
	var target *RAG
	if m == nil {
		return target
	}
	target = &RAG{}

	if h, ok := interface{}(m.GetDatastore()).(clone.Cloner); ok {
		target.Datastore = h.Clone().(*RAG_DataStore)
	} else {
		target.Datastore = proto.Clone(m.GetDatastore()).(*RAG_DataStore)
	}

	if h, ok := interface{}(m.GetEmbedding()).(clone.Cloner); ok {
		target.Embedding = h.Clone().(*Embedding)
	} else {
		target.Embedding = proto.Clone(m.GetEmbedding()).(*Embedding)
	}

	target.PromptTemplate = m.GetPromptTemplate()

	return target
}

// Clone function
func (m *RateLimiting) Clone() proto.Message {
	var target *RateLimiting
	if m == nil {
		return target
	}
	target = &RateLimiting{}

	if m.GetRateLimitConfigs() != nil {
		target.RateLimitConfigs = make([]string, len(m.GetRateLimitConfigs()))
		for idx, v := range m.GetRateLimitConfigs() {

			target.RateLimitConfigs[idx] = v

		}
	}

	return target
}

// Clone function
func (m *AIPromptEnrichment) Clone() proto.Message {
	var target *AIPromptEnrichment
	if m == nil {
		return target
	}
	target = &AIPromptEnrichment{}

	if m.GetPrepend() != nil {
		target.Prepend = make([]*AIPromptEnrichment_Message, len(m.GetPrepend()))
		for idx, v := range m.GetPrepend() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Prepend[idx] = h.Clone().(*AIPromptEnrichment_Message)
			} else {
				target.Prepend[idx] = proto.Clone(v).(*AIPromptEnrichment_Message)
			}

		}
	}

	if m.GetAppend() != nil {
		target.Append = make([]*AIPromptEnrichment_Message, len(m.GetAppend()))
		for idx, v := range m.GetAppend() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Append[idx] = h.Clone().(*AIPromptEnrichment_Message)
			} else {
				target.Append[idx] = proto.Clone(v).(*AIPromptEnrichment_Message)
			}

		}
	}

	return target
}

// Clone function
func (m *AIPromptGaurd) Clone() proto.Message {
	var target *AIPromptGaurd
	if m == nil {
		return target
	}
	target = &AIPromptGaurd{}

	if h, ok := interface{}(m.GetRequest()).(clone.Cloner); ok {
		target.Request = h.Clone().(*AIPromptGaurd_Request)
	} else {
		target.Request = proto.Clone(m.GetRequest()).(*AIPromptGaurd_Request)
	}

	if h, ok := interface{}(m.GetResponse()).(clone.Cloner); ok {
		target.Response = h.Clone().(*AIPromptGaurd_Response)
	} else {
		target.Response = proto.Clone(m.GetResponse()).(*AIPromptGaurd_Response)
	}

	return target
}

// Clone function
func (m *UpstreamSpec_CustomHost) Clone() proto.Message {
	var target *UpstreamSpec_CustomHost
	if m == nil {
		return target
	}
	target = &UpstreamSpec_CustomHost{}

	target.Host = m.GetHost()

	target.Port = m.GetPort()

	return target
}

// Clone function
func (m *UpstreamSpec_OpenAI) Clone() proto.Message {
	var target *UpstreamSpec_OpenAI
	if m == nil {
		return target
	}
	target = &UpstreamSpec_OpenAI{}

	if h, ok := interface{}(m.GetAuthToken()).(clone.Cloner); ok {
		target.AuthToken = h.Clone().(*SingleAuthToken)
	} else {
		target.AuthToken = proto.Clone(m.GetAuthToken()).(*SingleAuthToken)
	}

	if h, ok := interface{}(m.GetCustomHost()).(clone.Cloner); ok {
		target.CustomHost = h.Clone().(*UpstreamSpec_CustomHost)
	} else {
		target.CustomHost = proto.Clone(m.GetCustomHost()).(*UpstreamSpec_CustomHost)
	}

	target.Model = m.GetModel()

	return target
}

// Clone function
func (m *UpstreamSpec_AzureOpenAI) Clone() proto.Message {
	var target *UpstreamSpec_AzureOpenAI
	if m == nil {
		return target
	}
	target = &UpstreamSpec_AzureOpenAI{}

	target.Endpoint = m.GetEndpoint()

	target.DeploymentName = m.GetDeploymentName()

	target.ApiVersion = m.GetApiVersion()

	switch m.AuthTokenSource.(type) {

	case *UpstreamSpec_AzureOpenAI_AuthToken:

		if h, ok := interface{}(m.GetAuthToken()).(clone.Cloner); ok {
			target.AuthTokenSource = &UpstreamSpec_AzureOpenAI_AuthToken{
				AuthToken: h.Clone().(*SingleAuthToken),
			}
		} else {
			target.AuthTokenSource = &UpstreamSpec_AzureOpenAI_AuthToken{
				AuthToken: proto.Clone(m.GetAuthToken()).(*SingleAuthToken),
			}
		}

	}

	return target
}

// Clone function
func (m *UpstreamSpec_Mistral) Clone() proto.Message {
	var target *UpstreamSpec_Mistral
	if m == nil {
		return target
	}
	target = &UpstreamSpec_Mistral{}

	if h, ok := interface{}(m.GetAuthToken()).(clone.Cloner); ok {
		target.AuthToken = h.Clone().(*SingleAuthToken)
	} else {
		target.AuthToken = proto.Clone(m.GetAuthToken()).(*SingleAuthToken)
	}

	if h, ok := interface{}(m.GetCustomHost()).(clone.Cloner); ok {
		target.CustomHost = h.Clone().(*UpstreamSpec_CustomHost)
	} else {
		target.CustomHost = proto.Clone(m.GetCustomHost()).(*UpstreamSpec_CustomHost)
	}

	target.Model = m.GetModel()

	return target
}

// Clone function
func (m *UpstreamSpec_Anthropic) Clone() proto.Message {
	var target *UpstreamSpec_Anthropic
	if m == nil {
		return target
	}
	target = &UpstreamSpec_Anthropic{}

	if h, ok := interface{}(m.GetAuthToken()).(clone.Cloner); ok {
		target.AuthToken = h.Clone().(*SingleAuthToken)
	} else {
		target.AuthToken = proto.Clone(m.GetAuthToken()).(*SingleAuthToken)
	}

	if h, ok := interface{}(m.GetCustomHost()).(clone.Cloner); ok {
		target.CustomHost = h.Clone().(*UpstreamSpec_CustomHost)
	} else {
		target.CustomHost = proto.Clone(m.GetCustomHost()).(*UpstreamSpec_CustomHost)
	}

	target.Version = m.GetVersion()

	target.Model = m.GetModel()

	return target
}

// Clone function
func (m *UpstreamSpec_MultiPool) Clone() proto.Message {
	var target *UpstreamSpec_MultiPool
	if m == nil {
		return target
	}
	target = &UpstreamSpec_MultiPool{}

	if m.GetPriorities() != nil {
		target.Priorities = make([]*UpstreamSpec_MultiPool_Priority, len(m.GetPriorities()))
		for idx, v := range m.GetPriorities() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Priorities[idx] = h.Clone().(*UpstreamSpec_MultiPool_Priority)
			} else {
				target.Priorities[idx] = proto.Clone(v).(*UpstreamSpec_MultiPool_Priority)
			}

		}
	}

	return target
}

// Clone function
func (m *UpstreamSpec_MultiPool_Backend) Clone() proto.Message {
	var target *UpstreamSpec_MultiPool_Backend
	if m == nil {
		return target
	}
	target = &UpstreamSpec_MultiPool_Backend{}

	switch m.Llm.(type) {

	case *UpstreamSpec_MultiPool_Backend_Openai:

		if h, ok := interface{}(m.GetOpenai()).(clone.Cloner); ok {
			target.Llm = &UpstreamSpec_MultiPool_Backend_Openai{
				Openai: h.Clone().(*UpstreamSpec_OpenAI),
			}
		} else {
			target.Llm = &UpstreamSpec_MultiPool_Backend_Openai{
				Openai: proto.Clone(m.GetOpenai()).(*UpstreamSpec_OpenAI),
			}
		}

	case *UpstreamSpec_MultiPool_Backend_Mistral:

		if h, ok := interface{}(m.GetMistral()).(clone.Cloner); ok {
			target.Llm = &UpstreamSpec_MultiPool_Backend_Mistral{
				Mistral: h.Clone().(*UpstreamSpec_Mistral),
			}
		} else {
			target.Llm = &UpstreamSpec_MultiPool_Backend_Mistral{
				Mistral: proto.Clone(m.GetMistral()).(*UpstreamSpec_Mistral),
			}
		}

	case *UpstreamSpec_MultiPool_Backend_Anthropic:

		if h, ok := interface{}(m.GetAnthropic()).(clone.Cloner); ok {
			target.Llm = &UpstreamSpec_MultiPool_Backend_Anthropic{
				Anthropic: h.Clone().(*UpstreamSpec_Anthropic),
			}
		} else {
			target.Llm = &UpstreamSpec_MultiPool_Backend_Anthropic{
				Anthropic: proto.Clone(m.GetAnthropic()).(*UpstreamSpec_Anthropic),
			}
		}

	case *UpstreamSpec_MultiPool_Backend_AzureOpenai:

		if h, ok := interface{}(m.GetAzureOpenai()).(clone.Cloner); ok {
			target.Llm = &UpstreamSpec_MultiPool_Backend_AzureOpenai{
				AzureOpenai: h.Clone().(*UpstreamSpec_AzureOpenAI),
			}
		} else {
			target.Llm = &UpstreamSpec_MultiPool_Backend_AzureOpenai{
				AzureOpenai: proto.Clone(m.GetAzureOpenai()).(*UpstreamSpec_AzureOpenAI),
			}
		}

	}

	return target
}

// Clone function
func (m *UpstreamSpec_MultiPool_Priority) Clone() proto.Message {
	var target *UpstreamSpec_MultiPool_Priority
	if m == nil {
		return target
	}
	target = &UpstreamSpec_MultiPool_Priority{}

	if m.GetPool() != nil {
		target.Pool = make([]*UpstreamSpec_MultiPool_Backend, len(m.GetPool()))
		for idx, v := range m.GetPool() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Pool[idx] = h.Clone().(*UpstreamSpec_MultiPool_Backend)
			} else {
				target.Pool[idx] = proto.Clone(v).(*UpstreamSpec_MultiPool_Backend)
			}

		}
	}

	target.Priority = m.GetPriority()

	return target
}

// Clone function
func (m *Embedding_OpenAI) Clone() proto.Message {
	var target *Embedding_OpenAI
	if m == nil {
		return target
	}
	target = &Embedding_OpenAI{}

	switch m.AuthTokenSource.(type) {

	case *Embedding_OpenAI_AuthToken:

		if h, ok := interface{}(m.GetAuthToken()).(clone.Cloner); ok {
			target.AuthTokenSource = &Embedding_OpenAI_AuthToken{
				AuthToken: h.Clone().(*SingleAuthToken),
			}
		} else {
			target.AuthTokenSource = &Embedding_OpenAI_AuthToken{
				AuthToken: proto.Clone(m.GetAuthToken()).(*SingleAuthToken),
			}
		}

	}

	return target
}

// Clone function
func (m *Embedding_AzureOpenAI) Clone() proto.Message {
	var target *Embedding_AzureOpenAI
	if m == nil {
		return target
	}
	target = &Embedding_AzureOpenAI{}

	target.ApiVersion = m.GetApiVersion()

	target.Endpoint = m.GetEndpoint()

	target.DeploymentName = m.GetDeploymentName()

	switch m.AuthTokenSource.(type) {

	case *Embedding_AzureOpenAI_AuthToken:

		if h, ok := interface{}(m.GetAuthToken()).(clone.Cloner); ok {
			target.AuthTokenSource = &Embedding_AzureOpenAI_AuthToken{
				AuthToken: h.Clone().(*SingleAuthToken),
			}
		} else {
			target.AuthTokenSource = &Embedding_AzureOpenAI_AuthToken{
				AuthToken: proto.Clone(m.GetAuthToken()).(*SingleAuthToken),
			}
		}

	}

	return target
}

// Clone function
func (m *SemanticCache_Redis) Clone() proto.Message {
	var target *SemanticCache_Redis
	if m == nil {
		return target
	}
	target = &SemanticCache_Redis{}

	target.ConnectionString = m.GetConnectionString()

	target.ScoreThreshold = m.GetScoreThreshold()

	return target
}

// Clone function
func (m *SemanticCache_Weaviate) Clone() proto.Message {
	var target *SemanticCache_Weaviate
	if m == nil {
		return target
	}
	target = &SemanticCache_Weaviate{}

	target.Host = m.GetHost()

	target.HttpPort = m.GetHttpPort()

	target.GrpcPort = m.GetGrpcPort()

	target.Insecure = m.GetInsecure()

	return target
}

// Clone function
func (m *SemanticCache_DataStore) Clone() proto.Message {
	var target *SemanticCache_DataStore
	if m == nil {
		return target
	}
	target = &SemanticCache_DataStore{}

	switch m.Datastore.(type) {

	case *SemanticCache_DataStore_Redis:

		if h, ok := interface{}(m.GetRedis()).(clone.Cloner); ok {
			target.Datastore = &SemanticCache_DataStore_Redis{
				Redis: h.Clone().(*SemanticCache_Redis),
			}
		} else {
			target.Datastore = &SemanticCache_DataStore_Redis{
				Redis: proto.Clone(m.GetRedis()).(*SemanticCache_Redis),
			}
		}

	case *SemanticCache_DataStore_Weaviate:

		if h, ok := interface{}(m.GetWeaviate()).(clone.Cloner); ok {
			target.Datastore = &SemanticCache_DataStore_Weaviate{
				Weaviate: h.Clone().(*SemanticCache_Weaviate),
			}
		} else {
			target.Datastore = &SemanticCache_DataStore_Weaviate{
				Weaviate: proto.Clone(m.GetWeaviate()).(*SemanticCache_Weaviate),
			}
		}

	}

	return target
}

// Clone function
func (m *RAG_DataStore) Clone() proto.Message {
	var target *RAG_DataStore
	if m == nil {
		return target
	}
	target = &RAG_DataStore{}

	switch m.Datastore.(type) {

	case *RAG_DataStore_Postgres:

		if h, ok := interface{}(m.GetPostgres()).(clone.Cloner); ok {
			target.Datastore = &RAG_DataStore_Postgres{
				Postgres: h.Clone().(*Postgres),
			}
		} else {
			target.Datastore = &RAG_DataStore_Postgres{
				Postgres: proto.Clone(m.GetPostgres()).(*Postgres),
			}
		}

	}

	return target
}

// Clone function
func (m *AIPromptEnrichment_Message) Clone() proto.Message {
	var target *AIPromptEnrichment_Message
	if m == nil {
		return target
	}
	target = &AIPromptEnrichment_Message{}

	target.Role = m.GetRole()

	target.Content = m.GetContent()

	return target
}

// Clone function
func (m *AIPromptGaurd_Request) Clone() proto.Message {
	var target *AIPromptGaurd_Request
	if m == nil {
		return target
	}
	target = &AIPromptGaurd_Request{}

	if m.GetMatches() != nil {
		target.Matches = make([]string, len(m.GetMatches()))
		for idx, v := range m.GetMatches() {

			target.Matches[idx] = v

		}
	}

	target.CustomResponseMessage = m.GetCustomResponseMessage()

	return target
}

// Clone function
func (m *AIPromptGaurd_Response) Clone() proto.Message {
	var target *AIPromptGaurd_Response
	if m == nil {
		return target
	}
	target = &AIPromptGaurd_Response{}

	if m.GetMatches() != nil {
		target.Matches = make([]string, len(m.GetMatches()))
		for idx, v := range m.GetMatches() {

			target.Matches[idx] = v

		}
	}

	if m.GetBuiltins() != nil {
		target.Builtins = make([]AIPromptGaurd_Response_BuiltIn, len(m.GetBuiltins()))
		for idx, v := range m.GetBuiltins() {

			target.Builtins[idx] = v

		}
	}

	return target
}
