package mongo

import ( 
  "gopkg.in/mgo.v2"
  "log"
)

type Mongo struct {
  Database  *mgo.Database
  Session   *mgo.Session
}

func NewMongo(cs string, db string) *Mongo {
  session, err := mgo.Dial(cs)
  if err != nil {
    panic(err)
  }

  database := session.DB(db)

  return &Mongo {
    Database : database,
    Session  : session,
  }
}

func (this *Mongo) FindAll(collectionName string, result interface{})  {
  err := this.Database.C(collectionName).Find(nil).All(result)

  if err != nil {
    log.Print(err)
  }
}


func (this *Mongo) Exists(collectionName string, id interface{}) bool {
  collection := this.Database.C(collectionName)
  count, err := collection.FindId(id).Limit(1).Count()
  if err != nil {
      log.Print(err)
  }
  if count > 0 {
      return true
  }  
  return false
}

func (this *Mongo) Create(collectionName string, document ...interface{}) {
  collection := this.Database.C(collectionName)
  err := collection.Insert(document)
  if err != nil {
    log.Print(err)
  }
}

func (this *Mongo) Close() {
  this.Session.Close()
}