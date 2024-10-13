package goqty

import "math"

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
	"<micro>":  makeUnit("prefix", []string{"µ", "u", "Micro", "micro"}, 1e-06, nil, nil),
	"<nano>":   makeUnit("prefix", []string{"n", "Nano", "nano"}, 1e-09, nil, nil),
	"<pico>":   makeUnit("prefix", []string{"p", "Pico", "pico"}, 1e-12, nil, nil),
	"<femto>":  makeUnit("prefix", []string{"f", "Femto", "femto"}, 1e-15, nil, nil),
	"<atto>":   makeUnit("prefix", []string{"a", "Atto", "atto"}, 1e-18, nil, nil),
	"<zepto>":  makeUnit("prefix", []string{"z", "Zepto", "zepto"}, 1e-21, nil, nil),
	"<yocto>":  makeUnit("prefix", []string{"y", "Yocto", "yocto"}, 1e-24, nil, nil),
}
var prefixesByAlias = makeUnitAliasMap(prefixes)

var units = map[string]Unit{
	"<1>":       unityUnit,
	"<decibel>": makeUnit("logarithmic", []string{"dB", "decibel", "decibels"}, 1, []string{"decibel"}, nil),

	// length
	"<meter>":        makeUnit("length", []string{"m", "meter", "meters", "metre", "metres"}, 1, []string{"meter"}, nil),
	"<inch>":         makeUnit("length", []string{"in", "inch", "inches", "\""}, 0.0254, []string{"meter"}, nil),
	"<foot>":         makeUnit("length", []string{"ft", "foot", "feet", "'"}, 0.3048, []string{"meter"}, nil),
	"<yard>":         makeUnit("length", []string{"yd", "yard", "yards"}, 0.9144, []string{"meter"}, nil),
	"<mile>":         makeUnit("length", []string{"mi", "mile", "miles"}, 1609.344, []string{"meter"}, nil),
	"<naut-mile>":    makeUnit("length", []string{"nmi", "naut-mile"}, 1852, []string{"meter"}, nil),
	"<league>":       makeUnit("length", []string{"league", "leagues"}, 4828, []string{"meter"}, nil),
	"<furlong>":      makeUnit("length", []string{"furlong", "furlongs"}, 201.2, []string{"meter"}, nil),
	"<rod>":          makeUnit("length", []string{"rd", "rod", "rods"}, 5.029, []string{"meter"}, nil),
	"<mil>":          makeUnit("length", []string{"mil", "mils"}, 0.0000254, []string{"meter"}, nil),
	"<angstrom>":     makeUnit("length", []string{"ang", "angstrom", "angstroms"}, 1e-10, []string{"meter"}, nil),
	"<fathom>":       makeUnit("length", []string{"fathom", "fathoms"}, 1.829, []string{"meter"}, nil),
	"<pica>":         makeUnit("length", []string{"pc", "pica", "picas"}, 0.00423333333, []string{"meter"}, nil),
	"<point>":        makeUnit("length", []string{"pt", "point", "points"}, 0.000352777778, []string{"meter"}, nil),
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
	"<kelvin>":     makeUnit("temperature", []string{"degK", "kelvin"}, 1.0, []string{"<kelvin>"}, nil),
	"<celsius>":    makeUnit("temperature", []string{"degC", "celsius", "celsius", "centigrade"}, 1.0, []string{"<kelvin>"}, nil),
	"<fahrenheit>": makeUnit("temperature", []string{"degF", "fahrenheit"}, 5.0/9.0, []string{"<kelvin>"}, nil),
	"<rankine>":    makeUnit("temperature", []string{"degR", "rankine"}, 5.0/9.0, []string{"<kelvin>"}, nil),
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
}
var unitsByAlias = makeUnitAliasMap(units)

// var valuesByUnitAlias = makeUnitValuesMap(units)
var outputs = makeOutputsMap(units)
var baseUnits = []string{"<meter>", "<kilogram>", "<second>", "<mole>", "<ampere>", "<radian>", "<kelvin>", "<temp-K>", "<byte>", "<dollar>", "<candela>", "<each>", "<steradian>", "<decibel>"}

// export var UNITS = {

//   /* energy */
//   "<joule>" :  [["J","joule","Joule","joules","Joules"], 1.0, "energy", ["<meter>","<meter>","<kilogram>"], ["<second>","<second>"]],
//   "<erg>"   :  [["erg","ergs"], 1e-7, "energy", ["<meter>","<meter>","<kilogram>"], ["<second>","<second>"]],
//   "<btu>"   :  [["BTU","btu","BTUs"], 1055.056, "energy", ["<meter>","<meter>","<kilogram>"], ["<second>","<second>"]],
//   "<calorie>" :  [["cal","calorie","calories"], 4.18400, "energy",["<meter>","<meter>","<kilogram>"], ["<second>","<second>"]],
//   "<Calorie>" :  [["Cal","Calorie","Calories"], 4184.00, "energy",["<meter>","<meter>","<kilogram>"], ["<second>","<second>"]],
//   "<therm-US>" : [["th","therm","therms","Therm","therm-US"], 105480400, "energy",["<meter>","<meter>","<kilogram>"], ["<second>","<second>"]],
//   "<Wh>" : [["Wh"], 3600, "energy",["<meter>","<meter>","<kilogram>"], ["<second>","<second>"]],
//   "<electronvolt>" : [["eV", "electronvolt", "electronvolts"], 1.602176634E-19, "energy", ["<meter>","<meter>","<kilogram>"], ["<second>","<second>"]],

