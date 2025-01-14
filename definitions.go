package goqty

import (
	"math"
	"slices"
)

type Unit struct {
	kind        string
	scalar      float64
	aliases     []string
	numerator   []string
	denominator []string
}

// type NormalizedUnit struct {
// 	prefix string // the unit name, eg. <meter>
// 	unit   string
// }

func makeUnit(kind string, aliases []string, scalar float64, numerator []string, denominator []string) Unit {
	return Unit{kind, scalar, aliases, numerator, denominator}
}

var unity = "<1>"
var unityArray = []string{unity}
var unityUnit = makeUnit("", []string{"1", "<1>"}, 1, nil, nil)

var prefixes = map[string]Unit{
	// prefixes
	"<googol>": makeUnit("prefix", []string{"googol"}, 1e100, nil, nil),
	"<kibi>":   makeUnit("prefix", []string{"Ki", "Kibi", "kibi"}, math.Pow(2, 10), nil, nil),
	"<mibi>":   makeUnit("prefix", []string{"Mi", "Mibi", "mibi"}, math.Pow(2, 20), nil, nil),
	"<gibi>":   makeUnit("prefix", []string{"Gi", "Gibi", "gibi"}, math.Pow(2, 30), nil, nil),
	"<tibi>":   makeUnit("prefix", []string{"Ti", "Tibi", "tibi"}, math.Pow(2, 40), nil, nil),
	"<pibi>":   makeUnit("prefix", []string{"Pi", "Pibi", "pibi"}, math.Pow(2, 50), nil, nil),
	"<eibi>":   makeUnit("prefix", []string{"Ei", "Eibi", "eibi"}, math.Pow(2, 60), nil, nil),
	"<zibi>":   makeUnit("prefix", []string{"Zi", "Zibi", "zibi"}, math.Pow(2, 70), nil, nil),
	"<yibi>":   makeUnit("prefix", []string{"Yi", "Yibi", "yibi"}, math.Pow(2, 80), nil, nil),
	"<yotta>":  makeUnit("prefix", []string{"Y", "Yotta", "yotta"}, 1e24, nil, nil),
	"<zetta>":  makeUnit("prefix", []string{"Z", "Zetta", "zetta"}, 1e21, nil, nil),
	"<exa>":    makeUnit("prefix", []string{"E", "Exa", "exa"}, 1e18, nil, nil),
	"<peta>":   makeUnit("prefix", []string{"P", "Peta", "peta"}, 1e15, nil, nil),
	"<tera>":   makeUnit("prefix", []string{"T", "Tera", "tera"}, 1e12, nil, nil),
	"<giga>":   makeUnit("prefix", []string{"G", "Giga", "giga"}, 1e09, nil, nil),
	"<mega>":   makeUnit("prefix", []string{"M", "Mega", "mega"}, 1e06, nil, nil),
	"<kilo>":   makeUnit("prefix", []string{"k", "Kilo", "kilo"}, 1e03, nil, nil),
	"<hecto>":  makeUnit("prefix", []string{"h", "Hecto", "hecto"}, 1e02, nil, nil),
	"<deca>":   makeUnit("prefix", []string{"da", "Deca", "deca", "Deka", "deka"}, 1e01, nil, nil),
	"<deci>":   makeUnit("prefix", []string{"d", "Deci", "deci"}, 1e-01, nil, nil),
	"<centi>":  makeUnit("prefix", []string{"c", "Centi", "centi"}, 1e-02, nil, nil),
	"<milli>":  makeUnit("prefix", []string{"m", "Milli", "milli"}, 1e-03, nil, nil),
	"<micro>":  makeUnit("prefix", []string{"\u00B5", "\u03BC", "u", "Micro", "micro"}, 1e-06, nil, nil),
	"<nano>":   makeUnit("prefix", []string{"n", "Nano", "nano"}, 1e-09, nil, nil),
	"<pico>":   makeUnit("prefix", []string{"p", "Pico", "pico"}, 1e-12, nil, nil),
	"<femto>":  makeUnit("prefix", []string{"f", "Femto", "femto"}, 1e-15, nil, nil),
	"<atto>":   makeUnit("prefix", []string{"a", "Atto", "atto"}, 1e-18, nil, nil),
	"<zepto>":  makeUnit("prefix", []string{"z", "Zepto", "zepto"}, 1e-21, nil, nil),
	"<yocto>":  makeUnit("prefix", []string{"y", "Yocto", "yocto"}, 1e-24, nil, nil),
}
var prefixesByAlias = makeUnitAliasMap(prefixes)

