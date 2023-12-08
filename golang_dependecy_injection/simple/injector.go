//go:build wireinject
// +build wireinject

package simple

import "github.com/google/wire"

func InitializeService(isError bool) (*SimpleService, error) {
	wire.Build(NewSimpleRepository, NewSimpleService)
	return nil, nil
}

var fooSet = wire.NewSet(NewFooRepository, NewFooService)
var barSet = wire.NewSet(NewBarRepository, NewBarService)

func InitializeFooBarService() *FooBarService {
	wire.Build(fooSet, barSet, NewFooBarService)
	return nil
}

var HelloSet = wire.NewSet(NewSayHelloImpl, wire.Bind(new(SayHello), new(*SayHelloImpl)))

func InitializeHelloService() *HelloService {
	wire.Build(HelloSet, NewHelloService)
	return nil
}

var FooBarSet = wire.NewSet(NewFoo, NewBar)

func InitializeFooBar() *FooBar {
	wire.Build(FooBarSet, wire.Struct(new(FooBar), "Foo", "Bar"))
	return nil
}

func InitializeFooBarWithValue() *FooBar {
	wire.Build(wire.Value(&Foo{}), wire.Value(&Bar{}), wire.Struct(new(FooBar), "Foo", "Bar"))
	return nil
}

func InitializeConfiguration() *Configuration {
	wire.Build(NewApplication, wire.FieldsOf(new(*Application), "Configuration"))
	return nil
}

func InitializeConnection(name string) (*Connection, func()) {
	wire.Build(NewConnection, NewFile)
	return nil, nil
}