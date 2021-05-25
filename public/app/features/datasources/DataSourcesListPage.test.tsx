import React from 'react';
import { shallow } from 'enzyme';
import { DataSourceSettings, NavModel } from '@grafana/data';

import { DataSourcesListPage, Props } from './DataSourcesListPage';
import { getMockDataSources } from './__mocks__/dataSourcesMocks';
import { setDataSourcesSearchQuery } from './state/reducers';

const setup = (propOverrides?: object) => {
  const props: Props = {
    dataSources: [] as DataSourceSettings[],
    loadDataSources: jest.fn(),
    navModel: {
      main: {
        text: 'Configuration',
      },
      node: {
        text: 'Data Sources',
      },
    } as NavModel,
    dataSourcesCount: 0,
    searchQuery: '',
    setDataSourcesSearchQuery,
    hasFetched: false,
  };

  Object.assign(props, propOverrides);

  return shallow(<DataSourcesListPage {...props} />);
};

describe('Render', () => {
  it('should render component', () => {
    const wrapper = setup();

    expect(wrapper).toMatchSnapshot();
  });

  it('should render action bar and datasources', () => {
    const wrapper = setup({
      dataSources: getMockDataSources(5),
      dataSourcesCount: 5,
      hasFetched: true,
    });

    expect(wrapper).toMatchSnapshot();
  });
});
