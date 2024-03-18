package services

import (
	"go.uber.org/fx"

	"github.com/CoopHive/hive/services/dealmaker"
	"github.com/CoopHive/hive/services/jobcreator"
	"github.com/CoopHive/hive/services/mediator"
	"github.com/CoopHive/hive/services/resourceprovider"
	"github.com/CoopHive/hive/services/root"
	"github.com/CoopHive/hive/services/run"
	"github.com/CoopHive/hive/services/solver"
	"github.com/CoopHive/hive/services/version"
)

var Module = fx.Options(
	dealmaker.Module, // TODO: refactor to internal
	version.Module,
	run.Module,
	jobcreator.Module,
	resourceprovider.Module,
	mediator.Module,
	solver.Module,
	root.Module,
)

var ModuleWithoutRoot = fx.Options(
	dealmaker.Module,
	version.Module,
	run.Module,
	jobcreator.Module,
	resourceprovider.Module,
	mediator.Module,
	solver.Module,
)
