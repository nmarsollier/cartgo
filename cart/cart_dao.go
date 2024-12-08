package cart

import (
	"context"
	"errors"
	"strconv"
	"sync"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/nmarsollier/cartgo/tools/env"
	"github.com/nmarsollier/cartgo/tools/errs"
)

var tableName = "cart"

var (
	once     sync.Once
	instance CartDao
)

type CartDao interface {
	FindById(id string) (token *Cart, err error)
	FindByUserId(userId string) (token *Cart, err error)
	Save(article *Cart) (err error)
	Disable(cartId string, userId string) (err error)
}

func GetCartDao(deps ...interface{}) (CartDao, error) {
	for _, o := range deps {
		if client, ok := o.(CartDao); ok {
			return client, nil
		}
	}

	var conn_err error
	once.Do(func() {
		customCreds := aws.NewCredentialsCache(credentials.NewStaticCredentialsProvider(
			env.Get().AwsAccessKeyId,
			env.Get().AwsSecret,
			"",
		))

		cfg, err := config.LoadDefaultConfig(context.TODO(),
			config.WithRegion(env.Get().AwsRegion),
			config.WithCredentialsProvider(customCreds),
		)
		if err != nil {
			conn_err = err
			return
		}

		instance = &cartDao{
			client: dynamodb.NewFromConfig(cfg),
		}
	})

	return instance, conn_err
}

type cartDao struct {
	client *dynamodb.Client
}

func (r *cartDao) FindById(id string) (cart *Cart, err error) {
	response, err := r.client.GetItem(context.TODO(), &dynamodb.GetItemInput{
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{
				Value: id,
			}},
		TableName: &tableName,
	})

	if err != nil || response == nil || response.Item == nil {
		return nil, err
	}

	err = attributevalue.UnmarshalMap(response.Item, &cart)
	return
}

func (r *cartDao) FindByUserId(id string) (*Cart, error) {

	expr, err := expression.NewBuilder().WithKeyCondition(
		expression.Key("userId_enabled").Equal(expression.Value(id + "_" + strconv.FormatBool(true))),
	).Build()

	if err != nil {
		return nil, err
	}

	response, err := r.client.Query(context.TODO(), &dynamodb.QueryInput{
		TableName:                 &tableName,
		IndexName:                 aws.String("userId_enabled-index"),
		KeyConditionExpression:    expr.KeyCondition(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
	})

	if temp := new(types.ResourceNotFoundException); err != nil && !errors.As(err, &temp) {
		return nil, errs.NotFound
	}

	if err != nil || len(response.Items) == 0 {
		return nil, errs.NotFound
	}

	var cart Cart
	err = attributevalue.UnmarshalMap(response.Items[0], &cart)
	if err != nil {
		return nil, err
	}

	return &cart, nil
}

func (r *cartDao) Save(cart *Cart) (err error) {
	cart.UserIdEnabled = cart.UserId + "_" + strconv.FormatBool(cart.Enabled)
	articleToInsert, err := attributevalue.MarshalMap(cart)
	if err != nil {
		return
	}

	_, err = r.client.PutItem(
		context.TODO(),
		&dynamodb.PutItemInput{
			TableName: &tableName,
			Item:      articleToInsert,
		},
	)
	return
}

func (r *cartDao) Disable(cartId string, userId string) (err error) {
	key, err := attributevalue.MarshalMap(map[string]interface{}{
		"id": cartId,
	})
	if err != nil {
		return
	}

	update, err := attributevalue.MarshalMap(map[string]interface{}{
		":enabled":        false,
		":userId_enabled": userId + "_" + strconv.FormatBool(false),
	})
	if err != nil {
		return
	}

	_, err = r.client.UpdateItem(
		context.TODO(),
		&dynamodb.UpdateItemInput{
			TableName:                 &tableName,
			Key:                       key,
			UpdateExpression:          aws.String("SET enabled = :enabled, userId_enabled = :userId_enabled"),
			ExpressionAttributeValues: update,
		},
	)
	return
}
