import * as React from 'react';
/** @jsx jsx */
import { jsx } from '@emotion/core';

import styled from '@emotion/styled/macro';

import { useFormValidation } from 'Hooks/useFormValidation';
import { SoloInput } from 'Components/Common/SoloInput';
import { ErrorText } from '../Details/ExtAuthForm';
import { SoloButton } from 'Components/Common/SoloButton';
import { useCreateVirtualService } from 'Api';
import { ResourceRef } from 'proto/github.com/solo-io/solo-kit/api/v1/ref_pb';
import {
  CreateVirtualServiceRequest,
  VirtualServiceInput
} from 'proto/github.com/solo-io/solo-projects/projects/grpcserver/api/v1/virtualservice_pb';
import { NamespacesContext } from 'GlooIApp';
import { SoloTypeahead } from 'Components/Common/SoloTypeahead';

const Footer = styled.div`
  display: flex;
  flex-direction: row;
  justify-content: flex-end;
`;

const InputContainer = styled.div`
  display: grid;
  grid-template-columns: 1fr 1fr;
  grid-gap: 10px;
`;

let initialValues = {
  virtualServiceName: '',
  displayName: '',
  addDomain: '',
  namespace: ''
};

const validate = (values: typeof initialValues) => {
  let errors = {} as typeof initialValues;
  if (!values.virtualServiceName) {
    errors.virtualServiceName = 'Name is required';
  }

  if (!values.namespace) {
    errors.namespace = 'Namespace is required';
  }

  return errors;
};

interface Props {
  onCompletion?: (succeeded?: { namespace: string; name: string }) => any;
}

export const CreateVirtualServiceForm = (props: Props) => {
  const namespaces = React.useContext(NamespacesContext);
  // this is to match the value displayed by the typeahead
  initialValues.namespace = namespaces[0] || '';
  const {
    handleSubmit,
    handleChange,
    handleBlur,
    values,
    errors,
    isSubmitting,
    isDifferent
  } = useFormValidation(initialValues, validate, createVirtualService);

  const { refetch: makeRequest } = useCreateVirtualService(null);

  function createVirtualService(values: typeof initialValues) {
    let vsRequest = new CreateVirtualServiceRequest();
    let vsInput = new VirtualServiceInput();

    let vsRef = new ResourceRef();
    vsRef.setName(values.virtualServiceName);
    vsRef.setNamespace(values.namespace);
    vsInput.setRef(vsRef);

    vsInput.setDisplayName(values.displayName);

    vsRequest.setInput(vsInput);
    makeRequest(vsRequest);

    if (!!props.onCompletion) {
      props.onCompletion({
        name: values.virtualServiceName,
        namespace: values.namespace
      });
    }
  }

  return (
    <div>
      <InputContainer>
        <div>
          <SoloInput
            title='Virtual Service Name'
            name='virtualServiceName'
            value={values.virtualServiceName}
            placeholder={'Virtual Service Name'}
            onChange={handleChange}
            onBlur={handleBlur}
          />
          {errors && <ErrorText>{errors.virtualServiceName}</ErrorText>}
        </div>
        <div>
          <SoloInput
            title='Display Name'
            name='displayName'
            value={values.displayName}
            placeholder={'Display Name'}
            onChange={handleChange}
            onBlur={handleBlur}
          />
          {errors && <ErrorText>{errors.displayName}</ErrorText>}
        </div>
        <div>
          <SoloTypeahead
            title='Virtual Service Namespace'
            defaultValue={values.namespace}
            onChange={e => handleChange(e, 'namespace')}
            presetOptions={namespaces}
          />
          {errors && <ErrorText>{errors.namespace}</ErrorText>}
        </div>
      </InputContainer>
      <Footer>
        <SoloButton
          onClick={handleSubmit}
          text='Create Virtual Service'
          disabled={isSubmitting}
        />
      </Footer>
    </div>
  );
};
