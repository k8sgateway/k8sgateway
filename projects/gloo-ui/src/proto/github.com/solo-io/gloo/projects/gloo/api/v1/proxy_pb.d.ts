// package: gloo.solo.io
// file: github.com/solo-io/gloo/projects/gloo/api/v1/proxy.proto

import * as jspb from "google-protobuf";
import * as google_protobuf_wrappers_pb from "google-protobuf/google/protobuf/wrappers_pb";
import * as gogoproto_gogo_pb from "../../../../../../../gogoproto/gogo_pb";
import * as github_com_solo_io_solo_kit_api_v1_metadata_pb from "../../../../../../../github.com/solo-io/solo-kit/api/v1/metadata_pb";
import * as github_com_solo_io_solo_kit_api_v1_status_pb from "../../../../../../../github.com/solo-io/solo-kit/api/v1/status_pb";
import * as github_com_solo_io_solo_kit_api_v1_ref_pb from "../../../../../../../github.com/solo-io/solo-kit/api/v1/ref_pb";
import * as github_com_solo_io_gloo_projects_gloo_api_v1_ssl_pb from "../../../../../../../github.com/solo-io/gloo/projects/gloo/api/v1/ssl_pb";
import * as github_com_solo_io_gloo_projects_gloo_api_v1_subset_pb from "../../../../../../../github.com/solo-io/gloo/projects/gloo/api/v1/subset_pb";
import * as github_com_solo_io_gloo_projects_gloo_api_v1_plugins_pb from "../../../../../../../github.com/solo-io/gloo/projects/gloo/api/v1/plugins_pb";

export class Proxy extends jspb.Message {
  clearListenersList(): void;
  getListenersList(): Array<Listener>;
  setListenersList(value: Array<Listener>): void;
  addListeners(value?: Listener, index?: number): Listener;

  hasStatus(): boolean;
  clearStatus(): void;
  getStatus(): github_com_solo_io_solo_kit_api_v1_status_pb.Status | undefined;
  setStatus(value?: github_com_solo_io_solo_kit_api_v1_status_pb.Status): void;

  hasMetadata(): boolean;
  clearMetadata(): void;
  getMetadata(): github_com_solo_io_solo_kit_api_v1_metadata_pb.Metadata | undefined;
  setMetadata(value?: github_com_solo_io_solo_kit_api_v1_metadata_pb.Metadata): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Proxy.AsObject;
  static toObject(includeInstance: boolean, msg: Proxy): Proxy.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Proxy, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Proxy;
  static deserializeBinaryFromReader(message: Proxy, reader: jspb.BinaryReader): Proxy;
}

export namespace Proxy {
  export type AsObject = {
    listenersList: Array<Listener.AsObject>,
    status?: github_com_solo_io_solo_kit_api_v1_status_pb.Status.AsObject,
    metadata?: github_com_solo_io_solo_kit_api_v1_metadata_pb.Metadata.AsObject,
  }
}

export class Listener extends jspb.Message {
  getName(): string;
  setName(value: string): void;

  getBindAddress(): string;
  setBindAddress(value: string): void;

  getBindPort(): number;
  setBindPort(value: number): void;

  hasHttpListener(): boolean;
  clearHttpListener(): void;
  getHttpListener(): HttpListener | undefined;
  setHttpListener(value?: HttpListener): void;

  clearSslConfigurationsList(): void;
  getSslConfigurationsList(): Array<github_com_solo_io_gloo_projects_gloo_api_v1_ssl_pb.SslConfig>;
  setSslConfigurationsList(value: Array<github_com_solo_io_gloo_projects_gloo_api_v1_ssl_pb.SslConfig>): void;
  addSslConfigurations(value?: github_com_solo_io_gloo_projects_gloo_api_v1_ssl_pb.SslConfig, index?: number): github_com_solo_io_gloo_projects_gloo_api_v1_ssl_pb.SslConfig;

  hasUseProxyProto(): boolean;
  clearUseProxyProto(): void;
  getUseProxyProto(): google_protobuf_wrappers_pb.BoolValue | undefined;
  setUseProxyProto(value?: google_protobuf_wrappers_pb.BoolValue): void;

