import * as React from 'react';
/** @jsx jsx */
import { jsx } from '@emotion/core';
import { isEmpty } from 'lodash';
import styled from '@emotion/styled/macro';
import { colors, soloConstants } from 'Styles';
import {
  SoloFormTemplate,
  InputRow
} from 'Components/Common/Form/SoloFormTemplate';
import {
  SoloFormInput,
  SoloFormDropdown,
  SoloFormMultiselect,
  SoloFormMultipartStringCardsList,
  SoloFormMetadataBasedDropdown
} from 'Components/Common/Form/SoloFormField';
import { Field, Formik } from 'formik';
import * as yup from 'yup';

import { VirtualService } from 'proto/github.com/solo-io/gloo/projects/gateway/api/v1/virtual_service_pb';
import { Upstream } from 'proto/github.com/solo-io/gloo/projects/gloo/api/v1/upstream_pb';
import {
  CreateRouteRequest,
  RouteInput,
  ListVirtualServicesRequest
} from 'proto/github.com/solo-io/solo-projects/projects/grpcserver/api/v1/virtualservice_pb';
import {
  useCreateRoute,
  useGetUpstreamsList,
  useListVirtualServices
} from 'Api';
import { ResourceRef } from 'proto/github.com/solo-io/solo-kit/api/v1/ref_pb';
import {
  Route,
  Matcher,
  HeaderMatcher,
  QueryParameterMatcher,
  RouteAction,
  Destination,
  KubernetesServiceDestination,
  ConsulServiceDestination,
  RedirectAction
} from 'proto/github.com/solo-io/gloo/projects/gloo/api/v1/proxy_pb';
import { DestinationSpec } from 'proto/github.com/solo-io/gloo/projects/gloo/api/v1/plugins_pb';
import { DestinationSpec as AWSDestinationSpec } from 'proto/github.com/solo-io/gloo/projects/gloo/api/v1/plugins/aws/aws_pb';
import { DestinationSpec as AzureDestinationSpec } from 'proto/github.com/solo-io/gloo/projects/gloo/api/v1/plugins/azure/azure_pb';
import { DestinationSpec as RestDestinationSpec } from 'proto/github.com/solo-io/gloo/projects/gloo/api/v1/plugins/rest/rest_pb';
import { DestinationSpec as GrpcDestinationSpec } from 'proto/github.com/solo-io/gloo/projects/gloo/api/v1/plugins/grpc/grpc_pb';
import { ListUpstreamsRequest } from 'proto/github.com/solo-io/solo-projects/projects/grpcserver/api/v1/upstream_pb';
import { Loading } from 'Components/Common/Loading';
import { ErrorText } from '../VirtualService/Details/ExtAuthForm';
import { NamespacesContext } from 'GlooIApp';
import {
  createUpstreamId,
  parseUpstreamId,
  createVirtualServiceId,
  parseVirtualServiceId
} from 'utils/helpers';
import { SoloButton } from 'Components/Common/SoloButton';
import { ButtonProgress } from 'Styles/CommonEmotions/button';

enum PathSpecifierCase { // From gloo -> proxy_pb -> Matcher's namespace
  PATH_SPECIFIER_NOT_SET = 0,
  PREFIX = 1,
  EXACT = 2,
  REGEX = 3
}

export const PATH_SPECIFIERS = [
  {
    key: 'PREFIX',
    value: 'PREFIX',
    displayValue: 'Prefix'
  },
  {
    key: 'EXACT',
    value: 'EXACT',
    displayValue: 'Exact'
  },
  {
    key: 'REGEX',
    value: 'REGEX',
    displayValue: 'Regex'
  }
];

export interface CreateRouteValuesType {
  virtualService: VirtualService.AsObject | undefined;
  upstream: Upstream.AsObject | undefined;
  path: string;
  matchType: 'PREFIX' | 'EXACT' | 'REGEX';
  headers: {
    name: string;
    value: string;
    isRegex: boolean;
  }[];
  queryParameters: {
    name: string;
    value: string;
    isRegex: boolean;
  }[];
  methods: {
    POST: boolean;
    PUT: boolean;
    GET: boolean;
    PATCH: boolean;
    DELETE: boolean;
    HEAD: boolean;
    OPTIONS: boolean;
  };
}

