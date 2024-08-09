package services

import (
	"github.com/appwrite/sdk-for-go/databases"
	"github.com/appwrite/sdk-for-go/permission"
	"github.com/open-runtimes/types-for-go/v4/openruntimes"
)

func DoesDatabaseExist(dbs databases.Databases, dbId string) bool {
	_, err := dbs.Get(dbId)
	if err != nil {
		return false
	}
	return true
}

func DoesCollectionExist(dbs databases.Databases, dbId string, collId string) bool {
	_, err := dbs.GetCollection(dbId, collId)
	if err != nil {
		return false
	}
	return true
}

func DoesAttributeExist(dbs databases.Databases, dbId string, collId string, attribId string) bool {
	_, err := dbs.GetAttribute(dbId, collId, attribId)
	if err != nil {
		return false
	}
	return true
}

func InitialiseDatabase(Context openruntimes.Context, dbs databases.Databases, dbId string, collId string) {
	doesDbExist := DoesDatabaseExist(dbs, dbId)
	if !doesDbExist {
		dbs.Create(
			dbId,
			"URL Databases",
		)
	}

	doesCollExist := DoesCollectionExist(dbs, dbId, collId)
	if !doesCollExist {
		dbs.CreateCollection(
			dbId,
			collId,
			"URLs",
			dbs.WithCreateCollectionPermissions([]string{permission.Read("any")}),
		)
	}

	doesAttribExist := DoesAttributeExist(dbs, dbId, collId, "longUrl")
	if !doesAttribExist {
		dbs.CreateUrlAttribute(
			dbId,
			collId,
			"longUrl",
			true,
			dbs.WithCreateUrlAttributeArray(false),
		)
	}
}