//   /* force */
//   "<newton>"  : [["N","Newton","newton"], 1.0, "force", ["<kilogram>","<meter>"], ["<second>","<second>"]],
//   "<dyne>"  : [["dyn","dyne"], 1e-5, "force", ["<kilogram>","<meter>"], ["<second>","<second>"]],
//   "<pound-force>"  : [["lbf","pound-force"], 4.448222, "force", ["<kilogram>","<meter>"], ["<second>","<second>"]],
//   "<kilogram-force>"  : [["kgf","kilogram-force", "kilopond", "kp"], 9.80665, "force", ["<kilogram>","<meter>"], ["<second>","<second>"]],
//   "<gram-force>"  : [["gf","gram-force"], 9.80665E-3, "force", ["<kilogram>","<meter>"], ["<second>","<second>"]],

//   /* frequency */
//   "<hertz>" : [["Hz","hertz","Hertz"], 1.0, "frequency", ["<1>"], ["<second>"]],

//   /* angle */
//   "<radian>" :[["rad","radian","radians"], 1.0, "angle", ["<radian>"]],
//   "<degree>" :[["deg","degree","degrees"], Math.PI / 180.0, "angle", ["<radian>"]],
//   "<arcminute>" :[["arcmin","arcminute","arcminutes"], Math.PI / 10800.0, "angle", ["<radian>"]],
//   "<arcsecond>" :[["arcsec","arcsecond","arcseconds"], Math.PI / 648000.0, "angle", ["<radian>"]],
//   "<gradian>"   :[["gon","grad","gradian","grads"], Math.PI / 200.0, "angle", ["<radian>"]],
//   "<steradian>"  : [["sr","steradian","steradians"], 1.0, "solid_angle", ["<steradian>"]],

//   /* rotation */
//   "<rotation>" : [["rotation"], 2.0 * Math.PI, "angle", ["<radian>"]],
//   "<rpm>"   :[["rpm"], 2.0 * Math.PI / 60.0, "angular_velocity", ["<radian>"], ["<second>"]],

//   /* information */
//   "<byte>"  :[["B","byte","bytes"], 1.0, "information", ["<byte>"]],
//   "<bit>"  :[["b","bit","bits"], 0.125, "information", ["<byte>"]],

//   /* information rate */
//   "<Bps>" : [["Bps"], 1.0, "information_rate", ["<byte>"], ["<second>"]],
//   "<bps>" : [["bps"], 0.125, "information_rate", ["<byte>"], ["<second>"]],

//   /* currency */
//   "<dollar>":[["USD","dollar"], 1.0, "currency", ["<dollar>"]],
//   "<cents>" :[["cents"], 0.01, "currency", ["<dollar>"]],

//   /* luminosity */
//   "<candela>" : [["cd","candela"], 1.0, "luminosity", ["<candela>"]],
//   "<lumen>" : [["lm","lumen"], 1.0, "luminous_power", ["<candela>","<steradian>"]],
//   "<lux>" :[["lux"], 1.0, "illuminance", ["<candela>","<steradian>"], ["<meter>","<meter>"]],

//   /* power */
//   "<watt>"  : [["W","watt","watts"], 1.0, "power", ["<kilogram>","<meter>","<meter>"], ["<second>","<second>","<second>"]],
//   "<volt-ampere>"  : [["VA","volt-ampere"], 1.0, "power", ["<kilogram>","<meter>","<meter>"], ["<second>","<second>","<second>"]],
//   "<volt-ampere-reactive>"  : [["var","Var","VAr","VAR","volt-ampere-reactive"], 1.0, "power", ["<kilogram>","<meter>","<meter>"], ["<second>","<second>","<second>"]],
//   "<horsepower>"  :  [["hp","horsepower"], 745.699872, "power", ["<kilogram>","<meter>","<meter>"], ["<second>","<second>","<second>"]],

//   /* radiation */
//   "<gray>" : [["Gy","gray","grays"], 1.0, "radiation", ["<meter>","<meter>"], ["<second>","<second>"]],
//   "<roentgen>" : [["R","roentgen"], 0.009330, "radiation", ["<meter>","<meter>"], ["<second>","<second>"]],
//   "<sievert>" : [["Sv","sievert","sieverts"], 1.0, "radiation", ["<meter>","<meter>"], ["<second>","<second>"]],
//   "<becquerel>" : [["Bq","becquerel","becquerels"], 1.0, "radiation", ["<1>"],["<second>"]],
//   "<curie>" : [["Ci","curie","curies"], 3.7e10, "radiation", ["<1>"],["<second>"]],

