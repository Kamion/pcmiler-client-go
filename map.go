package pcmiler

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
	Options         Options         `json:"Options"`
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
	ElevLimit          int          `json:"ElevLimit"`
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
