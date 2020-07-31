package Exec_250_0_test
//需要导入外部类库
import (
    "strconv"
    "time"
)

type  Exec_250_0_testProvider struct {

}

func  DataProvider() []interface{} {

    return []interface{}{
        []interface{}{
            []byte(`{"Tns":"172.30.0.6","User":"darkhold","Password":"6532b3c13C1491FB","Port":3306,"DB":"allinone"}`),
            1,
            map[string]interface{}{"sql":"insert into a (a1,dd) values(?,?)","binds":[]interface{}{"cc:"+time.Now().Format("2006-01-02 15:04:05"),"ok2" }},
            " - id=103",
        },
        []interface{}{
            []byte(`{"Tns":"172.30.0.6","User":"darkhold","Password":"6532b3c13C1491FB","Port":3306,"DB":"allinone"}`),
            3,
            map[string]interface{}{"sql":"update a set dd=? where id<7","binds":[]interface{}{"ok3" +strconv.FormatInt(time.Now().Unix(),10) }},
            " - id=104",
        },
	}

}