// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package cmd

import (
	"github.com/hashicorp/go-hclog"
	"github.com/leonardochaia/vendopunkto/clients"
	"github.com/leonardochaia/vendopunkto/internal/currency"
	"github.com/leonardochaia/vendopunkto/internal/invoice"
	"github.com/leonardochaia/vendopunkto/internal/pluginmgr"
	"github.com/leonardochaia/vendopunkto/internal/server"
	"github.com/leonardochaia/vendopunkto/internal/store"
	"github.com/leonardochaia/vendopunkto/internal/store/repositories"
)

import (
	_ "net/http/pprof"
)

// Injectors from wire.go:

func NewServer(globalLogger hclog.Logger) (*server.Server, error) {
	db, err := store.NewDB(globalLogger)
	if err != nil {
		return nil, err
	}
	invoiceRepository, err := repositories.NewPostgresInvoiceRepository(db)
	if err != nil {
		return nil, err
	}
	http := clients.NewHTTPClient()
	manager := pluginmgr.NewManager(globalLogger, http)
	invoiceManager, err := invoice.NewManager(invoiceRepository, manager, globalLogger)
	if err != nil {
		return nil, err
	}
	handler := invoice.NewHandler(invoiceManager, globalLogger, manager)
	vendoPunktoRouter, err := server.NewRouter(handler, globalLogger, db)
	if err != nil {
		return nil, err
	}
	currencyHandler, err := currency.NewHandler(manager, globalLogger)
	if err != nil {
		return nil, err
	}
	internalRouter, err := server.NewInternalRouter(handler, globalLogger, currencyHandler, db)
	if err != nil {
		return nil, err
	}
	serverServer, err := server.NewServer(vendoPunktoRouter, internalRouter, db, globalLogger, manager)
	if err != nil {
		return nil, err
	}
	return serverServer, nil
}