export const createRouteDefaultValues: CreateRouteValuesType = {
  virtualService: new VirtualService().toObject(),
  upstream: new Upstream().toObject(),
  path: '',
  matchType: 'PREFIX',
  headers: [],
  queryParameters: [],
  methods: {
    POST: false,
    PUT: false,
    GET: false,
    PATCH: false,
    DELETE: false,
    HEAD: false,
    OPTIONS: false
  }
};

const validationSchema = yup.object().shape({
  region: yup.string(),
  virtualService: yup.object(),
  upstream: yup.object(),
  path: yup
    .string()
    .test('Valid Path', 'Paths begin with /', val => val && val[0] === '/'),
  matchType: yup.string(),
  headers: yup.array().of(
    yup.object().shape({
      name: yup.string(),
      value: yup.string(),
      isRegex: yup.boolean()
    })
  ),
  queryParameters: yup.array().of(
    yup.object().shape({
      name: yup.string(),
      value: yup.string(),
      isRegex: yup.boolean()
    })
  ),
  methods: yup.object().shape({
    POST: yup.boolean(),
    PUT: yup.boolean(),
    GET: yup.boolean(),
    PATCH: yup.boolean(),
    DELETE: yup.boolean(),
    HEAD: yup.boolean(),
    OPTIONS: yup.boolean()
  })
});

const FormContainer = styled.div`
  display: flex;
  flex-direction: column;
`;

const HalfColumn = styled.div`
  width: calc(50% - 10px);
`;
const Footer = styled.div`
  display: flex;
  flex-direction: row;
  justify-content: flex-end;
  margin-top: 28px;
`;

interface Props {
  defaultVirtualService?: VirtualService.AsObject;
  defaultUpstream?: Upstream.AsObject;
  completeCreation: (newVirtualService?: VirtualService.AsObject) => any;
}

