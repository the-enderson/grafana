import { SelectableValue } from '@grafana/data';
import { BarGaugeDisplayMode } from '@grafana/schema';
import { SingleStatBaseOptions } from '@grafana/ui';

export interface BarGaugeOptions extends SingleStatBaseOptions {
  displayMode: BarGaugeDisplayMode;
  showUnfilled: boolean;
  minVizWidth: number;
  minVizHeight: number;
}

export const displayModes: Array<SelectableValue<string>> = [
  { value: BarGaugeDisplayMode.Gradient, label: 'Gradient' },
  { value: BarGaugeDisplayMode.Lcd, label: 'Retro LCD' },
  { value: BarGaugeDisplayMode.Basic, label: 'Basic' },
];
