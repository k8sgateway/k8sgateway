import * as React from 'react';
/** @jsx jsx */
import { jsx } from '@emotion/core';

import styled from '@emotion/styled/macro';
import { RouteComponentProps } from 'react-router';
import { colors, healthConstants } from 'Styles';
import {
  TableActionCircle,
  TableHealthCircleHolder,
  TableActions
} from 'Styles/table';
import {
  ListingFilter,
  StringFilterProps,
  TypeFilterProps,
  CheckboxFilterProps,
  RadioFilterProps
} from 'Components/Common/ListingFilter';
import { SoloTable } from 'Components/Common/SoloTable';
import { SectionCard } from 'Components/Common/SectionCard';
import { CatalogTableToggle } from 'Components/Common/CatalogTableToggle';
import { ReactComponent as Gloo } from 'assets/Gloo.svg';
import { Breadcrumb } from 'Components/Common/Breadcrumb';
import { CardsListing } from 'Components/Common/CardsListing';
import { useListVirtualServices } from 'Api';
import { ListVirtualServicesRequest } from 'proto/github.com/solo-io/solo-projects/projects/grpcserver/api/v1/virtualservice_pb';
import { VirtualService } from 'proto/github.com/solo-io/gloo/projects/gateway/api/v1/virtual_service_pb';
import { Status } from 'proto/github.com/solo-io/solo-kit/api/v1/status_pb';
import { NamespacesContext } from 'GlooIApp';
import { getResourceStatus, getVSDomains } from 'utils/helpers';
import { CreateVirtualServiceModal } from './Creation/CreateVirtualServiceModal';
import { HealthInformation } from 'Components/Common/HealthInformation';
import { HealthIndicator } from 'Components/Common/HealthIndicator';
import { SoloModal } from 'Components/Common/SoloModal';
import { CreateRouteModal } from 'Components/Features/Route/CreateRouteModal';

const TableLink = styled.div`
  cursor: pointer;
  color: ${colors.seaBlue};
`;

const TableDomains = styled.div`
  max-width: 200px;
  max-height: 70px;
  overflow: hidden;
  text-overflow: ellipsis;
`;

const getTableColumns = (
  startCreatingRoute: (vs: VirtualService.AsObject) => any
) => {
  return [
    {
      title: 'Name',
      dataIndex: 'name',
      render: (nameObject: {
        goToVirtualService: () => void;
        displayName: string;
      }) => {
        return (
          <TableLink onClick={nameObject.goToVirtualService}>
            {nameObject.displayName}
          </TableLink>
        );
      }
    },
    {
      title: 'Domain',
      dataIndex: 'domains',
      render: (domains: string) => {
        return <TableDomains>{domains}</TableDomains>;
      }
    },
    {
      title: 'Namespace',
      dataIndex: 'metadata.namespace'
    },
    {
      title: 'Version',
      dataIndex: 'metadata.resourceVersion'
    },
    {
      title: 'Status',
      dataIndex: 'status',
      render: (healthStatus: Status.AsObject) => (
        <div>
          <TableHealthCircleHolder>
            <HealthIndicator healthStatus={healthStatus.state} />
          </TableHealthCircleHolder>
          <HealthInformation healthStatus={healthStatus} />
        </div>
      )
    },
    {
      title: 'Routes',
      dataIndex: 'routes'
    },
    {
      title: 'BR Limit',
      dataIndex: 'brLimit'
    },
    {
      title: 'AR Limit',
      dataIndex: 'arLimit'
    },
    {
      title: 'Actions',
      dataIndex: 'actions',
      render: (vs: VirtualService.AsObject) => {
        return (
          <TableActions>
            <TableActionCircle onClick={() => startCreatingRoute(vs)}>
              +
            </TableActionCircle>
          </TableActions>
        );
      }
    }
  ];
};

const StringFilters: StringFilterProps[] = [
  {
    displayName: 'Filter By Name...',
    placeholder: 'Filter by name...',
    value: ''
  }
];

const Heading = styled.div`
  display: flex;
  justify-content: space-between;
  margin-bottom: 20px;
`;

const Action = styled.div`
  display: flex;
  flex-direction: row;
  align-items: center;
  align-items: baseline;
`;

interface Props extends RouteComponentProps {}