  getListenertypeCase(): Listener.ListenertypeCase;
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Listener.AsObject;
  static toObject(includeInstance: boolean, msg: Listener): Listener.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Listener, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Listener;
  static deserializeBinaryFromReader(message: Listener, reader: jspb.BinaryReader): Listener;
}

export namespace Listener {
  export type AsObject = {
    name: string,
    bindAddress: string,
    bindPort: number,
    httpListener?: HttpListener.AsObject,
    sslConfigurationsList: Array<github_com_solo_io_gloo_projects_gloo_api_v1_ssl_pb.SslConfig.AsObject>,
    useProxyProto?: google_protobuf_wrappers_pb.BoolValue.AsObject,
  }

  export enum ListenertypeCase {
    LISTENERTYPE_NOT_SET = 0,
    HTTP_LISTENER = 4,
  }
}

export class HttpListener extends jspb.Message {
  clearVirtualHostsList(): void;
  getVirtualHostsList(): Array<VirtualHost>;
  setVirtualHostsList(value: Array<VirtualHost>): void;
  addVirtualHosts(value?: VirtualHost, index?: number): VirtualHost;

  hasListenerPlugins(): boolean;
  clearListenerPlugins(): void;
  getListenerPlugins(): github_com_solo_io_gloo_projects_gloo_api_v1_plugins_pb.ListenerPlugins | undefined;
  setListenerPlugins(value?: github_com_solo_io_gloo_projects_gloo_api_v1_plugins_pb.ListenerPlugins): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): HttpListener.AsObject;
  static toObject(includeInstance: boolean, msg: HttpListener): HttpListener.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: HttpListener, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): HttpListener;
  static deserializeBinaryFromReader(message: HttpListener, reader: jspb.BinaryReader): HttpListener;
}

export namespace HttpListener {
  export type AsObject = {
    virtualHostsList: Array<VirtualHost.AsObject>,
    listenerPlugins?: github_com_solo_io_gloo_projects_gloo_api_v1_plugins_pb.ListenerPlugins.AsObject,
  }
}

export class VirtualHost extends jspb.Message {
  getName(): string;
  setName(value: string): void;

  clearDomainsList(): void;
  getDomainsList(): Array<string>;
  setDomainsList(value: Array<string>): void;
  addDomains(value: string, index?: number): string;

  clearRoutesList(): void;
  getRoutesList(): Array<Route>;
  setRoutesList(value: Array<Route>): void;
  addRoutes(value?: Route, index?: number): Route;

  hasVirtualHostPlugins(): boolean;
  clearVirtualHostPlugins(): void;
  getVirtualHostPlugins(): github_com_solo_io_gloo_projects_gloo_api_v1_plugins_pb.VirtualHostPlugins | undefined;
  setVirtualHostPlugins(value?: github_com_solo_io_gloo_projects_gloo_api_v1_plugins_pb.VirtualHostPlugins): void;

  hasCorsPolicy(): boolean;
  clearCorsPolicy(): void;
  getCorsPolicy(): CorsPolicy | undefined;
  setCorsPolicy(value?: CorsPolicy): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): VirtualHost.AsObject;
  static toObject(includeInstance: boolean, msg: VirtualHost): VirtualHost.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: VirtualHost, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): VirtualHost;
  static deserializeBinaryFromReader(message: VirtualHost, reader: jspb.BinaryReader): VirtualHost;
}

export namespace VirtualHost {
  export type AsObject = {
    name: string,
    domainsList: Array<string>,
    routesList: Array<Route.AsObject>,
    virtualHostPlugins?: github_com_solo_io_gloo_projects_gloo_api_v1_plugins_pb.VirtualHostPlugins.AsObject,
    corsPolicy?: CorsPolicy.AsObject,
  }
}

export class Route extends jspb.Message {
  hasMatcher(): boolean;
  clearMatcher(): void;
  getMatcher(): Matcher | undefined;
  setMatcher(value?: Matcher): void;

  hasRouteAction(): boolean;
  clearRouteAction(): void;
  getRouteAction(): RouteAction | undefined;
  setRouteAction(value?: RouteAction): void;