export const CreateRouteModal = (props: Props) => {
  const [
    allUsableVirtualServices,
    setAllUsableVirtualServices
  ] = React.useState<VirtualService.AsObject[]>([]);
  const [allUsableUpstreams, setAllUsableUpstreams] = React.useState<
    Upstream.AsObject[]
  >([]);

  const {
    data: createdVirtualServiceData,
    refetch: makeRequest
  } = useCreateRoute(null);
  let listVirtualServicesRequest = React.useRef(
    new ListVirtualServicesRequest()
  );
  let listUpstreamsRequest = React.useRef(new ListUpstreamsRequest());
  const namespaces = React.useContext(NamespacesContext);
  listVirtualServicesRequest.current.setNamespacesList(namespaces);
  listUpstreamsRequest.current.setNamespacesList(namespaces);

  React.useEffect(() => {
    if (!!createdVirtualServiceData) {
      props.completeCreation(createdVirtualServiceData.virtualService);
    }
  }, [createdVirtualServiceData]);

  const {
    data: upstreamsData,
    error: upstreamsError,
    loading: upstreamsLoading
  } = useGetUpstreamsList(listUpstreamsRequest.current);
  const {
    data: virtualServicesData,
    error: virtualServicesError,
    loading: virtualServicesLoading
  } = useListVirtualServices(listVirtualServicesRequest.current);

  React.useEffect(() => {
    setAllUsableVirtualServices(
      !!virtualServicesData
        ? virtualServicesData.virtualServicesList.filter(vs => !!vs.metadata)
        : []
    );
  }, [virtualServicesData]);
  React.useEffect(() => {
    setAllUsableUpstreams(
      !!upstreamsData
        ? upstreamsData.upstreamsList.filter(upstream => !!upstream.metadata)
        : []
    );
  }, [upstreamsData]);

  if (upstreamsLoading || virtualServicesLoading) {
    return <Loading />;
  }
  if (!!upstreamsError || !!virtualServicesError) {
    // @ts-ignore
    return <ErrorText>{upstreamsError || virtualServicesError}</ErrorText>;
  }

  const { defaultUpstream, defaultVirtualService } = props;

  const createRoute = (values: CreateRouteValuesType) => {
    let newRouteReq = new CreateRouteRequest();
    let reqRouteInput = new RouteInput();

    let virtualServiceResourceRef = new ResourceRef();
    virtualServiceResourceRef.setName(values.virtualService!.metadata!.name);
    virtualServiceResourceRef.setNamespace(
      values.virtualService!.metadata!.namespace
    );
    reqRouteInput.setVirtualServiceRef(virtualServiceResourceRef);

    //reqRouteInput.setIndex(vs.virtualHost!.routesList.length);

    /***
     *  ROUTE CREATION BEGINS
     * */
    let newRoute = new Route();
    let routeMatcher = new Matcher();
    switch (values.matchType) {
      case 'PREFIX':
        routeMatcher.setPrefix(values.path);
        break;
      case 'EXACT':
        routeMatcher.setExact(values.path);
        break;
      case 'REGEX':
        routeMatcher.setRegex(values.path);
        break;
    }

    let matcherHeaders: HeaderMatcher[] = values.headers.map(head => {
      const newMatcherHeader = new HeaderMatcher();
      newMatcherHeader.setName(head.name);
      newMatcherHeader.setValue(head.value);
      newMatcherHeader.setRegex(head.isRegex);

      return newMatcherHeader;
    });
    routeMatcher.setHeadersList(matcherHeaders);
    let matcherQueryParams: QueryParameterMatcher[] = values.queryParameters.map(
      queryParam => {
        const newMatcherQueryParam = new QueryParameterMatcher();
        newMatcherQueryParam.setName(queryParam.name);
        newMatcherQueryParam.setValue(queryParam.value);
        newMatcherQueryParam.setRegex(queryParam.isRegex);

        return newMatcherQueryParam;
      }
    );
    routeMatcher.setQueryParametersList(matcherQueryParams);
    routeMatcher.setMethodsList(
      //@ts-ignore
      Object.keys(values.methods).filter(key => values.methods[key])
    );
    newRoute.setMatcher(routeMatcher);

    /* Route->Destination Section */
    let newRouteAction = new RouteAction();
    let newDestination = new Destination();
    const upstreamSpec = values.upstream!.upstreamSpec!;
    let newDestinationResourceRef = new ResourceRef();
    newDestinationResourceRef.setName(values.upstream!.metadata!.name);
    newDestinationResourceRef.setNamespace(
      values.upstream!.metadata!.namespace
    );
    let newDestinationSpec = new DestinationSpec();

    if (!!upstreamSpec.aws) {
      newDestination.setUpstream(newDestinationResourceRef);
      let newAWSDestinationSpec = new AWSDestinationSpec();
      // TODO :: I have no idea what to set the values to
      //newAWSDestinationSpec.setInvocationStyle(0);
      newDestinationSpec.setAws(newAWSDestinationSpec);
    } else if (!!upstreamSpec.azure) {
      newDestination.setUpstream(newDestinationResourceRef);
      let newAzureDestinationSpec = new AzureDestinationSpec();
      // TODO :: I have no idea what to set the values to
      newDestinationSpec.setAzure(newAzureDestinationSpec);
    } /*else if (!!upstreamSpec.kube) {
      let newKubeServiceDestination = new KubernetesServiceDestination();
      newKubeServiceDestination.setRef(newDestinationResourceRef);
      // TODO :: I have no idea what to set the values to
      newDestination.setKube(newKubeServiceDestination);
      let newKubeDestinationSpec;
      // TODO:: How do we tell if it is rest or GRPC?
      //if() -> set DestinationSpec to grpc...
      newDestination.setDestinationSpec(newKubeDestinationSpec);
    } else if (!!upstreamSpec.consul) {
      let newConsulServiceDestination = new ConsulServiceDestination();
      // TODO :: I have no idea what to set the values to
      newDestination.setConsul(newConsulServiceDestination);
      let newConsulDestinationSpec;
      // TODO:: I have no idea what goes in this case
      newDestination.setDestinationSpec(newConsulDestinationSpec);
    }*/
    newDestination.setDestinationSpec(newDestinationSpec);
    newRouteAction.setSingle(newDestination);
    newRoute.setRouteAction(newRouteAction);

    // It looks like we don't see any of the other actions if
    // Route Action is taken??  But if they supplied
    // a path, shouldn't we do the redirect action?
    // Not clear on what the other actions would be based on?

    /*
    let newRedirectAction = new RedirectAction();
    //TODO:: Do we need to set anything else for this???
    if(values.matchType === PathSpecifierCase.PREFIX) {
      // TODO:: Is this correct??
      newRedirectAction.setPrefixRewrite("PREFIX");
    } else {
      newRedirectAction.setPathRedirect(values.path);
    }
    newRoute.setRedirectAction(newRedirectAction);*/

    reqRouteInput.setRoute(newRoute);
    /***
     *  ROUTE CREATION ENDS
     * */

    newRouteReq.setInput(reqRouteInput);
    makeRequest(newRouteReq);
  };

  const initialValues: CreateRouteValuesType = {
    ...createRouteDefaultValues,
    virtualService: defaultVirtualService
      ? defaultVirtualService
      : createRouteDefaultValues.virtualService,
    upstream: defaultUpstream
      ? defaultUpstream
      : createRouteDefaultValues.upstream
  };

  return (
    <Formik
      initialValues={initialValues}
      validationSchema={validationSchema}
      onSubmit={createRoute}>
      {({ values, isSubmitting, handleSubmit, isValid, errors }) => {
        return (
          <FormContainer>
            <SoloFormTemplate>
              <InputRow>
                {allUsableVirtualServices.length && (
                  <HalfColumn>
                    <SoloFormMetadataBasedDropdown
                      name='virtualService'
                      title='Virtual Service'
                      value={values.virtualService}
                      placeholder='Virtual Service...'
                      options={allUsableVirtualServices}
                    />
                  </HalfColumn>
                )}
                {allUsableUpstreams.length && (
                  <HalfColumn>
                    <SoloFormMetadataBasedDropdown
                      name='upstream'
                      title='Upstream'
                      value={values.upstream}
                      placeholder='Upstream...'
                      options={allUsableUpstreams}
                    />
                  </HalfColumn>
                )}
              </InputRow>
              <InputRow>
                <HalfColumn>
                  <SoloFormInput
                    name='path'
                    title='Path'
                    placeholder='Path...'
                  />
                </HalfColumn>
                <HalfColumn>
                  <SoloFormDropdown
                    name='matchType'
                    title='Match Type'
                    defaultValue={'PREFIX'}
                    options={PATH_SPECIFIERS}
                  />
                </HalfColumn>
              </InputRow>
              <InputRow>
                <SoloFormMultipartStringCardsList
                  name='headers'
                  title='Headers'
                  values={values.headers}
                  createNewNamePromptText={'Name...'}
                  createNewValuePromptText={'Value...'}
                />
              </InputRow>
              <InputRow>
                <SoloFormMultipartStringCardsList
                  name='queryParameters'
                  title='Query Parameters'
                  values={values.queryParameters}
                  createNewNamePromptText={'Name...'}
                  createNewValuePromptText={'Value...'}
                />
              </InputRow>
              <InputRow>
                <SoloFormMultiselect
                  name='methods'
                  title='Methods'
                  placeholder='Methods...'
                  options={Object.keys(createRouteDefaultValues.methods).map(
                    key => {
                      return {
                        key: key,
                        value: key
                      };
                    }
                  )}
                />
              </InputRow>
            </SoloFormTemplate>
            <Footer>
              <SoloButton
                onClick={handleSubmit}
                text='Create Route'
                disabled={!isEmpty(errors)}
                loading={isSubmitting}
                inProgressText={'Creating Route...'}>
                <ButtonProgress />
              </SoloButton>
            </Footer>
          </FormContainer>
        );
      }}
    </Formik>
  );
};
