package selectvarMiddle_271_0_test

type  SelectvarMiddle_271_0_testProvider struct {

}

func  DataProvider() []interface{} {

    return []interface{}{
        []interface{}{
            []byte(`{"Tns":"172.30.0.6","User":"darkhold","Password":"6532b3c13C1491FB","Port":3306,"DB":"allinone"}`),
            0,
            "select count(*) from a1",
            "使用中间件查询 - id=136",
        },
	}

}