package SqldbPool_249_0_test

type  SqldbPool_249_0_testProvider struct {

}

/**
测试数据库连接池方法，返回类型正确*/
func  DataProvider() []interface{} {

    return []interface{}{
        []interface{}{
            //输入
            []byte(`{"Tns":"172.30.0.6","User":"darkhold","Password":"6532b3c13C1491FB","Port":3306,"DB":"allinone"}`),
            // 预期结论
            "*sql.DB",
            //执行测试，附加参数
            "",
            //测试案例-唯一id
            " - id=102",
        },
	}

}