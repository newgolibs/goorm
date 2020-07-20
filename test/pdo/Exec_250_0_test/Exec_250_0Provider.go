package Exec_250_0_test

type  Exec_250_0_testProvider struct {

}

func  DataProvider() []interface{} {

    return []interface{}{
        []interface{}{
            map[string]interface{}{"tns":"172.30.0.6","user":"darkhold","password":"6532b3c13C1491FB","port":3306,"db":"allinone","sql":"insert into a (a1,dd) values(?,?)","binds":[]interface{}{"cc","ok"}},
            1,
            "",
            " - id=103",
        },
	}

}