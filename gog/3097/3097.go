package main

import "gopkg.in/mgo.v2/bson"

type M bson.M
type UniqueID struct{}

var fDeleted = "demo"
var fLocation_uid = "demo"
var fInDateTime = "demo"

func xxx(locID UniqueID, fromDateTime, toDateTime int64) {
	selector := M{
		fDeleted:      false,
		fLocation_uid: locID,
		"$and":        []M{{fInDateTime: M{"$gte": fromDateTime}}, {fInDateTime: M{"$lte": toDateTime}}},
	}

	_ = selector

	// more code down here
}

func main() {
	xxx(UniqueID{}, 1, 1)
}
