package SqldbPool_249_0_test

type  SqldbPool_249_0_testProvider struct {

}

func  DataProvider() []interface{} {

    return []interface{}{
        []interface{}{
            []byte(`{"Tns":"172.30.0.6","User":"darkhold","Password":"6532b3c13C1491FB","Port":3306,"DB":"allinone"}`),
            "*sql.DB",
            "",
            " - id=102",
        },
	}

}