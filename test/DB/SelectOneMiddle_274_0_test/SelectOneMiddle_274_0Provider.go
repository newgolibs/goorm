package SelectOneMiddle_274_0_test

type  SelectOneMiddle_274_0_testProvider struct {

}

/**
查询一行数据，返回map[string]interface{} 结构的一维map*/
func  DataProvider() []interface{} {

    return []interface{}{
        []interface{}{
            //输入
            []byte(`{"Tns":"172.30.0.6","User":"darkhold","Password":"6532b3c13C1491FB","Port":3306,"DB":"allinone"}`),
            // 预期结论
            "{\"a1\":\"a1\",\"dd\":\"hahaa\",\"id\":\"10\"}",
            //执行测试，附加参数
            map[string]interface{}{"sql":"select id,a1,dd from a where id=?","binds":[]interface{}{1}},
            //测试案例-唯一id
            " - id=141",
        },
	}

}