package cart

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/nmarsollier/cartgo/tools/db"
	"github.com/nmarsollier/cartgo/tools/errs"
	"github.com/nmarsollier/cartgo/tools/log"
	uuid "github.com/satori/go.uuid"
)

var tableName = "cart"

var ErrID = errs.NewValidation().Add("id", "Invalid")

func newCart(userId string) *Cart {
	return &Cart{
		ID:       uuid.NewV4().String(),
		UserId:   userId,
		Enabled:  true,
		Created:  time.Now(),
		Updated:  time.Now(),
		Articles: []*Article{},
	}
}

// findByUserId lee el cart activo del usuario
func findByUserId(cartId string, deps ...interface{}) (cart *Cart, err error) {
	expr, err := expression.NewBuilder().WithKeyCondition(
		expression.Key("userId_enabled").Equal(expression.Value(cartId + "_" + strconv.FormatBool(true))),
	).Build()

	if err != nil {
		return
	}

	response, err := db.Get(deps...).Query(context.TODO(), &dynamodb.QueryInput{
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

	err = attributevalue.UnmarshalMap(response.Items[0], &cart)
	if err != nil {
		return
	}

	return
}

func findById(cartId string, deps ...interface{}) (cart *Cart, err error) {
	response, err := db.Get(deps...).GetItem(context.TODO(), &dynamodb.GetItemInput{
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{
				Value: cartId,
			}},
		TableName: &tableName,
	})

	if err != nil || response == nil || response.Item == nil {
		log.Get(deps...).Error(err)

		return
	}

	err = attributevalue.UnmarshalMap(response.Item, &cart)
	if err != nil {
		log.Get(deps...).Error(err)
	}

	return
}

func save(cart *Cart, deps ...interface{}) (err error) {
	if err = cart.validateSchema(); err != nil {
		log.Get(deps...).Error(err)
		return
	}

	cart.UserIdEnabled = cart.UserId + "_" + strconv.FormatBool(cart.Enabled)
	articleToInsert, err := attributevalue.MarshalMap(cart)
	if err != nil {
		log.Get(deps...).Error(err)

		return
	}

	_, err = db.Get(deps...).PutItem(
		context.TODO(),
		&dynamodb.PutItemInput{
			TableName: &tableName,
			Item:      articleToInsert,
		},
	)

	if err != nil {
		log.Get(deps...).Error(err)
	}

	return
}

func invalidate(cart *Cart, deps ...interface{}) (err error) {
	if err = cart.validateSchema(); err != nil {
		log.Get(deps...).Error(err)
		return
	}

	key, err := attributevalue.MarshalMap(map[string]interface{}{
		"id": cart.ID,
	})
	if err != nil {
		return
	}

	update, err := attributevalue.MarshalMap(map[string]interface{}{
		":enabled":        false,
		":userId_enabled": cart.UserId + "_" + strconv.FormatBool(false),
	})
	if err != nil {
		return
	}

	_, err = db.Get(deps...).UpdateItem(
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
