import * as React from 'react';
/** @jsx jsx */
import { jsx } from '@emotion/core';

import styled from '@emotion/styled/macro';
import { colors } from 'Styles';
import { Divider } from 'antd';
import { SoloModal } from 'Components/Common/SoloModal';
import { CreateUpstreamForm } from './CreateUpstreamForm';
import { ReactComponent as GreenPlus } from 'assets/small-green-plus.svg';

interface Props {}

const StyledGreenPlus = styled(GreenPlus)`
  cursor: pointer;
  margin-right: 7px;
  .a {
    fill: ${colors.forestGreen};
  }
`;

const ModalContainer = styled.div`
  display: flex;
  flex-direction: row;
  align-content: center;
`;
const Legend = styled.div`
  background-color: ${colors.januaryGrey};
  padding: 13px 13px 13px 10px;
  margin-bottom: 23px;
`;

// TODO: use spec font
const ModalTrigger = styled.div`
  cursor: pointer;
  display: flex;
  align-items: center;
  padding: 0 10px;
  font-size: 14px;
`;
export const CreateUpstreamModal = (props: Props) => {
  const [showModal, setShowModal] = React.useState(false);

  return (
    <ModalContainer>
      <ModalTrigger onClick={() => setShowModal(s => !s)}>
        <React.Fragment>
          <StyledGreenPlus />
          Create Upstream
        </React.Fragment>
        <Divider type='vertical' style={{ height: '1.5em' }} />
      </ModalTrigger>
      <SoloModal
        visible={showModal}
        width={650}
        title='Create an Upstream'
        onClose={() => setShowModal(false)}>
        <React.Fragment>
          <Legend>
            Upstreams define destinations for routes. Upstreams tell Gloo what
            to route to and how to route to them. Gloo determines how to handle
            routing for the upstream based on its spec field. Upstreams have a
            type-specific spec field which must be used to provide routing
            information to Gloo.
          </Legend>
          <CreateUpstreamForm />
        </React.Fragment>
      </SoloModal>
    </ModalContainer>
  );
};
