package models

import (
	"go-clean/db"

	"gopkg.in/mgo.v2/bson"
)

type Machine struct {
	Sn         string `json:"_id" bson:"_id,omitempty"`
	PowerWall  string `json:"powerwall"  bson:"powerwall"`
	Pln        string `json:"pln"  bson:"pln"`
	SolarPanel string `json:"solarpanel"  bson:"solarpanel"`
	Turbine    string `json:"turbine"  bson:"turbine"`
	Status     string `json:"status"  bson:"status"`
}

type MachineModels struct{}

func (R *MachineModels) GetMachine(SN string) (data Machine, err error) {
	db := db.DBConnection{}
	db.NewConnection()
	defer db.Close()
	coll := db.Session.DB("db_acier").C("Machine")
	err = coll.Find(bson.M{"_id": SN}).One(&data)
	return data, err
}
