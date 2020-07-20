package SqldbPool_249_0_test

type  SqldbPool_249_0_testProvider struct {

}

func  DataProvider() []interface{} {

    return []interface{}{
        []interface{}{
            map[string]interface{}{"tns":"172.30.0.6","user":"darkhold","password":"6532b3c13C1491FB","port":3306,"db":"allinone"},
            "*sql.DB",
            "",
            " - id=102",
        },
	}

}