  hasRedirectAction(): boolean;
  clearRedirectAction(): void;
  getRedirectAction(): RedirectAction | undefined;
  setRedirectAction(value?: RedirectAction): void;

  hasDirectResponseAction(): boolean;
  clearDirectResponseAction(): void;
  getDirectResponseAction(): DirectResponseAction | undefined;
  setDirectResponseAction(value?: DirectResponseAction): void;

  hasRoutePlugins(): boolean;
  clearRoutePlugins(): void;
  getRoutePlugins(): github_com_solo_io_gloo_projects_gloo_api_v1_plugins_pb.RoutePlugins | undefined;
  setRoutePlugins(value?: github_com_solo_io_gloo_projects_gloo_api_v1_plugins_pb.RoutePlugins): void;

  getActionCase(): Route.ActionCase;
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Route.AsObject;
  static toObject(includeInstance: boolean, msg: Route): Route.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Route, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Route;
  static deserializeBinaryFromReader(message: Route, reader: jspb.BinaryReader): Route;
}

export namespace Route {
  export type AsObject = {
    matcher?: Matcher.AsObject,
    routeAction?: RouteAction.AsObject,
    redirectAction?: RedirectAction.AsObject,
    directResponseAction?: DirectResponseAction.AsObject,
    routePlugins?: github_com_solo_io_gloo_projects_gloo_api_v1_plugins_pb.RoutePlugins.AsObject,
  }

  export enum ActionCase {
    ACTION_NOT_SET = 0,
    ROUTE_ACTION = 2,
    REDIRECT_ACTION = 3,
    DIRECT_RESPONSE_ACTION = 4,
  }
}

export class Matcher extends jspb.Message {
  hasPrefix(): boolean;
  clearPrefix(): void;
  getPrefix(): string;
  setPrefix(value: string): void;

  hasExact(): boolean;
  clearExact(): void;
  getExact(): string;
  setExact(value: string): void;

  hasRegex(): boolean;
  clearRegex(): void;
  getRegex(): string;
  setRegex(value: string): void;

  clearHeadersList(): void;
  getHeadersList(): Array<HeaderMatcher>;
  setHeadersList(value: Array<HeaderMatcher>): void;
  addHeaders(value?: HeaderMatcher, index?: number): HeaderMatcher;

  clearQueryParametersList(): void;
  getQueryParametersList(): Array<QueryParameterMatcher>;
  setQueryParametersList(value: Array<QueryParameterMatcher>): void;
  addQueryParameters(value?: QueryParameterMatcher, index?: number): QueryParameterMatcher;

  clearMethodsList(): void;
  getMethodsList(): Array<string>;
  setMethodsList(value: Array<string>): void;
  addMethods(value: string, index?: number): string;

  getPathSpecifierCase(): Matcher.PathSpecifierCase;
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Matcher.AsObject;
  static toObject(includeInstance: boolean, msg: Matcher): Matcher.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Matcher, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Matcher;
  static deserializeBinaryFromReader(message: Matcher, reader: jspb.BinaryReader): Matcher;
}

export namespace Matcher {
  export type AsObject = {
    prefix: string,
    exact: string,
    regex: string,
    headersList: Array<HeaderMatcher.AsObject>,
    queryParametersList: Array<QueryParameterMatcher.AsObject>,
    methodsList: Array<string>,
  }

  export enum PathSpecifierCase {
    PATH_SPECIFIER_NOT_SET = 0,
    PREFIX = 1,
    EXACT = 2,
    REGEX = 3,
  }
}

export class HeaderMatcher extends jspb.Message {
  getName(): string;
  setName(value: string): void;

  getValue(): string;
  setValue(value: string): void;

  getRegex(): boolean;
  setRegex(value: boolean): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): HeaderMatcher.AsObject;
  static toObject(includeInstance: boolean, msg: HeaderMatcher): HeaderMatcher.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: HeaderMatcher, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): HeaderMatcher;
  static deserializeBinaryFromReader(message: HeaderMatcher, reader: jspb.BinaryReader): HeaderMatcher;
}

