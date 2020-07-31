package Rollback2_264_0_test

type  Rollback2_264_0_testProvider struct {

}

func  DataProvider() []interface{} {

    return []interface{}{
        []interface{}{
            []byte(`{"Tns":"172.30.0.6","User":"darkhold","Password":"6532b3c13C1491FB","Port":3306,"DB":"allinone"}`),
            1,
            []map[string]interface{}{
    {"sql":"update  a set dd=? where id =?","binds":[]interface{}{"wokao",13}},
    {"sql":"update  a set dd=? where id =?","binds":[]interface{}{"wokao2",14}},
    },
            " - id=128",
        },
	}

}