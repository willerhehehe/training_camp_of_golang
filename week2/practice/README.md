## 问题
1. 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。

> 我认为，不应该Wrap这个error抛给上层，而应该是返回一个`nil`。
> 
>因为从调用者来看，调用dao层获取某一个数据对象，只关心能否拿到这个对象以及对象的值，而不是`sql.ErrorNoRows`这条错误。
>
>此外依据`Once an error is handled, it is not allowed to be passed up the call stack any longer.
`的原则，一旦错误被处理，就不应该再被Wrap并抛给上层。


eg: `practice/repository/PersonDAO.GetPerson`
```golang
func GetPerson(id int) (*models.Person, error) {
	db, err := sql.Open("mysql", "root:123456@/demo")
	if err != nil {
		err = errors.Wrap(err, fmt.Sprintf("DB connect error "))
		return nil, err
	}
	defer db.Close()
	rows := db.QueryRow("SELECT id, name, age FROM Person WHERE id = ?", id)
	var person = models.Person{}
	err = rows.Scan(&person.Id, &person.Name, &person.Age)
	if err != nil {
		return nil, nil
	}
	return &person, nil
}
```