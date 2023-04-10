// Copyright 2023 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package phishingprotection

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"net/url"
	"time"

	phishingprotectionpb "cloud.google.com/go/phishingprotection/apiv1beta1/phishingprotectionpb"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	httptransport "google.golang.org/api/transport/http"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/encoding/protojson"
)

var newPhishingProtectionServiceV1Beta1ClientHook clientHook

// PhishingProtectionServiceV1Beta1CallOptions contains the retry settings for each method of PhishingProtectionServiceV1Beta1Client.
type PhishingProtectionServiceV1Beta1CallOptions struct {
	ReportPhishing []gax.CallOption
}

func defaultPhishingProtectionServiceV1Beta1GRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("phishingprotection.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("phishingprotection.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://phishingprotection.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultPhishingProtectionServiceV1Beta1CallOptions() *PhishingProtectionServiceV1Beta1CallOptions {
	return &PhishingProtectionServiceV1Beta1CallOptions{
		ReportPhishing: []gax.CallOption{},
	}
}

func defaultPhishingProtectionServiceV1Beta1RESTCallOptions() *PhishingProtectionServiceV1Beta1CallOptions {
	return &PhishingProtectionServiceV1Beta1CallOptions{
		ReportPhishing: []gax.CallOption{},
	}
}

// internalPhishingProtectionServiceV1Beta1Client is an interface that defines the methods available from Phishing Protection API.
type internalPhishingProtectionServiceV1Beta1Client interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	ReportPhishing(context.Context, *phishingprotectionpb.ReportPhishingRequest, ...gax.CallOption) (*phishingprotectionpb.ReportPhishingResponse, error)
}

// PhishingProtectionServiceV1Beta1Client is a client for interacting with Phishing Protection API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service to report phishing URIs.
type PhishingProtectionServiceV1Beta1Client struct {
	// The internal transport-dependent client.
	internalClient internalPhishingProtectionServiceV1Beta1Client

	// The call options for this service.
	CallOptions *PhishingProtectionServiceV1Beta1CallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *PhishingProtectionServiceV1Beta1Client) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *PhishingProtectionServiceV1Beta1Client) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *PhishingProtectionServiceV1Beta1Client) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// ReportPhishing reports a URI suspected of containing phishing content to be reviewed. Once
// the report review is complete, its result can be found in the Cloud
// Security Command Center findings dashboard for Phishing Protection. If the
// result verifies the existence of malicious phishing content, the site will
// be added the to Google’s Social Engineering
// lists (at https://support.google.com/webmasters/answer/6350487/) in order to
// protect users that could get exposed to this threat in the future.
func (c *PhishingProtectionServiceV1Beta1Client) ReportPhishing(ctx context.Context, req *phishingprotectionpb.ReportPhishingRequest, opts ...gax.CallOption) (*phishingprotectionpb.ReportPhishingResponse, error) {
	return c.internalClient.ReportPhishing(ctx, req, opts...)
}

