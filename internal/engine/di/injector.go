package di

import (
	"github.com/nmarsollier/cartgo/internal/cart"
	"github.com/nmarsollier/cartgo/internal/engine/db"
	"github.com/nmarsollier/cartgo/internal/engine/env"
	"github.com/nmarsollier/cartgo/internal/engine/httpx"
	"github.com/nmarsollier/cartgo/internal/engine/log"
	"github.com/nmarsollier/cartgo/internal/rabbit/consume"
	"github.com/nmarsollier/cartgo/internal/rabbit/emit"
	"github.com/nmarsollier/cartgo/internal/security"
	"github.com/nmarsollier/cartgo/internal/services"
	"go.mongodb.org/mongo-driver/mongo"
)

// Singletons
var database *mongo.Database
var httpClient httpx.HTTPClient
var cartCollection db.Collection

type Injector interface {
	Logger() log.LogRusEntry
	Database() *mongo.Database
	HttpClient() httpx.HTTPClient
	SecurityRepository() security.SecurityRepository
	SecurityService() security.SecurityService
	CartCollection() db.Collection
	CartRepository() cart.CartRepository
	CartService() cart.CartService
	ArticleExistConsumer() consume.ArticleExistConsumer
	LogoutConsumer() consume.LogoutConsumer
	ConsumeOrderPlaced() consume.OrderPlacedConsumer
	RabbitEmiter() emit.RabbitEmiter
	RabbitChannel() emit.RabbitChannel
	Service() services.Service
}

type Deps struct {
	CurrLog        log.LogRusEntry
	CurrHttpClient httpx.HTTPClient
	CurrDatabase   *mongo.Database
	CurrSecRepo    security.SecurityRepository
	CurrSecSvc     security.SecurityService
	CurrCartColl   db.Collection
	CurrCartRepo   cart.CartRepository
	CurrCartSvc    cart.CartService
	CurrExistCons  consume.ArticleExistConsumer
	CurrLogoutCons consume.LogoutConsumer
	CurrOrderCons  consume.OrderPlacedConsumer
	CurrEmiter     emit.RabbitEmiter
	CurrChannel    emit.RabbitChannel
	CurrService    services.Service
}

func NewInjector(log log.LogRusEntry) Injector {
	return &Deps{
		CurrLog: log,
	}
}

func (i *Deps) Logger() log.LogRusEntry {
	return i.CurrLog
}

func (i *Deps) Database() *mongo.Database {
	if i.CurrDatabase != nil {
		return i.CurrDatabase
	}

	if database != nil {
		return database
	}

	database, err := db.NewDatabase(env.Get().MongoURL, "catalog")
	if err != nil {
		i.CurrLog.Fatal(err)
		return nil
	}

	return database
}

func (i *Deps) HttpClient() httpx.HTTPClient {
	if i.CurrHttpClient != nil {
		return i.CurrHttpClient
	}

	if httpClient != nil {
		return httpClient
	}

	httpClient = httpx.Get()
	return httpClient
}

func (i *Deps) SecurityRepository() security.SecurityRepository {
	if i.CurrSecRepo != nil {
		return i.CurrSecRepo
	}
	i.CurrSecRepo = security.NewSecurityRepository(i.Logger(), i.HttpClient())
	return i.CurrSecRepo
}

func (i *Deps) SecurityService() security.SecurityService {
	if i.CurrSecSvc != nil {
		return i.CurrSecSvc
	}
	i.CurrSecSvc = security.NewSecurityService(i.Logger(), i.SecurityRepository())
	return i.CurrSecSvc
}

func (i *Deps) CartCollection() db.Collection {
	if i.CurrCartColl != nil {
		return i.CurrCartColl
	}

	if cartCollection != nil {
		return cartCollection
	}

	cartCollection, err := db.NewCollection(i.CurrLog, i.Database(), "cart")
	if err != nil {
		i.CurrLog.Fatal(err)
		return nil
	}
	return cartCollection
}

func (i *Deps) CartRepository() cart.CartRepository {
	if i.CurrCartRepo != nil {
		return i.CurrCartRepo
	}
	i.CurrCartRepo = cart.NewCartRepository(i.Logger(), i.CartCollection())
	return i.CurrCartRepo
}

func (i *Deps) CartService() cart.CartService {
	if i.CurrCartSvc != nil {
		return i.CurrCartSvc
	}
	i.CurrCartSvc = cart.NewCartService(i.Logger(), i.CartRepository())
	return i.CurrCartSvc
}

func (i *Deps) ArticleExistConsumer() consume.ArticleExistConsumer {
	if i.CurrExistCons != nil {
		return i.CurrExistCons
	}
	i.CurrExistCons = consume.NewArticleExistConsumer(env.Get().FluentUrl, env.Get().RabbitURL, i.CartService())
	return i.CurrExistCons
}

func (i *Deps) LogoutConsumer() consume.LogoutConsumer {
	if i.CurrLogoutCons != nil {
		return i.CurrLogoutCons
	}
	i.CurrLogoutCons = consume.NewLogoutConsumer(env.Get().FluentUrl, env.Get().RabbitURL, i.SecurityService())
	return i.CurrLogoutCons
}

func (i *Deps) ConsumeOrderPlaced() consume.OrderPlacedConsumer {
	if i.CurrOrderCons != nil {
		return i.CurrOrderCons
	}
	i.CurrOrderCons = consume.NewOrderPlacedConsumer(env.Get().FluentUrl, env.Get().RabbitURL, i.CartService())
	return i.CurrOrderCons
}

func (i *Deps) RabbitChannel() emit.RabbitChannel {
	if i.CurrChannel != nil {
		return i.CurrChannel
	}

	chn, err := emit.NewChannel(env.Get().RabbitURL, i.Logger())
	if err != nil {
		i.Logger().Fatal(err)
		return nil
	}

	i.CurrChannel = chn
	return i.CurrChannel
}

func (i *Deps) RabbitEmiter() emit.RabbitEmiter {
	if i.CurrEmiter != nil {
		return i.CurrEmiter
	}
	i.CurrEmiter = emit.NewRabbitEmiter(i.Logger(), i.RabbitChannel())
	return i.CurrEmiter
}

func (i *Deps) Service() services.Service {
	if i.CurrService != nil {
		return i.CurrService
	}
	i.CurrService = services.NewService(i.Logger(), i.HttpClient(), i.CartService(), i.RabbitEmiter())
	return i.CurrService
}
