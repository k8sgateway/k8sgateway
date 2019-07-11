import styled from '@emotion/styled/macro';
import { useCreateUpstream } from 'Api';
import {
  SoloFormDropdown,
  SoloFormInput,
  SoloFormTypeahead
} from 'Components/Common/Form/SoloFormField';
import { SoloFormTemplate } from 'Components/Common/Form/SoloFormTemplate';
import { SoloButton } from 'Components/Common/SoloButton';
import { Field, Formik } from 'formik';
import { NamespacesContext } from 'GlooIApp';
import { UpstreamSpec as AwsUpstreamSpec } from 'proto/github.com/solo-io/gloo/projects/gloo/api/v1/plugins/aws/aws_pb';
import { UpstreamSpec as AzureUpstreamSpec } from 'proto/github.com/solo-io/gloo/projects/gloo/api/v1/plugins/azure/azure_pb';
import { UpstreamSpec as KubeUpstreamSpec } from 'proto/github.com/solo-io/gloo/projects/gloo/api/v1/plugins/kubernetes/kubernetes_pb';
import { UpstreamSpec as StaticUpstreamSpec } from 'proto/github.com/solo-io/gloo/projects/gloo/api/v1/plugins/static/static_pb';
import { ResourceRef } from 'proto/github.com/solo-io/solo-kit/api/v1/ref_pb';
import {
  CreateUpstreamRequest,
  UpstreamInput
} from 'proto/github.com/solo-io/solo-projects/projects/grpcserver/api/v1/upstream_pb';
import * as React from 'react';
import { UPSTREAM_SPEC_TYPES, UPSTREAM_TYPES } from 'utils/upstreamHelpers';
import * as yup from 'yup';
import { awsInitialValues, AwsUpstreamForm } from './AwsUpstreamForm';
import { kubeInitialValues, KubeUpstreamForm } from './KubeUpstreamForm';
import { AzureUpstreamForm, azureInitialValues } from './AzureUpstreamForm';

interface Props {}

const FormContainer = styled.div`
  display: flex;
  flex-direction: column;
`;

export const InputContainer = styled.div`
  display: flex;
  flex-direction: row;
  justify-content: space-around;
  padding: 10px;
`;

const Footer = styled.div`
  display: flex;
  flex-direction: row;
  justify-content: flex-end;
`;

// TODO: better way to include all initial values?
let initialValues = {
  name: '',
  type: '',
  namespace: 'gloo-system',
  ...awsInitialValues,
  ...kubeInitialValues,
  ...azureInitialValues
};

// TODO combine validation schemas
const validationSchema = yup.object().shape({
  name: yup.string(),
  namespace: yup.string(),
  type: yup.string()
});

export const CreateUpstreamForm = (props: Props) => {
  const namespaces = React.useContext(NamespacesContext);

  const { refetch: makeRequest } = useCreateUpstream(null);

  // grpc request
  function createUpstream(values: typeof initialValues) {
    let newUpstreamReq = new CreateUpstreamRequest();
    let usInput = new UpstreamInput();

    let usResourceRef = new ResourceRef();

    usResourceRef.setName(values.name);
    usResourceRef.setNamespace(values.namespace);

    usInput.setRef(usResourceRef);

    //TODO: set up correct upstream spec
    // TODO: validation for specific fields
    switch (values.type) {
      case UPSTREAM_SPEC_TYPES.AWS:
        let awsSpec = new AwsUpstreamSpec();
        awsSpec.setRegion(values.region);
        let awsSecretRef = new ResourceRef();
        awsSecretRef.setName(values.awsSecretRefName);
        awsSecretRef.setNamespace(values.awsSecretRefNamespace);

        awsSpec.setSecretRef(awsSecretRef);
        usInput.setAws(awsSpec);
        break;
      case UPSTREAM_SPEC_TYPES.AZURE:
        let azureSpec = new AzureUpstreamSpec();
        let azureSecretRef = new ResourceRef();
        azureSecretRef.setName(values.azureSecretRefName);
        azureSecretRef.setNamespace(values.azureSecretRefNamespace);
        azureSpec.setSecretRef(azureSecretRef);
        usInput.setAzure(azureSpec);
        break;
      case UPSTREAM_SPEC_TYPES.KUBE:
        let kubeSpec = new KubeUpstreamSpec();
        kubeSpec.setServiceName(values.serviceName);
        kubeSpec.setServiceNamespace(values.serviceNamespace);
        kubeSpec.setServicePort(+values.servicePort);
        usInput.setKube(kubeSpec);
        break;
      case UPSTREAM_SPEC_TYPES.STATIC:
        let staticSpec = new StaticUpstreamSpec();
        usInput.setStatic(staticSpec);
        break;
      default:
        break;
    }

    newUpstreamReq.setInput(usInput);
    makeRequest(newUpstreamReq);
  }

  return (
    <Formik
      initialValues={initialValues}
      validationSchema={validationSchema}
      onSubmit={createUpstream}>
      {({ values, isSubmitting, handleSubmit }) => (
        <FormContainer>
          <SoloFormTemplate>
            <Field
              name='name'
              title='Upstream Name'
              placeholder='Upstream Name'
              component={SoloFormInput}
            />
            <Field
              name='type'
              title='Upstream Type'
              placeholder='Type'
              options={UPSTREAM_TYPES}
              component={SoloFormDropdown}
            />
            <Field
              name='namespace'
              title='Upstream Namespace'
              defaultValue='gloo-system'
              presetOptions={namespaces}
              component={SoloFormTypeahead}
            />
          </SoloFormTemplate>
          {values.type === UPSTREAM_SPEC_TYPES.AWS && <AwsUpstreamForm />}
          {values.type === UPSTREAM_SPEC_TYPES.KUBE && <KubeUpstreamForm />}
          {values.type === UPSTREAM_SPEC_TYPES.AZURE && <AzureUpstreamForm />}
          <Footer>
            <SoloButton
              onClick={handleSubmit}
              text='Create Upstream'
              disabled={isSubmitting}
            />
          </Footer>
        </FormContainer>
      )}
    </Formik>
  );
};
