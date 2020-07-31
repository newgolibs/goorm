package SelectOne_252_0_test

type  SelectOne_252_0_testProvider struct {

}

func  DataProvider() []interface{} {

    return []interface{}{
        []interface{}{
            []byte(`{"Tns":"172.30.0.6","User":"darkhold","Password":"6532b3c13C1491FB","Port":3306,"DB":"allinone"}`),
            "{\"a1\":\"a1\",\"dd\":\"hahaa\",\"id\":\"10\"}",
            map[string]interface{}{"sql":"select id,a1,dd from  a where id=?","binds":[]interface{}{10}},
            " - id=106",
        },
        []interface{}{
            []byte(`{"Tns":"172.30.0.6","User":"darkhold","Password":"6532b3c13C1491FB","Port":3306,"DB":"allinone"}`),
            "{}",
            map[string]interface{}{"sql":"select id,a1,dd from  a where id=?","binds":[]interface{}{0}},
            " - id=110",
        },
	}

}