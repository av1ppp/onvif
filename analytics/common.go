package analytics

import (
	"github.com/av1ppp/onvif/xsd"
	"github.com/av1ppp/onvif/xsd/onvif"
)

type MetadataInfo struct {
	Type        xsd.QName `xml:"tan:Type"` // required
	SampleFrame Frame     `xml:"tan:SampleFrame"`
}

type Frame struct {
	UtcTime        xsd.DateTime     `xml:"tan:UtcTime"` // required
	Colorspace     xsd.String       `xml:"tan:Colorspace"`
	Source         xsd.String       `xml:"tan:Source"`
	PTZStatus      onvif.PTZStatus  `xml:"tan:PTZStatus"`      // optional
	Transformation Transformation   `xml:"tan:Transformation"` // optional
	Object         Object           `xml:"tan:Object"`         // optional, unbounded
	ObjectTree     ObjectTree       `xml:"tan:ObjectTree"`     // optional, unbounded
	Extension      FrameExtension   `xml:"tan:Extension"`      // optional
	SceneImageRef  xsd.AnyURI       `xml:"tan:SceneImageRef"`  // optional
	SceneImage     xsd.Base64Binary `xml:"tan:SceneImage"`     // optional
}

type FrameExtension struct {
	MotionInCells MotionInCells   `xml:"tan:MotionInCells"` // optional
	Extension     FrameExtension2 `xml:"tan:Extension"`     // optional
}

type FrameExtension2 struct {
	// TODO
}

type MotionInCells struct {
	Columns xsd.Integer      `xml:"tan:Columns"`
	Rows    xsd.Integer      `xml:"tan:Rows"`
	Cells   xsd.Base64Binary `xml:"tan:Cells"`
}

type ObjectTree struct {
	Rename    Rename              `xml:"tan:Rename"`    // optional, unbounded
	Split     Split               `xml:"tan:Split"`     // optional, unbounded
	Merge     Merge               `xml:"tan:Merge"`     // optional, unbounded
	Delete    ObjectId            `xml:"Delete"`        // TODO
	Extension ObjectTreeExtension `xml:"tan:Extension"` // optional
}

type ObjectTreeExtension struct {
	// TODO
}

type Merge struct {
	From ObjectId `xml:"from"` // TODO
	To   ObjectId `xml:"to"`   // TODO
}

type Split struct {
	From ObjectId `xml:"from"` // TODO
	To   ObjectId `xml:"to"`   // TODO
}

type Rename struct {
	From ObjectId `xml:"from"` // TODO
	To   ObjectId `xml:"to"`   // TODO
}

type ObjectId struct {
	ObjectId xsd.Integer `xml:"tan:ObjectId"`
	UUID     xsd.String  `xml:"tan:UUID"`
}

type Transformation struct {
	// TODO
}

type Object struct {
	ObjectId   xsd.Integer     `xml:"tan:ObjectId"`
	UUID       xsd.String      `xml:"tan:UUID"` // optional
	Parent     xsd.Integer     `xml:"tan:Parent"`
	Appearance Appearance      // optional
	Behaviour  Behaviour       // optional
	Extension  ObjectExtension // optional
}

type Behaviour struct {
	Removed   xsd.AnyType          `xml:"tan:Removed"`   // optional
	Idle      xsd.AnyType          `xml:"tan:Idle"`      // optional
	Extension BehaviourExtension   `xml:"tan:Extension"` // optional
	Speed     xsd.Float            `xml:"tan:Speed"`
	Direction onvif.GeoOrientation `xml:"tan:Direction"`
}

type BehaviourExtension struct {
	// TODO
}

