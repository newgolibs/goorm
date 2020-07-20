package LinkString_248_0_test

type  LinkString_248_0_testProvider struct {

}

func  DataProvider() []interface{} {

    return []interface{}{
        []interface{}{
            map[string]interface{}{"tns":"127.0.0.1","user":"a","password":"b","port":3306,"db":"allinone"},
            "a:b@tcp(127.0.0.1:3306)/allinone?charset=utf8mb4",
            "",
            " - id=101",
        },
	}

}