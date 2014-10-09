// Package gotelemetry provides bindings for the Telemetry visualization service (http://telemetryapp.com).
//
// In order to use the package, you will need to sign up for an account and obtain your API token from
// https://www.telemetryapp.com/user/documentation/authentication. A full set of documents that explain
// how the Telemetry REST API works is also available at https://www.telemetryapp.com/user/documentation/.
//
// The package supports submitting flow data either individually or in batches to minimize network use.
// Please bear in mind that both the Flow and Batch submission mechanisms are not, by default,
// thread-safe; if thread-safety is desired, it must be provided by the caller.
package gotelemetry

import (
	"net/http"
	"reflect"
)

// Struct Flow identifies a flow, defined as the combination of a tag and
// the data associated with it, which must be a pointer to one of the structs declared
// in variants.go
//
// Note that Flow structs are not thread-safe by default, because they store a pointer to
// the underlying variant information. If you require thread-safety, you must mediate access
// to the flow through a synchronization mechanism of some kind, like a mutex.
//
// Flows are designed to be instantiated once and then modified as needed; you can grab
// a pointer to the appropriate underlying data by calling one of the *Data() methods
// of the struct.
type Flow struct {
	Tag             string `json:"tag"`
	Data            interface{}
	Variant         string `json:"variant"`
	Source_provider string `json:"source_provider,omitempty"`
	Filter          string `json:"filter,omitempty"`
	Params          string `json:"params,omitempty"`
}

// NewFlow() creates a new flow. Note that the `data` parameter *must* be a pointer to
// one of the variant structs defined in variant.go. If anything other than a pointer
// is passed, the function panics to prevent the creation of a silently immutable flow.
//
// If the flow is being submitted individually, the tag can be one of:
//
// ** The flow's named tag as entered in the Telemetry admin interface (e.g.: `gauge_1`)
//
// ** The flow's unique ID
//
// ** The flow's embed ID
//
// If, on the other hand, the flow is being submitted as part of a batch, only named
// tags are supported.
func NewFlow(tag string, data interface{}) *Flow {
	if reflect.TypeOf(data).Kind() != reflect.Ptr {
		panic("NewFlow() expects a pointer to a variant struct")
	}

	return &Flow{Tag: tag, Data: data}
}

// Publish() sends a flow to the Telemetry API servers. On output, the function return
// nil if the submission was successful, an instance of gotelemetry.Error if a REST
// error occurred, or a errors.Error instance otherwise.
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

	_, err = sendJSONRequest(r)

	return err
}

func (f *Flow) Read(credentials Credentials) error {
	req, err := buildRequest(
		"GET",
		credentials,
		"/flows/"+f.Tag+"/data",
		nil,
	)

	if err != nil {
		return err
	}

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return err
	}

	err = readJSONResponseBody(res, f.Data)

	return err
}

func (f *Flow) CreateFlow(credentials Credentials) error {
	request, err := buildRequest("POST", credentials, "/flows", f)

	if err != nil {
		return err
	}

	_, err = sendJSONRequest(request)

	return err
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