type Appearance struct {
	Transformation      Transformation      `xml:"tan:Transformation"`      // optional
	Shape               ShapeDescriptor     `xml:"tan:Shape"`               // optional
	Color               ColorDescriptor     `xml:"tan:Color"`               // optional
	Class               ClassDescriptor     `xml:"tan:Class"`               // optional
	Extension           AppearanceExtension `xml:"tan:Extension"`           // optional
	GeoLocation         onvif.GeoLocation   `xml:"tan:GeoLocation"`         // optional
	VehicleInfo         VehicleInfo         `xml:"tan:VehicleInfo"`         // optional
	LicensePlateInfo    LicensePlateInfo    `xml:"tan:LicensePlateInfo"`    // optional
	HumanFace           HumanFace           `xml:"tan:HumanFace"`           // optional
	HumanBody           HumanBody           `xml:"tan:HumanBody"`           // optional
	ImageRef            xsd.AnyURI          `xml:"tan:ImageRef"`            // optional
	BarcodeInfo         BarcodeInfo         `xml:"tan:BarcodeInfo"`         // optional
	SphericalCoordinate SphericalCoordinate `xml:"tan:SphericalCoordinate"` // optional
	Label               LabelInfo           `xml:"tan:Label"`               // optional, unbounded
}

type LabelInfo struct {
	Likelihood xsd.Float   `xml:"tan:Likelihood"`
	Authority  xsd.String  `xml:"tan:Authority"`
	ID         xsd.AnyType `xml:"tan:ID"` // idk
}

type SphericalCoordinate struct {
	// TODO
}

type BarcodeInfo struct {
	Data StringLikelihood `xml:"tan:Data"`
	Type StringLikelihood `xml:"tan:Type"` // optional
	PPM  xsd.Float        `xml:"tan:PPM"`  // optional
}

type HumanBody struct {
	// TODO
}

type HumanFace struct {
	// TODO
}

type LicensePlateInfo struct {
	PlateNumber   StringLikelihood `xml:"tan:PlateNumber"`
	PlateType     StringLikelihood `xml:"tan:PlateType"`     // optional
	CountryCode   StringLikelihood `xml:"tan:CountryCode"`   // optional
	IssuingEntity StringLikelihood `xml:"tan:IssuingEntity"` // optional
}

type VehicleInfo struct {
	Type  StringLikelihood `xml:"tan:Type"`
	Brand StringLikelihood `xml:"tan:Brand"` // optional
	Model StringLikelihood `xml:"tan:Model"` // optional
}

type AppearanceExtension struct {
	// TODO
}

type ColorDescriptor struct {
}

type ClassType string

const (
	ClassTypeAnimal  ClassType = "Animal"
	ClassTypeFace    ClassType = "Face"
	ClassTypeHuman   ClassType = "Human"
	ClassTypeVehical ClassType = "Vehical"
	ClassTypeOther   ClassType = "Other"
)

type ClassDescriptor struct {
	ClassCandidate struct {
		Type       ClassType `xml:"tan:Type"`
		Likelihood xsd.Float `xml:"tan:Likelihood"`
	} // optional, unbounded
	Extension ClassDescriptorExtension `xml:"tan:Extension"` // optional
}

type ClassDescriptorExtension struct {
	OtherTypes OtherType                 `xml:"tan:OtherTypes"` // unbounded
	Extension  ClassDescriptorExtension2 `xml:"tan:Extension"`  // optional
	Type       StringLikelihood          `xml:"tan:Type"`       // optional, unbounded
}

type StringLikelihood = xsd.String // TODO

type ClassDescriptorExtension2 struct {
	// TODO
}

type OtherType struct {
	Type       xsd.String `xml:"tan:Type"`
	Likelihood xsd.Float  `xml:"tan:Likelihood"`
}

type ShapeDescriptor struct {
	BoundingBox     onvif.Rectangle          `xml:"tan:BoundingBox"`
	CenterOfGravity onvif.Vector             `xml:"tan:CenterOfGravity"`
	Polygol         Polygon                  `xml:"tan:Polygol"`   // optional, unbounded
	Extension       ShapeDescriptorExtension `xml:"tan:Extension"` // optional
}

type ShapeDescriptorExtension struct {
	// TODO
}

type Polygon struct {
	// TODO
}

type ObjectExtension struct { // optional
	// TODO
}

type SupportedAnalyticsModules struct {
	Limit                                *xsd.Int
	AnalyticsModuleContentSchemaLocation *xsd.String
	AnalyticsModuleDescription           []AnalyticsModuleDescription
}