// phishingProtectionServiceV1Beta1GRPCClient is a client for interacting with Phishing Protection API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type phishingProtectionServiceV1Beta1GRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing PhishingProtectionServiceV1Beta1Client
	CallOptions **PhishingProtectionServiceV1Beta1CallOptions

	// The gRPC API client.
	phishingProtectionServiceV1Beta1Client phishingprotectionpb.PhishingProtectionServiceV1Beta1Client

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewPhishingProtectionServiceV1Beta1Client creates a new phishing protection service v1 beta1 client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service to report phishing URIs.
func NewPhishingProtectionServiceV1Beta1Client(ctx context.Context, opts ...option.ClientOption) (*PhishingProtectionServiceV1Beta1Client, error) {
	clientOpts := defaultPhishingProtectionServiceV1Beta1GRPCClientOptions()
	if newPhishingProtectionServiceV1Beta1ClientHook != nil {
		hookOpts, err := newPhishingProtectionServiceV1Beta1ClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	disableDeadlines, err := checkDisableDeadlines()
	if err != nil {
		return nil, err
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := PhishingProtectionServiceV1Beta1Client{CallOptions: defaultPhishingProtectionServiceV1Beta1CallOptions()}

	c := &phishingProtectionServiceV1Beta1GRPCClient{
		connPool:                               connPool,
		disableDeadlines:                       disableDeadlines,
		phishingProtectionServiceV1Beta1Client: phishingprotectionpb.NewPhishingProtectionServiceV1Beta1Client(connPool),
		CallOptions:                            &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *phishingProtectionServiceV1Beta1GRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *phishingProtectionServiceV1Beta1GRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *phishingProtectionServiceV1Beta1GRPCClient) Close() error {
	return c.connPool.Close()
}

// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type phishingProtectionServiceV1Beta1RESTClient struct {
	// The http endpoint to connect to.
	endpoint string

	// The http client.
	httpClient *http.Client

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD

	// Points back to the CallOptions field of the containing PhishingProtectionServiceV1Beta1Client
	CallOptions **PhishingProtectionServiceV1Beta1CallOptions
}

// NewPhishingProtectionServiceV1Beta1RESTClient creates a new phishing protection service v1 beta1 rest client.
//
// Service to report phishing URIs.
func NewPhishingProtectionServiceV1Beta1RESTClient(ctx context.Context, opts ...option.ClientOption) (*PhishingProtectionServiceV1Beta1Client, error) {
	clientOpts := append(defaultPhishingProtectionServiceV1Beta1RESTClientOptions(), opts...)
	httpClient, endpoint, err := httptransport.NewClient(ctx, clientOpts...)
	if err != nil {
		return nil, err
	}

	callOpts := defaultPhishingProtectionServiceV1Beta1RESTCallOptions()
	c := &phishingProtectionServiceV1Beta1RESTClient{
		endpoint:    endpoint,
		httpClient:  httpClient,
		CallOptions: &callOpts,
	}
	c.setGoogleClientInfo()

	return &PhishingProtectionServiceV1Beta1Client{internalClient: c, CallOptions: callOpts}, nil
}

func defaultPhishingProtectionServiceV1Beta1RESTClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("https://phishingprotection.googleapis.com"),
		internaloption.WithDefaultMTLSEndpoint("https://phishingprotection.mtls.googleapis.com"),
		internaloption.WithDefaultAudience("https://phishingprotection.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
	}
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *phishingProtectionServiceV1Beta1RESTClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "rest", "UNKNOWN")
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *phishingProtectionServiceV1Beta1RESTClient) Close() error {
	// Replace httpClient with nil to force cleanup.
	c.httpClient = nil
	return nil
}

// Connection returns a connection to the API service.
//
// Deprecated: This method always returns nil.
func (c *phishingProtectionServiceV1Beta1RESTClient) Connection() *grpc.ClientConn {
	return nil
}
func (c *phishingProtectionServiceV1Beta1GRPCClient) ReportPhishing(ctx context.Context, req *phishingprotectionpb.ReportPhishingRequest, opts ...gax.CallOption) (*phishingprotectionpb.ReportPhishingResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ReportPhishing[0:len((*c.CallOptions).ReportPhishing):len((*c.CallOptions).ReportPhishing)], opts...)
	var resp *phishingprotectionpb.ReportPhishingResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.phishingProtectionServiceV1Beta1Client.ReportPhishing(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// ReportPhishing reports a URI suspected of containing phishing content to be reviewed. Once
// the report review is complete, its result can be found in the Cloud
// Security Command Center findings dashboard for Phishing Protection. If the
// result verifies the existence of malicious phishing content, the site will
// be added the to Google’s Social Engineering
// lists (at https://support.google.com/webmasters/answer/6350487/) in order to
// protect users that could get exposed to this threat in the future.
func (c *phishingProtectionServiceV1Beta1RESTClient) ReportPhishing(ctx context.Context, req *phishingprotectionpb.ReportPhishingRequest, opts ...gax.CallOption) (*phishingprotectionpb.ReportPhishingResponse, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	jsonReq, err := m.Marshal(req)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1beta1/%v/phishing:report", req.GetParent())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))

	headers := buildHeaders(ctx, c.xGoogMetadata, md, metadata.Pairs("Content-Type", "application/json"))
	opts = append((*c.CallOptions).ReportPhishing[0:len((*c.CallOptions).ReportPhishing):len((*c.CallOptions).ReportPhishing)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &phishingprotectionpb.ReportPhishingResponse{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return err
		}

		buf, err := ioutil.ReadAll(httpRsp.Body)
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return maybeUnknownEnum(err)
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	return resp, nil
}
