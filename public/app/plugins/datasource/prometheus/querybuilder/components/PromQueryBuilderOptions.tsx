import React, { SyntheticEvent } from 'react';

import { CoreApp, SelectableValue } from '@grafana/data';
import { EditorRow, EditorField, EditorSwitch } from '@grafana/experimental';
import { AutoSizeInput, RadioButtonGroup, Select } from '@grafana/ui';

import { getQueryTypeChangeHandler, getQueryTypeOptions } from '../../components/PromExploreExtraField';
import { FORMAT_OPTIONS, INTERVAL_FACTOR_OPTIONS } from '../../components/PromQueryEditor';
import { PromQuery } from '../../types';
import { QueryOptionGroup } from '../shared/QueryOptionGroup';

import { getLegendModeLabel, PromQueryLegendEditor } from './PromQueryLegendEditor';

export interface UIOptions {
  exemplars: boolean;
  type: boolean;
  format: boolean;
  minStep: boolean;
  legend: boolean;
  resolution: boolean;
}

export interface Props {
  query: PromQuery;
  app?: CoreApp;
  onChange: (update: PromQuery) => void;
  onRunQuery: () => void;
  uiOptions: UIOptions;
}

export const PromQueryBuilderOptions = React.memo<Props>(({ query, app, onChange, onRunQuery, uiOptions }) => {
  const onChangeFormat = (value: SelectableValue<string>) => {
    onChange({ ...query, format: value.value });
    onRunQuery();
  };

  const onChangeStep = (evt: React.FormEvent<HTMLInputElement>) => {
    onChange({ ...query, interval: evt.currentTarget.value });
    onRunQuery();
  };

  const queryTypeOptions = getQueryTypeOptions(app === CoreApp.Explore);
  const onQueryTypeChange = getQueryTypeChangeHandler(query, onChange);

  const onExemplarChange = (event: SyntheticEvent<HTMLInputElement>) => {
    const isEnabled = event.currentTarget.checked;
    onChange({ ...query, exemplar: isEnabled });
    onRunQuery();
  };

  const onIntervalFactorChange = (value: SelectableValue<number>) => {
    onChange({ ...query, intervalFactor: value.value });
    onRunQuery();
  };

  const formatOption = FORMAT_OPTIONS.find((option) => option.value === query.format) || FORMAT_OPTIONS[0];
  const queryTypeValue = getQueryTypeValue(query);
  const queryTypeLabel = queryTypeOptions.find((x) => x.value === queryTypeValue)!.label;

  return (
    <EditorRow>
      <QueryOptionGroup
        title="Options"
        collapsedInfo={getCollapsedInfo(query, formatOption.label!, queryTypeLabel, uiOptions)}
      >
        {uiOptions.legend && (
          <PromQueryLegendEditor
            legendFormat={query.legendFormat}
            onChange={(legendFormat) => onChange({ ...query, legendFormat })}
            onRunQuery={onRunQuery}
          />
        )}
        {uiOptions.minStep && (
          <EditorField
            label="Min step"
            tooltip={
              <>
                An additional lower limit for the step parameter of the Prometheus query and for the{' '}
                <code>$__interval</code> and <code>$__rate_interval</code> variables.
              </>
            }
          >
            <AutoSizeInput
              type="text"
              aria-label="Set lower limit for the step parameter"
              placeholder={'auto'}
              minWidth={10}
              onCommitChange={onChangeStep}
              defaultValue={query.interval}
            />
          </EditorField>
        )}
        {uiOptions.format && (
          <EditorField label="Format">
            <Select value={formatOption} allowCustomValue onChange={onChangeFormat} options={FORMAT_OPTIONS} />
          </EditorField>
        )}
        {uiOptions.type && (
          <EditorField label="Type">
            <RadioButtonGroup options={queryTypeOptions} value={queryTypeValue} onChange={onQueryTypeChange} />
          </EditorField>
        )}
        {uiOptions.exemplars && shouldShowExemplarSwitch(query, app) && (
          <EditorField label="Exemplars">
            <EditorSwitch value={query.exemplar} onChange={onExemplarChange} />
          </EditorField>
        )}
        {uiOptions.resolution && query.intervalFactor && query.intervalFactor > 1 && (
          <EditorField label="Resolution">
            <Select
              aria-label="Select resolution"
              isSearchable={false}
              options={INTERVAL_FACTOR_OPTIONS}
              onChange={onIntervalFactorChange}
              value={INTERVAL_FACTOR_OPTIONS.find((option) => option.value === query.intervalFactor)}
            />
          </EditorField>
        )}
      </QueryOptionGroup>
    </EditorRow>
  );
});

function shouldShowExemplarSwitch(query: PromQuery, app?: CoreApp) {
  if (app === CoreApp.UnifiedAlerting || !query.range) {
    return false;
  }

  return true;
}

function getQueryTypeValue(query: PromQuery) {
  return query.range && query.instant ? 'both' : query.instant ? 'instant' : 'range';
}

function getCollapsedInfo(query: PromQuery, formatOption: string, queryType: string, uiOptions: UIOptions): string[] {
  const items: string[] = [];

  if (uiOptions.legend) {
    items.push(`Legend: ${getLegendModeLabel(query.legendFormat)}`);
  }
  if (uiOptions.format) {
    items.push(`Format: ${formatOption}`);
  }
  if (uiOptions.minStep && query.interval) {
    items.push(`Step ${query.interval}`);
  }
  if (uiOptions.type) {
    items.push(`Type: ${queryType}`);
  }

  if (uiOptions.exemplars && query.exemplar) {
    items.push(`Exemplars: true`);
  }

  return items;
}

PromQueryBuilderOptions.displayName = 'PromQueryBuilderOptions';
