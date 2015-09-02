package mongo

import ( 
  "gopkg.in/mgo.v2"
  "log"
)

type Mongo struct {
  database  *mgo.Database
  session   *mgo.Session
}

func NewMongo(cs string, db string) *Mongo {
  session, err := mgo.Dial(cs)
  if err != nil {
    panic(err)
  }

  database := session.DB(db)

  return &Mongo {
    database : database,
    session  : session,
  }
}

func (this *Mongo) documentExists(collectionName string, id interface{}) bool {
  collection := this.database.C(collectionName)
  count, err := collection.FindId(id).Limit(1).Count()
  if err != nil {
      log.Print(err)
  }
  if count > 0 {
      return true
  }  
  return false
}

func (this *Mongo) createDocument(collectionName string, document ...interface{}) {
  collection := this.database.C(collectionName)
  err := collection.Insert(document)
  if err != nil {
    log.Print(err)
  }
}

func (this *Mongo) close() {
  this.session.Close()
}