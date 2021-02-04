package Update_275_0_test

type  Update_275_0_testProvider struct {

}

/**
更新数据*/
func  DataProvider() []interface{} {

    return []interface{}{
        []interface{}{
            //输入
            []byte(`{"Tns":"172.30.0.6","User":"darkhold","Password":"6532b3c13C1491FB","Port":3306,"DB":"allinone"}`),
            // 预期结论
            "c",
            //执行测试，附加参数
            map[string]interface{}{"sql":"update  a set a1='c' where id=?","binds":[]interface{}{1}},
            //测试案例-唯一id
            " - id=142",
        },
	}

}