var units = map[string]Unit{
	"<1>": unityUnit,

	// length
	"<meter>":        makeUnit("length", []string{"m", "meter", "meters", "metre", "metres"}, 1, []string{"<meter>"}, nil),
	"<inch>":         makeUnit("length", []string{"in", "inch", "inches", "\""}, 0.0254, []string{"<meter>"}, nil),
	"<foot>":         makeUnit("length", []string{"ft", "foot", "feet", "'"}, 0.3048, []string{"<meter>"}, nil),
	"<yard>":         makeUnit("length", []string{"yd", "yard", "yards"}, 0.9144, []string{"<meter>"}, nil),
	"<mile>":         makeUnit("length", []string{"mi", "mile", "miles"}, 1609.344, []string{"<meter>"}, nil),
	"<naut-mile>":    makeUnit("length", []string{"nmi", "naut-mile"}, 1852, []string{"<meter>"}, nil),
	"<league>":       makeUnit("length", []string{"league", "leagues"}, 4828, []string{"<meter>"}, nil),
	"<furlong>":      makeUnit("length", []string{"furlong", "furlongs"}, 201.2, []string{"<meter>"}, nil),
	"<rod>":          makeUnit("length", []string{"rd", "rod", "rods"}, 5.029, []string{"<meter>"}, nil),
	"<mil>":          makeUnit("length", []string{"mil", "mils"}, 0.0000254, []string{"<meter>"}, nil),
	"<angstrom>":     makeUnit("length", []string{"ang", "angstrom", "angstroms"}, 1e-10, []string{"<meter>"}, nil),
	"<fathom>":       makeUnit("length", []string{"fathom", "fathoms"}, 1.829, []string{"<meter>"}, nil),
	"<pica>":         makeUnit("length", []string{"pc", "pica", "picas"}, 0.00423333333, []string{"<meter>"}, nil),
	"<point>":        makeUnit("length", []string{"pt", "point", "points"}, 0.000352777778, []string{"<meter>"}, nil),
	"<redshift>":     makeUnit("length", []string{"z", "red-shift", "redshift"}, 1.302773e26, []string{"<meter>"}, nil),
	"<AU>":           makeUnit("length", []string{"AU", "astronomical-unit"}, 149597900000, []string{"<meter>"}, nil),
	"<light-second>": makeUnit("length", []string{"ls", "light-second"}, 299792500, []string{"<meter>"}, nil),
	"<light-minute>": makeUnit("length", []string{"lmin", "light-minute"}, 17987550000, []string{"<meter>"}, nil),
	"<light-year>":   makeUnit("length", []string{"ly", "light-year"}, 9460528000000000, []string{"<meter>"}, nil),
	"<parsec>":       makeUnit("length", []string{"pc", "parsec", "parsecs"}, 30856780000000000, []string{"<meter>"}, nil),
	"<datamile>":     makeUnit("length", []string{"DM", "datamile"}, 1828.8, []string{"<meter>"}, nil),

	// mass
	"<kilogram>":   makeUnit("mass", []string{"kg", "kilogram", "kilograms"}, 1.0, []string{"<kilogram>"}, nil),
	"<AMU>":        makeUnit("mass", []string{"u", "AMU", "amu"}, 1.660538921e-27, []string{"<kilogram>"}, nil),
	"<dalton>":     makeUnit("mass", []string{"Da", "Dalton", "Daltons", "dalton", "daltons"}, 1.660538921e-27, []string{"<kilogram>"}, nil),
	"<slug>":       makeUnit("mass", []string{"slug", "slugs"}, 14.5939029, []string{"<kilogram>"}, nil),
	"<short-ton>":  makeUnit("mass", []string{"tn", "ton", "short-ton"}, 907.18474, []string{"<kilogram>"}, nil),
	"<metric-ton>": makeUnit("mass", []string{"t", "tonne", "metric-ton"}, 1000, []string{"<kilogram>"}, nil),
	"<carat>":      makeUnit("mass", []string{"ct", "carat", "carats"}, 0.0002, []string{"<kilogram>"}, nil),
	"<pound>":      makeUnit("mass", []string{"lbs", "lb", "pound", "pounds", "#"}, 0.45359237, []string{"<kilogram>"}, nil),
	"<ounce>":      makeUnit("mass", []string{"oz", "ounce", "ounces"}, 0.0283495231, []string{"<kilogram>"}, nil),
	"<gram>":       makeUnit("mass", []string{"g", "gram", "grams", "gramme", "grammes"}, 1e-3, []string{"<kilogram>"}, nil),
	"<grain>":      makeUnit("mass", []string{"grain", "grains", "gr"}, 6.479891e-5, []string{"<kilogram>"}, nil),
	"<dram>":       makeUnit("mass", []string{"dram", "drams", "dr"}, 0.0017718452, []string{"<kilogram>"}, nil),
	"<stone>":      makeUnit("mass", []string{"stone", "stones", "st"}, 6.35029318, []string{"<kilogram>"}, nil),

	// time
	"<second>":    makeUnit("time", []string{"s", "sec", "secs", "second", "seconds"}, 1.0, []string{"<second>"}, nil),
	"<minute>":    makeUnit("time", []string{"min", "mins", "minute", "minutes"}, 60.0, []string{"<second>"}, nil),
	"<hour>":      makeUnit("time", []string{"h", "hr", "hrs", "hour", "hours"}, 3600.0, []string{"<second>"}, nil),
	"<day>":       makeUnit("time", []string{"d", "day", "days"}, 3600*24, []string{"<second>"}, nil),
	"<week>":      makeUnit("time", []string{"wk", "week", "weeks"}, 7*3600*24, []string{"<second>"}, nil),
	"<fortnight>": makeUnit("time", []string{"fortnight", "fortnights"}, 1209600, []string{"<second>"}, nil),
	"<year>":      makeUnit("time", []string{"y", "yr", "year", "years", "annum"}, 31556926, []string{"<second>"}, nil),
	"<decade>":    makeUnit("time", []string{"decade", "decades"}, 315569260, []string{"<second>"}, nil),
	"<century>":   makeUnit("time", []string{"century", "centuries"}, 3155692600, []string{"<second>"}, nil),

	// substance
	"<mole>": makeUnit("substance", []string{"mol", "mole"}, 1.0, []string{"<mole>"}, nil),

	// current
	"<ampere>": makeUnit("current", []string{"A", "Ampere", "ampere", "amp", "amps"}, 1.0, []string{"<ampere>"}, nil),

	// area
	"<hectare>": makeUnit("area", []string{"hectare"}, 10000, []string{"<meter>", "<meter>"}, nil),
	"<acre>":    makeUnit("area", []string{"acre", "acres"}, 4046.85642, []string{"<meter>", "<meter>"}, nil),
	"<sqft>":    makeUnit("area", []string{"sqft"}, 1, []string{"<foot>", "<foot>"}, nil),

	// volume
	"<liter>":           makeUnit("volume", []string{"l", "L", "liter", "liters", "litre", "litres"}, 0.001, []string{"<meter>", "<meter>", "<meter>"}, nil),
	"<gallon>":          makeUnit("volume", []string{"gal", "gallon", "gallons"}, 0.0037854118, []string{"<meter>", "<meter>", "<meter>"}, nil),
	"<gallon-imp>":      makeUnit("volume", []string{"galimp", "gallon-imp", "gallons-imp"}, 0.0045460900, []string{"<meter>", "<meter>", "<meter>"}, nil),
	"<quart>":           makeUnit("volume", []string{"qt", "quart", "quarts"}, 0.00094635295, []string{"<meter>", "<meter>", "<meter>"}, nil),
	"<pint>":            makeUnit("volume", []string{"pt", "pint", "pints"}, 0.000473176475, []string{"<meter>", "<meter>", "<meter>"}, nil),
	"<pint-imp>":        makeUnit("volume", []string{"ptimp", "pint-imp", "pints-imp"}, 5.6826125e-4, []string{"<meter>", "<meter>", "<meter>"}, nil),
	"<cup>":             makeUnit("volume", []string{"cu", "cup", "cups"}, 0.000236588238, []string{"<meter>", "<meter>", "<meter>"}, nil),
	"<fluid-ounce>":     makeUnit("volume", []string{"floz", "fluid-ounce", "fluid-ounces"}, 2.95735297e-5, []string{"<meter>", "<meter>", "<meter>"}, nil),
	"<fluid-ounce-imp>": makeUnit("volume", []string{"flozimp", "floz-imp", "fluid-ounce-imp", "fluid-ounces-imp"}, 2.84130625e-5, []string{"<meter>", "<meter>", "<meter>"}, nil),
	"<tablespoon>":      makeUnit("volume", []string{"tb", "tbsp", "tbs", "tablespoon", "tablespoons"}, 1.47867648e-5, []string{"<meter>", "<meter>", "<meter>"}, nil),
	"<teaspoon>":        makeUnit("volume", []string{"tsp", "teaspoon", "teaspoons"}, 4.92892161e-6, []string{"<meter>", "<meter>", "<meter>"}, nil),
	"<bushel>":          makeUnit("volume", []string{"bu", "bsh", "bushel", "bushels"}, 0.035239072, []string{"<meter>", "<meter>", "<meter>"}, nil),
	"<oilbarrel>":       makeUnit("volume", []string{"bbl", "oilbarrel", "oilbarrels", "oil-barrel", "oil-barrels"}, 0.158987294928, []string{"<meter>", "<meter>", "<meter>"}, nil),
	"<beerbarrel>":      makeUnit("volume", []string{"bl", "bl-us", "beerbarrel", "beerbarrels", "beer-barrel", "beer-barrels"}, 0.1173477658, []string{"<meter>", "<meter>", "<meter>"}, nil),
	"<beerbarrel-imp>":  makeUnit("volume", []string{"blimp", "bl-imp", "beerbarrel-imp", "beerbarrels-imp", "beer-barrel-imp", "beer-barrels-imp"}, 0.16365924, []string{"<meter>", "<meter>", "<meter>"}, nil),

	// speed
	"<kph>":  makeUnit("speed", []string{"kph"}, 0.277777778, []string{"<meter>"}, []string{"<second>"}),
	"<mph>":  makeUnit("speed", []string{"mph"}, 0.44704, []string{"<meter>"}, []string{"<second>"}),
	"<knot>": makeUnit("speed", []string{"kt", "kn", "kts", "knot", "knots"}, 0.514444444, []string{"<meter>"}, []string{"<second>"}),
	"<fps>":  makeUnit("speed", []string{"fps"}, 0.3048, []string{"<meter>"}, []string{"<second>"}),

	// acceleration
	"<gee>": makeUnit("acceleration", []string{"gee"}, 9.80665, []string{"<meter>"}, []string{"<second>", "<second>"}),
	"<Gal>": makeUnit("acceleration", []string{"Gal"}, 1e-2, []string{"<meter>"}, []string{"<second>", "<second>"}),

	// temperature_difference
	"<kelvin>":     makeUnit("temperature", []string{"\u00b0K", "degK", "kelvin"}, 1.0, []string{"<kelvin>"}, nil),
	"<celsius>":    makeUnit("temperature", []string{"\u00b0C", "degC", "celsius", "celsius", "centigrade"}, 1.0, []string{"<kelvin>"}, nil),
	"<fahrenheit>": makeUnit("temperature", []string{"\u00b0F", "degF", "fahrenheit"}, 5.0/9.0, []string{"<kelvin>"}, nil),
	"<rankine>":    makeUnit("temperature", []string{"\u00b0R", "degR", "rankine"}, 5.0/9.0, []string{"<kelvin>"}, nil),
	"<temp-K>":     makeUnit("temperature", []string{"tempK", "temp-K"}, 1.0, []string{"<temp-K>"}, nil),
	"<temp-C>":     makeUnit("temperature", []string{"tempC", "temp-C"}, 1.0, []string{"<temp-K>"}, nil),
	"<temp-F>":     makeUnit("temperature", []string{"tempF", "temp-F"}, 5.0/9.0, []string{"<temp-K>"}, nil),
	"<temp-R>":     makeUnit("temperature", []string{"tempR", "temp-R"}, 5.0/9.0, []string{"<temp-K>"}, nil),

	// pressure
	"<pascal>": makeUnit("pressure", []string{"Pa", "pascal", "Pascal"}, 1.0, []string{"<kilogram>"}, []string{"<meter>", "<second>", "<second>"}),
	"<bar>":    makeUnit("pressure", []string{"bar", "bars"}, 100000, []string{"<kilogram>"}, []string{"<meter>", "<second>", "<second>"}),
	"<mmHg>":   makeUnit("pressure", []string{"mmHg"}, 133.322368, []string{"<kilogram>"}, []string{"<meter>", "<second>", "<second>"}),
	"<inHg>":   makeUnit("pressure", []string{"inHg"}, 3386.3881472, []string{"<kilogram>"}, []string{"<meter>", "<second>", "<second>"}),
	"<torr>":   makeUnit("pressure", []string{"torr"}, 133.322368, []string{"<kilogram>"}, []string{"<meter>", "<second>", "<second>"}),
	"<atm>":    makeUnit("pressure", []string{"atm", "ATM", "atmosphere", "atmospheres"}, 101325, []string{"<kilogram>"}, []string{"<meter>", "<second>", "<second>"}),
	"<psi>":    makeUnit("pressure", []string{"psi"}, 6894.76, []string{"<kilogram>"}, []string{"<meter>", "<second>", "<second>"}),
	"<cmh2o>":  makeUnit("pressure", []string{"cmH2O", "cmh2o"}, 98.0638, []string{"<kilogram>"}, []string{"<meter>", "<second>", "<second>"}),
	"<inh2o>":  makeUnit("pressure", []string{"inH2O", "inh2o"}, 249.082052, []string{"<kilogram>"}, []string{"<meter>", "<second>", "<second>"}),

	// viscosity
	"<poise>":  makeUnit("viscosity", []string{"P", "poise"}, 0.1, []string{"<kilogram>"}, []string{"<meter>", "<second>"}),
	"<stokes>": makeUnit("viscosity", []string{"St", "stokes"}, 1e-4, []string{"<meter>", "<meter>"}, []string{"<second>"}),

	// molar_concentration
	"<molar>":     makeUnit("molar_concentration", []string{"M", "molar"}, 1000, []string{"<mole>"}, []string{"<meter>", "<meter>", "<meter>"}),
	"<wtpercent>": makeUnit("molar_concentration", []string{"wt%", "wtpercent"}, 10, []string{"<kilogram>"}, []string{"<meter>", "<meter>", "<meter>"}),

	// activity
	"<katal>": makeUnit("activity", []string{"kat", "katal", "Katal"}, 1.0, []string{"<mole>"}, []string{"<second>"}),
	"<unit>":  makeUnit("activity", []string{"U", "enzUnit", "unit"}, 16.667e-16, []string{"<mole>"}, []string{"<second>"}),

	// capacitance
	"<farad>": makeUnit("capacitance", []string{"F", "farad", "Farad"}, 1.0, []string{"<second>", "<second>", "<second>", "<second>", "<ampere>", "<ampere>"}, []string{"<meter>", "<meter>", "<kilogram>"}),

	// charge
	"<coulomb>":           makeUnit("charge", []string{"C", "coulomb", "Coulomb"}, 1.0, []string{"<ampere>", "<second>"}, nil),
	"<Ah>":                makeUnit("charge", []string{"Ah"}, 3600, []string{"<ampere>", "<second>"}, nil),
	"<elementary-charge>": makeUnit("charge", []string{"e"}, 1.602176634e-19, []string{"<ampere>", "<second>"}, nil),

	// conductance
	"<siemens>": makeUnit("conductance", []string{"S", "Siemens", "siemens"}, 1.0, []string{"<second>", "<second>", "<second>", "<ampere>", "<ampere>"}, []string{"<kilogram>", "<meter>", "<meter>"}),

	// inductance
	"<henry>": makeUnit("inductance", []string{"H", "Henry", "henry"}, 1.0, []string{"<meter>", "<meter>", "<kilogram>"}, []string{"<second>", "<second>", "<ampere>", "<ampere>"}),

	// potential
	"<volt>": makeUnit("potential", []string{"V", "Volt", "volt", "volts"}, 1.0, []string{"<meter>", "<meter>", "<kilogram>"}, []string{"<second>", "<second>", "<second>", "<ampere>"}),

	// resistance
	// \u2126 is the Ohm sign, \u03a9 is the greek letter omega
	"<ohm>": makeUnit("resistance", []string{"\u2126", "\u03A9", "Ohm", "ohm"}, 1.0, []string{"<meter>", "<meter>", "<kilogram>"}, []string{"<second>", "<second>", "<second>", "<ampere>", "<ampere>"}),

	// magnetism
	"<weber>":   makeUnit("magnetism", []string{"Wb", "weber", "webers"}, 1.0, []string{"<meter>", "<meter>", "<kilogram>"}, []string{"<second>", "<second>", "<ampere>"}),
	"<tesla>":   makeUnit("magnetism", []string{"T", "tesla", "teslas"}, 1.0, []string{"<kilogram>"}, []string{"<second>", "<second>", "<ampere>"}),
	"<gauss>":   makeUnit("magnetism", []string{"G", "gauss"}, 1e-4, []string{"<kilogram>"}, []string{"<second>", "<second>", "<ampere>"}),
	"<maxwell>": makeUnit("magnetism", []string{"Mx", "maxwell", "maxwells"}, 1e-8, []string{"<meter>", "<meter>", "<kilogram>"}, []string{"<second>", "<second>", "<ampere>"}),
	"<oersted>": makeUnit("magnetism", []string{"Oe", "oersted", "oersteds"}, 250.0/math.Pi, []string{"<ampere>"}, []string{"<meter>"}),

	// energy
	"<joule>":        makeUnit("energy", []string{"J", "joule", "Joule", "joules", "Joules"}, 1.0, []string{"<meter>", "<meter>", "<kilogram>"}, []string{"<second>", "<second>"}),
	"<erg>":          makeUnit("energy", []string{"erg", "ergs"}, 1e-7, []string{"<meter>", "<meter>", "<kilogram>"}, []string{"<second>", "<second>"}),
	"<btu>":          makeUnit("energy", []string{"BTU", "btu", "BTUs"}, 1055.056, []string{"<meter>", "<meter>", "<kilogram>"}, []string{"<second>", "<second>"}),
	"<calorie>":      makeUnit("energy", []string{"cal", "calorie", "calories"}, 4.18400, []string{"<meter>", "<meter>", "<kilogram>"}, []string{"<second>", "<second>"}),
	"<Calorie>":      makeUnit("energy", []string{"Cal", "Calorie", "Calories"}, 4184.00, []string{"<meter>", "<meter>", "<kilogram>"}, []string{"<second>", "<second>"}),
	"<therm-US>":     makeUnit("energy", []string{"th", "therm", "therms", "Therm", "therm-US"}, 105480400, []string{"<meter>", "<meter>", "<kilogram>"}, []string{"<second>", "<second>"}),
	"<Wh>":           makeUnit("energy", []string{"Wh"}, 3600, []string{"<meter>", "<meter>", "<kilogram>"}, []string{"<second>", "<second>"}),
	"<electronvolt>": makeUnit("energy", []string{"eV", "electronvolt", "electronvolts"}, 1.602176634e-19, []string{"<meter>", "<meter>", "<kilogram>"}, []string{"<second>", "<second>"}),

	// force
	"<newton>":         makeUnit("force", []string{"N", "Newton", "newton"}, 1.0, []string{"<kilogram>", "<meter>"}, []string{"<second>", "<second>"}),
	"<dyne>":           makeUnit("force", []string{"dyn", "dyne"}, 1e-5, []string{"<kilogram>", "<meter>"}, []string{"<second>", "<second>"}),
	"<pound-force>":    makeUnit("force", []string{"lbf", "pound-force"}, 4.448222, []string{"<kilogram>", "<meter>"}, []string{"<second>", "<second>"}),
	"<kilogram-force>": makeUnit("force", []string{"kgf", "kilogram-force", "kilopond", "kp"}, 9.80665, []string{"<kilogram>", "<meter>"}, []string{"<second>", "<second>"}),
	"<gram-force>":     makeUnit("force", []string{"gf", "gram-force"}, 9.80665e-3, []string{"<kilogram>", "<meter>"}, []string{"<second>", "<second>"}),

	// frequency
	"<hertz>": makeUnit("frequency", []string{"Hz", "hertz", "Hertz"}, 1.0, []string{"<1>"}, []string{"<second>"}),

	// angle
	"<radian>":    makeUnit("angle", []string{"rad", "radian", "radians"}, 1.0, []string{"<radian>"}, nil),
	"<degree>":    makeUnit("angle", []string{"\u00b0", "deg", "degree", "degrees"}, math.Pi/180.0, []string{"<radian>"}, nil),
	"<arcminute>": makeUnit("angle", []string{"arcmin", "arcminute", "arcminutes"}, math.Pi/10800.0, []string{"<radian>"}, nil),
	"<arcsecond>": makeUnit("angle", []string{"arcsec", "arcsecond", "arcseconds"}, math.Pi/648000.0, []string{"<radian>"}, nil),
	"<gradian>":   makeUnit("angle", []string{"gon", "grad", "gradian", "grads"}, math.Pi/200.0, []string{"<radian>"}, nil),
	"<steradian>": makeUnit("solid_angle", []string{"sr", "steradian", "steradians"}, 1.0, []string{"<steradian>"}, nil),

	// rotation
	"<rotation>": makeUnit("angle", []string{"rotation"}, 2.0*math.Pi, []string{"<radian>"}, nil),
	"<rpm>":      makeUnit("angular_velocity", []string{"rpm"}, 2.0*math.Pi/60.0, []string{"<radian>"}, []string{"<second>"}),

	// information
	"<byte>": makeUnit("information", []string{"B", "byte", "bytes"}, 1.0, []string{"<byte>"}, nil),
	"<bit>":  makeUnit("information", []string{"b", "bit", "bits"}, 0.125, []string{"<byte>"}, nil),

	// information rate
	"<Bps>": makeUnit("information_rate", []string{"Bps"}, 1.0, []string{"<byte>"}, []string{"<second>"}),
	"<bps>": makeUnit("information_rate", []string{"bps"}, 0.125, []string{"<byte>"}, []string{"<second>"}),

	// currency
	// "<dollar>": makeUnit("currency", []string{"USD", "dollar"}, 1.0, []string{"<dollar>"}, nil),
	// "<cents>":  makeUnit("currency", []string{"cents"}, 0.01, []string{"<dollar>"}, nil),

	// luminosity
	"<candela>": makeUnit("luminosity", []string{"cd", "candela"}, 1.0, []string{"<candela>"}, nil),
	"<lumen>":   makeUnit("luminous_power", []string{"lm", "lumen"}, 1.0, []string{"<candela>", "<steradian>"}, nil),
	"<lux>":     makeUnit("illuminance", []string{"lux"}, 1.0, []string{"<candela>", "<steradian>"}, []string{"<meter>", "<meter>"}),

	// power
	"<watt>":                 makeUnit("power", []string{"W", "watt", "watts"}, 1.0, []string{"<kilogram>", "<meter>", "<meter>"}, []string{"<second>", "<second>", "<second>"}),
	"<volt-ampere>":          makeUnit("power", []string{"VA", "volt-ampere"}, 1.0, []string{"<kilogram>", "<meter>", "<meter>"}, []string{"<second>", "<second>", "<second>"}),
	"<volt-ampere-reactive>": makeUnit("power", []string{"var", "Var", "VAr", "VAR", "volt-ampere-reactive"}, 1.0, []string{"<kilogram>", "<meter>", "<meter>"}, []string{"<second>", "<second>", "<second>"}),
	"<horsepower>":           makeUnit("power", []string{"hp", "horsepower"}, 745.699872, []string{"<kilogram>", "<meter>", "<meter>"}, []string{"<second>", "<second>", "<second>"}),

	// radiation
	"<gray>":      makeUnit("radiation", []string{"Gy", "gray", "grays"}, 1.0, []string{"<meter>", "<meter>"}, []string{"<second>", "<second>"}),
	"<roentgen>":  makeUnit("radiation", []string{"R", "roentgen"}, 0.009330, []string{"<meter>", "<meter>"}, []string{"<second>", "<second>"}),
	"<sievert>":   makeUnit("radiation", []string{"Sv", "sievert", "sieverts"}, 1.0, []string{"<meter>", "<meter>"}, []string{"<second>", "<second>"}),
	"<becquerel>": makeUnit("radiation", []string{"Bq", "becquerel", "becquerels"}, 1.0, []string{"<1>"}, []string{"<second>"}),
	"<curie>":     makeUnit("radiation", []string{"Ci", "curie", "curies"}, 3.7e10, []string{"<1>"}, []string{"<second>"}),

	// rate
	"<cpm>": makeUnit("rate", []string{"cpm"}, 1.0/60.0, []string{"<count>"}, []string{"<second>"}),
	"<dpm>": makeUnit("rate", []string{"dpm"}, 1.0/60.0, []string{"<count>"}, []string{"<second>"}),
	"<bpm>": makeUnit("rate", []string{"bpm"}, 1.0/60.0, []string{"<count>"}, []string{"<second>"}),

	// resolution / typography
	"<dot>":   makeUnit("resolution", []string{"dot", "dots"}, 1, []string{"<each>"}, nil),
	"<pixel>": makeUnit("resolution", []string{"pixel", "px"}, 1, []string{"<each>"}, nil),
	"<ppi>":   makeUnit("resolution", []string{"ppi"}, 1, []string{"<pixel>"}, []string{"<inch>"}),
	"<dpi>":   makeUnit("typography", []string{"dpi"}, 1, []string{"<dot>"}, []string{"<inch>"}),

	// counting
	"<cell>":       makeUnit("counting", []string{"cells", "cell"}, 1, []string{"<each>"}, nil),
	"<each>":       makeUnit("counting", []string{"each"}, 1.0, []string{"<each>"}, nil),
	"<count>":      makeUnit("counting", []string{"count"}, 1.0, []string{"<each>"}, nil),
	"<base-pair>":  makeUnit("counting", []string{"bp", "base-pair"}, 1.0, []string{"<each>"}, nil),
	"<nucleotide>": makeUnit("counting", []string{"nt", "nucleotide"}, 1.0, []string{"<each>"}, nil),
	"<molecule>":   makeUnit("counting", []string{"molecule", "molecules"}, 1.0, []string{"<1>"}, nil),

	// prefix only
	"<dozen>":   makeUnit("prefix_only", []string{"doz", "dz", "dozen"}, 12.0, []string{"<each>"}, nil),
	"<percent>": makeUnit("prefix_only", []string{"%", "percent"}, 0.01, []string{"<1>"}, nil),
	"<ppm>":     makeUnit("prefix_only", []string{"ppm"}, 1e-6, []string{"<1>"}, nil),
	"<ppb>":     makeUnit("prefix_only", []string{"ppb"}, 1e-9, []string{"<1>"}, nil),
	"<ppt>":     makeUnit("prefix_only", []string{"ppt"}, 1e-12, []string{"<1>"}, nil),
	"<ppq>":     makeUnit("prefix_only", []string{"ppq"}, 1e-15, []string{"<1>"}, nil),
	"<gross>":   makeUnit("prefix_only", []string{"gr", "gross"}, 144.0, []string{"<dozen>", "<dozen>"}, nil),

	// logarithmic
	"<decibel>": makeUnit("logarithmic", []string{"dB", "decibel", "decibels"}, 1.0, []string{"<decibel>"}, nil),
}
var unitsByAlias = makeUnitAliasMap(units)

