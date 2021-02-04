package Insert_276_0_test

type  Insert_276_0_testProvider struct {

}

/**
*/
func  DataProvider() []interface{} {

    return []interface{}{
        []interface{}{
            //输入
            []byte(`{"Tns":"172.30.0.6","User":"darkhold","Password":"6532b3c13C1491FB","Port":3306,"DB":"allinone"}`),
            // 预期结论
            true,
            //执行测试，附加参数
            map[string]interface{}{"sql":"insert into a (a1) values (?)","binds":[]interface{}{"new a1"}},
            //测试案例-唯一id
            " - id=143",
        },
	}

}