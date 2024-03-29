package config 

type Configuration interface{

	GetString(name string) (configValue string, found bool)
	GetInt(name string) (configValue int, found bool)
	GetBool(name string) (configValue bool, found bool)
	GetFloat(name string) (configValue float64, found bool)

	GetStringDefault(name, defVal string) (configValue string)
	GetIntDefault(name string, defVal int) (configValue int)
	GetBoolDefault(name string, defVal bool) (configbalue bool)

	GetSection(sectionName string) (section Configuration, found bool)
}