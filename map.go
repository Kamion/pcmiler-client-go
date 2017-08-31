package pcmiler

// Map.Style
// Default		0	Indicates the default map style.
// Classic		1	Indicates a classic motif style.
// Monochrome	2	Indicates a monochrome style.
// RoadAtlas	3	Indicates a style that mimics a road atlas look and feel.
// Darkness		4	Indicates a dark style.
// Modern		5	Indicates a modern style.
// Contemporary	6	Indicates a contemporary style.
// Night		7	Indicates a style optimized for night viewing.
// Satelite		8	Indicates the style for satellite imagery.
// Lightness	9	Indicates the style for lightness style.
// Smooth		10	Indicates the style for smooth style.
// Terrain		11	Indicates the style of terrain imagery.

// Map.Drawers
// Airport					0	Indicates a airports.
// AvoidFavor				1	Indicates that avoids and favors should be shown on the map.
// City						2	Indicates cities.
// CityLarge				3	Indicates large cities.
// CommerciallyProhibited	4	Indicates commercially prohibited roads.
// CountryBoundaries		5	Indicates country boundaries.
// CountyBoundaries			6	Indicates county boundaries.
// HazMat					7	Indicates hazardous material restrictions.
// Land						8	Indicates land masses and in general this should always be included.
// LinkLabel				9	Indicates road names.
// MilitarBases				10	Indicates military bases.
// Network					11	Indicates roads.
// Ortholmage				12	Indicates satellite imagery.
// Parks					13	Indicates parks.
// Place					14	Indicates points of interest.
// PointDrawerWeb			15	Indicates user supplied points.
// Railroads				16	Indicates rail roads.
// Route					17	Indicates user supplied routes.
// Shield					18	Indicates road shields.
// ShieldInterstate			19	Indicates interstate road shields.
// StaaDesignation			20	Indicates truck designations.
// StateBoundaries			21	Indicates state and country boundaries.
// Stop						22	Indicates the name of stops along a route.
// TimeZone					23	Indicates time zone dividers.
// TruckRestrictions		24	Indicates truck restrictions.
// UrbanAreas				25	Indicates urban areas.
// Water					26	Indicates Water bodies such as oceans, lakes, river, etc.
// ExitLabel				27	Indicates exit labels for roads.
// PolygonLabel				28	Indicates polygon labels.
// Edge						29	Shows supertile edges - debugging only.

type Map struct {
	Viewport       Viewport   `json:"Viewport"`
	Projection     int        `json:"Projection"`
	Style          int        `json:"Style"`
	ImageOption    int        `json:"ImageOption"`
	Width          int        `json:"Width"`
	Height         int        `json:"Height"`
	Drawers        []int      `json:"Drawers"`
	LegendDrawer   []Legend   `json:"LegendDrawer"`
	GeometryDrawer []Geometry `json:"GeometryDrawer"`
	PinDrawer      PinDrawer  `json:"PinDrawer"`
}

type Route struct {
	RouteID         string          `json:"RouteId"`
	Stops           []StopLocation  `json:"Stops"`
	Options         *Options        `json:"Options"`
	FuelOptions     *FuelOptions    `json:"FuelOptions"`
	DrawLeastCost   bool            `json:"DrawLeastCost"`
	RouteLegOptions RouteLegOptions `json:"RouteLegOptions"`
	StopLabelDrawer int             `json:"StopLabelDrawer"`
}

type Options struct {
	AFSetIDs           string       `json:"AFSetIDs"`
	BordersOpen        bool         `json:"BordersOpen"`
	ClassOverrides     int          `json:"ClassOverrides"`
	DistanceUnits      int          `json:"DistanceUnits"`
	ElevLimit          int          `json:"ElevLimit,omitempty"`
	FerryDiscourage    bool         `json:"FerryDiscourage"`
	FuelRoute          bool         `json:"FuelRoute"`
	HazMatType         int          `json:"HazMatType"`
	HighwayOnly        bool         `json:"HighwayOnly"`
	HubRouting         bool         `json:"HubRouting"`
	OverrideRestrict   bool         `json:"OverrideRestrict"`
	RouteOptimization  int          `json:"RouteOptimization"`
	RoutingType        int          `json:"RoutingType"`
	TollDiscourage     bool         `json:"TollDiscourage"`
	TruckConfig        *TruckConfig `json:"TruckCfg,omitempty"`
	UseAvoidsAndFavors bool         `json:"UseAvoidsAndFavors"`
	VehicleType        int          `json:"VehicleType"`
}

type Viewport struct {
	Center       *Coordinates `json:"Center"`
	ScreenCenter *Point       `json:"ScreenCenter"`
	ZoomRadius   float64      `json:"ZoomRadius"`
	CornerA      *Coordinates `json:"CornerA"`
	CornerB      *Coordinates `json:"CornerB"`
	Region       int          `json:"Region"`
}

type Point struct {
	X int `json:"X"`
	Y int `json:"Y"`
}

type Legend struct {
	DrawOnMap bool `json:"DrawOnMap"`
	Type      int  `json:"Type"`
}

type Geometry struct {
	Color []RGB `json:"Color"`
}

type RGB struct {
	Red   int `json:"Red"`
	Blue  int `json:"Blue"`
	Green int `json:"Green"`
}

type Pin struct {
	Category int         `json:"Category"`
	ID       int         `json:"ID"`
	Image    string      `json:"Image"`
	Label    string      `json:"Label"`
	Point    Coordinates `json:"Point"`
}

type PinDrawer struct {
	DrawOnMap          bool  `json:"DrawOnMap"`
	Pins               []Pin `json:"Pins"`
	PointGroupDensity  int   `json:"PointGroupDensity"`
	PointSpreadInGroup int   `json:"PointSpreadInGroup"`
}

type StopLocation struct {
	Address    Address     `json:"Address"`
	Coords     Coordinates `json:"Coords"`
	Region     int         `json:"Region"`
	Label      string      `json:"Label"`
	PlaceName  string      `json:"PlaceName"`
	Costs      []Cost      `json:"Costs"`
	IsViaPoint bool        `json:"IsViaPoint"`
}

type Cost struct {
	CostOfStop   float64 `json:"CostOfStop"`
	HoursPerStop float64 `json:"HoursPerStop"`
	Loaded       bool    `json:"Loaded"`
	OnDuty       bool    `json:"OnDuty"`
	UseOrigin    bool    `json:"UseOrigin"`
}

type RouteLineOptions struct {
	Color RGB `json:"Color"`
	Width int `json:"Width"`
}

type LineOptions struct {
	RouteLineOptions RouteLineOptions `json:"RouteLineOptions"`
}

type RouteLabelOptions struct {
	Color    RGB `json:"Color"`
	FontSize int `json:"FontSize"`
}

type TextOptions struct {
	RouteLabelOptions RouteLabelOptions `json:"RouteLabelOptions"`
}

type RouteLegOptions struct {
	LineOptions LineOptions `json:"LineOptions"`
	TextOptions TextOptions `json:"TextOptions"`
}
