package di

import (
	"github.com/nmarsollier/cartgo/internal/cart"
	"github.com/nmarsollier/cartgo/internal/env"
	"github.com/nmarsollier/cartgo/internal/rabbit/rbtschema"
	"github.com/nmarsollier/cartgo/internal/services"
	"github.com/nmarsollier/commongo/db"
	"github.com/nmarsollier/commongo/httpx"
	"github.com/nmarsollier/commongo/log"
	"github.com/nmarsollier/commongo/rbt"
	"github.com/nmarsollier/commongo/security"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/x/mongo/driver/topology"
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
	ArticleValidatorPublisher() rbtschema.ArticleValidationPublisher
	PlacedDataPublisher() rbtschema.PlacedDataPublisher
	Service() services.Service
}

type Deps struct {
	CurrLog          log.LogRusEntry
	CurrHttpClient   httpx.HTTPClient
	CurrDatabase     *mongo.Database
	CurrSecRepo      security.SecurityRepository
	CurrSecSvc       security.SecurityService
	CurrCartColl     db.Collection
	CurrCartRepo     cart.CartRepository
	CurrCartSvc      cart.CartService
	CurrValPublisher rbtschema.ArticleValidationPublisher
	CurrPldPublisher rbtschema.PlacedDataPublisher
	CurrService      services.Service
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
	i.CurrSecRepo = security.NewSecurityRepository(i.Logger(), i.HttpClient(), env.Get().SecurityServerURL)
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

	cartCollection, err := db.NewCollection(i.CurrLog, i.Database(), "cart", IsDbTimeoutError)
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

func (i *Deps) Service() services.Service {
	if i.CurrService != nil {
		return i.CurrService
	}
	i.CurrService = services.NewService(
		i.Logger(),
		i.HttpClient(),
		i.CartService(),
		i.ArticleValidatorPublisher(),
		i.PlacedDataPublisher(),
	)

	return i.CurrService
}

func (i *Deps) ArticleValidatorPublisher() rbtschema.ArticleValidationPublisher {
	if i.CurrValPublisher != nil {
		return i.CurrValPublisher
	}

	i.CurrValPublisher, _ = rbt.NewRabbitPublisher[*rbtschema.ArticleValidationData](
		rbt.RbtLogger(env.Get().FluentURL, env.Get().ServerName, i.Logger().CorrelationId()),
		env.Get().RabbitURL,
		"article_exist",
		"direct",
		"article_exist",
	)

	return i.CurrValPublisher
}

func (i *Deps) PlacedDataPublisher() rbtschema.PlacedDataPublisher {

	if i.CurrPldPublisher != nil {
		return i.CurrPldPublisher
	}

	i.CurrPldPublisher, _ = rbt.NewRabbitPublisher[*rbtschema.PlacedData](
		rbt.RbtLogger(env.Get().FluentURL, env.Get().ServerName, i.Logger().CorrelationId()),
		env.Get().RabbitURL,
		"place_order",
		"direct",
		"place_order",
	)

	return i.CurrPldPublisher
}

// IsDbTimeoutError función a llamar cuando se produce un error de db
func IsDbTimeoutError(err error) {
	if err == topology.ErrServerSelectionTimeout {
		database = nil
		cartCollection = nil
	}
}