export namespace HeaderMatcher {
  export type AsObject = {
    name: string,
    value: string,
    regex: boolean,
  }
}

export class QueryParameterMatcher extends jspb.Message {
  getName(): string;
  setName(value: string): void;

  getValue(): string;
  setValue(value: string): void;

  getRegex(): boolean;
  setRegex(value: boolean): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): QueryParameterMatcher.AsObject;
  static toObject(includeInstance: boolean, msg: QueryParameterMatcher): QueryParameterMatcher.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: QueryParameterMatcher, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): QueryParameterMatcher;
  static deserializeBinaryFromReader(message: QueryParameterMatcher, reader: jspb.BinaryReader): QueryParameterMatcher;
}

export namespace QueryParameterMatcher {
  export type AsObject = {
    name: string,
    value: string,
    regex: boolean,
  }
}

export class RouteAction extends jspb.Message {
  hasSingle(): boolean;
  clearSingle(): void;
  getSingle(): Destination | undefined;
  setSingle(value?: Destination): void;

  hasMulti(): boolean;
  clearMulti(): void;
  getMulti(): MultiDestination | undefined;
  setMulti(value?: MultiDestination): void;

  hasUpstreamGroup(): boolean;
  clearUpstreamGroup(): void;
  getUpstreamGroup(): github_com_solo_io_solo_kit_api_v1_ref_pb.ResourceRef | undefined;
  setUpstreamGroup(value?: github_com_solo_io_solo_kit_api_v1_ref_pb.ResourceRef): void;

  getDestinationCase(): RouteAction.DestinationCase;
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RouteAction.AsObject;
  static toObject(includeInstance: boolean, msg: RouteAction): RouteAction.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: RouteAction, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RouteAction;
  static deserializeBinaryFromReader(message: RouteAction, reader: jspb.BinaryReader): RouteAction;
}

export namespace RouteAction {
  export type AsObject = {
    single?: Destination.AsObject,
    multi?: MultiDestination.AsObject,
    upstreamGroup?: github_com_solo_io_solo_kit_api_v1_ref_pb.ResourceRef.AsObject,
  }

  export enum DestinationCase {
    DESTINATION_NOT_SET = 0,
    SINGLE = 1,
    MULTI = 2,
    UPSTREAM_GROUP = 3,
  }
}

export class Destination extends jspb.Message {
  hasUpstream(): boolean;
  clearUpstream(): void;
  getUpstream(): github_com_solo_io_solo_kit_api_v1_ref_pb.ResourceRef | undefined;
  setUpstream(value?: github_com_solo_io_solo_kit_api_v1_ref_pb.ResourceRef): void;

  hasService(): boolean;
  clearService(): void;
  getService(): ServiceDestination | undefined;
  setService(value?: ServiceDestination): void;

  hasDestinationSpec(): boolean;
  clearDestinationSpec(): void;
  getDestinationSpec(): github_com_solo_io_gloo_projects_gloo_api_v1_plugins_pb.DestinationSpec | undefined;
  setDestinationSpec(value?: github_com_solo_io_gloo_projects_gloo_api_v1_plugins_pb.DestinationSpec): void;

  hasSubset(): boolean;
  clearSubset(): void;
  getSubset(): github_com_solo_io_gloo_projects_gloo_api_v1_subset_pb.Subset | undefined;
  setSubset(value?: github_com_solo_io_gloo_projects_gloo_api_v1_subset_pb.Subset): void;

  getDestinationTypeCase(): Destination.DestinationTypeCase;
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Destination.AsObject;
  static toObject(includeInstance: boolean, msg: Destination): Destination.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Destination, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Destination;
  static deserializeBinaryFromReader(message: Destination, reader: jspb.BinaryReader): Destination;
}

export namespace Destination {
  export type AsObject = {
    upstream?: github_com_solo_io_solo_kit_api_v1_ref_pb.ResourceRef.AsObject,
    service?: ServiceDestination.AsObject,
    destinationSpec?: github_com_solo_io_gloo_projects_gloo_api_v1_plugins_pb.DestinationSpec.AsObject,
    subset?: github_com_solo_io_gloo_projects_gloo_api_v1_subset_pb.Subset.AsObject,
  }

