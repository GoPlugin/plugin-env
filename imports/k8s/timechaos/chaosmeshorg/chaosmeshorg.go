// chaos-meshorg
package chaosmeshorg

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/goplugin/plugin-env/imports/k8s/timechaos/chaosmeshorg/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
	"github.com/goplugin/plugin-env/imports/k8s/timechaos/chaosmeshorg/internal"
)

// TimeChaos is the Schema for the timechaos API.
type TimeChaos interface {
	cdk8s.ApiObject
	// The group portion of the API version (e.g. `authorization.k8s.io`).
	ApiGroup() *string
	// The object's API version (e.g. `authorization.k8s.io/v1`).
	ApiVersion() *string
	// The chart in which this object is defined.
	Chart() cdk8s.Chart
	// The object kind.
	Kind() *string
	// Metadata associated with this API object.
	Metadata() cdk8s.ApiObjectMetadataDefinition
	// The name of the API object.
	//
	// If a name is specified in `metadata.name` this will be the name returned.
	// Otherwise, a name will be generated by calling
	// `Chart.of(this).generatedObjectName(this)`, which by default uses the
	// construct path to generate a DNS-compatible name for the resource.
	Name() *string
	// The tree node.
	Node() constructs.Node
	// Create a dependency between this ApiObject and other constructs.
	//
	// These can be other ApiObjects, Charts, or custom.
	AddDependency(dependencies ...constructs.IConstruct)
	// Applies a set of RFC-6902 JSON-Patch operations to the manifest synthesized for this API object.
	//
	// Example:
	//     kubePod.addJsonPatch(JsonPatch.replace('/spec/enableServiceLinks', true));
	//
	AddJsonPatch(ops ...cdk8s.JsonPatch)
	// Renders the object to Kubernetes JSON.
	ToJson() interface{}
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for TimeChaos
type jsiiProxy_TimeChaos struct {
	internal.Type__cdk8sApiObject
}

func (j *jsiiProxy_TimeChaos) ApiGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimeChaos) ApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimeChaos) Chart() cdk8s.Chart {
	var returns cdk8s.Chart
	_jsii_.Get(
		j,
		"chart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimeChaos) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimeChaos) Metadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimeChaos) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimeChaos) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Defines a "TimeChaos" API object.
func NewTimeChaos(scope constructs.Construct, id *string, props *TimeChaosProps) TimeChaos {
	_init_.Initialize()

	j := jsiiProxy_TimeChaos{}

	_jsii_.Create(
		"chaos-meshorg.TimeChaos",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Defines a "TimeChaos" API object.
func NewTimeChaos_Override(t TimeChaos, scope constructs.Construct, id *string, props *TimeChaosProps) {
	_init_.Initialize()

	_jsii_.Create(
		"chaos-meshorg.TimeChaos",
		[]interface{}{scope, id, props},
		t,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func TimeChaos_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"chaos-meshorg.TimeChaos",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Renders a Kubernetes manifest for "TimeChaos".
//
// This can be used to inline resource manifests inside other objects (e.g. as templates).
func TimeChaos_Manifest(props *TimeChaosProps) interface{} {
	_init_.Initialize()

	var returns interface{}

	_jsii_.StaticInvoke(
		"chaos-meshorg.TimeChaos",
		"manifest",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Returns the `ApiObject` named `Resource` which is a child of the given construct.
//
// If `c` is an `ApiObject`, it is returned directly. Throws an
// exception if the construct does not have a child named `Default` _or_ if
// this child is not an `ApiObject`.
func TimeChaos_Of(c constructs.IConstruct) cdk8s.ApiObject {
	_init_.Initialize()

	var returns cdk8s.ApiObject

	_jsii_.StaticInvoke(
		"chaos-meshorg.TimeChaos",
		"of",
		[]interface{}{c},
		&returns,
	)

	return returns
}

func TimeChaos_GVK() *cdk8s.GroupVersionKind {
	_init_.Initialize()
	var returns *cdk8s.GroupVersionKind
	_jsii_.StaticGet(
		"chaos-meshorg.TimeChaos",
		"GVK",
		&returns,
	)
	return returns
}

func (t *jsiiProxy_TimeChaos) AddDependency(dependencies ...constructs.IConstruct) {
	args := []interface{}{}
	for _, a := range dependencies {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		t,
		"addDependency",
		args,
	)
}

func (t *jsiiProxy_TimeChaos) AddJsonPatch(ops ...cdk8s.JsonPatch) {
	args := []interface{}{}
	for _, a := range ops {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		t,
		"addJsonPatch",
		args,
	)
}

func (t *jsiiProxy_TimeChaos) ToJson() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimeChaos) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// TimeChaos is the Schema for the timechaos API.
type TimeChaosProps struct {
	// Spec defines the behavior of a time chaos experiment.
	Spec *TimeChaosSpec `field:"required" json:"spec" yaml:"spec"`
	Metadata *cdk8s.ApiObjectMetadata `field:"optional" json:"metadata" yaml:"metadata"`
}

// Spec defines the behavior of a time chaos experiment.
type TimeChaosSpec struct {
	// Mode defines the mode to run chaos action.
	//
	// Supported mode: one / all / fixed / fixed-percent / random-max-percent.
	Mode TimeChaosSpecMode `field:"required" json:"mode" yaml:"mode"`
	// Selector is used to select pods that are used to inject chaos action.
	Selector *TimeChaosSpecSelector `field:"required" json:"selector" yaml:"selector"`
	// TimeOffset defines the delta time of injected program.
	//
	// It's a possibly signed sequence of decimal numbers, such as "300ms", "-1.5h" or "2h45m". Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".
	TimeOffset *string `field:"required" json:"timeOffset" yaml:"timeOffset"`
	// ClockIds defines all affected clock id All available options are ["CLOCK_REALTIME","CLOCK_MONOTONIC","CLOCK_PROCESS_CPUTIME_ID","CLOCK_THREAD_CPUTIME_ID", "CLOCK_MONOTONIC_RAW","CLOCK_REALTIME_COARSE","CLOCK_MONOTONIC_COARSE","CLOCK_BOOTTIME","CLOCK_REALTIME_ALARM", "CLOCK_BOOTTIME_ALARM"] Default value is ["CLOCK_REALTIME"].
	ClockIds *[]*string `field:"optional" json:"clockIds" yaml:"clockIds"`
	// ContainerNames indicates list of the name of affected container.
	//
	// If not set, all containers will be injected.
	ContainerNames *[]*string `field:"optional" json:"containerNames" yaml:"containerNames"`
	// Duration represents the duration of the chaos action.
	Duration *string `field:"optional" json:"duration" yaml:"duration"`
	// Value is required when the mode is set to `FixedPodMode` / `FixedPercentPodMod` / `RandomMaxPercentPodMod`.
	//
	// If `FixedPodMode`, provide an integer of pods to do chaos action. If `FixedPercentPodMod`, provide a number from 0-100 to specify the percent of pods the server can do chaos action. IF `RandomMaxPercentPodMod`,  provide a number from 0-100 to specify the max percent of pods to do chaos action
	Value *string `field:"optional" json:"value" yaml:"value"`
}

// Mode defines the mode to run chaos action.
//
// Supported mode: one / all / fixed / fixed-percent / random-max-percent.
type TimeChaosSpecMode string

const (
	// one.
	TimeChaosSpecMode_ONE TimeChaosSpecMode = "ONE"
	// all.
	TimeChaosSpecMode_ALL TimeChaosSpecMode = "ALL"
	// fixed.
	TimeChaosSpecMode_FIXED TimeChaosSpecMode = "FIXED"
	// fixed-percent.
	TimeChaosSpecMode_FIXED_PERCENT TimeChaosSpecMode = "FIXED_PERCENT"
	// random-max-percent.
	TimeChaosSpecMode_RANDOM_MAX_PERCENT TimeChaosSpecMode = "RANDOM_MAX_PERCENT"
)

// Selector is used to select pods that are used to inject chaos action.
type TimeChaosSpecSelector struct {
	// Map of string keys and values that can be used to select objects.
	//
	// A selector based on annotations.
	AnnotationSelectors *map[string]*string `field:"optional" json:"annotationSelectors" yaml:"annotationSelectors"`
	// a slice of label selector expressions that can be used to select objects.
	//
	// A list of selectors based on set-based label expressions.
	ExpressionSelectors *[]*TimeChaosSpecSelectorExpressionSelectors `field:"optional" json:"expressionSelectors" yaml:"expressionSelectors"`
	// Map of string keys and values that can be used to select objects.
	//
	// A selector based on fields.
	FieldSelectors *map[string]*string `field:"optional" json:"fieldSelectors" yaml:"fieldSelectors"`
	// Map of string keys and values that can be used to select objects.
	//
	// A selector based on labels.
	LabelSelectors *map[string]*string `field:"optional" json:"labelSelectors" yaml:"labelSelectors"`
	// Namespaces is a set of namespace to which objects belong.
	Namespaces *[]*string `field:"optional" json:"namespaces" yaml:"namespaces"`
	// Nodes is a set of node name and objects must belong to these nodes.
	Nodes *[]*string `field:"optional" json:"nodes" yaml:"nodes"`
	// Map of string keys and values that can be used to select nodes.
	//
	// Selector which must match a node's labels, and objects must belong to these selected nodes.
	NodeSelectors *map[string]*string `field:"optional" json:"nodeSelectors" yaml:"nodeSelectors"`
	// PodPhaseSelectors is a set of condition of a pod at the current time.
	//
	// supported value: Pending / Running / Succeeded / Failed / Unknown.
	PodPhaseSelectors *[]*string `field:"optional" json:"podPhaseSelectors" yaml:"podPhaseSelectors"`
	// Pods is a map of string keys and a set values that used to select pods.
	//
	// The key defines the namespace which pods belong, and the each values is a set of pod names.
	Pods *map[string]*[]*string `field:"optional" json:"pods" yaml:"pods"`
}

// A label selector requirement is a selector that contains values, a key, and an operator that relates the key and values.
type TimeChaosSpecSelectorExpressionSelectors struct {
	// key is the label key that the selector applies to.
	Key *string `field:"required" json:"key" yaml:"key"`
	// operator represents a key's relationship to a set of values.
	//
	// Valid operators are In, NotIn, Exists and DoesNotExist.
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// values is an array of string values.
	//
	// If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

