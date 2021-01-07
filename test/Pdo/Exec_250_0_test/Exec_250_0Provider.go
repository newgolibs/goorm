package Exec_250_0_test
//需要导入外部类库
import (
    "strconv"
    "time"
)

type  Exec_250_0_testProvider struct {

}

/**
直接执行sql语句*/
func  DataProvider() []interface{} {

    return []interface{}{
        []interface{}{
            //输入
            []byte(`{"Tns":"172.30.0.6","User":"darkhold","Password":"6532b3c13C1491FB","Port":3306,"DB":"allinone"}`),
            // 预期结论
            int64(0),
            //执行测试，附加参数
            map[string]interface{}{"sql":"delete from a  where id=? ","binds":[]interface{}{"1" }},
            //测试案例-唯一id
            " - id=137",
        },
        []interface{}{
            //输入
            []byte(`{"Tns":"172.30.0.6","User":"darkhold","Password":"6532b3c13C1491FB","Port":3306,"DB":"allinone"}`),
            // 预期结论
            int64(0),
            //执行测试，附加参数
            map[string]interface{}{"sql":"delete from a1  where id=? ","binds":[]interface{}{1 }},
            //测试案例-唯一id
            " - id=138",
        },
        []interface{}{
            //输入
            []byte(`{"Tns":"172.30.0.6","User":"darkhold","Password":"6532b3c13C1491FB","Port":3306,"DB":"allinone"}`),
            // 预期结论
            int64(1),
            //执行测试，附加参数
            map[string]interface{}{"sql":"insert into a (id,a1,dd) values(1,?,?)","binds":[]interface{}{"cc:"+time.Now().Format("2006-01-02 15:04:05"),"ok2" }},
            //测试案例-唯一id
            " - id=103",
        },
        []interface{}{
            //输入
            []byte(`{"Tns":"172.30.0.6","User":"darkhold","Password":"6532b3c13C1491FB","Port":3306,"DB":"allinone"}`),
            // 预期结论
            int64(1),
            //执行测试，附加参数
            map[string]interface{}{"sql":"update a set dd=? where id=1","binds":[]interface{}{"ok3" +strconv.FormatInt(time.Now().Unix(),10) }},
            //测试案例-唯一id
            " - id=104",
        },
	}

}