//   /* rate */
//   "<cpm>" : [["cpm"], 1.0 / 60.0, "rate", ["<count>"],["<second>"]],
//   "<dpm>" : [["dpm"], 1.0 / 60.0, "rate", ["<count>"],["<second>"]],
//   "<bpm>" : [["bpm"], 1.0 / 60.0, "rate", ["<count>"],["<second>"]],

//   /* resolution / typography */
//   "<dot>" : [["dot","dots"], 1, "resolution", ["<each>"]],
//   "<pixel>" : [["pixel","px"], 1, "resolution", ["<each>"]],
//   "<ppi>" : [["ppi"], 1, "resolution", ["<pixel>"], ["<inch>"]],
//   "<dpi>" : [["dpi"], 1, "typography", ["<dot>"], ["<inch>"]],

//   /* other */
//   "<cell>" : [["cells","cell"], 1, "counting", ["<each>"]],
//   "<each>" : [["each"], 1.0, "counting", ["<each>"]],
//   "<count>" : [["count"], 1.0, "counting", ["<each>"]],
//   "<base-pair>"  : [["bp","base-pair"], 1.0, "counting", ["<each>"]],
//   "<nucleotide>" : [["nt","nucleotide"], 1.0, "counting", ["<each>"]],
//   "<molecule>" : [["molecule","molecules"], 1.0, "counting", ["<1>"]],
//   "<dozen>" :  [["doz","dz","dozen"],12.0,"prefix_only", ["<each>"]],
//   "<percent>": [["%","percent"], 0.01, "prefix_only", ["<1>"]],
//   "<ppm>" :  [["ppm"],1e-6, "prefix_only", ["<1>"]],
//   "<ppb>" :  [["ppb"],1e-9, "prefix_only", ["<1>"]],
//   "<ppt>" :  [["ppt"],1e-12, "prefix_only", ["<1>"]],
//   "<ppq>" :  [["ppq"],1e-15, "prefix_only", ["<1>"]],
//   "<gross>" :  [["gr","gross"],144.0, "prefix_only", ["<dozen>","<dozen>"]],
//   "<decibel>"  : [["dB","decibel","decibels"], 1.0, "logarithmic", ["<decibel>"]]
// };

// export var BASE_UNITS = ["<meter>","<kilogram>","<second>","<mole>", "<ampere>","<radian>","<kelvin>","<temp-K>","<byte>","<dollar>","<candela>","<each>","<steradian>","<decibel>"];

// export var UNITY = "<1>";
// export var UNITY_ARRAY = [UNITY];

// // Setup

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

// export var PREFIX_VALUES = {};
// export var PREFIX_MAP = {};
// export var UNIT_VALUES = {};
// export var UNIT_MAP = {};
// export var OUTPUT_MAP = {};
// for (var unitDef in UNITS) {
//   if (UNITS.hasOwnProperty(unitDef)) {
//     var definition = UNITS[unitDef];
//     if (definition[2] === "prefix") {
//       PREFIX_VALUES[unitDef] = definition[1];
//       for (var i = 0; i < definition[0].length; i++) {
//         PREFIX_MAP[definition[0][i]] = unitDef;
//       }
//     }
//     else {
//       validateUnitDefinition(unitDef, definition);
//       UNIT_VALUES[unitDef] = {
//         scalar: definition[1],
//         numerator: definition[3],
//         denominator: definition[4]
//       };
//       for (var j = 0; j < definition[0].length; j++) {
//         UNIT_MAP[definition[0][j]] = unitDef;
//       }
//     }
//     OUTPUT_MAP[unitDef] = definition[0][0];
//   }
// }

// /**
//  * Returns a list of available units of kind
//  *
//  * @param {string} [kind] - kind of units
//  * @returns {array} names of units
//  * @throws {QtyError} if kind is unknown
//  */
// export function getUnits(kind) {
//   var i;
//   var units = [];
//   var unitKeys = Object.keys(UNITS);
//   if (typeof kind === "undefined") {
//     for (i = 0; i < unitKeys.length; i++) {
//       if (["", "prefix"].indexOf(UNITS[unitKeys[i]][2]) === -1) {
//         units.push(unitKeys[i].substr(1, unitKeys[i].length - 2));
//       }
//     }
//   }
//   else if (this.getKinds().indexOf(kind) === -1) {
//     throw new QtyError("Kind not recognized");
//   }
//   else {
//     for (i = 0; i < unitKeys.length; i++) {
//       if (UNITS[unitKeys[i]][2] === kind) {
//         units.push(unitKeys[i].substr(1, unitKeys[i].length - 2));
//       }
//     }
//   }

//   return units.sort(function(a, b) {
//     if (a.toLowerCase() < b.toLowerCase()) {
//       return -1;
//     }
//     if (a.toLowerCase() > b.toLowerCase()) {
//       return 1;
//     }
//     return 0;
//   });
// }

// /**
//  * Returns a list of alternative names for a unit
//  *
//  * @param {string} unitName - name of unit
//  * @returns {string[]} aliases for unit
//  * @throws {QtyError} if unit is unknown
//  */
// export function getAliases(unitName) {
//   if (!UNIT_MAP[unitName]) {
//     throw new QtyError("Unit not recognized");
//   }
//   return UNITS[UNIT_MAP[unitName]][0];
// }

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
	return result
}