  export enum DestinationTypeCase {
    DESTINATION_TYPE_NOT_SET = 0,
    UPSTREAM = 10,
    SERVICE = 11,
  }
}

export class ServiceDestination extends jspb.Message {
  hasRef(): boolean;
  clearRef(): void;
  getRef(): github_com_solo_io_solo_kit_api_v1_ref_pb.ResourceRef | undefined;
  setRef(value?: github_com_solo_io_solo_kit_api_v1_ref_pb.ResourceRef): void;

  getPort(): number;
  setPort(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ServiceDestination.AsObject;
  static toObject(includeInstance: boolean, msg: ServiceDestination): ServiceDestination.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ServiceDestination, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ServiceDestination;
  static deserializeBinaryFromReader(message: ServiceDestination, reader: jspb.BinaryReader): ServiceDestination;
}

export namespace ServiceDestination {
  export type AsObject = {
    ref?: github_com_solo_io_solo_kit_api_v1_ref_pb.ResourceRef.AsObject,
    port: number,
  }
}

export class UpstreamGroup extends jspb.Message {
  clearDestinationsList(): void;
  getDestinationsList(): Array<WeightedDestination>;
  setDestinationsList(value: Array<WeightedDestination>): void;
  addDestinations(value?: WeightedDestination, index?: number): WeightedDestination;

  hasStatus(): boolean;
  clearStatus(): void;
  getStatus(): github_com_solo_io_solo_kit_api_v1_status_pb.Status | undefined;
  setStatus(value?: github_com_solo_io_solo_kit_api_v1_status_pb.Status): void;

  hasMetadata(): boolean;
  clearMetadata(): void;
  getMetadata(): github_com_solo_io_solo_kit_api_v1_metadata_pb.Metadata | undefined;
  setMetadata(value?: github_com_solo_io_solo_kit_api_v1_metadata_pb.Metadata): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpstreamGroup.AsObject;
  static toObject(includeInstance: boolean, msg: UpstreamGroup): UpstreamGroup.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UpstreamGroup, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpstreamGroup;
  static deserializeBinaryFromReader(message: UpstreamGroup, reader: jspb.BinaryReader): UpstreamGroup;
}

export namespace UpstreamGroup {
  export type AsObject = {
    destinationsList: Array<WeightedDestination.AsObject>,
    status?: github_com_solo_io_solo_kit_api_v1_status_pb.Status.AsObject,
    metadata?: github_com_solo_io_solo_kit_api_v1_metadata_pb.Metadata.AsObject,
  }
}

export class MultiDestination extends jspb.Message {
  clearDestinationsList(): void;
  getDestinationsList(): Array<WeightedDestination>;
  setDestinationsList(value: Array<WeightedDestination>): void;
  addDestinations(value?: WeightedDestination, index?: number): WeightedDestination;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): MultiDestination.AsObject;
  static toObject(includeInstance: boolean, msg: MultiDestination): MultiDestination.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: MultiDestination, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): MultiDestination;
  static deserializeBinaryFromReader(message: MultiDestination, reader: jspb.BinaryReader): MultiDestination;
}

export namespace MultiDestination {
  export type AsObject = {
    destinationsList: Array<WeightedDestination.AsObject>,
  }
}

export class WeightedDestination extends jspb.Message {
  hasDestination(): boolean;
  clearDestination(): void;
  getDestination(): Destination | undefined;
  setDestination(value?: Destination): void;

  getWeight(): number;
  setWeight(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): WeightedDestination.AsObject;
  static toObject(includeInstance: boolean, msg: WeightedDestination): WeightedDestination.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: WeightedDestination, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): WeightedDestination;
  static deserializeBinaryFromReader(message: WeightedDestination, reader: jspb.BinaryReader): WeightedDestination;
}

export namespace WeightedDestination {
  export type AsObject = {
    destination?: Destination.AsObject,
    weight: number,
  }
}

export class RedirectAction extends jspb.Message {
  getHostRedirect(): string;
  setHostRedirect(value: string): void;

