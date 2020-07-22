package insert_251_0_test

type  Insert_251_0_testProvider struct {

}

func  DataProvider() []interface{} {

    return []interface{}{
        []interface{}{
            []byte(`{"Tns":"172.30.0.6","User":"darkhold","Password":"6532b3c13C1491FB","Port":3306,"DB":"allinone"}`),
            10,
            map[string]interface{}{"sql":"insert into a (id,a1,dd) values(10,?,?)","binds":[]interface{}{"a1","hahaa"}},
            " - id=105",
        },
        []interface{}{
            []byte(`{"Tns":"172.30.0.6","User":"darkhold","Password":"6532b3c13C1491FB","Port":3306,"DB":"allinone"}`),
            11,
            map[string]interface{}{"sql":"insert into a (id,a1,dd) values(11,?,?)","binds":[]interface{}{"a1","hahaa"}},
            " - id=109",
        },
	}

}