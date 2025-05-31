package forgefy

import (
	"fmt"
	"text/template"

	"github.com/followthepattern/forgefy/plugins"
	"github.com/followthepattern/forgefy/productmap"
	"github.com/followthepattern/forgefy/specification"
	"github.com/followthepattern/forgefy/specification/defaults"
	"github.com/followthepattern/forgefy/specification/models"
	"github.com/followthepattern/forgefy/specification/naming"
	"github.com/followthepattern/forgefy/specification/parsing"
	"github.com/followthepattern/forgefy/specification/types"
)

type Builder struct {
	productSpecification specification.Product
	plugins              []plugins.Plugin
	apps                 map[string][]plugins.App
}

func newBuilder(productSpecification specification.Product) Builder {
	return Builder{
		productSpecification: productSpecification,
		plugins:              make([]plugins.Plugin, 0),
		apps:                 make(map[string][]plugins.App),
	}
}

func (builder Builder) withPlugins(plugins ...plugins.Plugin) Builder {
	builder.plugins = plugins
	return builder
}

func (builder Builder) withApps(apps map[string][]plugins.App) Builder {
	builder.apps = apps
	return builder
}

func (builder Builder) Build(plugins ...plugins.Plugin) (productmap.ProductMap, error) {
	pm := productmap.NewProductMap()

	err := builder.addDefaultFiles(pm)
	if err != nil {
		return pm, err
	}

	pm = builder.addParsingFunctions(pm)

	for _, plugin := range builder.plugins {
		err = plugin.Build(pm, builder.productSpecification)
		if err != nil {
			return pm, err
		}
	}

	for _, appSpecification := range builder.productSpecification.Apps {
		err = builder.buildApps(pm, appSpecification)
		if err != nil {
			return pm, err
		}
	}

	return pm, nil
}

func (builder Builder) addParsingFunctions(pm productmap.ProductMap) productmap.ProductMap {
	pm.WithFuncMap(template.FuncMap{
		"slist":           parsing.SList,
		"IsID":            parsing.CreateIsID(types.Registered),
		"NoneID":          parsing.CreateNoneID(types.Registered),
		"TableViewFilter": parsing.CreateTableViewFilter(types.Registered),
		"IsUndefined":     parsing.CreateIsUndefined(types.Registered),
		"IsBoolean":       parsing.CreateIsBoolean(types.Registered),
		"HasBoolean":      parsing.CreateHasBoolean(types.Registered),
		"IsNumber":        parsing.CreateIsNumber(types.Registered),
		"HasNumber":       parsing.CreateHasNumber(types.Registered),
		"IsString":        parsing.CreateIsString(types.Registered),
		"HasString":       parsing.CreateHasString(types.Registered),
		"IsText":          parsing.CreateIsText(types.Registered),
		"HasText":         parsing.CreateHasText(types.Registered),
		"IsTime":          parsing.CreateIsTime(types.Registered),
		"HasTime":         parsing.CreateHasTime(types.Registered),
		"IsNotTimeBased":  parsing.CreateIsNotTimeBased(types.Registered),
		"IsDate":          parsing.CreateIsDate(types.Registered),
		"HasDate":         parsing.CreateHasDate(types.Registered),
		"IsFile":          parsing.CreateIsFile(types.Registered),
		"HasFile":         parsing.CreateHasFile(types.Registered),
		"Records":         parsing.CreateRecordsFunc(types.Registered),
		"EnvVariable":     naming.ToEnvVariable,
		"IDField":         parsing.CreateFindIDFunc(types.Registered),
		"IsType":          parsing.CreateIsType(types.Registered),
		"LowerSnakeCase":  naming.ToLowerSnakeCase,
		"CamelCase":       naming.ToCamelCase,
		"UpperCamelCase":  naming.ToUpperCamelCase,
		"LowerCamelCase":  naming.ToLowerCamelCase,
		"FileName":        naming.ToFileName,
		"NetworkName":     naming.ToNetworkName,
		"HumanReadable":   naming.ToHumanReadable,
	})

	return pm
}

func (builder Builder) addDefaultFiles(_ productmap.ProductMap) error { return nil }

func (b Builder) buildApps(pm productmap.ProductMap, appSpecification specification.App) error {
	apps, ok := b.apps[string(appSpecification.AppType)]
	if !ok {
		return fmt.Errorf("unknown app definition: %s", appSpecification.AppType)
	}

	appSpecification = b.setAppDefaults(appSpecification)

	for _, app := range apps {
		err := app.Build(pm, appSpecification)

		if err != nil {
			return err
		}
	}

	return nil
}

func (b Builder) setAppDefaults(app specification.App) specification.App {
	admin := defaults.AdminUser()
	app.Defaults.Users = []models.User{admin}

	roles := defaults.Roles(admin, app.FeaturesArray())
	app.Defaults.Roles = roles
	app.Defaults.UserRoles = defaults.UserRole(admin, roles)

	app.CountOfRandomValues = 30

	for i := range app.Features {
		if app.Features[i].FeatureType == "" {
			app.Features[i].FeatureType = specification.CRUD
		}
	}

	return app
}