type AnalyticsModuleDescription struct {
	Name         string `xml:"Name,attr"`
	Fixed        bool   `xml:"fixed,attr"`
	MaxInstances int    `xml:"maxInstances,attr"`
	Parameters   *Parameters
	Messages     *Messages
}

type AnalyticsModuleOptions struct {
	RuleType        string `xml:",attr"`
	Name            string `xml:",attr"`
	Type            string `xml:",attr"`
	AnalyticsModule string `xml:",attr"`
	IntRange        *IntRange
	StringItems     *StringItems
}

type IntRange struct {
	Min int
	Max int
}

type StringItems struct {
	Item []string
}

type SupportedRules struct {
	Limit                     *xsd.Int
	RuleContentSchemaLocation *xsd.String
	RuleDescription           []RuleDescription
}

type RuleDescription struct {
	Name         *xsd.String  `xml:",attr"`
	Fixed        *xsd.Boolean `xml:"fixed,attr"`
	MaxInstances *xsd.Int     `xml:"maxInstances,attr"`
	Parameters   Parameters
	Messages     Messages
}

type ItemListExtension xsd.AnyType

type RuleOptions struct {
	RuleType                  *xsd.String
	Name                      *xsd.String `xml:",attr"`
	Type                      *xsd.String `xml:",attr"`
	MinOccurs                 *xsd.String `xml:"minOccurs,attr"`
	MaxOccurs                 *xsd.String `xml:"maxOccurs,attr"`
	AnalyticsModule           *xsd.String
	IntRange                  *IntRange
	StringItems               *StringItems
	PolygonOptions            *PolygonOptions
	MotionRegionConfigOptions *MotionRegionConfigOptions
	StringList                *xsd.String
}

type PolygonOptions struct {
	VertexLimits VertexLimits
}

type VertexLimits struct {
	Min int
	Max int
}

type MotionRegionConfigOptions struct {
	DisarmSupport  bool
	PolygonSupport bool
	PolygonLimits  VertexLimits
}

type Capabilities struct {
	// Indication that the device supports the rules interface and the rules syntax
	RuleSupport xsd.Boolean

	// Indication that the device supports the scene analytics module interface
	AnalyticsModuleSupport xsd.Boolean

	// Indication that the device produces the cell based scene description
	CellBasedSceneDescriptionSupported xsd.Boolean

	// Indication that the device supports the GetRuleOptions operation on the rules interface
	RuleOptionsSupported xsd.Boolean

	// Indication that the device supports the [GetAnalyticsModuleOptions] operation on the
	// analytics interface
	AnalyticsModuleOptionsSupported xsd.Boolean

	// Indication that the device supports the [GetSupportedMetadata] operation
	SupportedMetadata xsd.Boolean

	// Indication what kinds of method that the device support for sending image, acceptable values
	// are defined in tt:ImageSendingType
	ImageSendingType []xsd.String
}

type Parameters struct {
	SimpleItemDescription  []SimpleItemDescription
	ElementItemDescription []ElementItemDescription
	Extension              *xsd.String
}

type SimpleItemDescription struct {
	Name  string `xml:",attr"`
	Type  string `xml:",attr"`
	Value string `xml:",attr"`
}

type ElementItemDescription struct {
	Name  string `xml:",attr"`
	Value string `xml:",attr"`
}

type Messages struct {
	IsProperty  *xsd.Boolean `xml:",attr"`
	Source      *Source
	Key         *Key
	Data        *Data
	Extension   *xsd.String
	ParentTopic *xsd.String
}

type Source struct {
	SimpleItemDescription  []SimpleItemDescription
	ElementItemDescription []ElementItemDescription
	Extension              *xsd.String
}

type Key struct {
	SimpleItemDescription  []SimpleItemDescription
	ElementItemDescription []ElementItemDescription
	Extension              *xsd.String
}

type Data struct {
	SimpleItemDescription  []SimpleItemDescription
	ElementItemDescription []ElementItemDescription
	Extension              *xsd.String
}