  hasPathRedirect(): boolean;
  clearPathRedirect(): void;
  getPathRedirect(): string;
  setPathRedirect(value: string): void;

  hasPrefixRewrite(): boolean;
  clearPrefixRewrite(): void;
  getPrefixRewrite(): string;
  setPrefixRewrite(value: string): void;

  getResponseCode(): RedirectAction.RedirectResponseCode;
  setResponseCode(value: RedirectAction.RedirectResponseCode): void;

  getHttpsRedirect(): boolean;
  setHttpsRedirect(value: boolean): void;

  getStripQuery(): boolean;
  setStripQuery(value: boolean): void;

  getPathRewriteSpecifierCase(): RedirectAction.PathRewriteSpecifierCase;
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RedirectAction.AsObject;
  static toObject(includeInstance: boolean, msg: RedirectAction): RedirectAction.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: RedirectAction, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RedirectAction;
  static deserializeBinaryFromReader(message: RedirectAction, reader: jspb.BinaryReader): RedirectAction;
}

export namespace RedirectAction {
  export type AsObject = {
    hostRedirect: string,
    pathRedirect: string,
    prefixRewrite: string,
    responseCode: RedirectAction.RedirectResponseCode,
    httpsRedirect: boolean,
    stripQuery: boolean,
  }

  export enum RedirectResponseCode {
    MOVED_PERMANENTLY = 0,
    FOUND = 1,
    SEE_OTHER = 2,
    TEMPORARY_REDIRECT = 3,
    PERMANENT_REDIRECT = 4,
  }

  export enum PathRewriteSpecifierCase {
    PATH_REWRITE_SPECIFIER_NOT_SET = 0,
    PATH_REDIRECT = 2,
    PREFIX_REWRITE = 5,
  }
}

export class DirectResponseAction extends jspb.Message {
  getStatus(): number;
  setStatus(value: number): void;

  getBody(): string;
  setBody(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DirectResponseAction.AsObject;
  static toObject(includeInstance: boolean, msg: DirectResponseAction): DirectResponseAction.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: DirectResponseAction, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DirectResponseAction;
  static deserializeBinaryFromReader(message: DirectResponseAction, reader: jspb.BinaryReader): DirectResponseAction;
}

export namespace DirectResponseAction {
  export type AsObject = {
    status: number,
    body: string,
  }
}

export class CorsPolicy extends jspb.Message {
  clearAllowOriginList(): void;
  getAllowOriginList(): Array<string>;
  setAllowOriginList(value: Array<string>): void;
  addAllowOrigin(value: string, index?: number): string;

  clearAllowOriginRegexList(): void;
  getAllowOriginRegexList(): Array<string>;
  setAllowOriginRegexList(value: Array<string>): void;
  addAllowOriginRegex(value: string, index?: number): string;

  clearAllowMethodsList(): void;
  getAllowMethodsList(): Array<string>;
  setAllowMethodsList(value: Array<string>): void;
  addAllowMethods(value: string, index?: number): string;

  clearAllowHeadersList(): void;
  getAllowHeadersList(): Array<string>;
  setAllowHeadersList(value: Array<string>): void;
  addAllowHeaders(value: string, index?: number): string;

  clearExposeHeadersList(): void;
  getExposeHeadersList(): Array<string>;
  setExposeHeadersList(value: Array<string>): void;
  addExposeHeaders(value: string, index?: number): string;

  getMaxAge(): string;
  setMaxAge(value: string): void;

  getAllowCredentials(): boolean;
  setAllowCredentials(value: boolean): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CorsPolicy.AsObject;
  static toObject(includeInstance: boolean, msg: CorsPolicy): CorsPolicy.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CorsPolicy, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CorsPolicy;
  static deserializeBinaryFromReader(message: CorsPolicy, reader: jspb.BinaryReader): CorsPolicy;
}

export namespace CorsPolicy {
  export type AsObject = {
    allowOriginList: Array<string>,
    allowOriginRegexList: Array<string>,
    allowMethodsList: Array<string>,
    allowHeadersList: Array<string>,
    exposeHeadersList: Array<string>,
    maxAge: string,
    allowCredentials: boolean,
  }
}

