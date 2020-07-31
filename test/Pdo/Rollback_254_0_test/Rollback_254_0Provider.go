package Rollback_254_0_test

type  Rollback_254_0_testProvider struct {

}

func  DataProvider() []interface{} {

    return []interface{}{
        []interface{}{
            []byte(`{"Tns":"172.30.0.6","User":"darkhold","Password":"6532b3c13C1491FB","Port":3306,"DB":"allinone"}`),
            nil,
            []map[string]interface{}{{"sql":"insert into a (id,a1,dd) values(12,?,?)","binds":[]interface{}{"a1","hahaa"}},{"sql":"insert into a (id,aa,dd) values(13,?,?)","binds":[]interface{}{"a1","hahaa"}},{"sql":"insert into a (id,a1,dd) values(14,?,?)","binds":[]interface{}{"a1","hahaa"}}},
            "在执行更新操作的时候，有一条语句失败，导致全部回滚 - id=112",
        },
	}

}