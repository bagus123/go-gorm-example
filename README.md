# go-gorm-example


### Summary
- model can extend gorm.Model, if extend gorm.Model will inject field ID, CreatedAt, UpdatedAt, and DeletedAt
- ID is primary key, but you can add custom ID using tag primary_key; example : `gorm:"primary_key"`
- table name using pluralized you create struct 'User' will create table 'users',
  but you also can make custom table name or make singular table name using db.SingularTable(true)
- column name will use snake case but you can override using tag column; example : `gorm:"column:beast_id"` 
- for model has DeletedAt will delete as soft delete
- if wanna delete permanently can do by `db.Unscoped().Delete(&order)`
- support several database : mysql, postgre, sqlite, mssql
- create record; will return true if primary key is blank, if already exist will return false (for same object)
- default values can set by tag default; example : `gorm:"default:'this is default value'"`
- support plain SQL and CRUD interface; example : Create, First, Take, Last, Find
- plain SQL using Where method
- support null value using pointer or scaner/valuer varible example : sql.NullInt64, sql.NullBool, sql.NullFloat64, sql.NullString, etc sql.Nullxxxx
- support transaction
- support hook (Before/After Create/Save/Update/Delete/Find)
- support scope for filtering
- support migration, AutoMigrate will ONLY create tables, missing columns and missing indexes, and WON’T change existing column’s type or delete unused columns to protect your data : example `db.AutoMigrate(&User{})`
- cascade can define using tag OnUpdate, OnDelete example `gorm:"ForeignKey:User.ID;OnUpdate:Cascade;OnDelete:Cascade"`




## Associations (relationship)
The relationship defines how the structs or models interact with each other

- "belongs to" relationship
- "has one" / "one to one" relationship
- "has many" / "one to many" relationship
- "many to many" relationship




sources
- https://gorm.io/docs/index.html
- https://jinzhu.me/gorm/
- https://www.mindbowser.com/golang-go-with-gorm/
- https://medium.com/aubergine-solutions/how-i-handled-null-possible-values-from-database-rows-in-golang-521fb0ee267
- http://blog.ralch.com/tutorial/golang-object-relation-mapping-with-gorm/
- http://learningprogramming.net/golang/gorm/call-stored-procedure-with-parameters-in-gorm/
