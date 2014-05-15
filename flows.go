package gotelemetry

import (
	"reflect"
)

type Flow struct {
	Tag  string
	Data interface{}
}

func NewFlow(tag string, data interface{}) *Flow {
	if reflect.TypeOf(data).Kind() != reflect.Ptr {
		panic("NewFlow() expects a pointer to a variant struct")
	}

	return &Flow{tag, data}
}

func (f *Flow) Publish(credentials Credentials) error {
	r, err := buildRequest(
		"PUT",
		credentials,
		"/flows/"+f.Tag+"/data",
		f.Data,
	)

	if err != nil {
		return err
	}

	return sendJSONRequest(r, nil)
}

func (f *Flow) BarchartData() (*Barchart, bool) {
	res, ok := f.Data.(*Barchart)

	return res, ok
}

func (f *Flow) BulletchartData() (*Bulletchart, bool) {
	res, ok := f.Data.(*Bulletchart)

	return res, ok
}

func (f *Flow) CountdownData() (*Countdown, bool) {
	res, ok := f.Data.(*Countdown)

	return res, ok
}

func (f *Flow) CustomData() (*Custom, bool) {
	res, ok := f.Data.(*Custom)

	return res, ok
}

func (f *Flow) FunnelchartData() (*Funnelchart, bool) {
	res, ok := f.Data.(*Funnelchart)

	return res, ok
}

func (f *Flow) GaugeData() (*Gauge, bool) {
	res, ok := f.Data.(*Gauge)

	return res, ok
}

func (f *Flow) GraphData() (*Graph, bool) {
	res, ok := f.Data.(*Graph)

	return res, ok
}

func (f *Flow) GridData() (*Grid, bool) {
	res, ok := f.Data.(*Grid)

	return res, ok
}

func (f *Flow) HistogramData() (*Histogram, bool) {
	res, ok := f.Data.(*Histogram)

	return res, ok
}

func (f *Flow) IconData() (*Icon, bool) {
	res, ok := f.Data.(*Icon)

	return res, ok
}

func (f *Flow) ImageData() (*Image, bool) {
	res, ok := f.Data.(*Image)

	return res, ok
}

func (f *Flow) LogData() (*Log, bool) {
	res, ok := f.Data.(*Log)

	return res, ok
}

func (f *Flow) MapData() (*Map, bool) {
	res, ok := f.Data.(*Map)

	return res, ok
}

func (f *Flow) MultigaugeData() (*Multigauge, bool) {
	res, ok := f.Data.(*Multigauge)

	return res, ok
}

func (f *Flow) MultivalueData() (*Multivalue, bool) {
	res, ok := f.Data.(*Multivalue)

	return res, ok
}

func (f *Flow) PiechartData() (*Piechart, bool) {
	res, ok := f.Data.(*Piechart)

	return res, ok
}

func (f *Flow) ScatterplotData() (*Scatterplot, bool) {
	res, ok := f.Data.(*Scatterplot)

	return res, ok
}

func (f *Flow) ServersData() (*Servers, bool) {
	res, ok := f.Data.(*Servers)

	return res, ok
}

func (f *Flow) StatusData() (*Status, bool) {
	res, ok := f.Data.(*Status)

	return res, ok
}

func (f *Flow) TableData() (*Table, bool) {
	res, ok := f.Data.(*Table)

	return res, ok
}

func (f *Flow) TextData() (*Text, bool) {
	res, ok := f.Data.(*Text)

	return res, ok
}

func (f *Flow) TickertapeData() (*Tickertape, bool) {
	res, ok := f.Data.(*Tickertape)

	return res, ok
}

func (f *Flow) TimelineData() (*Timeline, bool) {
	res, ok := f.Data.(*Timeline)

	return res, ok
}

func (f *Flow) TimeseriesData() (*Timeseries, bool) {
	res, ok := f.Data.(*Timeseries)

	return res, ok
}

func (f *Flow) UpstatusData() (*Upstatus, bool) {
	res, ok := f.Data.(*Upstatus)

	return res, ok
}

func (f *Flow) ValueData() (*Value, bool) {
	res, ok := f.Data.(*Value)

	return res, ok
}

func (f *Flow) WaterfallData() (*Waterfall, bool) {
	res, ok := f.Data.(*Waterfall)

	return res, ok
}
