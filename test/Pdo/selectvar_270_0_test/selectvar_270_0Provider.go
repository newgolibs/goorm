package selectvar_270_0_test

type  Selectvar_270_0_testProvider struct {

}

func  DataProvider() []interface{} {

    return []interface{}{
        []interface{}{
            []byte(`{"Tns":"172.30.0.6","User":"darkhold","Password":"6532b3c13C1491FB","Port":3306,"DB":"allinone"}`),
            "3",
            "select count(*) from a",
            " - id=134",
        },
        []interface{}{
            []byte(`{"Tns":"172.30.0.6","User":"darkhold","Password":"6532b3c13C1491FB","Port":3306,"DB":"allinone"}`),
            0,
            "select count(*) from a",
            "使用中间件查询 - id=135",
        },
	}

}