// var valuesByUnitAlias = makeUnitValuesMap(units)
var outputs = makeOutputsMap(units)
var baseUnits = []string{"<meter>", "<kilogram>", "<second>", "<mole>", "<ampere>", "<radian>", "<kelvin>", "<temp-K>", "<byte>", "<dollar>", "<candela>", "<each>", "<steradian>", "<decibel>"}

// /**
//  * Asserts unit definition is valid
//  *
//  * @param {string} unitDef - Name of unit to test
//  * @param {Object} definition - Definition of unit to test
//  *
//  * @returns {void}
//  * @throws {QtyError} if unit definition is not valid
//  */
// function validateUnitDefinition(unitDef, definition) {
//   var scalar = definition[1];
//   var numerator = definition[3] || [];
//   var denominator = definition[4] || [];
//   if (!isNumber(scalar)) {
//     throw new QtyError(unitDef + ": Invalid unit definition. " +
//                        "'scalar' must be a number");
//   }

//   numerator.forEach(function(unit) {
//     if (UNITS[unit] === undefined) {
//       throw new QtyError(unitDef + ": Invalid unit definition. " +
//                          "Unit " + unit + " in 'numerator' is not recognized");
//     }
//   });

//   denominator.forEach(function(unit) {
//     if (UNITS[unit] === undefined) {
//       throw new QtyError(unitDef + ": Invalid unit definition. " +
//                          "Unit " + unit + " in 'denominator' is not recognized");
//     }
//   });
// }

// returns a list of available units of kind
// returns an empty list if kind is unknown
func Units(kind string) []string {
	var result []string
	for name, unit := range units {
		if kind == "" || unit.kind == kind {
			result = append(result, name)
		}
	}
	slices.Sort(result)
	return result
}

// returns a list of unit aliases
// returns an empty list if unit is unknown
func UnitAliases(unit string) []string {
	if u, ok := units[unit]; ok {
		return u.aliases
	}
	return nil
}

// returns a list of unit aliases
// returns an empty list if unit is unknown
func PrefixAliases(prefix string) []string {
	if p, ok := prefixes[prefix]; ok {
		return p.aliases
	}
	return nil
}

func makeUnitAliasMap(units map[string]Unit) map[string]string {
	result := make(map[string]string)
	for name, unit := range units {
		for _, alias := range unit.aliases {
			result[alias] = name
		}
	}
	return result
}
func makeOutputsMap(units map[string]Unit) map[string]string {
	result := make(map[string]string)
	for name, unit := range units {
		result[name] = unit.aliases[0]
	}
	for name, prefix := range prefixes {
		result[name] = prefix.aliases[0]
	}
	return result
}
