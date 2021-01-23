package ShellLinkString_273_0_test

type  ShellLinkString_273_0_testProvider struct {

}

/**
用于命令行的字符串*/
func  DataProvider() []interface{} {

    return []interface{}{
        []interface{}{
            //输入
            []byte(`{"Tns":"172.30.0.6","User":"darkhold","Password":"6532b3c13C1491FB","Port":3306,"DB":"allinone"}`),
            // 预期结论
            "-h172.30.0.6 -P3306 -udarkhold -p6532b3c13C1491FB --default-character-set=utf8mb4 allinone",
            //执行测试，附加参数
            "",
            //测试案例-唯一id
            " - id=140",
        },
	}

}