export const VirtualServicesListing = (props: Props) => {
  let listVsRequest = React.useRef(new ListVirtualServicesRequest());
  const namespaces = React.useContext(NamespacesContext);

  listVsRequest.current.setNamespacesList(namespaces);
  const {
    data: vsListData,
    loading: vsLoading,
    error: vsError,
    refetch
  } = useListVirtualServices(listVsRequest.current);

  const [catalogNotTable, setCatalogNotTable] = React.useState(true);
  const [
    virtualServiceForRouteCreation,
    setVirtualServiceForRouteCreation
  ] = React.useState<VirtualService.AsObject | undefined>(undefined);
  const { history, match } = props;

  const getUsableCatalogData = (
    nameFilter: string,
    data: VirtualService.AsObject[]
  ) => {
    const dataUsed = data.map(virtualService => {
      return {
        ...virtualService,
        healthStatus: virtualService.status
          ? virtualService.status.state
          : healthConstants.Pending.value,
        cardTitle: virtualService.displayName || virtualService.metadata!.name,
        cardSubtitle: getVSDomains(virtualService),
        onRemovecard: (id: string): void => {},
        onExpanded: () => {},
        onClick: () => {
          history.push(
            `${match.path}${virtualService.metadata!.namespace}/${
              virtualService.metadata!.name
            }`
          );
        },
        onCreate: () => setVirtualServiceForRouteCreation(virtualService)
      };
    });

    return dataUsed.filter(row => row.cardTitle.includes(nameFilter));
  };

  const getUsableTableData = (
    nameFilter: string,
    data: VirtualService.AsObject[]
  ) => {
    const dataUsed = data.map(virtualService => {
      return {
        ...virtualService,
        name: {
          displayName: virtualService.metadata!.name,
          goToVirtualService: () => {
            history.push(
              `${match.path}${virtualService.metadata!.namespace}/${
                virtualService.metadata!.name
              }`
            );
          }
        },
        domains: getVSDomains(virtualService),
        routes: virtualService.virtualHost!.routesList.length,
        status: virtualService.status,
        key: `${virtualService.metadata!.name}`,
        actions: virtualService
      };
    });

    return dataUsed.filter(row => row.name.displayName.includes(nameFilter));
  };

  const listDisplay = (
    strings: StringFilterProps[],
    types: TypeFilterProps[],
    checkboxes: CheckboxFilterProps[],
    radios: RadioFilterProps[]
  ) => {
    const nameFilterValue: string = strings.find(
      s => s.displayName === 'Filter By Name...'
    )!.value!;

    if (!vsListData || vsLoading) {
      return <div>Loading...</div>;
    }
    return (
      <div>
        {catalogNotTable ? (
          <SectionCard cardName={'Virtual Services'} logoIcon={<Gloo />}>
            <CardsListing
              cardsData={getUsableCatalogData(
                nameFilterValue,
                vsListData.virtualServicesList
              )}
            />
          </SectionCard>
        ) : (
          <SoloTable
            dataSource={getUsableTableData(
              nameFilterValue,
              vsListData.virtualServicesList
            )}
            columns={getTableColumns(setVirtualServiceForRouteCreation)}
          />
        )}
      </div>
    );
  };

  const finishCreation = (succeeded?: {
    namespace: string;
    name: string;
  }): void => {
    //TODO : Proper way to do this is to be polling always and, once we see the VS that matches this exists, we then jump

    if (succeeded) {
      setTimeout(() => {
        history.push(`${match.path}${succeeded.namespace}/${succeeded.name}`);
      }, 500);
    }
  };

  return (
    <div>
      <Heading>
        <Breadcrumb />
        <Action>
          <CreateVirtualServiceModal finishCreation={finishCreation} />
          <CatalogTableToggle
            listIsSelected={!catalogNotTable}
            onToggle={() => {
              setCatalogNotTable(cNt => !cNt);
            }}
          />
        </Action>
      </Heading>
      <ListingFilter strings={StringFilters} filterFunction={listDisplay} />
      <SoloModal
        visible={!!virtualServiceForRouteCreation}
        width={500}
        title={'Create Route'}
        onClose={() => setVirtualServiceForRouteCreation(undefined)}>
        <CreateRouteModal
          defaultVirtualService={virtualServiceForRouteCreation}
          completeCreation={() => setVirtualServiceForRouteCreation(undefined)}
        />
      </SoloModal>
    </div>
  );
};
